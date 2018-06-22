package fmt

import (
	"fmt"
	"log"

	"github.com/zew/go-questionaire/qst"
	"github.com/zew/go-questionaire/tpl"
	"github.com/zew/go-questionaire/trl"
)

// Create creates a JSON file for a financial markets survey
func Create() *qst.QuestionaireT {

	q := qst.QuestionaireT{}
	q.Survey = qst.NewSurvey("fmt")
	q.LangCodes = map[string]string{"de": "Deutsch", "en": "English"}
	q.LangCode = "de" // default language
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
				inp.CSSLabel = "vert-wider"
				inp.ColSpanLabel = 3
				impr := trl.S{}
				for lc := range q.LangCodes {
					cnt, err := tpl.MarkDownFromFile("./static/doc/data-protection.md", lc)
					if err != nil {
						log.Print(err)
					}
					impr[lc] = cnt
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
				inp.CSSLabel = "vert-wider"

				inp.Label = trl.S{"de": " ", "en": " "}
				inp.Desc = trl.S{"de": "Sind Sie die angeschriebene Person?", "en": "Are you the addressee?"}

				inp.ColSpanLabel = 1
				inp.ColSpanControl = 2
				{
					rad := inp.AddRadio()
					rad.HAlign = qst.HLeft
					rad.HAlign = qst.HCenter
					rad.Label = trl.S{
						"de": "Ja, ich bin die angeschriebene Person.",
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

			// gr = p.AddGroup()
			// gr.Cols = 5
			// gr.Width = 75
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
				"de": "Die gesamtwirtschaftliche Situation wird sich mittelfristig (<b>6</b> Mo.)",
				"en": "The overall economic situation medium term (<b>6</b> months) will",
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
				"de": "Die gesamtwirtschaftliche Situation wird sich langfristig (<b>24</b> Mo.)",
				"en": "The overall economic situation long term (<b>24</b> months) will",
			}

		}

	}

	// page 2
	{
		p := q.AddPage()
		p.Label = trl.S{"de": "Wachstum", "en": "Growth"}
		p.AestheticCompensation = 8

		{
			gr := p.AddGroup()
			gr.Cols = 5 // necessary, otherwise no vspacers
			gr.Label = trl.S{"de": "3a.", "en": "3a."}
			{
				inp := gr.AddInput()
				inp.Type = "text"
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
				inp.Type = "text"
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
				inp.Type = "text"
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
				inp.Type = "text"
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
		p.Label = trl.S{"de": "Inflation und Zinsen", "en": "Inflation and Rates"}
		p.AestheticCompensation = 5

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
			gr.Width = 62
			gr.Label = trl.S{
				"de": "4.",
				"en": "4.",
			}
			gr.Desc = trl.S{
				"de": "Die jährl. gesamtwirtschaftl. Inflationsrate wird mittelfristig (<b>6</b> Mo.)",
				"en": "Medium term (<b>6</b> months) yearly overall inflation rate will",
			}

		}

		{
			gr := p.AddGroup()
			gr.Cols = 100 // necessary, otherwise no vspacers
			gr.Label = trl.S{"de": "5a.", "en": "5a."}
			{
				inp := gr.AddInput()
				inp.Type = "text"
				inp.Name = "i_ez_low"
				inp.MaxChars = 4
				inp.Validator = "inRange20"

				inp.ColSpanLabel = 40
				inp.CSSLabel = "vert-wider"
				inp.ColSpanControl = 5
				inp.Desc = trl.S{
					"de": `Die <b>kurzfristigen</b> Zinsen (3-Mo.-Interbanksätze) im <b>Euroraum</b> erwarten wir auf Sicht von 6&nbsp;Monaten<br>
					 [zentrales 90%-Konfidenzintervall] zwischen`,
					"en": "We expect <b>short term</b> interest rates (3 months interbank) in the <b>euro area</b> between",
				}
				inp.Suffix = trl.S{"de": "%", "en": "pct"}
				inp.HAlignLabel = qst.HLeft
				inp.HAlignControl = qst.HLeft
			}

			{
				inp := gr.AddInput()
				inp.Type = "text"
				inp.Name = "i_ez_high"
				inp.MaxChars = 4
				inp.Validator = "inRange20"

				inp.ColSpanLabel = 2
				inp.ColSpanControl = 10
				inp.Desc = trl.S{
					"de": "und",
					"en": "and",
				}
				inp.Suffix = trl.S{"de": "%", "en": "pct"}
				inp.HAlignLabel = qst.HLeft
				inp.HAlignControl = qst.HLeft
			}
		}

		{
			gr := p.AddGroup()
			gr.Cols = 100 // necessary, otherwise no vspacers
			gr.Label = trl.S{"de": "5b.", "en": "5b."}

			{
				inp := gr.AddInput()
				inp.Type = "text"
				inp.Name = "r_deu_low"
				inp.MaxChars = 4
				inp.Validator = "inRange100"

				inp.ColSpanLabel = 40
				inp.CSSLabel = "vert-wider"
				inp.ColSpanControl = 5
				inp.Desc = trl.S{
					"de": `Die <b>langfristigen</b> Zinsen (Renditen 10jg. Staatsanleihen) in <b>Deutschland</b> erwarten wir auf Sicht von 6&nbsp;Monaten<br> 
					[zentrales 90%-Konfidenzintervall] zwischen`,
					"en": "We expect <b>long term</b> interest rates in <b>Germany</b> in 6 months between",
				}
				inp.Suffix = trl.S{"de": "%", "en": "pct"}
				inp.HAlignLabel = qst.HLeft
				inp.HAlignControl = qst.HLeft
			}

			{
				inp := gr.AddInput()
				inp.Type = "text"
				inp.Name = "r_deu_high"
				inp.MaxChars = 4
				inp.Validator = "inRange100"

				inp.ColSpanLabel = 2
				inp.ColSpanControl = 10
				inp.Desc = trl.S{
					"de": "und",
					"en": "and",
				}
				inp.Suffix = trl.S{"de": "%", "en": "pct"}
				inp.HAlignLabel = qst.HLeft
				inp.HAlignControl = qst.HLeft
			}

		}

	}

	//
	// page 4 - Credit situation
	{
		p := q.AddPage()
		p.Section = trl.S{"de": "Kreditsituation", "en": "Credit situation"}
		p.Label = trl.S{"de": "Markt", "en": "Market"}

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
					"de": "mittelfristig (<b>6</b> Mo.)",
					"en": "medium term (<b>6</b> months)",
				},
				{
					"de": "langfristig (<b>24</b> Mo.)",
					"en": "long term (<b>24</b> months)",
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
				"de": "Die (saisonbereinigte) Kreditnachfrage wird mittelfristig (<b>6</b> Mo.)",
				"en": "The seasonally adjusted credit demand medium term (<b>6</b> months) will be",
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
				"de": "Wie schätzen Sie den Einfluss folgender Faktoren auf die mittelfristige (<b>6</b> Mo.) Veränderung des Kreditangebots ein?",
				"en": "How do you assess the influence of following factors on the medium term (<b>6</b> months) change of credit supply?",
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
				"de": "Die (saisonbereinigte) Kreditstandards für Neukredite werden mittelfristig (<b>6</b> Mo.)",
				"en": "The seasonally adjusted credit standards medium term (<b>6</b> months) will",
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
				"de": "Die folgenden Aktienindizes / Rohstoffpreise / Wechselkurse werden mittefristig (<b>6</b> Mo.)",
				"en": "Following stock indices / raw materials / exchange rates will medium term (<b>6</b> months)",
			}

			gr.Cols = 5 // necessary, otherwise no vspacers
			gr.OddRowsColoring = true
		}

		gr := p.AddGroup()
		gr.Label = trl.S{"de": "7b.", "en": "7b."}
		gr.Cols = 100
		{
			inp := gr.AddInput()
			inp.Type = "text"
			inp.Name = "dax_6"
			inp.MaxChars = 6
			inp.Validator = "inRange50000"

			inp.ColSpanLabel = 55
			inp.CSSLabel = "vert-wider"
			inp.ColSpanControl = 45
			inp.Desc = trl.S{
				"de": `Den DAX erwarten wir in 6 Monaten bei `,
				"en": "We expect the German DAX in 6 month at",
			}
			inp.Suffix = trl.S{"de": "Punkten", "en": "points"}
			inp.HAlignLabel = qst.HLeft
			inp.HAlignControl = qst.HLeft
		}

		{
			inp := gr.AddInput()
			inp.Type = "text"
			inp.Name = "dax_6_low"
			inp.MaxChars = 6
			inp.Validator = "inRange50000"

			inp.ColSpanLabel = 55
			inp.CSSLabel = "vert-wider"
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
			inp.Type = "text"
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
			inp := gr.AddInput()
			inp.Type = "radiogroup"
			inp.Name = "dax_fund"
			for i2, val := range labelsOvervaluedFairUndervalued() {
				rad := inp.AddRadio()
				rad.Label = val
				// rad.HAlign = qst.HRight
				rad.HAlign = qst.HLeft
				rad.Val = fmt.Sprintf("%v", i2)
			}
			gr.Label = trl.S{"de": "7c.", "en": "7c."}
			gr.Desc = trl.S{
				"de": "Aus Sicht der Fundamentaldaten der DAX-Unternehmen ist der DAX derzeit",
				"en": "The fundamentals of the companies comprising the DAX make the DAX currently",
			}
			gr.Cols = 3 // necessary, otherwise no vspacers
		}

		//
		//
		{
			gr := p.AddGroup()
			gr.Label = trl.S{"de": "8.", "en": "8."}
			gr.Desc = trl.S{
				"de": "Die Wahrscheinlichkeit für ein Extremereignis im deutschen Finanzmarkt liegt",
				"en": "The probability for an extreme event in the German financial markets is",
			}
			gr.Cols = 100
			{
				inp := gr.AddInput()
				inp.Type = "text"
				inp.Name = "crash_low"
				inp.MaxChars = 4
				inp.Validator = "inRange100"

				inp.ColSpanLabel = 22
				inp.CSSLabel = "vert-wider"
				inp.ColSpanControl = 34
				inp.Desc = trl.S{
					"de": ` mittefristig (<b>6</b> Mo.) bei  &nbsp;  `,
					"en": " medium term (<b>6</b> months) at  &nbsp;  ",
				}
				inp.Suffix = trl.S{
					"de": "%  &nbsp; und langfristig (<b>24</b> Mo.) bei &nbsp; ",
					"en": "pct  &nbsp; and long term (<b>24</b> months) at &nbsp; ",
				}
				inp.HAlignLabel = qst.HLeft
				inp.HAlignControl = qst.HLeft
			}

			{
				inp := gr.AddInput()
				inp.Type = "text"
				inp.Name = "crash_high"
				inp.MaxChars = 4
				inp.Validator = "inRange100"

				inp.ColSpanControl = 44
				inp.Suffix = trl.S{"de": "%", "en": "pct"}
				inp.HAlignLabel = qst.HLeft
				inp.HAlignControl = qst.HLeft
			}
		}

	}

	// Seasonal
	{

		// Special may questions revolve around the month of the Quarter.
		//
		// 	Business cycle & drivers: 		           Januar, April, Juli, October
		// 	Inflation, drivers, central bank rates:    Februar, May, August, November
		// 	Free special questoins:                    March, June, September, December
		if monthOfQuarter() == 1 || true {

			p := q.AddPage()
			p.Section = trl.S{"de": "Sonderfrage", "en": "Special"}
			p.Label = trl.S{"de": "Prognosetreiber Wachstum", "en": "Growth drivers"}
			p.Width = 80

			{
				gr := p.AddGroup()
				gr.Cols = 9
				gr.Label = trl.S{
					"de": "1.",
					"en": "1.",
				}
				gr.Desc = trl.S{
					"de": "Punktprognose der Wachstumsrate des deutschen BIP",
					"en": "Forecast growth rate German GDP",
				}

				{
					inp := gr.AddInput()
					inp.Type = "textblock"
					// inp.Label = trl.S{
					// 	"de": "Prognose:<br>\n",
					// 	"en": "Forecast:<br>\n",
					// }
					inp.Desc = trl.S{
						"de": " <br>\n Prognose Wachstum des BIP <b>je Quartal</b> <br>\n (real, saisonbereinigt, nicht annualisiert) <br>\n <br>\n ",
						"en": " <br>\n Forecast <b>quarterly</b> GDP growth <br>\n(real, seasonally adjusted, non annualized) <br>\n <br>\n ",
					}
					inp.ColSpanLabel = 3
				}
				{
					inp := gr.AddInput()
					inp.Type = "text"
					inp.Name = "xquart1"
					inp.MaxChars = 4
					inp.Validator = "inRange20"
					inp.Desc = trl.S{
						"de": nextQ(-1) + " &nbsp; ",
						"en": nextQ(-1) + " &nbsp; ",
					}
					inp.Suffix = trl.S{
						"de": "%",
						"en": "pct",
					}
					inp.HAlignLabel = qst.HRight

				}
				{
					inp := gr.AddInput()
					inp.Type = "text"
					inp.Name = "xquart2"
					inp.MaxChars = 4
					inp.Validator = "inRange20"
					inp.Desc = trl.S{
						"de": nextQ(0) + " &nbsp; ",
						"en": nextQ(0) + " &nbsp; ",
					}
					inp.Suffix = trl.S{
						"de": "%",
						"en": "pct",
					}
					inp.HAlignLabel = qst.HRight
				}
				{
					inp := gr.AddInput()
					inp.Type = "text"
					inp.Name = "xquart3"
					inp.MaxChars = 4
					inp.Validator = "inRange20"
					inp.Desc = trl.S{
						"de": nextQ() + " &nbsp; ",
						"en": nextQ() + " &nbsp; ",
					}
					inp.Suffix = trl.S{
						"de": "%",
						"en": "pct",
					}
					inp.HAlignLabel = qst.HRight

				}

				{
					inp := gr.AddInput()
					inp.Type = "textblock"
					// inp.Label = trl.S{
					// 	"de": "Prognose:<br>\n",
					// 	"en": "Forecast:<br>\n",
					// }
					inp.Desc = trl.S{
						"de": "Prognose Wachstum des BIP aufs&nbsp;<b>Jahr</b> <br>\n(real, saisonbereinigt)",
						"en": "Forecast GDP growth per&nbsp;<b>year</b> <br>\n(real, seasonally adjusted)",
					}
					inp.ColSpanLabel = 3

				}

				{
					inp := gr.AddInput()
					inp.Type = "text"
					inp.Name = "xyear1"
					inp.MaxChars = 4
					inp.Validator = "inRange20"
					inp.Desc = trl.S{
						"de": nextY(-1) + " &nbsp; ",
						"en": nextY(-1) + " &nbsp; ",
					}
					inp.Suffix = trl.S{
						"de": "%",
						"en": "pct",
					}
					inp.HAlignLabel = qst.HRight

				}
				{
					inp := gr.AddInput()
					inp.Type = "text"
					inp.Name = "xyear2"
					inp.MaxChars = 4
					inp.Validator = "inRange20"
					inp.Desc = trl.S{
						"de": nextY(0) + " &nbsp; ",
						"en": nextY(0) + " &nbsp; ",
					}
					inp.Suffix = trl.S{
						"de": "%",
						"en": "pct",
					}
					inp.HAlignLabel = qst.HRight
				}
				{
					inp := gr.AddInput()
					inp.Type = "text"
					inp.Name = "xyear3"
					inp.MaxChars = 4
					inp.Validator = "inRange20"
					inp.Desc = trl.S{
						"de": nextY() + " &nbsp; ",
						"en": nextY() + " &nbsp; ",
					}
					inp.Suffix = trl.S{
						"de": "%",
						"en": "pct",
					}
					inp.HAlignLabel = qst.HRight

				}
				// // in between label
				// {
				// 	inp := gr.AddInput()
				// 	inp.Type = "textblock"
				// 	inp.Desc = trl.S{
				// 		"de": `Wachstumsrate des realen BIP aufs Jahr`,
				// 		"en": "Real growth rate per year",
				// 	}
				// 	inp.ColSpanLabel = 7
				// 	inp.CSSLabel = "textblock-smaller"
				// }

			}

			// gr1
			{
				labels123Matrix := []trl.S{
					{
						"de": "Konjunkturdaten Deutschland",
						"en": "Business cycle data Germany",
					},
					{
						"de": "Exportmärkte",
						"en": "Export markets",
					},
					{
						"de": "Wechselkurse",
						"en": "Exchange rates",
					},
					{
						"de": "EZB-Geldpolitik",
						"en": "ECB monetary policy",
					},
					{
						"de": "US-Geldpolitik",
						"en": "FED monetary policy",
					},
					{
						"de": "Geopol. Ereignisse",
						"en": "Geopolitical events",
					},
					{
						"de": "Regierungsbildung Deutschland",
						"en": "Government formation Germany",
					},
					{
						"de": "Sonstige",
						"en": "Other",
					},
				}
				names1stMatrix := []string{
					"infl_cycle_data_deu",
					"infl_exp_markets",
					"infl_exch_rates",
					"infl_mp_ecb",
					"infl_mp_fed",
					"infl_geopol",
					"infl_gvt_form_deu",
					"infl_other",
				}
				gr := p.AddRadioMatrixGroup(labelsStronglyPositiveStronglyNegativeInfluence(),
					names1stMatrix, labels123Matrix, 2)
				gr.Cols = 8 // necessary, otherwise no vspacers
				gr.OddRowsColoring = true
				gr.Label = trl.S{
					"de": "2.",
					"en": "2.",
				}
				gr.Desc = trl.S{
					"de": "Haben Entwicklungen in den folgenden Bereichen Sie zu einer Revision (ggü. Vormonat) Ihrer Konjunkturprognosen für die deutsche Wirtschaft bewogen und wenn ja in welche Richtung?",
					"en": "Which developments have lead you to change your assessment of the business cycle outlook for the German economy compared to the previous month",
				}

				{
					inp := gr.AddInput()
					inp.Type = "text"
					inp.Name = "other_cycle_infl"
					inp.Desc = trl.S{
						"de": "Wenn sonstige - welche ?",
						"en": "If other - which?",
					}
					inp.MaxChars = 30
					inp.ColSpanLabel = 2
					inp.ColSpanControl = 6
					inp.CSSControl = "input-smaller"
				}

			}

		}
	}

	//
	//
	// Finish
	{
		p := q.AddPage()
		p.Section = trl.S{"de": "Abschluss", "en": "Finish"}
		p.Label = trl.S{"de": "", "en": ""}
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
			gr.Cols = 4 // necessary, otherwise no vspacers
			// gr.Desc = trl.S{"de": "Abschluss", "en": "Finish"}
			gr.Width = 100
			{

				inp := gr.AddInput()
				inp.Type = "radiogroup"
				inp.Name = "finished"
				inp.CSSLabel = "vert-wider"

				inp.Label = trl.S{"de": "Abschluss", "en": "Finish"}
				inp.Desc = trl.S{"de": "", "en": ""}

				inp.ColSpanLabel = 1
				inp.ColSpanControl = 1
				inp.Validator = "mustRadioGroup"

				{
					rad := inp.AddRadio()
					rad.HAlign = qst.HLeft
					rad.HAlign = qst.HCenter
					rad.Label = trl.S{
						"de": "Fragebogen ist abgeschlossen <br>\nund kann nicht mehr geöffnet werden.",
						"en": "Questionaire is finished.\nNo more edits.",
					}
					rad.Val = "1"
				}
				{
					rad := inp.AddRadio()
					rad.HAlign = qst.HLeft
					rad.HAlign = qst.HCenter
					rad.Label = trl.S{
						"de": "Zugang bleibt bestehen.  \nDaten können in weiteren Sitzungen \ngeändert/ergänzt werden.",
						"en": "Leave questionaire open. \nData  can be changed/completed     \nin later sessions.",
					}
					rad.Val = "2"
				}
			}

		}

		{
			gr := p.AddGroup()
			gr.Cols = 4 // necessary, otherwise no vspacers
			gr.Width = 80
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{"de": "", "en": ""}
				inp.Desc = trl.S{
					"de": "Durch Klicken auf 'OK' erhalten Sie eine Zusammenfassung Ihrer Antworten.",
					"en": "By Clicking 'OK' you receive a summary of your answers.",
				}
				inp.ColSpanLabel = 2
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
				inp.HAlignControl = qst.HCenter
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
				inp.CSSLabel = "vert-wider"
				inp.DynamicFunc = "RepsonseStatistics"
			}
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.CSSLabel = "vert-wider"
				impr := trl.S{}
				for lc := range q.LangCodes {
					cnt, err := tpl.MarkDownFromFile("./static/doc/site-imprint.md", lc)
					if err != nil {
						log.Print(err)
					}
					impr[lc] = cnt
				}
				inp.Desc = impr
			}
		}
	}

	// quest.ClosingTime = time.Now()

	err := q.Validate()
	if err != nil {
		log.Fatalf("Error validating questionaire: %v", err)
	}

	return &q
}
