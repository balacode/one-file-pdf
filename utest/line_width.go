// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-04-23 11:32:14 032D2C                          [utest/line_width.go]
// -----------------------------------------------------------------------------

package utest

import "fmt"     // standard
import "testing" // standard

import "github.com/balacode/one-file-pdf"

// LineWidth is the unit test for PDF.LineWidth()
func LineWidth(t *testing.T) {
	fmt.Println("utest.LineWidth")
	//
	// LineWidth of new PDF must be 1 point
	func() {
		var doc pdf.PDF
		TEqual(t, doc.LineWidth(), 1)
	}()
	func() {
		var doc = pdf.NewPDF("A4")
		TEqual(t, doc.LineWidth(), 1)
	}()
	//
	// SetLineWidth() has effect on the property?
	func() {
		var doc pdf.PDF
		doc.SetLineWidth(42)
		TEqual(t, doc.LineWidth(), 42)
	}()
	func() {
		var doc = pdf.NewPDF("A4")
		doc.SetLineWidth(7)
		TEqual(t, doc.LineWidth(), 7)
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
			DrawText("LineWidth Property")
		var expect = `
		%PDF-1.4
		%%EOF
        `
		pdfCompare(t, doc.Bytes(), expect, pdfStreamsInText)
		doc.SaveFile("~~test_line_width.pdf")
	}()

} //                                                                   LineWidth

//end
