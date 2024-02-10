package types

import (
	"testing"
)

func TestToString(t *testing.T) {
	for name, tt := range map[string]struct {
		board    Board
		expected string
	}{
		"valid": {
			board: Board{
				Dim: 10,
				Moves: Moves{
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
			expected: `10
----------
----------
--------m-
----------
--d-----m-
---m-m----
------m---
-m---mmhh-
--m-------
------m---
`,
		},
	} {
		t.Run(name, func(t *testing.T) {
			actual := tt.board.ToString()
			if tt.expected != actual {
				t.Fatalf("\nexpected\n%#v\ngot\n%#v\n", tt.expected, actual)
			}
		})
	}
}
