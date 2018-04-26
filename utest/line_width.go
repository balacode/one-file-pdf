// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-04-26 22:42:44 FF6555                          [utest/line_width.go]
// -----------------------------------------------------------------------------

package utest

import "fmt"     // standard
import "testing" // standard

import "github.com/balacode/one-file-pdf"

// Test_PDF_LineWidth_ is the unit test for PDF.LineWidth()
func Test_PDF_LineWidth_(t *testing.T) {
	fmt.Println("Test PDF.LineWidth()")
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

} //                                                         Test_PDF_LineWidth_

//end
