package adapter

import "fmt"

// Target是适配的目标接口
type Target interface {
	Request()
}

// Adaptee是被适配的目标接口
type Adaptee interface {
	SpecificRequest()
}

// AdapteeImpl是被适配的目标类
type AdapteeImpl struct{}

func (impl *AdapteeImpl) SpecificRequest() {
	fmt.Println("adapteeImpl SpecificRequest")
}

// NewAdaptee是被适配接口的工厂函数
func NewAdaptee() Adaptee {
	return &AdapteeImpl{}
}

// Adapter是转换Adaptee为Target接口的适配器
type Adapter struct {
	Adaptee
}

// Adapter适配器实现目标接口
func (a Adapter) Request() {
	a.SpecificRequest()
}

// NewAdapter是Adapter的工厂函数，将Adaptee转化为Target
func NewAdapter(adaptee Adaptee) Target {
	return &Adapter{Adaptee: adaptee}
}
