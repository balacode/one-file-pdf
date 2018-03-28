// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-03-28 03:14:14 188617                                   [utest/y.go]
// -----------------------------------------------------------------------------

package utest

import "fmt"     // standard
import "testing" // standard

import "github.com/balacode/one-file-pdf"

// Y is the unit test for PDF.Y()
func Y(t *testing.T) {
	fmt.Println("utest.Y")
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
		pdfCompare(t, doc.Bytes(), expect, pdfStreamsInText)
		doc.SaveFile("``test_x.pdf")
	}()

} //                                                                           Y

//end
