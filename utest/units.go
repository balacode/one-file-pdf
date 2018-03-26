// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-03-26 12:23:48 8C3256                               [utest/units.go]
// -----------------------------------------------------------------------------

package utest

import "fmt"     // standard
import "testing" // standard

import "github.com/balacode/one-file-pdf"

// Units tests PDF.Units() and SetUnits()
func Units(t *testing.T) {
	fmt.Println("utest.Units")

	// (pdf *PDF) Units() string
	//
	func() {
		var doc pdf.PDF // uninitialized PDF
		TEqual(t, doc.Units(), "POINT")
	}()

	func() {
		var doc = pdf.NewPDF("A4")
		TEqual(t, doc.Units(), "POINT")
	}()

	// (pdf *PDF) SetUnits(unitName string) *PDF
	//
	fmt.Println("utest.SetUnits")

	func() {
		var doc = pdf.NewPDF("A4")
		TEqual(t, len(doc.Errors()), 0)
		doc.SetUnits("cm")
		TEqual(t, len(doc.Errors()), 0)
		TEqual(t, doc.Units(), "CM")
	}()

	func() {
		var doc = pdf.NewPDF("A4")
		TEqual(t, len(doc.Errors()), 0)
		doc.SetUnits("fathoms")
		TEqual(t, len(doc.Errors()), 1)
		//
		if len(doc.Errors()) == 1 {
			TEqual(t,
				doc.Errors()[0],
				fmt.Errorf(`Unknown unit name "fathoms" @SetUnits`))
		}
		TEqual(t, doc.Units(), "POINT")
	}()

} //                                                                       Units

//end
