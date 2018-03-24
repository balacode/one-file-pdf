// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-03-24 18:56:20 931001                            [utest/to_units.go]
// -----------------------------------------------------------------------------

package utest

import "fmt"     // standard
import "testing" // standard

import "github.com/balacode/one-file-pdf"

// ToUnits is the unit test for
// ToUnits(points float64) float64
func ToUnits(t *testing.T) {
	fmt.Println("utest.ToUnits")

	func() {
		var doc pdf.PDF
		TEqual(t, doc.ToUnits(1), 1)
		//TODO: add test cases for ToUnits()
	}()

} //                                                                     ToUnits

//end
