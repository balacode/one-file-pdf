// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-04-29 23:42:24 C8940E                             [utest/new_pdf.go]
// -----------------------------------------------------------------------------

package utest

import (
	"fmt"
	"testing"

	"github.com/balacode/one-file-pdf"
)

// Test_NewPDF_ is the unit test for PDF.NewPDF
func Test_NewPDF_(t *testing.T) {
	fmt.Println("Test NewPDF()")

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
	xref
	0 5
	0000000000 65535 f
	0000000010 00000 n
	0000000056 00000 n
	0000000130 00000 n
	0000000189 00000 n
	trailer
	<</Size 5/Root 1 0 R>>
	startxref
	238
	%%EOF
	`

	// test NewPDF() and Bytes() while calling AddPage()
	func() {
		var doc = pdf.NewPDF("A4")
		var result = doc.SetCompression(false).AddPage().Bytes()
		pdfCompare(t, result, expect, pdfStreamsInText)
	}()

	// test NewPDF() and Bytes() without calling AddPage()
	func() {
		var doc = pdf.NewPDF("A4")
		var result = doc.SetCompression(false).Bytes()
		pdfCompare(t, result, expect, pdfStreamsInText)
	}()

} //                                                                Test_NewPDF_

//end
