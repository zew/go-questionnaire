package fmt

import (
	"github.com/zew/go-questionnaire/trl"
)

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
			"en": "no answer",
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
			"de": "nicht verändern",
			"en": "not change",
		},
		{
			"de": "verschlechtern",
			"en": "deteriorate",
		},
		{
			"de": "keine<br>Angabe",
			"en": "no answer",
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
			"de": "gleich bleiben",
			"en": "not change",
		},
		{
			"de": "schlechter",
			"en": "deteriorate",
		},
		{
			"de": "keine<br>Angabe",
			"en": "no answer",
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
			"de": "gleich bleiben",
			"en": "remain unchanged",
		},
		{
			"de": "sinken",
			"en": "decrease",
		},
		{
			"de": "keine<br>Angabe",
			"en": "no answer",
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
			"de": "gleich bleiben",
			"en": "remain unchanged",
		},
		{
			"de": "abwerten",
			"en": "depreciate",
		},
		{
			"de": "keine<br>Angabe",
			"en": "no answer",
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
			"en": "remain unchanged",
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

func labelsOvervaluedFairUndervalued() []trl.S {

	tm := []trl.S{
		{
			"de": "überbewertet",
			"en": "overvalued",
		},
		{
			"de": "fair bewertet",
			"en": "fair valued",
		},
		{
			"de": "unterbewertet",
			"en": "undervalued",
		},
		{
			"de": "keine Angabe",
			"en": "no answer",
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
			"de": "kein Einfluss",
			"en": "no influence",
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
