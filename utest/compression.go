// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2019-04-03 10:27:56 390A2D            one-file-pdf/utest/[compression.go]
// -----------------------------------------------------------------------------

package utest

import (
	"fmt"
	"testing"

	"github.com/balacode/one-file-pdf"
)

// Test_PDF_Compression_ tests PDF.Compression() and SetCompression()
func Test_PDF_Compression_(t *testing.T) {
	fmt.Println("Test PDF.Compression()")

	draw := func(doc *pdf.PDF) {
		doc.SetUnits("cm").
			SetXY(1, 1).
			SetFont("Helvetica", 10).
			DrawText("Hello World! Hello World!")
	}

	func() {
		var doc pdf.PDF // uninitialized PDF
		TEqual(t, doc.Compression(), true)
		FailIfHasErrors(t, doc.Errors)
	}()

	func() {
		doc := pdf.NewPDF("A4") // initialized PDF
		TEqual(t, doc.Compression(), true)
		FailIfHasErrors(t, doc.Errors)
	}()

	// generate a simple PDF with compression turned on
	func() {

		const want = `
		%PDF-1.4
		1 0 obj <</Type/Catalog/Pages 2 0 R>>
		endobj
		2 0 obj <</Type/Pages/Count 1/MediaBox[0 0 595 841]/Kids[3 0 R]>>
		endobj
		3 0 obj <</Type/Page/Parent 2 0 R/Contents 4 0 R
		/Resources <</Font <</FNT1 5 0 R>> >> >>
		endobj
		4 0 obj <</Filter/FlateDecode/Length 81>> stream
		0A 78 9C 72 0A 51 D0 77 F3 0B 31 54 30 34 50 08
		49 53 70 0D E1 52 30 D0 33 30 30 40 21 8B D2 B9
		30 05 83 DC B9 9C 42 14 8C 2C 14 2C 0C 8D 15 42
		52 14 34 3C 52 73 72 F2 15 C2 F3 8B 72 52 14 15
		90 39 9A 0A 21 59 20 93 01 01 00 00 FF FF F6 FE
		19 77 0A
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
		0000000378 00000 n
		trailer
		<</Size 6/Root 1 0 R>>
		startxref
		479
		%%EOF
		`

		doc := pdf.NewPDF("A4") // initialized PDF
		doc.SetCompression(true)
		draw(&doc)
		FailIfHasErrors(t, doc.Errors)
		ComparePDF(t, doc.Bytes(), want)
	}()

	// generate a simple PDF with compression turned off
	func() {
		want := `
		%PDF-1.4
		1 0 obj <</Type/Catalog/Pages 2 0 R>>
		endobj
		2 0 obj <</Type/Pages/Count 1/MediaBox[0 0 595 841]/Kids[3 0 R]>>
		endobj
		3 0 obj <</Type/Page/Parent 2 0 R/Contents 4 0 R
		/Resources <</Font <</FNT1 5 0 R>> >> >>
		endobj
		4 0 obj <</Length 108>> stream
		BT /FNT1 10 Tf ET
		0.000 0.000 0.000 rg
		0.000 0.000 0.000 RG
		BT 28 813 Td (Hello World! Hello World!) Tj ET
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
		0000000387 00000 n
		trailer
		<</Size 6/Root 1 0 R>>
		startxref
		488
		%%EOF
		`

		doc := pdf.NewPDF("A4") // initialized PDF
		doc.SetCompression(false)
		draw(&doc)
		FailIfHasErrors(t, doc.Errors)
		ComparePDF(t, doc.Bytes(), want)
	}()

} //                                                       Test_PDF_Compression_

//end
