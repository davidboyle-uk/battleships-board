package board

import (
	"fmt"
	"strconv"

	"github.com/dbx123/battleships-board/types"
)

func BoardFromString(s string) (types.Board, error) {

	dim, board, err := getBoardSize(s)
	if err != nil {
		return types.Board{}, fmt.Errorf("unable to parse board")
	}

	l := len(s)
	max := (dim * dim)
	if l < max {
		return types.Board{}, fmt.Errorf("board length is %v or more [%v]", max, l)
	}

	b := types.Board{
		Dim:   dim,
		Moves: make(types.Moves),
	}

	var x, y int
	for _, r := range board {
		if r == '\n' {
			continue
		}
		if x%dim == 0 {
			if x > 0 {
				y++
			}
			if y > dim-1 {
				break
			}
			x = 0
		}

		s := types.State(r)
		if s != types.SEA {
			b.Moves[types.Coord{X: x, Y: y}] = types.CoordState{
				State: s,
			}
		}
		x++
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
