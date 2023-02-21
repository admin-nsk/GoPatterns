//Flyweight is a structural design pattern that allows programs to support large
//numbers of objects while keeping memory usage low.
//
//The pattern accomplishes this by sharing parts of the object’s state between multiple objects.
//	In other words, the Flyweight saves RAM memory by caching the same information used by different objects.

package Structural

import "fmt"

const (
	TerroristDressType        = "tDress"
	CounterTerroristDressType = "ctDress"
)

var (
	dressFactorySingleInstance = &dressFactory{
		dressMap: make(map[string]dress),
	}
)

type dressFactory struct {
	dressMap map[string]dress
}

func (d *dressFactory) getDressByType(dressType string) (dress, error) {
	if d.dressMap[dressType] != nil {
		return d.dressMap[dressType], nil
	}

	if dressType == TerroristDressType {
		d.dressMap[dressType] = newTerroristDress()
		return d.dressMap[dressType], nil
	}

	if dressType == CounterTerroristDressType {
		d.dressMap[dressType] = newCounterTerroristDress()
		return d.dressMap[dressType], nil
	}
	return nil, fmt.Errorf("Wrong dress type passed")
}

func getDressFactorySingleInstance() *dressFactory {
	return dressFactorySingleInstance
}

type dress interface {
	getColor() string
}

type terroristDress struct {
	color string
}

func (t *terroristDress) getColor() string {
	return t.color
}

func newTerroristDress() *terroristDress {
	return &terroristDress{color: "red"}
}

type counterTerroristDress struct {
	color string
}

func (c *counterTerroristDress) getColor() string {
	return c.color
}

func newCounterTerroristDress() *counterTerroristDress {
	return &counterTerroristDress{color: "green"}
}

type player struct {
	dress      dress
	playerType string
	lat        int
	long       int
}

func newPlayer(playerType, dressType string) *player {
	dress, _ := getDressFactorySingleInstance().getDressByType(dressType)
	return &player{playerType: playerType, dress: dress}

}

func (p player) newLocation(lat, long int) {
	p.lat = lat
	p.long = long
}

type game struct {
	terrorist        []*player
	counterTerrorist []*player
}

func newGame() *game {
	return &game{
		terrorist:        make([]*player, 1),
		counterTerrorist: make([]*player, 1),
	}
}

func (g game) addTerrorist(dressType string) {
	player := newPlayer("T", dressType)
	g.terrorist = append(g.terrorist, player)
}

func (g game) addCounterTerrorist(dressType string) {
	player := newPlayer("CT", dressType)
	g.counterTerrorist = append(g.counterTerrorist, player)
	return
}

func runFlyweight() {
	game := newGame()

	game.addTerrorist(TerroristDressType)
	game.addTerrorist(TerroristDressType)
	game.addTerrorist(TerroristDressType)
	game.addTerrorist(TerroristDressType)

	game.addCounterTerrorist(CounterTerroristDressType)
	game.addCounterTerrorist(CounterTerroristDressType)
	game.addCounterTerrorist(CounterTerroristDressType)
	game.addCounterTerrorist(CounterTerroristDressType)

	dressFactoryInstance := getDressFactorySingleInstance()

	for dressType, dress := range dressFactoryInstance.dressMap {
		fmt.Printf("DressColorType: %s\nDressColor: %s\n", dressType, dress.getColor())
	}

}
