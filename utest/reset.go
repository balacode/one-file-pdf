// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2019-04-03 10:27:56 B458CF                  one-file-pdf/utest/[reset.go]
// -----------------------------------------------------------------------------

package utest

import (
	"fmt"
	"testing"

	"github.com/balacode/one-file-pdf"
)

// Test_PDF_Reset_ tests PDF.Reset()
func Test_PDF_Reset_(t *testing.T) {
	fmt.Println("Test PDF.Reset()")
	//
	// prepare a test PDF before calling Reset()
	doc := pdf.NewPDF("A4")
	{
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
		const want = `
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
		got := doc.Bytes()
		ComparePDF(t, got, want)
	}
	{
		doc.Reset()
		//
		// after calling Reset(), the PDF should just be a blank page:
		const want = `
		%PDF-1.4
		1 0 obj <</Type/Catalog/Pages 2 0 R>>
		endobj
		2 0 obj <</Type/Pages/Count 1/MediaBox[0 0 595 841]/Kids[3 0 R]>>
		endobj
		3 0 obj <</Type/Page/Parent 2 0 R/Contents 4 0 R>>
		endobj
		4 0 obj <</Length 0>> stream
		endstream
		endobj
		xref
		0 5
		0000000000 65535 f
		0000000010 00000 n
		0000000056 00000 n
		0000000130 00000 n
		0000000189 00000 n
		trailer
		<</Size 5/Root 1 0 R>>
		startxref
		238
		%%EOF
		`
		got := doc.SetCompression(false).Bytes()
		ComparePDF(t, got, want)
	}
	// TODO: add more test cases, test each property's state
} //                                                             Test_PDF_Reset_

//end
