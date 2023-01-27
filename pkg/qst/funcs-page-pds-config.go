package qst

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/css"
	"github.com/zew/go-questionnaire/pkg/trl"
)

const firstColLbl = float32(3)
const outline2Indent = "1.2rem"

var styleHeaderCols1 = css.NewStylesResponsive(nil)
var styleHeaderCols2 = css.NewStylesResponsive(nil)
var styleHeaderCols3 = css.NewStylesResponsive(nil)
var styleHeaderCols4 = css.NewStylesResponsive(nil)

func init() {

	//
	//
	styleHeaderCols1.Desktop.StyleText.FontSize = 85
	styleHeaderCols1.Desktop.StyleGridItem.JustifySelf = "start"
	styleHeaderCols1.Desktop.StyleGridItem.AlignSelf = "end"

	styleHeaderCols1.Desktop.StyleText.AlignHorizontal = "center" // good for Mezzanine..., but bad for Senior
	styleHeaderCols1.Desktop.StyleText.AlignHorizontal = "left"

	// left margin
	styleHeaderCols1.Desktop.StyleBox.Margin = "0 0 0 0.41rem"
	styleHeaderCols1.Mobile.StyleBox.Margin = " 0"
	styleHeaderCols1.Mobile.StyleBox.Padding = "0 0.25rem 0 0.25rem"
	styleHeaderCols1.Desktop.StyleBox.Width = "100%"
	styleHeaderCols1.Desktop.StyleBox.WidthMax = "6rem"

	//
	//
	styleHeaderCols2.Desktop.StyleText.FontSize = 85
	styleHeaderCols2.Desktop.StyleGridItem.JustifySelf = "center"
	styleHeaderCols2.Desktop.StyleGridItem.AlignSelf = "end"
	styleHeaderCols2.Desktop.StyleText.AlignHorizontal = "center"

	// right margin  - exclude range radio
	// bottom margin - dy from range display
	styleHeaderCols2.Desktop.StyleBox.Margin = "0 0.98rem 0.5rem 0"
	// styleHeaderCols2.Desktop.StyleBox.Margin = "0 4.18rem 0.5rem 0"
	// styleHeaderCols2.Desktop.StyleBox.Margin = "0 4.68rem 0.5rem 0"
	styleHeaderCols2.Desktop.StyleBox.Margin = "0 5.18rem 0.5rem 0"
	styleHeaderCols2.Mobile.StyleBox.Margin = "0 1.1rem 0.5rem 0"
	styleHeaderCols2.Desktop.StyleBox.Width = "100%"

	*styleHeaderCols3 = *styleHeaderCols2
	styleHeaderCols3.Desktop.StyleBox.Margin = ""
	styleHeaderCols3.Desktop.StyleBox.Position = "relative"
	styleHeaderCols3.Desktop.StyleBox.Top = "0.6rem"
	// styleHeaderCols3.Desktop.StyleBox.BackgroundColor = "lightgrey"

	*styleHeaderCols4 = *styleHeaderCols2
	styleHeaderCols4.Desktop.StyleGridItem.JustifySelf = "left"
	styleHeaderCols4.Desktop.StyleText.AlignHorizontal = "left"
	styleHeaderCols4.Desktop.StyleBox.Margin = "0 0 0.5rem 3.2rem"

}

// asset classes
type assetClass struct {
	NameUnused   string // unusued
	Prefix       string
	Lbl          trl.S
	LblOld       trl.S
	Short        trl.S // Short label
	TrancheTypes []trancheType
}

// tranche types
// strategies
type trancheType struct {
	NameUnused, Prefix string
	Lbl                trl.S
}

