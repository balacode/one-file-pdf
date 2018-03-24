// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-03-24 22:56:46 522E42                         [utest/page_height.go]
// -----------------------------------------------------------------------------

package utest

import "fmt"     // standard
import "testing" // standard

import "github.com/balacode/one-file-pdf"

func PageHeight(t *testing.T) {
	fmt.Println("utest.PageHeight")

	func() {
		var doc pdf.PDF // uninitialized PDF
		TEqual(t, doc.PageHeight(), 0.0)
	}()

	// 2.83464566929134 points per mm
	// 1 inch / 25.4mm per inch * 72 points per inch

	func() {
		var doc = pdf.NewPDF("A4")              // initialized PDF
		TEqual(t, doc.PageHeight(), 841.889764) // points
		//
		// A4 = 210mm width x 297mm height = 841.8897637795276 points
	}()

	func() {
		var doc = pdf.NewPDF("LETTER")          // initialized PDF
		TEqual(t, doc.PageHeight(), 790.866142) // points
		//
		// LETTER = 216mm width x 279mm height = 790.8661417322835 points
	}()

} //                                                                  PageHeight

//end
