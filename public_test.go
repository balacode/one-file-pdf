// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-03-18 01:39:08 AB118A                               [public_test.go]
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

// go test --run Test_Clean_
func Test_Clean_(t *testing.T) { utest.Clean(t) }

// go test --run Test_CurrentPage_
func Test_CurrentPage_(t *testing.T) { utest.CurrentPage(t) }

// go test --run Test_DrawUnitGrid_
func Test_DrawUnitGrid_(t *testing.T) { utest.DrawUnitGrid(t) }

// go test --run Test_Errors_
func Test_Errors_(t *testing.T) { utest.Errors(t) }

// go test --run Test_ToPoints_
func Test_ToPoints_(t *testing.T) { utest.ToPoints(t) }

// go test --run Test_NewPDF_
func Test_NewPDF_(t *testing.T) { utest.NewPDF(t) }

//end
