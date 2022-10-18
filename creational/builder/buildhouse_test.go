package builder

func ExampleBuildHouse() {
	builder := normalBuilder{&house{}}
	dir := director{builder: builder}
	house := dir.buildHouse()
	house.show()
	// Output:
	// window=Wooden Window,door=Wooden Door,num of floor=2
}

func ExampleBuildHouse_second() {
	builder := iglooBuilder{&house{}}
	dir := director{builder: builder}
	house := dir.buildHouse()
	house.show()
	// Output:
	// window=Snow Window,door=Snow Door,num of floor=1
}
