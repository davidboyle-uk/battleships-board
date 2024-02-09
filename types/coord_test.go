package types

import (
	"reflect"
	"testing"
)

func TestAddNum(t *testing.T) {
	for name, tt := range map[string]struct {
		origin    Coord
		direction Direction
		num       int
		expected  Coord
	}{
		"UP": {
			origin:    Coord{0, 0},
			direction: UP,
			num:       5,
			expected:  Coord{0, 5},
		},
	} {
		t.Run(name, func(t *testing.T) {
			actual := tt.origin.AddNum(PossibleMoves[tt.direction], tt.num)
			if !reflect.DeepEqual(tt.expected, actual) {
				t.Fatalf("expected %#v\ngot %#v\n", tt.expected, actual)
			}
		})
	}
}
