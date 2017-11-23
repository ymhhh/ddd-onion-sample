package domain

import "github.com/go-trellis/ddd-onion-sample/core"

// UserRepository 用户方法
type UserRepository interface {
	GetUsers() ([]core.User, error)
	GetUserByID(id string) (*core.User, error)
}
