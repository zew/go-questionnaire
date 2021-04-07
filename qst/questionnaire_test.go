package qst

import "testing"

func TestDelocalizeNumber(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"t1", args{"123"}, "123"},
		{"t2", args{"123.45"}, "123.45"},
		{"t3", args{"123,45"}, "123.45"},
		{"t4", args{"123,456.78"}, "123456.78"},
		{"t5", args{"123.456,78"}, "123456.78"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DelocalizeNumber(tt.args.s); got != tt.want {
				t.Errorf("DelocalizeNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
