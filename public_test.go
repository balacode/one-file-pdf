// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-06-07 20:25:46 6BF1EC                  one-file-pdf/[public_test.go]
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

To test a single method or property, use:
	go test --run Test_PDF_MethodName_

To generate a test coverage report, use:
	go test -coverprofile cover.out
	go tool cover -html=cover.out
*/

import (
	"testing"

	"github.com/balacode/one-file-pdf/utest"
)

// Every tested public method must be added here, or it won't be tested:

// -----------------------------------------------------------------------------
// # Constructor

// NewPDF(paperSize string) PDF
func Test_NewPDF_(t *testing.T) { utest.Test_NewPDF_(t) }

// -----------------------------------------------------------------------------
// # Read-Only Properties (ob *PDF)

// PageCount() int
func Test_PDF_PageCount_(t *testing.T) { utest.Test_PDF_PageCount_(t) }

// PageHeight() float64
func Test_PDF_PageHeight_(t *testing.T) { utest.Test_PDF_PageHeight_(t) }

// PageWidth() float64
func Test_PDF_PageWidth_(t *testing.T) { utest.Test_PDF_PageWidth_(t) }

// -----------------------------------------------------------------------------
// # Properties

// Color() color.RGBA
// SetColor(nameOrHTMLColor string) *PDF
// SetColorRGB(r, g, b byte) *PDF
func Test_PDF_Color_(t *testing.T) { utest.Test_PDF_Color_(t) }

// Compression() bool
// SetCompression(val bool) *PDF
func Test_PDF_Compression_(t *testing.T) { utest.Test_PDF_Compression_(t) }

// CurrentPage() int
// SetCurrentPage(pageNo int) *PDF
func Test_PDF_CurrentPage_(t *testing.T) { utest.Test_PDF_CurrentPage_(t) }

// DocAuthor() string
// SetDocAuthor(s string) *PDF
func Test_PDF_DocAuthor_(t *testing.T) { utest.Test_PDF_DocAuthor_(t) }

// DocCreator() string
// SetDocCreator(s string) *PDF
func Test_PDF_DocCreator_(t *testing.T) { utest.Test_PDF_DocCreator_(t) }

// DocKeywords() string
// SetDocKeywords(s string) *PDF
func Test_PDF_DocKeywords_(t *testing.T) { utest.Test_PDF_DocKeywords_(t) }

// DocSubject() string
// SetDocSubject(s string) *PDF
func Test_PDF_DocSubject_(t *testing.T) { utest.Test_PDF_DocSubject_(t) }

// DocTitle() string
// SetDocTitle(s string) *PDF
func Test_PDF_DocTitle_(t *testing.T) { utest.Test_PDF_DocTitle_(t) }

// FontName() string
// SetFontName(name string) *PDF
func Test_PDF_FontName_(t *testing.T) { utest.Test_PDF_FontName_(t) }

// FontSize() float64
// SetFontSize(points float64) *PDF
func Test_PDF_FontSize_(t *testing.T) { utest.Test_PDF_FontSize_(t) }

// SetFont(name string, points float64) *PDF
func Test_PDF_SetFont_(t *testing.T) { utest.Test_PDF_SetFont_(t) }

// HorizontalScaling() uint16
// SetHorizontalScaling(percent uint16) *PDF
func Test_PDF_HorizontalScaling_(t *testing.T) {
	utest.Test_PDF_HorizontalScaling_(t)
}

// LineWidth() float64
// SetLineWidth(points float64) *PDF
func Test_PDF_LineWidth_(t *testing.T) { utest.Test_PDF_LineWidth_(t) }

// Units() string
// SetUnits(units string) *PDF
func Test_PDF_Units_(t *testing.T) { utest.Test_PDF_Units_(t) }

// X() float64
// SetX(x float64) *PDF
func Test_PDF_X_(t *testing.T) { utest.Test_PDF_X_(t) }

// Y() float64
// SetY(y float64) *PDF
func Test_PDF_Y_(t *testing.T) { utest.Test_PDF_Y_(t) }

//TODO: func Test_PDF_SetXY_(t *testing.T) { utest.Test_PDF_SetXY_(t) }
// SetXY(x, y float64) *PDF

// -----------------------------------------------------------------------------
// # Methods (ob *PDF)

// AddPage() *PDF
//TODO: func Test_PDF_AddPage_(t *testing.T) { utest.Test_PDF_AddPage_(t) }

// Bytes() []byte
//TODO: func Test_PDF_Bytes_(t *testing.T) { utest.Test_PDF_Bytes_(t) }

// DrawBox(x, y, width, height float64, optFill ...bool) *PDF
func Test_PDF_DrawBox_(t *testing.T) { utest.Test_PDF_DrawBox_(t) }

