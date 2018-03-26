// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-03-26 12:23:48 B5E099                              [utest/errors.go]
// -----------------------------------------------------------------------------

package utest

import "fmt"     // standard
import "testing" // standard

import "github.com/balacode/one-file-pdf"

// Errors tests PDF.Errors()
func Errors(t *testing.T) {
	fmt.Println("utest.Errors")

	// Errors() should be []error{} on a non-initialized PDF:
	func() {
		var doc pdf.PDF // uninitialized PDF
		//
		//        result            expected
		TEqual(t, len(doc.Errors()), 0)
		TEqual(t, doc.Errors(), []error{})
	}()

	// same as above for a PDF properly initialized with NewPDF()
	// (also, call Errors() without chaining)
	func() {
		var doc = pdf.NewPDF("A4")
		TEqual(t, len(doc.Errors()), 0)
		TEqual(t, doc.Errors(), []error{})
	}()

} //                                                                      Errors

//end
