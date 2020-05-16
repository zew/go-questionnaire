package lt2020

import (
	"fmt"

	"github.com/zew/go-questionnaire/trl"
)

func labels9(minus4, neutral, plus4 string) []trl.S {

	tm := []trl.S{
		{
			"de": fmt.Sprintf("%v<span class='ordinals'><br>-4</span>", minus4),
		},
		{
			"de": "<span class='ordinals'><br>-3</span>",
		},
		{
			"de": "<span class='ordinals'><br>-2</span>",
		},
		{
			"de": "<span class='ordinals'><br>-1</span>",
		},
		{
			"de": fmt.Sprintf("%v<span class='ordinals'><br>0</span>", neutral),
		},
		{
			"de": "<span class='ordinals'><br>1</span>",
		},
		{
			"de": "<span class='ordinals'><br>2</span>",
		},
		{
			"de": "<span class='ordinals'><br>3</span>",
		},
		{
			"de": fmt.Sprintf("%v<span class='ordinals'><br>4</span>", plus4),
		},
	}

	return tm

}

func labelsFiverPercentages() []trl.S {

	tm := []trl.S{
		{
			"de": "weniger als -15%",
		},
		{
			"de": "-10 bis -15%",
		},
		{
			"de": " -5 bis -10% ",
		},
		{
			"de": "0 bis -5%",
		},
		{
			"de": "mehr als 0%",
		},
	}

	return tm

}

func labelsFiverWichtig() []trl.S {

	tm := []trl.S{
		{
			"de": "sehr wichtig",
		},
		{
			"de": "eher wichtig",
		},
		{
			"de": "eher unwichtig",
		},
		{
			"de": "sehr unwichtig",
		},
		{
			"de": "weder wichtig<br>noch unwichtig",
		},
	}

	return tm

}
func labelsFiverDafuerDagegen() []trl.S {

	tm := []trl.S{
		{
			"de": "sehr dafür",
		},
		{
			"de": "eher dafür",
		},
		{
			"de": "eher dagegen",
		},
		{
			"de": "sehr dagegen",
		},
		{
			"de": "weder dafür noch dagegen",
		},
	}

	return tm

}
