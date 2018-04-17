// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-04-17 23:15:55 111D46                    [utest/draw_text_in_box.go]
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
	4 0 obj <</Length 1275>>stream
	0.827 0.827 0.827 rg
	0.827 0.827 0.827 RG
	141.732 274.961 85.039 425.197 re b
	BT /FNT1 10 Tf ET
	0.000 0.000 0.000 rg
	0.000 0.000 0.000 RG
	BT 153 629 Td (Lorem ipsum ) Tj ET
	BT 151 619 Td (dolor sit amet, ) Tj ET
	BT 157 609 Td (consectetur ) Tj ET
	BT 142 599 Td (adipiscing elit, sed ) Tj ET
	BT 157 589 Td (do eiusmod ) Tj ET
	BT 144 579 Td (tempor incididunt ) Tj ET
	BT 142 569 Td (ut labore et dolore ) Tj ET
	BT 145 559 Td (magna aliqua. Ut ) Tj ET
	BT 150 549 Td (enim ad minim ) Tj ET
	BT 154 539 Td (veniam, quis ) Tj ET
	BT 166 529 Td (nostrud ) Tj ET
	BT 157 519 Td (exercitation ) Tj ET
	BT 149 509 Td (ullamco laboris ) Tj ET
	BT 147 499 Td (nisi ut aliquip ex ) Tj ET
	BT 153 489 Td (ea commodo ) Tj ET
	BT 147 479 Td (consequat. Duis ) Tj ET
	BT 143 469 Td (aute irure dolor in ) Tj ET
	BT 147 459 Td (reprehenderit in ) Tj ET
	BT 152 449 Td (voluptate velit ) Tj ET
	BT 142 439 Td (esse cillum dolore ) Tj ET
	BT 151 429 Td (eu fugiat nulla ) Tj ET
	BT 164 419 Td (pariatur. ) Tj ET
	BT 151 409 Td (Excepteur sint ) Tj ET
	BT 162 399 Td (occaecat ) Tj ET
	BT 152 389 Td (cupidatat non ) Tj ET
	BT 147 379 Td (proident, sunt in ) Tj ET
	BT 148 369 Td (culpa qui officia ) Tj ET
	BT 150 359 Td (deserunt mollit ) Tj ET
	BT 158 349 Td (anim id est ) Tj ET
	BT 164 339 Td (laborum.) Tj ET
	endstream
	5 0 obj<</Type/Font/Subtype/Type1/Name/FNT1/BaseFont/Helvetica\
	/Encoding/StandardEncoding>>
	endobj
	xref
	0 6
	0000000000 65535 f
	0000000009 00000 n
	0000000053 00000 n
	0000000125 00000 n
	0000000217 00000 n
	0000001534 00000 n
	trailer
	<</Size 6/Root 1 0 R>>
	startxref
	1632
	%%EOF
	`
	pdfCompare(t, doc.Bytes(), expect, pdfStreamsInText)
} //                                                               DrawTextInBox

//TODO: add test for Courier font
//TDOO: Courier font metrics are not correct

//end
