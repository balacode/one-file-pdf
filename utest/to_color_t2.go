// -----------------------------------------------------------------------------
// (c) markaum@gmail.com                                            License: MIT
// :v: 2018-03-26 12:23:48 002F1A                         [utest/to_color_t2.go]
// -----------------------------------------------------------------------------

package utest

import (
	"fmt"         // standard
	"image/color" // standard
	"testing"     // standard

	pdf "github.com/balacode/one-file-pdf"
)

// ToColorT2 is the second unit test for PDF.ToColor()
func ToColorT2(t *testing.T) {
	fmt.Println("utest.ToColorT2")
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
			err:         pdf.TMPErrBadColorCode{Code: "#wrongcolor"},
		},
		// X is not a valid hex char. Only valid values are: 0-9 and A-F
		{
			description: "hex with an invalid character",
			input:       "#845X76",
			color:       color.RGBA{A: 255},
			err:         pdf.TMPErrBadColorCode{Code: "#845X76"},
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
			err:         pdf.TMPErrUnknownColor{Color: "picasso"},
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
} //                                                                   ToColorT2

//end
