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

type Ship []Coord

type Ships []Ship

func (i Ships) AsString() string {
	var r string
	for _, s := range i {
		var ss string
		ss = fmt.Sprintln(s[0].String())
		if len(s) > 1 {
			ss = fmt.Sprintf("%s:%s\n", s[0].String(), s[len(s)-1].String())
		}
		r += ss
	}
	return r
}
