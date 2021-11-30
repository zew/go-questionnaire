package qst

import (
	"strings"
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

func Test_cleansePrefixes(t *testing.T) {

	// log.SetFlags(log.Lshortfile)

	type args struct {
		ss []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"t1",
			args{[]string{
				"Your growth estimate? Ger",
				"Your growth estimate? Ger",
				"Your growth estimate? Ger   US",
				"Your growth estimate? Ger   US",
				"Your growth estimate? Ger   US   China",
				"Your growth estimate? Ger   US   China",
			}}, []string{
				"Your growth estimate? Ger",
				"Your growth estimate? Ger",
				"US",
				"US",
				"China",
				"China",
			},
		},
		{
			"t2",
			args{[]string{
				"Are you Alice?        Yes",
				"Are you Alice?        Yes   No",
				"Are you Alice?        Yes   No   Perhaps",
			}}, []string{
				"Are you Alice?        Yes",
				"No",
				"Perhaps",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := cleansePrefixes(tt.args.ss)
			gots := strings.Join(got, "|")
			wnts := strings.Join(tt.want, "|")
			if gots != wnts {
				t.Errorf("cleansePrefixes()\n\tgot: %v  \n\twnt: %v", gots, wnts)
			}
		})
	}
}
