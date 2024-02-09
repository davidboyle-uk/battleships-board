package board

import (
	"fmt"
)

type fullBoard map[int][]string

func fullFromString(s string) (fullBoard, error) {

	dim, board, err := getBoardSize(s)
	if err != nil {
		return fullBoard{}, fmt.Errorf("unable to parse board")
	}

	l := len(s)
	max := (dim * dim)
	if l < max {
		return fullBoard{}, fmt.Errorf("board length is %v or more [%v]", max, l)
	}

	var b fullBoard = make(fullBoard)
	b[0] = []string{}

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

		b[x] = append(b[x], string(r))
		x++
	}

	return b, nil
}
