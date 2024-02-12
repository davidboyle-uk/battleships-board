package types

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type State string

const (
	SEA  State = "-"
	HIT  State = "h"
	MISS State = "m"
	SUNK State = "d"
)

type Moves map[string]CoordState

type CoordState struct {
	Ship  *Ship
	State State
}

type Board struct {
	Dim     int
	Moves   Moves
	ShipTot int
}

func CoordFromString(s string) Coord {
	bits := strings.Split(s, " ")
	if len(bits) < 2 {
		panic(fmt.Sprintf("invalid coord %v", s))
	}
	x, _ := strconv.Atoi(bits[0])
	y, _ := strconv.Atoi(bits[1])

	return Coord{
		X: x,
		Y: y,
	}
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
			hits = append(hits, CoordFromString(c))
		}
	}
	return hits
}

func (b Board) GetSunk() []Coord {
	var sunk = []Coord{}
	for c, move := range b.Moves {
		if move.State == SUNK {
			sunk = append(sunk, CoordFromString(c))
		}
	}
	return sunk
}

func (b Board) ToString() string {
	var s string
	s += fmt.Sprintln(strconv.Itoa(b.Dim))
	for y := 0; y <= b.Dim-1; y++ {
		for x := 0; x <= b.Dim-1; x++ {
			c := SEA
			if move, ok := b.Moves[Coord{X: x, Y: y}.String()]; ok {
				c = move.State
			}
			s += string(c)
		}
		s += "\n"
	}

	return s
}

func (b Board) AddShips(s Ships) {
	for _, sh := range s {
		ns := sh
		for _, c := range sh.Coords {
			key := c.String()
			b.Moves[key] = CoordState{
				Ship:  &ns,
				State: SEA,
			}
		}
	}
}

func toJSON(data interface{}) string {
	b, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	return string(b)
}
