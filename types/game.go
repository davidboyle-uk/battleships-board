package types

type Game struct {
	Players []*Player
}

type Player struct {
	Name  string
	Board Board
	Moves Moves
	Hits  int
}
