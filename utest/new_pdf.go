// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-03-22 02:48:21 286601                             [utest/new_pdf.go]
// -----------------------------------------------------------------------------

package utest

import "fmt"     // standard
import "testing" // standard

import "github.com/balacode/one-file-pdf"

// NewPDF is the unit test for PDF.NewPDF
func NewPDF(t *testing.T) {
	fmt.Println("utest.NewPDF")
	//
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
	xref
	0 5
	0000000000 65535 f
	0000000009 00000 n
	0000000053 00000 n
	0000000125 00000 n
	0000000182 00000 n
	trailer
	<</Size 5/Root 1 0 R>>
	startxref
	221
	%%EOF
	`

	// test NewPDF() and Bytes() while calling AddPage()
	func() {
		var ob = pdf.NewPDF("A4")
		var result = ob.SetCompression(false).AddPage().Bytes()
		pdfCompare(t, result, expect, pdfStreamsInText)
	}()

	// test NewPDF() and Bytes() without calling AddPage()
	func() {
		var ob = pdf.NewPDF("A4")
		var result = ob.SetCompression(false).Bytes()
		pdfCompare(t, result, expect, pdfStreamsInText)
	}()

} //                                                                      NewPDF

//end
