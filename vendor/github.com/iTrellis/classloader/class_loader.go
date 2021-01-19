// GNU GPL v3 License
// Copyright (c) 2017 github.com:go-trellis

package classloader

import (
	"fmt"
	"reflect"
	"sync"
)

// ClassLoader class loader funtions
type ClassLoader interface {
	GetParent() ClassLoader
	LoadClass(name string, v interface{})
	FindClass(name interface{}) (reflect.Type, bool)
	FindLoadedClass(name string) bool
}

// Default default class loader
var Default = NewClassLoader(nil)

type defaultClassLoader struct {
	parent    ClassLoader
	nameTypes map[string]reflect.Type
	pathTypes map[string]reflect.Type

	locker sync.RWMutex
}

// NewClassLoader generate class loader with giving parent class loader
func NewClassLoader(parent ClassLoader) ClassLoader {
	return &defaultClassLoader{
		parent:    parent,
		nameTypes: make(map[string]reflect.Type),
		pathTypes: make(map[string]reflect.Type),
	}
}

func (p *defaultClassLoader) FindLoadedClass(name string) (exist bool) {
	_, exist = p.FindClass(name)
	return
}

func (p *defaultClassLoader) GetParent() (loader ClassLoader) {
	p.locker.RLock()
	loader = p.parent
	p.locker.RUnlock()
	return
}

func (p *defaultClassLoader) LoadClass(name string, v interface{}) {

	vType := reflect.TypeOf(v)

	if name == "" {
		name = genKey(vType.Elem())
	}

	if p.FindLoadedClass(name) {
		return
	}

	p.locker.Lock()

	switch vType.Kind() {
	case reflect.Ptr, reflect.Array, reflect.Chan, reflect.Map, reflect.Slice:
		{
			// It panics if the type's Kind is not Array, Chan, Map, Ptr, or Slice.
		}
	default:
		// unsupported class' type
		p.locker.Unlock()
		return
	}

	if _, exist := p.FindClass(v); !exist {
		p.pathTypes[genKey(vType.Elem())] = vType.Elem()
	}

	if _, exist := p.FindClass(name); !exist {
		p.nameTypes[name] = vType.Elem()
	}

	p.locker.Unlock()
}

func (p *defaultClassLoader) FindClass(name interface{}) (typ reflect.Type, exist bool) {

	vType := reflect.TypeOf(name)

	switch vType.Kind() {
	case reflect.String:
		{
			typ, exist = p.nameTypes[name.(string)]
		}
	case reflect.Ptr, reflect.Array, reflect.Chan, reflect.Map, reflect.Slice:
		{
			typ, exist = p.pathTypes[genKey(vType.Elem())]
		}
	}
	if exist {
		return
	} else if p.parent != nil {
		return p.parent.FindClass(name)
	}
	return
}

func genKey(v reflect.Type) string {
	return fmt.Sprintf("%s.%s", v.PkgPath(), v.Name())
}
