// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-04-23 11:32:14 58C4F0                               [public_test.go]
// -----------------------------------------------------------------------------

package pdf_test

/*
This file provides the entry points to run unit tests on the public API.

The actual unit test functions are implemented in a separate 'utest'
package, each named after the PDF method it tests, and not prefixed
with 'Test'. That is because utest has to be imported as a normal
package. 'utest' is only imported here and only used for testing.

The test files are kept in a separate folder to avoid cluttering
this root folder. Every unit test is in a separate file named
after the tested method.

To generate a test coverage report use:
	go test -coverprofile cover.out
	go tool cover -html=cover.out
*/

import "testing" // standard

import "github.com/balacode/one-file-pdf/utest"

// Every tested public method must be added here, or it won't be tested:
// go test --run TestPublicAPI
func TestPublicAPI(t *testing.T) {
	for _, fn := range []func(*testing.T){

		// ---------------------------------------------------------------------
		// # Constructor

		utest.NewPDF, // NewPDF(paperSize string) PDF

		// ---------------------------------------------------------------------
		// # Read-Only Properties (ob *PDF)

		utest.CurrentPage, // CurrentPage() int
		utest.PageHeight,  // PageHeight() float64
		utest.PageWidth,   // PageWidth() float64

		// ---------------------------------------------------------------------
		// # Properties

		utest.Color,
		// Color() color.RGBA
		// SetColor(nameOrHTMLColor string) *PDF
		// SetColorRGB(red, green, blue int) *PDF

		utest.Compression,
		// Compression() bool
		// SetCompression(compress bool) *PDF

		utest.DocAuthor,
		// DocAuthor() string
		// SetDocAuthor(s string) *PDF

		utest.DocCreator,
		// DocCreator() string
		// SetDocCreator(s string) *PDF

		utest.DocKeywords,
		// DocKeywords() string
		// SetDocKeywords(s string) *PDF

		utest.DocSubject,
		// DocSubject() string
		// SetDocSubject(s string) *PDF

		utest.DocTitle,
		// DocTitle() string
		// SetDocTitle(s string) *PDF

		utest.FontName,
		// FontName() string
		// SetFontName(name string) *PDF

		utest.SetFont, // SetFont(name string, points float64) *PDF

		//TODO: utest.FontSize,
		// FontSize() float64
		// SetFontSize(points float64) *PDF

		utest.HorizontalScaling,
		// HorizontalScaling() uint16
		// SetHorizontalScaling(percent uint16) *PDF

		utest.LineWidth,
		// LineWidth() float64
		// SetLineWidth(points float64) *PDF

		utest.X,
		// X() float64
		// SetX(x float64) *PDF

		//TODO: utest.SetXY, // SetXY(x, y float64) *PDF

		utest.Y,
		// Y() float64
		// SetY(y float64) *PDF

		utest.Units,
		// Units() string
		// SetUnits(unitName string) *PDF

		// ---------------------------------------------------------------------
		// # Methods (ob *PDF)

		//TODO: utest.AddPage, // AddPage() *PDF

		//TODO: utest.Bytes, // Bytes() []byte

		utest.DrawBox,
		// DrawBox(x, y, width, height float64, fill ...bool) *PDF

		utest.DrawCircle, // DrawCircle(x, y, radius float64, fill ...bool) *PDF

		//TODO: utest.DrawEllipse,
		// DrawEllipse(x, y, xRadius, yRadius float64, fill ...bool) *PDF

		utest.DrawImage,
		// DrawImage(x, y, height float64, fileNameOrBytes interface{},
		//     backColor ...string) *PDF

		//TODO: utest.DrawLine, // DrawLine(x1, y1, x2, y2 float64) *PDF

		utest.DrawText, // DrawText(s string) *PDF

		//TODO: utest.DrawTextAlignedToBox,
		// DrawTextAlignedToBox(
		//     x, y, width, height float64, align, text string) *PDF

		utest.DrawTextAt, // DrawTextAt(x, y float64, text string) *PDF

		utest.DrawTextInBox,
		// DrawTextInBox(
		//     x, y, width, height float64, align, text string ) *PDF

		utest.DrawUnitGrid, // DrawUnitGrid() *PDF
		utest.FillBox,      // FillBox(x, y, width, height float64) *PDF
		utest.FillCircle,   // FillCircle(x, y, radius float64) *PDF

		//TODO: utest.FillEllipse,
		// FillEllipse(x, y, xRadius, yRadius float64) *PDF

		//TODO: utest.NextLine, // NextLine() *PDF

		utest.Reset, // Reset() *PDF

		//TODO: utest.SaveFile, // SaveFile(filename string) error

		//TODO: utest.SetColumnWidths,
		// SetColumnWidths(widths ...float64) *PDF

		// ---------------------------------------------------------------------
		// # Metrics Methods (ob *PDF)

		//TODO: utest.TextWidth, // TextWidth(s string) float64

		utest.ToColorT1,
		utest.ToColorT2,
		// ToColor(nameOrHTMLColor string) (color.RGBA, error)

		// ---------------------------------------------------------------------
		// # Metrics Methods (ob *PDF)

		utest.ToPoints, // ToPoints(numberAndUnit string) (float64, error)
		utest.ToUnits,  // ToUnits(points float64) float64

		//TODO: utest.WrapTextLines,
		// WrapTextLines(width float64, text string) (ret []string)

		// ---------------------------------------------------------------------
		// # Error Handling Methods (ob *PDF)

		utest.Clean,     // Clean() *PDF
		utest.Errors,    // Errors() []error
		utest.PullError, // PullError() error
	} {
		fn(t)
	}
} //                                                               TestPublicAPI

//end
