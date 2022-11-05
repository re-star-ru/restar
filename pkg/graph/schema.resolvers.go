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

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUser) (*model.User, error) {
	usr, err := r.user.Create(domain.User{
		Name: input.Name,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get create user: %w", err)
	}

	return &model.User{
		ID:   usr.ID,
		Name: usr.Name,
	}, nil
}

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

// UpdateDiagnostic is the resolver for the updateDiagnostic field.
func (r *mutationResolver) UpdateDiagnostic(ctx context.Context, input model.UpdateDiagnostic) (*model.Diagnostic, error) {
	if err := r.diagnostic.Update(ctx, &domain.Diagnostic{
		ID:            input.ID,
		DefinedNumber: input.DefinedNumber,
		SKU:           input.Sku,
	}); err != nil {
		return nil, fmt.Errorf("failed to update diagnostic: %w", err)
	}

	diag, err := r.diagnostic.Read(ctx, input.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to read diagnostic: %w", err)
	}

	return marshalDiagnostic(diag), nil
}

// UserList is the resolver for the userList field.
func (r *queryResolver) UserList(ctx context.Context) (*model.User, error) {
	panic(fmt.Errorf("not implemented: UserList - userList"))
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
