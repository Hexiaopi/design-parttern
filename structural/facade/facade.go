package facade

import "fmt"

type Facade struct {
	system1 Subsystem1
	system2 Subsystem2
	system3 Subsystem3
}

func (f Facade) Operation() {
	f.system1.Operation1()
	f.system2.Operation2()
	f.system3.Operation3()
}

type Subsystem1 struct{}

func (s Subsystem1) Operation1() {
	fmt.Println("subsystem1 operation")
}

type Subsystem2 struct{}

func (s Subsystem2) Operation2() {
	fmt.Println("subsystem2 operation")
}

type Subsystem3 struct{}

func (s Subsystem3) Operation3() {
	fmt.Println("subsystem3 operation")
}
