package old

import (
	"reflect"
)

type A struct {
}

func (this A) Run() {
	// ValueOf returns a new Value initialized to the concrete value
	// stored in the interface i.  ValueOf(nil) returns the zero Value.
	c := reflect.ValueOf(this)
	method := c.MethodByName("Test")
	println(method.IsValid())
}

type B struct {
	A
}

func (self B) Test(s string) {
	println("B")
}
