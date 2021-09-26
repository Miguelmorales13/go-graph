package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/Miguelmorales13/go-graph/graph/generated"
	"github.com/Miguelmorales13/go-graph/graph/model"
)

var todos  = []*model.Todo{}


func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	return  *model.Todo{ID    :"1", Text  :input.Text, Done : true},nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return todos
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
