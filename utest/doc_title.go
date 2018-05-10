// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-05-10 22:50:31 B8F322                           [utest/doc_title.go]
// -----------------------------------------------------------------------------

package utest

import (
	"fmt"
	"testing"

	"github.com/balacode/one-file-pdf"
)

// Test_PDF_DocTitle_ is the unit test for
func Test_PDF_DocTitle_(t *testing.T) {

	// -------------------------------------------------------------------------
	// (ob *PDF) DocTitle() string
	//
	fmt.Println("Test PDF.DocTitle()")

	func() {
		var doc pdf.PDF // uninitialized PDF
		TEqual(t, doc.DocTitle(), "")
	}()

	func() {
		var doc = pdf.NewPDF("A4") // initialized PDF
		TEqual(t, doc.DocTitle(), "")
	}()

	// -------------------------------------------------------------------------
	// (ob *PDF) SetDocTitle(s string) *PDF
	//
	fmt.Println("Test PDF.SetDocTitle()")

	func() {
		var doc pdf.PDF // uninitialized PDF
		TEqual(t, doc.SetDocTitle("Abcdefg").DocTitle(), "Abcdefg")
	}()

	func() {
		var doc = pdf.NewPDF("A4") // initialized PDF
		TEqual(t, doc.SetDocTitle("Abcdefg").DocTitle(), "Abcdefg")
	}()

	// -------------------------------------------------------------------------
	// Test PDF generation

	func() {
		var doc = pdf.NewPDF("A4")
		doc.SetCompression(false).SetDocTitle("'Title' metadata entry")

		var expect = `
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
		5 0 obj <</Type/Info/Title ('Title' metadata entry)>>
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
		300
		%%EOF
		`

		ComparePDF(t, doc.Bytes(), expect)
	}()

} //                                                          Test_PDF_DocTitle_

//end
