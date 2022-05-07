package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/a6e5h1/ent-sample/pkg/api/interfaces/graphql/generated"
	"github.com/a6e5h1/ent-sample/pkg/api/interfaces/graphql/model"
)

func (r *postResolver) Comments(ctx context.Context, obj *model.Post) ([]*model.Comment, error) {
	panic(fmt.Errorf("not implemented"))
}

// Comment returns generated.CommentResolver implementation.
func (r *Resolver) Comment() generated.CommentResolver { return &commentResolver{r} }

type commentResolver struct{ *Resolver }
