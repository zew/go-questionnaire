package tpl

import (
	"testing"
)

func TestContained(t *testing.T) {

	tests := []struct {
		in   string
		want bool
	}{
		{
			in:   "./a/b/../../../etc/",
			want: false,
		},
		{
			in:   "./a/b/../../../etc/passwd",
			want: false,
		},
		{
			in:   "./a/b/../../etc/passwd",
			want: true,
		},
		{
			in:   "./a/../b/../etc/../../passwd",
			want: false,
		},
	}
	for idx, tt := range tests {
		got := contained(tt.in)
		if got != tt.want {
			t.Errorf("idx%2v: %-36v is %v should be %v", idx, tt.in, got, tt.want)
		} else {
			t.Logf("idx%2v: %-36v is %-5v indeed", idx, tt.in, got)
		}
	}
}
