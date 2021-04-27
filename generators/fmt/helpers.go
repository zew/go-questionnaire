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
