// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-04-29 23:42:24 53F3F1                               [utest/units.go]
// -----------------------------------------------------------------------------

package utest

import (
	"fmt"
	"testing"

	"github.com/balacode/one-file-pdf"
)

// Test_PDF_Units_ tests PDF.Units() and SetUnits()
func Test_PDF_Units_(t *testing.T) {

	// (ob *PDF) Units() string
	//
	fmt.Println("Test PDF.Units()")
	func() {
		var doc pdf.PDF // uninitialized PDF
		TEqual(t, doc.Units(), "POINT")
	}()

	func() {
		var doc = pdf.NewPDF("A4")
		TEqual(t, doc.Units(), "POINT")
	}()

	// (ob *PDF) SetUnits(units string) *PDF
	//
	fmt.Println("Test PDF.SetUnits()")

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
			//TODO: rename 'unit name' in message
		}
		TEqual(t, doc.Units(), "POINT")
	}()

} //                                                             Test_PDF_Units_

//end
