package memento

type Originator struct {
	state string
}

func NewOriginator(value string) *Originator {
	return &Originator{state: value}
}

func (o *Originator) ChangeState(value string) {
	o.state = value
}

func (o *Originator) Save() Memento {
	return NewMemento(o.state)
}

func (o *Originator) Restore(memento Memento) {
	o.state = memento.GetState()
}

type Memento struct {
	state string
}

func NewMemento(state string) Memento {
	return Memento{state: state}
}

func (m Memento) GetState() string {
	return m.state
}

type Caretaker struct {
	history []Memento
}

func (c *Caretaker) AddMemento(memento Memento) {
	c.history = append(c.history, memento)
}

func (c *Caretaker) GetMemento(index int) Memento {
	return c.history[index]
}
