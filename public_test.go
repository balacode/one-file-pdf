// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-03-19 23:43:04 144294                               [public_test.go]
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
// # Methods (pdf *PDF)

// DrawImage(x, y, height float64, fileNameOrBytes interface{},
//     backColor ...string) *PDF
// go test --run Test_DrawImage_
func Test_DrawImage_(t *testing.T) { utest.DrawImage(t) }

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
