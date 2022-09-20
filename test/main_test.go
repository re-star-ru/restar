package test

import (
	"context"
	"github.com/avast/retry-go/v4"
	"github.com/jackc/pgx/v4"
	"github.com/stretchr/testify/require"
	"testing"
)

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

func Test(t *testing.T) {
	ctx := context.Background()
	conn, err := setupPG(ctx)

	require.NoError(t, err)
	require.NotNil(t, conn)
}
