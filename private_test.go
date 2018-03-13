// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-03-13 23:41:43 D21FD0                              [private_test.go]
// -----------------------------------------------------------------------------

package pdf

import "fmt"     // standard
import "testing" // standard

// go test --run Test_getPapreSize_
func Test_getPapreSize_(t *testing.T) {
	//
	var test = func(input, name string, widthPt, heightPt float64, err error) {
		var pdf PDF
		var got, gotErr = pdf.getPaperSize(name)
		//
		if gotErr != err {
			t.Errorf("'error' mismatch: expected: %v returned %v", err, gotErr)
			t.Fail()
		}
		if got.name != name {
			mismatch(t, input+" 'name'", name, got.name)
		}
		if floatStr(got.widthPt) != floatStr(widthPt) {
			mismatch(t, input+" 'widthPt'", widthPt, got.widthPt)
		}
		if floatStr(got.heightPt) != floatStr(heightPt) {
			mismatch(t, input+" 'heightPt'", heightPt, got.heightPt)
		}
	}
	test("A4", "A4", 595.276, 841.89, nil)
	test("A4-L", "A4-L", 841.89, 595.276, nil)
	// TODO: add more test cases
} //                                                          Test_getPapreSize_

// -----------------------------------------------------------------------------
// # Helper Functions

// floatStr returns a float64 as a string, with val rounded to 3 decimals
func floatStr(val float64) string {
	return fmt.Sprintf("%0.3f", val)
} //                                                                    floatStr

// mismatch formats and raises a test error
func mismatch(t *testing.T, tag string, expected, got interface{}) {
	var expStr = fmt.Sprintf("%v", expected)
	var gotStr = fmt.Sprintf("%v", got)
	t.Errorf("%s mismatch: expected: %s got: %s", tag, expStr, gotStr)
	t.Fail()
} //                                                                    mismatch

//end
