package registry

import (
	"github.com/a6e5h1/ent-sample/pkg/core/services/comment"
	"github.com/a6e5h1/ent-sample/pkg/core/services/comment/usecase"
)

func (r *registry) newUserService() comment.Service {
	return usecase.New()
}
