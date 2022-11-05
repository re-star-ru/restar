package graph

import (
	"restar/pkg/diagnostic"
	"restar/pkg/domain"
	"restar/pkg/graph/model"
	"restar/pkg/user"
)

//go:generate go run github.com/99designs/gqlgen generate

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

func NewResolver(dUsecase *diagnostic.Usecase, userUsecase *user.Usecase) *Resolver {
	return &Resolver{
		diagnostic: dUsecase,
		user:       userUsecase,
	}
}

type Resolver struct {
	diagnostic *diagnostic.Usecase
	user       *user.Usecase
}

func marshalDiagnostic(d domain.Diagnostic) *model.Diagnostic {
	return &model.Diagnostic{
		ID:            d.ID,
		Version:       d.Version,
		CreatedAt:     d.CreatedAt,
		UpdatedAt:     d.UpdatedAt,
		DefinedNumber: d.DefinedNumber,
		Sku:           d.SKU,
	}
}
