package polymorphism

//golang语言中没有继承，但是可以依靠组合来模拟继承和多态。
type ST struct {
}

func (s *ST) Show() {
	println("ST")
}

func (s *ST) Show2() {
	println("ST:Show2()")
}

type ST2 struct {
	ST
	I int
}

func (s *ST2) Show() {
	s.ST.Show()
	println("ST2")
}
