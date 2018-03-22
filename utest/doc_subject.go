// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-03-22 03:14:56 8BF24C                         [utest/doc_subject.go]
// -----------------------------------------------------------------------------

package utest

import "fmt"     // standard
import "testing" // standard

import "github.com/balacode/one-file-pdf"

// DocSubject is the unit test for
func DocSubject(t *testing.T) {

	// -------------------------------------------------------------------------
	// (pdf *PDF) DocSubject() string
	//
	fmt.Println("utest.DocSubject")

	func() {
		var ob pdf.PDF // uninitialized PDF
		TEqual(t, ob.DocSubject(), "")
	}()

	func() {
		var ob = pdf.NewPDF("A4") // initialized PDF
		TEqual(t, ob.DocSubject(), "")
	}()

	// -------------------------------------------------------------------------
	// (pdf *PDF) SetDocSubject(s string) *PDF
	//
	fmt.Println("utest.SetDocSubject")

	func() {
		var ob pdf.PDF // uninitialized PDF
		TEqual(t, ob.SetDocSubject("Abcdefg").DocSubject(), "Abcdefg")
	}()

	func() {
		var ob = pdf.NewPDF("A4") // initialized PDF
		TEqual(t, ob.SetDocSubject("Abcdefg").DocSubject(), "Abcdefg")
	}()

	// -------------------------------------------------------------------------
	// Test PDF generation
	func() {
		var pdf = pdf.NewPDF("A4")
		pdf.SetCompression(false).
			SetDocSubject("'Subject' metadata entry")
		var expect = `
		%PDF-1.4
		1 0 obj<</Type/Catalog/Pages 2 0 R>>
		endobj
		2 0 obj<</Type/Pages/Count 1/MediaBox[0 0 595 841]/Kids[3 0 R]>>
		endobj
		3 0 obj<</Type/Page/Parent 2 0 R/Contents 4 0 R>>
		endobj
		4 0 obj <</Length 0>>stream
		endstream
		5 0 obj<</Type/Info/Subject ('Subject' metadata entry)>>
		endobj
		xref
		0 6
		0000000000 65535 f
		0000000009 00000 n
		0000000053 00000 n
		0000000125 00000 n
		0000000182 00000 n
		0000000221 00000 n
		trailer
		<</Size 6/Root 1 0 R/Info 5 0 R>>
		startxref
		285
		%%EOF
        `
		pdfCompare(t, pdf.Bytes(), expect, pdfStreamsInText)
	}()

} //                                                                  DocSubject

//end
