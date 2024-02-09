package types

type Direction int

const (
	UP Direction = iota
	DOWN
	LEFT
	RIGHT
)

var PossibleMoves = map[Direction]Coord{
	UP: {
		Y: 1,
	},
	DOWN: {
		Y: -1,
	},
	LEFT: {
		X: -1,
	},
	RIGHT: {
		X: 1,
	},
}

type CoordSpace struct {
	Coord Coord
	MinX  int
	MaxX  int
	MinY  int
	MaxY  int
}

type Probabilities map[int][]Coord
