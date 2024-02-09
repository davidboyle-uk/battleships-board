package types

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
)

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

func (c Coord) Add(m Coord) Coord {
	return Coord{
		X: c.X + m.X,
		Y: c.Y + m.Y,
	}
}

func (c Coord) AddNum(m Coord, num int) Coord {
	res := Coord{
		X: c.X,
		Y: c.Y,
	}

	for i := 0; i <= num-1; i++ {
		res = res.Add(m)
	}

	return res
}

func (c Coord) Dist(m Coord) float64 {
	first := math.Pow(float64(c.X-m.X), 2)
	second := math.Pow(float64(c.Y-m.Y), 2)
	return math.Sqrt(first + second)
}

func RandNum(boardSize int) int {
	bg := big.NewInt(int64(boardSize))
	n, _ := rand.Int(rand.Reader, bg)
	return int(n.Int64())
}
