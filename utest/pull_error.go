// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-03-24 18:56:20 C8FE22                          [utest/pull_error.go]
// -----------------------------------------------------------------------------

package utest

import "fmt"     // standard
import "testing" // standard

import "github.com/balacode/one-file-pdf"

// PullError is the unit test for PullError() error
func PullError(t *testing.T) {

	fmt.Println("utest.PullError")

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

} //                                                                   PullError

//end