// ultra short abbreviations would be
//
//	CDL / RED / ID
//
// so far unused
var PDSAssetClasses = []assetClass{
	{
		NameUnused: "ac1_corplending",
		Prefix:     "ac1",
		LblOld: trl.S{
			"en": "Corporate direct lending",
			"de": "Corporate direct lending",
		},
		Lbl: trl.S{
			"en": "European Corporate Direct Lending",
			"de": "European Corporate Direct Lending",
		},
		Short: trl.S{
			"en": "Corporate<br>Direct Lending",
			"de": "Corporate<br>Direct Lending",
		},
		TrancheTypes: []trancheType{
			{
				NameUnused: "tt1_senior",
				Prefix:     "tt1",
				Lbl: trl.S{
					"en": "Senior",
					"de": "Senior",
				},
			},
			{
				NameUnused: "tt2_unittranche",
				Prefix:     "tt2",
				Lbl: trl.S{
					"en": "Uni&shy;tranche&nbsp;&nbsp;",
					"de": "Uni&shy;tranche&nbsp;&nbsp;",
				},
			},
			{
				NameUnused: "tt3_subordinated",
				Prefix:     "tt3",
				Lbl: trl.S{
					// "en": "Subordinated&nbsp;/<br>&nbsp;&nbsp;Mezzanine&nbsp;/<br>&nbsp;&nbsp;&nbsp;Other",
					// "en": "<span style='font-size:85%'>Subordinated<br>   &nbsp;&nbsp;Mezzanine<br>  &nbsp;&nbsp;&nbsp;&nbsp;Other<span>",
					"en": "Subordinated",
					"de": "Subordinated",
				},
			},
		},
	},
	{
		NameUnused: "ac2_realestate",
		Prefix:     "ac2",
		Lbl: trl.S{
			"en": "European  Real Estate Debt",
			"de": "European  Real Estate Debt",
		},
		Short: trl.S{
			"en": "Real Estate<br> Debt",
			"de": "Real Estate<br> Debt",
		},
		TrancheTypes: []trancheType{
			{
				NameUnused: "tt1_wholeloan",
				Prefix:     "tt1",
				Lbl: trl.S{
					"en": "Whole loan",
					"de": "Whole loan",
				},
			},
			{
				NameUnused: "tt2_subordinated",
				Prefix:     "tt2",
				Lbl: trl.S{
					"en": "Subordinated",
					"de": "Subordinated",
				},
			},
		},
	},
	{
		NameUnused: "ac3_infrastruct",
		Prefix:     "ac3",
		Lbl: trl.S{
			"en": "European Infrastructure Debt",
			"de": "European Infrastructure Debt",
		},
		Short: trl.S{
			"en": "Infrastructure <br> Debt",
			"de": "Infrastructure <br> Debt",
		},
		TrancheTypes: []trancheType{
			{
				NameUnused: "tt1_senior",
				Prefix:     "tt1",
				Lbl: trl.S{
					"en": "Senior",
					"de": "Senior",
				},
			},
			{
				NameUnused: "tt2_subordinated",
				Prefix:     "tt2",
				Lbl: trl.S{
					"en": "Subordinated",
					"de": "Subordinated",
				},
			},
		},
	},
}

var lblDont = trl.S{
	"de": "Don´t know",
	"en": "Don´t know",
}

var PDSLbls = map[string][]trl.S{
	"teamsize": {
		{
			"en": "<5",
			"de": "<5",
		},
		{
			"en": "5-10",
			"de": "5-10",
		},
		{
			"en": "11-20",
			"de": "11-20",
		},
		{
			"en": ">20",
			"de": ">20",
		},
	},
	"relevance1-5": {
		{
			"en": "not rele&shy;vant<br>(1)",
			"de": "not rele&shy;vant<br>(1)",
		},
		{
			"en": "some&shy;what<br>rele&shy;vant<br>(2)",
			"de": "some&shy;what<br>rele&shy;vant<br>(2)",
		},
		{
			"en": "rele&shy;vant<br>(3)",
			"de": "rele&shy;vant<br>(3)",
		},
		{
			"en": "core prin&shy;ciple<br>(4)",
			"de": "core prin&shy;ciple<br>(4)",
		},
		{
			"en": "potential<br>deal&shy;breaker<br>(5)",
			"de": "potential<br>deal&shy;breaker<br>(5)",
		},
	},
	// "improveDecline1-5": {
	// 	{
	// 		"en": "im&shy;proved",
	// 		"de": "im&shy;proved",
	// 	},
	// 	{
	// 		"en": "&nbsp;",
	// 		"de": "&nbsp;",
	// 	},
	// 	{
	// 		// yes - its terrible
	// 		"en": "sa&shy;me",
	// 		"de": "sa&shy;me",
	// 	},
	// 	{
	// 		"en": "&nbsp;",
	// 		"de": "&nbsp;",
	// 	},
	// 	{
	// 		"en": "de&shy;clined",
	// 		"de": "de&shy;clined",
	// 	},
	// },
	"improveDecline1-5-prev": {
		{
			"en": "bad&nbsp;&nbsp;&nbsp;&nbsp;",
			"de": "bad&nbsp;&nbsp;&nbsp;&nbsp;",
		},
		{
			"en": "&nbsp;",
			"de": "&nbsp;",
		},
		{
			// yes - its terrible
			"en": "satis&shy;factory",
			"de": "satis&shy;factory",
		},
		{
			"en": "&nbsp;",
			"de": "&nbsp;",
		},
		{
			"en": "&nbsp;&nbsp;good",
			"de": "&nbsp;&nbsp;good",
		},
	},
	"improveDecline1-5-prev-spec": {
		// {
		// 	"en": "un&shy;attrac&shy;tive",
		// 	"de": "un&shy;attrac&shy;tive",
		// },
		{
			"en": "bad&nbsp;&nbsp;&nbsp;&nbsp;",
			"de": "bad&nbsp;&nbsp;&nbsp;&nbsp;",
		},
		{
			"en": "&nbsp;",
			"de": "&nbsp;",
		},
		{
			// yes - its terrible
			"en": "satis&shy;factory",
			"de": "satis&shy;factory",
		},
		{
			"en": "&nbsp;",
			"de": "&nbsp;",
		},
		// {
		// 	"en": "attrac&shy;tive",
		// 	"de": "attrac&shy;tive",
		// },
		{
			"en": "&nbsp;&nbsp;good",
			"de": "&nbsp;&nbsp;good",
		},
	},
	"improveDecline1-5-next": {
		{
			"en": "will<br>de&shy;cline",
			"de": "will<br>de&shy;cline",
		},
		{
			"en": "&nbsp;",
			"de": "&nbsp;",
		},
		{
			// yes - its terrible
			"en": "sa&shy;me",
			"de": "sa&shy;me",
		},
		{
			"en": "&nbsp;",
			"de": "&nbsp;",
		},
		{
			"en": "will<br>im&shy;prove",
			"de": "will<br>im&shy;prove",
		},
	},
	"closing-time-weeks": {
		{
			"en": "<<br>4",
			"de": "<<br>4",
		},

		{
			"en": "&nbsp;<br>4&#8209;8",
			"de": "&nbsp;<br>4&#8209;8",
		},
		{
			"en": "weeks<br><br>8&#8209;12",
			"de": "weeks<br><br>8&#8209;12",
		},
		{
			"en": "&nbsp;<br>12&#8209;16",
			"de": "&nbsp;<br>12&#8209;16",
		},
		{
			"en": "><br>16",
			"de": "><br>16",
		},
	},

	"covenants-per-credit": {
		// &#8209; - non breaking dash, non-breaking hyphen
		//   to prevent
		// 			0-
		// 			1
		{
			"en": "0&#8209;1",
			"de": "0&#8209;1",
		},
		{
			"en": "2&#8209;3",
			"de": "2&#8209;3",
		},
		{
			"en": "4&#8209;5",
			"de": "4&#8209;5",
		},
		{
			"en": "&nbsp;>5",
			"de": "&nbsp;>5",
		},
	},
}

