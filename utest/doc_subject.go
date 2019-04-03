// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2019-04-03 10:27:56 3626FF            one-file-pdf/utest/[doc_subject.go]
// -----------------------------------------------------------------------------

package utest

import (
	"fmt"
	"testing"

	"github.com/balacode/one-file-pdf"
)

// Test_PDF_DocSubject_ is the unit test for
func Test_PDF_DocSubject_(t *testing.T) {

	// -------------------------------------------------------------------------
	// (ob *PDF) DocSubject() string
	//
	fmt.Println("Test PDF.DocSubject()")

	func() {
		var doc pdf.PDF // uninitialized PDF
		TEqual(t, doc.DocSubject(), "")
	}()

	func() {
		doc := pdf.NewPDF("A4") // initialized PDF
		TEqual(t, doc.DocSubject(), "")
	}()

	// -------------------------------------------------------------------------
	// (ob *PDF) SetDocSubject(s string) *PDF
	//
	fmt.Println("Test PDF.SetDocSubject()")

	func() {
		var doc pdf.PDF // uninitialized PDF
		TEqual(t, doc.SetDocSubject("Abcdefg").DocSubject(), "Abcdefg")
	}()

	func() {
		doc := pdf.NewPDF("A4") // initialized PDF
		TEqual(t, doc.SetDocSubject("Abcdefg").DocSubject(), "Abcdefg")
	}()

	// -------------------------------------------------------------------------
	// Test PDF generation
	func() {
		doc := pdf.NewPDF("A4") // initialized PDF
		doc.SetCompression(false).SetDocSubject("'Subject' metadata entry")

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
		5 0 obj <</Type/Info/Subject ('Subject' metadata entry)>>
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

		ComparePDF(t, doc.Bytes(), want)
	}()

} //                                                        Test_PDF_DocSubject_

//end
