package factory_method

import "fmt"

// 抽象产品
type Product interface {
	Show()
}

// 抽象工厂
type AbstractFactory interface {
	NewProduct() Product
}

// 具体产品1
type ConcreteProduct1 struct{}

func (*ConcreteProduct1) Show() {
	fmt.Println("Product1 show")
}

// 具体工厂1
type ConcreteFactory1 struct{}

func (ConcreteFactory1) NewProduct() Product {
	return &ConcreteProduct1{}
}

// 具体产品2
type ConcreteProduct2 struct{}

func (*ConcreteProduct2) Show() {
	fmt.Println("Product2 show")
}

// 具体工厂2
type ConcreteFactory2 struct{}

func (ConcreteFactory2) NewProduct() Product {
	return &ConcreteProduct2{}
}
