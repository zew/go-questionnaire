package qst

// config multiple choice
type configMC struct {
	KeyLabels          string // key to a map of labels
	Cols               float32
	InpColspan         float32
	LabelBottom        bool
	DontKnow           bool
	GroupBottomSpacers int

	GroupLeftIndent string

	XDisplacements []string
}

var (
	mCh2 = configMC{
		KeyLabels:          "teamsize",
		Cols:               4,
		InpColspan:         1,
		LabelBottom:        true,
		DontKnow:           false,
		GroupBottomSpacers: 3,
	}

	mCh2a = configMC{
		KeyLabels:          "covenants-per-credit",
		Cols:               4,
		InpColspan:         1,
		LabelBottom:        false,
		DontKnow:           false,
		GroupBottomSpacers: 3,
		GroupLeftIndent:    outline2Indent,

		XDisplacements: []string{
			"1.6rem",
			"0.62rem",
			"0.62rem",
			"1.6rem",
		},
	}

	mCh3 = configMC{
		KeyLabels:   "relevance1-5",
		Cols:        5,
		InpColspan:  1,
		LabelBottom: false,
		DontKnow:    false,
	}

	mCh4Prev = configMC{
		KeyLabels:       "improveDecline1-5-prev",
		Cols:            5,
		InpColspan:      1,
		LabelBottom:     false,
		DontKnow:        false,
		GroupLeftIndent: outline2Indent,
		XDisplacements: []string{
			"1.6rem",
			"0.79rem",
			"",
			"0.79rem",
			"1.6rem",
		},
	}
	mCh4Next = configMC{
		KeyLabels:       "improveDecline1-5-next",
		Cols:            5,
		InpColspan:      1,
		LabelBottom:     false,
		DontKnow:        false,
		GroupLeftIndent: outline2Indent,
		XDisplacements: []string{
			"1.6rem",
			"0.79rem",
			"",
			"0.79rem",
			"1.6rem",
		},
	}
	mCh5 = configMC{
		KeyLabels:   "closing-time-weeks",
		Cols:        5,
		InpColspan:  1,
		LabelBottom: false,
		DontKnow:    false,

		// not yet
		// GroupLeftIndent: outline2Indent,

		// XDisplacements: []string{
		// 	"1.46rem",
		// 	"1.27rem",
		// 	"0.64rem",
		// 	"",
		// 	"0.64rem",
		// 	"1.27rem",
		// 	"1.46rem",
		// },
		XDisplacements: []string{
			"1.6rem",
			"0.79rem",
			"",
			"0.79rem",
			"1.6rem",
		},
	}
)
