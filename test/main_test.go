package test

import (
	"context"
	"github.com/avast/retry-go/v4"
	"github.com/jackc/pgx/v4"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
	"net"
	"restar/pkg/diagnostic"
	"restar/pkg/diagnostic/pb"
	"testing"
)

const testConn = "127.0.0.1:40444"

func setupPG(ctx context.Context) (conn *pgx.Conn, err error) {
	retryErr := retry.Do(
		func() error {
			conn, err = pgx.Connect(ctx, "postgresql://test:test@localhost:5433/test")
			if err != nil {
				return err
			}

			if err = conn.Ping(ctx); err != nil {
				return err
			}

			return nil
		},
	)

	if retryErr != nil {
		return nil, retryErr
	}

	return conn, nil
}

func setupServer() (*grpc.Server, error) {
	ctx := context.Background()
	conn, err := setupPG(ctx)
	if err != nil {
		return nil, err
	}

	listen, err := net.Listen("tcp", testConn)
	if err != nil {
		return nil, err
	}

	srv := grpc.NewServer()

	drepo := diagnostic.NewPostgresRepo(conn)
	ducase := diagnostic.NewUsecase(drepo)
	diagnostic.RegisterService(srv, ducase)

	go func() {
		if err = srv.Serve(listen); err != nil {
			log.Fatal().Err(err)
		}
	}()

	return srv, nil
}

func setupClient() (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(testConn, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return conn, nil
}

var diagClient pb.DiagnosticServiceClient

func TestGRPC(t *testing.T) {
	srv, err := setupServer()
	require.NoError(t, err)
	require.NotNil(t, srv)

	conn, err := setupClient()
	require.NoError(t, err)
	require.NotNil(t, conn)

	diagClient = pb.NewDiagnosticServiceClient(conn)

	ctx := context.Background()

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			tc.test(ctx, t)
		})
	}
}

type testCase struct {
	name string
	test func(ctx context.Context, t *testing.T)
}

var tcs = []testCase{
	{
		name: "first document",
		test: func(ctx context.Context, t *testing.T) {
			create, err := diagClient.Create(ctx, &emptypb.Empty{})
			require.NoError(t, err)
			require.Equal(t, 1, create.Id)
		},
	},
	{
		name: "second document",
		test: func(ctx context.Context, t *testing.T) {
			create, err := diagClient.Create(ctx, &emptypb.Empty{})
			require.NoError(t, err)
			require.Equal(t, 2, create.Id)
		},
	},
	{
		name: "update first document",
		test: func(ctx context.Context, t *testing.T) {
			_, err := diagClient.Update(ctx, &pb.Diagnostic{
				Version:       1,
				DefinedNumber: "defined",
				SKU:           "updated",
			})

			require.NoError(t, err)
		},
	},
}