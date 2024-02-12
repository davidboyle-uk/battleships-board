package types

import "fmt"

type ShipType int

const (
	SUBMARINE ShipType = iota + 1
	DESTROYER
	CRUISER
	BATTLESHIP
	CARRIER
)

type ShipDirection int

const (
	HORIZONTAL ShipDirection = iota
	VERTICAL
)

type Ship struct {
	Coords []Coord
	Hits   int
}

type Ships []Ship

func (i Ships) AsString() string {
	var r string
	for _, s := range i {
		var ss string
		ss = fmt.Sprintln(s.Coords[0].String())
		if len(s.Coords) > 1 {
			ss = fmt.Sprintf("%s:%s\n", s.Coords[0].String(), s.Coords[len(s.Coords)-1].String())
		}
		r += ss
	}
	return r
}

func (i Ships) GetVolume() int {
	var vol int
	for _, s := range i {
		vol += len(s.Coords)
	}
	return vol
}
