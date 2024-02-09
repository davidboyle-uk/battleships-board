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
		"valid": {
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
				Dim: 10,
				Moves: types.Moves{
					{1, 7}: {State: 109},
					{2, 4}: {State: 100},
					{2, 8}: {State: 109},
					{3, 5}: {State: 109},
					{5, 5}: {State: 109},
					{5, 7}: {State: 109},
					{6, 6}: {State: 109},
					{6, 7}: {State: 109},
					{6, 9}: {State: 109},
					{7, 7}: {State: 104},
					{8, 2}: {State: 109},
					{8, 4}: {State: 109},
					{8, 7}: {State: 104},
				},
			},
		},
	} {
		t.Run(name, func(t *testing.T) {
			actual, err := BoardFromString(tt.in)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(tt.out, actual) {
				t.Fatalf("\nexpected\n%#v\ngot\n%#v\n", tt.out, actual)
			}
		})
	}
}
