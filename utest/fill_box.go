// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-03-24 18:56:20 6451B5                            [utest/fill_box.go]
// -----------------------------------------------------------------------------

package utest

import "fmt"     // standard
import "testing" // standard

import "github.com/balacode/one-file-pdf"

// FillBox is the unit test for
// PDF.FillBox(x, y, width, height float64) *PDF
//
// Runs the test by filling the shape of a Monolith from 2001 Space Odyssey
func FillBox(t *testing.T) {
	fmt.Println("utest.FillBox")
	var (
		doc    = pdf.NewPDF("A4")
		x      = 6.5
		y      = 6.0
		width  = 8.0
		height = 18.0
	)
	doc.SetCompression(false).
		SetUnits("cm").
		SetColor("#1B1B1B EerieBlack").
		FillBox(x, y, width, height)
	//
	var expect = `
	%PDF-1.4
	1 0 obj<</Type/Catalog/Pages 2 0 R>>
	endobj
	2 0 obj<</Type/Pages/Count 1/MediaBox[0 0 595 841]/Kids[3 0 R]>>
	endobj
	3 0 obj<</Type/Page/Parent 2 0 R/Contents 4 0 R>>
	endobj
	4 0 obj <</Length 80>>stream
	 0.106 0.106 0.106 rg
	0.106 0.106 0.106 RG
	184.252 161.575 226.772 510.236 re b
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
	302
	%%EOF
	`
	pdfCompare(t, doc.Bytes(), expect, pdfStreamsInText)
} //                                                                     FillBox

//end
