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
		"de": "Sonderfrage: Einlagensicherung",
		"en": "Special: Deposit insurance",
	}
	page.Short = trl.S{
		"de": "Sonderfrage:<br>Einlagensicherung",
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

					The loss at Bremen Greensill Bank, 
					cost public and private insurers 3,1&nbsp;Billion Euro.  
					This has put the issue of deposit insurance back into the spotlight.
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
					Public creditors of the Greensill back were compensated.
					Among others the utilities of the city of Cologne as well as North German Broadcasting Corporation. 
					
					Should German deposit insurance be allowed to compensate only private investors in future?
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
				"en": "Investors pay a small co-payment in the event of a loss (i.e.&nbsp;5-15%)",
			},
			{
				"de": "Schärfere Prüfung und Überwachung der Institute durch den Prüfungsverband der deutschen Banken",
				"en": "Stricter oversight of the banks by the oversight board of the German banking industry",
			},
			{
				"de": "Schärfere Prüfung und Überwachung der Institute durch die Bafin",
				"en": "Stricter oversight of the banks by the state regulator Bafin",
			},
			{
				"de": "Ausgeprägtere Risikogewichtung der Beitragsprämien von Banken zum Einlagensicherungsfonds",
				"en": "More risk-weighted premiums of banks for the deposit insurance",
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
					Which reforms are suited, to reduce the risk of bankruptcies? 
					Which reforms are suited to limit the exposure of the deposit insurance? 

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
					Wählen Sie bitte die Antwortmöglichkeit aus, der Sie am meisten zustimmen.
				</p>
				`,
				"en": `
				<p style=''>
					<b>3.</b> 
					Some municipal governments deposited money with the Greensill Bank, to benefit from
					higher interest rates. 
					They will not be compensated, since they are excluded from deposit insurance.
					Please choose the answer most closely to your position.
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
				"en": "Yes, investment choices of public households should by restricted by law.",
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
				"en": "Yes, investment choices of public households should be made more transparent by regulation.",
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
				"en": "No, voter punishments of local governments are sufficient.",
			}
			rad.StyleLbl = lblStyleLeft
		}
		{
			rad := gr.AddInput()
			rad.Type = "radio"
			rad.Name = "di_no_answer"
			rad.ValueRadio = "no_answer"
			rad.ColSpan = 12
			rad.ColSpanLabel = 11
			rad.ColSpanControl = 1
			rad.Label = trl.S{
				"de": "Keine Antwort",
				"en": "No Answer",
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
				"en": "agree",
			},
			{
				"de": "stimme nicht zu",
				"en": "dont agree",
			},
			{
				"de": "keine<br>Angabe",
				"en": "no answer",
			},
		}

		rowLabelsReforms := []trl.S{
			{
				"de": "… sollten zu mehr Risikoaufklärung gegenüber ihren Kunden verpflichtet werden.",
				"en": "… should be obliged to educate their customers more about risks.",
			},
			{
				"de": "… sollten dazu verpflichtet werden, das Risiko der Anbieterbanken zu überwachen und ggf.  Banken von der Plattform auszuschließen.",
				"en": "… should be obliged to supervise the risk of the banks and if necessary exclude banks from their platform",
			},
			{
				"de": "… sollten im Entschädigungsfall einen Teil der vermittelten Einlagen ersetzen müssen.",
				"en": "… should have to refund the losses partially, in the event of a loss.",
			},
			{
				"de": "… sollten nicht weiter reguliert werden, da es die Aufgabe der Privatanleger ist, zu prüfen ob ihre Einlagen bei den Banken sicher sind.",
				"en": "… should not be regulated any further, since it is the responsibility of the private investor to do due diligence.",
			},
			{
				"de": "… sollten nicht weiter reguliert werden, weil sie aus eigenem Interesse dafür sorgen, dass Banken mit übermäßig riskanten Geschäftsmodellen nicht auf ihrer Plattform Anbieter sind.",
				"en": "… should not be regulated any further, since they have an incentive to exclude banks with outsized risky business models from their platforms.",
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
				Loan platforms, intermediating deposits to banks, subject to deposit insurance. 

					`,
		}
		gr := page.AddGrid(gb)
		gr.OddRowsColoring = true
	}

	return nil

}
