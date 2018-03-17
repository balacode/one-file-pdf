// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-03-17 11:57:31 23E3EC                           [utest/to_points.go]
// -----------------------------------------------------------------------------

package utest

import "fmt"     // standard
import "testing" // standard

import "github.com/balacode/one-file-pdf"

func ToPoints(t *testing.T) {
	fmt.Println("utest.ToPoints")
	//
	var test = func(
		expectVal float64, expectErr error, inputParts ...[]string,
	) {
		for _, s := range permuteStrings(inputParts...) {
			var ob pdf.PDF
			var gotVal, gotErr = ob.ToPoints(s)
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
		_, _, _, _, _, _ = cm, inches, mm, points, twips, spc
	)
	// if unit is not specified at all, there's no error, but assume it's points
	test(123, nil, spc, []string{"123"}, spc)
	//
	// test single units
	test(72, nil, spc, []string{"1"}, spc, inches, spc)   // 1 inch = 72 points
	test(2.835, nil, spc, []string{"1"}, spc, mm, spc)    // 1 mm = 2.835 points
	test(28.346, nil, spc, []string{"1"}, spc, cm, spc)   // 1 cm = 28.346 points
	test(0.050, nil, spc, []string{"1"}, spc, twips, spc) // 1 twip = 0.05 points
	test(1, nil, spc, []string{"1"}, spc, points, spc)    // 1 point = 1 point :)
	//
	// test negative number with decimals
	test(-888.840, nil, spc, []string{"-12.345"}, spc, inches, spc)
	test(-34.994, nil, spc, []string{"-12.345"}, spc, mm, spc)
	test(-349.937, nil, spc, []string{"-12.345"}, spc, cm, spc)
	test(-0.617, nil, spc, []string{"-12.345"}, spc, twips, spc)
	test(-12.345, nil, spc, []string{"-12.345"}, spc, points, spc)
	//
	test(1, nil, spc, []string{"20"}, spc, twips, spc) // 1 point = 20 twips
	test(-1, nil, spc, []string{"-20"}, spc, twips, spc)
	//
	// test some bad units
	test(0, fmt.Errorf(`Unknown unit name: "km"`), []string{"1km"})
} //                                                                    ToPoints

//end
