package game

func ExampleBuilder() {
	director := &Director{}
	wb := &WarriorBuilder{&Warrior{}}
	director.SetBuilder(wb)
	warrior := director.Construct()
	warrior.Show()
	// Output:
	// name:Warrior
	// level:10
	// equipment:[Sword Shield]
	// skills:[Slash Block]
}
