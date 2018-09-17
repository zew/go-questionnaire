package main

import (
	"reflect"
	"testing"
)

func TestSuperset(t *testing.T) {
	type args struct {
		keys [][]string
	}
	tests := []struct {
		name         string
		args         [][]string
		wantSuperset []string
	}{
		{
			name: "one longer at the end - one in between",
			args: [][]string{
				[]string{"a1", "a2", "a3"},
				[]string{"a1", "a2", "a3", "a4"},
				[]string{"a1", "a3"},
				[]string{"a1", "a12", "a3"},
				[]string{"a1", "a2", "a21"},
				[]string{"a3", "a31", "a32", "a1", "a13"},
				[]string{"a5"},
			},
			wantSuperset: []string{
				"a1", "a13", "a12", // imperfect, but as good as it gets
				"a2", "a21",
				"a3", "a31", "a32",
				"a4",
				"a5",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSuperset := Superset(tt.args); !reflect.DeepEqual(gotSuperset, tt.wantSuperset) {
				t.Errorf("Superset() = \ngot %v, \nwnt %v", gotSuperset, tt.wantSuperset)
			}
		})
	}
}
