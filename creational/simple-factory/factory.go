package simple_factory

import "fmt"

// 抽象产品
type Product interface {
	Show()
}

const (
	Product1 = 1+iota
	Product2
)

// 工厂
type SimpleFactory struct {}

func (SimpleFactory)NewProduct(i int) Product {
	switch i {
	case Product1:
		return &ConcreteProduct1{}
	case Product2:
		return &ConcreteProduct2{}
	}
	return nil
}

// 具体产品1
type ConcreteProduct1 struct{}

func (*ConcreteProduct1) Show() {
	fmt.Println("Product1 show")
}

// 具体产品2
type ConcreteProduct2 struct{}

func (*ConcreteProduct2) Show() {
	fmt.Println("Product2 show")
}
