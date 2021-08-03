package builder

import "fmt"

type Product struct {
	partA string
	partB string
	partC string
}

func (p *Product) setPartA(partA string) {
	p.partA = partA
}

func (p *Product) setPartB(partB string) {
	p.partB = partB
}

func (p *Product) setPartC(partC string) {
	p.partC = partC
}

func (p Product) show() {
	fmt.Printf("A:%s\tB:%s\tC:%s", p.partA, p.partB, p.partC)
}

type Builder interface {
	buildPartA()
	buildPartB()
	buildPartC()
	getResult() Product
}

type Director struct {
	builder Builder
}

func (d *Director) Director(builder Builder) {
	d.builder = builder
}

func (d Director) construct() Product {
	d.builder.buildPartA()
	d.builder.buildPartB()
	d.builder.buildPartC()
	return d.builder.getResult()
}

type ConcreteBuilder1 struct {
	product *Product
}

func (c ConcreteBuilder1) buildPartA() {
	c.product.setPartA("A1")
}

func (c ConcreteBuilder1) buildPartB() {
	c.product.setPartB("B1")
}

func (c ConcreteBuilder1) buildPartC() {
	c.product.setPartC("C1")
}

func (c ConcreteBuilder1) getResult() Product {
	return *c.product
}

type ConcreteBuilder2 struct {
	product *Product
}

func (c ConcreteBuilder2) buildPartA() {
	c.product.setPartA("A2")
}

func (c ConcreteBuilder2) buildPartB() {
	c.product.setPartB("B2")
}

func (c ConcreteBuilder2) buildPartC() {
	c.product.setPartC("C2")
}

func (c ConcreteBuilder2) getResult() Product {
	return *c.product
}
