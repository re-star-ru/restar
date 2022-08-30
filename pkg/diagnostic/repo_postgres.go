package diagnostic

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/rs/zerolog/log"
	"restar/pkg/domain"
)

type PostgresRepo struct {
	db *pgx.Conn
}

func NewPostgresRepo(conn *pgx.Conn) *PostgresRepo {
	return &PostgresRepo{conn}
}

func (p PostgresRepo) Create(ctx context.Context, diag domain.Diagnostic) (domain.Diagnostic, error) {
	tx, err := p.db.Begin(ctx)
	if err != nil {
		return domain.Diagnostic{}, err
	}
	defer tx.Rollback(ctx)

	id := 0
	if err := tx.QueryRow(ctx,
		`insert into diagnostic(definednumber, sku) values ($1, $2) returning id`,
		diag.DefinedNumber, diag.SKU,
	).Scan(&id); err != nil {
		return domain.Diagnostic{}, fmt.Errorf("cant scan id %w", err)
	}

	irows := make([][]interface{}, len(diag.Images))
	for i, v := range diag.Images {
		irows[i] = []interface{}{v.Path, v.Name}
	}

	ins, err := tx.CopyFrom(ctx, pgx.Identifier{"image"}, []string{"name", "path"}, pgx.CopyFromRows(irows))
	if err != nil {
		return domain.Diagnostic{}, err
	}

	if err = tx.Commit(ctx); err != nil {
		return domain.Diagnostic{}, err
	}
	log.Debug().Msgf("inserted rows: %v", ins)

	return domain.Diagnostic{
		ID:            int64(id),
		DefinedNumber: diag.DefinedNumber,
		SKU:           diag.SKU,
		Images:        diag.Images,
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
