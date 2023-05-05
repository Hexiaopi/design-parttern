package visitor

import "fmt"

type Visitor interface {
	VisitConcreteElementA(element *ConcreteElementA)
	VisitConcreteElementB(element *ConcreteElementB)
}

type ConcreteVisitor struct{}

func (c *ConcreteVisitor) VisitConcreteElementA(element *ConcreteElementA) {
	fmt.Println("Visit ConcreteElementA")
}

func (c *ConcreteVisitor) VisitConcreteElementB(element *ConcreteElementB) {
	fmt.Println("Visit ConcreteElementB")
}

type Element interface {
	Accept(visitor Visitor)
}

type ConcreteElementA struct{}

func (c *ConcreteElementA) Accept(visitor Visitor) {
	visitor.VisitConcreteElementA(c)
}

type ConcreteElementB struct{}

func (c *ConcreteElementB) Accept(visitor Visitor) {
	visitor.VisitConcreteElementB(c)
}
