package trl

import (
	"fmt"
	"log"
	"reflect"
	"regexp"
	"testing"
)

// a primitve regex splitter - useful only for testing
var aAndPossiblyB = regexp.MustCompile(`ab?`)

func Test_regexSplit(t *testing.T) {

	type args struct {
		s  string
		re *regexp.Regexp
	}
	tests := []struct {
		name  string
		args  args
		want1 [][]int
		want2 string
	}{
		{
			name: "t1: ab?",
			want1: [][]int{
				{1, 3},
				{10, 12},
				{18, 19},
			},
			args: args{
				//   012345678901234567890
				//    _ _      _ _     __
				s:  `tablett  hablo espanol`,
				re: aAndPossiblyB,
			},
		},
		{
			name: "t2",
			want1: [][]int{
				{6, 14},
			},
			args: args{
				//   01234567890123456789
				//         _       _
				s:  `start for/ward stop`,
				re: pathSplitter,
			},
		},
		{
			name: "t3",
			want1: [][]int{
				{48, 61},
				{62, 70},
				{77, 86},
			},
			args: args{
				s:  `word <div style='font-size:90%' alt="attr" >div content</div> for/ward -and- back\ward eol`,
				re: pathSplitter,
			},
		},
		{
			name: "t4",
			want1: [][]int{
				{0, 8},
				{32, 41},
			},
			args: args{
				//  01234567890123456789
				//  _       _
				s: `for/ward 
				      -and- 
					back\ward`,
				re: pathSplitter,
			},
		},
	}
	for idx, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nm := tt.name
			if nm == "" {
				nm = fmt.Sprintf("t%v", idx)
			}
			got1 := regexSplit(tt.args.re, tt.args.s)
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("%10s - regexSplit() = \n\t got  %v\n\t want %v",
					nm,
					got1,
					tt.want1,
				)
			} else {
				log.Printf("%10s - is fine -  %v", nm, got1)
			}
			got2 := indexesToStringSlice(tt.args.s, got1)
			got2Str := ""
			for _, s := range got2 {
				got2Str += s
			}
			if !reflect.DeepEqual(got2Str, tt.args.s) {
				t.Errorf("%10s - recombine() = \n\t got  %v\n\t want %v",
					nm,
					got2Str,
					tt.args.s,
				)
			} else {
				log.Printf("%10s - recombined to %s", nm, got2)
			}
		})
	}

}
