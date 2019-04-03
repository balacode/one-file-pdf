// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2019-04-03 10:03:38 842983             one-file-pdf/utest/[line_width.go]
// -----------------------------------------------------------------------------

package utest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/balacode/one-file-pdf"
)

// Test_PDF_LineWidth_ is the unit test for PDF.LineWidth()
// go test --run Test_PDF_LineWidth_
func Test_PDF_LineWidth_(t *testing.T) {
	fmt.Println("Test PDF.LineWidth()")
	//
	// LineWidth of new PDF must be 1 point
	func() {
		var doc pdf.PDF
		TEqual(t, doc.LineWidth(), 1)
	}()
	func() {
		doc := pdf.NewPDF("A4")
		TEqual(t, doc.LineWidth(), 1)
	}()
	//
	// SetLineWidth() has effect on the property?
	func() {
		var doc pdf.PDF
		doc.SetLineWidth(42)
		TEqual(t, doc.LineWidth(), 42)
	}()
	func() {
		doc := pdf.NewPDF("A4")
		doc.SetLineWidth(7)
		TEqual(t, doc.LineWidth(), 7)
	}()

	// -------------------------------------------------------------------------
	// Test PDF generation

	func() {
		doc := pdf.NewPDF("A4")
		doc.SetCompression(false).
			SetUnits("cm").
			//DrawUnitGrid().
			SetXY(1, 1.5).SetColor("Indigo").
			SetFont("Helvetica", 16).
			DrawText("Test PDF.LineWidth()").
			SetFont("Helvetica", 9)
		y := 2.0
		for _, w := range []float64{0.1, 0.2, 0.5, 1, 5, 10, 15, 20, 25} {
			doc.SetColor("Dark Gray").
				SetXY(1.0, y+0.3).DrawText(" y = "+strconv.Itoa(int(y))).
				SetXY(2.3, y+0.3).DrawText(" w = "+fmt.Sprintf("%0.1f", w)).
				SetColor("Gray").SetLineWidth(0.1).DrawLine(1, y, 20, y).
				SetColor("Indigo").SetLineWidth(w).DrawLine(4, y, 15, y)
			y += 1
		}
		const expect = `
		%PDF-1.4
		1 0 obj <</Type/Catalog/Pages 2 0 R>>
		endobj
		2 0 obj <</Type/Pages/Count 1/MediaBox[0 0 595 841]/Kids[3 0 R]>>
		endobj
		3 0 obj <</Type/Page/Parent 2 0 R/Contents 4 0 R
		/Resources <</Font <</FNT1 5 0 R>> >> >>
		endobj
		4 0 obj <</Length 2617>> stream
		BT /FNT1 16 Tf ET
		0.294 0.000 0.510 rg
		0.294 0.000 0.510 RG
		BT 28 799 Td (Test PDF.LineWidth\(\)) Tj ET
		BT /FNT1 9 Tf ET
		0.663 0.663 0.663 rg
		0.663 0.663 0.663 RG
		BT 28 776 Td ( y = 2) Tj ET
		BT 65 776 Td ( w = 0.1) Tj ET
		0.745 0.745 0.745 rg
		0.745 0.745 0.745 RG
		0.100 w
		28.346 785.197 m 566.929 785.197 l S
		0.294 0.000 0.510 rg
		0.294 0.000 0.510 RG
		113.386 785.197 m 425.197 785.197 l S
		0.663 0.663 0.663 rg
		0.663 0.663 0.663 RG
		BT 28 748 Td ( y = 3) Tj ET
		BT 65 748 Td ( w = 0.2) Tj ET
		0.745 0.745 0.745 rg
		0.745 0.745 0.745 RG
		28.346 756.850 m 566.929 756.850 l S
		0.294 0.000 0.510 rg
		0.294 0.000 0.510 RG
		0.200 w
		113.386 756.850 m 425.197 756.850 l S
		0.663 0.663 0.663 rg
		0.663 0.663 0.663 RG
		BT 28 720 Td ( y = 4) Tj ET
		BT 65 720 Td ( w = 0.5) Tj ET
		0.745 0.745 0.745 rg
		0.745 0.745 0.745 RG
		0.100 w
		28.346 728.504 m 566.929 728.504 l S
		0.294 0.000 0.510 rg
		0.294 0.000 0.510 RG
		0.500 w
		113.386 728.504 m 425.197 728.504 l S
		0.663 0.663 0.663 rg
		0.663 0.663 0.663 RG
		BT 28 691 Td ( y = 5) Tj ET
		BT 65 691 Td ( w = 1.0) Tj ET
		0.745 0.745 0.745 rg
		0.745 0.745 0.745 RG
		0.100 w
		28.346 700.157 m 566.929 700.157 l S
		0.294 0.000 0.510 rg
		0.294 0.000 0.510 RG
		1.000 w
		113.386 700.157 m 425.197 700.157 l S
		0.663 0.663 0.663 rg
		0.663 0.663 0.663 RG
		BT 28 663 Td ( y = 6) Tj ET
		BT 65 663 Td ( w = 5.0) Tj ET
		0.745 0.745 0.745 rg
		0.745 0.745 0.745 RG
		0.100 w
		28.346 671.811 m 566.929 671.811 l S
		0.294 0.000 0.510 rg
		0.294 0.000 0.510 RG
		5.000 w
		113.386 671.811 m 425.197 671.811 l S
		0.663 0.663 0.663 rg
		0.663 0.663 0.663 RG
		BT 28 634 Td ( y = 7) Tj ET
		BT 65 634 Td ( w = 10.0) Tj ET
		0.745 0.745 0.745 rg
		0.745 0.745 0.745 RG
		0.100 w
		28.346 643.465 m 566.929 643.465 l S
		0.294 0.000 0.510 rg
		0.294 0.000 0.510 RG
		10.000 w
		113.386 643.465 m 425.197 643.465 l S
		0.663 0.663 0.663 rg
		0.663 0.663 0.663 RG
		BT 28 606 Td ( y = 8) Tj ET
		BT 65 606 Td ( w = 15.0) Tj ET
		0.745 0.745 0.745 rg
		0.745 0.745 0.745 RG
		0.100 w
		28.346 615.118 m 566.929 615.118 l S
		0.294 0.000 0.510 rg
		0.294 0.000 0.510 RG
		15.000 w
		113.386 615.118 m 425.197 615.118 l S
		0.663 0.663 0.663 rg
		0.663 0.663 0.663 RG
		BT 28 578 Td ( y = 9) Tj ET
		BT 65 578 Td ( w = 20.0) Tj ET
		0.745 0.745 0.745 rg
		0.745 0.745 0.745 RG
		0.100 w
		28.346 586.772 m 566.929 586.772 l S
		0.294 0.000 0.510 rg
		0.294 0.000 0.510 RG
		20.000 w
		113.386 586.772 m 425.197 586.772 l S
		0.663 0.663 0.663 rg
		0.663 0.663 0.663 RG
		BT 28 549 Td ( y = 10) Tj ET
		BT 65 549 Td ( w = 25.0) Tj ET
		0.745 0.745 0.745 rg
		0.745 0.745 0.745 RG
		0.100 w
		28.346 558.425 m 566.929 558.425 l S
		0.294 0.000 0.510 rg
		0.294 0.000 0.510 RG
		25.000 w
		113.386 558.425 m 425.197 558.425 l S
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
		0000002897 00000 n
		trailer
		<</Size 6/Root 1 0 R>>
		startxref
		2998
		%%EOF
        `
		ComparePDF(t, doc.Bytes(), expect)
		// doc.SaveFile("~~Test_PDF_LineWidth_.pdf")
	}()

} //                                                         Test_PDF_LineWidth_

//end
