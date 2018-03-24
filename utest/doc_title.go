// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-03-24 18:56:20 491B6F                           [utest/doc_title.go]
// -----------------------------------------------------------------------------

package utest

import "fmt"     // standard
import "testing" // standard

import "github.com/balacode/one-file-pdf"

// DocTitle is the unit test for
func DocTitle(t *testing.T) {

	// -------------------------------------------------------------------------
	// (pdf *PDF) DocTitle() string
	//
	fmt.Println("utest.DocTitle")

	func() {
		var doc pdf.PDF // uninitialized PDF
		TEqual(t, doc.DocTitle(), "")
	}()

	func() {
		var doc = pdf.NewPDF("A4") // initialized PDF
		TEqual(t, doc.DocTitle(), "")
	}()

	// -------------------------------------------------------------------------
	// (pdf *PDF) SetDocTitle(s string) *PDF
	//
	fmt.Println("utest.SetDocTitle")

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
		1 0 obj<</Type/Catalog/Pages 2 0 R>>
		endobj
		2 0 obj<</Type/Pages/Count 1/MediaBox[0 0 595 841]/Kids[3 0 R]>>
		endobj
		3 0 obj<</Type/Page/Parent 2 0 R/Contents 4 0 R>>
		endobj
		4 0 obj <</Length 0>>stream
		endstream
		5 0 obj<</Type/Info/Title ('Title' metadata entry)>>
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
		281
		%%EOF
        `
		pdfCompare(t, doc.Bytes(), expect, pdfStreamsInText)
	}()

} //                                                                    DocTitle

//end
