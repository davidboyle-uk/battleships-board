package board

import (
	"fmt"
	"strconv"

	"battleships-board/types"
)

const (
	HIT       = 'h'
	MISS      = 'm'
	DESTROYED = 'd'
)

func boardFromString(s string) (types.Board, error) {

	dim, board, err := getBoardSize(s)
	if err != nil {
		return types.Board{}, fmt.Errorf("unable to parse board")
	}

	l := len(s)
	max := (dim * dim)
	if l < max {
		return types.Board{}, fmt.Errorf("board length is %v or more [%v]", max, l)
	}

	var b types.Board

	var x, y int
	for _, r := range board {
		if r == '\n' {
			continue
		}
		if y%dim == 0 {
			if y > 0 {
				x++
			}
			if x > dim-1 {
				break
			}
			y = 0
		}
		switch r {
		case HIT:
			b.Hit = append(b.Hit, types.Coord{X: x, Y: y})
		case MISS:
			b.Miss = append(b.Miss, types.Coord{X: x, Y: y})
		case DESTROYED:
			b.Destoyed = append(b.Destoyed, types.Coord{X: x, Y: y})
		}
		y++
	}

	return b, nil
}

func getBoardSize(s string) (int, string, error) {
	for i, r := range s {
		if r == '\n' {
			bss := s[:i]
			bs, err := strconv.Atoi(bss)
			if err != nil {
				return 0, "", err
			}
			return bs, s[i:], nil
		}
	}
	return 0, "", fmt.Errorf("invalid input")
}
