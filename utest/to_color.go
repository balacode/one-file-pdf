// -----------------------------------------------------------------------------
// (c) markaum@gmail.com                                            License: MIT
// :v: 2018-03-23 17:07:40                                   [utest/to_color.go]
// -----------------------------------------------------------------------------

package utest

import (
	"fmt"         // standard
	"image/color" // standard
	"testing"     // standard

	pdf "github.com/balacode/one-file-pdf"
)

// ToColor is the unit test for PDF.ToColor()
func ToColor(t *testing.T) {
	fmt.Println("utest.ToColor")

	testCases := []struct {
		description string
		input       string
		color       color.RGBA
		err         error
	}{
		{
			description: "valid hex",
			input:       "#c83296",
			color:       color.RGBA{200, 50, 150, 255},
		},
		{
			description: "hex with more than seven characters",
			input:       "#c83296XXXXXXX",
			color:       color.RGBA{200, 50, 150, 255},
		},
		{
			description: "invalid hex",
			input:       "#wrongcolor",
			color:       color.RGBA{A: 255},
			err:         pdf.ErrBadColorCode{Code: "#wrongcolor"},
		},

		// X is not a valid hex char. Only valid values are: 0-9 and A-F
		{
			description: "hex with an invalid character",
			input:       "#845X76",
			color:       color.RGBA{A: 255},
			err:         pdf.ErrBadColorCode{Code: "#845X76"},
		},
		{
			description: "valid color name",
			input:       "MEDIUMPURPLE",
			color:       color.RGBA{147, 112, 219, 255},
		},
		{
			description: "valid lowercase color name",
			input:       "mediumpurple",
			color:       color.RGBA{147, 112, 219, 255},
		},
		{
			description: "unknown color name",
			input:       "picasso",
			color:       color.RGBA{A: 255},
			err:         pdf.ErrUnknownColor{Color: "picasso"},
		},
	}

	for _, test := range testCases {
		var ob pdf.PDF

		t.Run(test.description, func(t *testing.T) {
			color, err := ob.ToColor(test.input)

			if err != test.err {
				t.Fatalf("expected err %v got %v", test.err, err)
			}

			if test.color != color {
				t.Fatalf("expected color %v got %v", test.color, color)
			}
		})
	}

} //                                                                      ToColor

//end
