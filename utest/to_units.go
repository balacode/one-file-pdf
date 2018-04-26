// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-04-26 22:42:44 37AD48                            [utest/to_units.go]
// -----------------------------------------------------------------------------

package utest

import "fmt"     // standard
import "testing" // standard

import "github.com/balacode/one-file-pdf"

// Test_PDF_ToUnits_ is the unit test for
// ToUnits(points float64) float64
func Test_PDF_ToUnits_(t *testing.T) {
	fmt.Println("Test PDF.ToUnits()")

	func() {
		var doc pdf.PDF
		TEqual(t, doc.ToUnits(1), 1)
		//TODO: add test cases for ToUnits()
	}()

} //                                                           Test_PDF_ToUnits_

//end
