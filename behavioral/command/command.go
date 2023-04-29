package command

import "fmt"

type Command interface {
	Execute()
}

type Receiver struct{}

func (r *Receiver) operation1(a string) {
	fmt.Println("operation1:", a)
}

func (r *Receiver) operation2(a, b, c string) {
	fmt.Println("operation2:", a, b, c)
}

type Invoker struct {
	cmd Command
}

func (i *Invoker) SetCommand(cmd Command) {
	i.cmd = cmd
}

func (i *Invoker) ExecuteCommand() {
	i.cmd.Execute()
}

type ConcreateCommand1 struct {
	name     string
	receiver *Receiver
}

func (c *ConcreateCommand1) Execute() {
	c.receiver.operation1(c.name)
}

type ConcreateCommand2 struct {
	name     string
	desc     string
	address  string
	receiver *Receiver
}

func (c *ConcreateCommand2) Execute() {
	c.receiver.operation2(c.name, c.desc, c.address)
}
