package registry

import (
	"github.com/a6e5h1/ent-sample/ent"
	"github.com/a6e5h1/ent-sample/pkg/core/services/comment"
	"github.com/a6e5h1/ent-sample/pkg/core/services/post"
	"github.com/a6e5h1/ent-sample/pkg/core/services/user"
	"go.uber.org/zap"
)

type registry struct {
	logger   *zap.Logger
	dbClient *ent.Client
}

type Services struct {
	Post    post.Service
	User    user.Service
	Comment comment.Service
}

func NewServices(logger *zap.Logger, dbClient *ent.Client) *Services {
	reg := &registry{
		logger:   logger,
		dbClient: dbClient,
	}

	return &Services{
		Post:    reg.newPostService(),
		User:    reg.newUserService(),
		Comment: reg.newCommentService(),
	}
}
