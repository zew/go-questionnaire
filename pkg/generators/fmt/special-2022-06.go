package fmt

import (
	"github.com/zew/go-questionnaire/pkg/css"
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

func special202206(q *qst.QuestionnaireT) error {

	cond := false
	cond = cond || q.Survey.Year == 2022 && q.Survey.Month == 6
	if !cond {
		return nil
	}

	{
		page := q.AddPage()
		page.Short = trl.S{
			"de": "Sonderfrage:<br>Stablecoins",
			"en": "Special:<br>Stablecoins",
		}
		page.Label = trl.S{
			"de": "",
			"en": "",
		}
		// page.NoNavigation = true
		// page.SuppressProgressbar = true

		page.WidthMax("calc(100% - 1.2rem)")
		page.WidthMax("46rem")

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
						<h3>Sonderfragen zu den systemischen Risiken von Stablecoins </h3>

						<p style="text-align: justify:">
						Inmitten von allgemeinen Turbulenzen an den Märkten für Kryptowerten kollabierte Mitte Mai 2022 die digitale Währung Terra. Terra war als Stablecoin konzipiert, d.h. eine digitale Währung, die den Wert von einem US-Dollar haben soll. Auch Tether, der größte Stablecoin nach Marktkapitalisierung  (am 1. Juni, 2022 circa 72,5 Milliarden US-Dollar), entfernte sich von seinem Zielpreis von 1 US-Dollar und fiel bis auf einem Marktpreis von knapp 0,95 US-Dollar. Zwar hat sich der Marktpreis von Tether seither wieder erholt. Jedoch war die Differenz zwischen dem Zielpreis von 1 US-Dollar und dem Marktpreis am 1. Juni 2022 noch deutlich größer als vor den Turbulenzen. 
						</p>

						<p style="text-align: justify:">
						Tether ist ein so genannter Asset Backed Stablecoin. Der Emittent des Stablecoins verspricht üblicherweise für jeden emittierten Stablecoin Finanzwerte im Wert von 1 US-Dollar (oder Euro etc.) zu halten. Außerdem verspricht der Emittent den Haltern des Stablecoins, dass sie diese jederzeit im Verhältnis von 1:1 gegen US-Dollar eintauschen können. Unter normalen Marktbedingungen sorgt ein Arbitragemechanismus dafür, dass der Marktpreis des Stablecoins immer (ungefähr) 1 US-Dollar beträgt. So kaufen Arbitrageure wenn der Preis eines Stablecoins unter 1 US-Dollar fällt die entsprechenden Stablecoins auf den Kryptomärkten und tauschen diese beim Emittenten gegen jeweils 1 US-Dollar ein, was den Preis des Stablecoins nach oben Richtung 1 US-Dollar treibt. Wenn der Preis des Stablecoins über 1 US-Dollar liegt, tauschen Arbitrageure beim Emittenten des Stablecoins US-Dollar gegen neue Stablecoins und verkaufen diese auf Kryptomärkten, was den Preis nach unten Richtung 1 US-Dollar treibt.  
						</p>
					`,
					"en": `

						<h3>Additional questions on the financial stability implications of stable coins</h3>

						<p style="text-align: justify:">
						The middle of May 2022 saw the collapse of the coin Terra amid general turbulences on markets for cryptocurrencies. Terra was conceived as a crypto asset that should always be worth 1 U.S. Dollar (USD), i.e. a so-called stable coin. During the unraveling of Terra, also Tether, the largest stable coin by market value (about 72.5 billion USD as of June 1, 2022), temporarily lost its peg of 1 USD, trading as low as 0.95 USD. Although Tether has recovered most of its losses since then, the difference between its target price of 1 USD and its market price on June 1, 2022 still was considerably higher than before the turbulences.  
						</p>

						<p style="text-align: justify:">
						Tether is a so-called asset-backed stable coin. The entities that issue asset-backed stable coins usually promise to hold financial assets worth 1 USD (or Euro etc.) for every stable coin issued. They also promise holders of their stable coins that they can always redeem these for USD at a ratio of 1:1. Under normal market conditions, arbitrage ensures that the stable coins have a market price of 1 USD. More specifically, if the price of a stable coin is below 1 USD, arbitrageurs buy it on the market and redeem it for 1 USD each with the stable coin’s issuer, thereby driving up the stable coin’s market price. If the price of a stable coin is above 1 USD, arbitrageurs exchange USD against stable coins at the stable coin’s issuer and sell these on the market, thereby driving down the stable coin’s market price.
						</p>
					`,
				}
			}

			//
			//
			names1 := []string{
				"qs4a_trad_system",
				"qs4b1_channels",
				"qs4b2_firesales",
				"qs4b3_indirect",
				"qs4b4_open",
			}
			mainLbls1 := []trl.S{
				{
					"de": `
						<p style="margin-left: -3.1rem">
						<b>4a.</b> &nbsp;
						Sollte ein großer Asset Backed Stablecoin zusammenbrechen, auf einer Skala zwischen 0 (kein Effekt) und 5 (extrem großer Effekt), wie groß schätzen Sie den direkten, negativen Effekt des Zusammenbruchs für das traditionelle Finanzsystem ein? Bitte berücksichtigen Sie nur Stablecoins, die mit Finanzwerten aus dem traditionellen Finanzsystem abgesichert werden, d.h. z.B. Staats- oder Unternehmensanleihen. 
						</p>
					`,
					"en": `
						<p style="margin-left: -3.1rem">
						<b>4a.</b> &nbsp;
						If a large asset-backed stable coin were to collapse, on a scale between 0 (no effect) and 5 (extremely large), how large do you think would the immediate negative effect be for the traditional financial system? Please consider only stable coins that are backed by traditional financial assets, e.g. government or corporate bonds.
						</p>
					`,
				},
				// q4b1
				{
					"de": `
						<p style="margin-left: -3.1rem; margin-bottom: 0.3rem;">
						<b>4b.</b> &nbsp;
						Wie wichtig wären die folgenden Kanäle für die Transmission von Shocks in das traditionelle Finanzsystem, 
						die vom Zusammenbruch eines großen Asset Backed Stablecoins ausgehen? 
						 
						Bitte berücksichtigen Sie nur Stablecoins, 
						die mit Finanzwerten aus dem traditionellen Finanzsystem abgesichert sind, 
						d.h. z.B. Staats- oder Unternehmensanleihen.
						</p>

						<p style="margin: 0 0 -0.67rem 0; ">
							<b>1.</b> &nbsp;	
							Fire Sales, die vom Emittenten des Stablecoins ausgehen: Der große Andrang von Haltern des Stablecoins, die diese gegen US-Dollar eintauschen wollen, zwingt den Emittenten des Stablecoins dazu, schnell Finanzwerte zu verkaufen, was zu Preisverfällen bei diesen Finanzwerten führt. 
						</p>
					`,
					"en": `
						<p style="margin-left: -3.1rem; margin-bottom: 0.3rem;">
						<b>4b.</b> &nbsp;
						How important would the following channels be for the transmission of shocks originating from the collapse of a large asset-backed stable coin to the traditional financial system? (0: not important, 5: extremely important). Please consider only stable coins that are backed by traditional financial assets, e.g. government or corporate bonds.
						</p>

						<p style="margin: 0 0 -0.67rem 0; ">
							<b>1.</b> &nbsp;	
							Fire-sales, originating from the stable coin's issuer: Large redemptions force the issuer to quickly sell traditional assets, causing dislocations in the respective assets' prices. 
							
						</p>
					`,
				},
				// q4b2
				{
					"de": `
						<p style="margin: 0 0 -0.67rem 0; ">
							<b>2.</b>	&nbsp;  Fire Sales, die von Haltern des Stablecoins ausgehen: Akteure des traditionellen Finanzsystems, die den Stablecoin halten, erleiden Verluste. Diese Verluste zwingen die Akteure des traditionellen Finanzsystems traditionelle Finanzwerte zu verkaufen, was zu Preisverfällen bei diesen Finanzwerten führt.
						</p>
					`,
					"en": `
						<p style="margin: 0 0 -0.67rem 0; ">
							<b>2.</b>	&nbsp;  Fire-sales, originating from holders of asset-backed stable coins: Actors from the traditional financial system, who own the collapsing stable coin, suffer losses. These losses then force them to sell non-crypto assets, causing dislocations on traditional asset markets.
						</p>
					`,
				},
				// q4b3
				{
					"de": `
						<p style="margin: 0 0 -0.67rem 0; ">
							<b>3.</b>	&nbsp;  Kreditbeziehungen: Akteure des traditionellen Finanzsystems haben Kredite an Unternehmen und Haushalte vergeben, die unmittelbar vom Zusammenbruch des Stablecoins betroffen sind. 
						</p>
					`,
					"en": `
						<p style="margin: 0 0 -0.67rem 0; ">
							<b>3.</b>	&nbsp;  Credit exposures: Actors from the traditional financial system have lending exposures to clients, who are directly exposed to the collapsing stable coin.
						</p>
					`,
				},
				// q4b4
				{
					"de": `
						<p style="margin: 0 0 -0.67rem 0; ">
							<b>4.</b>	&nbsp;  
							Offen
						</p>
					`,
					"en": `
						<p style="margin: 0 0 -0.67rem 0; ">
							<b>4.</b>	&nbsp;  
							Open
						</p>
					`,
				},
			}

			for idx, name := range names1 {

				names := []string{name}

				hdrLabels := importanceZeroToFive()
				if idx == 0 {
					hdrLabels = effectZeroToFive()
				}

				{

					if idx == len(names1)-1 {
						gr := page.AddGroup()
						gr.Cols = 1
						gr.BottomVSpacers = 1
						{

							inp := gr.AddInput()
							inp.Type = "text"
							inp.Name = "qs4b4_open_label"
							inp.Label = mainLbls1[idx]
							inp.ColSpan = 1
							inp.MaxChars = 20
							inp.ColSpanLabel = 2
							inp.ColSpanControl = 7
						}
						gr.Style = css.NewStylesResponsive(gr.Style)
						gr.Style.Desktop.StyleBox.Margin = "0 0 0 3.1rem"
					}

					gb := qst.NewGridBuilderRadios(
						columnTemplate7NFCL,
						hdrLabels,
						names,
						radioVals7,
						[]trl.S{{"de": ``, "en": ``}},
					)

					if idx == len(names1)-1 {
						// see above
					} else {
						gb.MainLabel = mainLbls1[idx]
					}
					gr := page.AddGrid(gb)

					gr.BottomVSpacers = 3
					if idx == 1 || idx == 2 {
						gr.BottomVSpacers = 2
					}

					gr.Style = css.NewStylesResponsive(gr.Style)
					gr.Style.Desktop.StyleBox.Margin = "0 0 0 3.1rem"

				}

			}

			{

				{
					gb := qst.NewGridBuilderRadios(
						columnTemplate6NFCL,
						agree6(),
						[]string{"qs5a"},
						radioVals6,
						[]trl.S{{"de": ``, "en": ``}},
					)

					gb.MainLabel = trl.S{
						"de": `
						<p style="margin-left: -3.1rem">
						<b>5a.</b> &nbsp;
						Stimmen Sie der folgenden Aussage zu oder nicht zu? Emittenten von Stablecoins sollten generell reguliert werden. 
						</p>
						`,
						"en": `
						<p style="margin-left: -3.1rem">
						<b>5a.</b> &nbsp;
						Do you agree or disagree with the following statement: Issuers of stable coins should be regulated in general.
						</p>
					`}

					gr := page.AddGrid(gb)

					gr.BottomVSpacers = 3

					gr.Style = css.NewStylesResponsive(gr.Style)
					gr.Style.Desktop.StyleBox.Margin = "0 0 0 3.1rem"

				}

			}

			//
			//
			//
			//
			mainLbl := trl.S{
				"de": `
					<p style="margin-bottom: -0.2rem">
					<b>5b.</b> &nbsp; Was glauben Sie, wie würden sich die folgenden Regulierungen auf das systemische Risiko von Stablecoins auswirken? 
					<br>
					&nbsp; &nbsp;  &nbsp; &nbsp;  (++: Stark positiv, +: positiv, 0: kein Effekt, -: negativ, --: stark negativ)

					</p>
				`,
				"en": `
					<p style="margin-bottom: -0.2rem">
					<b>5b.</b> &nbsp;
					How do you think would the following regulations affect the systemic risk of stable coins? 
					<br>
					&nbsp; &nbsp;  &nbsp; &nbsp;  (++: strongly positive, +: positive, 0: no effect, -: negative, --: strongly negative)

					</p>
				`,
			}
			names2 := []string{
				"qs5b1",
				"qs5b2",
				"qs5b3",
				"qs5b4",
			}
			colOneLbls := []trl.S{
				{
					"de": `
					Einführung einer Finanzaufsicht über die Emittenten von Stablecoins 
					`,
					"en": `Introduction of financial supervision of the entities that issue the stable coins`,
				},
				{
					"de": `
					Einschränkungen bei den Finanzwerten, welche die Emittenten von Stablecoins zur Unterlegung des Stablecoins halten dürfen, z.B. mit Blick auf Liquiditäts- und Kreditrisiko.
					`,
					"en": `Constraints on which assets can be held by the issuers of stable coins in order to back the stable coin, e.g. concerning liquidity and credit risk.`,
				},
				{
					"de": `
					Minimum Eigenkapitalvorgaben für die Emittenten von Stablecoins
					`,
					"en": `Minimum capital requirements for the issuers of stable coins`,
				},
				{
					"de": `
					Stablecoins vollständig verbieten
					`,
					"en": `Forbidding stable coins entirely`,
				},
			}

			{
				gb := qst.NewGridBuilderRadios(
					[]float32{
						9, 1, // wide col
						0, 1,
						0, 1,
						0, 1,
						0, 1,
						0.4, 1.3,
					},
					improvedDeterioratedPlusMinus6(),
					names2,
					radioVals6,
					colOneLbls,
				)

				gb.MainLabel = mainLbl
				gr := page.AddGrid(gb)

				gr.BottomVSpacers = 3

				// gr.Style = css.NewStylesResponsive(gr.Style)
				// gr.Style.Desktop.StyleBox.Margin = "0 0 0 3.1rem"

			}

		} // page

	}

	return nil

}
