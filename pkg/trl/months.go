package trl

import (
	"fmt"
	"time"
)

func EnglishMonths() []string {

	loc := time.Now().Location()

	mn := make([]string, 0, 12) // month names

	for i := 1; i < 13; i++ {
		t := time.Date(2022, time.Month(i), 2, 0, 0, 0, 0, loc) // 2022 is not relevant, any year
		m := fmt.Sprintln(t.Month())
		mn = append(mn, m)
	}

	return mn
}

var monthsTranslations = []S{
	{
		"en": "January",
		"de": "Januar",
	},
	{
		"en": "February",
		"de": "Februar",
	},
	{
		"en": "March",
		"de": "März",
	},
	{
		"en": "April",
		"de": "April",
	},
	{
		"en": "May",
		"de": "Mai",
	},
	{
		"en": "June",
		"de": "Juni",
	},
	{
		"en": "July",
		"de": "Juli",
	},
	{
		"en": "August",
		"de": "August",
	},
	{
		"en": "September",
		"de": "September",
	},
	{
		"en": "October",
		"de": "Oktober",
	},
	{
		"en": "November",
		"de": "November",
	},
	{
		"en": "December",
		"de": "Dezember",
	},
}

// MonthsShift could use date.Add - but it would be wasteful
func MonthsShift(i, shift int) int {
	if shift >= 0 {
		//
	} else {
		for (i + shift) < 1 {
			// this loop could be obviated
			i += 12
		}
	}

	shifted := i + shift
	if shifted > 12 {
		shifted = shifted % 12
	}
	return shifted

}

// MonthByInt maps 1 to January, 12 to December
func MonthByInt(i int) S {
	if i < 1 || i > 12 {
		return S{
			"en": fmt.Sprintf("error_unknown_month_idx__%v", i),
			"de": fmt.Sprintf("error_unknown_month_idx__%v", i),
		}
	}
	return monthsTranslations[i-1]
}
