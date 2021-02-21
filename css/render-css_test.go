package css

import "testing"

func Test_boxStyle_CSS(t *testing.T) {
	tests := []struct {
		name  string
		csser CSSer
		want  string
	}{
		{
			name:  "box styles",
			csser: boxStyleExample1(),
			want:  boxStyleExample1Want(),
		},
		{
			name:  "text styles",
			csser: textStyleExample1(),
			want:  textStyleExample1Want(),
		},
		{
			name:  "grid container styles",
			csser: gridContainerStyleExample1(),
			want:  gridContainerStyleExample1Want(),
		},
		{
			name:  "grid item styles",
			csser: gridItemStyleExample1(),
			want:  gridItemStyleExample1Want(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.csser.CSS(); got != tt.want {
				t.Errorf("boxStyle.CSS() = \n%v, \nwant \n%v", got, tt.want)
			}
		})
	}
}
