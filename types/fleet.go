package types

import "fmt"

type Flotilla struct {
	ShipType ShipType
	Num      int
}

type Fleet []Flotilla

var (
	ShipDirections = map[int]ShipDirection{
		0: HORIZONTAL,
		1: VERTICAL,
	}
)

func NewFleet() Fleet {
	return Fleet{
		{SUBMARINE, 2},
		{DESTROYER, 2},
		{CRUISER, 1},
		{BATTLESHIP, 1},
		{CARRIER, 1},
	}
}

func (f *Fleet) RemoveShip(t ShipType) error {
	tmp := *f
	for k, flotilla := range tmp {
		if flotilla.ShipType == t {
			if flotilla.Num < 1 {
				return fmt.Errorf("empty flotilla")
			}
			if flotilla.Num == 1 {
				tmp = append(tmp[:k], tmp[k+1:]...)
				*f = tmp
				return nil
			}
			flotilla.Num--
			tmp[k] = flotilla
			*f = tmp
			return nil
		}
	}
	return fmt.Errorf("shiptype not in fleet")
}
