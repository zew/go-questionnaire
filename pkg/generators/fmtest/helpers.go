package fmtest

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

func addEchartsExamplePage(q *qst.QuestionnaireT) {

	{
		page := q.AddPage()
		page.Label = trl.S{
			"de": "",
			"en": "",
		}
		page.WidthMax("calc(100% - 1.2rem)")
		page.WidthMax("40rem")
		{
			gr := page.AddGroup()
			gr.Cols = 1
			{
				inp := gr.AddInput()
				inp.Type = "dyn-textblock"
				inp.DynamicFunc = "RenderStaticContent"
				inp.DynamicFuncParamset = "/echart/inner.html"
				inp.ColSpan = 1
				inp.ColSpanLabel = 1
			}
			{
				inp := gr.AddInput()
				inp.Type = "text"
				inp.Name = "chart_output"
				inp.MaxChars = 40
				inp.ColSpan = 1
				inp.ColSpanLabel = 1
				inp.ColSpanControl = 1
			}
		}
	}

}

func colTemplateWithFreeRow() ([]float32, []float32, string) {

	var columnTemplateLocal = []float32{
		3.6, 1, // separated - see below
		0.0, 1,
		0.0, 1,
		0.0, 1,
		0.0, 1,
		0.4, 1,
	}
	// additional row below each block
	colsBelow1 := append([]float32{1.0}, columnTemplateLocal...)
	colsBelow1 = []float32{
		1.4, 2.2, //   3.6, 1 => 4.6 - separated to two cols - part 1
		0.0, 1, //     3.6, 1 => 4.6 - separated to two cols - part 2
		0.0, 1,
		0.0, 1,
		0.0, 1,
		0.0, 1,
		0.4, 1,
	}
	colsBelow2 := []float32{}
	for i := 0; i < len(colsBelow1); i += 2 {
		colsBelow2 = append(colsBelow2, colsBelow1[i]+colsBelow1[i+1])
	}
	stl := ""
	for colIdx := 0; colIdx < len(colsBelow2); colIdx++ {
		stl = fmt.Sprintf(
			"%v   %vfr ",
			stl,
			colsBelow2[colIdx],
		)
	}

	return columnTemplateLocal, colsBelow1, stl

}

func labelsGoodBad() []trl.S {

	tm := []trl.S{
		{
			"de": "gut",
			"en": "good",
		},
		{
			"de": "normal",
			"en": "normal",
		},
		{
			"de": "schlecht",
			"en": "bad",
		},
		{
			"de": "keine<br>Angabe",
			"en": "no estimate",
		},
	}

	return tm

}

func labelsImproveDeteriorate() []trl.S {

	tm := []trl.S{
		{
			"de": "verbessern",
			"en": "improve",
		},
		{
			"de": "nicht<br>verändern",
			"en": "not change",
		},
		{
			"de": "verschlechtern",
			"en": "worsen",
		},
		{
			"de": "keine<br>Angabe",
			"en": "no estimate",
		},
	}

	return tm

}

func labelsImproveDeteriorateSectoral() []trl.S {

	tm := []trl.S{
		{
			"de": "besser",
			"en": "improve",
		},
		{
			"de": "gleich<br>bleiben",
			"en": "not change",
		},
		{
			"de": "schlechter",
			"en": "worsen",
		},
		{
			"de": "keine<br>Angabe",
			"en": "no estimate",
		},
	}

	return tm

}

func labelsIncreaseDecrease() []trl.S {

	tm := []trl.S{
		{
			"de": "steigen",
			"en": "increase",
		},
		{
			"de": "gleich<br>bleiben",
			"en": "not change",
		},
		{
			"de": "sinken",
			"en": "decrease",
		},
		{
			"de": "keine<br>Angabe",
			"en": "no estimate",
		},
	}

	return tm

}

func labelsIncreaseDecreaseCurrency() []trl.S {

	tm := []trl.S{
		{
			"de": "aufwerten",
			"en": "appreciate",
		},
		{
			"de": "gleich<br>bleiben",
			"en": "stay constant",
		},
		{
			"de": "abwerten",
			"en": "depreciate",
		},
		{
			"de": "keine<br>Angabe",
			"en": "no estimate",
		},
	}

	return tm

}

