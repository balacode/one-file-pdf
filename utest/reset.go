// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-03-26 13:13:57 FD0666                               [utest/reset.go]
// -----------------------------------------------------------------------------

package utest

import "fmt"     // standard
import "testing" // standard

import "github.com/balacode/one-file-pdf"

// Reset tests PDF.Reset()
func Reset(t *testing.T) {
	fmt.Println("utest.Reset")
	//
	// prepare a test PDF before calling Reset()
	var doc = pdf.NewPDF("A4")
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
		var expect = `
		%PDF-1.4
		1 0 obj<</Type/Catalog/Pages 2 0 R>>
		endobj
		2 0 obj<</Type/Pages/Count 1/MediaBox[0 0 595 841]/Kids[3 0 R]>>
		endobj
		3 0 obj<</Type/Page/Parent 2 0 R/Contents 4 0 R\
		/Resources<</Font <</FNT1 5 0 R>>>>>>
		endobj
		4 0 obj <</Length 143>>stream
		BT /FNT1 10 Tf ET
		 0.000 0.420 0.235 rg
		0.000 0.420 0.235 RG
		BT 0 700 Td (FIRST) Tj ET
		BT 28 700 Td (SECOND) Tj ET
		BT 141 700 Td (THIRD) Tj ET
		endstream
		5 0 obj<</Type/Font/Subtype/Type1/Name/F1/BaseFont/Helvetica-Bold\
		/Encoding/WinAnsiEncoding>>
		endobj
		xref
		0 6
		0000000000 65535 f
		0000000009 00000 n
		0000000053 00000 n
		0000000125 00000 n
		0000000217 00000 n
		0000000401 00000 n
		trailer
		<</Size 6/Root 1 0 R>>
		startxref
		501
		%%EOF
	`
		var result = doc.Bytes()
		pdfCompare(t, result, expect, pdfStreamsInText)
	}
	{
		doc.Reset()
		//
		// after calling Reset(), the PDF should just be a blank page:
		var expect = `
		%PDF-1.4
		1 0 obj<</Type/Catalog/Pages 2 0 R>>
		endobj
		2 0 obj<</Type/Pages/Count 1/MediaBox[0 0 595 841]/Kids[3 0 R]>>
		endobj
		3 0 obj<</Type/Page/Parent 2 0 R/Contents 4 0 R>>
		endobj
		4 0 obj <</Length 0>>stream
		endstream
		xref
		0 5
		0000000000 65535 f
		0000000009 00000 n
		0000000053 00000 n
		0000000125 00000 n
		0000000182 00000 n
		trailer
		<</Size 5/Root 1 0 R>>
		startxref
		221
		%%EOF
		`
		var result = doc.SetCompression(false).Bytes()
		pdfCompare(t, result, expect, pdfStreamsInText)
	}
	// TODO: add more test cases, test each property's state
} //                                                                       Reset

//end
