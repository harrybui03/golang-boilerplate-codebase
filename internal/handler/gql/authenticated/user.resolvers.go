package authenticated

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"golang-boilerplate/ent"
	"golang-boilerplate/internal/model"
)

// User is the resolver for the user field.
func (r *mutationResolver) User(ctx context.Context) (*ent.UserOps, error) {
	panic(fmt.Errorf("not implemented: User - user"))
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context) (*model.UserQueries, error) {
	panic(fmt.Errorf("not implemented: User - user"))
}

// ID is the resolver for the id field.
func (r *userResolver) ID(ctx context.Context, obj *ent.User) (string, error) {
	panic(fmt.Errorf("not implemented: ID - id"))
}

// Create is the resolver for the create field.
func (r *userOpsResolver) Create(ctx context.Context, obj *ent.UserOps, input model.CreateUserInput) (*ent.User, error) {
	panic(fmt.Errorf("not implemented: Create - create"))
}

// List is the resolver for the list field.
func (r *userQueriesResolver) List(ctx context.Context, obj *model.UserQueries, filter *ent.UserFilterInput) (*ent.UserConnection, error) {
	panic(fmt.Errorf("not implemented: List - list"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// User returns UserResolver implementation.
func (r *Resolver) User() UserResolver { return &userResolver{r} }

// UserOps returns UserOpsResolver implementation.
func (r *Resolver) UserOps() UserOpsResolver { return &userOpsResolver{r} }

// UserQueries returns UserQueriesResolver implementation.
func (r *Resolver) UserQueries() UserQueriesResolver { return &userQueriesResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
type userOpsResolver struct{ *Resolver }
type userQueriesResolver struct{ *Resolver }
