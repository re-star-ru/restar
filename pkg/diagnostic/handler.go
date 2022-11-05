package diagnostic

import (
	"google.golang.org/grpc"

	"restar/pkg/diagnostic/pb"
)

type GRPCHandler struct {
	pb.UnimplementedDiagnosticServiceServer
	usecase *Usecase
}

func RegisterService(srv grpc.ServiceRegistrar, usecase *Usecase) {
	pb.RegisterDiagnosticServiceServer(srv, NewGRPCHandler(usecase))
}

func NewGRPCHandler(usecase *Usecase) *GRPCHandler {
	return &GRPCHandler{
		usecase: usecase,
	}
}

//func (ds GRPCHandler) Create(ctx context.Context, _ *emptypb.Empty) (*pb.Diagnostic, error) {
//	diag, err := ds.usecase.Create(ctx, domain.Diagnostic{})
//	if err != nil {
//		log.Err(err).Msg("cant create diagnostic")
//
//		return nil, err
//	}
//
//	return &pb.Diagnostic{
//		Id:        diag.ID,
//		Version:   diag.Version,
//		CreatedAt: timestamppb.New(diag.CreatedAt),
//		UpdatedAt: timestamppb.New(diag.UpdatedAt),
//	}, nil
//}
//
//func (ds GRPCHandler) Update(ctx context.Context, diagnostic *pb.Diagnostic) (*pb.Diagnostic, error) {
//	diag := &domain.Diagnostic{
//		ID:            diagnostic.Id,
//		Version:       diagnostic.Version,
//		CreatedAt:     diagnostic.CreatedAt.AsTime(),
//		UpdatedAt:     diagnostic.UpdatedAt.AsTime(),
//		DefinedNumber: diagnostic.DefinedNumber,
//		SKU:           diagnostic.Sku,
//	}
//
//	if err := ds.usecase.Update(ctx, diag); err != nil {
//		return nil, err
//	}
//
//	return marshalDiagnostic(*diag), nil
//}
//
//func (ds GRPCHandler) List(ctx context.Context, _ *emptypb.Empty) (*pb.DiagnosticList, error) {
//	list, err := ds.usecase.List(ctx)
//	if err != nil {
//		return nil, fmt.Errorf("cant get list of diagnostic %w", err)
//	}
//
//	pbList := make([]*pb.Diagnostic, len(list))
//	for i, v := range list {
//		pbList[i] = marshalDiagnostic(v)
//	}
//
//	return &pb.DiagnosticList{Lists: pbList}, nil
//}
//
//func (ds GRPCHandler) Read(ctx context.Context, id *pb.ID) (*pb.Diagnostic, error) {
//	diag, err := ds.usecase.Read(ctx, id.Id)
//	if err != nil {
//		return nil, fmt.Errorf("failed to read document, %w", err)
//	}
//
//	return marshalDiagnostic(diag), nil
//}
//
//func marshalDiagnostic(diag domain.Diagnostic) *pb.Diagnostic {
//	return &pb.Diagnostic{
//		Id:            diag.ID,
//		Version:       diag.Version,
//		CreatedAt:     timestamppb.New(diag.CreatedAt),
//		UpdatedAt:     timestamppb.New(diag.UpdatedAt),
//		DefinedNumber: diag.DefinedNumber,
//		Sku:           diag.SKU,
//	}
//}
