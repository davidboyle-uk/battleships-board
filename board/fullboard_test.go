package board

import (
	"reflect"
	"testing"
)

func TestFullBoardFromString(t *testing.T) {
	for name, tt := range map[string]struct {
		in  string
		out fullBoard
	}{
		"valid a": {
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
			out: fullBoard{
				0: map[int]string{0: "-", 1: "-", 2: "-", 3: "-", 4: "-", 5: "-", 6: "-", 7: "-", 8: "-", 9: "-"},
				1: map[int]string{0: "-", 1: "-", 2: "-", 3: "-", 4: "-", 5: "-", 6: "-", 7: "-", 8: "-", 9: "-"},
				2: map[int]string{0: "-", 1: "-", 2: "-", 3: "-", 4: "-", 5: "-", 6: "-", 7: "-", 8: "m", 9: "-"},
				3: map[int]string{0: "-", 1: "-", 2: "-", 3: "-", 4: "-", 5: "-", 6: "-", 7: "-", 8: "-", 9: "-"},
				4: map[int]string{0: "-", 1: "-", 2: "d", 3: "-", 4: "-", 5: "-", 6: "-", 7: "-", 8: "m", 9: "-"},
				5: map[int]string{0: "-", 1: "-", 2: "-", 3: "m", 4: "-", 5: "m", 6: "-", 7: "-", 8: "-", 9: "-"},
				6: map[int]string{0: "-", 1: "-", 2: "-", 3: "-", 4: "-", 5: "-", 6: "m", 7: "-", 8: "-", 9: "-"},
				7: map[int]string{0: "-", 1: "m", 2: "-", 3: "-", 4: "-", 5: "m", 6: "m", 7: "h", 8: "h", 9: "-"},
				8: map[int]string{0: "-", 1: "-", 2: "m", 3: "-", 4: "-", 5: "-", 6: "-", 7: "-", 8: "-", 9: "-"},
				9: map[int]string{0: "-", 1: "-", 2: "-", 3: "-", 4: "-", 5: "-", 6: "m", 7: "-", 8: "-", 9: "-"},
			},
		},
	} {
		t.Run(name, func(t *testing.T) {
			actual, err := fullFromString(tt.in)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(tt.out, actual) {
				t.Fatalf("\nexpected\n%#v\ngot\n%#v\n", tt.out, actual)
			}
		})
	}
}
