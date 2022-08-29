package diagnostic

import (
	"context"
	"github.com/jackc/pgx/v4"
	"restar/pkg/domain"
)

type PostgresRepo struct {
	db *pgx.Conn
}

func NewPostgresRepo(conn *pgx.Conn) *PostgresRepo {
	return &PostgresRepo{conn}
}

func (p PostgresRepo) Create(ctx context.Context, diag domain.Diagnostic) (domain.Diagnostic, error) {
	return domain.Diagnostic{
		ID:            1,
		DefinedNumber: "23",
		SKU:           "23",
		Images:        nil,
	}, nil
}

func (p PostgresRepo) Update(ctx context.Context, diag *domain.Diagnostic) error {
	return nil
}

func (p PostgresRepo) List(ctx context.Context) ([]domain.Diagnostic, error) {
	return []domain.Diagnostic{{
		ID:            1,
		DefinedNumber: "32",
		SKU:           "32",
		Images:        nil,
	}}, nil
}
