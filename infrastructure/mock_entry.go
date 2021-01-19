// GNU GPL v3 License
// Copyright (c) 2017 github.com:iTrellis

package infrastructure

import (
	"context"
	"fmt"
	"log"
)

// MockEntryGetUser 测试获取用户
func MockEntryGetUser() {

	_, ok := mapHandlers["test"]
	if !ok {
		fmt.Printf("not found test handler\n")
	}

	ctx := context.WithValue(context.Background(), "id", "1")

	userHandler, ok := mapHandlers["user/get"]
	if !ok {
		log.Fatalln("not found test handler:", "user/get")
	}
	ctx = context.WithValue(ctx, "method", "post")
	_, err := userHandler(ctx)
	if err != nil {
		fmt.Println(err)
	}

	ctx = context.WithValue(ctx, "method", "get")
	resp, err := userHandler(ctx)
	if err != nil {
		log.Fatalln("not found user John by id: 1")
	}
	fmt.Println(resp)

	ctx = context.WithValue(ctx, "id", "3")
	_, err = userHandler(ctx)
	if err != nil {
		log.Fatalln("not found user by id:3")
	}
}
