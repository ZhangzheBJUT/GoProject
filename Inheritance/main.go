package main

import (
	h "Inheritance/composition"
	n "Inheritance/new"
	o "Inheritance/old"
	p "Inheritance/polymorphism"
	"fmt"
	"reflect"
)

func main() {

	//分析问题，首先这里“继承”两个词就用错了，在go中不应该提及“继承”这个词，我更选择使用“嵌套”这个词。
	//B是嵌套了A，所以这里的b.Run()实际上是语法糖，调用的是b.A.Run()。这里Run的全部环境都在A中。
	//所以是找不到A的Test的。

	//所以在golang中要模拟普通的继承，除了使用嵌套之外，还需要在父类中“注册”子类的信息！

	b := new(o.B)
	b.Run()

	c := new(n.B)
	c.A.Son = c
	c.Run()

	d := n.SharedB()
	d.Run()

	//i := reflect.ValueOf(2)
	//fmt.Println(i)

	v := reflect.ValueOf("hello world")
	t := reflect.TypeOf("hello world")
	fmt.Println(v)
	fmt.Println(t)

	mark := h.Student{h.Human{"Mark", 25, "222-222-YYYY"}, "MIT"}
	sam := h.Employee{h.Human{"Sam", 45, "111-888-XXXX"}, "Golang Inc"}

	mark.SayHi()
	sam.SayHi()
	sam.Human.SayHi()

	s := p.ST2{I: 5}
	s.Show()
	s.Show2()
	println(s.I)

}
