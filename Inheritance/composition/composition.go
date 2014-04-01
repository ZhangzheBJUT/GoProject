package composition

import (
	"fmt"

//	"image/color"
)

type Human struct {
	Name  string
	Age   int
	Phone string
}

type Student struct {
	Human         // 匿名字段
	School string // 聚合字段

}

type Employee struct {
	Human
	Company string
}

func (h *Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.Name, h.Phone)
}

//Employee的method重写Human的method
func (e *Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.Name,
		e.Company, e.Phone) //Yes you can split into 2 lines here.
}
