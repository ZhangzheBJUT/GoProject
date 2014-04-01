package new

import (
	"reflect"
)

type A struct {
	Son interface{}
}

func (this A) Run() {
	// ValueOf returns a new Value initialized to the concrete value
	// stored in the interface i.  ValueOf(nil) returns the zero Value.
	c := reflect.ValueOf(this.Son)
	method := c.MethodByName("Test")
	println(method.IsValid())
}

type B struct {
	A
}

func (self B) Test(s string) {
	println("B")
}

func (self B) Run() {
	self.A.Run()
}

func SharedB() *B {
	c := new(B)
	c.A.Son = c
	return c
}
