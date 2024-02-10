package ships

import (
	"reflect"
	"testing"

	"battleships-board/types"
)

func TestRandCoord(t *testing.T) {
	c := randCoord(10)
	if (c != types.Coord{}) {
		t.Fatalf("%#v", c)
	}
}

func TestPrintInPlayShips(t *testing.T) {
	ships := GenerateShips(10)
	t.Fatalf("%v\n%s", ships, ships.AsString())
}

func TestExpandShip(t *testing.T) {
	for name, tt := range map[string]struct {
		c         types.Coord
		boardSize int
		t         types.ShipType
		dir       types.ShipDirection
		out       types.Ship
		err       bool
	}{
		"carrier h": {
			c:         types.Coord{0, 0},
			boardSize: 10,
			t:         types.CARRIER,
			dir:       types.HORIZONTAL,
			out: types.Ship{
				Coords: []types.Coord{
					{0, 0},
					{1, 0},
					{2, 0},
					{3, 0},
					{4, 0},
				},
			},
		},
		"carrier v": {
			c:         types.Coord{0, 0},
			boardSize: 10,
			t:         types.CARRIER,
			dir:       types.VERTICAL,
			out: types.Ship{
				Coords: []types.Coord{
					{0, 0},
					{0, 1},
					{0, 2},
					{0, 3},
					{0, 4},
				},
			},
		},
	} {
		t.Run(name, func(t *testing.T) {
			s, err := expandShip(tt.c, tt.boardSize, tt.t, tt.dir)
			if err != nil && !tt.err {
				t.Fatalf("unexpected error %v", err)
			}
			if !reflect.DeepEqual(s, tt.out) {
				t.Fatalf("expected %#v, got %#v", tt.out, s)
			}
		})
	}
}

func TestGenerateShips(t *testing.T) {
	ships := GenerateShips(10)
	if len(ships) != 7 {
		t.Fatalf("%#v", len(ships))
	}
}

func TestHasCollisions(t *testing.T) {
	for name, tt := range map[string]struct {
		ship     types.Ship
		inPlay   types.Ships
		expected bool
	}{
		"does": {
			ship: types.Ship{
				Coords: []types.Coord{
					{0, 1},
				},
			},
			inPlay: types.Ships{
				types.Ship{
					Coords: []types.Coord{
						{0, 0},
						{0, 1},
						{0, 2},
						{0, 3},
						{0, 4},
						{0, 5},
					},
				},
			},
			expected: true,
		},
		"doesnt": {
			ship: types.Ship{
				Coords: []types.Coord{
					{6, 1},
				},
			},
			inPlay: types.Ships{
				types.Ship{
					Coords: []types.Coord{
						{0, 0},
						{0, 1},
						{0, 2},
						{0, 3},
						{0, 4},
						{0, 5},
					},
				},
			},
			expected: false,
		},
	} {
		t.Run(name, func(t *testing.T) {
			actual := hasCollisions(tt.ship, tt.inPlay)
			if tt.expected != actual {
				t.Fatalf("expected %#v\ngot %#v\n", tt.expected, actual)
			}
		})
	}
}
