package board

import (
	"fmt"
)

type fullBoard map[int]map[int]string

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
	b[0] = make(map[int]string)

	var row, col int
	for _, r := range board {
		if r == '\n' {
			continue
		}
		if col%dim == 0 {
			if col > 0 {
				row++
			}
			if row > dim-1 {
				break
			}
			col = 0
			b[row] = make(map[int]string)
		}
		b[row][col] = string(r)
		col++
	}

	return b, nil
}
