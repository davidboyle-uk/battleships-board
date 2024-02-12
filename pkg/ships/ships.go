package ships

import (
	"fmt"
	"sort"

	"github.com/dbx123/battleships-board/types"
)

func GenerateShips(boardSize int) types.Ships {
	var ships types.Ships
	for _, s := range types.NewFleet() {
		for i := 1; i <= s.Num; i++ {
			ships = append(ships, generateShip(boardSize, ships, s.ShipType))
		}
	}
	return ships
}

func generateShip(boardSize int, g types.Ships, t types.ShipType) types.Ship {
	for {
		for _, dir := range types.ShipDirections {
			s, err := expandShip(randCoord(boardSize), boardSize, t, dir)
			if err != nil {
				continue
			}
			if !hasCollisions(s, g) {
				return s
			}
		}
	}
}

func hasCollisions(s types.Ship, g types.Ships) bool {
	for _, i := range g {
		for _, t := range i.Coords {
			for _, c := range s.Coords {
				if c == t {
					return true
				}
			}
		}
	}
	return false
}

func expandShip(c types.Coord, boardSize int, t types.ShipType, dir types.ShipDirection) (types.Ship, error) {
	ship := types.Ship{
		Coords: []types.Coord{c},
	}

	l := int(t)
	if l == 1 {
		return ship, nil
	}

	var next types.Coord
	for i := 1; i <= int(t)-1; i++ {
		switch dir {
		case types.HORIZONTAL:
			next = types.Coord{
				c.X + i,
				c.Y,
			}
			if next.IsOutOfBounds(boardSize) {
				return types.Ship{}, fmt.Errorf("out of bounds")
			}
			ship.Coords = append(ship.Coords, next)
		case types.VERTICAL:
			next = types.Coord{
				c.X,
				c.Y + i,
			}
			if next.IsOutOfBounds(boardSize) {
				return types.Ship{}, fmt.Errorf("out of bounds")
			}
			ship.Coords = append(ship.Coords, next)
		}
	}

	return ship, nil
}

func randCoord(boardSize int) types.Coord {
	return types.Coord{
		types.RandNum(boardSize),
		types.RandNum(boardSize),
	}
}

func sortShips(s types.Ships) {
	sort.Slice(s, func(i, j int) bool {
		return len(s[i].Coords) < len(s[j].Coords)
	})
}
