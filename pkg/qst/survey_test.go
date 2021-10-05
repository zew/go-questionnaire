package qst

import (
	"testing"

	"github.com/zew/go-questionnaire/pkg/cfg"
)

func Test_surveyT_WaveIDFuncs(t *testing.T) {

	cfg.LoadFakeConfigForTests()

	tests := []struct {
		name  string
		s     SurveyT
		want1 string
		want2 string
	}{
		{"t1", SurveyT{Year: 0, Month: 1}, "0000-01", ""},
		{"t2", SurveyT{Year: 2021, Month: 1}, "2021-01", "2021-01"},
		{"t3", SurveyT{Year: 2021, Month: 12}, "2021-12", "2021-12"},
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
		s    SurveyT
		args args
		want string
	}{
		// start of quarter - January
		{"t1a", SurveyT{Year: 0, Month: 1}, args{offs: []int{0}}, "Q1&nbsp;0"},
		{"t2a", SurveyT{Year: 2021, Month: 1}, args{offs: []int{0}}, "Q1&nbsp;2021"},
		{"t2aa", SurveyT{Year: 2021, Month: 1}, args{nil}, "Q1&nbsp;2021"},
		{"t3a", SurveyT{Year: 2021, Month: 1}, args{offs: []int{1}}, "Q2&nbsp;2021"},
		{"t4a", SurveyT{Year: 2021, Month: 1}, args{offs: []int{3}}, "Q4&nbsp;2021"},
		{"t5a", SurveyT{Year: 2021, Month: 1}, args{offs: []int{4}}, "Q1&nbsp;2022"},
		{"t6a", SurveyT{Year: 2021, Month: 1}, args{offs: []int{5}}, "Q2&nbsp;2022"},

		// end of quarter - March
		{"t1b", SurveyT{Year: 0, Month: 3}, args{offs: []int{0}}, "Q1&nbsp;0"},
		{"t2b", SurveyT{Year: 2021, Month: 3}, args{offs: []int{0}}, "Q1&nbsp;2021"},
		{"t3b", SurveyT{Year: 2021, Month: 3}, args{offs: []int{1}}, "Q2&nbsp;2021"},
		{"t4b", SurveyT{Year: 2021, Month: 3}, args{offs: []int{3}}, "Q4&nbsp;2021"},
		{"t5b", SurveyT{Year: 2021, Month: 3}, args{offs: []int{4}}, "Q1&nbsp;2022"},

		// start of quarter - April
		{"t2c", SurveyT{Year: 2021, Month: 4}, args{offs: []int{0}}, "Q2&nbsp;2021"},
		{"t3c", SurveyT{Year: 2021, Month: 4}, args{offs: []int{1}}, "Q3&nbsp;2021"},
		{"t4c", SurveyT{Year: 2021, Month: 4}, args{offs: []int{3}}, "Q1&nbsp;2022"}, // overflow
		{"t5c", SurveyT{Year: 2021, Month: 4}, args{offs: []int{4}}, "Q2&nbsp;2022"},

		// end of quarter - October - year overflow
		{"t6a", SurveyT{Year: 2021, Month: 10}, args{offs: []int{0}}, "Q4&nbsp;2021"},
		{"t7a", SurveyT{Year: 2021, Month: 10}, args{offs: []int{1}}, "Q1&nbsp;2022"},
		{"t8a", SurveyT{Year: 2021, Month: 10}, args{offs: []int{4}}, "Q4&nbsp;2022"},
		{"t9a", SurveyT{Year: 2021, Month: 10}, args{offs: []int{10}}, "Q2&nbsp;2024"}, // overflow several years

		// end of quarter - December - year overflow
		{"t6b", SurveyT{Year: 2021, Month: 12}, args{offs: []int{0}}, "Q4&nbsp;2021"},
		{"t7b", SurveyT{Year: 2021, Month: 12}, args{offs: []int{1}}, "Q1&nbsp;2022"},
		{"t8b", SurveyT{Year: 2021, Month: 12}, args{offs: []int{4}}, "Q4&nbsp;2022"},
		{"t9b", SurveyT{Year: 2021, Month: 12}, args{offs: []int{10}}, "Q2&nbsp;2024"}, // overflow several years

		// underflow
		{"tUf1", SurveyT{Year: 2021, Month: 1}, args{offs: []int{-1}}, "Q4&nbsp;2020"},
		{"tUf2", SurveyT{Year: 2021, Month: 3}, args{offs: []int{-1}}, "Q4&nbsp;2020"},
		{"tUf3", SurveyT{Year: 2021, Month: 4}, args{offs: []int{-1}}, "Q1&nbsp;2021"},
		{"tUf4", SurveyT{Year: 2021, Month: 1}, args{offs: []int{-4}}, "Q1&nbsp;2020"},
		{"tUf5", SurveyT{Year: 2021, Month: 1}, args{offs: []int{-5}}, "Q4&nbsp;2019"},

		// destatis parameter - quarter offset
		{"tDest1", SurveyT{Year: 2021, Month: 1, Params: []ParamT{{"destatis", "-1"}}}, args{offs: []int{0}}, "Q4&nbsp;2020"},
		{"tDest2", SurveyT{Year: 2021, Month: 3, Params: []ParamT{{"destatis", "-1"}}}, args{offs: []int{0}}, "Q4&nbsp;2020"},
		{"tDest3", SurveyT{Year: 2021, Month: 4, Params: []ParamT{{"destatis", "-1"}}}, args{offs: []int{0}}, "Q1&nbsp;2021"},
		{"tDest4", SurveyT{Year: 2021, Month: 10, Params: []ParamT{{"destatis", "-1"}}}, args{offs: []int{0}}, "Q3&nbsp;2021"},
		{"tDest5", SurveyT{Year: 2021, Month: 6, Params: []ParamT{{"destatis", "2"}}}, args{offs: []int{0}}, "Q4&nbsp;2021"},
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

func Test_surveyT_MonthOfQuarter(t *testing.T) {
	type args struct {
		offs []int
	}
	tests := []struct {
		name string
		s    SurveyT
		want int
	}{
		// start of quarter - January
		{"t1", SurveyT{Year: 2021, Month: 1}, 1},
		{"t2", SurveyT{Year: 2021, Month: 2}, 2},
		{"t3", SurveyT{Year: 2021, Month: 3}, 3},
		{"t4", SurveyT{Year: 2021, Month: 4}, 1},
		{"t5", SurveyT{Year: 2021, Month: 5}, 2},
		{"t6", SurveyT{Year: 2021, Month: 6}, 3},
		{"t7", SurveyT{Year: 2021, Month: 7}, 1},

		{"t10", SurveyT{Year: 2021, Month: 10}, 1},
		{"t11", SurveyT{Year: 2021, Month: 11}, 2},
		{"t12", SurveyT{Year: 2021, Month: 12}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.MonthOfQuarter(); got != tt.want {
				t.Errorf("surveyT.MonthOfQuarter() = %v, want %v", got, tt.want)
			}
		})
	}
}
