package types

import (
	"encoding/json"
	"reflect"
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
					"1 7": {State: "m"},
					"2 4": {State: "d"},
					"2 8": {State: "m"},
					"3 5": {State: "m"},
					"5 5": {State: "m"},
					"5 7": {State: "m"},
					"6 6": {State: "m"},
					"6 7": {State: "m"},
					"6 9": {State: "m"},
					"7 7": {State: "h"},
					"8 2": {State: "m"},
					"8 4": {State: "m"},
					"8 7": {State: "h"},
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

func TestBoardToStringFromJSON(t *testing.T) {
	for name, tt := range map[string]struct {
		boardJSON string
		expected  string
	}{
		"CPU": {
			boardJSON: `{"Dim":10,"Moves":{"0 0":{"Ship":{"Coords":[{"X":9,"Y":5},{"X":9,"Y":6},{"X":9,"Y":7},{"X":9,"Y":8},{"X":9,"Y":9}],"Hits":1},"State":"h"},"0 9":{"Ship":{"Coords":[{"X":9,"Y":5},{"X":9,"Y":6},{"X":9,"Y":7},{"X":9,"Y":8},{"X":9,"Y":9}],"Hits":0},"State":"-"},"1 0":{"Ship":{"Coords":[{"X":9,"Y":5},{"X":9,"Y":6},{"X":9,"Y":7},{"X":9,"Y":8},{"X":9,"Y":9}],"Hits":1},"State":"h"},"2 0":{"Ship":{"Coords":[{"X":9,"Y":5},{"X":9,"Y":6},{"X":9,"Y":7},{"X":9,"Y":8},{"X":9,"Y":9}],"Hits":1},"State":"h"},"3 0":{"Ship":{"Coords":[{"X":9,"Y":5},{"X":9,"Y":6},{"X":9,"Y":7},{"X":9,"Y":8},{"X":9,"Y":9}],"Hits":1},"State":"h"},"5 2":{"Ship":{"Coords":[{"X":9,"Y":5},{"X":9,"Y":6},{"X":9,"Y":7},{"X":9,"Y":8},{"X":9,"Y":9}],"Hits":0},"State":"-"},"5 5":{"Ship":{"Coords":[{"X":9,"Y":5},{"X":9,"Y":6},{"X":9,"Y":7},{"X":9,"Y":8},{"X":9,"Y":9}],"Hits":1},"State":"h"},"5 6":{"Ship":{"Coords":[{"X":9,"Y":5},{"X":9,"Y":6},{"X":9,"Y":7},{"X":9,"Y":8},{"X":9,"Y":9}],"Hits":1},"State":"h"},"5 7":{"Ship":{"Coords":[{"X":9,"Y":5},{"X":9,"Y":6},{"X":9,"Y":7},{"X":9,"Y":8},{"X":9,"Y":9}],"Hits":0},"State":"-"},"6 5":{"Ship":{"Coords":[{"X":9,"Y":5},{"X":9,"Y":6},{"X":9,"Y":7},{"X":9,"Y":8},{"X":9,"Y":9}],"Hits":0},"State":"-"},"7 5":{"Ship":{"Coords":[{"X":9,"Y":5},{"X":9,"Y":6},{"X":9,"Y":7},{"X":9,"Y":8},{"X":9,"Y":9}],"Hits":0},"State":"-"},"7 9":{"Ship":{"Coords":[{"X":9,"Y":5},{"X":9,"Y":6},{"X":9,"Y":7},{"X":9,"Y":8},{"X":9,"Y":9}],"Hits":0},"State":"-"},"8 9":{"Ship":{"Coords":[{"X":9,"Y":5},{"X":9,"Y":6},{"X":9,"Y":7},{"X":9,"Y":8},{"X":9,"Y":9}],"Hits":0},"State":"-"},"9 5":{"Ship":{"Coords":[{"X":9,"Y":5},{"X":9,"Y":6},{"X":9,"Y":7},{"X":9,"Y":8},{"X":9,"Y":9}],"Hits":0},"State":"-"},"9 6":{"Ship":{"Coords":[{"X":9,"Y":5},{"X":9,"Y":6},{"X":9,"Y":7},{"X":9,"Y":8},{"X":9,"Y":9}],"Hits":0},"State":"-"},"9 7":{"Ship":{"Coords":[{"X":9,"Y":5},{"X":9,"Y":6},{"X":9,"Y":7},{"X":9,"Y":8},{"X":9,"Y":9}],"Hits":0},"State":"-"},"9 8":{"Ship":{"Coords":[{"X":9,"Y":5},{"X":9,"Y":6},{"X":9,"Y":7},{"X":9,"Y":8},{"X":9,"Y":9}],"Hits":0},"State":"-"},"9 9":{"Ship":{"Coords":[{"X":9,"Y":5},{"X":9,"Y":6},{"X":9,"Y":7},{"X":9,"Y":8},{"X":9,"Y":9}],"Hits":0},"State":"-"}},"ShipTot":18}`,
			expected: `10
hhhh------
----------
----------
----------
----------
-----h----
-----h----
----------
----------
----------
`,
		},
		"DBX": {
			boardJSON: `{"Dim":10,"Moves":{"1 1":{"Ship":{"Coords":[{"X":1,"Y":1},{"X":1,"Y":2},{"X":1,"Y":3},{"X":1,"Y":4},{"X":1,"Y":5}],"Hits":1},"State":"h"},"1 2":{"Ship":{"Coords":[{"X":1,"Y":1},{"X":1,"Y":2},{"X":1,"Y":3},{"X":1,"Y":4},{"X":1,"Y":5}],"Hits":1},"State":"h"},"1 3":{"Ship":{"Coords":[{"X":1,"Y":1},{"X":1,"Y":2},{"X":1,"Y":3},{"X":1,"Y":4},{"X":1,"Y":5}],"Hits":1},"State":"h"},"1 4":{"Ship":{"Coords":[{"X":1,"Y":1},{"X":1,"Y":2},{"X":1,"Y":3},{"X":1,"Y":4},{"X":1,"Y":5}],"Hits":1},"State":"h"},"1 5":{"Ship":{"Coords":[{"X":1,"Y":1},{"X":1,"Y":2},{"X":1,"Y":3},{"X":1,"Y":4},{"X":1,"Y":5}],"Hits":1},"State":"h"},"3 0":{"Ship":{"Coords":[{"X":1,"Y":1},{"X":1,"Y":2},{"X":1,"Y":3},{"X":1,"Y":4},{"X":1,"Y":5}],"Hits":0},"State":"-"},"3 1":{"Ship":{"Coords":[{"X":1,"Y":1},{"X":1,"Y":2},{"X":1,"Y":3},{"X":1,"Y":4},{"X":1,"Y":5}],"Hits":0},"State":"-"},"3 4":{"Ship":{"Coords":[{"X":1,"Y":1},{"X":1,"Y":2},{"X":1,"Y":3},{"X":1,"Y":4},{"X":1,"Y":5}],"Hits":1},"State":"h"},"4 0":{"Ship":{"Coords":[{"X":1,"Y":1},{"X":1,"Y":2},{"X":1,"Y":3},{"X":1,"Y":4},{"X":1,"Y":5}],"Hits":0},"State":"-"},"4 4":{"Ship":{"Coords":[{"X":1,"Y":1},{"X":1,"Y":2},{"X":1,"Y":3},{"X":1,"Y":4},{"X":1,"Y":5}],"Hits":1},"State":"h"},"5 0":{"Ship":{"Coords":[{"X":1,"Y":1},{"X":1,"Y":2},{"X":1,"Y":3},{"X":1,"Y":4},{"X":1,"Y":5}],"Hits":0},"State":"-"},"5 3":{"Ship":{"Coords":[{"X":1,"Y":1},{"X":1,"Y":2},{"X":1,"Y":3},{"X":1,"Y":4},{"X":1,"Y":5}],"Hits":0},"State":"-"},"6 0":{"Ship":{"Coords":[{"X":1,"Y":1},{"X":1,"Y":2},{"X":1,"Y":3},{"X":1,"Y":4},{"X":1,"Y":5}],"Hits":0},"State":"-"},"6 2":{"Ship":{"Coords":[{"X":1,"Y":1},{"X":1,"Y":2},{"X":1,"Y":3},{"X":1,"Y":4},{"X":1,"Y":5}],"Hits":1},"State":"h"},"6 3":{"Ship":{"Coords":[{"X":1,"Y":1},{"X":1,"Y":2},{"X":1,"Y":3},{"X":1,"Y":4},{"X":1,"Y":5}],"Hits":1},"State":"h"},"9 1":{"Ship":{"Coords":[{"X":1,"Y":1},{"X":1,"Y":2},{"X":1,"Y":3},{"X":1,"Y":4},{"X":1,"Y":5}],"Hits":0},"State":"-"},"9 2":{"Ship":{"Coords":[{"X":1,"Y":1},{"X":1,"Y":2},{"X":1,"Y":3},{"X":1,"Y":4},{"X":1,"Y":5}],"Hits":0},"State":"-"},"9 3":{"Ship":{"Coords":[{"X":1,"Y":1},{"X":1,"Y":2},{"X":1,"Y":3},{"X":1,"Y":4},{"X":1,"Y":5}],"Hits":0},"State":"-"}},"ShipTot":18}`,
			expected: `10
----------
-h--------
-h----h---
-h----h---
-h-hh-----
-h--------
----------
----------
----------
----------
`,
		},
	} {
		t.Run(name, func(t *testing.T) {
			var b Board
			err := json.Unmarshal([]byte(tt.boardJSON), &b)
			if err != nil {
				t.Fatal(err)
			}
			actual := b.ToString()
			if !reflect.DeepEqual(tt.expected, actual) {
				t.Fatalf("\nexpected\n%s\ngot\n%s\n", tt.expected, actual)
			}
		})
	}
}
