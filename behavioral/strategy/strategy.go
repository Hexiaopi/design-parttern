package strategy

import "fmt"

type Strategy interface {
	Execute()
}

type ConcreteStrategy1 struct {
}

func (s ConcreteStrategy1) Execute() {
	fmt.Println("concrete strategy1 execute")
}

type ConcreteStrategy2 struct {
}

func (s ConcreteStrategy2) Execute() {
	fmt.Println("concrete strategy2 execute")
}

type Context struct {
	strategy Strategy
}

func (c *Context) SetStrategy(strategy Strategy) {
	c.strategy = strategy
}

func (c *Context) Do() {
	fmt.Println("do something")
	c.strategy.Execute()
}
