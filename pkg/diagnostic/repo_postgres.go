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

	var id uint64
	if err = tx.QueryRow(ctx, `
		insert into diagnostic(defined_number, sku, created_at, updated_at) 
		values ($1, $2, timestamptz(current_timestamp), timestamptz(current_timestamp)) 
		returning id`,
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
		ID:            id,
		DefinedNumber: diag.DefinedNumber,
		SKU:           diag.SKU,
		Images:        diag.Images,
	}, nil
}

func (p PostgresRepo) Update(ctx context.Context, diag *domain.Diagnostic) error {
	return nil
}

func (p PostgresRepo) List(ctx context.Context) ([]domain.Diagnostic, error) {
	rows, err := p.db.Query(ctx, `
		select id, "version", created_at, updated_at, sku, defined_number
		from diagnostic_view 
		limit 50`)
	if err != nil {
		return nil, fmt.Errorf("cant query list diagnostic %w", err)
	}
	defer rows.Close()

	diagRows := make([]domain.Diagnostic, 0, 50)
	for rows.Next() {
		var diag domain.Diagnostic
		if err = rows.Scan(
			&diag.ID, &diag.Version, &diag.CreatedAt, &diag.UpdatedAt,
			&diag.SKU, &diag.DefinedNumber,
		); err != nil {
			return nil, fmt.Errorf("cant scan row %w", err)
		}

		diagRows = append(diagRows, diag)
	}

	return diagRows, nil
}
