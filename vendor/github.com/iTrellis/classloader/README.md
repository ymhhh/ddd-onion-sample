# class_loader
class loader in Go

## [Example](example/main.go)

```go
package main

import (
	"fmt"
	"reflect"

	"github.com/go-trellis/class_loader"
)

type Test struct{}

func (p *Test) Construct() {}

func (p Test) Hello(name string) {
	fmt.Println("hello:", name)
}

func main() {
	c := class_loader.Default

	c.LoadClass("class_loader:Test", (*Test)(nil))

	c.LoadClass("", (*Test)(nil))

	// unsupported class' type: string
	c.LoadClass("class_loader:Aha", "name")

	NameTest()

	ClassLoaderTest()

	DynamicAccess()
}

func ClassLoaderTest() {

	t, e := class_loader.Default.FindClass((*Test)(nil))
	if !e {
		fmt.Println("error:", "class_loader:Test not exist")
		return
	}

	test := reflect.New(t)

	if !test.IsValid() {
		fmt.Println("error:", "class_loader:Test is invalid")
		return
	}

	testI, _ := test.Interface().(*Test)

	testI.Hello("class loaded")

}

func NameTest() {

	t, e := class_loader.Default.FindClass("class_loader:Test")
	if !e {
		fmt.Println("error:", "name class not exist")
		return
	}

	test := reflect.New(t)

	if !test.IsValid() {
		fmt.Println("error:", "name class is invalid")
		return
	}

	testI, _ := test.Interface().(*Test)

	testI.Hello("name class loaded")
}

func DynamicAccess() {
	da := class_loader.NewReflectiveDynamicAccess(class_loader.Default)

	fmt.Println(da.GetClassFor("class_loader:Test"))

	ins, e := da.CreateInstanceByName("class_loader:Test")
	if e != nil {
		fmt.Println("error:", e)
		return
	}

	testI, _ := ins.(*Test)

	testI.Hello("DynamicAccess")
}
```