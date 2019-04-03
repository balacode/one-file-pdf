// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2019-04-03 10:27:56 BC711A           one-file-pdf/utest/[doc_keywords.go]
// -----------------------------------------------------------------------------

package utest

import (
	"fmt"
	"testing"

	"github.com/balacode/one-file-pdf"
)

// Test_PDF_DocKeywords_ is the unit test for
func Test_PDF_DocKeywords_(t *testing.T) {

	// -------------------------------------------------------------------------
	// (ob *PDF) DocKeywords() string
	//
	fmt.Println("Test PDF.DocKeywords()")

	func() {
		var doc pdf.PDF // uninitialized PDF
		TEqual(t, doc.DocKeywords(), "")
	}()

	func() {
		doc := pdf.NewPDF("A4") // initialized PDF
		TEqual(t, doc.DocKeywords(), "")
	}()

	// -------------------------------------------------------------------------
	// (ob *PDF) SetDocKeywords(s string) *PDF
	//
	fmt.Println("Test PDF.SetDocKeywords()")

	func() {
		var doc pdf.PDF // uninitialized PDF
		TEqual(t, doc.SetDocKeywords("Abcdefg").DocKeywords(), "Abcdefg")
	}()

	func() {
		doc := pdf.NewPDF("A4") // initialized PDF
		TEqual(t, doc.SetDocKeywords("Abcdefg").DocKeywords(), "Abcdefg")
	}()

	// -------------------------------------------------------------------------
	// Test PDF generation
	//
	func() {
		doc := pdf.NewPDF("A4") // initialized PDF
		doc.SetCompression(false).SetDocKeywords("'Keywords' metadata entry")

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
		5 0 obj <</Type/Info/Keywords ('Keywords' metadata entry)>>
		endobj
		xref
		0 6
		0000000000 65535 f
		0000000010 00000 n
		0000000056 00000 n
		0000000130 00000 n
		0000000189 00000 n
		0000000238 00000 n
		trailer
		<</Size 6/Root 1 0 R/Info 5 0 R>>
		startxref
		306
		%%EOF
		`

		ComparePDF(t, doc.Bytes(), want)
	}()

} //                                                       Test_PDF_DocKeywords_

//end
