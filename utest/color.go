// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-03-21 01:14:03 055364                               [utest/color.go]
// -----------------------------------------------------------------------------

package utest

import "fmt"         // standard
import "testing"     // standard
import "image/color" // standard

import "github.com/balacode/one-file-pdf"

func Color(t *testing.T) {

	// -------------------------------------------------------------------------
	// Color() color.RGBA

	fmt.Println("utest.Color")
	func() {
		var ob pdf.PDF
		TEqual(t, ob.Color(), color.RGBA{})
	}()
	// test property getters of initialized PDF
	func() {
		var ob = pdf.NewPDF("A4")
		TEqual(t, ob.Color(), color.RGBA{A: 255})
	}()

	// -------------------------------------------------------------------------
	// SetColor(nameOrHTMLColor string) *PDF

	// TODO: test various named colors
	// TODO: test setting HTML color codes
	// TODO: test if alpha is always 255
	// TODO: test name trimming
	// TODO: test case-insensitivity
	// TODO: test ignoring '-' and '_'

	fmt.Println("utest.SetColor")
	for _, name := range permuteStrings(
		[]string{"", " ", "  "},
		[]string{"red", "Red", "RED"},
		[]string{"", " ", "  "},
	) {
		func() {
			var a = pdf.NewPDF("A4")
			var b = a.SetColor(name)
			TEqual(t, a.Color(), color.RGBA{R: 255, A: 255})
			TEqual(t, &a, b)
		}()
	}
	// try setting a blank color name
	func() {
		var ob pdf.PDF
		TEqual(t, len(ob.Errors()), 0)
		ob.SetColor("")
		TEqual(t, len(ob.Errors()), 1)
		//
		if len(ob.Errors()) == 1 {
			TEqual(t,
				ob.Errors()[0],
				fmt.Errorf(`Unknown color name ""`))
		}
		TEqual(t, ob.Color(), color.RGBA{A: 255})
	}()
	// try setting an unknown color name
	func() {
		var ob pdf.PDF
		TEqual(t, len(ob.Errors()), 0)
		ob.SetColor("TheColourOutOfSpace")
		TEqual(t, len(ob.Errors()), 1)
		//
		if len(ob.Errors()) == 1 {
			TEqual(t,
				ob.Errors()[0],
				fmt.Errorf(`Unknown color name "TheColourOutOfSpace"`))
		}
		TEqual(t, ob.Color(), color.RGBA{A: 255})
	}()

	// -------------------------------------------------------------------------
	// SetColorRGB(red, green, blue int) *PDF

	fmt.Println("utest.SetColorRGB")
	func() {
		// red
		var a = pdf.NewPDF("A4")
		var b = a.SetColorRGB(128, 0, 0)
		TEqual(t, a.Color(), color.RGBA{R: 128, A: 255})
		TEqual(t, &a, b)
	}()
	func() {
		// green
		var a = pdf.NewPDF("A4")
		var b = a.SetColorRGB(0, 128, 0)
		TEqual(t, a.Color(), color.RGBA{G: 128, A: 255})
		TEqual(t, &a, b)
	}()
	func() {
		// blue
		var a = pdf.NewPDF("A4")
		var b = a.SetColorRGB(0, 0, 128)
		TEqual(t, a.Color(), color.RGBA{B: 128, A: 255})
		TEqual(t, &a, b)
	}()
} //                                                                       Color

//end
