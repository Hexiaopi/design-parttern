package command

type Commander interface {
	Execute()
	Undo()
}

type AddCommand struct {
	receiver *Origin
	value    int
}

func (c *AddCommand) Execute() {
	c.receiver.Add(c.value)
}

func (c *AddCommand) Undo() {
	c.receiver.Subtract(c.value)
}

type SubtractCommand struct {
	receiver *Origin
	value    int
}

func (c *SubtractCommand) Execute() {
	c.receiver.Subtract(c.value)
}

func (c *SubtractCommand) Undo() {
	c.receiver.Add(c.value)
}

type MultiplyCommand struct {
	receiver *Origin
	value    int
}

func (c *MultiplyCommand) Execute() {
	c.receiver.Multiply(c.value)
}

func (c *MultiplyCommand) Undo() {
	c.receiver.Divide(c.value)
}

type DivideCommand struct {
	receiver *Origin
	value    int
}

func (c *DivideCommand) Execute() {
	c.receiver.Divide(c.value)
}

func (c *DivideCommand) Undo() {
	c.receiver.Multiply(c.value)
}

type Origin struct {
	value int
}

func (r *Origin) Add(value int) {
	r.value += value
}

func (r *Origin) Subtract(value int) {
	r.value -= value
}

func (r *Origin) Multiply(value int) {
	r.value *= value
}

func (r *Origin) Divide(value int) {
	r.value /= value
}

type Computer struct {
	commands []Commander
}

func (i *Computer) AddCommand(command Commander) {
	i.commands = append(i.commands, command)
}

func (i *Computer) ExecuteCommands() {
	for _, command := range i.commands {
		command.Execute()
	}
}

func (i *Computer) UndoLastCommand() {
	if len(i.commands) > 0 {
		cmd := i.commands[len(i.commands)-1]
		cmd.Undo()
		i.commands = i.commands[0 : len(i.commands)-1]
	}
}
