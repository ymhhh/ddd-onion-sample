// GNU GPL v3 License
// Copyright (c) 2017 github.com:go-trellis

package classloader

import (
	"fmt"
	"reflect"
)

// ReflectiveDynamicAccess is the default implementation akka.DynamicAccess
type ReflectiveDynamicAccess struct {
	classLoader ClassLoader
}

// NewReflectiveDynamicAccess obtain the ReflectiveDynamicAccess
func NewReflectiveDynamicAccess(classLoader ClassLoader) DynamicAccess {

	da := &ReflectiveDynamicAccess{
		classLoader: Default,
	}
	if classLoader != nil {
		da.classLoader = classLoader
	}
	return da
}

// CreateInstanceByType Convenience method which given a Class[_] object and a constructor description
// will create a new instance of that class.
func (p *ReflectiveDynamicAccess) CreateInstanceByType(typ reflect.Type, args ...interface{}) (
	ins interface{}, err error) {
	value := reflect.New(typ)
	if !value.IsValid() {
		return nil, ErrFailedCreateInstance
	}

	if err = p.constructInstance(value, args...); err != nil {
		return
	}

	ins = value.Interface()

	return
}

// CreateInstanceByName Obtain an object conforming to the type T, which is expected to be instantiated from a class designated by the fully-qualified class name given, where the constructor is selected and invoked according to the args argument.
func (p *ReflectiveDynamicAccess) CreateInstanceByName(name string, args ...interface{}) (
	ins interface{}, err error) {
	typ, exist := p.classLoader.FindClass(name)
	if !exist {
		return nil, fmt.Errorf("[ErrNotFoundInClassLoader]: %s", name)
	}

	return p.CreateInstanceByType(typ, args...)
}

func (p *ReflectiveDynamicAccess) constructInstance(val reflect.Value, args ...interface{}) (err error) {

	value := val.MethodByName("Construct")

	if !value.IsValid() {
		return ErrNotFoundConstructMethod
	}

	if numOut := value.Type().NumOut(); numOut > 1 ||
		numOut == 1 && value.Type().Out(0) != ErrorType {
		return ErrBadActorInitFuncOutNumber
	}

	var valArgs []reflect.Value
	for _, arg := range args {
		valArgs = append(valArgs, reflect.ValueOf(arg))
	}

	fnValues := value.Call(valArgs)

	if len(fnValues) > 0 &&
		fnValues[0].IsValid() &&
		!fnValues[0].IsNil() {
		err = fnValues[0].Interface().(error)
	}
	return
}

// GetClassFor Obtain a Class[_] object loaded with the right class loader (i.e. the one returned by classLoader).
func (p *ReflectiveDynamicAccess) GetClassFor(name string) (reflect.Type, bool) {
	return p.classLoader.FindClass(name)
}
