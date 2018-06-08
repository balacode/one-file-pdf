// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-06-08 21:53:17 0E6BBA                one-file-pdf/utest/[set_x_y.go]
// -----------------------------------------------------------------------------

package utest

import (
	"fmt"
	"testing"

	"github.com/balacode/one-file-pdf"
)

// Test_PDF_SetXY_ is the unit test for PDF.SetXY()
func Test_PDF_SetXY_(t *testing.T) {
	fmt.Println("Test PDF.SetXY()")
	//
	// SetXY() sets X and Y properties properly?
	func() {
		var doc pdf.PDF
		doc.SetXY(123, 456)
		TEqual(t, doc.X(), 123)
		TEqual(t, doc.Y(), 456)
	}()
	func() {
		var doc = pdf.NewPDF("A4")
		doc.SetXY(123, 456)
		TEqual(t, doc.X(), 123)
		TEqual(t, doc.Y(), 456)
	}()
	// -------------------------------------------------------------------------
	// Test PDF generation
	func() {
		var doc = pdf.NewPDF("A4")
		doc.SetCompression(false).SetUnits("cm").SetFont("Helvetica", 10).
			SetXY(1, 3).DrawText("X=1cm Y=3cm").
			SetXY(3, 1).DrawText("X=3cm Y=1cm").
			SetXY(10, 5).DrawText("X=10cm Y=5cm").
			SetXY(5, 10).DrawText("X=5cm Y=10cm")

		const expect = `
		%PDF-1.4
		1 0 obj <</Type/Catalog/Pages 2 0 R>>
		endobj
		2 0 obj <</Type/Pages/Count 1/MediaBox[0 0 595 841]/Kids[3 0 R]>>
		endobj
		3 0 obj <</Type/Page/Parent 2 0 R/Contents 4 0 R
		/Resources <</Font <</FNT1 5 0 R>> >> >>
		endobj
		4 0 obj <</Length 197>> stream
		BT /FNT1 10 Tf ET
		0.000 0.000 0.000 rg
		0.000 0.000 0.000 RG
		BT 28 756 Td (X=1cm Y=3cm) Tj ET
		BT 85 813 Td (X=3cm Y=1cm) Tj ET
		BT 283 700 Td (X=10cm Y=5cm) Tj ET
		BT 141 558 Td (X=5cm Y=10cm) Tj ET
		endstream
		endobj
		5 0 obj <</Type/Font/Subtype/Type1/Name/FNT1
		/BaseFont/Helvetica
		/Encoding/StandardEncoding>>
		endobj
		xref
		0 6
		0000000000 65535 f
		0000000010 00000 n
		0000000056 00000 n
		0000000130 00000 n
		0000000228 00000 n
		0000000476 00000 n
		trailer
		<</Size 6/Root 1 0 R>>
		startxref
		577
		%%EOF
		`

		ComparePDF(t, doc.Bytes(), expect)
		doc.SaveFile("~~Test_PDF_SetXY_.pdf")
	}()
} //                                                             Test_PDF_SetXY_

//end
