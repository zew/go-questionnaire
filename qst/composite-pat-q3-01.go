package qst

import (
	"fmt"
	"log"
)

func getQ3Labels(version0to15 int, aOrB string) []string {

	//  0 => 0
	//  1 => 0
	//  2 => 1
	//  3 => 1
	//  4 => 2
	//  5 => 2
	//  6 => 3
	//   ...
	// 14 => 7
	// 15 => 7
	version0to7 := version0to15 / 2

	//                                    2*0 + 1            2*0 + 2
	//                                    2*1 + 1            2*1 + 2
	//                                    2*3 + 1            2*3 + 2
	key := fmt.Sprintf("version-%02v-%02v", 2*version0to7+1, 2*version0to7+2)
	log.Printf("key is %v", key)

	if aOrB == "a" {
		return dataQ3[key][0]
	}
	if aOrB == "b" {
		return dataQ3[key][1]
	}

	log.Panicf("pat-getQ3 - second argument must be 'a' or 'b' - was %v", aOrB)
	return []string{"invalid", "invalid", "invalid"}

}

var dataQ3 = map[string][][]string{

	"version-01-02": {
		// a
		{
			"0 € sofort und 18 € in 6 Monaten",
			"3 € sofort und 12 € in 6 Monaten",
			"6 € sofort und  3 € in 6 Monaten",
		},
		// b
		{
			"0 € in 1 Monat und 24 € in 7 Monaten",
			"4 € in 1 Monat und 16 € in 7 Monaten",
			"8 € in 1 Monat und  4 € in 7 Monaten",
		},
	},
	"version-03-04": {
		{
			"0 € in 1 Monat und 18 € in 7 Monaten",
			"3 € in 1 Monat und 12 € in 7 Monaten",
			"6 € in 1 Monat und  3 € in 7 Monaten",
		},
		{
			"0 € sofort und 24 € in 6 Monaten",
			"4 € sofort und 16 € in 6 Monaten",
			"8 € sofort und  4 € in 6 Monaten",
		},
	},
	"version-05-06": {
		{
			"0 € sofort und 24 € in 6 Monaten",
			"4 € sofort und 16 € in 6 Monaten",
			"8 € sofort und  4 € in 6 Monaten",
		},
		{
			"0 € in 1 Monat und 18 € in 7 Monaten",
			"3 € in 1 Monat und 12 € in 7 Monaten",
			"6 € in 1 Monat und  3 € in 7 Monaten",
		},
	},
	"version-07-08": {
		{
			"0 € in 1 Monat und 24 € in 7 Monaten",
			"4 € in 1 Monat und 16 € in 7 Monaten",
			"8 € in 1 Monat und  4 € in 7 Monaten",
		},
		{
			"0 € sofort und 18 € in 6 Monaten",
			"3 € sofort und 12 € in 6 Monaten",
			"6 € sofort und  3 € in 6 Monaten",
		},
	},
	"version-09-10": {
		{
			"0 € sofort und 16 € in 6 Monaten",
			"2 € sofort und 10 € in 6 Monaten",
			"4 € sofort und  2 € in 6 Monaten",
		},
		{
			"0 € in 1 Monat und 24 € in 7 Monaten",
			"3 € in 1 Monat und 15 € in 7 Monaten",
			"6 € in 1 Monat und  3 € in 7 Monaten",
		},
	},
	"version-11-12": {
		{
			"0 € in 1 Monat und 16 € in 7 Monaten",
			"2 € in 1 Monat und 10 € in 7 Monaten",
			"4 € in 1 Monat und  2 € in 7 Monaten",
		},
		{
			"0 € sofort und 24 € in 6 Monaten",
			"3 € sofort und 15 € in 6 Monaten",
			"6 € sofort und  3 € in 6 Monaten",
		},
	},
	"version-13-14": {
		{
			"0 € sofort und 24 € in 6 Monaten",
			"3 € sofort und 15 € in 6 Monaten",
			"6 € sofort und  3 € in 6 Monaten",
		},
		{
			"0 € in 1 Monat und 16 € in 7 Monaten",
			"2 € in 1 Monat und 10 € in 7 Monaten",
			"4 € in 1 Monat und  2 € in 7 Monaten",
		},
	},
	"version-15-16": {
		{
			"0 € in 1 Monat und 24 € in 7 Monaten",
			"3 € in 1 Monat und 15 € in 7 Monaten",
			"6 € in 1 Monat und  3 € in 7 Monaten",
		},
		{
			"0 € sofort und 16 € in 6 Monaten",
			"2 € sofort und 10 € in 6 Monaten",
			"4 € sofort und  2 € in 6 Monaten",
		},
	},
}
