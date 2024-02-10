package board

import (
	"reflect"
	"testing"

	"github.com/dbx123/battleships-board/types"
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
					{1, 7}: {State: "m"},
					{2, 4}: {State: "d"},
					{2, 8}: {State: "m"},
					{3, 5}: {State: "m"},
					{5, 5}: {State: "m"},
					{5, 7}: {State: "m"},
					{6, 6}: {State: "m"},
					{6, 7}: {State: "m"},
					{6, 9}: {State: "m"},
					{7, 7}: {State: "h"},
					{8, 2}: {State: "m"},
					{8, 4}: {State: "m"},
					{8, 7}: {State: "h"},
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
