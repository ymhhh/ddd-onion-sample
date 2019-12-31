// GNU GPL v3 License
// Copyright (c) 2017 github.com:go-trellis

package domain

import "github.com/go-trellis/ddd-onion-sample/core"

// UserRepository 用户方法
type UserRepository interface {
	GetUsers() ([]core.User, error)
	GetUserByID(id string) (*core.User, error)
}
