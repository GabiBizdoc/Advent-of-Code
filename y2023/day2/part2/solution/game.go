package solution

import "errors"

type Game struct {
	GameID int
	Sets   []*GameSet
}

func (g *Game) ComputeCubePower() int {
	colorMap := make(map[string]int)

	for _, set := range g.Sets {
		for color, count := range set.colorMap {
			old, ok := colorMap[color]
			if ok {
				colorMap[color] = max(old, count)
			} else {
				colorMap[color] = count
			}
		}
	}

	sum := 1
	// TODO: maybe return error if len(colorMap) is 0
	for _, value := range colorMap {
		sum *= value
	}
	return sum
}

func NewGame(gameID int) *Game {
	return &Game{GameID: gameID}
}

type GameSet struct {
	colorMap map[string]int
}

func NewGameSet() *GameSet {
	return &GameSet{colorMap: make(map[string]int)}
}

func (g *GameSet) AddCube(color string, cnt int) {
	g.colorMap[color] += cnt
}

func (g *GameSet) CountCubes(color string) int {
	return g.colorMap[color]
}

//func (g *GameSet) IsValid() bool {
//	return g.CountCubes("red") <= 12 &&
//		g.CountCubes("green") <= 13 &&
//		g.CountCubes("blue") <= 14
//}

var InvalidGameSet = errors.New("invalid game set")
