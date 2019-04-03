// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2019-04-03 10:27:56 664F60               one-file-pdf/utest/[set_font.go]
// -----------------------------------------------------------------------------

package utest

import (
	"fmt"
	"testing"

	"github.com/balacode/one-file-pdf"
)

// Test_PDF_SetFont_ is the unit test for PDF.SetFont()
func Test_PDF_SetFont_(t *testing.T) {
	fmt.Println("Test PDF.SetFont()")
	//
	// setting font on uninitialized pdf
	for _, tc := range []struct {
		size    float64
		inName  string
		expName string
	}{
		{size: 10, inName: "Courier", expName: "Courier"},
		{size: 20, inName: "Zapf-Dingbats", expName: "Zapf-Dingbats"},
		{size: 30, inName: "ZapfDingbats", expName: "ZapfDingbats"},
		{size: 40, inName: "YeOldeScript", expName: "YeOldeScript"},
	} {
		var doc pdf.PDF // uninitialized PDF
		doc.SetFont(tc.inName, tc.size)
		TEqual(t, doc.FontName(), tc.expName)
		TEqual(t, doc.FontSize(), tc.size)
	}

	// -------------------------------------------------------------------------
	// Test PDF generation

	func() {
		doc := pdf.NewPDF("A4")
		doc.SetCompression(false).
			SetUnits("cm").
			SetFont("Times-Bold", 20).
			SetXY(1, 1).
			DrawText("Built-in PDF Fonts")
		for i, font := range []string{
			"Courier",
			"Courier-Bold",
			"Courier-BoldOblique",
			"Courier-Oblique",
			"Helvetica",
			"Helvetica-Bold",
			"Helvetica-BoldOblique",
			"Helvetica-Oblique",
			"Symbol",
			"Times-Bold",
			"Times-BoldItalic",
			"Times-Italic",
			"Times-Roman",
			"ZapfDingbats",
		} {
			y := 2.5 + float64(i)*1.8
			doc.SetXY(1, y).
				SetFont("Helvetica", 10).
				DrawText(font).
				SetXY(1, y+0.7).
				SetFont(font, 20).
				DrawText("Five hexing wizard bots jump quickly")

		}

		const want = `
		%PDF-1.4
		1 0 obj <</Type/Catalog/Pages 2 0 R>>
		endobj
		2 0 obj <</Type/Pages/Count 1/MediaBox[0 0 595 841]/Kids[3 0 R]>>
		endobj
		3 0 obj <</Type/Page/Parent 2 0 R/Contents 4 0 R
		/Resources <</Font <<
		/FNT1 5 0 R
		/FNT2 6 0 R
		/FNT3 7 0 R
		/FNT4 8 0 R
		/FNT5 9 0 R
		/FNT6 10 0 R
		/FNT7 11 0 R
		/FNT8 12 0 R
		/FNT9 13 0 R
		/FNT10 14 0 R
		/FNT11 15 0 R
		/FNT12 16 0 R
		/FNT13 17 0 R
		/FNT14 18 0 R>> >> >>
		endobj
		4 0 obj <</Length 1910>> stream
		BT /FNT1 20 Tf ET
		0.000 0.000 0.000 rg
		0.000 0.000 0.000 RG
		BT 28 813 Td (Built-in PDF Fonts) Tj ET
		BT /FNT2 10 Tf ET
		BT 28 771 Td (Courier) Tj ET
		BT /FNT3 20 Tf ET
		BT 28 751 Td (Five hexing wizard bots jump quickly) Tj ET
		BT /FNT2 10 Tf ET
		BT 28 720 Td (Courier-Bold) Tj ET
		BT /FNT4 20 Tf ET
		BT 28 700 Td (Five hexing wizard bots jump quickly) Tj ET
		BT /FNT2 10 Tf ET
		BT 28 668 Td (Courier-BoldOblique) Tj ET
		BT /FNT5 20 Tf ET
		BT 28 649 Td (Five hexing wizard bots jump quickly) Tj ET
		BT /FNT2 10 Tf ET
		BT 28 617 Td (Courier-Oblique) Tj ET
		BT /FNT6 20 Tf ET
		BT 28 598 Td (Five hexing wizard bots jump quickly) Tj ET
		BT /FNT2 10 Tf ET
		BT 28 566 Td (Helvetica) Tj ET
		BT /FNT2 20 Tf ET
		BT 28 547 Td (Five hexing wizard bots jump quickly) Tj ET
		BT /FNT2 10 Tf ET
		BT 28 515 Td (Helvetica-Bold) Tj ET
		BT /FNT7 20 Tf ET
		BT 28 496 Td (Five hexing wizard bots jump quickly) Tj ET
		BT /FNT2 10 Tf ET
		BT 28 464 Td (Helvetica-BoldOblique) Tj ET
		BT /FNT8 20 Tf ET
		BT 28 445 Td (Five hexing wizard bots jump quickly) Tj ET
		BT /FNT2 10 Tf ET
		BT 28 413 Td (Helvetica-Oblique) Tj ET
		BT /FNT9 20 Tf ET
		BT 28 394 Td (Five hexing wizard bots jump quickly) Tj ET
		BT /FNT2 10 Tf ET
		BT 28 362 Td (Symbol) Tj ET
		BT /FNT10 20 Tf ET
		BT 28 342 Td (Five hexing wizard bots jump quickly) Tj ET
		BT /FNT2 10 Tf ET
		BT 28 311 Td (Times-Bold) Tj ET
		BT /FNT1 20 Tf ET
		BT 28 291 Td (Five hexing wizard bots jump quickly) Tj ET
		BT /FNT2 10 Tf ET
		BT 28 260 Td (Times-BoldItalic) Tj ET
		BT /FNT11 20 Tf ET
		BT 28 240 Td (Five hexing wizard bots jump quickly) Tj ET
		BT /FNT2 10 Tf ET
		BT 28 209 Td (Times-Italic) Tj ET
		BT /FNT12 20 Tf ET
		BT 28 189 Td (Five hexing wizard bots jump quickly) Tj ET
		BT /FNT2 10 Tf ET
		BT 28 158 Td (Times-Roman) Tj ET
		BT /FNT13 20 Tf ET
		BT 28 138 Td (Five hexing wizard bots jump quickly) Tj ET
		BT /FNT2 10 Tf ET
		BT 28 107 Td (ZapfDingbats) Tj ET
		BT /FNT14 20 Tf ET
		BT 28 87 Td (Five hexing wizard bots jump quickly) Tj ET
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
		7 0 obj <</Type/Font/Subtype/Type1/Name/FNT3
		/BaseFont/Courier
		/Encoding/StandardEncoding>>
		endobj
		8 0 obj <</Type/Font/Subtype/Type1/Name/FNT4
		/BaseFont/Courier-Bold
		/Encoding/StandardEncoding>>
		endobj
		9 0 obj <</Type/Font/Subtype/Type1/Name/FNT5
		/BaseFont/Courier-BoldOblique
		/Encoding/StandardEncoding>>
		endobj
		10 0 obj <</Type/Font/Subtype/Type1/Name/FNT6
		/BaseFont/Courier-Oblique
		/Encoding/StandardEncoding>>
		endobj
		11 0 obj <</Type/Font/Subtype/Type1/Name/FNT7
		/BaseFont/Helvetica-Bold
		/Encoding/StandardEncoding>>
		endobj
		12 0 obj <</Type/Font/Subtype/Type1/Name/FNT8
		/BaseFont/Helvetica-BoldOblique
		/Encoding/StandardEncoding>>
		endobj
		13 0 obj <</Type/Font/Subtype/Type1/Name/FNT9
		/BaseFont/Helvetica-Oblique
		/Encoding/StandardEncoding>>
		endobj
		14 0 obj <</Type/Font/Subtype/Type1/Name/FNT10
		/BaseFont/Symbol
		/Encoding/StandardEncoding>>
		endobj
		15 0 obj <</Type/Font/Subtype/Type1/Name/FNT11
		/BaseFont/Times-BoldItalic
		/Encoding/StandardEncoding>>
		endobj
		16 0 obj <</Type/Font/Subtype/Type1/Name/FNT12
		/BaseFont/Times-Italic
		/Encoding/StandardEncoding>>
		endobj
		17 0 obj <</Type/Font/Subtype/Type1/Name/FNT13
		/BaseFont/Times-Roman
		/Encoding/StandardEncoding>>
		endobj
		18 0 obj <</Type/Font/Subtype/Type1/Name/FNT14
		/BaseFont/ZapfDingbats
		/Encoding/StandardEncoding>>
		endobj
		xref
		0 19
		0000000000 65535 f
		0000000010 00000 n
		0000000056 00000 n
		0000000130 00000 n
		0000000399 00000 n
		0000002361 00000 n
		0000002463 00000 n
		0000002564 00000 n
		0000002663 00000 n
		0000002767 00000 n
		0000002878 00000 n
		0000002986 00000 n
		0000003093 00000 n
		0000003207 00000 n
		0000003317 00000 n
		0000003417 00000 n
		0000003527 00000 n
		0000003633 00000 n
		0000003738 00000 n
		trailer
		<</Size 19/Root 1 0 R>>
		startxref
		3844
		%%EOF
		`

		ComparePDF(t, doc.Bytes(), want)
		doc.SaveFile("~~font_sample.pdf")
	}()

} //                                                           Test_PDF_SetFont_

//end
