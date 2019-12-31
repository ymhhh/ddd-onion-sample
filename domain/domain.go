// GNU GPL v3 License
// Copyright (c) 2017 github.com:go-trellis

package domain

import "github.com/go-trellis/classloader"

// ClassLoader 类加载器
var ClassLoader = classloader.NewClassLoader(classloader.Default)

func init() {
	ClassLoader.LoadClass("user_repository_mock", NewUserMockRepo())
	// added other repositories like user_repository_mock
}

// BaseRepository 初始化基础类
type BaseRepository interface {
	Init(map[string]interface{}) error
}
