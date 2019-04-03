// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2019-04-03 10:27:56 6F33E4            one-file-pdf/utest/[draw_circle.go]
// -----------------------------------------------------------------------------

package utest

import (
	"fmt"
	"testing"

	"github.com/balacode/one-file-pdf"
)

// Test_PDF_DrawCircle_ is the unit test for
// PDF.DrawCircle(x, y, radius float64, fill ...bool) *PDF
//
// Runs the test by drawing three concentric
// circles and one small filled circle
func Test_PDF_DrawCircle_(t *testing.T) {
	fmt.Println("Test PDF.DrawCircle()")
	doc := pdf.NewPDF("20cm x 20cm")
	x, y := 10.0, 10.0 // center of page
	doc.SetCompression(false).
		SetUnits("cm").
		SetLineWidth(5).
		SetColor("Black").DrawCircle(x, y, 0.5, true). // fill
		SetColor("Red").DrawCircle(x, y, 2).
		SetColor("DarkGreen").DrawCircle(x, y, 4.5).
		SetColor("Blue").DrawCircle(x, y, 8.5)

	const want = `
	%PDF-1.4
	1 0 obj <</Type/Catalog/Pages 2 0 R>>
	endobj
	2 0 obj <</Type/Pages/Count 1/MediaBox[0 0 566 566]/Kids[3 0 R]>>
	endobj
	3 0 obj <</Type/Page/Parent 2 0 R/Contents 4 0 R>>
	endobj
	4 0 obj <</Length 1003>> stream
	0.000 0.000 0.000 rg
	0.000 0.000 0.000 RG
	5.000 w
	269.291 283.465 m
	269.291 291.292 275.637 297.638 283.465 297.638 c
	291.292 297.638 297.638 291.292 297.638 283.465 c
	297.638 275.637 291.292 269.291 283.465 269.291 c
	275.637 269.291 269.291 275.637 269.291 283.465 c
	b
	1.000 0.000 0.000 RG
	226.772 283.465 m
	226.772 314.775 252.154 340.157 283.465 340.157 c
	314.775 340.157 340.157 314.775 340.157 283.465 c
	340.157 252.154 314.775 226.772 283.465 226.772 c
	252.154 226.772 226.772 252.154 226.772 283.465 c
	S
	0.000 0.392 0.000 RG
	155.906 283.465 m
	155.906 353.913 213.016 411.024 283.465 411.024 c
	353.913 411.024 411.024 353.913 411.024 283.465 c
	411.024 213.016 353.913 155.906 283.465 155.906 c
	213.016 155.906 155.906 213.016 155.906 283.465 c
	S
	0.000 0.000 1.000 RG
	42.520 283.465 m
	42.520 416.535 150.394 524.409 283.465 524.409 c
	416.535 524.409 524.409 416.535 524.409 283.465 c
	524.409 150.394 416.535 42.520 283.465 42.520 c
	150.394 42.520 42.520 150.394 42.520 283.465 c
	S
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
	1244
	%%EOF
	`

	ComparePDF(t, doc.Bytes(), want)
} //                                                        Test_PDF_DrawCircle_

//end
