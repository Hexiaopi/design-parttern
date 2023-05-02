package memento

import "fmt"

func ExampleMemento() {
	caretaker := Caretaker{history: make([]Memento, 0)}
	originator := NewOriginator("one")
	caretaker.AddMemento(originator.Save())
	fmt.Println(originator.state)
	originator.ChangeState("two")
	caretaker.AddMemento(originator.Save())
	fmt.Println(originator.state)
	originator.ChangeState("three")
	caretaker.AddMemento(originator.Save())
	fmt.Println(originator.state)
	originator.Restore(caretaker.GetMemento(1))
	fmt.Println(originator.state)
	// Output:
	// one
	// two
	// three
	// two
}
