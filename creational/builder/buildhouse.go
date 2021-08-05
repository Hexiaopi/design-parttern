package builder

import "fmt"

type house struct {
	windowType string
	doorType   string
	floorNum   int
}

func (h *house) setWindowType(window string) {
	h.windowType = window
}

func (h *house) setDoorType(door string) {
	h.doorType = door
}

func (h *house) setFloorNum(num int) {
	h.floorNum = num
}

func (h house) show() {
	fmt.Printf("window=%s,door=%s,num of floor=%d\n", h.windowType, h.doorType, h.floorNum)
}

type iBuilder interface {
	buildWindow()
	buildDoor()
	buildFloor()
	getHouse() house
}

type normalBuilder struct {
	*house
}

func (builder normalBuilder) buildWindow() {
	builder.house.setWindowType("Wooden Window")
}

func (builder normalBuilder) buildDoor() {
	builder.house.setDoorType("Wooden Door")
}

func (builder normalBuilder) buildFloor() {
	builder.house.setFloorNum(2)
}

func (builder normalBuilder) getHouse() house {
	return *builder.house
}

type iglooBuilder struct {
	*house
}

func (builder iglooBuilder) buildWindow() {
	builder.house.setWindowType("Snow Window")
}

func (builder iglooBuilder) buildDoor() {
	builder.house.setDoorType("Snow Door")
}

func (builder iglooBuilder) buildFloor() {
	builder.house.setFloorNum(1)
}

func (builder iglooBuilder) getHouse() house {
	return *builder.house
}

type director struct {
	builder iBuilder
}

func (d director) buildHouse() house {
	d.builder.buildWindow()
	d.builder.buildDoor()
	d.builder.buildFloor()
	return d.builder.getHouse()
}
