package fmt

import (
	"github.com/zew/go-questionnaire/css"
	"github.com/zew/go-questionnaire/qst"
	"github.com/zew/go-questionnaire/trl"
)

func special202105(q *qst.QuestionnaireT) error {

	if q.Survey.Year != 2021 && q.Survey.Month != 5 {
		return nil
	}

	lblStyleRight := css.NewStylesResponsive(nil)
	lblStyleRight.Desktop.StyleText.AlignHorizontal = "right"
	lblStyleRight.Desktop.StyleBox.Padding = "0 1.0rem 0 0"
	lblStyleRight.Mobile.StyleBox.Padding = " 0 0.5rem 0 0"

	lblStyleLeft := css.NewStylesResponsive(nil)
	lblStyleLeft.Desktop.StyleText.AlignHorizontal = "left"
	lblStyleLeft.Desktop.StyleBox.Padding = "0 0 0 1rem"
	lblStyleLeft.Mobile.StyleBox.Padding = " 0 0 0 2rem"

	page := q.AddPage()
	page.Label = trl.S{
		"de": "Sonderfragen zur Einlagensicherung",
		"en": "Special: Deposit insurance",
	}
	page.Short = trl.S{
		"de": "Sonderfragen:<br>Einlagensicherung",
		"en": "Special:<br>Deposit Insurance",
	}
	page.Style = css.DesktopWidthMaxForPages(page.Style, "46rem")

	// gr1 with intro
	{
		gr := page.AddGroup()
		gr.Cols = 12

		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = 12
			inp.Label = trl.S{
				"de": `
				<p style=''>
					Der Entschädigungsfall der Bremer Greensill Bank, 
					der die gesetzlichen und freiwilligen Sicherungssysteme 
					des privaten Bankenverbands 3,1&nbsp;Milliarden Euro gekostet hat, 
					hat das Thema Einlagensicherung wieder ins Bewusstsein gerufen.
				</p>
				`,
				"en": `
				<p style=''>
					Der Entschädigungsfall der Bremer Greensill Bank, 
					der die gesetzlichen und freiwilligen Sicherungssysteme 
					des privaten Bankenverbands 3,1&nbsp;Milliarden Euro gekostet hat, 
					hat das Thema Einlagensicherung wieder ins Bewusstsein gerufen.
				</p>
				`,
			}
		}

		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = 12
			inp.Label = trl.S{
				"de": `
				<p style=''>
					<b>1.</b> 
					Auch öffentliche Gläubiger der Greensill Bank wurden entschädigt, unter anderem der Kölner Stadtwerke-Konzern sowie die Rundfunkanstalten NDR und Südwestrundfunk. Sollten die deutschen Einlagensicherungsfonds zukünftig nur private Anleger entschädigen dürfen?
				</p>
				`,
				"en": `
				<p style=''>
					<b>1.</b> 
					Auch öffentliche Gläubiger der Greensill Bank wurden entschädigt, unter anderem der Kölner Stadtwerke-Konzern sowie die Rundfunkanstalten NDR und Südwestrundfunk. Sollten die deutschen Einlagensicherungsfonds zukünftig nur private Anleger entschädigen dürfen?
				</p>
				`,
			}
		}
		{
			rad := gr.AddInput()
			rad.Type = "radio"
			rad.Name = "di_private_investors_only"
			rad.ValueRadio = "yes"
			rad.ColSpan = 4
			rad.ColSpanLabel = 4
			rad.ColSpanControl = 1
			rad.Label = trl.S{
				"de": "Ja",
				"en": "Yes",
			}
			rad.StyleLbl = lblStyleRight
		}
		{
			rad := gr.AddInput()
			rad.Type = "radio"
			rad.Name = "di_private_investors_only"
			rad.ValueRadio = "no"
			rad.ColSpan = 4
			rad.ColSpanLabel = 4
			rad.ColSpanControl = 1
			rad.Label = trl.S{
				"de": "Nein",
				"en": "No",
			}
			rad.StyleLbl = lblStyleRight
		}
	}

	//
	//
	//
	//
	// gr2
	{
		var columnTemplateLocal = []float32{
			4, 1,
			0, 1,
			0, 1,
			0, 1,
			0, 1,
			0.4, 1,
		}
		rowLabelsDefaultRisk := []trl.S{
			{
				"de": "Anleger tragen im Entschädigungsfall einen kleinen Selbstanteil (z.B.&nbsp;5-15%)",
				"en": "Anleger tragen im Entschädigungsfall einen kleinen Selbstanteil (z.B.&nbsp;5-15%)",
			},
			{
				"de": "Schärfere Prüfung und Überwachung der Institute durch den Prüfungsverband der deutschen Banken",
				"en": "Schärfere Prüfung und Überwachung der Institute durch den Prüfungsverband der deutschen Banken",
			},
			{
				"de": "Schärfere Prüfung und Überwachung der Institute durch die Bafin",
				"en": "Schärfere Prüfung und Überwachung der Institute durch die Bafin",
			},
			{
				"de": "Ausgeprägtere Risikogewichtung der Beitragsprämien von Banken zum Einlagensicherungsfonds",
				"en": "Ausgeprägtere Risikogewichtung der Beitragsprämien von Banken zum Einlagensicherungsfonds",
			},
		}

		gb := qst.NewGridBuilderRadios(
			columnTemplateLocal,
			labelsConducive1to5(),
			[]string{"di_liability_investor", "di_self_regulation", "di_govt_regulation", "di_collateral"},
			radioVals6,
			rowLabelsDefaultRisk,
		)
		gb.MainLabel = trl.S{
			"de": `<b>2.</b> 
					Welche Reformen sind geeignet, um das Schadensfallrisiko im Falle einer Bankeninsolvenz zu mindern bzw. die Schadensfallsumme für die Einlagensicherung einzugrenzen? 

					`,
			"en": `<b>2.</b> 
					Welche Reformen sind geeignet, um das Schadensfallrisiko im Falle einer Bankeninsolvenz zu mindern bzw. die Schadensfallsumme für die Einlagensicherung einzugrenzen? 

					`,
		}
		gr := page.AddGrid(gb)
		gr.OddRowsColoring = true
	}

	// gr3
	{
		gr := page.AddGroup()
		gr.Cols = 12

		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = 12
			inp.Label = trl.S{
				"de": `
				<p style=''>
					<b>3.</b> 
					Auch einige Kommunen haben Geld bei der Greensill Bank angelegt, um von den höheren Zinsen zu profitieren. Da sie von der Einlagensicherung seit 2017 ausgenommen sind, werden sie keine Entschädigung erhalten. Besteht hier Handlungsbedarf?
				</p>
				`,
				"en": `
				<p style=''>
					<b>3.</b> 
					Auch einige Kommunen haben Geld bei der Greensill Bank angelegt, um von den höheren Zinsen zu profitieren. Da sie von der Einlagensicherung seit 2017 ausgenommen sind, werden sie keine Entschädigung erhalten. Besteht hier Handlungsbedarf?
				</p>
				`,
			}
		}
		{
			rad := gr.AddInput()
			rad.Type = "radio"
			rad.Name = "di_municipalities"
			rad.ValueRadio = "yes_prohibit"
			rad.ColSpan = 12
			rad.ColSpanLabel = 11
			rad.ColSpanControl = 1
			rad.Label = trl.S{
				"de": "Ja, die Anlagemöglichkeiten öffentlicher Haushalte sollten gesetzlich eingeschränkt werden.",
				"en": "Ja, die Anlagemöglichkeiten öffentlicher Haushalte sollten gesetzlich eingeschränkt werden.",
			}
			rad.StyleLbl = lblStyleLeft
		}
		{
			rad := gr.AddInput()
			rad.Type = "radio"
			rad.Name = "di_municipalities"
			rad.ValueRadio = "yes_transparent"
			rad.ColSpan = 12
			rad.ColSpanLabel = 11
			rad.ColSpanControl = 1
			rad.Label = trl.S{
				"de": "Ja, die öffentlichen Haushalte sollten zu mehr Transparenz bei ihren Geldanlagen verpflichtet werden.",
				"en": "Ja, die öffentlichen Haushalte sollten zu mehr Transparenz bei ihren Geldanlagen verpflichtet werden.",
			}
			rad.StyleLbl = lblStyleLeft
		}
		{
			rad := gr.AddInput()
			rad.Type = "radio"
			rad.Name = "di_municipalities"
			rad.ValueRadio = "no"
			rad.ColSpan = 12
			rad.ColSpanLabel = 11
			rad.ColSpanControl = 1
			rad.Label = trl.S{
				"de": "Nein, es ist ausreichend, dass die öffentlichen Haushalte durch die Wähler sanktioniert werden können.",
				"en": "Nein, es ist ausreichend, dass die öffentlichen Haushalte durch die Wähler sanktioniert werden können.",
			}
			rad.StyleLbl = lblStyleLeft
		}
	}

	//
	//
	//
	//
	// gr4
	{

		var columnTemplateLocal = []float32{
			6, 1,
			0, 1,
			0.4, 1,
		}
		headerLabelsLocal := []trl.S{
			{
				"de": "stimme zu",
				"en": "stimme zu",
			},
			{
				"de": "stimme nicht zu",
				"en": "stimme nicht zu",
			},
			{
				"de": "keine<br>Angabe",
				"en": "no answer",
			},
		}

		rowLabelsReforms := []trl.S{
			{
				"de": "… sollten zu mehr Risikoaufklärung gegenüber ihren Kunden verpflichtet werden.",
				"en": "… sollten zu mehr Risikoaufklärung gegenüber ihren Kunden verpflichtet werden.",
			},
			{
				"de": "… sollten dazu verpflichtet werden, das Risiko der Anbieterbanken zu überwachen und ggf.  Banken von der Plattform auszuschließen.",
				"en": "… sollten dazu verpflichtet werden, das Risiko der Anbieterbanken zu überwachen und ggf.  Banken von der Plattform auszuschließen.",
			},
			{
				"de": "… sollten im Entschädigungsfall einen Teil der vermittelten Einlagen ersetzen müssen.",
				"en": "… sollten im Entschädigungsfall einen Teil der vermittelten Einlagen ersetzen müssen.",
			},
			{
				"de": "… sollten nicht weiter reguliert werden, da es die Aufgabe der Privatanleger ist, zu prüfen ob ihre Einlagen bei den Banken sicher sind.",
				"en": "… sollten nicht weiter reguliert werden, da es die Aufgabe der Privatanleger ist, zu prüfen ob ihre Einlagen bei den Banken sicher sind.",
			},
			{
				"de": "… sollten nicht weiter reguliert werden, weil sie aus eigenem Interesse dafür sorgen, dass Banken mit übermäßig riskanten Geschäftsmodellen nicht auf ihrer Plattform Anbieter sind.",
				"en": "… sollten nicht weiter reguliert werden, weil sie aus eigenem Interesse dafür sorgen, dass Banken mit übermäßig riskanten Geschäftsmodellen nicht auf ihrer Plattform Anbieter sind.",
			},
		}

		gb := qst.NewGridBuilderRadios(
			columnTemplateLocal,
			headerLabelsLocal,
			[]string{"di_plattform_education", "di_plattform_surveillance", "di_plattform_liability", "di_plattform_laissez_faire", "di_plattform_incentive"},
			radioVals6,
			rowLabelsReforms,
		)
		gb.MainLabel = trl.S{
			"de": `<b>4.</b> 
				Zinsplattformen, die Einlagen an der Einlagensicherung unterliegende Banken vermitteln, 

					`,
			"en": `<b>4.</b> 
				Zinsplattformen, die Einlagen an der Einlagensicherung unterliegende Banken vermitteln, 

					`,
		}
		gr := page.AddGrid(gb)
		gr.OddRowsColoring = true
	}

	return nil

}
