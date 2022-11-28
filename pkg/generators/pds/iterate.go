package pds

import (
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

func iterate(
	page *qst.WrappedPageT,
	inpNames []string,
	overTypes []string,
	lbls []trl.S,
	rangeCfgs []*rangeConf,
) {

	for idx1, inpName := range inpNames {

		if overTypes[idx1] == "restricted-text-million" {
			restrTextRowLabelsTop(
				page,
				inpName,
				lbls[idx1],
				rTSingleRowMill,
			)
		}

		if overTypes[idx1] == "restricted-text-pct" {
			restrTextRowLabelsTop(
				page,
				inpName,
				lbls[idx1],
				rTSingleRowPercent,
			)
		}

		if overTypes[idx1] == "range-pct" {
			slidersPctRowLabelsTop(
				page,
				inpName,
				lbls[idx1],
				*rangeCfgs[idx1],
			)
		}

		if overTypes[idx1] == "radios1-4" {
			chapter3(
				page,
				inpName,
				lbls[idx1],
				mCh2a,
			)
		}

	}

}
