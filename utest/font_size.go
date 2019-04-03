// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2019-04-03 10:03:38 02D294              one-file-pdf/utest/[font_size.go]
// -----------------------------------------------------------------------------

package utest

import (
	"fmt"
	"testing"

	"github.com/balacode/one-file-pdf"
)

// Test_PDF_FontSize_ is the unit test for PDF.FontSize() and SetFontSize()
func Test_PDF_FontSize_(t *testing.T) {

	// -------------------------------------------------------------------------
	// (ob *PDF) FontSize() string
	//
	fmt.Println("Test PDF.FontSize()")

	func() {
		var doc pdf.PDF // uninitialized PDF
		TEqual(t, doc.FontSize(), 10)
	}()

	func() {
		doc := pdf.NewPDF("A4") // initialized PDF
		TEqual(t, doc.FontSize(), 10)
	}()

	// -------------------------------------------------------------------------
	// (ob *PDF) SetFontSize(name string) *PDF
	//
	fmt.Println("Test PDF.SetFontSize()")

	func() {
		var doc pdf.PDF // uninitialized PDF
		TEqual(t, doc.SetFontSize(15).FontSize(), 15)
	}()

	func() {
		doc := pdf.NewPDF("A4") // initialized PDF
		TEqual(t, doc.SetFontSize(20).FontSize(), 20)
	}()

	// -------------------------------------------------------------------------
	// Test PDF generation

	func() {
		doc := pdf.NewPDF("A4")
		doc.SetCompression(false).SetUnits("cm")
		pt10 := doc.ToUnits(10) // 10pt
		const w = 10
		doc.SetFont("Times-Bold", 20).SetXY(1, 1).DrawText("Font Sizes")
		for i, size := range []float64{5, 6, 7, 8, 9, 10, 15, 20, 25, 30} {
			y := 2 + float64(i)*3*pt10
			doc.SetLineWidth(0.1).
				SetXY(1, y+0.5).SetColor("Gray").
				DrawBox(1, y+0*pt10, w, pt10).
				DrawBox(1, y+1*pt10, w, pt10).
				DrawBox(1, y+2*pt10, w, pt10).
				SetXY(1, y+0.5).SetColor("Black").SetFont("Helvetica", size).
				DrawBox(1, y, w, 3*pt10).
				DrawTextInBox(1, y, w, 3*pt10, "TC",
					fmt.Sprintf("Helvetica %1.0f", size))
		}

		const expect = `
		%PDF-1.4
		1 0 obj <</Type/Catalog/Pages 2 0 R>>
		endobj
		2 0 obj <</Type/Pages/Count 1/MediaBox[0 0 595 841]/Kids[3 0 R]>>
		endobj
		3 0 obj <</Type/Page/Parent 2 0 R/Contents 4 0 R
		/Resources <</Font <<
		/FNT1 5 0 R
		/FNT2 6 0 R>> >> >>
		endobj
		4 0 obj <</Length 2440>> stream
		BT /FNT1 20 Tf ET
		0.000 0.000 0.000 rg
		0.000 0.000 0.000 RG
		BT 28 813 Td (Font Sizes) Tj ET
		0.745 0.745 0.745 RG
		0.100 w
		28.346 775.197 283.465 10.000 re S
		28.346 765.197 283.465 10.000 re S
		28.346 755.197 283.465 10.000 re S
		0.000 0.000 0.000 RG
		28.346 755.197 283.465 30.000 re S
		BT /FNT2 5 Tf ET
		BT 157 780 Td (Helvetica 5) Tj ET
		0.745 0.745 0.745 RG
		28.346 745.197 283.465 10.000 re S
		28.346 735.197 283.465 10.000 re S
		28.346 725.197 283.465 10.000 re S
		0.000 0.000 0.000 RG
		28.346 725.197 283.465 30.000 re S
		BT /FNT2 6 Tf ET
		BT 155 749 Td (Helvetica 6) Tj ET
		0.745 0.745 0.745 RG
		28.346 715.197 283.465 10.000 re S
		28.346 705.197 283.465 10.000 re S
		28.346 695.197 283.465 10.000 re S
		0.000 0.000 0.000 RG
		28.346 695.197 283.465 30.000 re S
		BT /FNT2 7 Tf ET
		BT 152 718 Td (Helvetica 7) Tj ET
		0.745 0.745 0.745 RG
		28.346 685.197 283.465 10.000 re S
		28.346 675.197 283.465 10.000 re S
		28.346 665.197 283.465 10.000 re S
		0.000 0.000 0.000 RG
		28.346 665.197 283.465 30.000 re S
		BT /FNT2 8 Tf ET
		BT 150 687 Td (Helvetica 8) Tj ET
		0.745 0.745 0.745 RG
		28.346 655.197 283.465 10.000 re S
		28.346 645.197 283.465 10.000 re S
		28.346 635.197 283.465 10.000 re S
		0.000 0.000 0.000 RG
		28.346 635.197 283.465 30.000 re S
		BT /FNT2 9 Tf ET
		BT 147 656 Td (Helvetica 9) Tj ET
		0.745 0.745 0.745 RG
		28.346 625.197 283.465 10.000 re S
		28.346 615.197 283.465 10.000 re S
		28.346 605.197 283.465 10.000 re S
		0.000 0.000 0.000 RG
		28.346 605.197 283.465 30.000 re S
		BT /FNT2 10 Tf ET
		BT 142 625 Td (Helvetica 10) Tj ET
		0.745 0.745 0.745 RG
		28.346 595.197 283.465 10.000 re S
		28.346 585.197 283.465 10.000 re S
		28.346 575.197 283.465 10.000 re S
		0.000 0.000 0.000 RG
		28.346 575.197 283.465 30.000 re S
		BT /FNT2 15 Tf ET
		BT 128 590 Td (Helvetica 15) Tj ET
		0.745 0.745 0.745 RG
		28.346 565.197 283.465 10.000 re S
		28.346 555.197 283.465 10.000 re S
		28.346 545.197 283.465 10.000 re S
		0.000 0.000 0.000 RG
		28.346 545.197 283.465 30.000 re S
		BT /FNT2 20 Tf ET
		BT 115 555 Td (Helvetica 20) Tj ET
		0.745 0.745 0.745 RG
		28.346 535.197 283.465 10.000 re S
		28.346 525.197 283.465 10.000 re S
		28.346 515.197 283.465 10.000 re S
		0.000 0.000 0.000 RG
		28.346 515.197 283.465 30.000 re S
		BT /FNT2 25 Tf ET
		BT 101 520 Td (Helvetica 25) Tj ET
		0.745 0.745 0.745 RG
		28.346 505.197 283.465 10.000 re S
		28.346 495.197 283.465 10.000 re S
		28.346 485.197 283.465 10.000 re S
		0.000 0.000 0.000 RG
		28.346 485.197 283.465 30.000 re S
		BT /FNT2 30 Tf ET
		BT 87 485 Td (Helvetica 30) Tj ET
		endstream
		endobj
		5 0 obj <</Type/Font/Subtype/Type1/Name/FNT1
		/BaseFont/Times-Bold
		/Encoding/StandardEncoding>>
		endobj
		6 0 obj <</Type/Font/Subtype/Type1/Name/FNT2
		/BaseFont/Helvetica
		/Encoding/StandardEncoding>>
		endobj
		xref
		0 7
		0000000000 65535 f
		0000000010 00000 n
		0000000056 00000 n
		0000000130 00000 n
		0000000241 00000 n
		0000002733 00000 n
		0000002835 00000 n
		trailer
		<</Size 7/Root 1 0 R>>
		startxref
		2936
		%%EOF
		`

		ComparePDF(t, doc.Bytes(), expect)
		// doc.SaveFile("~~Test_PDF_FontSize_.pdf")
	}()

} //                                                          Test_PDF_FontSize_

//end
