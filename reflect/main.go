package main

import (
	"errors"
	"fmt"
	"reflect"
)

type MyInt int

func TestKind() {
	var x MyInt = 7
	fmt.Println("kind:", reflect.ValueOf(x).Kind()) //kind: int
}
func TestSet() {
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	fmt.Println("settability of v:", v.CanSet()) //settability of v: false
	//v.SetFloat(7.1)                              //TestSet() //panic: reflect: reflect.Value.SetFloat using unaddressable valu

	p := reflect.ValueOf(&x)
	fmt.Println("type of p:", p.Type())                        //type of p: *float64
	fmt.Println("settability of p:", p.CanSet())               //settability of p: false
	fmt.Println("settability of p.Elem():", p.Elem().CanSet()) //of p.Elem(): true

	reflect.ValueOf(&x).Elem().SetFloat(7.1)
	fmt.Println(x)                              //7.1
	fmt.Println(reflect.ValueOf(x).Interface()) //7.1
}

type T struct {
	A int
	B string
	C error
}

func TestStruct() {

	t := T{A: 23, B: "Hello world", C: nil}
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
	fmt.Println("t.c is nil:", reflect.ValueOf(t).Field(2).IsNil()) // t.c is nil: true
	t.C = errors.New("Error")
	fmt.Println("t.c is nil:", reflect.ValueOf(t).Field(2).IsNil()) // t.c is nil: false
}
func main() {
	var x float64 = 3.4
	fmt.Println("type:", reflect.TypeOf(x))                                       //type: float64
	fmt.Println("value:", reflect.ValueOf(x))                                     //value: <float64 Value>
	fmt.Println("kind is float64:", reflect.ValueOf(x).Kind() == reflect.Float64) //kind is float64: true
	fmt.Println("value:", reflect.ValueOf(x).Float())                             //value: 3.4

	TestKind()

	fmt.Println("value:", reflect.ValueOf(x).Interface().(float64)) //value: 3.4
	fmt.Printf("value:%7.1e\n", reflect.ValueOf(x).Interface())     //value:3.4e+00

	TestSet()
	TestStruct()

}