var suffixWeeks = trl.S{
	"en": "weeks",
	"de": "Wochen",
}

var suffixYears = trl.S{
	"en": "years",
	"de": "Jahre",
}

var suffixEBITDA = trl.S{
	"en": "x EBITDA",
	"de": "x EBITDA",
}

var suffixInvestedCapital = trl.S{
	"en": "x Invested Cap.",
	"de": "x Invested Cap.",
}
var suffixPercent = trl.S{
	"en": "%",
	"de": "%",
}

var suffixDebtService = trl.S{
	"en": "x Debt Service",
	"de": "x Debt Service",
}

var suffixInterestPayment = trl.S{
	"en": "x Interest Payment",
	"de": "x Interest Payment",
}

var suffixNumDeals = trl.S{
	"en": "deals",
	"de": "Deals",
	// "en": "transactions",
	// "de": "Transaktionen",
}

var suffixMillionEuro = trl.S{
	// capitalizemytitle.com/how-to-abbreviate-million/
	// "en": "million €",
	// "en": "MM €",
	"en": "mn&nbsp;€",
	"de": "Mio&nbsp;€",
}

var placeHolderNum = trl.S{
	"en": "#",
	"de": "#",
}

//	var placeHolderMillion = trl.S{
//		"en": "million Euro",
//		"de": "Millionen Euro",
//	}
var placeHolderMillion = trl.S{
	"en": "mn Euro",
	"de": "Millionen Euro",
}

var placeHolder100percent = trl.S{
	"en": "100",
	"de": "100",
}

func onlySelectedTranchTypes(q *QuestionnaireT, ac assetClass) assetClass {

	ln := len(ac.TrancheTypes)

	// iterate over all
	names := make([]string, 0, ln)
	for i := 0; i < ln; i++ {
		//                               ("ac1_tt1_q031")
		names = append(names, fmt.Sprintf("%v_%v_q031", ac.Prefix, ac.TrancheTypes[i].Prefix))
	}

	newTTs := make([]trancheType, 0, ln)
	for i, name := range names {
		inp := q.ByName(name)
		if inp.Response != "" && inp.Response != "0" {
			newTTs = append(newTTs, ac.TrancheTypes[i])
		}
	}

	acRet := ac
	acRet.TrancheTypes = newTTs

	return acRet
}
