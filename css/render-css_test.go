package css

import (
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
