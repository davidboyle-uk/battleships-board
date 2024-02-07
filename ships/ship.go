package ship

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"sort"

	"battleships-board/types"
)

var (
	gameShips = []struct {
		ShipType types.ShipType
		Num      int
	}{
		{types.SUBMARINE, 2},
		{types.DESTROYER, 2},
		{types.CRUISER, 1},
		{types.BATTLESHIP, 1},
		{types.CARRIER, 1},
	}
	shipDirections = map[int]types.ShipDirection{
		0: types.HORIZONTAL,
		1: types.VERTICAL,
	}
)

func orderShips(s types.Ships) {
	// Sort values by length
	sort.Slice(s, func(i, j int) bool {
		return len(s[i]) < len(s[j])
	})
}

func generateShips(boardSize int) types.Ships {
	var ships types.Ships
	for _, s := range gameShips {
		for i := 1; i <= s.Num; i++ {
			ships = append(ships, generateShip(boardSize, ships, s.ShipType))
		}
	}
	return ships
}

func generateShip(boardSize int, g types.Ships, t types.ShipType) types.Ship {
	for {
		for _, dir := range shipDirections {
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
		for _, t := range i {
			for _, c := range s {
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
		c,
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
			ship = append(ship, next)
		case types.VERTICAL:
			next = types.Coord{
				c.X,
				c.Y + i,
			}
			if next.IsOutOfBounds(boardSize) {
				return types.Ship{}, fmt.Errorf("out of bounds")
			}
			ship = append(ship, next)
		}
	}

	return ship, nil
}

func randCoord(boardSize int) types.Coord {
	return types.Coord{
		randNum(boardSize),
		randNum(boardSize),
	}
}

func randNum(boardSize int) int {
	bg := big.NewInt(int64(boardSize))
	n, _ := rand.Int(rand.Reader, bg)
	return int(n.Int64())
}
