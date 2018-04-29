// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-04-29 23:42:24 D142EE                           [utest/warn_test.go]
// -----------------------------------------------------------------------------

package utest

import (
	"fmt"
	"testing"
)

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
