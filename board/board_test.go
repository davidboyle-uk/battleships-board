package board

import (
	"reflect"
	"testing"

	"battleships-board/types"
)

func TestBoardFromString(t *testing.T) {
	for name, tt := range map[string]struct {
		in  string
		out types.Board
	}{
		"valid a": {
			in: `10
----------
----------
--------m-
----------
--d-----m-
---m-m----
------m---
-m---mmhh-
--m-------
------m---`,
			out: types.Board{
				Hit: []types.Coord{
					{7, 7},
					{7, 8},
				},
				Miss: []types.Coord{
					{2, 8},
					{4, 8},
					{5, 3},
					{5, 5},
					{6, 6},
					{7, 1},
					{7, 5},
					{7, 6},
					{8, 2},
					{9, 6},
				},
				Destoyed: []types.Coord{
					{X: 4, Y: 2},
				},
			},
		},
	} {
		t.Run(name, func(t *testing.T) {
			actual, err := boardFromString(tt.in)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(tt.out, actual) {
				t.Fatalf("\nexpected\n%#v\ngot\n%#v\n", tt.out, actual)
			}
		})
	}
}
