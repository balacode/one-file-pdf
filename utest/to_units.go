// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-04-29 23:42:24 CB363F                            [utest/to_units.go]
// -----------------------------------------------------------------------------

package utest

import (
	"fmt"
	"testing"

	"github.com/balacode/one-file-pdf"
)

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
