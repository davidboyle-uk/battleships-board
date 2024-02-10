package ai

import (
	"reflect"
	"testing"

	"github.com/dbx123/battleships-board/pkg/board"
	"github.com/dbx123/battleships-board/types"
)

func TestCombineProbabilities(t *testing.T) {
	for name, tt := range map[string]struct {
		probabilities types.Probabilities
		expected      types.Probabilities
	}{
		"a": {
			probabilities: types.Probabilities{
				5: []types.Coord{{1, 6}},
				6: []types.Coord{{2, 7}, {4, 7}, {6, 7}, {7, 6}},
				7: []types.Coord{{3, 6}},
				8: []types.Coord{{2, 5}, {3, 6}, {4, 5}, {5, 6}, {6, 5}, {5, 6}},
			},
			expected: types.Probabilities{
				10: []types.Coord{{1, 6}},
				12: []types.Coord{{2, 7}, {4, 7}, {6, 7}, {7, 6}},
				16: []types.Coord{{2, 5}, {4, 5}, {6, 5}},
				22: []types.Coord{{3, 6}},
				23: []types.Coord{{3, 6}},
				24: []types.Coord{{5, 6}},
			},
		},
	} {
		t.Run(name, func(t *testing.T) {
			actual := combineProbabilities(tt.probabilities)
			if !reflect.DeepEqual(tt.expected, actual) {
				t.Fatalf("\nexpected\n%#v\ngot\n%#v\n", tt.expected, actual)
			}
		})
	}
}

func TestCalcProbabilities(t *testing.T) {
	for name, tt := range map[string]struct {
		boardString     string
		hitTrajectories types.Trajectories
		out             types.Probabilities
	}{
		"a": {
			boardString: `10
m-m-mm--h-
-m-mm--m-m
m-mm--m-m-
-hm--m-m--
mh--dmhhhh
-m-mmm--m-
m-h-h-h--m
-m-d---m--
m-m--m--m-
-m--m-m--h`,
			hitTrajectories: types.Trajectories{
				{1, 4}: types.Trajectory{1: []types.Coord{{1, 3}}},
			},
			out: types.Probabilities{7: []types.Coord{types.Coord{X: 1, Y: 2}}},
		},
		"b": {
			boardString: `10
----------
----------
----------
----------
----------
----------
--h-h-h---
----------
----------
----------`,
			hitTrajectories: types.Trajectories{
				{2, 6}: types.Trajectory{},
				{4, 6}: types.Trajectory{},
				{6, 6}: types.Trajectory{},
			},
			out: types.Probabilities{
				5: []types.Coord{{1, 6}},
				6: []types.Coord{{2, 7}, {4, 7}, {6, 7}, {7, 6}},
				7: []types.Coord{{3, 6}},
				8: []types.Coord{{2, 5}, {3, 6}, {4, 5}, {5, 6}, {6, 5}, {5, 6}}},
		},
	} {
		t.Run(name, func(t *testing.T) {
			b, err := board.BoardFromString(tt.boardString)
			if err != nil {
				t.Fatal(err)
			}

			actual := calcProbabilities(b, tt.hitTrajectories, types.NewFleet())
			if !reflect.DeepEqual(tt.out, actual) {
				t.Fatalf("\nexpected\n%#v\ngot\n%#v\n", tt.out, actual)
			}
		})
	}
}

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
		"b": {
			boardString: `10
m-m-mm--h-
-m-mm--m-m
m-mm--m-m-
-hm--m-m--
mh--dmhhhh
-m-mmm--m-
m-h-h-h--m
-m-d---m--
m-m--m--m-
-m--m-m--h`,
			out: types.Coord{6, 5},
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
				Coord: types.Coord{8, 0},
				MinX:  0,
				MaxX:  9,
				MinY:  0,
				MaxY:  1,
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
				Coord: types.Coord{4, 0},
				MinX:  0,
				MaxX:  9,
				MinY:  0,
				MaxY:  9,
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
