package templatemethod

import "fmt"

type Templater interface {
	Template()
}

type Implement interface {
	step1()
	step2()
	step3()
}

type Abstract struct {
	Implement
}

func newAbstract(impl Implement) *Abstract {
	return &Abstract{Implement: impl}
}

func (t *Abstract) Template() {
	t.Implement.step1()
	t.Implement.step2()
	t.Implement.step3()
}

func (t *Abstract) step2() {
	fmt.Println("default step2")
}

type Concreate1 struct {
	*Abstract
}

func NewConcreate1() Templater {
	c := &Concreate1{}
	a := newAbstract(c)
	c.Abstract = a
	return c
}

func (*Concreate1) step1() {
	fmt.Println("concreate1 step1")
}

func (*Concreate1) step2() {
	fmt.Println("concreate1 step2")
}

func (*Concreate1) step3() {
	fmt.Println("concreate1 step3")
}

type Concreate2 struct {
	*Abstract
}

func NewConcreate2() Templater {
	c := &Concreate2{}
	a := newAbstract(c)
	c.Abstract = a
	return c
}

func (*Concreate2) step1() {
	fmt.Println("concreate2 step1")
}

func (*Concreate2) step3() {
	fmt.Println("concreate2 step3")
}
