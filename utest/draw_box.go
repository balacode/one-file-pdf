// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-03-21 01:14:03 9108C6                            [utest/draw_box.go]
// -----------------------------------------------------------------------------

package utest

import "fmt"     // standard
import "testing" // standard

import "github.com/balacode/one-file-pdf"

// DrawBox is the unit test for
// PDF.DrawBox(x, y, width, height float64, fill ...bool) *PDF
//
// Runs the test by drawing three rectangles and one filled rectangle
func DrawBox(t *testing.T) {
	fmt.Println("utest.DrawBox")
	var (
		ob = pdf.NewPDF("18cm x 18cm")
		x  = 1.0
		y  = 1.0
	)
	ob.SetCompression(false).
		SetUnits("cm").
		SetLineWidth(5).
		SetColor("Black").DrawBox(x, y, 1, 1, true). // fill
		SetColor("Red").DrawBox(x, y, 4, 4).
		SetColor("DarkGreen").DrawBox(x, y, 9, 9).
		SetColor("Blue").DrawBox(x, y, 16, 16)
	//
	var expect = `
	%PDF-1.4
	1 0 obj<</Type/Catalog/Pages 2 0 R>>
	endobj
	2 0 obj<</Type/Pages/Count 1/MediaBox[0 0 510 510]/Kids[3 0 R]>>
	endobj
	3 0 obj<</Type/Page/Parent 2 0 R/Contents 4 0 R>>
	endobj
	4 0 obj <</Length 255>>stream
	0.000 0.000 0.000 rg
	0.000 0.000 0.000 RG
	5.000 w
	28.346 453.543 28.346 28.346 re b
	1.000 0.000 0.000 RG
	28.346 368.504 113.386 113.386 re S
	0.000 0.392 0.000 RG
	28.346 226.772 255.118 255.118 re S
	0.000 0.000 1.000 RG
	28.346 28.346 453.543 453.543 re S
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
	478
	%%EOF
	`
	pdfCompare(t, ob.Bytes(), expect, pdfStreamsInText)
} //                                                                     DrawBox

//end
