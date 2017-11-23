package domain

import "github.com/go-trellis/class_loader"

// ClassLoader 类加载器
var ClassLoader = class_loader.NewClassLoader(class_loader.Default)

func init() {
	ClassLoader.LoadClass("user_repository_mock", NewUserMockRepo())
	// added other repositories like user_repository_mock
}

// BaseRepository 初始化基础类
type BaseRepository interface {
	Init(map[string]interface{}) error
}
