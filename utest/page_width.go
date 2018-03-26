// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-03-26 12:23:48 087D54                          [utest/page_width.go]
// -----------------------------------------------------------------------------

package utest

import "fmt"     // standard
import "testing" // standard

import "github.com/balacode/one-file-pdf"

// PageWidth tests PDF.PageWidth()
func PageWidth(t *testing.T) {
	fmt.Println("utest.PageWidth")

	func() {
		var doc pdf.PDF // uninitialized PDF
		TEqual(t, doc.PageWidth(), 0.0)
	}()

	// 2.83464566929134 points per mm
	// 1 inch / 25.4mm per inch * 72 points per inch

	func() {
		var doc = pdf.NewPDF("A4")             // initialized PDF
		TEqual(t, doc.PageWidth(), 595.275591) // points
		//
		// A4 = 210mm width x 297mm height = 595.2755905511811 points
	}()

	func() {
		var doc = pdf.NewPDF("LETTER")         // initialized PDF
		TEqual(t, doc.PageWidth(), 612.283465) // points
		//
		// LETTER = 216mm width x 279mm height = 612.2834645669291 points
	}()

} //                                                                   PageWidth

//end
