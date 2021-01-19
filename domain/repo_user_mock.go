// GNU GPL v3 License
// Copyright (c) 2017 github.com:iTrellis

package domain

import (
	"fmt"

	"github.com/iTrellis/ddd-onion-sample/core"
)

// 模拟用户数据
var mockUser = map[string]*core.User{
	"1": &core.User{ID: "1", Name: "John"},
	"2": &core.User{ID: "2", Name: "Toni"},
}

// UserMockRepo 用户信息mock
type UserMockRepo struct {
	BaseRepository
}

// NewUserMockRepo 获取用户操作对象
func NewUserMockRepo() UserRepository {
	return &UserMockRepo{}
}

// Init 初始化
func (*UserMockRepo) Init(map[string]interface{}) error {
	fmt.Println("user repo to do something with params")
	return nil
}

// GetUsers 获取所有用户
func (*UserMockRepo) GetUsers() ([]core.User, error) {
	var users []core.User
	for _, v := range mockUser {
		item := *v
		users = append(users, item)
	}
	return users, nil
}

// GetUserByID 获取指定用户
func (*UserMockRepo) GetUserByID(id string) (*core.User, error) {
	user, ok := mockUser[id]
	if !ok {
		return nil, fmt.Errorf("unkown user, id: %s", id)
	}
	return user, nil
}
