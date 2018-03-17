// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-03-17 12:46:38 D93961                               [utest/clean.go]
// -----------------------------------------------------------------------------

package utest

import "fmt"     // standard
import "testing" // standard

import "github.com/balacode/one-file-pdf"

func Clean(t *testing.T) {
	fmt.Println("utest.Clean")
	//
	// calling Clean() multiple times on a non-initialized PDF:
	// (you should not do this normally, use NewPDF() to create a PDF)
	// - should not panic
	// - length of Errors() should be zero
	// - Errors() should be []error{}, not nil
	func() {
		var ob pdf.PDF
		ob.Clean().Clean().Clean()
		//        result            expected
		TEqual(t, len(ob.Errors()), 0)
		TEqual(t, ob.Errors(), []error{})
	}()
	//
	// same as above for a PDF properly initialized with NewPDF()
	// (also, call Clean() without chaining)
	func() {
		var ob = pdf.NewPDF("A4")
		ob.Clean()
		ob.Clean()
		ob.Clean()
		//        result            expected
		TEqual(t, len(ob.Errors()), 0)
		TEqual(t, ob.Errors(), []error{})
	}()
	//
	// create a new PDF with an unknown page size, then call Clean()
	// first, Errors should have 1 error
	// after Clean(), Errors should be zero-length
	func() {
		var ob = pdf.NewPDF("Parchment")
		//        result            expected
		TEqual(t, len(ob.Errors()), 1)
		ob.Clean()
		//        result            expected
		TEqual(t, len(ob.Errors()), 0)
		TEqual(t, ob.Errors(), []error{})
	}()
} //                                                                       Clean

//end
