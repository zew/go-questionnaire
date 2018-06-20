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

			{
				inp := gr.AddInput()
				inp.Type = "button"
				inp.Name = "submitBtn"
				inp.Response = "1"
				inp.Label = trl.S{
					"de": "Weiter",
					"en": "next",
				}
				inp.ColSpanControl = 1
				inp.HAlignControl = qst.HCenter
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
			inp.Validator = "inRange10000"

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
			inp.Validator = "inRange10000"

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
			inp.Validator = "inRange10000"

			inp.ColSpanControl = 24
			inp.Suffix = trl.S{"de": "Punkten liegen.", "en": "points"}
			inp.HAlignLabel = qst.HLeft
			inp.HAlignControl = qst.HLeft
		}

		{
			names3rdMatrix := []string{
				"dax_fund",
			}
			labels123Matrix := []trl.S{}

			gr := p.AddRadioMatrixGroup(labelsOvervaluedFairUndervalued(), names3rdMatrix, labels123Matrix)
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
			gr.Cols = 100
			{
				inp := gr.AddInput()
				inp.Type = "text"
				inp.Name = "crash_low"
				inp.MaxChars = 4
				inp.Validator = "inRange100"

				inp.ColSpanLabel = 55
				inp.CSSLabel = "vert-wider"
				inp.ColSpanControl = 33
				inp.Desc = trl.S{
					"de": `Die Wahrscheinlichkeit für ein Extremereignis im deutschen Finanzmarkt liegt mittefristig (<b>6</b> Mo.) bei`,
					"en": "The probability for an extreme event in the German financial markets medium term (<b>6</b> months) is at ",
				}
				inp.Suffix = trl.S{
					"de": "%  &nbsp; und langfristig (<b>24</b> Mo.) bei",
					"en": "pct  &nbsp; and long term (<b>24</b> months) at ",
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

				inp.ColSpanControl = 12
				inp.Suffix = trl.S{"de": "%", "en": "pct"}
				inp.HAlignLabel = qst.HLeft
				inp.HAlignControl = qst.HLeft
			}
		}

	}

	//
	//
	// Finish
	{
		p := q.AddPage()
		p.Section = trl.S{"de": "Abschluss", "en": " &nbsp; Finish"}
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
				inp.ColSpanControl = 1 // ignored for radiogroup

				{
					rad := inp.AddRadio()
					rad.HAlign = qst.HLeft
					rad.HAlign = qst.HCenter
					rad.Label = trl.S{
						"de": "\nFragebogen ist abgeschlossen \nund kann nicht mehr geöffnet werden.",
						"en": "\nQuestionaire is finished.\nNo more edits.",
					}
				}
				{
					rad := inp.AddRadio()
					rad.HAlign = qst.HLeft
					rad.HAlign = qst.HCenter
					rad.Label = trl.S{
						"de": "Zugang bleibt bestehen.  \nDaten können in weiteren Sitzungen \ngeändert/ergänzt werden.",
						"en": "Leave questionaire open. \nData  can be changed/completed     \nin later sessions.",
					}
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
