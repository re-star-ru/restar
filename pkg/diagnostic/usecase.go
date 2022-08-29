package diagnostic

import (
	"context"
	"restar/pkg/domain"
)

type Repo interface {
}

type Usecase struct {
	repo Repo
}

func NewUsecase(repo Repo) *Usecase {
	return &Usecase{repo}
}

func (uc *Usecase) Create(ctx context.Context) (domain.Diagnostic, error) {
	return domain.Diagnostic{}, nil
}

func (uc *Usecase) Update(ctx context.Context, diag *domain.Diagnostic) error {
	return nil
}

func (uc *Usecase) Delete(ctx context.Context, id int64) error {
	return nil
}

func (uc *Usecase) Read(ctx context.Context, id int64) (domain.Diagnostic, error) {
	return domain.Diagnostic{}, nil
}

func (uc *Usecase) List(ctx context.Context) ([]domain.Diagnostic, error) {
	return nil, nil
}
