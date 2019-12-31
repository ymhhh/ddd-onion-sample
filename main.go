// GNU GPL v3 License
// Copyright (c) 2017 github.com:go-trellis

package main

import "github.com/go-trellis/ddd-onion-sample/infrastructure"

func main() {
	// do somethings with build

	// 初始化主流程
	infrastructure.MainEntry()

	// 测试获取用户信息
	infrastructure.MockEntryGetUser()
}
