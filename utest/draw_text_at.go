// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-03-22 03:10:56 869A01                        [utest/draw_text_at.go]
// -----------------------------------------------------------------------------

package utest

import "fmt"     // standard
import "testing" // standard

import "github.com/balacode/one-file-pdf"

// DrawTextAt is the unit test for
// DrawTextAt(x, y float64, text string) *PDF
func DrawTextAt(t *testing.T) {
	fmt.Println("utest.DrawTextAt")
	var pdf = pdf.NewPDF("A4")
	pdf.
		SetCompression(false).
		SetUnits("cm").
		SetColor("#36454F Charcoal").
		SetFont("Helvetica-Bold", 20).
		DrawTextAt(5, 5, ""). // no effect
		DrawTextAt(5, 5, "(5,5)").
		DrawTextAt(10, 10, ""). // no effect
		DrawTextAt(10, 10, "(10,10)").
		DrawTextAt(15, 15, ""). // no effect
		DrawTextAt(15, 15, "(15,15)").
		SetColor("#E03C31 CGRed").
		FillBox(5, 5, 0.1, 0.1).
		FillBox(10, 10, 0.1, 0.1).
		FillBox(15, 15, 0.1, 0.1)
	//
	var expect = `
	%PDF-1.4
	1 0 obj<</Type/Catalog/Pages 2 0 R>>
	endobj
	2 0 obj<</Type/Pages/Count 1/MediaBox[0 0 595 841]/Kids[3 0 R]>>
	endobj
	3 0 obj<</Type/Page/Parent 2 0 R/Contents 4 0 R\
	/Resources<</Font <</FNT1 5 0 R>>>>>>
	endobj
	4 0 obj <</Length 297>>stream
	BT /FNT1 20 Tf ET
	0.212 0.271 0.310 rg
	0.212 0.271 0.310 RG
	BT 141 700 Td (\(5,5\)) Tj ET
	BT 283 558 Td (\(10,10\)) Tj ET
	BT 425 416 Td (\(15,15\)) Tj ET
	0.878 0.235 0.192 rg
	0.878 0.235 0.192 RG
	141.732 697.323 2.835 2.835 re b
	283.465 555.591 2.835 2.835 re b
	425.197 413.858 2.835 2.835 re b
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
	0000000555 00000 n
	trailer
	<</Size 6/Root 1 0 R>>
	startxref
	655
	%%EOF
	`
	pdfCompare(t, pdf.Bytes(), expect, pdfStreamsInText)
} //                                                                  DrawTextAt

//end
