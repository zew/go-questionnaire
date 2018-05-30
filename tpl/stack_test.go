package tpl

import (
	"fmt"
	"testing"
)

func TestStackT_Push(t *testing.T) {
	type args struct {
		pushee string
		expect *StackT
	}
	tests := []struct {
		name string
		sp   *StackT
		args args
	}{
		// {
		// 	name: "nil",
		// 	sp:   nil,
		// 	args: args{pushee: "tpl3.html", expect: &StackT{"tpl3.html"}},
		// },
		{
			name: "empty",
			sp:   &StackT{},
			args: args{pushee: "tpl3.html", expect: &StackT{"tpl3.html"}},
		},
		{
			name: "append full",
			sp:   &StackT{"tpl1.html", "tpl2.html"},
			args: args{pushee: "tpl3.html", expect: &StackT{"tpl1.html", "tpl2.html", "tpl3.html"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.sp.Push(tt.args.pushee)
			got := fmt.Sprintf("%+v", tt.sp)
			wnt := fmt.Sprintf("%+v", tt.args.expect)
			if got != wnt {
				t.Fatalf("\nwnt %v\ngot %v", wnt, got)
			}
		})
	}
}
