// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-03-24 22:56:46 DD03C1                               [public_test.go]
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
		// # Read-Only Properties (pdf *PDF)

		utest.CurrentPage, // CurrentPage() int
		utest.PageHeight,  // PageHeight() float64
		utest.PageWidth,   // PageWidth() float64

		// ---------------------------------------------------------------------
		// # Properties

		// Color() color.RGBA
		// SetColor(nameOrHTMLColor string) *PDF
		// SetColorRGB(red, green, blue int) *PDF
		utest.Color,

		utest.Compression, // Compression() bool

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

		// Units() string
		// SetUnits(unitName string) *PDF
		utest.Units,

		// ---------------------------------------------------------------------
		// # Methods (pdf *PDF)

		// DrawBox(x, y, width, height float64, fill ...bool) *PDF
		utest.DrawBox,

		// DrawCircle(x, y, radius float64, fill ...bool) *PDF
		utest.DrawCircle,

		// DrawImage(x, y, height float64, fileNameOrBytes interface{},
		//     backColor ...string) *PDF
		utest.DrawImage,

		// DrawText(s string) *PDF
		utest.DrawText,

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

		// ---------------------------------------------------------------------
		// # Metrics Methods (pdf *PDF)

		// ToColor(nameOrHTMLColor string) (color.RGBA, error)
		utest.ToColor,

		// ---------------------------------------------------------------------
		// # Metrics Methods (pdf *PDF)

		// ToPoints(numberAndUnit string) (float64, error)
		utest.ToPoints,

		// ToUnits(points float64) float64
		utest.ToUnits,

		// ---------------------------------------------------------------------
		// # Error Handling Methods (pdf *PDF)

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
