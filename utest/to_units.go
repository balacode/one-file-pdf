// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-05-03 23:17:35 3C9F4C               one-file-pdf/utest/[to_units.go]
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
		//
		doc.SetUnits("cm")
		TEqual(t, doc.ToUnits(1), 0.035278)         // 1 point = 0.035278 cm
		TEqual(t, doc.ToUnits(28.3464566929134), 1) // ~28.3 points = 1 cm
		//
		doc.SetUnits("in")
		TEqual(t, doc.ToUnits(1), 0.0138888888888889) // 1 point = ~0.0138 in.
		TEqual(t, doc.ToUnits(72), 1)                 // 72 points = 1 inch
		//
		doc.SetUnits("mm")
		TEqual(t, doc.ToUnits(1), 0.3527777777777776) // 1 point = ~0.3527 mm
		TEqual(t, doc.ToUnits(2.83464566929134), 1)   // ~2.8346 points = 1 mm
		//
		doc.SetUnits("point")
		TEqual(t, doc.ToUnits(1), 1) // 1 point = 1 point
		//
		doc.SetUnits("twip")
		TEqual(t, doc.ToUnits(1), 20)   // 1 point = 20 twips
		TEqual(t, doc.ToUnits(0.05), 1) // 0.05 points = 1 twip
	}()

} //                                                           Test_PDF_ToUnits_

//end
