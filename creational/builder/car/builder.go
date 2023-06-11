package car

type Car struct {
	Wheel         string // 车轮
	Engine        string // 发动机
	Body          string // 车身
	SteeringWheel string //方向盘
}

type Builder interface {
	BuildWheel() Builder
	BuildEngine() Builder
	BuildBody() Builder
	BuildSteeringWheel() Builder
	GetResult() *Car
}

type BenzBuilder struct {
	car *Car
}

func NewBenzBuilder() *BenzBuilder {
	return &BenzBuilder{car: &Car{}}
}

func (b *BenzBuilder) BuildWheel() Builder {
	b.car.Wheel = "Benz Wheel"
	return b
}

func (b *BenzBuilder) BuildEngine() Builder {
	b.car.Engine = "Benz Engine"
	return b
}

func (b *BenzBuilder) BuildBody() Builder {
	b.car.Body = "Benz Body"
	return b
}

func (b *BenzBuilder) BuildSteeringWheel() Builder {
	b.car.SteeringWheel = "Benz Steering Wheel"
	return b
}

func (b *BenzBuilder) GetResult() *Car {
	return b.car
}

type BMWBuilder struct {
	car *Car
}

func NewBMWBuilder() *BMWBuilder {
	return &BMWBuilder{car: &Car{}}
}

func (b *BMWBuilder) BuildWheel() Builder {
	b.car.Wheel = "BMW Wheel"
	return b
}

func (b *BMWBuilder) BuildEngine() Builder {
	b.car.Engine = "BMW Engine"
	return b
}

func (b *BMWBuilder) BuildBody() Builder {
	b.car.Body = "BMW Body"
	return b
}

func (b *BMWBuilder) BuildSteeringWheel() Builder {
	b.car.SteeringWheel = "BMW Steering Wheel"
	return b
}

func (b *BMWBuilder) GetResult() *Car {
	return b.car
}

type Director struct {
	build Builder
}

func (d *Director) SetBuilder(build Builder) {
	d.build = build
}

func (d *Director) Construct() *Car {
	d.build.BuildWheel()
	d.build.BuildEngine()
	d.build.BuildBody()
	d.build.BuildSteeringWheel()
	return d.build.GetResult()
}
