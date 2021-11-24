package qst

import (
	"testing"
)

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

func TestQuestionnaireT_LabelIsOutline(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		q    *QuestionnaireT
		args args
		want bool
	}{

		{"t1", &QuestionnaireT{}, args{"3     text text"}, true},
		{"t2", &QuestionnaireT{}, args{"2.    text text"}, true},
		{"t3", &QuestionnaireT{}, args{"2.)   text text"}, true},
		{"t4", &QuestionnaireT{}, args{"2b.   text text"}, true},
		{"t1", &QuestionnaireT{}, args{"2b.)  text text"}, true},
		{"t6", &QuestionnaireT{}, args{"2b.)\t\n\ttext text"}, true},
		{"t3", &QuestionnaireT{}, args{" 2.)   text text"}, false}, // leading space
		{"t3", &QuestionnaireT{}, args{"a2     text text"}, false}, // leading space

		{"t7", &QuestionnaireT{}, args{"b.)   text text"}, true},
		{"t8", &QuestionnaireT{}, args{"ab.)  text text"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.q.LabelIsOutline(tt.args.s); got != tt.want {
				t.Errorf("QuestionnaireT.LabelIsOutline() = %v, want %v", got, tt.want)
			}
		})
	}
}
