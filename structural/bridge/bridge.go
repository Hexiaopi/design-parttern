package bridge

import "fmt"

type Implementor interface {
	OperationImpl()
}

type ConcreateImplA struct{}

func (ConcreateImplA) OperationImpl() {
	fmt.Println("ConcreateImplA")
}

type ConcreateImplB struct{}

func (ConcreateImplB) OperationImpl() {
	fmt.Println("ConcreateImplB")
}

type Abstraction struct {
	Impl Implementor
}

func (a Abstraction) Operation() {
	fmt.Println("Abstraction")
	a.Impl.OperationImpl()
}

type RefineAbstraction struct {
	Impl Implementor
}

func (ra RefineAbstraction) Operation() {
	fmt.Println("RefineAbstraction")
	ra.Impl.OperationImpl()
}
