package qst

import "github.com/zew/go-questionnaire/pkg/trl"

func createRows(
	page *pageT,
	ac assetClass,
	inpNames []string,
	overTypes []string,
	lbls []trl.S,
	rangeCfgs []*rangeConf,
) {

	if len(ac.TrancheTypes) == 0 {
		return
	}

	for idx1, inpName := range inpNames {

		if overTypes[idx1] == "range-pct" {
			rangesRowLabelsTop(
				page,
				ac,
				inpName,
				lbls[idx1],
				*rangeCfgs[idx1],
			)
		}

		if overTypes[idx1] == "radios1-4" {
			radiosLabelsTop(
				page,
				ac,
				inpName,
				lbls[idx1],
				mCh2a,
			)
		}

		if overTypes[idx1] == "restricted-text-million" {
			restrTextRowLabelsTop(
				page,
				ac,
				inpName,
				lbls[idx1],
				rTSingleRowMill,
			)
		}

		if overTypes[idx1] == "restricted-text-pct" {
			restrTextRowLabelsTop(
				page,
				ac,
				inpName,
				lbls[idx1],
				rTSingleRowPercent,
			)
		}

	}

}
