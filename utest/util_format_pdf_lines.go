// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-03-12 21:58:10 CC7777               [utest/util_format_pdf_lines.go]
// -----------------------------------------------------------------------------

package utest

import "strings" // standard

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
