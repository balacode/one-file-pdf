// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-03-21 01:11:49 BB78F8                               [public_test.go]
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

// -----------------------------------------------------------------------------
// # Constructor

// NewPDF(paperSize string) PDF
// go test --run Test_NewPDF_
func Test_NewPDF_(t *testing.T) { utest.NewPDF(t) }

// -----------------------------------------------------------------------------
// # Read-Only Properties (pdf *PDF)

// CurrentPage() int
// go test --run Test_CurrentPage_
func Test_CurrentPage_(t *testing.T) { utest.CurrentPage(t) }

// -----------------------------------------------------------------------------
// # Properties

// Color() color.RGBA
// SetColor(nameOrHTMLColor string) *PDF
// SetColorRGB(red, green, blue int) *PDF
// go test --run Test_Color_
func Test_Color_(t *testing.T) { utest.Color(t) }

// Units() string
// SetUnits(unitName string) *PDF
// go test --run Test_Units_
func Test_Units_(t *testing.T) { utest.Units(t) }

// -----------------------------------------------------------------------------
// # Methods (pdf *PDF)

// DrawBox(x, y, width, height float64, fill ...bool) *PDF
// go test --run Test_DrawBox_
func Test_DrawBox_(t *testing.T) { utest.DrawBox(t) }

// DrawCircle(x, y, radius float64, fill ...bool) *PDF
// go test --run Test_DrawCircle_
func Test_DrawCircle_(t *testing.T) { utest.DrawCircle(t) }

// DrawImage(x, y, height float64, fileNameOrBytes interface{},
//     backColor ...string) *PDF
// go test --run Test_DrawImage_
func Test_DrawImage_(t *testing.T) { utest.DrawImage(t) }

// DrawText(s string) *PDF
// go test --run Test_DrawText_
func Test_DrawText_(t *testing.T) { utest.DrawText(t) }

// DrawTextAt(x, y float64, text string) *PDF
// go test --run Test_DrawTextAt_
func Test_DrawTextAt_(t *testing.T) { utest.DrawTextAt(t) }

// DrawTextInBox(
//     x, y, width, height float64, align, text string ) *PDF
// go test --run Test_DrawTextInBox_
func Test_DrawTextInBox_(t *testing.T) { utest.DrawTextInBox(t) }

// DrawUnitGrid() *PDF
// go test --run Test_DrawUnitGrid_
func Test_DrawUnitGrid_(t *testing.T) { utest.DrawUnitGrid(t) }

// FillBox(x, y, width, height float64) *PDF
// go test --run Test_FillBox_
func Test_FillBox_(t *testing.T) { utest.FillBox(t) }

// FillCircle(x, y, radius float64) *PDF
// go test --run Test_FillCircle_
func Test_FillCircle_(t *testing.T) { utest.FillCircle(t) }

// -----------------------------------------------------------------------------
// # Metrics Methods (pdf *PDF)

// ToPoints(numberAndUnit string) (float64, error)
// go test --run Test_ToPoints_
func Test_ToPoints_(t *testing.T) { utest.ToPoints(t) }

// -----------------------------------------------------------------------------
// # Error Handling Methods (pdf *PDF)

// Clean() *PDF
// go test --run Test_Clean_
func Test_Clean_(t *testing.T) { utest.Clean(t) }

// Errors() []error
// go test --run Test_Errors_
func Test_Errors_(t *testing.T) { utest.Errors(t) }

//end
