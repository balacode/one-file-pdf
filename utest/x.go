// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-04-29 23:42:24 6FE7A2                                   [utest/x.go]
// -----------------------------------------------------------------------------

package utest

import (
	"fmt"
	"testing"

	"github.com/balacode/one-file-pdf"
)

// Test_PDF_X_ is the unit test for PDF.X()
func Test_PDF_X_(t *testing.T) {
	fmt.Println("Test PDF.X()")
	//
	// X of new PDF must be -1
	func() {
		var doc pdf.PDF
		TEqual(t, doc.X(), -1)
	}()
	func() {
		var doc = pdf.NewPDF("A4")
		TEqual(t, doc.X(), -1)
	}()
	// SetX() has effect on the property?
	func() {
		var doc pdf.PDF
		doc.SetX(220)
		TEqual(t, doc.X(), 220)
	}()
	func() {
		var doc = pdf.NewPDF("A4")
		doc.SetX(110)
		TEqual(t, doc.X(), 110)
	}()

	// -------------------------------------------------------------------------
	// Test PDF generation

	return // TODO: implement test case

	func() {
		var doc = pdf.NewPDF("A4")
		doc.SetCompression(false).
			SetUnits("cm").
			SetFont("Times-Bold", 20).
			SetXY(1, 1).
			DrawText("X Property")
		var expect = `
		%PDF-1.4
		%%EOF
        `
		pdfCompare(t, doc.Bytes(), expect, pdfStreamsInText)
		doc.SaveFile("``test_x.pdf")
	}()

} //                                                                 Test_PDF_X_

//end
