// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2019-04-03 10:10:37 AD0537              one-file-pdf/utest/util/[func.go]
// -----------------------------------------------------------------------------

package util

import (
	"bytes"
	"fmt"
	"strconv"
	str "strings"
	"testing"
)

// ComparePDF compares generated result bytes to the expected PDF content:
// - convert result to a string
// - format both result and expected string using formatLines()
// - compare result and expected lines
// - raise an error if there are diffs (report up to 5 differences)
func ComparePDF(t *testing.T, result []byte, expect string) {
	//
	const formatStreams = true
	var (
		results    = formatLines(string(result), formatStreams)
		expects    = formatLines(expect, !formatStreams)
		lenResults = len(results)
		lenExpects = len(expects)
		errCount   = 0
		mismatch   = false
		max        = lenResults
	)
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
				str.Join(results, "\n")+"\n"+
				"`\n")
	}
} //                                                                  ComparePDF

// FailIfHasErrors raises a test failure if the supplied PDF has errors
func FailIfHasErrors(t *testing.T, errors func() []error) {
	if len(errors()) == 0 {
		return
	}
	for i, err := range errors() {
		t.Errorf("ERROR %d: %s\n\n", i+1, err)
	}
	t.Fail()
} //                                                             FailIfHasErrors

// formatLines accepts an uncompressed PDF document as a string,
// and returns an array of trimmed, non-empty lines
func formatLines(s string, formatStreams bool) []string {
	//
	// format streams
	if formatStreams {
		s = pdfFormatStreams(s)
	}
	//
	// change all newlines to "\n"
	s = str.Replace(s, "\r\n", "\n", -1)
	s = str.Replace(s, "\r", "\n", -1)
	//
	// change all other white-spaces to spaces
	for _, space := range "\a\b\f\t\v" {
		s = str.Replace(s, string(space), " ", -1)
	}
	// remove all repeated spaces
	for str.Contains(s, "  ") {
		s = str.Replace(s, "  ", " ", -1)
	}
	// trim and copy non-blank lines to result
	// also, continue lines that end with '\'
	var (
		ar   = str.Split(s, "\n")
		ret  = make([]string, 0, len(ar))
		prev = ""
	)
	for _, line := range ar {
		line = str.Trim(line, " \a\b\f\n\r\t\v")
		if line == "" {
			continue
		}
		// join lines split with '\'
		if prev != "" {
			line = prev + line
		}
		if str.HasSuffix(line, "\\") {
			prev = str.TrimRight(line, "\\")
			continue
		}
		// append line to result
		prev = ""
		ret = append(ret, line)
	}
	return ret
} //                                                                 formatLines

// pdfFormatStreams formats content of all streams in s as hex strings
func pdfFormatStreams(s string) string {
	const (
		STREAM    = ">> stream"
		ENDSTREAM = "endstream"
		BPL       = 16 // bytes per line
	)
	buf := bytes.NewBuffer(make([]byte, 0, len(s)))
	for part, s := range str.Split(s, " obj ") {
		if part > 0 {
			buf.WriteString(" obj ")
		}
		// write the stream as-is if not compressed/image
		i := str.Index(s, STREAM)
		if i == -1 ||
			(!str.Contains(s[:i], "/FlateDecode") &&
				!str.Contains(s[:i], "/Image")) {
			buf.WriteString(s)
			continue
		}
		// write the part before stream's data without changing it
		i += len(STREAM)
		buf.WriteString(s[:i])
		s = s[i:]
		//
		// write the stream's data as hex numbers (each line with BPL columns)
		buf.WriteString("\n")
		n := str.Index(s, ENDSTREAM)
		if n == -1 {
			n = len(s)
		}
		c := 0
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
		// write the part after the stream's data
		s = s[n:]
		if len(s) > 0 {
			buf.WriteString(s)
		}
	}
	return buf.String()
} //                                                            pdfFormatStreams

//end
