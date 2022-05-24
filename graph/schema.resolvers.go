package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/romshark/gqlgen-2195-reproduction/graph/generated"
	"github.com/romshark/gqlgen-2195-reproduction/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	return &model.Todo{
		ID:   "new",
		Text: "This is a newly created todo",
		Done: false,
		User: &model.User{
			ID:   "you",
			Name: "It's You",
		},
	}, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return []*model.Todo{
		{
			ID:   "first",
			Text: "This is the first todo",
			Done: true,
			User: &model.User{
				ID:   "first_creator",
				Name: "First User",
			},
		},
		{
			ID:   "second",
			Text: "This is the second todo",
			Done: false,
			User: &model.User{
				ID:   "second_creator",
				Name: "Second User",
			},
		},
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
