package fmt

import (
	"fmt"
	"log"
	"strings"

	"github.com/zew/go-questionnaire/ctr"
	"github.com/zew/go-questionnaire/qst"
	"github.com/zew/go-questionnaire/tpl"
	"github.com/zew/go-questionnaire/trl"
)

// Create creates a JSON file for a financial markets survey
func Create(params []qst.ParamT) (*qst.QuestionnaireT, error) {

	ctr.Reset()

	// qst.RadioVali = "mustRadioGroup"
	qst.RadioVali = ""
	qst.CSSLabelHeader = ""
	qst.CSSLabelRow = ""

	q := qst.QuestionnaireT{}
	q.Survey = qst.NewSurvey("fmt")
	q.Survey.Params = params
	q.LangCodes = []string{"de", "en"} // governs default language code

	q.Survey.Org = trl.S{"de": "ZEW", "en": "ZEW"}
	q.Survey.Name = trl.S{"de": "Finanzmarkttest", "en": "Financial Markets Survey"}

	// Page 0

	{
		p := q.AddPage()
		p.Label = trl.S{"de": "Begrüßung", "en": "Greeting"}
		p.NoNavigation = true

		{
			gr := p.AddGroup()
			gr.Cols = 3

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.CSSLabel = "special-line-height-higher"
				inp.ColSpanLabel = 3
				impr := trl.S{}
				for _, lc := range q.LangCodes {
					w1 := &strings.Builder{}
					err := tpl.RenderStaticContent(
						w1, "data-protection.md", q.Survey.Type, lc,
					)
					if err != nil {
						log.Print(err)
					}
					impr[lc] = w1.String()

				}
				inp.Desc = impr
			}

			gr = p.AddGroup()
			gr.Cols = 3
			gr.Width = 75
			{
				inp := gr.AddInput()
				inp.Type = "radiogroup"
				inp.Name = "proxy"
				inp.CSSLabel = "special-line-height-higher"

				inp.Label = trl.S{"de": " ", "en": " "}
				inp.Desc = trl.S{"de": "Sind Sie die angeschriebene Person?", "en": "Are you the addressee?"}

				inp.ColSpanLabel = 1
				inp.ColSpanControl = 2
				{
					rad := inp.AddRadio()
					rad.HAlign = qst.HLeft
					rad.HAlign = qst.HCenter
					rad.Label = trl.S{
						"de": " &nbsp; <br>Ja, ich bin die angeschriebene Person.",
						"en": "Yes, I am the addressee.",
					}
				}
				{
					rad := inp.AddRadio()
					rad.HAlign = qst.HLeft
					rad.HAlign = qst.HCenter
					rad.Label = trl.S{
						"de": "Nein, ich fülle den Fragebogen in Vertretung der angeschriebenen Person aus.",
						"en": "No. I am filling in for the addressee.",
					}
				}
			}

			gr = p.AddGroup()
			gr.Cols = 3
			gr.Width = 75
			{
				inp := gr.AddInput()
				inp.Type = "textarea"
				inp.Name = "address_change"
				inp.Desc = trl.S{
					"de": "Meine Adresse hat sich geändert",
					"en": "My address has changed",
				}
				inp.MaxChars = 150
				inp.ColSpanLabel = 1
				inp.ColSpanControl = 2
			}

			gr = p.AddGroup()
			gr.Cols = 1
			gr.Width = 75
			{
				inp := gr.AddInput()
				inp.Type = "button"
				inp.Name = "submitBtn"
				inp.Response = "1"
				inp.Label = trl.S{
					"de": "Weiter",
					"en": "Next",
				}
				inp.AccessKey = "n"
				inp.ColSpanControl = 1
				inp.HAlignControl = qst.HRight
			}

		}
	}

	// Page 1
	{
		p := q.AddPage()
		p.Section = trl.S{"de": "Konjunktur", "en": "Business cycle"}
		p.Label = trl.S{"de": "Status und Ausblick", "en": "Status and outlook"}
		p.Short = trl.S{"de": "Konjunktur:<br>Status,<br>Ausblick", "en": "Business cycle:<br>Status,<br>Outlook"}
		p.Width = 70

		//
		//
		labels123Matrix := []trl.S{
			{
				"de": "Euroraum",
				"en": "Euro area",
			},
			{
				"de": "Deutschland",
				"en": "Germany",
			},
			{
				"de": "USA",
				"en": "US",
			},
			{
				"de": "Weltwirtschaft",
				"en": "Global economy",
			},
		}

		// gr1
		{
			names1stMatrix := []string{
				"y0_ez",
				"y0_deu",
				"y0_usa",
				"y0_glob",
			}
			gr := p.AddRadioMatrixGroup(labelsGoodBad(), names1stMatrix, labels123Matrix)
			gr.Cols = 5 // necessary, otherwise no vspacers
			gr.OddRowsColoring = true
			gr.Label = trl.S{
				"de": "1.",
				"en": "1.",
			}
			gr.Desc = trl.S{
				"de": "Die gesamtwirtschaftliche Situation beurteilen wir als",
				"en": "We assess the overall economic situation as",
			}

		}

		//
		// gr2
		{
			names2stMatrix := []string{
				"y_ez",
				"y_deu",
				"y_usa",
				"y_glob",
			}
			gr := p.AddRadioMatrixGroup(labelsImproveDeteriorate(), names2stMatrix, labels123Matrix)
			gr.Cols = 5 // necessary, otherwise no vspacers
			gr.OddRowsColoring = true
			gr.Label = trl.S{
				"de": "2a.",
				"en": "2a.",
			}
			gr.Desc = trl.S{
				"de": "Die gesamtwirtschaftliche Situation wird sich mittelfristig (<b>6</b>&nbsp;Mo.)",
				"en": "The overall economic situation medium term (<b>6</b>&nbsp;months) will",
			}

		}

		//
		// gr3
		{
			names3rdMatrix := []string{
				"y24_ez",
				"y24_deu",
				"y24_usa",
				"y24_glob",
			}

			gr := p.AddRadioMatrixGroup(labelsImproveDeteriorate(), names3rdMatrix, labels123Matrix)
			gr.Cols = 5 // necessary, otherwise no vspacers
			gr.OddRowsColoring = true
			gr.Label = trl.S{
				"de": "2b.",
				"en": "2b.",
			}
			gr.Desc = trl.S{
				"de": "Die gesamtwirtschaftliche Situation wird sich langfristig (<b>24</b>&nbsp;Mo.)",
				"en": "The overall economic situation long term (<b>24</b>&nbsp;months) will",
			}

		}

	}

	// page 2
	{
		p := q.AddPage()
		p.Label = trl.S{"de": "Wachstum", "en": "Growth"}
		p.AestheticCompensation = 8
		p.Width = 90

		{
			gr := p.AddGroup()
			gr.Cols = 5 // necessary, otherwise no vspacers
			gr.Label = trl.S{"de": "3a.", "en": "3a."}
			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = "y_q_deu"
				inp.MaxChars = 4
				inp.Validator = "inRange20"

				inp.ColSpanLabel = 4
				inp.Desc = trl.S{
					"de": fmt.Sprintf("Unsere Prognose für das <b>deutsche</b> BIP Wachstum in %v (real, saisonbereinigt, nicht annualisiert):", nextQ()),
					"en": fmt.Sprintf("Our estimate for the <b>German</b> GDP growth in %v (real, seasonally adjusted, non annualized):", nextQ()),
				}
				inp.Suffix = trl.S{"de": "%", "en": "pct"}
				inp.HAlignLabel = qst.HLeft
				inp.HAlignControl = qst.HLeft
			}

			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = "y_y_deu"
				inp.MaxChars = 4
				inp.Validator = "inRange20"

				inp.ColSpanLabel = 4
				inp.Desc = trl.S{
					"de": fmt.Sprintf("Unsere Prognose für das BIP Wachstum für das Jahr %v (real, saisonbereinigt):", nextY()),
					"en": fmt.Sprintf("Our estimate for the GDP growth in %v (real, seasonally adjusted):", nextY()),
				}
				inp.Suffix = trl.S{"de": "%", "en": "pct"}
				inp.HAlignLabel = qst.HLeft
				inp.HAlignControl = qst.HLeft
			}
		}

		{
			gr := p.AddGroup()
			gr.Cols = 5 // necessary, otherwise no vspacers
			gr.Label = trl.S{"de": "3b.", "en": "3b."}

			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = "yshr_q_deu"
				inp.MaxChars = 4
				inp.Validator = "inRange100"

				inp.ColSpanLabel = 4
				inp.Desc = trl.S{
					"de": fmt.Sprintf("Die Wahrscheinlichkeit eines negativen Wachstums des <b>deutschen</b> BIP in %v liegt bei:", nextQ()),
					"en": fmt.Sprintf("The probability of negative growth for the <b>German</b> GDP in %v is:", nextQ()),
				}
				inp.Suffix = trl.S{"de": "%", "en": "pct"}
				inp.HAlignLabel = qst.HLeft
				inp.HAlignControl = qst.HLeft
			}

			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = "yshr_y_deu"
				inp.MaxChars = 4
				inp.Validator = "inRange100"

				inp.ColSpanLabel = 4
				inp.Desc = trl.S{
					"de": fmt.Sprintf("Die Wahrscheinlichkeit einer Rezession in Deutschland (mind. 2&nbsp;Quartale neg. Wachstum) bis Q4 %v liegt bei:", nextY()),
					"en": fmt.Sprintf("The probability of a recession in Germany (at least 2&nbsp;quarters neg. growth) until Q4 %v is:", nextY()),
				}
				inp.Suffix = trl.S{"de": "%", "en": "pct"}
				inp.HAlignLabel = qst.HLeft
				inp.HAlignControl = qst.HLeft
			}

		}

	}

	//
	// page 3 - inflation
	{
		p := q.AddPage()
		p.Label = trl.S{"de": "Inflation und Zinsen", "en": "Inflation and rates"}
		p.Short = trl.S{"de": "Inflation,<br>Zinsen", "en": "Inflation,<br>Rates"}
		p.AestheticCompensation = 5
		p.Width = 80

		//
		// gr1
		{
			labels123Matrix := []trl.S{
				{
					"de": "Euroraum",
					"en": "Euro area",
				},
				{
					"de": "Deutschland",
					"en": "Germany",
				},
			}
			names1stMatrix := []string{
				"pi_ez",
				"pi_deu",
			}
			gr := p.AddRadioMatrixGroup(labelsIncreaseDecrease(), names1stMatrix, labels123Matrix)
			gr.Cols = 5 // necessary, otherwise no vspacers
			gr.OddRowsColoring = true
			gr.Width = 90
			gr.Label = trl.S{
				"de": "4.",
				"en": "4.",
			}
			gr.Desc = trl.S{
				"de": "Die jährl. gesamtwirtschaftl. Inflationsrate wird mittelfristig (<b>6</b>&nbsp;Mo.)",
				"en": "Medium term (<b>6</b>&nbsp;months) yearly overall inflation rate will",
			}

		}

		{
			gr := p.AddGroup()
			gr.Cols = 100 // necessary, otherwise no vspacers
			gr.Label = trl.S{"de": "5a.", "en": "5a."}
			gr.Desc = trl.S{
				"de": "Die <b>kurzfristigen</b> Zinsen (3-Mo.-Interbanksätze) im <b>Euroraum</b> erwarten wir auf Sicht von 6&nbsp;Monaten",
				"en": "We expect <b>short term</b> interest rates (3 months interbank) in the <b>euro area</b>",
			}
			gr.HeaderBottomVSpacers = 1

			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = "i_ez_low"
				inp.MaxChars = 4
				inp.Validator = "inRange20"

				inp.ColSpanLabel = 10
				// inp.CSSLabel = "special-line-height-higher"
				inp.ColSpanControl = 12
				inp.Desc = trl.S{
					"de": "zwischen&nbsp;",
					"en": "between&nbsp;",
				}
				inp.Suffix = trl.S{"de": "%", "en": "pct"}
				inp.HAlignLabel = qst.HRight
				inp.HAlignControl = qst.HLeft
			}

			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = "i_ez_high"
				inp.MaxChars = 4
				inp.Validator = "inRange20"

				inp.ColSpanLabel = 4
				inp.ColSpanControl = 74
				inp.Desc = trl.S{
					"de": "und",
					"en": "and",
				}
				inp.Suffix = trl.S{"de": "%", "en": "pct"}
				inp.HAlignLabel = qst.HLeft
				inp.HAlignControl = qst.HLeft
			}

			// row below
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpanLabel = 11
				inp.Desc = trl.S{
					"de": " &nbsp;",
					"en": " &nbsp;",
				}
			}
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpanLabel = 89
				inp.Desc = trl.S{
					"de": " [zentrales 90% Konfidenzintervall]",
					"en": " [central 90&nbsp;pct confidence interval]",
				}
				inp.CSSLabel = "special-input-textblock-smaller"
			}

		}

		{
			gr := p.AddGroup()
			gr.Cols = 100 // necessary, otherwise no vspacers
			gr.Label = trl.S{"de": "5b.", "en": "5b."}
			gr.Desc = trl.S{
				"de": "Die <b>langfristigen</b> Zinsen (Renditen 10jg. Staatsanleihen) in <b>Deutschland</b> erwarten wir auf Sicht von 6&nbsp;Monaten",
				"en": "We expect <b>long term</b> interest rates in <b>Germany</b> in 6&nbsp;months",
			}
			gr.HeaderBottomVSpacers = 1

			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = "r_deu_low"
				inp.MaxChars = 4
				inp.Validator = "inRange100"

				inp.ColSpanLabel = 10
				// inp.CSSLabel = "special-line-height-higher"
				inp.ColSpanControl = 12
				inp.Desc = trl.S{
					"de": "zwischen&nbsp;",
					"en": "between&nbsp;",
				}
				inp.Suffix = trl.S{"de": "%", "en": "pct"}
				inp.HAlignLabel = qst.HRight
				inp.HAlignControl = qst.HLeft
			}

			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = "r_deu_high"
				inp.MaxChars = 4
				inp.Validator = "inRange100"

				inp.ColSpanLabel = 4
				inp.ColSpanControl = 74
				inp.Desc = trl.S{
					"de": "und",
					"en": "and",
				}
				inp.Suffix = trl.S{"de": "%", "en": "pct"}
				inp.HAlignLabel = qst.HLeft
				inp.HAlignControl = qst.HLeft
			}

			// row below
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpanLabel = 11
				inp.Desc = trl.S{
					"de": " &nbsp;",
					"en": " &nbsp;",
				}
			}
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpanLabel = 89
				inp.Desc = trl.S{
					"de": " [zentrales 90% Konfidenzintervall]",
					"en": " [central 90&nbsp;pct confidence interval]",
				}
				inp.CSSLabel = "special-input-textblock-smaller"
			}

		}

	}

	//
	// page 4 - Credit situation
	{
		p := q.AddPage()
		p.Section = trl.S{"de": "Kreditsituation", "en": "Credit situation"}
		p.Label = trl.S{"de": "Markt", "en": "Market"}
		p.Short = trl.S{"de": "Kredit-<br>situation:<br>Markt", "en": "Credit<br>situation:<br>Market"}

		{
			names3rdMatrix := []string{
				"cd_deu",
				"cs_deu",
			}
			labels123Matrix := []trl.S{
				{
					"de": "Kreditnachfrage",
					"en": "Credit demand",
				},
				{
					"de": "Kreditangebot",
					"en": "Credit supply",
				},
			}

			gr := p.AddRadioMatrixGroup(labelsVeryHighVeryLow(), names3rdMatrix, labels123Matrix)
			gr.Label = trl.S{"de": "6a.", "en": "6a."}
			gr.Desc = trl.S{
				"de": "Wie schätzen Sie die Kreditsituation in Deutschland ein?",
				"en": "How do you assess credit conditions in Germany?",
			}
			gr.Cols = 7 // necessary, otherwise no vspacers
			gr.OddRowsColoring = true
		}

		{
			names3rdMatrix := []string{
				"c0_6",
				"c0_24",
			}
			labels123Matrix := []trl.S{
				{
					"de": "mittelfristig (<b>6</b>&nbsp;Mo.)",
					"en": "medium term (<b>6</b>&nbsp;months)",
				},
				{
					"de": "langfristig (<b>24</b>&nbsp;Mo.)",
					"en": "long term (<b>24</b>&nbsp;months)",
				},
			}

			gr := p.AddRadioMatrixGroup(labelsStrongIncreaseStrongDecrease(), names3rdMatrix, labels123Matrix)
			gr.Label = trl.S{"de": "6b.", "en": "6b."}
			gr.Desc = trl.S{
				"de": "Das (saisonbereinigte) Gesamtvolumen der Neukreditvergabe in Deutschland wird",
				"en": "The seasonally adjusted volume of new credit in Germany will",
			}
			gr.Cols = 7 // necessary, otherwise no vspacers
			gr.OddRowsColoring = true
		}

		{
			names3rdMatrix := []string{
				"cd_24_le",
				"cd_24_sme",
				"cd_24_re",
				"cd_24_co",
			}
			labels123Matrix := []trl.S{
				{
					"de": "Großunternehmen",
					"en": "Large enterprises",
				},
				{
					"de": "KMU",
					"en": "Small+medium enterprises",
				},
				{
					"de": "Immobilienkredite",
					"en": "Real estate credit",
				},
				{
					"de": "Konsumentenkredite",
					"en": "Consumer credit",
				},
			}

			gr := p.AddRadioMatrixGroup(labelsStrongIncreaseStrongDecrease(), names3rdMatrix, labels123Matrix)
			gr.Label = trl.S{"de": "6c.", "en": "6c."}
			gr.Desc = trl.S{
				"de": "Die (saisonbereinigte) Kreditnachfrage wird mittelfristig (<b>6</b>&nbsp;Mo.)",
				"en": "The seasonally adjusted credit demand medium term (<b>6</b>&nbsp;months) will be",
			}
			gr.Cols = 7 // necessary, otherwise no vspacers
			gr.OddRowsColoring = true
		}

	}

	//
	// page 5 - Credit influence factors
	{
		p := q.AddPage()
		// page.Section = trl.S{"de": "Kreditsituation", "en": "Credit situation"}
		p.Label = trl.S{"de": "Einflussfaktoren", "en": "Influence factors"}
		p.Short = trl.S{"de": "Einfluss-<br>faktoren", "en": "Influence<br>factors"}

		{
			names3rdMatrix := []string{
				"c_inf_6_dr",
				"c_inf_6_ri",
				"c_inf_6_re",
				"c_inf_6_ce",
				"c_inf_6_rg",
				"c_inf_6_ep",
			}
			labels123Matrix := []trl.S{
				{
					"de": "Ausfallrisiken",
					"en": "Default risk",
				},
				{
					"de": "Risikotragfähigkeit",
					"en": "Risk profile",
				},
				{
					"de": "Refinanzierung",
					"en": "Refinancing",
				},
				{
					"de": "Wettbewerbssituation",
					"en": "Competitive environment",
				},
				{
					"de": "Regulierung",
					"en": "Regulation",
				},
				{
					"de": "EZB Politik",
					"en": "ECB policy",
				},
			}

			gr := p.AddRadioMatrixGroup(labelsVeryPositiveVeryNegative(), names3rdMatrix, labels123Matrix)
			gr.Label = trl.S{"de": "6d.", "en": "6d."}
			gr.Desc = trl.S{
				"de": "Wie schätzen Sie den Einfluss folgender Faktoren auf die mittelfristige (<b>6</b>&nbsp;Mo.) Veränderung des Kreditangebots ein?",
				"en": "How do you assess the influence of following factors on the medium term (<b>6</b>&nbsp;months) change of credit supply?",
			}
			gr.Cols = 7 // necessary, otherwise no vspacers
			gr.OddRowsColoring = true
		}

		{
			names3rdMatrix := []string{
				"c_std_6_le",
				"c_std_6_sme",
				"c_std_6_re",
				"c_std_6_co",
			}
			labels123Matrix := []trl.S{
				{
					"de": "Großunternehmen",
					"en": "Large enterprises",
				},
				{
					"de": "KMU",
					"en": "Small+medium enterprises",
				},
				{
					"de": "Immobilienkredite",
					"en": "Real estate credit",
				},
				{
					"de": "Konsumentenkredite",
					"en": "Consumer credit",
				},
			}

			gr := p.AddRadioMatrixGroup(labelsStrongIncreaseStrongDecrease(), names3rdMatrix, labels123Matrix)
			gr.Label = trl.S{"de": "6e.", "en": "6e."}
			gr.Desc = trl.S{
				"de": "Die (saisonbereinigte) Kreditstandards für Neukredite werden mittelfristig (<b>6</b>&nbsp;Mo.)",
				"en": "The seasonally adjusted credit standards medium term (<b>6</b>&nbsp;months) will",
			}
			gr.Cols = 7 // necessary, otherwise no vspacers
			gr.OddRowsColoring = true
		}

	}

	//
	// page 6 - Financial markets
	{
		p := q.AddPage()
		p.Section = trl.S{"de": "Finanzmärkte", "en": "Financial markets"}
		p.Label = trl.S{"de": "Preise", "en": "Prices"}
		p.Short = trl.S{"de": "Finanz-<br>märkte:<br>Preise", "en": "Financial<br>markets:<br>Prices"}

		p.Width = 80

		{
			names3rdMatrix := []string{
				"sto_dax",
				"oil",
				"gold",
				"fx_usa",
			}
			labels123Matrix := []trl.S{
				{
					"de": "DAX",
					"en": "German DAX",
				},
				{
					"de": "Rohöl (Nordsee Brent)",
					"en": "Brent Crude",
				},
				{
					"de": "Gold",
					"en": "Gold",
				},
				{
					"de": "US-Dollar (ggü. €)",
					"en": "Dollar / Euro",
				},
			}

			gr := p.AddRadioMatrixGroup(labelsIncreaseDecrease(), names3rdMatrix, labels123Matrix)
			gr.Label = trl.S{"de": "7a.", "en": "7a."}
			gr.Desc = trl.S{
				"de": "Die folgenden Aktienindizes / Rohstoffpreise / Wechselkurse werden mittelfristig (<b>6</b>&nbsp;Mo.)",
				"en": "Following stock indices / raw materials / exchange rates will medium term (<b>6</b>&nbsp;months)",
			}

			gr.Cols = 5 // necessary, otherwise no vspacers
			gr.OddRowsColoring = true
		}

		gr := p.AddGroup()
		gr.Cols = 100
		gr.Label = trl.S{"de": "7b.", "en": "7b."}

		{
			inp := gr.AddInput()
			inp.Type = "number"
			inp.Name = "dax_6"
			inp.MaxChars = 6
			inp.Validator = "inRange50000"

			inp.ColSpanLabel = 55
			// inp.CSSLabel = "special-line-height-higher"
			inp.ColSpanControl = 45
			inp.Desc = trl.S{
				"de": `Den DAX erwarten wir in 6&nbsp;Monaten bei `,
				"en": "We expect the German DAX in 6&nbsp;month at",
			}
			inp.Suffix = trl.S{"de": "Punkten", "en": "points"}
			inp.HAlignLabel = qst.HLeft
			inp.HAlignControl = qst.HLeft
		}

		{
			inp := gr.AddInput()
			inp.Type = "number"
			inp.Name = "dax_6_low"
			inp.MaxChars = 6
			inp.Validator = "inRange50000"

			inp.ColSpanLabel = 55
			inp.ColSpanControl = 21
			inp.Desc = trl.S{
				"de": `Mit einer Wahrscheinlichkeit von 90&nbsp;Prozent wird der DAX dann zwischen `,
				"en": "With 90&nbsp;percent probability, the DAX will then be between",
			}
			inp.Suffix = trl.S{"de": "Punkten  &nbsp; und ", "en": "points &nbsp; and "}
			inp.HAlignLabel = qst.HLeft
			inp.HAlignControl = qst.HLeft
		}

		{
			inp := gr.AddInput()
			inp.Type = "number"
			inp.Name = "dax_6_high"
			inp.MaxChars = 6
			inp.Validator = "inRange50000"

			inp.ColSpanControl = 24
			inp.Suffix = trl.S{"de": "Punkten liegen.", "en": "points"}
			inp.HAlignLabel = qst.HLeft
			inp.HAlignControl = qst.HLeft
		}

		{

			// gr := p.AddRadioMatrixGroup(labelsOvervaluedFairUndervalued(), names3rdMatrix, labels123Matrix)
			gr := p.AddGroup()
			gr.Cols = 6 // necessary, otherwise no vspacers
			gr.Label = trl.S{"de": "7c.", "en": "7c."}
			gr.Desc = trl.S{
				"de": "Aus Sicht der Fundamentaldaten der DAX-Unternehmen ist der DAX derzeit",
				"en": "The fundamentals of the companies comprising the DAX make the DAX currently",
			}
			gr.HeaderBottomVSpacers = 1

			{
				inp := gr.AddInput()
				inp.Type = "radiogroup"
				inp.Name = "dax_fund"
				inp.CSSLabel = "special-input-left-padding"
				for i2, val := range labelsOvervaluedFairUndervalued() {
					rad := inp.AddRadio()
					rad.Label = val
					// rad.HAlign = qst.HRight
					rad.HAlign = qst.HLeft
					rad.Val = fmt.Sprintf("%v", i2)
				}
			}

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpanLabel = 2
				inp.Desc = trl.S{"de": " &nbsp; ", "en": " &nbsp; "}
			}

		}

		//
		//
		{
			gr := p.AddGroup()
			gr.Cols = 100
			gr.Label = trl.S{"de": "8.", "en": "8."}
			gr.Desc = trl.S{
				"de": "Die Wahrscheinlichkeit für ein Extremereignis im deutschen Finanzmarkt liegt",
				"en": "The probability for an extreme event in the German financial markets is",
			}
			gr.HeaderBottomVSpacers = 1

			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = "crash_low"
				inp.MaxChars = 4
				inp.Validator = "inRange100"

				inp.ColSpanLabel = 22
				inp.ColSpanControl = 12
				inp.Desc = trl.S{
					"de": " mittelfristig (<b>6</b>&nbsp;Mo.) bei  ",
					"en": " medium term (<b>6</b>&nbsp;months) at ",
				}
				inp.Suffix = trl.S{"de": "%", "en": "pct"}
				inp.HAlignLabel = qst.HRight
				inp.HAlignControl = qst.HLeft
			}

			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = "crash_high"
				inp.MaxChars = 4
				inp.Validator = "inRange100"

				inp.ColSpanLabel = 26
				inp.ColSpanControl = 40
				inp.Desc = trl.S{
					"de": "und langfristig (<b>24</b>&nbsp;Mo.) bei ",
					"en": "and long term (<b>24</b>&nbsp;months) at ",
				}

				inp.Suffix = trl.S{"de": "%", "en": "pct"}
				inp.HAlignLabel = qst.HRight
				inp.HAlignControl = qst.HLeft
			}
		}

	}

	err := addSeasonal1(&q)
	if err != nil {
		return nil, fmt.Errorf("Error adding seasonal1: %v", err)
	}

	err = addSeasonal2(&q)
	if err != nil {
		return nil, fmt.Errorf("Error adding seasonal2: %v", err)
	}

	//
	//
	// Finish questionnaire?  - one before last page
	{
		p := q.AddPage()
		p.Section = trl.S{"de": "Abschluss", "en": "Finish"}
		p.Label = trl.S{"de": "", "en": ""}
		p.Short = trl.S{"de": "Abschluss", "en": "Finish"}
		p.Width = 65

		{
			gr := p.AddGroup()
			gr.Cols = 1 // necessary, otherwise no vspacers
			gr.Label = trl.S{"de": "", "en": ""}
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{"de": "Kommentar zur Umfrage: ", "en": "Comment on the survey: "}
				inp.Desc = trl.S{
					"de": "Wollen Sie uns noch etwas mitteilen?",
					"en": "Any remarks or advice for us?",
				}
				inp.ColSpanLabel = 1
			}
			{
				inp := gr.AddInput()
				inp.Type = "textarea"
				inp.Name = "remark"
				inp.MaxChars = 300
				inp.ColSpanLabel = 1
				inp.ColSpanControl = 1
			}
		}

		{
			gr := p.AddGroup()
			// gr.Desc = trl.S{"de": "Abschluss", "en": "Finish"}
			gr.Width = 100
			gr.Cols = 2 // necessary, otherwise no vspacers
			{

				{
					inp := gr.AddInput()
					inp.Type = "radiogroup"
					inp.Name = "finished"
					inp.CSSLabel = "special-line-height-higher"
					inp.ColSpanLabel = 1
					inp.ColSpanControl = 1
					inp.Validator = "mustRadioGroup"
					{
						rad := inp.AddRadio()
						rad.HAlign = qst.HLeft
						// rad.HAlign = qst.HCenter
						rad.Label = trl.S{
							"de": "Zugang bleibt bestehen.  Daten können in weiteren Sitzungen geändert/ergänzt werden. <br>\n &nbsp;",
							"en": "Leave questionnaire open. Data  can be changed/completed&nbsp;in later sessions. <br>\n &nbsp;",
						}
						rad.Val = "2" // any other non null value
					}
					{
						rad := inp.AddRadio()
						rad.HAlign = qst.HLeft
						// rad.HAlign = qst.HCenter
						rad.Label = trl.S{
							"de": "Fragebogen ist abgeschlossen und kann nicht mehr geöffnet werden. <br>\n &nbsp;",
							"en": "Questionnaire is finished. No more edits. <br>\n &nbsp;",
						}
						rad.Val = qst.ValSet
					}

				}
			}

		}

		{
			gr := p.AddGroup()
			gr.Cols = 2 // necessary, otherwise no vspacers
			// gr.Width = 80
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{"de": "", "en": ""}
				inp.Desc = trl.S{
					"de": "Durch Klicken auf 'OK' erhalten Sie eine Zusammenfassung Ihrer Antworten",
					"en": "By Clicking 'OK' you receive a summary of your answers",
				}
				inp.ColSpanLabel = 1
				inp.CSSLabel = "special-line-height-higher"
			}
			{
				inp := gr.AddInput()
				inp.Type = "button"
				inp.Name = "submitBtn"
				inp.Response = fmt.Sprintf("%v", len(q.Pages)-1+1) // +1 since one page is appended below
				inp.Label = trl.S{"de": "", "en": ""}
				inp.Desc = trl.S{
					"de": "OK",
					"en": "OK",
				}
				inp.ColSpanControl = 1
				inp.AccessKey = "n"
				inp.HAlignControl = qst.HCenter
				inp.HAlignControl = qst.HLeft
			}
		}

		// page.ExampleSixColumnsLabelRight()

	}

	// Report of results
	{
		p := q.AddPage()
		p.Label = trl.S{"de": "Ihre Eingaben", "en": "Summary of results"}
		p.NoNavigation = true
		{
			gr := p.AddGroup()
			gr.Cols = 1
			{
				inp := gr.AddInput()
				inp.Type = "dynamic"
				inp.CSSLabel = "special-line-height-higher"
				inp.DynamicFunc = "RepsonseStatistics"
			}
			{
				inp := gr.AddInput()
				inp.Type = "dynamic"
				inp.CSSLabel = "special-line-height-higher"
				inp.DynamicFunc = "PersonalLink"
			}
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.CSSLabel = "special-line-height-higher"
				impr := trl.S{}
				for _, lc := range q.LangCodes {
					w1 := &strings.Builder{}
					err := tpl.RenderStaticContent(
						w1, "site-imprint.md", q.Survey.Type, lc,
					)
					if err != nil {
						log.Print(err)
					}
					impr[lc] = w1.String()

				}
				inp.Desc = impr
			}
		}
	}

	// quest.ClosingTime = time.Now()

	err = q.Validate()
	if err != nil {
		return nil, fmt.Errorf("Error validating questionnaire: %v", err)
	}
	q.Hyphenize()
	q.ComputeMaxGroups()
	if err := (&q).TranslationCompleteness(); err != nil {
		return &q, err
	}
	if err := (&q).Validate(); err != nil {
		return &q, err
	}

	return &q, nil
}
