// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2019-04-03 10:27:56 91D784                 one-file-pdf/utest/[errors.go]
// -----------------------------------------------------------------------------

package utest

import (
	"fmt"
	"testing"

	"github.com/balacode/one-file-pdf"
)

// Test_PDF_Errors_ tests PDF.Errors()
func Test_PDF_Errors_(t *testing.T) {
	fmt.Println("Test PDF.Errors()")

	// Errors() should be []error{} on a non-initialized PDF:
	func() {
		var doc pdf.PDF // uninitialized PDF
		//
		//        got                want
		TEqual(t, len(doc.Errors()), 0)
		TEqual(t, doc.Errors(), []error{})
	}()

	// same as above for a PDF properly initialized with NewPDF()
	// (also, call Errors() without chaining)
	func() {
		doc := pdf.NewPDF("A4")
		TEqual(t, len(doc.Errors()), 0)
		TEqual(t, doc.Errors(), []error{})
	}()

} //                                                            Test_PDF_Errors_

//end
