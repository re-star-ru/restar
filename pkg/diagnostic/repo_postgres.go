package diagnostic

import (
	"context"
	"fmt"

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
	tx, err := p.db.Begin(ctx)
	if err != nil {
		return domain.Diagnostic{}, fmt.Errorf("failed to begint tx: %w", err)
	}

	defer tx.Rollback(ctx)

	resp := domain.Diagnostic{}
	if err = tx.QueryRow(ctx, `
		insert into diagnostic("version", created_at, updated_at, defined_number, sku) 
		values (1, timestamptz(current_timestamp), timestamptz(current_timestamp), $1, $2) 
		returning id, "version", created_at, updated_at`,
		diag.DefinedNumber, diag.SKU,
	).Scan(&resp.ID, &resp.Version, &resp.CreatedAt, &resp.UpdatedAt); err != nil {
		return domain.Diagnostic{}, fmt.Errorf("cant scan id %w", err)
	}

	if err = tx.Commit(ctx); err != nil {
		return domain.Diagnostic{}, fmt.Errorf("cant commit new dignostic %w", err)
	}

	return resp, nil
}

func (p PostgresRepo) Read(ctx context.Context, id int) (domain.Diagnostic, error) {
	diag := domain.Diagnostic{ID: id}

	err := p.db.QueryRow(ctx, `
		select version, created_at, updated_at, defined_number, sku
		from diagnostic_view
		where id = $1`, id,
	).Scan(&diag.Version, &diag.CreatedAt, &diag.UpdatedAt, &diag.DefinedNumber, &diag.SKU)
	if err != nil {
		return domain.Diagnostic{}, fmt.Errorf("faield to query: %w", err)
	}

	return diag, nil
}

func (p PostgresRepo) Update(ctx context.Context, diag *domain.Diagnostic) error {
	tx, err := p.db.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin tx: %w", err)
	}

	defer tx.Rollback(ctx)

	err = tx.QueryRow(ctx, `
		insert into diagnostic(id, "version", created_at, updated_at, defined_number, sku)
		select id, "version"+1, created_at, timestamptz(current_timestamp), $2, $3
		from diagnostic_view 
		where id=$1
		returning "version", updated_at`,
		diag.ID, diag.DefinedNumber, diag.SKU,
	).Scan(&diag.Version, &diag.UpdatedAt)

	if err != nil {
		return fmt.Errorf("failed to update: %w", err)
	}

	if err = tx.Commit(ctx); err != nil {
		return fmt.Errorf("failed to commin update: %w", err)
	}

	return nil
}

func (p PostgresRepo) List(ctx context.Context) ([]domain.Diagnostic, error) {
	rows, err := p.db.Query(ctx, `
		select id, "version", created_at, updated_at, sku, defined_number
		from diagnostic_view 
		limit 50
		offset 0`)
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
