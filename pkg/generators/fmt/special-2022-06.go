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
			"de": "Crypto",
			"en": "Crypto",
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
					"de": `todo`,
					"en": `

						<h3>Additional questions on the financial stability implications of stable coins</h3>

						<p style="text-align: justify:">
						The middle of May 2022 saw the collapse of the coin Terra amid general turbulences on markets for cryptocurrencies. Terra was conceived as a crypto asset that should always be worth exactly 1 USD, i.e. a so-called stable coin. During the unraveling of Terra, also Tether, the largest stable coin by market value (about 73 billion USD as of May 25, 2022), temporarily lost its peg of 1 USD, trading as low as 0.97 USD. Tether has recovered most of its losses since then, but still trades below 1 USD as of May 25, 2022.
						</p>

						<p style="text-align: justify:">
						Tether is a so-called asset-backed stable coin. The entities that issue asset-backed stable coins usually promise to hold financial assets worth 1 USD (the peg) for every stable coin issued. They also promise holders of their stable coins that they can always redeem these for USD at a ratio of 1:1. Under normal market conditions, arbitrage ensures that the stable coins have a market price of 1 USD. More specifically, if the price of a stable coin is below 1 USD, arbitrageurs buy it on the market and redeem it for 1 USD each with the stable coin’s issuer, thereby driving up the stable coin’s market price. If the price of a stable coin is above 1 USD, arbitrageurs exchange USD against stable coins at the stable coin’s issuer and sell these on the market, thereby driving down the stable coin’s market price.
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
			}
			mainLbls1 := []trl.S{
				{
					"de": `todo`,
					"en": `
						<p style="margin-left: -3.1rem">
						<b>4a.</b> &nbsp;
						If a large asset-backed stable coin were to collapse, on a scale between 0 (no effect) and 5 (extremely large), how large do you think would the immediate negative effect be for the traditional financial system? Please consider only stable coins that are backed by traditional financial assets, e.g. government or corporate bonds.
						</p>
					`,
				},
				{
					"de": `todo`,
					"en": `

						<p style="margin-left: -3.1rem; margin-bottom: 0.3rem;">
						<b>4b.</b> &nbsp;
						How important would the following channels be for the transmission of shocks originating from the collapse of a large asset-backed stable coin to the traditional financial system? (0: not important, 5: extremely important). Please consider only stable coins that are backed by traditional financial assets, e.g. government or corporate bonds.
						</p>

						<p style="margin: 0 0 -0.67rem 0; ">
							<b>1.</b> &nbsp;	Fire-sales, originating from the stable coin’s issuer: Large redemptions force the issuer to quickly sell traditional assets, causing dislocations in the respective assets’ prices. Holders of these assets from the traditional financial system suffer mark-to-market losses.
						</p>
					`,
				},
				{
					"de": `todo`,
					"en": `
						<p style="margin: 0 0 -0.67rem 0; ">
							<b>2.</b>	&nbsp;  Fire-sale, originating from holders of asset-backed stable coins: Actors from the traditional financial system, who own the collapsing stable coin, suffer losses. These losses then force them to sell non-crypto assets, causing dislocations on traditional asset markets.
						</p>
					`,
				},
				{
					"de": `todo`,
					"en": `
						<p style="margin: 0 0 -0.67rem 0; ">
							<b>3.</b>	&nbsp;  Indirect exposures: Actors from the traditional financial system have lending exposures to clients, who are directly exposed to the collapsing stable coin.
						</p>
					`,
				},
			}

			for idx, name := range names1 {

				names := []string{name}

				hdrLabels := zeroToFive()
				if idx == 4 {
					hdrLabels = zeroToFive()
				}

				{
					gb := qst.NewGridBuilderRadios(
						columnTemplate7NFCL,
						hdrLabels,
						names,
						radioVals7,
						[]trl.S{{"de": ``, "en": ``}},
					)

					gb.MainLabel = mainLbls1[idx]
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
						"de": `todo`,
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
				"de": `todo`,
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
					"de": `todo`,
					"en": `Supervision of the entities that issue the stable coins `,
				},
				{
					"de": `todo`,
					"en": `Constraints on which assets can be held by the issuers of stable coins in order to back the stable coin, e.g. concerning liquidity and credit risk.`,
				},
				{
					"de": `todo`,
					"en": `Minimum capital requirements for the issuers of stable coins`,
				},
				{
					"de": `todo`,
					"en": `Forbidding stable coins`,
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
