package graph

import "restar/pkg/diagnostic"

//go:generate go run github.com/99designs/gqlgen generate

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

func NewResolver(dUsecase *diagnostic.Usecase) *Resolver {
	return &Resolver{
		diagnostic: dUsecase,
	}
}

type Resolver struct {
	diagnostic *diagnostic.Usecase
}
