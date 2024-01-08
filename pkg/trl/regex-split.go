package trl

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

/*
regexSplit splits s and returns the *variable* pieces of glue and the spandrels in between as equal parts.
The all parts concatenated must again be equal to the input of s.
It's an enlargement of Split(s, sep) where sep is *variable*,
and the resulting slice contains 'sep' and 'in between' alternatingly.

If the regexp produces overlapping separations, the func returns s (no change).
Thus regexSplit yields a *complete* intervall nesting (https://en.wikipedia.org/wiki/Nested_intervals).
This is of interest, because we want a lossless restoration of the input of s,
though some parts could be modified.

regexSplit uses re.FindAllStringIndex(s, -1) which splits s by matches to regex re.
It returns the boundaries of the splitters in an array of integers [][]int.
Example:

	<ul style='a:b' class="c2"/>
		<li>Here<br>are</li>
		<li>some spandrels</li>
	</ul>

should be split by any HTML tag
using    <[^>]+>

Test it using <https://regexr.com>

re.FindAllStringIndex() now returns

	[][]int{
		{ 0, 28},   // <ul.../>
		{32, 35},   // <li>
		{40, 43},   // <br>
		{48, 53},   // </li>
		{ ...  },   // <li>
		{ ...  },   // </li>
		{ ...  },   // </ul>
	}

We now want a slice of strings containing the splitters and the spandrels.
We need to know, whether the first slice element is a splitter or a spandrel.
Then we can modify either the splitters or the spandrels specifically,
and then recombine them to an enhanced s.
*/
func regexSplit(re *regexp.Regexp, s string) [][]int {
	return re.FindAllStringIndex(s, -1)
}

// indexesToStringSlice takes the split string indexes from a call of splitPaths
// and restores s
func indexesToStringSlice(s string, ps [][]int) []string {

	if len(ps) == 0 {
		return []string{s}
	}

	w := &strings.Builder{}
	ret := make([]string, 0, 2*len(ps))

	first := ps[0][0]
	if first > 0 {
		fmt.Fprint(w, s[0:first])
		ret = append(ret, s[0:first])
	}

	for i := 0; i < len(ps); i++ {
		o0 := ps[i]                   // occurrence
		fmt.Fprint(w, s[o0[0]:o0[1]]) // splitter
		ret = append(ret, s[o0[0]:o0[1]])
		if i < len(ps)-1 { // not last
			o1 := ps[i+1]
			fmt.Fprint(w, s[o0[1]:o1[0]])
			ret = append(ret, s[o0[1]:o1[0]])
		}
	}

	last := ps[len(ps)-1][1]
	if last < len(s)-1 {
		fmt.Fprint(w, s[last:]) // up till the end
		ret = append(ret, s[last:])
	}

	if s != w.String() {
		log.Fatalf("indexesToStringSlice() extracted \n\t%v  from \n\t%v", w.String(), s)
	}

	return ret
}

/*
https://regexr.com/
https://pkg.go.dev/regexp#Regexp.FindAllStringIndex

spaces   some-path/to\directory    whitespace
([\s]*   ([\S]*   [\\/]   [\S]*)   [\s]*)

condensed

	([\S]*[\\/][\S]*)
*/
var pathSplitter = regexp.MustCompile(`[\S]*[\\/][\S]*`)
var tagSplitter = regexp.MustCompile(`<[^>]+>`)
