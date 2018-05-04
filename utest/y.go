// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-05-04 23:51:59 9C2AEF                                   [utest/y.go]
// -----------------------------------------------------------------------------

package utest

import (
	"fmt"
	"testing"

	"github.com/balacode/one-file-pdf"
)

// Test_PDF_Y_ is the unit test for PDF.Y()
func Test_PDF_Y_(t *testing.T) {
	fmt.Println("Test PDF.Y()")
	//
	//TODO: fix Y(), then enable test cases
	if false {
		// Y of new PDF must be -1
		func() {
			var doc pdf.PDF
			TEqual(t, doc.Y(), -1) // TODO: fix: returns 1 instead of -1
		}()
		func() {
			var doc = pdf.NewPDF("A4")
			TEqual(t, doc.Y(), -1) // TODO: fix: returns 842.889764 instead of -1
		}()
	}
	// SetY() has effect on the property?
	func() {
		var doc pdf.PDF
		doc.SetY(220)
		TEqual(t, doc.Y(), 220)
	}()
	func() {
		var doc = pdf.NewPDF("A4")
		doc.SetY(110)
		TEqual(t, doc.Y(), 110)
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
			DrawText("Y Property")
		var expect = `
		%PDF-1.4
		%%EOF
        `
		ComparePDF(t, doc.Bytes(), expect, StreamsInText)
		doc.SaveFile("``test_x.pdf")
	}()

} //                                                                 Test_PDF_Y_

//end
