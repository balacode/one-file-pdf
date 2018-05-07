// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-05-07 19:15:31 49F4B6                                   [utest/x.go]
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
	// SetX() sets the property?
	func() {
		var doc pdf.PDF
		doc.SetX(123)
		TEqual(t, doc.X(), 123)
	}()
	func() {
		var doc = pdf.NewPDF("A4")
		doc.SetX(456)
		TEqual(t, doc.X(), 456)
	}()
	// -------------------------------------------------------------------------
	// Test PDF generation
	func() {
		var doc = pdf.NewPDF("A4")
		doc.SetCompression(false).
			SetUnits("cm").
			SetXY(10, 1).
			SetFont("Times-Bold", 20).
			DrawText("X=10 Y=1")
		var expect = `
		%PDF-1.4
		1 0 obj <</Type/Catalog/Pages 2 0 R>>
		endobj
		2 0 obj <</Type/Pages/Count 1/MediaBox[0 0 595 841]/Kids[3 0 R]>>
		endobj
		3 0 obj <</Type/Page/Parent 2 0 R/Contents 4 0 R
		/Resources <</Font <</FNT1 5 0 R>> >> >>
		endobj
		4 0 obj <</Length 92>> stream
		BT /FNT1 20 Tf ET
		0.000 0.000 0.000 rg
		0.000 0.000 0.000 RG
		BT 283 813 Td (X=10 Y=1) Tj ET
		endstream
		endobj
		5 0 obj <</Type/Font/Subtype/Type1/Name/FNT1
		/BaseFont/Times-Bold
		/Encoding/StandardEncoding>>
		endobj
		xref
		0 6
		0000000000 65535 f
		0000000010 00000 n
		0000000056 00000 n
		0000000130 00000 n
		0000000228 00000 n
		0000000370 00000 n
		trailer
		<</Size 6/Root 1 0 R>>
		startxref
		472
		%%EOF
		`
		ComparePDF(t, doc.Bytes(), expect, StreamsInText)
		// doc.SaveFile("~~Test_PDF_X_.pdf")
	}()
} //                                                                 Test_PDF_X_

//end
