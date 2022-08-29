package diagnostic

import (
	"context"
	"restar/pkg/domain"
)

type Usecase struct {
}

func NewUsecase() *Usecase {
	return &Usecase{}
}

func (uc *Usecase) Create(ctx context.Context) (domain.Diagnostic, error) {
	return domain.Diagnostic{}, nil
}

func (uc *Usecase) Update(ctx context.Context, diag *domain.Diagnostic) error {
	return nil
}

func (uc *Usecase) Read(ctx context.Context, id int64) (domain.Diagnostic, error) {
	return domain.Diagnostic{}, nil
}

func (uc *Usecase) List(ctx context.Context) ([]domain.Diagnostic, error) {
	return nil, nil
}
