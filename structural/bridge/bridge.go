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

type Abstraction interface {
	Operation()
}

type RefineAbstractionA struct {
	Impl Implementor
}

func (ra RefineAbstractionA) Operation() {
	fmt.Println("RefineAbstractionA")
	ra.Impl.OperationImpl()
}

type RefineAbstractionB struct {
	Impl Implementor
}

func (ra RefineAbstractionB) Operation() {
	fmt.Println("RefineAbstractionB")
	ra.Impl.OperationImpl()
}
