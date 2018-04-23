// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-04-23 11:32:14 303080                           [utest/util_func.go]
// -----------------------------------------------------------------------------

package utest

import "bytes"   // standard
import "fmt"     // standard
import "strconv" // standard
import "strings" // standard
import "testing" // standard

import "github.com/balacode/one-file-pdf"

type pdfStreamFmt int

const (
	pdfStreamsInText = iota
	pdfStreamsInHex  = iota
)

// pdfCompare compares generated result bytes to the expected PDF content:
// - convert result to a string
// - format both result and expected string using pdfFormatLines()
// - compare result and expected lines
// - raise an error if there are diffs (report up to 5 differences)
func pdfCompare(t *testing.T, result []byte, expect string, sfmt pdfStreamFmt) {
	//
	var results = pdfFormatLines(string(result), sfmt)
	var expects = pdfFormatLines(expect, pdfStreamsInText)
	var lenResults = len(results)
	var lenExpects = len(expects)
	var errCount = 0
	var mismatch = false
	var max = lenResults
	if max < lenExpects {
		max = lenExpects
	}
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
		mismatch = true
		errCount++
		if errCount > 5 {
			break
		}
		t.Errorf("%s",
			"\n"+
				"/*\n"+
				"LOCATION: "+TCaller()+":\n"+
				"MISMATCH: L"+strconv.Itoa(i+1)+":\n"+
				"EXPECTED: "+expect+"\n"+
				"PRODUCED: "+result+"\n"+
				"*/\n")
	}
	if mismatch {
		t.Errorf("%s",
			"\n"+
				"// RETURNED-PDF:\n"+
				"// "+TCaller()+"\n"+
				"`\n"+
				strings.Join(results, "\n")+"\n"+
				"`\n")
	}
} //                                                                  pdfCompare

// pdfFailIfErrors raises a test failure if the supplied PDF has errors
func pdfFailIfErrors(t *testing.T, doc *pdf.PDF) {
	if len(doc.Errors()) == 0 {
		return
	}
	for i, err := range doc.Errors() {
		t.Errorf("ERROR %d: %s\n\n", i+1, err)
	}
	t.Fail()
} //                                                             pdfFailIfErrors

// pdfFormatLines accepts an uncompressed PDF document as a string,
// and returns an array of trimmed, non-empty lines
func pdfFormatLines(s string, sfmt pdfStreamFmt) []string {
	if sfmt == pdfStreamsInHex {
		s = pdfFormatStreamsInHex(s)
	}
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
		// join lines split with '\'
		if prev != "" {
			line = prev + line
		}
		if strings.HasSuffix(line, "\\") {
			prev = strings.TrimRight(line, "\\")
			continue
		}
		// append line to result
		prev = ""
		ret = append(ret, line)
	}
	return ret
} //                                                              pdfFormatLines

// pdfFormatStreamsInHex formats content of all streams in s as hex strings
func pdfFormatStreamsInHex(s string) string {
	const (
		STREAM    = ">> stream"
		ENDSTREAM = "endstream"
		BPL       = 16 // bytes per line
	)
	var buf = bytes.NewBuffer(make([]byte, 0, len(s)))
	for {
		// exit if there are no more streams
		var i = strings.Index(s, STREAM)
		if i == -1 {
			break
		}
		// write the part before stream's data without changing it
		i += len(STREAM)
		buf.WriteString(s[:i])
		s = s[i:]
		//
		// write the stream data as hex bytes (in BPL columns of bytes)
		buf.WriteString("\n")
		var n = strings.Index(s, ENDSTREAM)
		if n == -1 {
			n = len(s)
		}
		var c = 0
		for _, b := range []byte(s[:n]) {
			buf.WriteString(fmt.Sprintf(" %02X", b))
			c++
			if c >= BPL {
				buf.WriteString("\n")
				c = 0
			}
		}
		buf.WriteString("\n")
		//
		// advance to next position
		s = s[n:]
	}
	if len(s) > 0 {
		buf.WriteString(s)
	}
	return buf.String()
} //                                                       pdfFormatStreamsInHex

//end
