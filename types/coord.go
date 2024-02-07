package types

import "fmt"

type Coord struct {
	X int
	Y int
}

func (c Coord) String() string {
	return fmt.Sprintf("%v %v", c.X, c.Y)
}

func (c Coord) IsOutOfBounds(boardSize int) bool {
	return c.X > boardSize-1 ||
		c.Y > boardSize-1 ||
		c.X < 0 ||
		c.Y < 0
}