func labelsVeryHighVeryLow() []trl.S {

	tm := []trl.S{
		{
			"de": "sehr hoch",
			"en": "very high",
		},
		{
			"de": "hoch",
			"en": "high",
		},
		{
			"de": "normal",
			"en": "normal",
		},
		{
			"de": "niedrig",
			"en": "low",
		},
		{
			"de": "sehr niedrig",
			"en": "very low",
		},
		{
			"de": "keine<br>Angabe",
			"en": "no answer",
		},
	}

	return tm

}

func labelsStrongIncreaseStrongDecrease() []trl.S {

	tm := []trl.S{
		{
			"de": "stark steigen",
			"en": "strongly increase",
		},
		{
			"de": "steigen",
			"en": "increase",
		},
		{
			"de": "gleich bleiben",
			"en": "not change",
		},
		{
			"de": "sinken",
			"en": "decrease",
		},
		{
			"de": "stark sinken",
			"en": "strongly decrease",
		},
		{
			"de": "keine<br>Angabe",
			"en": "no answer",
		},
	}

	return tm

}

func labelsStrongIncreaseStrongDecrease2() []trl.S {

	tm := []trl.S{
		{
			"de": "stark steigen",
			"en": "increase strongly",
		},
		{
			"de": "leicht steigen",
			"en": "increase slightly",
		},
		{
			"de": "unverändert bleiben",
			"en": "stay the same",
		},
		{
			"de": "leicht fallen",
			"en": "decrease slightly",
		},
		{
			"de": "stark fallen",
			"en": "decrease strongly",
		},
		{
			"de": "keine<br>Angabe",
			"en": "don't know",
		},
	}

	return tm

}

func labelsVeryPositiveVeryNegative() []trl.S {

	tm := []trl.S{
		{
			"de": "sehr positiv",
			"en": "very positive",
		},
		{
			"de": "positiv",
			"en": "positive",
		},
		{
			"de": "neutral",
			"en": "neutral",
		},
		{
			"de": "negativ",
			"en": "negative",
		},
		{
			"de": "sehr negativ",
			"en": "very negative",
		},
		{
			"de": "keine<br>Angabe",
			"en": "no answer",
		},
	}

	return tm

}

func labelsStronglyPositiveStronglyNegativeInfluence() []trl.S {

	tm := []trl.S{
		{
			"de": "stark positiv",
			"en": "strongly positive",
		},
		{
			"de": "positiv",
			"en": "positive",
		},
		{
			"de": "kein Einfluss",
			"en": "no influence",
		},
		{
			"de": "negativ",
			"en": "negative",
		},
		{
			"de": "stark negativ",
			"en": "strongly negative",
		},
		{
			"de": "keine<br>Angabe",
			"en": "no answer",
		},
	}

	return tm

}

func labelsPositiveNeutralNegative() []trl.S {

	tm := []trl.S{
		{
			"de": "stark<br>positiv",
			"en": "strongly<br>positive",
		},
		{
			"de": "leicht<br>positiv",
			"en": "slightly<br>positive",
		},
		{
			"de": "neutral",
			"en": "neutral",
		},
		{
			"de": "leicht<br>negativ",
			"en": "slightly<br>negative",
		},
		{
			"de": "stark<br>negativ",
			"en": "strongly<br>negative",
		},
		{
			"de": "keine<br>Angabe",
			"en": "no answer",
		},
	}

	return tm

}

func labelsOvervaluedFairUndervalued() []trl.S {

	tm := []trl.S{
		{
			"de": "überbewertet",
			"en": "over-priced",
		},
		{
			"de": "fair bewertet",
			"en": "fairly priced",
		},
		{
			"de": "unterbewertet",
			"en": "under-priced",
		},
		{
			"de": "keine Angabe",
			"en": "no estimate",
		},
	}

	return tm

}

func labelsConducive1to5() []trl.S {

	tm := []trl.S{
		{
			"de": "sehr geeignet",
			"en": "well suited",
		},
		{
			"de": "geeignet",
			"en": "suited",
		},
		{
			"de": "neutral",
			"en": "neutral",
		},
		{
			"de": "wenig geeignet",
			"en": "not suited",
		},
		{
			"de": "überhaupt nicht geeignet",
			"en": "not at all suited",
		},
		{
			"de": "keine<br>Angabe",
			"en": "no answer",
		},
	}

	return tm

}

func improvedDeterioratedPlusMinus6() []trl.S {
	return labelsPlusPlusMinusMinus()
}

func labelsPlusPlusMinusMinus() []trl.S {

	tm := []trl.S{
		{
			"de": "++",
			"en": "++",
		},
		{
			"de": "+",
			"en": "+",
		},
		{
			"de": "0",
			"en": "0",
		},
		{
			"de": "-",
			"en": "-",
		},
		{
			"de": "--",
			"en": "--",
		},
		{
			"de": "keine<br>Angabe",
			"en": "no answer",
		},
	}

	return tm

}

