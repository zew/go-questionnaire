package cppat

func init() {

	fourPermutationsOf6x3x3[0] = []preferences3x3T{
		{
			ID: 1, // Frage 1
			Ppls: [][]int{
				{3, 0, 2},
				{2, 3, 0},
				{0, 2, 3},
			},
		},
		{
			ID: 2, // Frage 2
			Ppls: [][]int{
				{3, 0, 2},
				{2, 2, 1},
				{0, 3, 2},
			},
		},
		{
			ID: 3, // Frage 3
			Ppls: [][]int{
				{3, 0, 2},
				{0, 5, 0},
				{2, 0, 3},
			},
		},
		{
			ID: 4, // Frage 4
			Ppls: [][]int{
				{3, 1, 1},
				{1, 4, 0},
				{1, 0, 4},
			},
		},
		{
			ID: 5, // Frage 5 - rows permutation 1
			Ppls: [][]int{
				{1, 4, 0}, // 2021-05-17 sandro; previously {3, 0, 2}
				{2, 0, 3}, // 2021-05-17 sandro; previously {0, 4, 1}
				{2, 1, 2}, // 2021-05-17 sandro; unchanged
			},
		},
		{
			ID: 6, // Frage 6
			Ppls: [][]int{
				{4, 0, 1},
				{0, 5, 0},
				{1, 0, 4},
			},
		},
	}

}
