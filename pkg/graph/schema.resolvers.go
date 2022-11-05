package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"restar/pkg/domain"
	"restar/pkg/graph/generated"
	"restar/pkg/graph/model"
)

// CreateDiagnostic is the resolver for the createDiagnostic field.
func (r *mutationResolver) CreateDiagnostic(ctx context.Context) (*model.Diagnostic, error) {
	d, err := r.diagnostic.Create(ctx, domain.Diagnostic{})
	if err != nil {
		return nil, fmt.Errorf("failed to create diagnostic: %w", err)
	}

	return &model.Diagnostic{
		ID:            d.ID,
		Version:       d.Version,
		CreatedAt:     d.CreatedAt,
		UpdatedAt:     d.UpdatedAt,
		DefinedNumber: d.DefinedNumber,
		Sku:           d.SKU,
	}, nil
}

// DiagnosticList is the resolver for the diagnosticList field.
func (r *queryResolver) DiagnosticList(ctx context.Context) ([]*model.Diagnostic, error) {
	list, err := r.diagnostic.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get diagnostic list: %w", err)
	}

	mList := make([]*model.Diagnostic, len(list))
	for i := range list {
		mList[i] = marshalDiagnostic(list[i])
	}

	return mList, nil
}

// Diagnostic is the resolver for the diagnostic field.
func (r *queryResolver) Diagnostic(ctx context.Context, id int) (*model.Diagnostic, error) {
	diag, err := r.diagnostic.Read(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get diagnostic: %w", err)
	}

	return marshalDiagnostic(diag), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type (
	mutationResolver struct{ *Resolver }
	queryResolver    struct{ *Resolver }
)

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
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
