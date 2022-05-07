package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/a6e5h1/ent-sample/pkg/api/interfaces/graphql/model"
)

func (r *commentResolver) Commenter(ctx context.Context, obj *model.Comment) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *postResolver) Poster(ctx context.Context, obj *model.Post) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}
