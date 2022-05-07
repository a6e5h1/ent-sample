package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/a6e5h1/ent-sample/pkg/api/interfaces/graphql/generated"
	"github.com/a6e5h1/ent-sample/pkg/api/interfaces/graphql/model"
)

func (r *queryResolver) Post(ctx context.Context, id string) (*model.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

// Post returns generated.PostResolver implementation.
func (r *Resolver) Post() generated.PostResolver { return &postResolver{r} }

type postResolver struct{ *Resolver }
