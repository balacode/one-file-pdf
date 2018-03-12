// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-03-12 14:25:34 23D9D0                       [test/to_points_test.go]
// -----------------------------------------------------------------------------

package main

/*
to generate a test coverage report use:
	go test -coverprofile cover.out
	go tool cover -html=cover.out
*/

import "fmt"     // standard
import "testing" // standard

import "github.com/balacode/one-file-pdf"

// go test --run Test_ToPoints_
func Test_ToPoints_(t *testing.T) {
	//
	// test function
	var test = func(expect float64, parts ...[]string) {
		for _, s := range permuteStrings(parts...) {
			var ob pdf.PDF
			TEqual(t,
				fmt.Sprintf("%0.03f", ob.ToPoints(s)),
				fmt.Sprintf("%0.03f", expect),
			)
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
		gaps = []string{
			"", " ", "  ", "\t",
		}
		_, _, _, _, _, _ = cm, inches, mm, points, twips, gaps
	)
	// test single units
	test(72, gaps, []string{"1"}, gaps, inches, gaps)   // 1 inch = 72 points
	test(2.835, gaps, []string{"1"}, gaps, mm, gaps)    // 1 mm = 2.835 points
	test(28.346, gaps, []string{"1"}, gaps, cm, gaps)   // 1 cm = 28.346 points
	test(0.050, gaps, []string{"1"}, gaps, twips, gaps) // 1 twip = 0.05 points
	test(1, gaps, []string{"1"}, gaps, points, gaps)    // 1 point = 1 point
	//
	// test negative number with decimals
	test(-888.840, gaps, []string{"-12.345"}, gaps, inches, gaps)
	test(-34.994, gaps, []string{"-12.345"}, gaps, mm, gaps)
	test(-349.937, gaps, []string{"-12.345"}, gaps, cm, gaps)
	test(-0.617, gaps, []string{"-12.345"}, gaps, twips, gaps)
	test(-12.345, gaps, []string{"-12.345"}, gaps, points, gaps)
	//
	test(1, gaps, []string{"20"}, gaps, twips, gaps) // 1 point = 20 twips
	test(-1, gaps, []string{"-20"}, gaps, twips, gaps)
} //                                                              Test_ToPoints_

//end
