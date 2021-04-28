package qst

import (
	"fmt"
	"log"
)

func getQ2Labels(version0to15 int, aOrB string) []string {

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
	// log.Printf("key is %v", key)

	if aOrB == "a" {
		return dataQ2[key][0]
	}
	if aOrB == "b" {
		return dataQ2[key][1]
	}

	log.Panicf("pat-getQ2 - second argument must be 'a' or 'b' - was %v", aOrB)
	return []string{"invalid", "invalid", "invalid"}

}

var dataQ2 = map[string][][]string{

	"version-01-02": {
		// q2a
		{
			"<b> 0&nbsp;€ sofort</b> und<br><b>16&nbsp;€</b> in <b>6&nbsp;Monaten</b>",
			"<b> 2&nbsp;€ sofort</b> und<br><b>10&nbsp;€</b> in <b>6&nbsp;Monaten</b>",
			"<b> 4&nbsp;€ sofort</b> und<br> <b>2&nbsp;€</b> in <b>6&nbsp;Monaten</b>",
		},
		// q2b
		{
			"<b> 0&nbsp;€</b> in <b>1&nbsp;Monat</b> und<br><b>24&nbsp;€</b> in <b>7&nbsp;Monaten</b>",
			"<b> 3&nbsp;€</b> in <b>1&nbsp;Monat</b> und<br><b>15&nbsp;€</b> in <b>7&nbsp;Monaten</b>",
			"<b> 6&nbsp;€</b> in <b>1&nbsp;Monat</b> und<br> <b>3&nbsp;€</b> in <b>7&nbsp;Monaten</b>",
		},
	},
	"version-03-04": {
		{
			"<b> 0&nbsp;€</b> in <b>1&nbsp;Monat</b> und<br><b>16&nbsp;€</b> in <b>7&nbsp;Monaten</b>",
			"<b> 2&nbsp;€</b> in <b>1&nbsp;Monat</b> und<br><b>10&nbsp;€</b> in <b>7&nbsp;Monaten</b>",
			"<b> 4&nbsp;€</b> in <b>1&nbsp;Monat</b> und<br> <b>2&nbsp;€</b> in <b>7&nbsp;Monaten</b>",
		},
		{
			"<b> 0&nbsp;€ sofort</b> und<br><b>24&nbsp;€</b> in <b>6&nbsp;Monaten</b>",
			"<b> 3&nbsp;€ sofort</b> und<br><b>15&nbsp;€</b> in <b>6&nbsp;Monaten</b>",
			"<b> 6&nbsp;€ sofort</b> und<br> <b>3&nbsp;€</b> in <b>6&nbsp;Monaten</b>",
		},
	},
	"version-05-06": {
		{
			"<b> 0&nbsp;€ sofort</b> und<br><b>24&nbsp;€</b> in <b>6&nbsp;Monaten</b>",
			"<b> 3&nbsp;€ sofort</b> und<br><b>15&nbsp;€</b> in <b>6&nbsp;Monaten</b>",
			"<b> 6&nbsp;€ sofort</b> und<br> <b>3&nbsp;€</b> in <b>6&nbsp;Monaten</b>",
		},
		{
			"<b> 0&nbsp;€</b> in <b>1&nbsp;Monat</b> und<br><b>16&nbsp;€</b> in <b>7&nbsp;Monaten</b>",
			"<b> 2&nbsp;€</b> in <b>1&nbsp;Monat</b> und<br><b>10&nbsp;€</b> in <b>7&nbsp;Monaten</b>",
			"<b> 4&nbsp;€</b> in <b>1&nbsp;Monat</b> und<br> <b>2&nbsp;€</b> in <b>7&nbsp;Monaten</b>",
		},
	},
	"version-07-08": {
		{
			"<b> 0&nbsp;€</b> in <b>1&nbsp;Monat</b> und<br><b>24&nbsp;€</b> in <b>7&nbsp;Monaten</b>",
			"<b> 3&nbsp;€</b> in <b>1&nbsp;Monat</b> und<br><b>15&nbsp;€</b> in <b>7&nbsp;Monaten</b>",
			"<b> 6&nbsp;€</b> in <b>1&nbsp;Monat</b> und<br> <b>3&nbsp;€</b> in <b>7&nbsp;Monaten</b>",
		},
		{
			"<b> 0&nbsp;€ sofort</b> und<br><b>16&nbsp;€</b> in <b>6&nbsp;Monaten</b>",
			"<b> 2&nbsp;€ sofort</b> und<br><b>10&nbsp;€</b> in <b>6&nbsp;Monaten</b>",
			"<b> 4&nbsp;€ sofort</b> und<br> <b>2&nbsp;€</b> in <b>6&nbsp;Monaten</b>",
		},
	},
	"version-09-10": {
		{
			"<b> 0&nbsp;€ sofort</b> und<br><b>18&nbsp;€</b> in <b>6&nbsp;Monaten</b>",
			"<b> 3&nbsp;€ sofort</b> und<br><b>12&nbsp;€</b> in <b>6&nbsp;Monaten</b>",
			"<b> 6&nbsp;€ sofort</b> und<br> <b>3&nbsp;€</b> in <b>6&nbsp;Monaten</b>",
		},
		{
			"<b> 0&nbsp;€</b> in <b>1&nbsp;Monat</b> und<br><b>24&nbsp;€</b> in <b>7&nbsp;Monaten</b>",
			"<b> 4&nbsp;€</b> in <b>1&nbsp;Monat</b> und<br><b>16&nbsp;€</b> in <b>7&nbsp;Monaten</b>",
			"<b> 8&nbsp;€</b> in <b>1&nbsp;Monat</b> und<br> <b>4&nbsp;€</b> in <b>7&nbsp;Monaten</b>",
		},
	},
	"version-11-12": {
		{
			"<b> 0&nbsp;€</b> in <b>1&nbsp;Monat</b> und<br><b>18&nbsp;€</b> in <b>7&nbsp;Monaten</b>",
			"<b> 3&nbsp;€</b> in <b>1&nbsp;Monat</b> und<br><b>12&nbsp;€</b> in <b>7&nbsp;Monaten</b>",
			"<b> 6&nbsp;€</b> in <b>1&nbsp;Monat</b> und<br> <b>3&nbsp;€</b> in <b>7&nbsp;Monaten</b>",
		},
		{
			"<b> 0&nbsp;€ sofort</b> und<br><b>24&nbsp;€</b> in <b>6&nbsp;Monaten</b>",
			"<b> 4&nbsp;€ sofort</b> und<br><b>16&nbsp;€</b> in <b>6&nbsp;Monaten</b>",
			"<b> 8&nbsp;€ sofort</b> und<br> <b>4&nbsp;€</b> in <b>6&nbsp;Monaten</b>",
		},
	},
	"version-13-14": {
		{
			"<b> 0&nbsp;€ sofort</b> und<br><b>24&nbsp;€</b> in <b>6&nbsp;Monaten</b>",
			"<b> 4&nbsp;€ sofort</b> und<br><b>16&nbsp;€</b> in <b>6&nbsp;Monaten</b>",
			"<b> 8&nbsp;€ sofort</b> und<br> <b>4&nbsp;€</b> in <b>6&nbsp;Monaten</b>",
		},
		{
			"<b> 0&nbsp;€</b> in <b>1&nbsp;Monat</b> und<br><b>18&nbsp;€</b> in <b>7&nbsp;Monaten</b>",
			"<b> 3&nbsp;€</b> in <b>1&nbsp;Monat</b> und<br><b>12&nbsp;€</b> in <b>7&nbsp;Monaten</b>",
			"<b> 6&nbsp;€</b> in <b>1&nbsp;Monat</b> und<br> <b>3&nbsp;€</b> in <b>7&nbsp;Monaten</b>",
		},
	},
	"version-15-16": {
		{
			"<b> 0&nbsp;€</b> in <b>1&nbsp;Monat</b> und<br><b>24&nbsp;€</b> in <b>7&nbsp;Monaten</b>",
			"<b> 4&nbsp;€</b> in <b>1&nbsp;Monat</b> und<br><b>16&nbsp;€</b> in <b>7&nbsp;Monaten</b>",
			"<b> 8&nbsp;€</b> in <b>1&nbsp;Monat</b> und<br> <b>4&nbsp;€</b> in <b>7&nbsp;Monaten</b>",
		},
		{
			"<b> 0&nbsp;€ sofort</b> und<br><b>18&nbsp;€</b> in <b>6&nbsp;Monaten</b>",
			"<b> 3&nbsp;€ sofort</b> und<br><b>12&nbsp;€</b> in <b>6&nbsp;Monaten</b>",
			"<b> 6&nbsp;€ sofort</b> und<br> <b>3&nbsp;€</b> in <b>6&nbsp;Monaten</b>",
		},
	},
}
