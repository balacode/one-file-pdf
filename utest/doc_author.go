// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-05-13 01:54:23 0A0CC1             one-file-pdf/utest/[doc_author.go]
// -----------------------------------------------------------------------------

package utest

import (
	"fmt"
	"testing"

	"github.com/balacode/one-file-pdf"
)

// Test_PDF_DocAuthor_ is the unit test for
func Test_PDF_DocAuthor_(t *testing.T) {

	// -------------------------------------------------------------------------
	// (ob *PDF) DocAuthor() string
	//
	fmt.Println("Test PDF.DocAuthor()")

	func() {
		var doc pdf.PDF // uninitialized PDF
		TEqual(t, doc.DocAuthor(), "")
	}()

	func() {
		var doc = pdf.NewPDF("A4") // initialized PDF
		TEqual(t, doc.DocAuthor(), "")
	}()

	// -------------------------------------------------------------------------
	// (ob *PDF) SetDocAuthor(s string) *PDF
	//
	fmt.Println("Test PDF.SetDocAuthor()")

	func() {
		var doc pdf.PDF // uninitialized PDF
		TEqual(t, doc.SetDocAuthor("Abcdefg").DocAuthor(), "Abcdefg")
	}()

	func() {
		var doc = pdf.NewPDF("A4") // initialized PDF
		TEqual(t, doc.SetDocAuthor("Abcdefg").DocAuthor(), "Abcdefg")
	}()

	// -------------------------------------------------------------------------
	// Test PDF generation
	//
	func() {
		var doc = pdf.NewPDF("A4") // initialized PDF
		doc.SetCompression(false).SetDocAuthor("'Author' metadata entry")

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
		5 0 obj <</Type/Info/Author ('Author' metadata entry)>>
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
		302
		%%EOF
		`

		ComparePDF(t, doc.Bytes(), expect)
	}()

} //                                                         Test_PDF_DocAuthor_

//end
