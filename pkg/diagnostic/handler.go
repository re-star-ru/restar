package diagnostic

import (
	"context"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"restar/pkg/diagnostic/pb"
	"restar/pkg/domain"
)

type GRPCHandler struct {
	pb.UnimplementedDiagnosticServiceServer
	usecase *Usecase
}

func RegisterService(srv grpc.ServiceRegistrar, usecase *Usecase) {
	pb.RegisterDiagnosticServiceServer(srv, &GRPCHandler{usecase: usecase})
}

func (ds GRPCHandler) Create(ctx context.Context, _ *emptypb.Empty) (*pb.Diagnostic, error) {
	diag, err := ds.usecase.Create(ctx, domain.Diagnostic{})
	if err != nil {
		log.Err(err).Msg("cant create diagnostic")
		return nil, err
	}

	return &pb.Diagnostic{
		Id:        diag.ID,
		CreatedAt: timestamppb.New(diag.CreatedAt),
		UpdatedAt: timestamppb.New(diag.CreatedAt),
	}, nil
}

func (ds GRPCHandler) Update(ctx context.Context, diagnostic *pb.Diagnostic) (*pb.Diagnostic, error) {
	//TODO implement me
	panic("implement me")
}

func (ds GRPCHandler) List(ctx context.Context, _ *emptypb.Empty) (*pb.DiagnosticList, error) {
	//TODO implement me
	panic("implement me")
}
