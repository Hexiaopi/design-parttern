package game

import "fmt"

type Role interface {
	Show()
}

type Warrior struct {
	name      string   //名字
	level     int      //等级
	equipment []string //装备
	skills    []string //技能
}

func (w Warrior) Show() {
	fmt.Printf("name:%s\nlevel:%d\nequipment:%v\nskills:%s",
		w.name, w.level, w.equipment, w.skills)
}

type Builder interface {
	BuildName() Builder
	BuildLevel() Builder
	BuildEquipment() Builder
	BuildSkills() Builder
	GetResult() Role
}

type WarriorBuilder struct {
	role *Warrior
}

func (w *WarriorBuilder) BuildName() Builder {
	w.role.name = "Warrior"
	return w
}

func (w *WarriorBuilder) BuildLevel() Builder {
	w.role.level = 10
	return w
}

func (w *WarriorBuilder) BuildEquipment() Builder {
	w.role.equipment = []string{"Sword", "Shield"}
	return w
}

func (w *WarriorBuilder) BuildSkills() Builder {
	w.role.skills = []string{"Slash", "Block"}
	return w
}

func (w *WarriorBuilder) GetResult() Role {
	return w.role
}

type Director struct {
	builder Builder
}

func (d *Director) SetBuilder(builder Builder) {
	d.builder = builder
}

func (d *Director) Construct() Role {
	d.builder.BuildName().BuildLevel().BuildEquipment().BuildSkills()
	return d.builder.GetResult()
}
