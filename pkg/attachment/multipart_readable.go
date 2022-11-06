package attachment

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime"
	"net/http"
	"strconv"
	"strings"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/errcode"
	"github.com/rs/zerolog/log"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// ReadableMultipartForm the Multipart request spec https://github.com/jaydenseric/graphql-multipart-request-spec
type ReadableMultipartForm struct {
	// MaxUploadSize sets the maximum number of bytes used to parse a request body
	// as multipart/form-data.
	MaxUploadSize int64
}

var _ graphql.Transport = ReadableMultipartForm{}

func (f ReadableMultipartForm) Supports(r *http.Request) bool {
	if r.Header.Get("Upgrade") != "" {
		return false
	}

	mediaType, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		return false
	}

	return r.Method == "POST" && mediaType == "multipart/form-data"
}

func (f ReadableMultipartForm) maxUploadSize() int64 {
	if f.MaxUploadSize == 0 {
		return 32 << 20
	}
	return f.MaxUploadSize
}

func (f ReadableMultipartForm) Do(w http.ResponseWriter, r *http.Request, exec graphql.GraphExecutor) {
	w.Header().Set("Content-Type", "application/json")

	start := graphql.Now()

	if r.ContentLength > f.maxUploadSize() {
		writeJsonError(w, "failed to parse multipart form, request body too large")

		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, f.maxUploadSize())
	defer r.Body.Close()

	var err error

	mr, err := r.MultipartReader()
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		writeJsonError(w, "failed to parse multipart form")
		return
	}

	part, err := mr.NextPart()
	if err != nil || part.FormName() != "operations" {
		w.WriteHeader(http.StatusUnprocessableEntity)
		writeJsonError(w, "first part must be operations")

		return
	}

	var params graphql.RawParams
	if err = jsonDecode(part, &params); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		writeJsonError(w, "operations form field could not be decoded")

		return
	}

	part, err = mr.NextPart()
	if err != nil || part.FormName() != "map" {
		w.WriteHeader(http.StatusUnprocessableEntity)
		writeJsonError(w, "second part must be map")

		return
	}

	uploadsMap := map[string][]string{}
	if err = json.NewDecoder(part).Decode(&uploadsMap); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		writeJsonError(w, "map form field could not be decoded")

		return
	}

	var (
		ctx       context.Context
		responses graphql.ResponseHandler
	)

	part, err = mr.NextPart()

	if err != nil {
		if !errors.Is(err, io.EOF) {
			w.WriteHeader(http.StatusUnprocessableEntity)
			writeJsonErrorf(w, "failed to next part")

			return
		}
	}

	if err != nil {
		return
	}

	partSize := partSizeFromHeader(part.Header.Get("Content-Disposition"))
	log.Debug().Msgf("PART SIZE: %d", partSize)

	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		writeJsonErrorf(w, "failed to parse part")

		return
	}

	key := part.FormName()
	filename := part.FileName()
	contentType := part.Header.Get("Content-Type")

	paths := uploadsMap[key]
	if len(paths) == 0 {
		w.WriteHeader(http.StatusUnprocessableEntity)
		writeJsonErrorf(w, "invalid empty operations paths list for key %s", key)

		return
	}

	delete(uploadsMap, key)

	for key := range uploadsMap {
		w.WriteHeader(http.StatusUnprocessableEntity)
		writeJsonErrorf(w, "failed to get key %s from form", key)

		return
	}

	var upload graphql.Upload
	for _, path := range paths {
		upload = graphql.Upload{
			File:        &bytesReader{Reader: part, i: 0},
			Size:        partSize,
			Filename:    filename,
			ContentType: contentType,
		}

		if gerr := params.AddUpload(upload, key, path); gerr != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			writeJsonGraphqlError(w, gerr)

			return
		}
	}

	params.Headers = r.Header
	params.ReadTime = graphql.TraceTiming{
		Start: start,
		End:   graphql.Now(),
	}

	rc, gerr := exec.CreateOperationContext(r.Context(), &params)
	if gerr != nil {
		resp := exec.DispatchError(graphql.WithOperationContext(r.Context(), rc), gerr)
		w.WriteHeader(statusFor(gerr))
		writeJson(w, resp)

		return
	}

	responses, ctx = exec.DispatchOperation(r.Context(), rc)

	writeJson(w, responses(ctx))
}

/// bytesReader

type bytesReader struct {
	io.Reader
	i int64 // current reading index
}

func (r *bytesReader) Seek(_ int64, _ int) (int64, error) {
	panic("SEEK NOT NEED")
}

///// utils

func partSizeFromHeader(str string) (partSize int64) {
	strs := strings.Split(str, ";")
	for _, v := range strs {
		if strings.Contains(v, "size=") {
			sizeStr := strings.Split(v, "=")
			if len(sizeStr) != 2 {
				continue
			}

			partsz := 0

			partsz, err := strconv.Atoi(sizeStr[1])
			if err != nil {
				return 0
			}

			partSize = int64(partsz)
		}
	}

	return
}

func writeJson(w io.Writer, response *graphql.Response) {
	b, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}
	w.Write(b)
}

func writeJsonError(w io.Writer, msg string) {
	writeJson(w, &graphql.Response{Errors: gqlerror.List{{Message: msg}}})
}

func writeJsonErrorf(w io.Writer, format string, args ...interface{}) {
	writeJson(w, &graphql.Response{Errors: gqlerror.List{{Message: fmt.Sprintf(format, args...)}}})
}

func writeJsonGraphqlError(w io.Writer, err ...*gqlerror.Error) {
	writeJson(w, &graphql.Response{Errors: err})
}

func jsonDecode(r io.Reader, val interface{}) error {
	dec := json.NewDecoder(r)
	dec.UseNumber()
	return dec.Decode(val)
}

func statusFor(errs gqlerror.List) int {
	switch errcode.GetErrorKind(errs) {
	case errcode.KindProtocol:
		return http.StatusUnprocessableEntity
	default:
		return http.StatusOK
	}
}
