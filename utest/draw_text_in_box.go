// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-05-04 23:51:59 F87BFC                    [utest/draw_text_in_box.go]
// -----------------------------------------------------------------------------

package utest

import (
	"fmt"
	"testing"

	"github.com/balacode/one-file-pdf"
)

// Test_PDF_DrawTextInBox_ is the unit test for
// DrawTextInBox(
//     x, y, width, height float64, align, text string) *PDF
func Test_PDF_DrawTextInBox_(t *testing.T) {
	fmt.Println("Test PDF.DrawTextInBox()")
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

	func() {
		var doc = pdf.NewPDF("A4")
		doc.SetCompression(false).
			SetUnits("cm").
			SetFont("Helvetica", 10).
			SetColor("Light Gray").
			FillBox(5, 5, 3, 15).
			SetColor("Black").
			DrawTextInBox(5, 5, 3, 15, "C", ""). // no effect
			DrawTextInBox(5, 5, 3, 15, "C", lorem)

		var expect = `
		%PDF-1.4
		1 0 obj <</Type/Catalog/Pages 2 0 R>>
		endobj
		2 0 obj <</Type/Pages/Count 1/MediaBox[0 0 595 841]/Kids[3 0 R]>>
		endobj
		3 0 obj <</Type/Page/Parent 2 0 R/Contents 4 0 R
		/Resources <</Font <</FNT1 5 0 R>> >> >>
		endobj
		4 0 obj <</Length 1275>> stream
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
		0000001555 00000 n
		trailer
		<</Size 6/Root 1 0 R>>
		startxref
		1656
		%%EOF
		`

		ComparePDF(t, doc.Bytes(), expect, StreamsInText)
	}()

	func() {
		var doc = pdf.NewPDF("A4")
		doc.SetCompression(false).
			SetUnits("cm").
			SetFont("Courier", 10).
			SetColor("Light Gray").
			FillBox(5, 5, 3, 15).
			SetColor("Black").
			DrawTextInBox(5, 5, 3, 15, "C", ""). // no effect
			DrawTextInBox(5, 5, 3, 15, "C", lorem)

		var expect = `
		%PDF-1.4
		1 0 obj <</Type/Catalog/Pages 2 0 R>>
		endobj
		2 0 obj <</Type/Pages/Count 1/MediaBox[0 0 595 841]/Kids[3 0 R]>>
		endobj
		3 0 obj <</Type/Page/Parent 2 0 R/Contents 4 0 R
		/Resources <</Font <</FNT1 5 0 R>> >> >>
		endobj
		4 0 obj <</Length 1505>> stream
		0.827 0.827 0.827 rg
		0.827 0.827 0.827 RG
		141.732 274.961 85.039 425.197 re b
		BT /FNT1 10 Tf ET
		0.000 0.000 0.000 rg
		0.000 0.000 0.000 RG
		BT 148 679 Td (Lorem ipsum ) Tj ET
		BT 154 669 Td (dolor sit ) Tj ET
		BT 166 659 Td (amet, ) Tj ET
		BT 148 649 Td (consectetur ) Tj ET
		BT 151 639 Td (adipiscing ) Tj ET
		BT 145 629 Td (elit, sed do ) Tj ET
		BT 160 619 Td (eiusmod ) Tj ET
		BT 163 609 Td (tempor ) Tj ET
		BT 142 599 Td (incididunt ut ) Tj ET
		BT 154 589 Td (labore et ) Tj ET
		BT 145 579 Td (dolore magna ) Tj ET
		BT 151 569 Td (aliqua. Ut ) Tj ET
		BT 142 559 Td (enim ad minim ) Tj ET
		BT 145 549 Td (veniam, quis ) Tj ET
		BT 160 539 Td (nostrud ) Tj ET
		BT 145 529 Td (exercitation ) Tj ET
		BT 160 519 Td (ullamco ) Tj ET
		BT 145 509 Td (laboris nisi ) Tj ET
		BT 142 499 Td (ut aliquip ex ) Tj ET
		BT 151 489 Td (ea commodo ) Tj ET
		BT 151 479 Td (consequat. ) Tj ET
		BT 154 469 Td (Duis aute ) Tj ET
		BT 148 459 Td (irure dolor ) Tj ET
		BT 175 449 Td (in ) Tj ET
		BT 142 439 Td (reprehenderit ) Tj ET
		BT 145 429 Td (in voluptate ) Tj ET
		BT 151 419 Td (velit esse ) Tj ET
		BT 142 409 Td (cillum dolore ) Tj ET
		BT 154 399 Td (eu fugiat ) Tj ET
		BT 166 389 Td (nulla ) Tj ET
		BT 154 379 Td (pariatur. ) Tj ET
		BT 154 369 Td (Excepteur ) Tj ET
		BT 142 359 Td (sint occaecat ) Tj ET
		BT 142 349 Td (cupidatat non ) Tj ET
		BT 154 339 Td (proident, ) Tj ET
		BT 142 329 Td (sunt in culpa ) Tj ET
		BT 148 319 Td (qui officia ) Tj ET
		BT 157 309 Td (deserunt ) Tj ET
		BT 148 299 Td (mollit anim ) Tj ET
		BT 139 289 Td (id est laborum.) Tj ET
		endstream
		endobj
		5 0 obj <</Type/Font/Subtype/Type1/Name/FNT1
		/BaseFont/Courier
		/Encoding/StandardEncoding>>
		endobj
		xref
		0 6
		0000000000 65535 f
		0000000010 00000 n
		0000000056 00000 n
		0000000130 00000 n
		0000000228 00000 n
		0000001785 00000 n
		trailer
		<</Size 6/Root 1 0 R>>
		startxref
		1884
		%%EOF
		`

		ComparePDF(t, doc.Bytes(), expect, StreamsInText)
	}()

} //                                                     Test_PDF_DrawTextInBox_

//end
