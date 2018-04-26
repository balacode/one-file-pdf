// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-04-26 22:42:44 9BB7CB                          [utest/pull_error.go]
// -----------------------------------------------------------------------------

package utest

import "fmt"     // standard
import "testing" // standard

import "github.com/balacode/one-file-pdf"

// Test_PDF_PullError_ is the unit test for PullError() error
func Test_PDF_PullError_(t *testing.T) {
	fmt.Println("Test PDF.PullError()")

	func() {
		var doc = pdf.NewPDF("Papyrus")
		//
		// errors should have one error
		TEqual(t, len(doc.Errors()), 1)
		//
		// fetch and remove this error from Errors()
		var err = doc.PullError()
		TEqual(t, err, fmt.Errorf(`Unknown paper size "Papyrus" @NewPDF`))
		//
		// Errors() should now be empty
		TEqual(t, len(doc.Errors()), 0)
		//
		// if we try to pull another error (there is none), 'err' will be nil
		err = doc.PullError()
		TEqual(t, len(doc.Errors()), 0)
		TEqual(t, err, nil)
	}()

} //                                                         Test_PDF_PullError_

//end
