package qst

import (
	"testing"

	"github.com/zew/go-questionnaire/cfg"
)

func Test_surveyT_WaveIDFuncs(t *testing.T) {

	cfg.LoadFakeConfigForTests()

	tests := []struct {
		name  string
		s     surveyT
		want1 string
		want2 string
	}{
		{"t1", surveyT{Year: 0, Month: 1}, "0000-01", ""},
		{"t2", surveyT{Year: 2021, Month: 1}, "2021-01", "2021-01"},
		{"t3", surveyT{Year: 2021, Month: 12}, "2021-12", "2021-12"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.WaveID(); got != tt.want1 {
				t.Errorf("surveyT.WaveID() = %v, want %v", got, tt.want1)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.WaveIDPretty(); got != tt.want2 {
				t.Errorf("surveyT.WaveIDPretty() = %v, want %v", got, tt.want2)
			}
		})
	}
}

func Test_surveyT_Quarter(t *testing.T) {
	type args struct {
		offs []int
	}
	tests := []struct {
		name string
		s    surveyT
		args args
		want string
	}{
		// start of quarter - January
		{"t1a", surveyT{Year: 0, Month: 1}, args{offs: []int{0}}, "Q1&nbsp;0"},
		{"t2a", surveyT{Year: 2021, Month: 1}, args{offs: []int{0}}, "Q1&nbsp;2021"},
		{"t2aa", surveyT{Year: 2021, Month: 1}, args{nil}, "Q1&nbsp;2021"},
		{"t3a", surveyT{Year: 2021, Month: 1}, args{offs: []int{1}}, "Q2&nbsp;2021"},
		{"t4a", surveyT{Year: 2021, Month: 1}, args{offs: []int{3}}, "Q4&nbsp;2021"},
		{"t5a", surveyT{Year: 2021, Month: 1}, args{offs: []int{4}}, "Q1&nbsp;2022"},
		{"t6a", surveyT{Year: 2021, Month: 1}, args{offs: []int{5}}, "Q2&nbsp;2022"},

		// end of quarter - March
		{"t1b", surveyT{Year: 0, Month: 3}, args{offs: []int{0}}, "Q1&nbsp;0"},
		{"t2b", surveyT{Year: 2021, Month: 3}, args{offs: []int{0}}, "Q1&nbsp;2021"},
		{"t3b", surveyT{Year: 2021, Month: 3}, args{offs: []int{1}}, "Q2&nbsp;2021"},
		{"t4b", surveyT{Year: 2021, Month: 3}, args{offs: []int{3}}, "Q4&nbsp;2021"},
		{"t5b", surveyT{Year: 2021, Month: 3}, args{offs: []int{4}}, "Q1&nbsp;2022"},

		// start of quarter - April
		{"t2c", surveyT{Year: 2021, Month: 4}, args{offs: []int{0}}, "Q2&nbsp;2021"},
		{"t3c", surveyT{Year: 2021, Month: 4}, args{offs: []int{1}}, "Q3&nbsp;2021"},
		{"t4c", surveyT{Year: 2021, Month: 4}, args{offs: []int{3}}, "Q1&nbsp;2022"}, // overflow
		{"t5c", surveyT{Year: 2021, Month: 4}, args{offs: []int{4}}, "Q2&nbsp;2022"},

		// end of quarter - October - year overflow
		{"t6a", surveyT{Year: 2021, Month: 10}, args{offs: []int{0}}, "Q4&nbsp;2021"},
		{"t7a", surveyT{Year: 2021, Month: 10}, args{offs: []int{1}}, "Q1&nbsp;2022"},
		{"t8a", surveyT{Year: 2021, Month: 10}, args{offs: []int{4}}, "Q4&nbsp;2022"},
		{"t9a", surveyT{Year: 2021, Month: 10}, args{offs: []int{10}}, "Q2&nbsp;2024"}, // overflow several years

		// end of quarter - December - year overflow
		{"t6b", surveyT{Year: 2021, Month: 12}, args{offs: []int{0}}, "Q4&nbsp;2021"},
		{"t7b", surveyT{Year: 2021, Month: 12}, args{offs: []int{1}}, "Q1&nbsp;2022"},
		{"t8b", surveyT{Year: 2021, Month: 12}, args{offs: []int{4}}, "Q4&nbsp;2022"},
		{"t9b", surveyT{Year: 2021, Month: 12}, args{offs: []int{10}}, "Q2&nbsp;2024"}, // overflow several years

		// underflow
		{"tUf1", surveyT{Year: 2021, Month: 1}, args{offs: []int{-1}}, "Q4&nbsp;2020"},
		{"tUf2", surveyT{Year: 2021, Month: 3}, args{offs: []int{-1}}, "Q4&nbsp;2020"},
		{"tUf3", surveyT{Year: 2021, Month: 4}, args{offs: []int{-1}}, "Q1&nbsp;2021"},
		{"tUf4", surveyT{Year: 2021, Month: 1}, args{offs: []int{-4}}, "Q1&nbsp;2020"},
		{"tUf5", surveyT{Year: 2021, Month: 1}, args{offs: []int{-5}}, "Q4&nbsp;2019"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Quarter(tt.args.offs...); got != tt.want {
				t.Errorf("surveyT.Quarter() = %v, want %v", got, tt.want)
			}

			// pretty output on verbose
			if false {
				offs := 0
				if len(tt.args.offs) > 0 {
					offs = tt.args.offs[0]
				}
				shift := ">>"
				if offs < 0 {
					shift = "<<"

				}
				t.Logf("%-10v of %4v -  %v%+3v Quarters   => %v",
					tt.s.Month, tt.s.Year, shift, offs, tt.s.Quarter(offs))
			}
		})
	}
}
