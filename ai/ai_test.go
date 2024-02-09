package ai

import (
	"reflect"
	"testing"

	"battleships-board/board"
	"battleships-board/types"
)

func TestCalcMoveBasedOnProbability(t *testing.T) {
	for name, tt := range map[string]struct {
		boardString string
		out         types.Coord
	}{
		"a": {
			boardString: `10
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
			out: types.Coord{9, 7},
		},
	} {
		t.Run(name, func(t *testing.T) {
			b, err := board.BoardFromString(tt.boardString)
			if err != nil {
				t.Fatal(err)
			}

			actual := calcMoveBasedOnProbability(b)
			if !reflect.DeepEqual(tt.out, actual) {
				t.Fatalf("\nexpected\n%#v\ngot\n%#v\n", tt.out, actual)
			}
		})
	}
}

func TestTraceTrajectories(t *testing.T) {
	for name, tt := range map[string]struct {
		boardString string
		out         types.Trajectories
	}{
		"a": {
			boardString: `10
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
			out: types.Trajectories{
				{7, 7}: types.Trajectory{
					3: []types.Coord{
						{8, 7},
					},
				},
				{8, 7}: types.Trajectory{
					2: []types.Coord{
						{7, 7},
					},
				},
			},
		},
	} {
		t.Run(name, func(t *testing.T) {
			b, err := board.BoardFromString(tt.boardString)
			if err != nil {
				t.Fatal(err)
			}

			hitsOnBoard := b.GetHits()
			actual := traceTrajectories(b, hitsOnBoard)
			if !reflect.DeepEqual(tt.out, actual) {
				t.Fatalf("\nexpected\n%#v\ngot\n%#v\n", tt.out, actual)
			}
		})
	}
}

func TestCalcMoveWhenNoHits(t *testing.T) {
	for name, tt := range map[string]struct {
		boardString string
		out         types.Coord
	}{
		"empty": {
			boardString: `10
----------
----------
----------
----------
----------
----------
----------
----------
----------
----------`,
			out: types.Coord{4, 4},
		},
		"a": {
			boardString: `10
----------
----------
--------m-
----------
--d-----m-
---m-m----
------m---
-m---mmmm-
--m-------
------m---`,
			out: types.Coord{4, 3},
		},
	} {
		t.Run(name, func(t *testing.T) {
			b, err := board.BoardFromString(tt.boardString)
			if err != nil {
				t.Fatal(err)
			}

			actual := calcMoveWhenNoHits(b)
			if !reflect.DeepEqual(tt.out, actual) {
				t.Fatalf("\nexpected\n%#v\ngot\n%#v\n", tt.out, actual)
			}
		})
	}
}

func TestCalcCoordSpace(t *testing.T) {
	for name, tt := range map[string]struct {
		boardString string
		target      types.Coord
		out         types.CoordSpace
	}{
		"0,0": {
			boardString: `10
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
			out: types.CoordSpace{
				MinX: 0,
				MaxX: 9,
				MinY: 0,
				MaxY: 9,
			},
		},
		"8,0": {
			boardString: `10
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
			target: types.Coord{8, 0},
			out: types.CoordSpace{
				MinX: 0,
				MaxX: 9,
				MinY: 0,
				MaxY: 1,
			},
		},
		"4,0": {
			boardString: `10
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
			target: types.Coord{4, 0},
			out: types.CoordSpace{
				MinX: 0,
				MaxX: 9,
				MinY: 0,
				MaxY: 9,
			},
		},
	} {
		t.Run(name, func(t *testing.T) {
			b, err := board.BoardFromString(tt.boardString)
			if err != nil {
				t.Fatal(err)
			}
			actual := calcCoordSpace(tt.target, b)
			if !reflect.DeepEqual(tt.out, actual) {
				t.Fatalf("\nexpected\n%#v\ngot\n%#v\n", tt.out, actual)
			}
		})
	}
}