func labels202508() []trl.S {

	tm := []trl.S{
		{
			"de": "Starker<br>Anstieg ",
			"en": "Strong<br>increase",
		},
		{
			"de": "Leichter<br>Anstieg",
			"en": "Slight<br>increase",
		},
		{
			"de": "Keine<br>Auswirkung",
			"en": "No effect",
		},
		{
			"de": "Leichter<br>Rückgang",
			"en": "Slight<br>decrease",
		},
		{
			"de": "Starker<br>Rückgang",
			"en": "Strong<br>decrease",
		},
		{
			"de": "Keine<br>Antwort",
			"en": "No answer",
		},
	}

	return tm

}

func positiveNegative5() []trl.S {

	tm := []trl.S{
		{
			"de": "stark positiv",
			"en": "strongly positive",
		},
		{
			"de": "leicht positiv",
			"en": "slightly positive",
		},
		{
			"de": "leicht negativ",
			"en": "slightly negative",
		},
		{
			"de": "stark negativ",
			"en": "strongly negative",
		},
		{
			"de": "keine<br>Angabe",
			"en": "no answer",
		},
	}

	return tm

}

func positiveNegative5HardBroken() []trl.S {

	tm := []trl.S{
		{
			"de": "stark <br> positiv",
			"en": "strongly <br> positive",
		},
		{
			"de": "leicht <br> positiv",
			"en": "slightly <br> positive",
		},
		{
			"de": "leicht <br> negativ",
			"en": "slightly <br> negative",
		},
		{
			"de": "stark <br> negativ",
			"en": "strongly <br> negative",
		},
		{
			"de": "keine <br> Angabe",
			"en": "no <br> answer",
		},
	}

	return tm

}

func importanceZeroToFive() []trl.S {

	tm := []trl.S{
		{
			"de": "0<br>(nicht<br>wichtig)",
			"en": "0<br>(not<br>important)",
		},
		{
			"de": "1",
			"en": "1",
		},
		{
			"de": "2",
			"en": "2",
		},
		{
			"de": "3",
			"en": "3",
		},
		{
			"de": "4",
			"en": "4",
		},
		{
			"de": "5<br>(extrem<br>wichtig)",
			"en": "5<br>(extremely<br>important)",
		},
		{
			"de": "keine<br>Angabe",
			"en": "no answer",
		},
	}

	return tm

}

func effectZeroToFive() []trl.S {

	tm := []trl.S{
		{
			"de": "0<br>(kein<br>Effekt)",
			"en": "0<br>(no<br>effect)",
		},
		{
			"de": "1",
			"en": "1",
		},
		{
			"de": "2",
			"en": "2",
		},
		{
			"de": "3",
			"en": "3",
		},
		{
			"de": "4",
			"en": "4",
		},
		{
			"de": "5<br>(extrem großer Effekt)",
			"en": "5<br>(extremely large effect)",
		},
		{
			"de": "keine<br>Angabe",
			"en": "no answer",
		},
	}

	return tm

}

func improvedDeteriorated6() []trl.S {

	tm := []trl.S{
		{
			"de": "stark verbessert",
			"en": "strongly improved",
		},
		{
			"de": "leicht verbessert",
			"en": "slightly improved",
		},
		{
			"de": "nicht verändert",
			"en": "unchanged",
		},
		{
			// extra aggressive hyphenization
			"de": "leicht ver&shy;schl&shy;ech&shy;tert",
			"en": "slightly deteriorated",
		},
		{
			// extra aggressive hyphenization
			"de": "stark ver&shy;schl&shy;ech&shy;tert",
			"en": "strongly deteriorated",
		},

		{
			"de": "keine<br>Angabe",
			"en": "no answer",
		},
	}

	return tm

}

func improvedDeteriorated6b() []trl.S {

	tm := []trl.S{
		{
			"de": "deutlich besser",
			"en": "significantly better",
		},
		{
			"de": "leicht<br>besser",
			"en": "slightly better",
		},
		{

			// "de": "unverändert",
			"de": "un&shy;ver-<br>än&shy;dert",
			"en": "the same",
		},
		{
			"de": "leicht schlechter",
			"en": "slightly worse",
		},
		{
			"de": "deutlich schlechter",
			"en": "significantly worse",
		},

		{
			"de": "keine<br>Angabe",
			"en": "no answer",
		},
	}

	return tm

}

