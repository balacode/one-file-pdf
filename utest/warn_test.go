// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-03-12 22:27:16 83825C                           [utest/warn_test.go]
// -----------------------------------------------------------------------------

package utest

import "fmt"     // standard
import "testing" // standard

// TestWarning gets displayed if someone tries to run 'go test' from 'utest'.
func TestWarning(t *testing.T) {
	fmt.Printf(`
		------------------------------------------------------------
		PLEASE NOTE:
		Do not run 'go test' from package 'one-file-pdf/utest'.
		Instead, run 'go test' from its parent folder 'one-file-pdf'
		which makes use of 'utest' (see notes in run_test.go)
		------------------------------------------------------------

	`)
}

//end
