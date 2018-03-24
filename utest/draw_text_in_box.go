// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-03-24 18:56:20 CB4C89                    [utest/draw_text_in_box.go]
// -----------------------------------------------------------------------------

package utest

import "fmt"     // standard
import "testing" // standard

import "github.com/balacode/one-file-pdf"

// DrawTextInBox is the unit test for
// DrawTextInBox(
//     x, y, width, height float64, align, text string) *PDF
func DrawTextInBox(t *testing.T) {
	fmt.Println("utest.DrawTextInBox")
	//
	const lorem = "Lorem ipsum dolor sit amet," +
		" consectetur adipiscing elit," +
		" sed do eiusmod tempor incididunt ut" +
		" labore et dolore magna aliqua." +
		" Ut enim ad minim veniam," +
		" quis nostrud exercitation ullamco laboris" +
		" nisi ut aliquip ex ea commodo consequat." +
		" Duis aute irure dolor in reprehenderit in voluptate velit" +
		" esse cillum dolore eu fugiat nulla pariatur." +
		" Excepteur sint occaecat cupidatat non proident," +
		" sunt in culpa qui officia deserunt mollit anim id est laborum."
	//
	var doc = pdf.NewPDF("A4")
	doc.SetCompression(false).
		SetUnits("cm").
		SetFont("Helvetica", 10).
		SetColor("Light Gray").
		FillBox(5, 5, 3, 15).
		SetColor("Black").
		DrawTextInBox(5, 5, 3, 15, "C", ""). // no effect
		DrawTextInBox(5, 5, 3, 15, "C", lorem)
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
	4 0 obj <</Length 1280>>stream
	 0.827 0.827 0.827 rg
	0.827 0.827 0.827 RG
	141.732 274.961 85.039 425.197 re b
	BT /FNT1 10 Tf ET
	 0.000 0.000 0.000 rg
	0.000 0.000 0.000 RG
	BT 153 624 Td (Lorem ipsum ) Tj ET
	BT 150 614 Td ( dolor sit amet, ) Tj ET
	BT 155 604 Td ( consectetur ) Tj ET
	BT 150 594 Td ( adipiscing elit, ) Tj ET
	BT 146 584 Td ( sed do eiusmod ) Tj ET
	BT 143 574 Td ( tempor incididunt ) Tj ET
	BT 156 564 Td ( ut labore et ) Tj ET
	BT 150 554 Td ( dolore magna ) Tj ET
	BT 148 544 Td ( aliqua. Ut enim ) Tj ET
	BT 142 534 Td ( ad minim veniam, ) Tj ET
	BT 154 524 Td ( quis nostrud ) Tj ET
	BT 155 514 Td ( exercitation ) Tj ET
	BT 148 504 Td ( ullamco laboris ) Tj ET
	BT 145 494 Td ( nisi ut aliquip ex ) Tj ET
	BT 152 484 Td ( ea commodo ) Tj ET
	BT 145 474 Td ( consequat. Duis ) Tj ET
	BT 142 464 Td ( aute irure dolor in ) Tj ET
	BT 146 454 Td ( reprehenderit in ) Tj ET
	BT 150 444 Td ( voluptate velit ) Tj ET
	BT 156 434 Td ( esse cillum ) Tj ET
	BT 147 424 Td ( dolore eu fugiat ) Tj ET
	BT 151 414 Td ( nulla pariatur. ) Tj ET
	BT 149 404 Td ( Excepteur sint ) Tj ET
	BT 161 394 Td ( occaecat ) Tj ET
	BT 151 384 Td ( cupidatat non ) Tj ET
	BT 145 374 Td ( proident, sunt in ) Tj ET
	BT 147 364 Td ( culpa qui officia ) Tj ET
	BT 148 354 Td ( deserunt mollit ) Tj ET
	BT 137 344 Td ( anim id est laborum.) Tj ET
	endstream
	5 0 obj<</Type/Font/Subtype/Type1/Name/F1/BaseFont/Helvetica\
	/Encoding/WinAnsiEncoding>>
	endobj
	xref
	0 6
	0000000000 65535 f
	0000000009 00000 n
	0000000053 00000 n
	0000000125 00000 n
	0000000217 00000 n
	0000001539 00000 n
	trailer
	<</Size 6/Root 1 0 R>>
	startxref
	1634
	%%EOF
	`
	pdfCompare(t, doc.Bytes(), expect, pdfStreamsInText)
} //                                                               DrawTextInBox

//TODO: add test for Courier font
//TDOO: Courier font metrics are not correct

//end
