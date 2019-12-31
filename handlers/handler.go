// GNU GPL v3 License
// Copyright (c) 2017 github.com:go-trellis

package handlers

import (
	"context"
	"fmt"
	"reflect"

	"github.com/go-trellis/classloader"
)

// HFunc 处理方法
type HFunc func(ctx context.Context) (resp interface{}, err error)

// BaseHandler 基础操作类
type BaseHandler interface {
	Init(params map[string]interface{}) (BaseHandler, error)
	AccessPath() string
	SupportMethod() []string
	Executor(ctx context.Context) (resp interface{}, err error)
}

// BaseWorkers 基础操作类
type BaseWorkers struct {
	Path string

	HandlerWorkers map[string]HFunc
}

// AccessPath 工作路由地址
func (p *BaseWorkers) AccessPath() string {
	return p.Path
}

// SupportMethod 支持的方法
func (p *BaseWorkers) SupportMethod() []string {
	supMethods := []string{}
	for method := range p.HandlerWorkers {
		supMethods = append(supMethods, method)
	}
	return supMethods
}

// Executor 处理方法
func (p *BaseWorkers) Executor(ctx context.Context) (interface{}, error) {
	method, ok := ctx.Value("method").(string)
	if !ok {
		return nil, fmt.Errorf("unknown method, type %v", reflect.TypeOf(ctx.Value("method")))
	}
	worker, ok := p.HandlerWorkers[method]
	if !ok {
		return nil, fmt.Errorf("unknown method: %v", method)
	}
	return worker(ctx)
}

// ClassLoader 类加载器
var ClassLoader = classloader.NewClassLoader(classloader.Default)

func init() {
	ClassLoader.LoadClass("user_handler", (*UserHandler)(nil))
	// added other repositories like user_handler
}
