// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2019-04-28 20:57:28 F688A9                 one-file-pdf/[private_test.go]
// -----------------------------------------------------------------------------

package pdf

/*
This file contains unit tests for internal methods/functions.

To generate a test coverage report use:
	go test -coverprofile cover.out
	go tool cover -html=cover.out
*/

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/balacode/one-file-pdf/utest/util"
)

// go test --run Test_getPapreSize_
func Test_getPapreSize_(t *testing.T) {
	//
	// subtest: tests a specific paper size against width and height in points
	subtest := func(paperSize, permuted string, w, h float64, err error) {
		var doc PDF
		doc.SetUnits("mm")
		got, gotErr := doc.getPaperSize(permuted)
		if gotErr != err {
			t.Errorf("'error' mismatch: expected: %v returned %v", err, gotErr)
			t.Fail()
		}
		if got.name != paperSize {
			mismatch(t, permuted+" 'name'", paperSize, got.name)
		}
		if floatStr(got.widthPt) != floatStr(w) {
			mismatch(t, permuted+" 'widthPt'", w, got.widthPt)
		}
		if floatStr(got.heightPt) != floatStr(h) {
			mismatch(t, permuted+" 'heightPt'", h, got.heightPt)
		}
	}
	// test: tests the given paper size in portrait and landscape orientations
	// w and h are the paper width and height in mm
	// - permutes the paper size by including spaces
	// - converts units from mm to points
	test := func(paperSize string, w, h float64, err error) {
		const PTperMM = 2.83464566929134
		spaces := []string{"", " ", "  ", "   ", "\r", "\n", "\t"}
		for _, orient := range []string{"", "-l", "-L"} {
			permuted := util.PermuteStrings(
				spaces,
				[]string{
					strings.ToLower(paperSize),
					strings.ToUpper(paperSize),
				},
				spaces,
				[]string{orient},
				spaces,
			)
			for _, s := range permuted {
				size := paperSize + strings.ToUpper(orient)
				if orient == "-l" || orient == "-L" {
					subtest(size, s, h*PTperMM, w*PTperMM, err)
				} else {
					subtest(size, s, w*PTperMM, h*PTperMM, err)
				}
			}
		}
	}
	test("A4", 210, 297, nil)
	test("A0", 841, 1189, nil)
	test("A1", 594, 841, nil)
	test("A2", 420, 594, nil)
	test("A3", 297, 420, nil)
	test("A4", 210, 297, nil)
	test("A5", 148, 210, nil)
	test("A6", 105, 148, nil)
	test("A7", 74, 105, nil)
	test("A8", 52, 74, nil)
	test("A9", 37, 52, nil)
	test("A10", 26, 37, nil)
	test("B0", 1000, 1414, nil)
	test("B1", 707, 1000, nil)
	test("B2", 500, 707, nil)
	test("B3", 353, 500, nil)
	test("B4", 250, 353, nil)
	test("B5", 176, 250, nil)
	test("B6", 125, 176, nil)
	test("B7", 88, 125, nil)
	test("B8", 62, 88, nil)
	test("B9", 44, 62, nil)
	test("B10", 31, 44, nil)
	test("C0", 917, 1297, nil)
	test("C1", 648, 917, nil)
	test("C2", 458, 648, nil)
	test("C3", 324, 458, nil)
	test("C4", 229, 324, nil)
	test("C5", 162, 229, nil)
	test("C6", 114, 162, nil)
	test("C7", 81, 114, nil)
	test("C8", 57, 81, nil)
	test("C9", 40, 57, nil)
	test("C10", 28, 40, nil)
	test("LEDGER", 432, 279, nil)
	test("LEGAL", 216, 356, nil)
	test("LETTER", 216, 279, nil)
	test("TABLOID", 279, 432, nil)
} //                                                          Test_getPapreSize_

// -----------------------------------------------------------------------------
// # Helper Functions

// floatStr returns a float64 as a string, with val rounded to 3 decimals
func floatStr(val float64) string {
	return fmt.Sprintf("%0.3f", val)
} //                                                                    floatStr

// mismatch formats and raises a test error
func mismatch(t *testing.T, tag string, want, got interface{}) {
	ws := fmt.Sprintf("%v", want)
	gs := fmt.Sprintf("%v", got)
	t.Errorf("%s mismatch: expected: %s got: %s\n"+"%s",
		tag, ws, gs, getStack())
} //                                                                    mismatch

// getStack returns a list of line numbers and function names on the call stack
func getStack() string {
	buf := make([]byte, 8192)
	runtime.Stack(buf, true)
	var ar []string
	for _, s := range strings.Split(string(buf), "\n") {
		if strings.Contains(s, "\t") && !strings.Contains(s, "/testing.go") {
			ar = append(ar, "<- "+filepath.Base(s))
		}
	}
	return strings.Join(ar, "\n")
} //                                                                    getStack

//end
