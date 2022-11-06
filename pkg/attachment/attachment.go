package attachment

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"

	"github.com/rs/zerolog/log"

	"restar/configs"
	"restar/pkg/domain"
)

type Repo struct {
	client *http.Client
	cfg    configs.Config
}

func NewAttachment(cfg configs.Config) *Repo {
	return &Repo{client: &http.Client{}, cfg: cfg}
}

func (r *Repo) PutFile(ctx context.Context, upload domain.Upload) {
	log.Debug().Msgf("upload: %v, %v, %v, %v", upload.File, upload.Filename, upload.ContentType, upload.Size)

	//ffff, err := io.ReadAll(upload.File)
	//if err != nil {
	//	log.Err(err).Msg("caint read all file")
	//}
	//
	//log.Debug().Msgf("FILE: size %v", len(ffff))

	resp, status, err := r.uploadd(ctx, upload.File, upload.Filename)

	log.Debug().Err(err).Msgf("status: %s, code: %d", string(resp), status)
}

func (r *Repo) uploadd(
	ctx context.Context, filereader io.Reader, filename string) (
	respBody []byte, statusCode int, err error,
) {
	rp, w := io.Pipe()
	mw := multipart.NewWriter(w)
	result := writePart(mw, w, filereader, filename)

	var resp *http.Response

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, fmt.Sprintf("%s/%s", r.cfg.StorePath, filename), rp)
	if err != nil {
		log.Panic().Err(err).Msg("failed to create new request with context")
	}

	req.Header.Set("Content-Type", mw.FormDataContentType())
	req.Header.Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, filename))

	resp, err = r.client.Do(req)
	_ = rp.Close()

	if err == nil {
		if respBody, statusCode, err = readAll(resp); err != nil {
			err = <-result
		}
	}

	return
}

func writePart(mw *multipart.Writer, w *io.PipeWriter, fileReader io.Reader, filename string) <-chan error {
	result := make(chan error, 1)

	go func() {
		h := make(textproto.MIMEHeader)
		h.Set("Content-Disposition", fmt.Sprintf(`form-data; name="file"; filename="%s"`, filename))
		h.Set("Content-Type", "image/jpeg")

		part, err := mw.CreatePart(h)
		if err == nil {
			_, err = io.Copy(part, fileReader)
		}

		if err == nil {
			if err = mw.Close(); err != nil {
				err = w.Close()
			} else {
				_ = w.Close()
			}
		} else {
			_ = mw.Close()
			_ = w.Close()
		}

		result <- err
	}()

	return result
}

func readAll(r *http.Response) (body []byte, statusCode int, err error) {
	statusCode = r.StatusCode
	body, err = io.ReadAll(r.Body)
	r.Body.Close()

	return
}
