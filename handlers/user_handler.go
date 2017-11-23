package handlers

import (
	"context"
	"fmt"
	"reflect"

	"github.com/go-trellis/ddd-onion-sample/domain"
)

// UserHandler 用户处理对象
type UserHandler struct {
	*BaseWorkers

	UserRepository domain.UserRepository
}

// Init 获取用户处理对象
func (p *UserHandler) Init(params map[string]interface{}) (BaseHandler, error) {
	p = &UserHandler{}
	p.BaseWorkers = &BaseWorkers{
		Path: "user/",
		HandlerWorkers: map[string]HFunc{
			"get": p.GetUserInfo,
		},
	}

	repo, ok := params["user_repository"].(domain.UserRepository)
	if !ok {
		return nil, fmt.Errorf("unknown user_repository, type: %v",
			reflect.TypeOf(params["user_repository"]))
	}
	p.UserRepository = repo
	return p, nil
}

// GetUserInfo 获取用户信息
func (p *UserHandler) GetUserInfo(ctx context.Context) (interface{}, error) {
	id, ok := ctx.Value("id").(string)
	if !ok {
		return nil, fmt.Errorf("id is not string: %v", id)
	}
	user, err := p.UserRepository.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
