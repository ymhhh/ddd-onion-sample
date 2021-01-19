// GNU GPL v3 License
// Copyright (c) 2017 github.com:iTrellis

package infrastructure

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/iTrellis/ddd-onion-sample/domain"
	"github.com/iTrellis/ddd-onion-sample/handlers"
)

var mapHandlers = make(map[string]handlers.HFunc)

// MainEntry 主入口函数
func MainEntry() {
	InitFramework()
}

// InitFramework 针对Handler对使用框架进行初始化
// 参数可能包含可能拥有的所有对象，包含数据处理，如：redis、mysql等
func InitFramework() {
	//可能是一些已经初始化的对象，比如数据库等
	params := map[string]interface{}{}
	// 可以通过配置文件告知要加载哪些类，也可以在这里写死，当做类加载
	repoNames := []string{"user_repository_mock"}
	for _, v := range repoNames {
		value, ok := domain.ClassLoader.FindClass(v)
		if !ok {
			panic(fmt.Errorf("repo not found: %s", v))
		}
		repoI := reflect.New(value)
		if !repoI.IsValid() {
			panic(fmt.Errorf("repo is invalid: %+v", repoI))
		}
		repo, ok := repoI.Interface().(domain.BaseRepository)
		if !ok {
			panic(fmt.Errorf("repo has no BaseRepository: %+v", reflect.TypeOf(repoI.Interface())))
		}
		if err := repo.Init(params); err != nil {
			panic(err)
		}
		params[strings.TrimSuffix(v, "_mock")] = repo
	}
	handlerNames := []string{"user_handler"}
	for _, v := range handlerNames {
		value, ok := handlers.ClassLoader.FindClass(v)
		if !ok {
			panic(fmt.Errorf("handler not found: %s", v))
		}
		handlerI := reflect.New(value)
		if !handlerI.IsValid() {
			panic(fmt.Errorf("handler is invalid: %+v", handlerI))
		}
		handler, ok := handlerI.Interface().(handlers.BaseHandler)
		if !ok {
			panic(fmt.Errorf("handler has no BaseHandler: %+v", reflect.TypeOf(handlerI.Interface())))
		}
		handler, err := handler.Init(params)
		if err != nil {
			panic(err)
		}
		for _, m := range handler.SupportMethod() {
			mapHandlers[handler.AccessPath()+m] = handler.Executor
		}
	}
}
