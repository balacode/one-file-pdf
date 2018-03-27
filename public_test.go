// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-03-27 22:10:13 C842AF                               [public_test.go]
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
// go  test --run Test_<name>_
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

		// Color() color.RGBA
		// SetColor(nameOrHTMLColor string) *PDF
		// SetColorRGB(red, green, blue int) *PDF
		utest.Color,

		// Compression() bool
		// SetCompression(compress bool) *PDF
		utest.Compression,

		// DocAuthor() string
		// SetDocAuthor(s string) *PDF
		utest.DocAuthor,

		// DocCreator() string
		// SetDocCreator(s string) *PDF
		utest.DocCreator,

		// DocKeywords() string
		// SetDocKeywords(s string) *PDF
		utest.DocKeywords,

		// DocSubject() string
		// SetDocSubject(s string) *PDF
		utest.DocSubject,

		// DocTitle() string
		// SetDocTitle(s string) *PDF
		utest.DocTitle,

		// FontName() string
		// SetFontName(name string) *PDF
		utest.FontName,

		//TODO: SetFont(name string, points float64) *PDF

		//TODO: FontSize() float64
		//TODO: SetFontSize(points float64) *PDF

		//TODO: HorizontalScaling() uint16
		//TODO: SetHorizontalScaling(percent uint16) *PDF

		//TODO: LineWidth() float64
		//TODO: SetLineWidth(points float64) *PDF

		//TODO: SetX(x float64) *PDF
		//TODO: SetXY(x, y float64) *PDF
		//TODO: SetY(y float64) *PDF

		// Units() string
		// SetUnits(unitName string) *PDF
		utest.Units,

		//TODO: X() float64
		//TODO: Y() float64

		// ---------------------------------------------------------------------
		// # Methods (ob *PDF)

		//TODO: AddPage() *PDF

		//TODO: Bytes() []byte

		// DrawBox(x, y, width, height float64, fill ...bool) *PDF
		utest.DrawBox,

		// DrawCircle(x, y, radius float64, fill ...bool) *PDF
		utest.DrawCircle,

		//TODO: DrawEllipse(x, y, xRadius, yRadius float64, fill ...bool) *PDF

		// DrawImage(x, y, height float64, fileNameOrBytes interface{},
		//     backColor ...string) *PDF
		utest.DrawImage,

		//TODO: DrawLine(x1, y1, x2, y2 float64) *PDF

		// DrawText(s string) *PDF
		utest.DrawText,

		//TODO: DrawTextAlignedToBox(
		//          x, y, width, height float64, align, text string) *PDF

		// DrawTextAt(x, y float64, text string) *PDF
		utest.DrawTextAt,

		// DrawTextInBox(
		//     x, y, width, height float64, align, text string ) *PDF
		utest.DrawTextInBox,

		// DrawUnitGrid() *PDF
		utest.DrawUnitGrid,

		// FillBox(x, y, width, height float64) *PDF
		utest.FillBox,

		// FillCircle(x, y, radius float64) *PDF
		utest.FillCircle,

		//TODO: FillEllipse(x, y, xRadius, yRadius float64) *PDF

		//TODO: NextLine() *PDF

		// Reset() *PDF
		utest.Reset,

		//TODO: SaveFile(filename string) error

		//TODO: SetColumnWidths(widths ...float64) *PDF

		// ---------------------------------------------------------------------
		// # Metrics Methods (ob *PDF)

		//TODO: TextWidth(s string) float64

		// ToColor(nameOrHTMLColor string) (color.RGBA, error)
		utest.ToColorT1,
		utest.ToColorT2,

		// ---------------------------------------------------------------------
		// # Metrics Methods (ob *PDF)

		// ToPoints(numberAndUnit string) (float64, error)
		utest.ToPoints,

		// ToUnits(points float64) float64
		utest.ToUnits,

		//TODO: WrapTextLines(width float64, text string) (ret []string)

		// ---------------------------------------------------------------------
		// # Error Handling Methods (ob *PDF)

		// Clean() *PDF
		utest.Clean,

		// Errors() []error
		utest.Errors,

		// PullError() error
		utest.PullError,
	} {
		fn(t)
	}
} //                                                               TestPublicAPI

//end
