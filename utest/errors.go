// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-03-22 02:48:21 6EAA73                              [utest/errors.go]
// -----------------------------------------------------------------------------

package utest

import "fmt"     // standard
import "testing" // standard

import "github.com/balacode/one-file-pdf"

func Errors(t *testing.T) {
	fmt.Println("utest.Errors")

	// Errors() should be []error{} on a non-initialized PDF:
	func() {
		var ob pdf.PDF // uninitialized PDF
		//
		//        result            expected
		TEqual(t, len(ob.Errors()), 0)
		TEqual(t, ob.Errors(), []error{})
	}()

	// same as above for a PDF properly initialized with NewPDF()
	// (also, call Errors() without chaining)
	func() {
		var ob = pdf.NewPDF("A4")
		TEqual(t, len(ob.Errors()), 0)
		TEqual(t, ob.Errors(), []error{})
	}()

} //                                                                      Errors

//end