func raiseDecrease6() []trl.S {

	tm := []trl.S{
		{
			"de": "deutlich<br>erhöhen",
			"en": "strong<br>increase",
		},
		{
			"de": "leicht<br>erhöhen",
			"en": "slight<br>increase",
		},
		{
			"de": "nicht<br>verändern",
			"en": "no change",
		},
		{
			"de": "leicht<br>senken",
			"en": "slight<br>decrease",
		},
		{
			"de": "stark<br>senken",
			"en": "strong<br>decrease",
		},

		{
			"de": "keine<br>Angabe",
			"en": "no answer",
		},
	}

	return tm

}

func raiseDecrease6b() []trl.S {

	tm := []trl.S{
		{
			"de": "stark<br>erhöht",
			"en": "strongly<br>increased",
		},
		{
			"de": "erhöht",
			"en": "increased",
		},
		{
			"de": "nicht<br>verändert",
			"en": "not changed",
		},
		{
			"de": "reduziert",
			"en": "decreased",
		},
		{
			"de": "stark<br>reduziert",
			"en": "strongly<br>decreased",
		},

		{
			"de": "keine<br>Angabe",
			"en": "no answer",
		},
	}

	return tm

}

func special202108A() []trl.S {

	tm := []trl.S{
		{
			"de": "erleichtern",
			"en": "easier",
		},
		{
			"de": "unverändert lassen",
			"en": "equal to before July&nbsp;8",
		},
		{
			"de": "erschweren",
			"en": "more complicated",
		},

		{
			"de": "keine Angabe",
			"en": "no answer",
		},
	}

	return tm

}

func special202108B() []trl.S {

	tm := []trl.S{
		{
			"de": "erhöhen ",
			"en": "better",
		},
		{
			"de": "unverändert lassen",
			"en": "equal to before July&nbsp;8",
		},
		{
			"de": "senken",
			"en": "worse",
		},

		{
			"de": "keine Angabe",
			"en": "no answer",
		},
	}

	return tm

}

func special202108C() []trl.S {

	tm := []trl.S{
		{
			"de": "fördern  ",
			"en": "easier",
		},
		{
			"de": "unverändert lassen",
			"en": "equal to before July&nbsp;8",
		},
		{
			"de": "erschweren",
			"en": "more difficult",
		},

		{
			"de": "keine Angabe",
			"en": "no answer",
		},
	}

	return tm

}

func agree6() []trl.S {

	tm := []trl.S{
		{
			"de": "stimme voll zu",
			"en": "strongly agree",
		},
		{
			"de": "stimme zu",
			"en": "agree",
		},
		{
			"de": "unsicher",
			"en": "undecided",
		},
		{
			"de": "stimme nicht zu",
			"en": "disagree",
		},
		{
			"de": "stimme überhaupt nicht zu",
			"en": "strongly disagree",
		},

		{
			"de": "keine<br>Angabe",
			"en": "no answer",
		},
	}

	return tm

}

var columnTemplate6b = []float32{
	6, 1,
	0, 1,
	0, 1,
	0, 1,
	0, 1,
	0.4, 1,
}

func labelsUnlikely() []trl.S {

	tm := []trl.S{
		{
			"de": "Sehr un&shy;wahr&shy;schein&shy;lich",
			"en": "Very unlikely",
		},
		{
			"de": "Un&shy;wahr&shy;schein&shy;lich",
			"en": "Unlikely",
		},
		{
			"de": "Neutral",
			"en": "Neutral",
		},
		{
			"de": "Wahr&shy;schein&shy;lich",
			"en": "Likely",
		},
		{
			"de": "Sehr wahr&shy;schein&shy;lich",
			"en": "Very likely",
		},
		{
			"de": "Keine<br>Angabe",
			"en": "No answer",
		},
	}

	return tm

}

var columnTemplate5a = []float32{
	0, 1,
	0, 1,
	0, 1,
	0, 1,
	0, 1,
}

func labelsCertainty() []trl.S {

	tm := []trl.S{
		{
			"de": "Sehr unsicher",
			"en": "Very uncertain",
		},
		{
			"de": "Unsicher",
			"en": "Uncertain",
		},
		{
			"de": "Neutral",
			"en": "Neutral",
		},
		{
			"de": "Sicher",
			"en": "Certain",
		},
		{
			"de": "Sehr sicher",
			"en": "Very certain",
		},
	}

	return tm

}
