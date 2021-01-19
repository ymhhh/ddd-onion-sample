// GNU GPL v3 License
// Copyright (c) 2017 github.com:go-trellis

package classloader

import "reflect"

// The DynamicAccess implementation is the class which is used for loading all configurable parts of an actor system (the ReflectiveDynamicAccess is the default implementation).
// This is an internal facility and users are not expected to encounter it unless they are extending Akka in ways which go beyond simple Extensions.
type DynamicAccess interface {
	// Convenience method which given a Class[_] object and a constructor description will create a new instance of that class.
	CreateInstanceByType(typ reflect.Type, args ...interface{}) (ins interface{}, err error)
	// Obtain an object conforming to the type T, which is expected to be instantiated from a class designated by the fully-qualified class name given, where the constructor is selected and invoked according to the args argument.
	CreateInstanceByName(name string, args ...interface{}) (ins interface{}, err error)
	// Obtain a Class[_] object loaded with the right class loader (i.e. the one returned by classLoader).
	GetClassFor(name string) (reflect.Type, bool)
}
