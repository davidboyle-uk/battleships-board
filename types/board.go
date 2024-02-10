package types

import (
	"fmt"
	"strconv"
)

type State string

const (
	SEA  State = "-"
	HIT  State = "h"
	MISS State = "m"
	SUNK State = "d"
)

type Moves map[Coord]CoordState

type CoordState struct {
	Ship  Ship
	State State
}

type Board struct {
	Dim     int
	Moves   Moves
	ShipTot int
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

func (b Board) ToString() string {
	var s string
	s += fmt.Sprintln(strconv.Itoa(b.Dim))
	for y := 0; y <= b.Dim-1; y++ {
		for x := 0; x <= b.Dim-1; x++ {
			if move, ok := b.Moves[Coord{X: x, Y: y}]; ok {
				s += string(move.State)
			} else {
				s += string(SEA)
			}
		}
		s += "\n"
	}

	return s
}
