// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-04-29 23:42:24 1E9A9E             one-file-pdf/utest/[page_width.go]
// -----------------------------------------------------------------------------

package utest

import (
	"fmt"
	"testing"

	"github.com/balacode/one-file-pdf"
)

// Test_PDF_PageWidth_ tests PDF.PageWidth()
func Test_PDF_PageWidth_(t *testing.T) {
	fmt.Println("Test PDF.PageWidth()")

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

} //                                                         Test_PDF_PageWidth_

//end
