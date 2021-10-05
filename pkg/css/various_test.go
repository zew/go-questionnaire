package css

import (
	"fmt"
	"testing"
)

func Test_renderStyles(t *testing.T) {
	tests := []struct {
		name  string
		csser CSSerSimple
		want  string
	}{
		{
			name:  "box-style-1",
			csser: boxStyleExample1(),
			want:  boxStyleExample1Want(),
		},
		{
			name:  "text-style-1",
			csser: textStyleExample1(),
			want:  textStyleExample1Want(),
		},
		{
			name:  "grid-container-style-1",
			csser: gridContainerStyleExample1(),
			want:  gridContainerStyleExample1Want(),
		},
		{
			name:  "grid-item-style-1",
			csser: gridItemStyleExample1(),
			want:  gridItemStyleExample1Want(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.csser.CSS(); got != tt.want {
				t.Errorf("csser.CSS() = \n%v, \nwant \n%v", got, tt.want)
			}
		})
	}
}

func Test_renderResponsive(t *testing.T) {
	tests := []struct {
		name  string
		csser CSSer
		want  string
	}{
		{
			name:  "test-1",
			csser: stylesResponsiveExample(),
			want:  stylesResponsiveExampleWant("test-1"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.csser.CSS(tt.name); got != tt.want {
				t.Errorf("responsive.CSS() = \n%v, \nwant \n%v", got, tt.want)
			}
		})
	}
}

func Test_Combine(t *testing.T) {

	tests := []struct {
		name  string
		sr, b *StyleGridContainer
		want  *StyleGridContainer
	}{
		{
			name: "test-1",
			sr:   &StyleGridContainer{AutoColumns: "aa", JustifyContent: "dd"},
			b:    &StyleGridContainer{AutoFlow: "bb", TemplateRows: "cc", JustifyContent: "ee"},
			want: &StyleGridContainer{AutoColumns: "aa", AutoFlow: "bb", TemplateRows: "cc", JustifyContent: "dd"},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				tt.sr.Combine(*tt.b)
				got := fmt.Sprintf("%+v", tt.sr)
				want1 := fmt.Sprintf("%+v", tt.want)
				t.Logf("\n%v\n%v", got, want1)
				if got != want1 {
					t.Errorf("sr.Combine() got \n%v\nwant \n%v", got, want1)
				}
			},
		)
	}
}
func Test_CombineAll(t *testing.T) {

	tests := []struct {
		name  string
		sr, b *StylesResponsive
		want  *StylesResponsive
	}{
		{
			name: "test-1",
			sr:   &StylesResponsive{Desktop: Styles{StyleGridContainer: StyleGridContainer{AutoColumns: "aa"}}},
			b:    &StylesResponsive{Desktop: Styles{StyleGridContainer: StyleGridContainer{AutoFlow: "bb", TemplateColumns: "cc"}}},
			want: &StylesResponsive{Desktop: Styles{StyleGridContainer: StyleGridContainer{AutoColumns: "aa", AutoFlow: "bb", TemplateColumns: "cc"}}},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				tt.sr.Combine(*tt.b)
				got := fmt.Sprintf("%+v", tt.sr)
				want1 := fmt.Sprintf("%+v", tt.want)
				// t.Logf("\n%v\n%v", got, want1)
				if got != want1 {
					t.Errorf("sr.Combine() got \n%v\nwant \n%v", got, want1)
				}
			},
		)
	}
}
