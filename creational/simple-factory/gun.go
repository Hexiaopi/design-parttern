package simple_factory

type iGun interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}

func NewGun(name string) iGun {
	switch name {
	case "ak47":
		return newAk47()
	case "m4a1":
		return newM4a1()
	}
	return nil
}

type gun struct {
	name  string
	power int
}

func (g *gun) setName(name string) {
	g.name = name
}

func (g *gun) getName() string {
	return g.name
}

func (g *gun) setPower(power int) {
	g.power = power
}

func (g *gun) getPower() int {
	return g.power
}

type ak47 struct {
	gun
}

func newAk47() iGun {
	return &ak47{gun{
		name:  "AK47 Gun",
		power: 4,
	}}
}

type m4a1 struct {
	gun
}

func newM4a1() iGun {
	return &m4a1{gun{
		name:  "M4A1 Gun",
		power: 3,
	}}
}
