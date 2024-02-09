package types

import (
	"reflect"
	"testing"
)

func TestRemoveShip(t *testing.T) {
	for name, tt := range map[string]struct {
		fleet    Fleet
		toRemove ShipType
		expected Fleet
		err      bool
	}{
		"CARRIER": {
			fleet:    NewFleet(),
			toRemove: CARRIER,
			expected: Fleet{
				Flotilla{ShipType: 1, Num: 2},
				Flotilla{ShipType: 2, Num: 2},
				Flotilla{ShipType: 3, Num: 1},
				Flotilla{ShipType: 4, Num: 1},
			},
		},
		"DESTROYER": {
			fleet:    NewFleet(),
			toRemove: DESTROYER,
			expected: Fleet{
				Flotilla{ShipType: 1, Num: 2},
				Flotilla{ShipType: 2, Num: 1},
				Flotilla{ShipType: 3, Num: 1},
				Flotilla{ShipType: 4, Num: 1},
				Flotilla{ShipType: 5, Num: 1},
			},
		},
		"BATTLESHIP": {
			fleet:    NewFleet(),
			toRemove: BATTLESHIP,
			expected: Fleet{
				Flotilla{ShipType: 1, Num: 2},
				Flotilla{ShipType: 2, Num: 2},
				Flotilla{ShipType: 3, Num: 1},
				Flotilla{ShipType: 5, Num: 1},
			},
		},
		"SUBMARINE": {
			fleet:    NewFleet(),
			toRemove: SUBMARINE,
			expected: Fleet{
				Flotilla{ShipType: 1, Num: 1},
				Flotilla{ShipType: 2, Num: 2},
				Flotilla{ShipType: 3, Num: 1},
				Flotilla{ShipType: 4, Num: 1},
				Flotilla{ShipType: 5, Num: 1},
			},
		},
	} {
		t.Run(name, func(t *testing.T) {
			err := tt.fleet.RemoveShip(tt.toRemove)
			if err != nil && !tt.err {
				t.Fatalf("unexpected err %v", err)
			}
			if !reflect.DeepEqual(tt.expected, tt.fleet) {
				t.Fatalf("expected %#v\ngot %#v\n", tt.expected, tt.fleet)
			}
		})
	}
}
