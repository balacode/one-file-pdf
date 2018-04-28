// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-04-28 22:37:48 31C825                           [utest/to_points.go]
// -----------------------------------------------------------------------------

package utest

import "fmt"     // standard
import "testing" // standard

import "github.com/balacode/one-file-pdf"

// Test_PDF_ToPoints_ is the unit test for PDF.ToPoints()
func Test_PDF_ToPoints_(t *testing.T) {
	fmt.Println("Test PDF.ToPoints()")
	//
	var test = func(
		expectVal float64, expectErr error, inputParts ...[]string,
	) {
		for _, s := range permuteStrings(inputParts...) {
			var doc pdf.PDF // uninitialized PDF
			var gotVal, gotErr = doc.ToPoints(s)
			TEqual(t,
				fmt.Sprintf("%0.03f", gotVal),
				fmt.Sprintf("%0.03f", expectVal),
			)
			TEqual(t, gotErr, expectErr)
		}
	}
	var (
		cm     = []string{"CM", "Cm", "cM", "cm"}
		inches = []string{
			"IN", "INCH", "INCHES",
			"In", "Inch", "Inches",
			"in", "inch", "inches",
			`"`,
		}
		mm     = []string{"MM", "mm", "Mm", "mM"}
		points = []string{
			"PT", "POINT", "POINTS",
			"Pt", "Point", "Points",
			"pt", "point", "points",
		}
		twips = []string{
			"TW", "TWIP", "TWIPS",
			"Tw", "Twip", "Twips",
			"tw", "twip", "twips",
		}
		spc = []string{ // various spaces
			"", " ", "  ", "\t",
		}
	)
	// if unit is not specified at all, there's no error, but assume it's points
	test(123, nil, spc, []string{"123"}, spc)
	//
	// test single units
	var one = []string{"1"}
	test(72, nil, spc, one, spc, inches, spc)   // 1 inch = 72 points
	test(2.835, nil, spc, one, spc, mm, spc)    // 1 mm = 2.835 points
	test(28.346, nil, spc, one, spc, cm, spc)   // 1 cm = 28.346 points
	test(0.050, nil, spc, one, spc, twips, spc) // 1 twip = 0.05 points
	test(1, nil, spc, one, spc, points, spc)    // 1 point = 1 point :)
	//
	// test negative number with decimals
	var negative = []string{"-12.345"}
	test(-888.840, nil, spc, negative, spc, inches, spc)
	test(-34.994, nil, spc, negative, spc, mm, spc)
	test(-349.937, nil, spc, negative, spc, cm, spc)
	test(-0.617, nil, spc, negative, spc, twips, spc)
	test(-12.345, nil, spc, negative, spc, points, spc)
	//
	test(1, nil, spc, []string{"20"}, spc, twips, spc) // 1 point = 20 twips
	test(-1, nil, spc, []string{"-20"}, spc, twips, spc)
	//
	// test some bad units
	test(0, fmt.Errorf(`Unknown unit name "km"`), []string{"1km"})
	//TODO: rename 'unit name' in message
	test(0, fmt.Errorf(`Invalid number "1.0.1"`), []string{"1.0.1mm"})
} //                                                          Test_PDF_ToPoints_

//end
