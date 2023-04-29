package command

import "fmt"

func ExampleComputer() {
	computer := Computer{}
	origin := Origin{10}
	addCmd := AddCommand{&origin, 20}
	subCmd := SubtractCommand{&origin, 15}
	mulCmd := MultiplyCommand{&origin, 2}
	divCmd := DivideCommand{&origin, 3}
	computer.AddCommand(&addCmd)
	computer.AddCommand(&subCmd)
	computer.AddCommand(&mulCmd)
	computer.AddCommand(&divCmd)
	computer.ExecuteCommands()
	fmt.Println(origin.value)
	// Output:
	// 10
}

func ExampleComputerUndo(){
	computer := Computer{}
	origin := Origin{10}
	addCmd := AddCommand{&origin, 20}
	subCmd := SubtractCommand{&origin, 15}
	mulCmd := MultiplyCommand{&origin, 2}
	divCmd := DivideCommand{&origin, 3}
	computer.AddCommand(&addCmd)
	computer.AddCommand(&subCmd)
	computer.AddCommand(&mulCmd)
	computer.AddCommand(&divCmd)
	computer.ExecuteCommands()
	fmt.Println(origin.value)
	computer.UndoLastCommand()
	fmt.Println(origin.value)
	computer.UndoLastCommand()
	fmt.Println(origin.value)
	computer.UndoLastCommand()
	fmt.Println(origin.value)
	// Output:
	// 10
	// 30
  // 15
	// 30
}