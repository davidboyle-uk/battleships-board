package game

import (
	"encoding/json"
	"testing"

	"github.com/davidboyle-uk/battleships-board/types"
)

func TestTakeShot(t *testing.T) {
	for name, tt := range map[string]struct {
		p1_before string
		p2_before string
		move      types.Coord
		p1_after  string
		p2_after  string
	}{
		"first hit": {
			p1_before: `{"Name":"DBX","Board":{"Dim":10},"Moves":{}}`,
			p2_before: `{"Name":"CPU","Board":{"Dim":10,"Moves":{"0 0":{"Ship":{"Coords":[{"X":0,"Y":0},{"X":0,"Y":1},{"X":0,"Y":2}],"Hits":0},"State":"-"},"0 1":{"Ship":{"Coords":[{"X":0,"Y":0},{"X":0,"Y":1},{"X":0,"Y":2}],"Hits":0},"State":"-"},"0 2":{"Ship":{"Coords":[{"X":0,"Y":0},{"X":0,"Y":1},{"X":0,"Y":2}],"Hits":0},"State":"-"}},"ShipTot":3},"Moves":{},"Hits":0}`,
			move:      types.Coord{0, 0},
			p1_after:  `{"Name":"DBX","Board":{"Dim":10,"Moves":null,"ShipTot":0},"Moves":{"0 0":{"Ship":null,"State":"h"}},"Hits":1}`,
			p2_after:  `{"Name":"CPU","Board":{"Dim":10,"Moves":{"0 0":{"Ship":{"Coords":[{"X":0,"Y":0},{"X":0,"Y":1},{"X":0,"Y":2}],"Hits":1},"State":"h"},"0 1":{"Ship":{"Coords":[{"X":0,"Y":0},{"X":0,"Y":1},{"X":0,"Y":2}],"Hits":1},"State":"-"},"0 2":{"Ship":{"Coords":[{"X":0,"Y":0},{"X":0,"Y":1},{"X":0,"Y":2}],"Hits":1},"State":"-"}},"ShipTot":3},"Moves":{},"Hits":0}`,
		},
		"second hit": {
			p1_before: `{"Name":"DBX","Board":{"Dim":10,"Moves":null,"ShipTot":0},"Moves":{"0 0":{"Ship":null,"State":"h"}},"Hits":1}`,
			p2_before: `{"Name":"CPU","Board":{"Dim":10,"Moves":{"0 0":{"Ship":{"Coords":[{"X":0,"Y":0},{"X":0,"Y":1},{"X":0,"Y":2}],"Hits":1},"State":"h"},"0 1":{"Ship":{"Coords":[{"X":0,"Y":0},{"X":0,"Y":1},{"X":0,"Y":2}],"Hits":1},"State":"-"},"0 2":{"Ship":{"Coords":[{"X":0,"Y":0},{"X":0,"Y":1},{"X":0,"Y":2}],"Hits":1},"State":"-"}},"ShipTot":3},"Moves":{},"Hits":0}`,
			move:      types.Coord{0, 1},
			p1_after:  `{"Name":"DBX","Board":{"Dim":10,"Moves":null,"ShipTot":0},"Moves":{"0 0":{"Ship":null,"State":"h"},"0 1":{"Ship":null,"State":"h"}},"Hits":2}`,
			p2_after:  `{"Name":"CPU","Board":{"Dim":10,"Moves":{"0 0":{"Ship":{"Coords":[{"X":0,"Y":0},{"X":0,"Y":1},{"X":0,"Y":2}],"Hits":2},"State":"h"},"0 1":{"Ship":{"Coords":[{"X":0,"Y":0},{"X":0,"Y":1},{"X":0,"Y":2}],"Hits":2},"State":"h"},"0 2":{"Ship":{"Coords":[{"X":0,"Y":0},{"X":0,"Y":1},{"X":0,"Y":2}],"Hits":2},"State":"-"}},"ShipTot":3},"Moves":{},"Hits":0}`,
		},
		"sunk": {
			p1_before: `{"Name":"DBX","Board":{"Dim":10,"Moves":null,"ShipTot":0},"Moves":{"0 0":{"Ship":null,"State":"h"},"0 1":{"Ship":null,"State":"h"}},"Hits":2}`,
			p2_before: `{"Name":"CPU","Board":{"Dim":10,"Moves":{"0 0":{"Ship":{"Coords":[{"X":0,"Y":0},{"X":0,"Y":1},{"X":0,"Y":2}],"Hits":2},"State":"h"},"0 1":{"Ship":{"Coords":[{"X":0,"Y":0},{"X":0,"Y":1},{"X":0,"Y":2}],"Hits":2},"State":"h"},"0 2":{"Ship":{"Coords":[{"X":0,"Y":0},{"X":0,"Y":1},{"X":0,"Y":2}],"Hits":2},"State":"-"}},"ShipTot":3},"Moves":{},"Hits":0}`,
			move:      types.Coord{0, 2},
			p1_after:  `{"Name":"DBX","Board":{"Dim":10,"Moves":null,"ShipTot":0},"Moves":{"0 0":{"Ship":null,"State":"d"},"0 1":{"Ship":null,"State":"d"},"0 2":{"Ship":null,"State":"d"}},"Hits":3}`,
			p2_after:  `{"Name":"CPU","Board":{"Dim":10,"Moves":{"0 0":{"Ship":{"Coords":[{"X":0,"Y":0},{"X":0,"Y":1},{"X":0,"Y":2}],"Hits":3},"State":"d"},"0 1":{"Ship":{"Coords":[{"X":0,"Y":0},{"X":0,"Y":1},{"X":0,"Y":2}],"Hits":3},"State":"d"},"0 2":{"Ship":{"Coords":[{"X":0,"Y":0},{"X":0,"Y":1},{"X":0,"Y":2}],"Hits":3},"State":"d"}},"ShipTot":3},"Moves":{},"Hits":0}`,
		},
	} {
		t.Run(name, func(t *testing.T) {
			p1 := types.Player{}
			err := json.Unmarshal([]byte(tt.p1_before), &p1)
			if err != nil {
				t.Fatal(err)
			}
			p2 := types.Player{}
			err = json.Unmarshal([]byte(tt.p2_before), &p2)
			if err != nil {
				t.Fatal(err)
			}
			TakeShot(&p1, &p2, tt.move)
			if err != nil {
				t.Fatal(err)
			}
			if tt.p1_after != toJSON(p1) {
				t.Fatalf("\nexpected\n%s\ngot\n%s\n", tt.p1_after, toJSON(p1))
			}
			if tt.p2_after != toJSON(p2) {
				t.Fatalf("\nexpected\n%s\ngot\n%s\n", tt.p2_after, toJSON(p2))
			}
		})
	}
}

func prettyPrint(data interface{}) string {
	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(b)
}

func toJSON(data interface{}) string {
	b, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	return string(b)
}
