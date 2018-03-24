// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-03-24 18:56:20 438C56                               [utest/clean.go]
// -----------------------------------------------------------------------------

package utest

import "fmt"     // standard
import "testing" // standard

import "github.com/balacode/one-file-pdf"

// Clean is the unit test for PDF.Clean()
func Clean(t *testing.T) {
	fmt.Println("utest.Clean")

	// calling Clean() multiple times on a non-initialized PDF:
	// (you should not do this normally, use NewPDF() to create a PDF)
	// - should not panic
	// - length of Errors() should be zero
	// - Errors() should be []error{}, not nil
	func() {
		var doc pdf.PDF // uninitialized PDF
		doc.Clean().Clean().Clean()
		//
		//        result            expected
		TEqual(t, len(doc.Errors()), 0)
		TEqual(t, doc.Errors(), []error{})
	}()

	// same as above for a PDF properly initialized with NewPDF()
	// (also, call Clean() without chaining)
	func() {
		var doc = pdf.NewPDF("A4")
		doc.Clean()
		doc.Clean()
		doc.Clean()
		//        result            expected
		TEqual(t, len(doc.Errors()), 0)
		TEqual(t, doc.Errors(), []error{})
	}()

	// create a new PDF with an unknown page size, then call Clean()
	// first, Errors should have 1 error
	// after Clean(), Errors should be zero-length
	func() {
		var doc = pdf.NewPDF("Parchment")
		//        result             expected
		TEqual(t, len(doc.Errors()), 1)
		doc.Clean()
		//        result             expected
		TEqual(t, len(doc.Errors()), 0)
		TEqual(t, doc.Errors(), []error{})
	}()

} //                                                                       Clean

//end
