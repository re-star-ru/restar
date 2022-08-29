package diagnostic

import (
	"context"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"restar/pkg/diagnostic/pb"
	"restar/pkg/domain"
)

type GRPCHandler struct {
	pb.UnimplementedDiagnosticServiceServer
	usecase *Usecase
}

func RegisterService(srv grpc.ServiceRegistrar, usecase *Usecase) {
	srv.RegisterService(
		&pb.DiagnosticService_ServiceDesc,
		&GRPCHandler{
			usecase: usecase,
		},
	)
}

func (ds GRPCHandler) Create(ctx context.Context, diagnostic *pb.Diagnostic) (*pb.Diagnostic, error) {
	diag, err := ds.usecase.Create(ctx, domain.Diagnostic{
		ID:            diagnostic.Id,
		DefinedNumber: diagnostic.DefinedNumber,
		SKU:           diagnostic.Sku,
	})
	if err != nil {
		log.Err(err).Msg("cant create diagnostic")
		return nil, err
	}

	return &pb.Diagnostic{
		Id:            diag.ID,
		DefinedNumber: diag.DefinedNumber,
		Sku:           diag.SKU,
		Images:        []*pb.Image{},
	}, nil
}

func (ds GRPCHandler) Update(ctx context.Context, diagnostic *pb.Diagnostic) (*pb.Diagnostic, error) {
	//TODO implement me
	panic("implement me")
}

func (ds GRPCHandler) List(ctx context.Context, empty *pb.Empty) (*pb.DiagnosticList, error) {
	//TODO implement me
	panic("implement me")
}
