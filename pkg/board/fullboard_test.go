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
				0: []string{"-", "-", "-", "-", "-", "-", "-", "-", "-", "-"},
				1: []string{"-", "-", "-", "-", "-", "-", "-", "m", "-", "-"},
				2: []string{"-", "-", "-", "-", "d", "-", "-", "-", "m", "-"},
				3: []string{"-", "-", "-", "-", "-", "m", "-", "-", "-", "-"},
				4: []string{"-", "-", "-", "-", "-", "-", "-", "-", "-", "-"},
				5: []string{"-", "-", "-", "-", "-", "m", "-", "m", "-", "-"},
				6: []string{"-", "-", "-", "-", "-", "-", "m", "m", "-", "m"},
				7: []string{"-", "-", "-", "-", "-", "-", "-", "h", "-", "-"},
				8: []string{"-", "-", "m", "-", "m", "-", "-", "h", "-", "-"},
				9: []string{"-", "-", "-", "-", "-", "-", "-", "-", "-", "-"},
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
