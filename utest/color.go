// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-05-28 13:59:12 FC09D6                               [utest/color.go]
// -----------------------------------------------------------------------------

package utest

import (
	"fmt"
	"image/color"
	str "strings"
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
	fmt.Println("Test PDF.SetColor()")

	// test various named colors and codes
	for _, iter := range []struct {
		in     string
		expect color.RGBA
	}{
		// test HTML color codes
		{"#000000", color.RGBA{R: 0, G: 0, B: 0, A: 255}},
		{"#010203", color.RGBA{R: 1, G: 2, B: 3, A: 255}},
		{"#FFFFFF", color.RGBA{R: 255, G: 255, B: 255, A: 255}},
		//
		// test all named colors
		{"ALICE BLUE", color.RGBA{R: 240, G: 248, B: 255, A: 255}},
		{"ANTIQUE WHITE", color.RGBA{R: 250, G: 235, B: 215, A: 255}},
		{"AQUA", color.RGBA{R: 000, G: 255, B: 255, A: 255}},
		{"AQUAMARINE", color.RGBA{R: 127, G: 255, B: 212, A: 255}},
		{"AZURE", color.RGBA{R: 240, G: 255, B: 255, A: 255}},
		{"BEIGE", color.RGBA{R: 245, G: 245, B: 220, A: 255}},
		{"BISQUE", color.RGBA{R: 255, G: 228, B: 196, A: 255}},
		{"BLACK", color.RGBA{R: 000, G: 000, B: 000, A: 255}},
		{"BLANCHED ALMOND", color.RGBA{R: 255, G: 235, B: 205, A: 255}},
		{"BLUE", color.RGBA{R: 000, G: 000, B: 255, A: 255}},
		{"BLUE VIOLET", color.RGBA{R: 138, G: 43, B: 226, A: 255}},
		{"BROWN", color.RGBA{R: 165, G: 42, B: 42, A: 255}},
		{"BURLYWOOD", color.RGBA{R: 222, G: 184, B: 135, A: 255}},
		{"CADET BLUE", color.RGBA{R: 95, G: 158, B: 160, A: 255}},
		{"CHARTREUSE", color.RGBA{R: 127, G: 255, B: 000, A: 255}},
		{"CHOCOLATE", color.RGBA{R: 210, G: 105, B: 30, A: 255}},
		{"CORAL", color.RGBA{R: 255, G: 127, B: 80, A: 255}},
		{"CORNFLOWER BLUE", color.RGBA{R: 100, G: 149, B: 237, A: 255}},
		{"CORNSILK", color.RGBA{R: 255, G: 248, B: 220, A: 255}},
		{"CRIMSON", color.RGBA{R: 220, G: 20, B: 60, A: 255}},
		{"CYAN", color.RGBA{R: 000, G: 255, B: 255, A: 255}},
		{"DARK BLUE", color.RGBA{R: 000, G: 000, B: 139, A: 255}},
		{"DARK CYAN", color.RGBA{R: 000, G: 139, B: 139, A: 255}},
		{"DARK GOLDEN ROD", color.RGBA{R: 184, G: 134, B: 11, A: 255}},
		{"DARK GRAY", color.RGBA{R: 169, G: 169, B: 169, A: 255}},
		{"DARK GREEN", color.RGBA{R: 000, G: 100, B: 000, A: 255}},
		{"DARK KHAKI", color.RGBA{R: 189, G: 183, B: 107, A: 255}},
		{"DARK MAGENTA", color.RGBA{R: 139, G: 000, B: 139, A: 255}},
		{"DARK OLIVE GREEN", color.RGBA{R: 85, G: 107, B: 47, A: 255}},
		{"DARK ORANGE", color.RGBA{R: 255, G: 140, B: 000, A: 255}},
		{"DARK ORCHID", color.RGBA{R: 153, G: 50, B: 204, A: 255}},
		{"DARK RED", color.RGBA{R: 139, G: 000, B: 000, A: 255}},
		{"DARK SALMON", color.RGBA{R: 233, G: 150, B: 122, A: 255}},
		{"DARK SEA GREEN", color.RGBA{R: 143, G: 188, B: 143, A: 255}},
		{"DARK SLATE BLUE", color.RGBA{R: 72, G: 61, B: 139, A: 255}},
		{"DARK SLATE GRAY", color.RGBA{R: 47, G: 79, B: 79, A: 255}},
		{"DARK TURQUOISE", color.RGBA{R: 000, G: 206, B: 209, A: 255}},
		{"DARK VIOLET", color.RGBA{R: 148, G: 000, B: 211, A: 255}},
		{"DEEP PINK", color.RGBA{R: 255, G: 20, B: 147, A: 255}},
		{"DEEP SKY BLUE", color.RGBA{R: 000, G: 191, B: 255, A: 255}},
		{"DIM GRAY", color.RGBA{R: 105, G: 105, B: 105, A: 255}},
		{"DODGER BLUE", color.RGBA{R: 30, G: 144, B: 255, A: 255}},
		{"FIRE BRICK", color.RGBA{R: 178, G: 34, B: 34, A: 255}},
		{"FLORAL WHITE", color.RGBA{R: 255, G: 250, B: 240, A: 255}},
		{"FOREST GREEN", color.RGBA{R: 34, G: 139, B: 34, A: 255}},
		{"FUCHSIA", color.RGBA{R: 255, G: 000, B: 255, A: 255}},
		{"GAINSBORO", color.RGBA{R: 220, G: 220, B: 220, A: 255}},
		{"GHOST WHITE", color.RGBA{R: 248, G: 248, B: 255, A: 255}},
		{"GOLD", color.RGBA{R: 255, G: 215, B: 000, A: 255}},
		{"GOLDEN ROD", color.RGBA{R: 218, G: 165, B: 32, A: 255}},
		{"GRAY", color.RGBA{R: 190, G: 190, B: 190, A: 255}},
		{"GREEN", color.RGBA{R: 000, G: 255, B: 000, A: 255}},
		{"GREEN YELLOW", color.RGBA{R: 173, G: 255, B: 47, A: 255}},
		{"HONEY DEW", color.RGBA{R: 240, G: 255, B: 240, A: 255}},
		{"HOT PINK", color.RGBA{R: 255, G: 105, B: 180, A: 255}},
		{"INDIAN RED", color.RGBA{R: 205, G: 92, B: 92, A: 255}},
		{"INDIGO", color.RGBA{R: 75, G: 000, B: 130, A: 255}},
		{"IVORY", color.RGBA{R: 255, G: 255, B: 240, A: 255}},
		{"KHAKI", color.RGBA{R: 240, G: 230, B: 140, A: 255}},
		{"LAVENDER", color.RGBA{R: 230, G: 230, B: 250, A: 255}},
		{"LAVENDER BLUSH", color.RGBA{R: 255, G: 240, B: 245, A: 255}},
		{"LAWN GREEN", color.RGBA{R: 124, G: 252, B: 000, A: 255}},
		{"LEMON CHIFFON", color.RGBA{R: 255, G: 250, B: 205, A: 255}},
		{"LIGHT BLUE", color.RGBA{R: 173, G: 216, B: 230, A: 255}},
		{"LIGHT CORAL", color.RGBA{R: 240, G: 128, B: 128, A: 255}},
		{"LIGHT CYAN", color.RGBA{R: 224, G: 255, B: 255, A: 255}},
		{"LIGHT GOLDENROD YELLOW", color.RGBA{R: 250, G: 250, B: 210, A: 255}},
		{"LIGHT GRAY", color.RGBA{R: 211, G: 211, B: 211, A: 255}},
		{"LIGHT GREEN", color.RGBA{R: 144, G: 238, B: 144, A: 255}},
		{"LIGHT PINK", color.RGBA{R: 255, G: 182, B: 193, A: 255}},
		{"LIGHT SALMON", color.RGBA{R: 255, G: 160, B: 122, A: 255}},
		{"LIGHT SEA GREEN", color.RGBA{R: 32, G: 178, B: 170, A: 255}},
		{"LIGHT SKY BLUE", color.RGBA{R: 135, G: 206, B: 250, A: 255}},
		{"LIGHT SLATE GRAY", color.RGBA{R: 119, G: 136, B: 153, A: 255}},
		{"LIGHT STEEL BLUE", color.RGBA{R: 176, G: 196, B: 222, A: 255}},
		{"LIGHT YELLOW", color.RGBA{R: 255, G: 255, B: 224, A: 255}},
		{"LIME", color.RGBA{R: 000, G: 255, B: 000, A: 255}},
		{"LIME GREEN", color.RGBA{R: 50, G: 205, B: 50, A: 255}},
		{"LINEN", color.RGBA{R: 250, G: 240, B: 230, A: 255}},
		{"MAGENTA", color.RGBA{R: 255, G: 000, B: 255, A: 255}},
		{"MAROON", color.RGBA{R: 176, G: 48, B: 96, A: 255}},
		{"MEDIUM AQUAMARINE", color.RGBA{R: 102, G: 205, B: 170, A: 255}},
		{"MEDIUM BLUE", color.RGBA{R: 000, G: 000, B: 205, A: 255}},
		{"MEDIUM ORCHID", color.RGBA{R: 186, G: 85, B: 211, A: 255}},
		{"MEDIUM PURPLE", color.RGBA{R: 147, G: 112, B: 219, A: 255}},
		{"MEDIUM SEA GREEN", color.RGBA{R: 60, G: 179, B: 113, A: 255}},
		{"MEDIUM SLATE BLUE", color.RGBA{R: 123, G: 104, B: 238, A: 255}},
		{"MEDIUM SPRING GREEN", color.RGBA{R: 000, G: 250, B: 154, A: 255}},
		{"MEDIUM TURQUOISE", color.RGBA{R: 72, G: 209, B: 204, A: 255}},
		{"MEDIUM VIOLET RED", color.RGBA{R: 199, G: 21, B: 133, A: 255}},
		{"MIDNIGHT BLUE", color.RGBA{R: 25, G: 25, B: 112, A: 255}},
		{"MINT CREAM", color.RGBA{R: 245, G: 255, B: 250, A: 255}},
		{"MISTY ROSE", color.RGBA{R: 255, G: 228, B: 225, A: 255}},
		{"MOCCASIN", color.RGBA{R: 255, G: 228, B: 181, A: 255}},
		{"NAVAJO WHITE", color.RGBA{R: 255, G: 222, B: 173, A: 255}},
		{"NAVY", color.RGBA{R: 000, G: 000, B: 128, A: 255}},
		{"OLD LACE", color.RGBA{R: 253, G: 245, B: 230, A: 255}},
		{"OLIVE", color.RGBA{R: 128, G: 128, B: 000, A: 255}},
		{"OLIVE DRAB", color.RGBA{R: 107, G: 142, B: 35, A: 255}},
		{"ORANGE", color.RGBA{R: 255, G: 165, B: 000, A: 255}},
		{"ORANGE RED", color.RGBA{R: 255, G: 69, B: 000, A: 255}},
		{"ORCHID", color.RGBA{R: 218, G: 112, B: 214, A: 255}},
		{"PALE GOLDEN ROD", color.RGBA{R: 238, G: 232, B: 170, A: 255}},
		{"PALE GREEN", color.RGBA{R: 152, G: 251, B: 152, A: 255}},
		{"PALE TURQUOISE", color.RGBA{R: 175, G: 238, B: 238, A: 255}},
		{"PALE VIOLET RED", color.RGBA{R: 219, G: 112, B: 147, A: 255}},
		{"PAPAYA WHIP", color.RGBA{R: 255, G: 239, B: 213, A: 255}},
		{"PEACH PUFF", color.RGBA{R: 255, G: 218, B: 185, A: 255}},
		{"PERU", color.RGBA{R: 205, G: 133, B: 63, A: 255}},
		{"PINK", color.RGBA{R: 255, G: 192, B: 203, A: 255}},
		{"PLUM", color.RGBA{R: 221, G: 160, B: 221, A: 255}},
		{"POWDER BLUE", color.RGBA{R: 176, G: 224, B: 230, A: 255}},
		{"PURPLE", color.RGBA{R: 160, G: 32, B: 240, A: 255}},
		{"REBECCA PURPLE", color.RGBA{R: 102, G: 51, B: 153, A: 255}},
		{"RED", color.RGBA{R: 255, G: 000, B: 000, A: 255}},
		{"ROSY BROWN", color.RGBA{R: 188, G: 143, B: 143, A: 255}},
		{"ROYAL BLUE", color.RGBA{R: 65, G: 105, B: 225, A: 255}},
		{"SADDLE BROWN", color.RGBA{R: 139, G: 69, B: 19, A: 255}},
		{"SALMON", color.RGBA{R: 250, G: 128, B: 114, A: 255}},
		{"SANDY BROWN", color.RGBA{R: 244, G: 164, B: 96, A: 255}},
		{"SEA GREEN", color.RGBA{R: 46, G: 139, B: 87, A: 255}},
		{"SEASHELL", color.RGBA{R: 255, G: 245, B: 238, A: 255}},
		{"SIENNA", color.RGBA{R: 160, G: 82, B: 45, A: 255}},
		{"SILVER", color.RGBA{R: 192, G: 192, B: 192, A: 255}},
		{"SKY BLUE", color.RGBA{R: 135, G: 206, B: 235, A: 255}},
		{"SLATE BLUE", color.RGBA{R: 106, G: 90, B: 205, A: 255}},
		{"SLATE GRAY", color.RGBA{R: 112, G: 128, B: 144, A: 255}},
		{"SNOW", color.RGBA{R: 255, G: 250, B: 250, A: 255}},
		{"SPRING GREEN", color.RGBA{R: 000, G: 255, B: 127, A: 255}},
		{"STEEL BLUE", color.RGBA{R: 70, G: 130, B: 180, A: 255}},
		{"TAN", color.RGBA{R: 210, G: 180, B: 140, A: 255}},
		{"TEAL", color.RGBA{R: 000, G: 128, B: 128, A: 255}},
		{"THISTLE", color.RGBA{R: 216, G: 191, B: 216, A: 255}},
		{"TOMATO", color.RGBA{R: 255, G: 99, B: 71, A: 255}},
		{"TURQUOISE", color.RGBA{R: 64, G: 224, B: 208, A: 255}},
		{"VIOLET", color.RGBA{R: 238, G: 130, B: 238, A: 255}},
		{"WEB GRAY", color.RGBA{R: 128, G: 128, B: 128, A: 255}},
		{"WEB GREEN", color.RGBA{R: 000, G: 128, B: 000, A: 255}},
		{"WEB MAROON", color.RGBA{R: 127, G: 000, B: 000, A: 255}},
		{"WEB PURPLE", color.RGBA{R: 127, G: 000, B: 127, A: 255}},
		{"WHEAT", color.RGBA{R: 245, G: 222, B: 179, A: 255}},
		{"WHITE", color.RGBA{R: 255, G: 255, B: 255, A: 255}},
		{"WHITE SMOKE", color.RGBA{R: 245, G: 245, B: 245, A: 255}},
		{"YELLOW", color.RGBA{R: 255, G: 255, B: 000, A: 255}},
		{"YELLOW GREEN", color.RGBA{R: 154, G: 205, B: 50, A: 255}},
	} {
		for pass := 0; pass < 3; pass++ {
			var s = iter.in
			switch pass {
			case 0: // do nothing
			case 1:
				s = str.ToLower(s)
			case 2:
				s = str.Replace(s, " ", "", -1)
			case 3:
				s = str.ToLower(str.Replace(s, " ", "", -1))
			case 4:
				s = str.Replace(s, " ", "-", -1)
			case 5:
				s = str.ToLower(str.Replace(s, " ", "-", -1))
			case 6:
				s = str.Replace(s, " ", "_", -1)
			case 7:
				s = str.ToLower(str.Replace(s, " ", "_", -1))
			}
			var doc = pdf.NewPDF("A4")
			var doc2 = doc.SetColor(s)
			if doc2 != &doc {
				t.Errorf(
					`Address of pointer returned by SetColor("%s") is wrong`,
					iter.in)
			}
			if doc.Color() != iter.expect {
				t.Errorf(
					`After SetColor("%s"), Color() returned %v instead of %v`,
					iter.in, doc.Color(), iter.expect)
			}
		}
	}
	// test color names with trimming and case insensitivity
	for _, name := range PermuteStrings(
		[]string{"", " ", "  "},
		[]string{"red", "Red", "RED"},
		[]string{"", " ", "  "},
	) {
		var doc = pdf.NewPDF("A4")
		var doc2 = doc.SetColor(name)
		TEqual(t, doc.Color(), color.RGBA{R: 255, A: 255})
		TEqual(t, &doc, doc2)
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
