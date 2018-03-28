// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-03-28 03:14:14 C3366F                                   [utest/x.go]
// -----------------------------------------------------------------------------

package utest

import "fmt"     // standard
import "testing" // standard

import "github.com/balacode/one-file-pdf"

// X is the unit test for PDF.X()
func X(t *testing.T) {
	fmt.Println("utest.X")
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

} //                                                                           X

//end