// DrawCircle(x, y, radius float64, optFill ...bool) *PDF
func Test_PDF_DrawCircle_(t *testing.T) { utest.Test_PDF_DrawCircle_(t) }

// DrawEllipse(x, y, xRadius, yRadius float64,
//     optFill ...bool) *PDF
//TODO: func Test_PDF_DrawEllipse_(t *testing.T) {
//          utest.Test_PDF_DrawEllipse_(t)
//      }

// DrawImage(x, y, height float64, fileNameOrBytes interface{},
//     backColor ...string) *PDF
func Test_PDF_DrawImage_(t *testing.T) { utest.Test_PDF_DrawImage_(t) }

// DrawLine(x1, y1, x2, y2 float64) *PDF
//TODO: func Test_PDF_DrawLine_(t *testing.T) { utest.Test_PDF_DrawLine_(t) }

// DrawText(s string) *PDF
func Test_PDF_DrawText_(t *testing.T) { utest.Test_PDF_DrawText_(t) }

// DrawTextAlignedToBox(
//     x, y, width, height float64, align, text string) *PDF
//TODO: func Test_PDF_DrawTextAlignedToBox_(t *testing.T) {
//          utest.Test_PDF_DrawTextAlignedToBox(t)
//      }

// DrawTextAt(x, y float64, text string) *PDF
func Test_PDF_DrawTextAt_(t *testing.T) { utest.Test_PDF_DrawTextAt_(t) }

// DrawTextInBox(
//     x, y, width, height float64, align, text string) *PDF
func Test_PDF_DrawTextInBox_(t *testing.T) { utest.Test_PDF_DrawTextInBox_(t) }

// DrawUnitGrid() *PDF
func Test_PDF_DrawUnitGrid_(t *testing.T) { utest.Test_PDF_DrawUnitGrid_(t) }

// FillBox(x, y, width, height float64) *PDF
func Test_PDF_FillBox_(t *testing.T) { utest.Test_PDF_FillBox_(t) }

// FillCircle(x, y, radius float64) *PDF
func Test_PDF_FillCircle_(t *testing.T) { utest.Test_PDF_FillCircle_(t) }

// FillEllipse(x, y, xRadius, yRadius float64) *PDF
//TODO: func Test_PDF_FillEllipse_(t *testing.T) {
//          utest.Test_PDF_FillEllipse_(t)
// }

// NextLine() *PDF
//TODO: func Test_PDF_NextLine_(t *testing.T) { utest.Test_PDF_NextLine_(t) }

// Reset() *PDF
func Test_PDF_Reset_(t *testing.T) { utest.Test_PDF_Reset_(t) }

// SaveFile(filename string) error
//TODO: func Test_PDF_SaveFile_(t *testing.T) { utest.Test_PDF_SaveFile_(t) }

// SetColumnWidths(widths ...float64) *PDF
//TODO: func Test_PDF_SetColumnWidths_(t *testing.T) {
//          utest.Test_PDF_SetColumnWidths(t)
//      }

// -----------------------------------------------------------------------------
// # Metrics Methods (ob *PDF)

// TextWidth(s string) float64
//TODO: func Test_PDF_TextWidth_(t *testing.T) { utest.Test_PDF_TextWidth_(t) }

// ToColor(nameOrHTMLColor string) (color.RGBA, error)
func Test_PDF_ToColor_1_(t *testing.T) { utest.Test_PDF_ToColor_1_(t) }
func Test_PDF_ToColor_2_(t *testing.T) { utest.Test_PDF_ToColor_2_(t) }

// -----------------------------------------------------------------------------
// # Metrics Methods (ob *PDF)

// ToPoints(numberAndUnit string) (float64, error)
func Test_PDF_ToPoints_(t *testing.T) { utest.Test_PDF_ToPoints_(t) }

// ToUnits(points float64) float64
func Test_PDF_ToUnits_(t *testing.T) { utest.Test_PDF_ToUnits_(t) }

// WrapTextLines(width float64, text string) (ret []string)
//TODO: func Test_PDF_WrapTextLines_(t *testing.T) {
//          utest.Test_PDF_WrapTextLines(t)
//      }

// -----------------------------------------------------------------------------
// # Error Handling Methods (ob *PDF)

// Clean() *PDF
func Test_PDF_Clean_(t *testing.T) { utest.Test_PDF_Clean_(t) }

// ErrorInfo(err error) (ret struct {
//     ID            int
//     Msg, Src, Val string
// })

// Errors() []error
func Test_PDF_Errors_(t *testing.T) { utest.Test_PDF_Errors_(t) }

// PullError() error
func Test_PDF_PullError_(t *testing.T) { utest.Test_PDF_PullError_(t) }

//end
