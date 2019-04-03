// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2019-04-03 10:03:38 C067A2             one-file-pdf/utest/[pull_error.go]
// -----------------------------------------------------------------------------

package utest

import (
	"fmt"
	"testing"

	"github.com/balacode/one-file-pdf"
)

// Test_PDF_PullError_ is the unit test for PullError() error
func Test_PDF_PullError_(t *testing.T) {
	fmt.Println("Test PDF.PullError()")

	func() {
		doc := pdf.NewPDF("Papyrus")
		//
		// errors should have one error
		TEqual(t, len(doc.Errors()), 1)
		//
		// fetch and remove this error from Errors()
		err := doc.PullError()
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
