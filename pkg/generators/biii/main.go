package biii

import (
	"fmt"
	"math"

	"github.com/zew/go-questionnaire/pkg/cfg"
	"github.com/zew/go-questionnaire/pkg/css"
	"github.com/zew/go-questionnaire/pkg/ctr"
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

// Create creates a JSON file for a financial markets survey
func Create(s qst.SurveyT) (*qst.QuestionnaireT, error) {

	ctr.Reset()

	q := qst.QuestionnaireT{}
	q.Survey = s
	q.LangCodes = []string{"de"} // governs default language code

	q.Survey.Org = trl.S{"de": "BIII", "en": "BIII"}

	q.Survey.Name = trl.S{
		"de": "Erhebung 2022 - Marktstudie Impact Investment",
	}

	// page 0
	{
		page := q.AddPage()
		page.Short = trl.S{"de": "Start"}
		page.Label = trl.S{"de": ""}
		// page.NoNavigation = true
		page.WidthMax("42rem")

		// gr0
		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 1
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 1
				inp.ColSpanLabel = 1
				inp.Label = trl.S{
					"de": `

					<p style='text-align: justify; font-size: 130%'>
						Willkommen zur Marktstudie der <a href='https://bundesinitiative-impact-investing.de/'>Bundesinitiative Impact Investing (BIII)</a>
						 im Auftrag der AIR GmbH 
						  –  ausgeführt durch das ZEW Mannheim					
					</p>

					<p style='text-align: justify;'>
					Im Rahmen dieser Erhebung nutzen wir bewusst eine breite Definition
					 von Impact Investments. 
					 
					 Diese spiegelt das Verständnis des Global Impact Investing Networks (GIIN) wider. 
					 
					 Demnach sind Impact Investments "Investitionen, die mit der Absicht getätigt werden, 
					 neben einer finanziellen Rendite auch eine positive, 
					 messbare soziale und ökologische Wirkung zu erzielen" (GIIN, 2017). 
					 
					 Auch in akademischen Studien wird ähnlich argumentiert. 
					 
					 Busch et al. (2021) bezeichnen Impact bezogene Investments 
					 als "Investitionen, die sich auf Veränderungen in der realen Welt 
					 im Hinblick auf die Lösung sozialer Probleme und/oder die Milderung 
					 ökologischer Schäden konzentrieren". 
					 
					 Ob entsprechende Investments aktiv und zusätzlich zu sozialen
					  und ökologischen Lösungen und Veränderungen beitragen, 
					  wird im Rahmen der Erhebung separat ermittelt. 					
					</p>
					`,
				}
			}
			{
				inp := gr.AddInput()
				inp.Type = "dyn-textblock"
				inp.DynamicFunc = "RenderStaticContent"
				inp.Param = "page-0-data-protection.md"
				inp.ColSpan = 1
				inp.ColSpanLabel = 1
			}
		}

		// gr1
		{
			gb := qst.NewGridBuilderRadios(
				columnTemplate3,
				roleOrFunctionQ1,
				[]string{"q1_role"},
				radioValsQ1,
				[]trl.S{{"de": "<b>1.</b> &nbsp;	Sind Sie…?"}},
			)
			// gb.MainLabel = ...
			gr := page.AddGrid(gb)
			_ = gr
		}

		// gr2
		{
			labels := []trl.S{
				{"de": "Privat Investor"},
				{"de": "Business Angel"},
				{"de": "VC / PE Impact Fondsmanager"},
				{"de": "Privates Finanzinstitut (einschließlich traditioneller Banken und ethischer Banken)"},
				{"de": "Versicherungsgesellschaft oder Pensionsfonds"},
				{"de": "Mikrofinanzinstitution"},
				{"de": "Crowdfunding-Plattform"},
				{"de": "Stiftung"},
				{"de": "Family Office"},
				{"de": "Investmentfondsmanager eines börsennotierten Unternehmens"},
				{"de": "Entwicklungsfinanzierungsagentur oder -einrichtung"},
				{"de": "Öffentlicher Finanzierungsfonds oder -einrichtung"},
				{"de": "Inkubator und Beschleuniger"},
			}
			radioValues := []string{
				"private_investor",
				"business_angel",
				"impact_fund_mgr",
				"private_bank",
				"insurance_fund",
				"micro_institution",
				"crowdfunding_plattform",
				"foundation",
				"family_office",
				"mgr_listed_comp",
				"development_agency",
				"public_fund",
				"incubator",
			}
			gr := page.AddGroup()
			gr.Cols = 2
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{"de": "<b>2.</b> &nbsp;	Als welche Art von Organisation ordnen Sie sich ein?"}
				inp.ColSpan = gr.Cols
			}
			for idx, labl := range labels {
				rad := gr.AddInput()
				rad.Type = "radio"
				rad.Name = "q2"
				rad.ValueRadio = radioValues[idx]

				rad.ColSpan = 1
				rad.ColSpanLabel = 1
				rad.ColSpanControl = 6

				rad.Label = labl

				rad.ControlFirst()
			}
			{
				inp := gr.AddInput()
				inp.Type = "text"
				inp.Name = "q2other"
				inp.MaxChars = 20
				inp.Label = trl.S{"de": "Andere, bitte nennen"}
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 2
				inp.ColSpanControl = 3
				inp.StyleLbl = css.NewStylesResponsive(inp.StyleLbl)
				inp.StyleLbl.Desktop.StyleBox.Padding = "0 0 0 3.4rem"

				inp.Style = css.NewStylesResponsive(inp.Style)
				inp.Style.Desktop.StyleBox.Margin = "1.2rem 0 0 0"
			}
		}

		// gr3
		{
			labels := []trl.S{
				{"de": "Ausschlüsse oder Negative Screening"},
				{"de": "Best-in-Class"},
				{"de": "ESG (ökologische, soziale oder Governance) Integration"},
				{"de": "Norms-based Screening (z.B. UN's Global Compact, OECD Guidelines for Multinational Enterprises, …)"},
				{"de": "Thematische Funds oder Themenbezogene Produkte (z.B. Klima, Menschenrechte, Gesundheit, ...)"},
				{"de": "Impact Investments"},
				{"de": "Wir tätigen keine Investments"},
			}
			radioValues := []string{
				"exclusions",
				"best_in_class",
				"esg",
				"norms_based",
				"theme_funds",
				"impact_investments",
				"no_investing",
			}
			gr := page.AddGroup()
			gr.Cols = 2
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{"de": "<b>3.</b> &nbsp;	Welchen Fokus haben Sie bei Ihrer Investment-Strategie bzw. welche Produktgestaltung nutzen Sie? "}
				inp.ColSpan = gr.Cols
			}
			for idx, label := range labels {
				rad := gr.AddInput()
				rad.Type = "radio"
				rad.Name = "q3"
				rad.ValueRadio = radioValues[idx]

				rad.ColSpan = 1
				rad.ColSpanLabel = 1
				rad.ColSpanControl = 6

				rad.Label = label

				rad.ControlFirst()
			}
			{
				inp := gr.AddInput()
				inp.Type = "text"
				inp.Name = "q3other"
				inp.MaxChars = 20
				inp.Label = trl.S{"de": "Andere, bitte nennen"}
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 2
				inp.ColSpanControl = 3
				inp.StyleLbl = css.NewStylesResponsive(inp.StyleLbl)
				inp.StyleLbl.Desktop.StyleBox.Padding = "0 0 0 3.4rem"

				inp.Style = css.NewStylesResponsive(inp.Style)
				inp.Style.Desktop.StyleBox.Margin = "1.2rem 0 0 0"
			}
		}

	}

	// page 1
	{
		page := q.AddPage()
		page.Label = trl.S{"de": "Ihre Entscheidung zu Impact Investing"}
		page.Short = trl.S{"de": "Entscheidung"}
		page.WidthMax("38rem")

		page.ValidationFuncName = "biiiPage1"
		page.ValidationFuncMsg = trl.S{
			"de": "no javascript dialog message needed",
		}

		// gr0
		{
			labels := []trl.S{
				{"de": "Gegenwärtig"},
				{"de": "Möglicherweise in der Zukunft"},
				{"de": "In Planung"},
				{"de": "Nein"},
			}
			radioValues := []string{
				"now",
				"in_future",
				"in_planning",
				"no",
			}
			gr := page.AddGroup()
			gr.Cols = 5
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{"de": "<b>4.</b> &nbsp;	Arbeitet Ihre Organisation mit Impact Investments?"}
				inp.ColSpan = gr.Cols
				inp.ColSpan = 4

			}

			// composit validation
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = gr.Cols
				inp.ColSpan = 1
				inp.ColSpanLabel = 1
				inp.Validator = "biii_branch1"

				inp.Style = css.NewStylesResponsive(inp.Style)
				inp.Style.Desktop.Position = "relative"
				inp.Style.Desktop.Top = "7rem"
				inp.Style.Desktop.Left = "-6rem"

			}

			for idx, label := range labels {
				rad := gr.AddInput()
				rad.Type = "radio"
				rad.Name = "q4"
				rad.ValueRadio = radioValues[idx]

				rad.ColSpan = gr.Cols
				rad.ColSpanLabel = 1
				rad.ColSpanControl = 6

				rad.Label = label

				rad.ControlFirst()
				// rad.Validator = "must"
				// rad.ErrMsg = "must_central_biii"

				//
				if idx == 0 {
					labels := []trl.S{
						{"de": "seit >10 Jahren "},
						{"de": "seit >5 Jahren"},
						{"de": "seit >3 Jahren"},
						{"de": "seit <=3 Jahren"},
						{"de": "erst seit kurzem"},
					}
					radioValues := []string{
						"10yrs",
						"5yrs",
						"3yrs",
						"lessthan3",
						"recently",
					}
					// gr := page.AddGroup()
					// gr.Cols = 1
					for idx, label := range labels {
						rad := gr.AddInput()
						rad.Type = "radio"
						rad.Name = "q4a"
						rad.ValueRadio = radioValues[idx]

						rad.ColSpan = gr.Cols
						rad.ColSpanLabel = 1
						rad.ColSpanControl = 6

						rad.Label = label

						rad.ControlFirst()

						rad.Style = css.NewStylesResponsive(rad.Style)
						rad.Style.Desktop.StyleBox.Margin = "0 0 0 3.2rem"

					}
				} // idx==0

			} // range labels

		}
	}

	// page 2a-01 II now yes
	{
		page := q.AddPage()
		page.Short = trl.S{"de": "II Now"}
		page.Label = trl.S{"de": ""}
		page.NavigationCondition = "BIIINow"
		// page.NoNavigation = true
		page.WidthMax("42rem")

		// gr0
		{
			labels := []trl.S{
				{"de": "Wir arbeiten nur mit Impact Investments"},
				{"de": "Impact Investing ist Teil unserer SRI / ESG Aktivitäten"},
				{"de": "Impact Investing ist ein unabhängiger Bereich neben unseren konventionellen und SRI / ESG Aktivitäten"},
			}
			radioValues := []string{
				"all",
				"partly",
				"separate_dept",
			}
			gr := page.AddGroup()
			gr.Cols = 1
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{"de": "<b>5.</b> &nbsp;	Wenn Sie gegenwärtig mit Impact Investments arbeiten, welchen Platz haben Impact Investments in Ihrer Organisation?"}
				inp.ColSpan = gr.Cols
			}
			for idx, label := range labels {
				rad := gr.AddInput()
				rad.Type = "radio"
				rad.Name = "q5"
				rad.ValueRadio = radioValues[idx]

				rad.ColSpan = 1
				rad.ColSpanLabel = 1
				rad.ColSpanControl = 6

				rad.Label = label

				rad.ControlFirst()
			}
		}

		//
		//
		// gr1
		{
			labels := []trl.S{
				{"de": "Impact Investments am 31/12/2021"},
				{"de": "Andere Investments mit ESG Bezug am 31/12/2021"},
				{"de": "Konventionelle Investments am 31/12/2021"},
			}
			radioValues := []string{
				"impact",
				"other",
				"conventional",
			}
			gr := page.AddGroup()
			gr.Cols = 1
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{"de": "<b>6.</b> &nbsp;	Welches Investitionsvolumen (Assets under Management, Kreditsumme, investiertes Kapital) verzeichnet Ihre Organisation <u>insgesamt</u>?"}
				inp.ColSpan = gr.Cols
			}
			for idx, label := range labels {
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = fmt.Sprintf("q6_%v", radioValues[idx])
				inp.Label = label

				inp.ColSpan = 1
				inp.ColSpanLabel = 6
				inp.ColSpanControl = 2
				inp.Min = 0
				inp.Max = math.MaxFloat64
				inp.Step = 1000
				inp.Step = 1
				inp.MaxChars = 18
				inp.Suffix = trl.S{"de": "€"}
				inp.Placeholder = trl.S{"de": "0.000.000"}
			}
		}

	}

	// page 2b-01 IE no or later
	{
		page := q.AddPage()
		page.Short = trl.S{"de": "II not or later"}
		page.Label = trl.S{"de": ""}
		page.NavigationCondition = "BIIILater"
		// page.NoNavigation = true
		page.WidthMax("42rem")

		// gr0
		{
			labels := []trl.S{
				{"de": "Wir bereiten einen Markteintritt strategisch vor"},
				{"de": "Wir prüfen künftige Marktchancen sorgfältig "},
				{"de": "Wir beginnen, Informationen zu diesem Markt zu sammeln"},
				{"de": "Wir sehen keine Notwendigkeit, uns damit zu befassen "},
			}
			radioValues := []string{
				"strategic_prep",
				"due_dilligence",
				"collect_info",
				"no_necessity",
			}
			gr := page.AddGroup()
			gr.Cols = 1
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{"de": "<b>42.</b> &nbsp;	Falls Ihre Organisation bisher noch nicht im Impact Investing-Markt tätig ist, wie bewerten Sie dieses Feld für die Zukunft: "}
				// (bitte nur eine Auswahl)
				inp.ColSpan = gr.Cols
			}
			for idx, label := range labels {
				rad := gr.AddInput()
				rad.Type = "radio"
				rad.Name = "q42"
				rad.ValueRadio = radioValues[idx]

				rad.ColSpan = 1
				rad.ColSpanLabel = 1
				rad.ColSpanControl = 6

				rad.Label = label

				rad.ControlFirst()
			}
			{
				inp := gr.AddInput()
				inp.Type = "textarea"
				inp.Name = "q42other"
				inp.MaxChars = 150
				inp.Label = trl.S{"de": "Wir sehen diesen Markt eher skeptisch aus den folgenden Gründen:"}
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 2
				inp.ColSpanControl = 3
				// inp.StyleLbl = css.NewStylesResponsive(inp.StyleLbl)
				// inp.StyleLbl.Desktop.StyleBox.Padding = "0 0 0 3.4rem"
				// inp.Style = css.NewStylesResponsive(inp.Style)
				// inp.Style.Desktop.StyleBox.Margin = "1.2rem 0 0 0"
			}
		}

	}

	//
	// page 7 - after seasonal
	// Finish questionnaire?  - one before last page
	{
		page := q.AddPage()
		page.Label = trl.S{"de": "Abschluss", "en": "Finish"}
		page.Short = trl.S{"de": "Abschluss", "en": "Finish"}
		page.WidthMax("36rem")

		// gr1
		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 2
			gr.Style = css.NewStylesResponsive(gr.Style)
			gr.Style.Desktop.StyleGridContainer.GapRow = "0.2rem"
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{"de": "Kommentar zur Umfrage: ", "en": "Comment on the survey: "}
				inp.Label = trl.S{
					"de": "Wollen Sie uns noch etwas mitteilen?",
					"en": "Any remarks or advice for us?",
				}
				inp.ColSpanLabel = 1
			}
			{
				inp := gr.AddInput()
				inp.Type = "textarea"
				inp.Name = "comment"
				inp.MaxChars = 300
				inp.ColSpanLabel = 0
				inp.ColSpanControl = 1
			}
		}

		// gr2
		{
			gr := page.AddGroup()
			gr.Cols = 1
			{
				rad := gr.AddInput()
				rad.Type = "radio"
				rad.Name = "finished"
				rad.ValueRadio = qst.RemainOpen
				rad.ColSpan = 1
				rad.ColSpanLabel = 6
				rad.ColSpanControl = 1
				rad.Label = trl.S{
					"de": "Zugang bleibt bestehen.  Daten können in weiteren Sitzungen geändert/ergänzt werden. ",
					"en": "Leave the questionnaire open. Data  can be changed/completed&nbsp;in later sessions.     ",
				}
			}
			{
				rad := gr.AddInput()
				rad.Type = "radio"
				rad.Name = "finished"
				rad.ValueRadio = qst.Finished
				rad.ColSpan = 1
				rad.ColSpanLabel = 6
				rad.ColSpanControl = 1
				rad.Label = trl.S{
					"de": "Fragebogen ist abgeschlossen und kann nicht mehr geöffnet werden. ",
					"en": "The questionnaire is finished. No more edits.                         ",
				}
			}
		}

		// gr3
		{
			gr := page.AddGroup()
			gr.Style = css.NewStylesResponsive(gr.Style)
			gr.Cols = 2
			gr.Style.Desktop.StyleGridContainer.TemplateColumns = "3fr 1fr"
			// gr.Width = 80

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{"de": "", "en": ""}
				inp.Label = trl.S{
					"de": "Durch Klicken erhalten Sie eine Zusammenfassung Ihrer Antworten",
					"en": "By clicking, you will receive a summary of your answers.",
				}
				inp.ColSpan = 1
				inp.ColSpanLabel = 1
			}
			{
				inp := gr.AddInput()
				inp.Type = "button"
				inp.Name = "submitBtn"
				inp.Response = fmt.Sprintf("%v", len(q.Pages)-1+1) // +1 since one page is appended below
				inp.Label = cfg.Get().Mp["end"]
				inp.Label = cfg.Get().Mp["finish_questionnaire"]
				inp.ColSpan = 1
				inp.ColSpanControl = 1
				inp.AccessKey = "n"
				inp.StyleCtl = css.NewStylesResponsive(inp.StyleCtl)
				inp.StyleCtl.Desktop.StyleGridItem.JustifySelf = "end"
				// inp.StyleCtl.Desktop.StyleBox.WidthMin = "8rem" // does not help with button
			}
		}

		// pge.ExampleSixColumnsLabelRight()

	}

	// page 8 - after seasonal
	// Report of results
	{
		page := q.AddPage()
		page.Label = trl.S{"de": "Ihre Eingaben", "en": "Summary of results"}
		page.NoNavigation = true
		page.WidthMax("calc(100% - 1.2rem)")
		page.WidthMax("40rem")
		{
			gr := page.AddGroup()
			gr.Cols = 1
			{
				inp := gr.AddInput()
				inp.Type = "dyn-textblock"
				inp.ColSpanControl = 1
				inp.DynamicFunc = "ResponseStatistics"
			}
			{
				inp := gr.AddInput()
				inp.Type = "dyn-textblock"
				inp.ColSpanControl = 1
				inp.DynamicFunc = "PersonalLink"
			}
			{
				inp := gr.AddInput()
				inp.Type = "dyn-textblock"
				inp.DynamicFunc = "RenderStaticContent"
				inp.Param = "site-imprint.md"
				inp.ColSpan = 1
				inp.ColSpanLabel = 1
			}
		}
	}

	q.Hyphenize()
	q.ComputeMaxGroups()
	if err := q.TranslationCompleteness(); err != nil {
		return &q, err
	}
	if err := q.Validate(); err != nil {
		return &q, err
	}
	return &q, nil

}
