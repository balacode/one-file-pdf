// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2019-04-03 10:03:38 8D822A            one-file-pdf/utest/[doc_creator.go]
// -----------------------------------------------------------------------------

package utest

import (
	"fmt"
	"testing"

	"github.com/balacode/one-file-pdf"
)

// Test_PDF_DocCreator_ is the unit test for
func Test_PDF_DocCreator_(t *testing.T) {

	// -------------------------------------------------------------------------
	// (ob *PDF) DocCreator() string
	//
	fmt.Println("Test PDF.DocCreator()")

	func() {
		var doc pdf.PDF // uninitialized PDF
		TEqual(t, doc.DocCreator(), "")
	}()

	func() {
		doc := pdf.NewPDF("A4") // initialized PDF
		TEqual(t, doc.DocCreator(), "")
	}()

	// -------------------------------------------------------------------------
	// (ob *PDF) SetDocCreator(s string) *PDF
	//
	fmt.Println("Test PDF.SetDocCreator()")

	func() {
		var doc pdf.PDF // uninitialized PDF
		TEqual(t, doc.SetDocCreator("Abcdefg").DocCreator(), "Abcdefg")
	}()

	func() {
		doc := pdf.NewPDF("A4") // initialized PDF
		TEqual(t, doc.SetDocCreator("Abcdefg").DocCreator(), "Abcdefg")
	}()

	// -------------------------------------------------------------------------
	// Test PDF generation
	//
	func() {
		doc := pdf.NewPDF("A4") // initialized PDF
		doc.SetCompression(false).SetDocCreator("'Creator' metadata entry")

		const expect = `
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
		5 0 obj <</Type/Info/Creator ('Creator' metadata entry)>>
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
		304
		%%EOF
		`

		ComparePDF(t, doc.Bytes(), expect)
	}()

} //                                                        Test_PDF_DocCreator_

//end
