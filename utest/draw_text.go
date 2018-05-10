// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-05-10 22:50:31 8FFC62                           [utest/draw_text.go]
// -----------------------------------------------------------------------------

package utest

import (
	"fmt"
	"testing"

	"github.com/balacode/one-file-pdf"
)

// Test_PDF_DrawText_ is the unit test for
// DrawText(s string) *PDF
func Test_PDF_DrawText_(t *testing.T) {
	fmt.Println("Test PDF.DrawText()")
	//
	func() {
		var doc = pdf.NewPDF("A4")
		doc.SetCompression(false).
			SetUnits("cm").
			SetColumnWidths(1, 4, 9).
			SetColor("#006B3C CadmiumGreen").
			SetFont("Helvetica-Bold", 10).
			SetX(5).
			SetY(5).
			DrawText("FIRST").
			DrawText("SECOND").
			DrawText("THIRD")

		var expect = `
		%PDF-1.4
		1 0 obj <</Type/Catalog/Pages 2 0 R>>
		endobj
		2 0 obj <</Type/Pages/Count 1/MediaBox[0 0 595 841]/Kids[3 0 R]>>
		endobj
		3 0 obj <</Type/Page/Parent 2 0 R/Contents 4 0 R
		/Resources <</Font <</FNT1 5 0 R>> >> >>
		endobj
		4 0 obj <</Length 143>> stream
		BT /FNT1 10 Tf ET
		0.000 0.420 0.235 rg
		0.000 0.420 0.235 RG
		BT 0 700 Td (FIRST) Tj ET
		BT 28 700 Td (SECOND) Tj ET
		BT 141 700 Td (THIRD) Tj ET
		endstream
		endobj
		5 0 obj <</Type/Font/Subtype/Type1/Name/FNT1
		/BaseFont/Helvetica-Bold
		/Encoding/StandardEncoding>>
		endobj
		xref
		0 6
		0000000000 65535 f
		0000000010 00000 n
		0000000056 00000 n
		0000000130 00000 n
		0000000228 00000 n
		0000000422 00000 n
		trailer
		<</Size 6/Root 1 0 R>>
		startxref
		528
		%%EOF
		`

		ComparePDF(t, doc.Bytes(), expect)
	}()

	func() {
		var doc = pdf.NewPDF("A4")
		doc.SetCompression(false).
			SetUnits("cm").
			SetFont("Ye-Olde-Scriptte", 10).
			SetXY(5, 5).
			SetHorizontalScaling(150).
			DrawText("Ye-Olde-Scriptte")

		var expect = `
		%PDF-1.4
		1 0 obj <</Type/Catalog/Pages 2 0 R>>
		endobj
		2 0 obj <</Type/Pages/Count 1/MediaBox[0 0 595 841]/Kids[3 0 R]>>
		endobj
		3 0 obj <</Type/Page/Parent 2 0 R/Contents 4 0 R
		/Resources <</Font <</FNT1 5 0 R>> >> >>
		endobj
		4 0 obj <</Length 113>> stream
		BT /FNT1 10 Tf ET
		BT 150 Tz ET
		0.000 0.000 0.000 rg
		0.000 0.000 0.000 RG
		BT 141 700 Td (Ye-Olde-Scriptte) Tj ET
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
		0000000392 00000 n
		trailer
		<</Size 6/Root 1 0 R>>
		startxref
		493
		%%EOF
		`

		ComparePDF(t, doc.Bytes(), expect)
		TEqual(t, len(doc.Errors()), 1)
		TEqual(t, doc.PullError(),
			fmt.Errorf(`Invalid font "Ye-Olde-Scriptte" @DrawText`))
	}()

} //                                                          Test_PDF_DrawText_

//end
