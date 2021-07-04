package abstract_factory

import "fmt"

// 抽象产品1
type Product1 interface {
	Show()
}

// 抽象产品2
type Product2 interface {
	Use()
}

// 抽象工厂
type AbstractFactory interface {
	NewProduct1() Product1
	NewProduct2() Product2
}


// 具体产品11
type ConcreteProduct11 struct{}

func (*ConcreteProduct11) Show() {
	fmt.Println("Product11 show")
}

// 具体产品12
type ConcreteProduct12 struct {}

func (*ConcreteProduct12)Use(){
	fmt.Println("Product12 use")
}

// 具体工厂1
type ConcreteFactory1 struct{}

func (ConcreteFactory1) NewProduct1() Product1 {
	return &ConcreteProduct11{}
}

func (ConcreteFactory1)NewProduct2() Product2{
	return &ConcreteProduct12{}
}

// 具体产品21
type ConcreteProduct21 struct{}

func (*ConcreteProduct21) Show() {
	fmt.Println("Product21 show")
}

type ConcreteProduct22 struct {}

func (*ConcreteProduct22) Use(){
	fmt.Println("Product22 use")
}


// 具体工厂2
type ConcreteFactory2 struct{}

func (ConcreteFactory2) NewProduct1() Product1 {
	return &ConcreteProduct21{}
}

func (ConcreteFactory2)NewProduct2()Product2{
	return &ConcreteProduct22{}
}
