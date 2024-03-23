package ai

import (
	"reflect"
	"testing"

	"github.com/davidboyle-uk/battleships-board/pkg/board"
	"github.com/davidboyle-uk/battleships-board/types"
)

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
			out: types.Trajectories{
				{2, 6}: types.Trajectory{},
				{4, 6}: types.Trajectory{},
				{6, 6}: types.Trajectory{},
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
