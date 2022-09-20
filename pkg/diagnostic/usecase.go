package diagnostic

import (
	"context"
	"fmt"

	"restar/pkg/domain"
)

type Repo interface {
	Create(context.Context, domain.Diagnostic) (domain.Diagnostic, error)
	Update(context.Context, *domain.Diagnostic) error
	List(context.Context) ([]domain.Diagnostic, error)
}

type Usecase struct {
	repo Repo
}

func NewUsecase(repo Repo) *Usecase {
	return &Usecase{repo}
}

func (uc *Usecase) Create(ctx context.Context, diag domain.Diagnostic) (domain.Diagnostic, error) {
	d, err := uc.repo.Create(ctx, diag)
	if err != nil {
		return domain.Diagnostic{}, fmt.Errorf("failed to create diagnostic :%w", err)
	}

	return d, nil
}

func (uc *Usecase) Update(ctx context.Context, diag *domain.Diagnostic) error {
	err := uc.repo.Update(ctx, diag)
	if err != nil {
		return fmt.Errorf("failed to update diag, %w, %+v", err, diag)
	}

	return nil
}

func (uc *Usecase) Delete(ctx context.Context, id int64) error {
	return nil
}

func (uc *Usecase) Read(ctx context.Context, id int64) (domain.Diagnostic, error) {
	return domain.Diagnostic{}, nil
}

func (uc *Usecase) List(ctx context.Context) ([]domain.Diagnostic, error) {
	list, err := uc.repo.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to %w", err)
	}

	return list, nil
}
