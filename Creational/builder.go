package Creational

import "fmt"

type human struct {
	hType     string
	strong    float32
	hasGuitar bool
	hasGun    bool
	hasHelm   bool
	hasHat    bool
	hasSword  bool
}

type humanBuilder struct {
	human human
}

func newHumanBuilder() *humanBuilder {
	return &humanBuilder{
		human: human{},
	}
}

func (hb *humanBuilder) setStrong(strong float32) *humanBuilder {
	hb.human.strong = strong
	return hb
}

func (hb *humanBuilder) setType(t string) *humanBuilder {
	hb.human.hType = t
	return hb
}

func (hb *humanBuilder) withGuitar() *humanBuilder {
	hb.human.hasGuitar = true
	return hb
}

func (hb *humanBuilder) withGun() *humanBuilder {
	hb.human.hasGun = true
	return hb
}

func (hb *humanBuilder) withHelm() *humanBuilder {
	hb.human.hasHelm = true
	return hb
}

func (hb *humanBuilder) withHat() *humanBuilder {
	hb.human.hasHat = true
	return hb
}

func (hb *humanBuilder) withSword() *humanBuilder {
	hb.human.hasSword = true
	return hb
}

func (hb *humanBuilder) build() human {
	return hb.human
}

func run() {
	humanFighter := newHumanBuilder().
		setStrong(100.00).
		setType("Fighter").
		withGun().
		withHelm().
		withSword()

	humanSinger := newHumanBuilder().
		setStrong(50.00).
		setType("Singer").
		withGuitar().
		withHat()
	printHumanInfo(humanFighter.build())
	fmt.Println("---------------")
	printHumanInfo(humanSinger.build())
}

func printHumanInfo(human human) {
	humanInfMsg := fmt.Sprintf("Type: %s \n"+
		"Strong: %f \n"+
		"hasGuitar: %t \n"+
		"hasGun: %t \n"+
		"hasHelm: %t \n"+
		"hasHat: %t \n"+
		"hasSword: %t \n",
		human.hType,
		human.strong,
		human.hasGuitar,
		human.hasGun,
		human.hasHelm,
		human.hasHat,
		human.hasSword)
	fmt.Println(humanInfMsg)
}
