package mediator

import "fmt"

type Mediator interface {
	Coordinate(colleague Colleague, message string)
}

type Colleague interface {
	SendMessage(message string)
	ReceiveMessage(message string)
}

type ConcreteColleague1 struct {
	mediator Mediator
}

func (c *ConcreteColleague1) SendMessage(message string) {
	fmt.Println("ConcreateColleague1 send message:", message)
	c.mediator.Coordinate(c, message)
}

func (c *ConcreteColleague1) ReceiveMessage(message string) {
	fmt.Println("ConcreateColleague1 receive message:", message)
}

type ConcreteColleague2 struct {
	mediator Mediator
}

func (c *ConcreteColleague2) SendMessage(message string) {
	fmt.Println("ConcreateColleague2 send message:", message)
	c.mediator.Coordinate(c, message)
}

func (c *ConcreteColleague2) ReceiveMessage(message string) {
	fmt.Println("ConcreateColleague2 receive message:", message)
}

type ConcreteMediator struct {
	ConcreteColleague1
	ConcreteColleague2
}

func (c ConcreteMediator) Coordinate(colleague Colleague, message string) {
	switch colleague.(type) {
	case *ConcreteColleague1:
		c.ConcreteColleague2.ReceiveMessage(message)
	case *ConcreteColleague2:
		c.ConcreteColleague1.ReceiveMessage(message)
	}
}
