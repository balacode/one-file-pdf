// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-03-19 23:06:24 71D523                           [utest/util_func.go]
// -----------------------------------------------------------------------------

package utest

import "strings" // standard
import "testing" // standard

// pdfCompare compares generated result bytes to the expected PDF content:
// - convert result to a string
// - format both result and expected string using formatPDFLines()
// - compare result and expected lines
// - raise an error if there are diffs (report up to 5 differences)
func pdfCompare(t *testing.T, result []byte, expect string) {
	//
	var results = formatPDFLines(string(result))
	var expects = formatPDFLines(expect)
	var lenResults = len(results)
	var lenExpects = len(expects)
	var max = lenResults
	if max < lenExpects {
		max = lenExpects
	}
	var errCount = 0
	for i := 0; i < max; i++ {
		//
		// get the expected and the result line at i
		// if the slice is too short, leave it blank
		var expect, result string
		if i < lenExpects {
			expect = expects[i]
		}
		if i < lenResults {
			result = results[i]
		}
		if expect == result { // no problem, move along
			continue
		}
		// only report the first 5 mismatches
		errCount++
		if errCount > 5 {
			break
		}
		t.Errorf("MISMATCH ON LINE %d:\n"+
			"EXPECTED: %s\n"+
			"PRODUCED: %s\n"+
			"\n", i+1, expect, result)
	}
} //                                                                  pdfCompare

// formatPDFLines accepts an uncompressed PDF document as a string,
// and returns an array of trimmed, non-empty lines
func formatPDFLines(s string) []string {
	// change all newlines to "\n"
	s = strings.Replace(s, "\r\n", "\n", -1)
	s = strings.Replace(s, "\r", "\n", -1)
	//
	// change all other white-spaces to spaces
	for _, space := range "\a\b\f\t\v" {
		s = strings.Replace(s, string(space), " ", -1)
	}
	// remove all repeated spaces
	for strings.Contains(s, "  ") {
		s = strings.Replace(s, "  ", " ", -1)
	}
	// trim and copy non-blank lines to result
	// also, continue lines that end with '\'
	var ar = strings.Split(s, "\n")
	var ret = make([]string, 0, len(ar))
	var prev = ""
	for _, line := range ar {
		line = strings.Trim(line, " \a\b\f\n\r\t\v")
		if line == "" {
			continue
		}
		if prev != "" {
			line = prev + line
		}
		if strings.HasSuffix(line, "\\") {
			prev = strings.TrimRight(line, "\\")
			continue
		}
		prev = ""
		ret = append(ret, line)
	}
	return ret
} //                                                              formatPDFLines

//end
