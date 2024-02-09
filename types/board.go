package types

type State rune

const (
	SEA  State = '-'
	HIT  State = 'h'
	MISS State = 'm'
	SUNK State = 'd'
)

func (s State) String() string {
	return string(s)
}

type Moves map[Coord]CoordState

type CoordState struct {
	Ship  *Ship
	State State
}

type Board struct {
	Dim   int
	Moves Moves
	Ships []*Ship
}

func (b Board) HasHits() bool {
	for _, move := range b.Moves {
		if move.State == HIT {
			return true
		}
	}
	return false
}

func (b Board) GetHits() []Coord {
	var hits = []Coord{}
	for c, move := range b.Moves {
		if move.State == HIT {
			hits = append(hits, c)
		}
	}
	return hits
}

func (b Board) GetSunk() []Coord {
	var sunk = []Coord{}
	for c, move := range b.Moves {
		if move.State == SUNK {
			sunk = append(sunk, c)
		}
	}
	return sunk
}
