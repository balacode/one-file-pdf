// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-04-29 23:43:39 220C7B                               [utest/color.go]
// -----------------------------------------------------------------------------

package utest

import (
	"fmt"
	"image/color"
	"testing"

	"github.com/balacode/one-file-pdf"
)

// Test_PDF_Color_ tests PDF.Color() and SetColor()
func Test_PDF_Color_(t *testing.T) {

	// -------------------------------------------------------------------------
	// (ob *PDF) Color() color.RGBA
	//
	fmt.Println("Test PDF.Color()")

	func() {
		var doc pdf.PDF // uninitialized PDF
		TEqual(t, doc.Color(), color.RGBA{A: 255})
	}()

	func() {
		var doc = pdf.NewPDF("A4") // initialized PDF
		TEqual(t, doc.Color(), color.RGBA{A: 255})
	}()

	// -------------------------------------------------------------------------
	// (ob *PDF) SetColor(nameOrHTMLColor string) *PDF
	//
	// TODO: test various named colors
	// TODO: test setting HTML color codes
	// TODO: test if alpha is always 255
	// TODO: test name trimming
	// TODO: test case-insensitivity
	// TODO: test ignoring '-' and '_'

	fmt.Println("Test PDF.SetColor()")
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
		var doc pdf.PDF // uninitialized PDF
		TEqual(t, len(doc.Errors()), 0)
		doc.SetColor("")
		TEqual(t, len(doc.Errors()), 1)
		//
		if len(doc.Errors()) == 1 {
			TEqual(t,
				doc.Errors()[0],
				fmt.Errorf(`Unknown color name "" @SetColor`))
		}
		TEqual(t, doc.Color(), color.RGBA{A: 255})
	}()

	// try setting an unknown color name
	func() {
		var doc pdf.PDF // uninitialized PDF
		TEqual(t, len(doc.Errors()), 0)
		doc.SetColor("TheColourOutOfSpace")
		TEqual(t, len(doc.Errors()), 1)
		//
		if len(doc.Errors()) == 1 {
			TEqual(t,
				doc.Errors()[0],
				fmt.Errorf(
					`Unknown color name "TheColourOutOfSpace" @SetColor`))
		}
		TEqual(t, doc.Color(), color.RGBA{A: 255})
	}()

	// -------------------------------------------------------------------------
	// SetColorRGB(red, green, blue int) *PDF
	//
	fmt.Println("Test PDF.SetColorRGB()")

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

} //                                                             Test_PDF_Color_

//end
