// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-03-11 00:23:05 A2F37D                             [test/pdf_test.go]
// -----------------------------------------------------------------------------

package main

import "testing" // standard

// go test --run TestFail
func TestFail(t *testing.T) {
	t.Errorf("Fail deliberately to test Travis CI")
} //                                                                    TestFail

//end
