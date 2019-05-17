// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2019-05-17 19:38:47 4CF777               one-file-pdf/utest/[pdf_test.go]
// -----------------------------------------------------------------------------

package pdf

// # Public Tests:
//   Test_NewPDF_
//   Test_PDF_Clean_
//   Test_PDF_Color_
//   Test_PDF_Compression_
//   Test_PDF_CurrentPage_
//   Test_PDF_DocAuthor_
//   Test_PDF_DocCreator_
//   Test_PDF_DocKeywords_
//   Test_PDF_DocSubject_
//   Test_PDF_DocTitle_
//   Test_PDF_DrawBox_
//   Test_PDF_DrawCircle_
//   Test_PDF_DrawImage_
//   Test_PDF_DrawTextAt_
//   Test_PDF_DrawTextInBox_
//   Test_PDF_DrawText_
//   Test_PDF_DrawUnitGrid_
//   Test_PDF_Errors_
//   Test_PDF_FillBox_
//   Test_PDF_FillCircle_
//   Test_PDF_FontName_
//   Test_PDF_FontSize_
//   Test_PDF_HorizontalScaling_
//   Test_PDF_LineWidth_
//   Test_PDF_PageCount_
//   Test_PDF_PageHeight_
//   Test_PDF_PageWidth_
//   Test_PDF_PullError_
//   Test_PDF_Reset_
//   Test_PDF_SetFont_
//   Test_PDF_SetXY_
//   Test_PDF_ToColor_1_
//   Test_PDF_ToColor_2_
//   Test_PDF_ToPoints_
//   Test_PDF_ToUnits_
//   Test_PDF_Units_
//   Test_PDF_X_
//   Test_PDF_Y_
//
// # Internal Tests
//   Test_getPapreSize_
//
// # Helper Functions
//   callerList() []string
//   failIfHasErrors(t *testing.T, errors func() []error)
//   floatStr(val float64) string
//   formatLines(s string, formatStreams bool) []string
//   getStack() string
//   mismatch(t *testing.T, tag string, want, got interface{})
//   pdfCompare(t *testing.T, got []byte, want string)
//   pdfFormatStreams(s string) string
//   permuteStrings(parts ...[]string) (ret []string)
//   tCaller() string
//   tEqual(t *testing.T, got interface{}, want interface{}) bool

//  This file contains unit tests for internal methods/functions.
//
//  To generate a test coverage report use:
//  	go test -coverprofile cover.out
//  	go tool cover -html=cover.out

import (
	"bytes"
	"fmt"
	"image/color"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"testing"
	"time"
)

// to run all tests

// Test_PDF_Clean_ is the unit test for PDF.Clean()
func Test_PDF_Clean_(t *testing.T) {
	//
	// calling Clean() multiple times on a non-initialized PDF:
	// (you should not do this normally, use NewPDF() to create a PDF)
	// - should not panic
	// - length of Errors() should be zero
	// - Errors() should be []error{}, not nil
	//
	func() {
		var doc PDF // uninitialized PDF
		doc.Clean().Clean().Clean()
		//
		//        got                want
		tEqual(t, len(doc.Errors()), 0)
		tEqual(t, doc.Errors(), []error{})
	}()
	// same as above for a PDF properly initialized with NewPDF()
	// (also, call Clean() without chaining)
	func() {
		doc := NewPDF("A4")
		doc.Clean()
		doc.Clean()
		doc.Clean()
		//        got                want
		tEqual(t, len(doc.Errors()), 0)
		tEqual(t, doc.Errors(), []error{})
	}()
	// create a new PDF with an unknown page size, then call Clean()
	// first, Errors should have 1 error
	// after Clean(), Errors should be zero-length
	func() {
		doc := NewPDF("Parchment")
		//        got                want
		tEqual(t, len(doc.Errors()), 1)
		doc.Clean()
		//        got                want
		tEqual(t, len(doc.Errors()), 0)
		tEqual(t, doc.Errors(), []error{})
	}()

} //                                                             Test_PDF_Clean_

// Test_PDF_Color_ tests PDF.Color() and SetColor()
func Test_PDF_Color_(t *testing.T) {
	//
	// (ob *PDF) Color() color.RGBA
	//
	func() {
		var doc PDF // uninitialized PDF
		tEqual(t, doc.Color(), color.RGBA{A: 255})
	}()
	func() {
		doc := NewPDF("A4") // initialized PDF
		tEqual(t, doc.Color(), color.RGBA{A: 255})
	}()
	//
	// (ob *PDF) SetColor(nameOrHTMLColor string) *PDF
	//
	// test various named colors and codes
	for _, test := range []struct {
		input string
		want  color.RGBA
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
			s := test.input
			switch pass {
			case 0:
				// do nothing
			case 1:
				s = strings.ToLower(s)
			case 2:
				s = strings.Replace(s, " ", "", -1)
			case 3:
				s = strings.ToLower(strings.Replace(s, " ", "", -1))
			case 4:
				s = strings.Replace(s, " ", "-", -1)
			case 5:
				s = strings.ToLower(strings.Replace(s, " ", "-", -1))
			case 6:
				s = strings.Replace(s, " ", "_", -1)
			case 7:
				s = strings.ToLower(strings.Replace(s, " ", "_", -1))
			}
			doc := NewPDF("A4")
			doc2 := doc.SetColor(s)
			if doc2 != &doc {
				t.Errorf(
					`Address of pointer returned by SetColor("%s") is wrong`,
					test.input)
			}
			if doc.Color() != test.want {
				t.Errorf(
					`After SetColor("%s"), Color() returned %v instead of %v`,
					test.input, doc.Color(), test.want)
			}
		}
	}
	// test color names with trimming and case insensitivity
	for _, name := range permuteStrings(
		[]string{"", " ", "  "},
		[]string{"red", "Red", "RED"},
		[]string{"", " ", "  "},
	) {
		doc := NewPDF("A4")
		doc2 := doc.SetColor(name)
		tEqual(t, doc.Color(), color.RGBA{R: 255, A: 255})
		tEqual(t, &doc, doc2)
	}
	// try setting a blank color name
	func() {
		var doc PDF // uninitialized PDF
		tEqual(t, len(doc.Errors()), 0)
		doc.SetColor("")
		tEqual(t, len(doc.Errors()), 1)
		//
		if len(doc.Errors()) == 1 {
			tEqual(t,
				doc.Errors()[0],
				fmt.Errorf(`Unknown color name "" @SetColor`))
		}
		tEqual(t, doc.Color(), color.RGBA{A: 255})
	}()
	// try setting an unknown color name
	func() {
		var doc PDF // uninitialized PDF
		tEqual(t, len(doc.Errors()), 0)
		doc.SetColor("TheColourOutOfSpace")
		tEqual(t, len(doc.Errors()), 1)
		//
		if len(doc.Errors()) == 1 {
			tEqual(t,
				doc.Errors()[0],
				fmt.Errorf(
					`Unknown color name "TheColourOutOfSpace" @SetColor`))
		}
		tEqual(t, doc.Color(), color.RGBA{A: 255})
	}()
	//
	// SetColorRGB(red, green, blue int) *PDF
	//
	func() {
		// red
		a := NewPDF("A4")
		b := a.SetColorRGB(128, 0, 0)
		tEqual(t, a.Color(), color.RGBA{R: 128, A: 255})
		tEqual(t, &a, b)
	}()
	func() {
		// green
		a := NewPDF("A4")
		b := a.SetColorRGB(0, 128, 0)
		tEqual(t, a.Color(), color.RGBA{G: 128, A: 255})
		tEqual(t, &a, b)
	}()
	func() {
		// blue
		a := NewPDF("A4")
		b := a.SetColorRGB(0, 0, 128)
		tEqual(t, a.Color(), color.RGBA{B: 128, A: 255})
		tEqual(t, &a, b)
	}()
} //                                                             Test_PDF_Color_

// Test_PDF_Compression_ tests PDF.Compression() and SetCompression()
func Test_PDF_Compression_(t *testing.T) {
	draw := func(doc *PDF) {
		doc.SetUnits("cm").
			SetXY(1, 1).
			SetFont("Helvetica", 10).
			DrawText("Hello World! Hello World!")
	}
	func() {
		var doc PDF // uninitialized PDF
		tEqual(t, doc.Compression(), true)
		failIfHasErrors(t, doc.Errors)
	}()
	func() {
		doc := NewPDF("A4") // initialized PDF
		tEqual(t, doc.Compression(), true)
		failIfHasErrors(t, doc.Errors)
	}()
	// generate a simple PDF with compression turned on
	func() {
		const want = `
		%PDF-1.4
		1 0 obj <</Type/Catalog/Pages 2 0 R>>
		endobj
		2 0 obj <</Type/Pages/Count 1/MediaBox[0 0 595 841]/Kids[3 0 R]>>
		endobj
		3 0 obj <</Type/Page/Parent 2 0 R/Contents 4 0 R
		/Resources <</Font <</FNT1 5 0 R>> >> >>
		endobj
		4 0 obj <</Filter/FlateDecode/Length 81>> stream
		0A 78 9C 72 0A 51 D0 77 F3 0B 31 54 30 34 50 08
		49 53 70 0D E1 52 30 D0 33 30 30 40 21 8B D2 B9
		30 05 83 DC B9 9C 42 14 8C 2C 14 2C 0C 8D 15 42
		52 14 34 3C 52 73 72 F2 15 C2 F3 8B 72 52 14 15
		90 39 9A 0A 21 59 20 93 01 01 00 00 FF FF F6 FE
		19 77 0A
		endstream
		endobj
		5 0 obj <</Type/Font/Subtype/Type1/Name/FNT1
		/BaseFont/Helvetica
		/Encoding/StandardEncoding>>
		endobj
		xref
		0 6
		0000000000 65535 f
		0000000010 00000 n
		0000000056 00000 n
		0000000130 00000 n
		0000000228 00000 n
		0000000378 00000 n
		trailer
		<</Size 6/Root 1 0 R>>
		startxref
		479
		%%EOF
		`
		doc := NewPDF("A4") // initialized PDF
		doc.SetCompression(true)
		draw(&doc)
		failIfHasErrors(t, doc.Errors)
		pdfCompare(t, doc.Bytes(), want)
	}()
	// generate a simple PDF with compression turned off
	func() {
		want := `
		%PDF-1.4
		1 0 obj <</Type/Catalog/Pages 2 0 R>>
		endobj
		2 0 obj <</Type/Pages/Count 1/MediaBox[0 0 595 841]/Kids[3 0 R]>>
		endobj
		3 0 obj <</Type/Page/Parent 2 0 R/Contents 4 0 R
		/Resources <</Font <</FNT1 5 0 R>> >> >>
		endobj
		4 0 obj <</Length 108>> stream
		BT /FNT1 10 Tf ET
		0.000 0.000 0.000 rg
		0.000 0.000 0.000 RG
		BT 28 813 Td (Hello World! Hello World!) Tj ET
		endstream
		endobj
		5 0 obj <</Type/Font/Subtype/Type1/Name/FNT1
		/BaseFont/Helvetica
		/Encoding/StandardEncoding>>
		endobj
		xref
		0 6
		0000000000 65535 f
		0000000010 00000 n
		0000000056 00000 n
		0000000130 00000 n
		0000000228 00000 n
		0000000387 00000 n
		trailer
		<</Size 6/Root 1 0 R>>
		startxref
		488
		%%EOF
		`
		doc := NewPDF("A4") // initialized PDF
		doc.SetCompression(false)
		draw(&doc)
		failIfHasErrors(t, doc.Errors)
		pdfCompare(t, doc.Bytes(), want)
	}()
} //                                                       Test_PDF_Compression_

// Test_PDF_CurrentPage_ tests PDF.CurrentPage()
func Test_PDF_CurrentPage_(t *testing.T) {
	func() {
		var doc PDF // uninitialized PDF
		//
		// before calling AddPage(), returns 1
		tEqual(t, doc.CurrentPage(), 1)
		//
		// since AddPage() is called without any drawing method,
		// the page is added implicitly: therefore still on page 1
		doc.AddPage()
		tEqual(t, doc.CurrentPage(), 1)
		//
		// the next call to AddPage(), returns 2, and so on
		doc.AddPage()
		tEqual(t, doc.CurrentPage(), 2)
		//
		doc.AddPage()
		tEqual(t, doc.CurrentPage(), 3)
		//
		doc.AddPage()
		tEqual(t, doc.CurrentPage(), 4)
	}()
	func() {
		doc := NewPDF("LETTER")
		//
		// before calling AddPage(), returns 1
		tEqual(t, doc.CurrentPage(), 1)
		//
		// since AddPage() is called without any drawing method,
		// the page is added implicitly: therefore still on page 1
		doc.AddPage()
		tEqual(t, doc.CurrentPage(), 1)
		//
		// the next call to AddPage(), returns 2, and so on
		doc.AddPage()
		tEqual(t, doc.CurrentPage(), 2)
		//
		doc.AddPage()
		tEqual(t, doc.CurrentPage(), 3)
		//
		doc.AddPage()
		tEqual(t, doc.CurrentPage(), 4)
	}()
} //                                                       Test_PDF_CurrentPage_

// Test_PDF_DocAuthor_ is the unit test for
func Test_PDF_DocAuthor_(t *testing.T) {
	//
	// (ob *PDF) DocAuthor() string
	//
	func() {
		var doc PDF // uninitialized PDF
		tEqual(t, doc.DocAuthor(), "")
	}()
	func() {
		doc := NewPDF("A4") // initialized PDF
		tEqual(t, doc.DocAuthor(), "")
	}()
	//
	// (ob *PDF) SetDocAuthor(s string) *PDF
	//
	func() {
		var doc PDF // uninitialized PDF
		tEqual(t, doc.SetDocAuthor("Abcdefg").DocAuthor(), "Abcdefg")
	}()
	func() {
		doc := NewPDF("A4") // initialized PDF
		tEqual(t, doc.SetDocAuthor("Abcdefg").DocAuthor(), "Abcdefg")
	}()
	//
	// Test PDF generation
	//
	func() {
		doc := NewPDF("A4") // initialized PDF
		doc.SetCompression(false).SetDocAuthor("'Author' metadata entry")
		const want = `
		%PDF-1.4
		1 0 obj <</Type/Catalog/Pages 2 0 R>>
		endobj
		2 0 obj <</Type/Pages/Count 1/MediaBox[0 0 595 841]/Kids[3 0 R]>>
		endobj
		3 0 obj <</Type/Page/Parent 2 0 R/Contents 4 0 R>>
		endobj
		4 0 obj <</Length 0>> stream
		endstream
		endobj
		5 0 obj <</Type/Info/Author ('Author' metadata entry)>>
		endobj
		xref
		0 6
		0000000000 65535 f
		0000000010 00000 n
		0000000056 00000 n
		0000000130 00000 n
		0000000189 00000 n
		0000000238 00000 n
		trailer
		<</Size 6/Root 1 0 R/Info 5 0 R>>
		startxref
		302
		%%EOF
		`
		pdfCompare(t, doc.Bytes(), want)
	}()
} //                                                         Test_PDF_DocAuthor_

// Test_PDF_DocCreator_ is the unit test for
func Test_PDF_DocCreator_(t *testing.T) {
	//
	// (ob *PDF) DocCreator() string
	//
	func() {
		var doc PDF // uninitialized PDF
		tEqual(t, doc.DocCreator(), "")
	}()
	func() {
		doc := NewPDF("A4") // initialized PDF
		tEqual(t, doc.DocCreator(), "")
	}()
	//
	// (ob *PDF) SetDocCreator(s string) *PDF
	//
	func() {
		var doc PDF // uninitialized PDF
		tEqual(t, doc.SetDocCreator("Abcdefg").DocCreator(), "Abcdefg")
	}()
	func() {
		doc := NewPDF("A4") // initialized PDF
		tEqual(t, doc.SetDocCreator("Abcdefg").DocCreator(), "Abcdefg")
	}()
	//
	// Test PDF generation
	//
	func() {
		doc := NewPDF("A4") // initialized PDF
		doc.SetCompression(false).SetDocCreator("'Creator' metadata entry")
		const want = `
		%PDF-1.4
		1 0 obj <</Type/Catalog/Pages 2 0 R>>
		endobj
		2 0 obj <</Type/Pages/Count 1/MediaBox[0 0 595 841]/Kids[3 0 R]>>
		endobj
		3 0 obj <</Type/Page/Parent 2 0 R/Contents 4 0 R>>
		endobj
		4 0 obj <</Length 0>> stream
		endstream
		endobj
		5 0 obj <</Type/Info/Creator ('Creator' metadata entry)>>
		endobj
		xref
		0 6
		0000000000 65535 f
		0000000010 00000 n
		0000000056 00000 n
		0000000130 00000 n
		0000000189 00000 n
		0000000238 00000 n
		trailer
		<</Size 6/Root 1 0 R/Info 5 0 R>>
		startxref
		304
		%%EOF
		`
		pdfCompare(t, doc.Bytes(), want)
	}()
} //                                                        Test_PDF_DocCreator_

// Test_PDF_DocKeywords_ is the unit test for
func Test_PDF_DocKeywords_(t *testing.T) {
	//
	// (ob *PDF) DocKeywords() string
	//
	func() {
		var doc PDF // uninitialized PDF
		tEqual(t, doc.DocKeywords(), "")
	}()
	func() {
		doc := NewPDF("A4") // initialized PDF
		tEqual(t, doc.DocKeywords(), "")
	}()
	//
	// (ob *PDF) SetDocKeywords(s string) *PDF
	//
	func() {
		var doc PDF // uninitialized PDF
		tEqual(t, doc.SetDocKeywords("Abcdefg").DocKeywords(), "Abcdefg")
	}()
	func() {
		doc := NewPDF("A4") // initialized PDF
		tEqual(t, doc.SetDocKeywords("Abcdefg").DocKeywords(), "Abcdefg")
	}()
	//
	// Test PDF generation
	//
	func() {
		doc := NewPDF("A4") // initialized PDF
		doc.SetCompression(false).SetDocKeywords("'Keywords' metadata entry")
		const want = `
		%PDF-1.4
		1 0 obj <</Type/Catalog/Pages 2 0 R>>
		endobj
		2 0 obj <</Type/Pages/Count 1/MediaBox[0 0 595 841]/Kids[3 0 R]>>
		endobj
		3 0 obj <</Type/Page/Parent 2 0 R/Contents 4 0 R>>
		endobj
		4 0 obj <</Length 0>> stream
		endstream
		endobj
		5 0 obj <</Type/Info/Keywords ('Keywords' metadata entry)>>
		endobj
		xref
		0 6
		0000000000 65535 f
		0000000010 00000 n
		0000000056 00000 n
		0000000130 00000 n
		0000000189 00000 n
		0000000238 00000 n
		trailer
		<</Size 6/Root 1 0 R/Info 5 0 R>>
		startxref
		306
		%%EOF
		`
		pdfCompare(t, doc.Bytes(), want)
	}()
} //                                                       Test_PDF_DocKeywords_

// Test_PDF_DocSubject_ is the unit test for
func Test_PDF_DocSubject_(t *testing.T) {
	//
	// (ob *PDF) DocSubject() string
	//
	func() {
		var doc PDF // uninitialized PDF
		tEqual(t, doc.DocSubject(), "")
	}()
	func() {
		doc := NewPDF("A4") // initialized PDF
		tEqual(t, doc.DocSubject(), "")
	}()
	//
	// (ob *PDF) SetDocSubject(s string) *PDF
	//
	func() {
		var doc PDF // uninitialized PDF
		tEqual(t, doc.SetDocSubject("Abcdefg").DocSubject(), "Abcdefg")
	}()
	func() {
		doc := NewPDF("A4") // initialized PDF
		tEqual(t, doc.SetDocSubject("Abcdefg").DocSubject(), "Abcdefg")
	}()
	//
	// Test PDF generation
	//
	func() {
		doc := NewPDF("A4") // initialized PDF
		doc.SetCompression(false).SetDocSubject("'Subject' metadata entry")
		const want = `
		%PDF-1.4
		1 0 obj <</Type/Catalog/Pages 2 0 R>>
		endobj
		2 0 obj <</Type/Pages/Count 1/MediaBox[0 0 595 841]/Kids[3 0 R]>>
		endobj
		3 0 obj <</Type/Page/Parent 2 0 R/Contents 4 0 R>>
		endobj
		4 0 obj <</Length 0>> stream
		endstream
		endobj
		5 0 obj <</Type/Info/Subject ('Subject' metadata entry)>>
		endobj
		xref
		0 6
		0000000000 65535 f
		0000000010 00000 n
		0000000056 00000 n
		0000000130 00000 n
		0000000189 00000 n
		0000000238 00000 n
		trailer
		<</Size 6/Root 1 0 R/Info 5 0 R>>
		startxref
		304
		%%EOF
		`
		pdfCompare(t, doc.Bytes(), want)
	}()
} //                                                        Test_PDF_DocSubject_

// Test_PDF_DocTitle_ is the unit test for
func Test_PDF_DocTitle_(t *testing.T) {
	//
	// (ob *PDF) DocTitle() string
	//
	func() {
		var doc PDF // uninitialized PDF
		tEqual(t, doc.DocTitle(), "")
	}()
	func() {
		doc := NewPDF("A4") // initialized PDF
		tEqual(t, doc.DocTitle(), "")
	}()
	//
	// (ob *PDF) SetDocTitle(s string) *PDF
	//
	func() {
		var doc PDF // uninitialized PDF
		tEqual(t, doc.SetDocTitle("Abcdefg").DocTitle(), "Abcdefg")
	}()
	func() {
		doc := NewPDF("A4") // initialized PDF
		tEqual(t, doc.SetDocTitle("Abcdefg").DocTitle(), "Abcdefg")
	}()
	//
	// Test PDF generation
	//
	func() {
		doc := NewPDF("A4")
		doc.SetCompression(false).SetDocTitle("'Title' metadata entry")
		const want = `
		%PDF-1.4
		1 0 obj <</Type/Catalog/Pages 2 0 R>>
		endobj
		2 0 obj <</Type/Pages/Count 1/MediaBox[0 0 595 841]/Kids[3 0 R]>>
		endobj
		3 0 obj <</Type/Page/Parent 2 0 R/Contents 4 0 R>>
		endobj
		4 0 obj <</Length 0>> stream
		endstream
		endobj
		5 0 obj <</Type/Info/Title ('Title' metadata entry)>>
		endobj
		xref
		0 6
		0000000000 65535 f
		0000000010 00000 n
		0000000056 00000 n
		0000000130 00000 n
		0000000189 00000 n
		0000000238 00000 n
		trailer
		<</Size 6/Root 1 0 R/Info 5 0 R>>
		startxref
		300
		%%EOF
		`
		pdfCompare(t, doc.Bytes(), want)
	}()
} //                                                          Test_PDF_DocTitle_

// Test_PDF_DrawBox_ is the unit test for
// PDF.DrawBox(x, y, width, height float64, fill ...bool) *PDF
//
// Runs the test by drawing three rectangles and one filled rectangle
func Test_PDF_DrawBox_(t *testing.T) {
	var (
		doc = NewPDF("18cm x 18cm")
		x   = 1.0
		y   = 1.0
	)
	{
		doc.SetCompression(false).
			SetUnits("cm").
			SetLineWidth(5).
			SetColor("Black").DrawBox(x, y, 1, 1, true). // fill
			SetColor("Red").DrawBox(x, y, 4, 4).
			SetColor("DarkGreen").DrawBox(x, y, 9, 9).
			SetColor("Blue").DrawBox(x, y, 16, 16)
	}
	const want = `
	%PDF-1.4
	1 0 obj <</Type/Catalog/Pages 2 0 R>>
	endobj
	2 0 obj <</Type/Pages/Count 1/MediaBox[0 0 510 510]/Kids[3 0 R]>>
	endobj
	3 0 obj <</Type/Page/Parent 2 0 R/Contents 4 0 R>>
	endobj
	4 0 obj <</Length 255>> stream
	0.000 0.000 0.000 rg
	0.000 0.000 0.000 RG
	5.000 w
	28.346 453.543 28.346 28.346 re b
	1.000 0.000 0.000 RG
	28.346 368.504 113.386 113.386 re S
	0.000 0.392 0.000 RG
	28.346 226.772 255.118 255.118 re S
	0.000 0.000 1.000 RG
	28.346 28.346 453.543 453.543 re S
	endstream
	endobj
	xref
	0 5
	0000000000 65535 f
	0000000010 00000 n
	0000000056 00000 n
	0000000130 00000 n
	0000000189 00000 n
	trailer
	<</Size 5/Root 1 0 R>>
	startxref
	495
	%%EOF
	`
	pdfCompare(t, doc.Bytes(), want)
} //                                                           Test_PDF_DrawBox_

// Test_PDF_DrawCircle_ is the unit test for
// PDF.DrawCircle(x, y, radius float64, fill ...bool) *PDF
//
// Runs the test by drawing three concentric
// circles and one small filled circle
func Test_PDF_DrawCircle_(t *testing.T) {
	var (
		doc  = NewPDF("20cm x 20cm")
		x, y = 10.0, 10.0 // center of page
	)
	{
		doc.SetCompression(false).
			SetUnits("cm").
			SetLineWidth(5).
			SetColor("Black").DrawCircle(x, y, 0.5, true). // fill
			SetColor("Red").DrawCircle(x, y, 2).
			SetColor("DarkGreen").DrawCircle(x, y, 4.5).
			SetColor("Blue").DrawCircle(x, y, 8.5)
	}
	const want = `
	%PDF-1.4
	1 0 obj <</Type/Catalog/Pages 2 0 R>>
	endobj
	2 0 obj <</Type/Pages/Count 1/MediaBox[0 0 566 566]/Kids[3 0 R]>>
	endobj
	3 0 obj <</Type/Page/Parent 2 0 R/Contents 4 0 R>>
	endobj
	4 0 obj <</Length 1003>> stream
	0.000 0.000 0.000 rg
	0.000 0.000 0.000 RG
	5.000 w
	269.291 283.465 m
	269.291 291.292 275.637 297.638 283.465 297.638 c
	291.292 297.638 297.638 291.292 297.638 283.465 c
	297.638 275.637 291.292 269.291 283.465 269.291 c
	275.637 269.291 269.291 275.637 269.291 283.465 c
	b
	1.000 0.000 0.000 RG
	226.772 283.465 m
	226.772 314.775 252.154 340.157 283.465 340.157 c
	314.775 340.157 340.157 314.775 340.157 283.465 c
	340.157 252.154 314.775 226.772 283.465 226.772 c
	252.154 226.772 226.772 252.154 226.772 283.465 c
	S
	0.000 0.392 0.000 RG
	155.906 283.465 m
	155.906 353.913 213.016 411.024 283.465 411.024 c
	353.913 411.024 411.024 353.913 411.024 283.465 c
	411.024 213.016 353.913 155.906 283.465 155.906 c
	213.016 155.906 155.906 213.016 155.906 283.465 c
	S
	0.000 0.000 1.000 RG
	42.520 283.465 m
	42.520 416.535 150.394 524.409 283.465 524.409 c
	416.535 524.409 524.409 416.535 524.409 283.465 c
	524.409 150.394 416.535 42.520 283.465 42.520 c
	150.394 42.520 42.520 150.394 42.520 283.465 c
	S
	endstream
	endobj
	xref
	0 5
	0000000000 65535 f
	0000000010 00000 n
	0000000056 00000 n
	0000000130 00000 n
	0000000189 00000 n
	trailer
	<</Size 5/Root 1 0 R>>
	startxref
	1244
	%%EOF
	`
	pdfCompare(t, doc.Bytes(), want)
} //                                                        Test_PDF_DrawCircle_

// Test_PDF_DrawImage_ is the unit test for
// PDF.DrawImage(x, y, height float64, fileNameOrBytes interface{},
//     backColor ...string) *PDF
//
// Runs the test by drawing rgbw64.png:
// a small 64 x 64 PNG split into pure red, green, blue
// and transparent gradient squares
func Test_PDF_DrawImage_(t *testing.T) {
	var (
		x       = 5.0
		y       = 5.0
		height  = 10.0
		pngData = []byte{
			0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A,
			0x00, 0x00, 0x00, 0x0D, 0x49, 0x48, 0x44, 0x52,
			0x00, 0x00, 0x00, 0x40, 0x00, 0x00, 0x00, 0x40,
			0x08, 0x02, 0x00, 0x00, 0x00, 0x25, 0x0B, 0xE6,
			0x89, 0x00, 0x00, 0x00, 0x01, 0x73, 0x52, 0x47,
			0x42, 0x00, 0xAE, 0xCE, 0x1C, 0xE9, 0x00, 0x00,
			0x00, 0x04, 0x67, 0x41, 0x4D, 0x41, 0x00, 0x00,
			0xB1, 0x8F, 0x0B, 0xFC, 0x61, 0x05, 0x00, 0x00,
			0x00, 0x09, 0x70, 0x48, 0x59, 0x73, 0x00, 0x00,
			0x0E, 0xC3, 0x00, 0x00, 0x0E, 0xC3, 0x01, 0xC7,
			0x6F, 0xA8, 0x64, 0x00, 0x00, 0x00, 0x9B, 0x49,
			0x44, 0x41, 0x54, 0x68, 0x43, 0xED, 0xCF, 0x21,
			0x0E, 0x00, 0x31, 0x10, 0xC3, 0xC0, 0xFE, 0xFF,
			0xD3, 0x3D, 0xEE, 0x03, 0x21, 0x0D, 0xB0, 0x94,
			0x68, 0x88, 0x97, 0xED, 0xB9, 0xA7, 0x8B, 0xAD,
			0xC3, 0xD6, 0x61, 0xEB, 0xB0, 0x75, 0xD8, 0x3A,
			0x6C, 0x1D, 0xB6, 0x0E, 0x5B, 0x87, 0xAD, 0xC3,
			0xD6, 0x61, 0xEB, 0xB0, 0x75, 0xD8, 0x3A, 0x6C,
			0x1D, 0xB6, 0x0E, 0x5B, 0x87, 0xAD, 0xC3, 0xD6,
			0x61, 0xEB, 0xB0, 0x75, 0xD8, 0x3A, 0x6C, 0x1D,
			0xB6, 0x0E, 0x5B, 0x87, 0xAD, 0xC3, 0xD6, 0x61,
			0xEB, 0xB0, 0x75, 0xD8, 0x3A, 0x6C, 0x1D, 0xB6,
			0x0E, 0xFB, 0xB9, 0xDF, 0xE1, 0xB1, 0xF6, 0xF6,
			0x40, 0xD2, 0xDE, 0x1E, 0x48, 0xDA, 0xDB, 0x03,
			0x49, 0x7B, 0x7B, 0x20, 0x69, 0x6F, 0x0F, 0x24,
			0xED, 0xED, 0x81, 0xA4, 0xBD, 0x3D, 0x90, 0xB4,
			0xB7, 0x07, 0x92, 0xF6, 0xF6, 0x40, 0xD2, 0xDE,
			0x1E, 0x48, 0xDA, 0xDB, 0x03, 0x49, 0x7B, 0x7B,
			0x20, 0x69, 0x6F, 0x0F, 0x24, 0xED, 0xED, 0x81,
			0xA4, 0xBD, 0x3D, 0x90, 0xB4, 0x27, 0x7F, 0xE0,
			0xDE, 0x0F, 0x44, 0xB5, 0xE9, 0x5A, 0xA4, 0x14,
			0xD5, 0xC4, 0x00, 0x00, 0x00, 0x00, 0x49, 0x45,
			0x4E, 0x44, 0xAE, 0x42, 0x60, 0x82,
		}
	)
	const wantOpaque = `
	%PDF-1.4
	1 0 obj <</Type/Catalog/Pages 2 0 R>>
	endobj
	2 0 obj <</Type/Pages/Count 1/MediaBox[0 0 566 566]/Kids[3 0 R]>>
	endobj
	3 0 obj <</Type/Page/Parent 2 0 R/Contents 4 0 R
	/Resources <</XObject <</IMG0 5 0 R>> >> >>
	endobj
	4 0 obj <</Filter/FlateDecode/Length 50>> stream
	0A 78 9C 2A E4 32 B2 30 D6 33 31 33 55 30 50 30
	50 80 B1 0D 4D 0C F5 CC 8D 8D E0 74 72 2E 97 BE
	A7 AF BB 81 82 4B 3E 57 20 17 20 00 00 FF FF 0E
	D0 0A 7E 0A
	endstream
	endobj
	5 0 obj <</Type/XObject/Subtype/Image
	/Width 64/Height 64/ColorSpace/DeviceRGB/BitsPerComponent 8
	/Filter/FlateDecode/Length 89>> stream
	0A 78 9C EC CE B1 0D 00 30 08 C4 40 F6 5F 9A 0C
	90 DA 05 D2 BD A8 CD ED 4C 7A 6D 7D EB 3C 3F 3F
	3F 3F 3F 3F 3F 3F 3F 3F 3F 3F 3F 3F 3F 3F 3F 3F
	3F 3F 3F FF 49 7F FD A0 1E 3F 3F 3F 3F 3F 3F 3F
	3F 3F 3F 3F 3F 3F 3F 3F 3F 3F 3F 3F 3F 3F 3F FF
	BF 17 00 00 FF FF 93 C7 E9 5A 0A
	endstream
	endobj
	xref
	0 6
	0000000000 65535 f
	0000000010 00000 n
	0000000056 00000 n
	0000000130 00000 n
	0000000231 00000 n
	0000000350 00000 n
	trailer
	<</Size 6/Root 1 0 R>>
	startxref
	596
	%%EOF
	`
	// generate image from an array of PNG bytes
	func() {
		doc := NewPDF("20cm x 20cm")
		doc.SetCompression(true).
			SetUnits("cm").
			DrawImage(x, y, height, pngData)
		failIfHasErrors(t, doc.Errors)
		pdfCompare(t, doc.Bytes(), wantOpaque)
	}()
	// the same test, but reading direcly from PNG file
	func() {
		doc := NewPDF("20cm x 20cm")
		doc.SetCompression(true).
			SetUnits("cm").
			DrawImage(x, y, height, "./image/rgbw64.png")
		failIfHasErrors(t, doc.Errors)
		pdfCompare(t, doc.Bytes(), wantOpaque)
	}()
	// PNG transparency test
	func() {
		var (
			x      = 5.0
			y      = 5.0
			height = 5.0
		)
		const wantTransparent = `
		%PDF-1.4
		1 0 obj <</Type/Catalog/Pages 2 0 R>>
		endobj
		2 0 obj <</Type/Pages/Count 1/MediaBox[0 0 566 566]/Kids[3 0 R]>>
		endobj
		3 0 obj <</Type/Page/Parent 2 0 R/Contents 4 0 R
		/Resources <</XObject <<
		/IMG0 5 0 R
		/IMG1 6 0 R>> >> >>
		endobj
		4 0 obj <</Filter/FlateDecode/Length 66>> stream
		0A 78 9C 2A E4 32 34 31 D4 33 37 36 52 30 50 30
		50 80 B1 61 B4 91 85 B1 9E 89 99 A9 42 72 2E 97
		BE A7 AF BB 81 82 4B 3E 57 20 17 7E 3D 86 86 C6
		7A C6 16 66 30 3D 86 10 3D 80 00 00 00 FF FF 37
		AA 14 E2 0A
		endstream
		endobj
		5 0 obj <</Type/XObject/Subtype/Image
		/Width 64/Height 64/ColorSpace/DeviceRGB/BitsPerComponent 8
		/Filter/FlateDecode/Length 862>> stream
		0A 78 9C EC D6 CD 4B 1B 69 00 C7 F1 DF 7F 53 68
		0F 39 74 C1 82 81 15 EC A1 85 16 5A 48 C5 2C 4D
		D8 94 B6 A8 A0 A0 12 25 BE 62 C4 C8 EA 6A C4 48
		E2 6E 22 51 9C 90 04 13 54 50 31 A2 A2 01 85 08
		0A 7A 50 50 88 A0 A0 82 82 07 3D 78 10 B2 CC 3C
		33 F3 CC 24 13 5F 0E CF 42 E1 19 E6 3C F9 F0 E5
		F7 3C 24 07 30 7D D9 7E 3D C7 FA F3 DC CF FD DC
		CF FD DC CF FD DC CF FD DC CF FD DC CF FD DC CF
		FD DC CF FD DC CF FD DC CF FD DC CF FD BF A4 FF
		7F F8 81 67 3D B9 67 3E DC CF FD DC CF FD DC CF
		FD DC 5F F8 BE 78 81 57 2F 61 32 E1 F5 6B 94 BC
		81 B9 14 65 65 78 FB 16 EF DF E3 E3 47 7C FE 8C
		2F 15 B0 FE 01 BB 1D DF BE E1 E7 4F D4 D4 A0 AE
		0E 8D 8D 68 6E 46 6B 2B 3A 3B D1 DD 8D DE 5E FC
		DD 8F A1 21 8C 8C E0 9F 51 84 C6 30 31 81 48 04
		F1 38 F3 FF 0F 2F 5F 29 F8 12 98 CD 32 FE DD 3B
		7C F8 20 E1 BF 88 78 DB 13 F0 5E 09 3F 2A E1 C7
		C7 11 11 44 7C 32 C9 DC 6F 32 E1 37 82 2F D5 E1
		3F 7D 42 45 05 AC 56 D8 6D 3A 7C 43 83 88 6F 6B
		93 F0 6E 0D DE 27 E2 C7 42 B4 7C 32 89 D9 59 E6
		7E 79 36 66 94 FD AE C3 8B E5 AD 62 79 87 E3 A1
		F2 FD 5A BC 54 5E 10 28 7E 61 81 B9 BF E4 0D 4A
		0B CA 13 BC DD A6 C3 93 F2 04 EF D6 E0 7D 12 3E
		14 92 66 13 41 4C C1 CF CF 63 69 89 B9 BF 18 DE
		56 30 9B A6 26 A5 BC 1B 1E 09 3F E4 95 36 1F 90
		F1 82 66 36 04 BF BA CA DC 2F CE A6 5C 7F 60 15
		FC 8F 1F 46 E5 DD 4A 79 AF 34 1B 15 2F E8 CA A7
		52 22 3E 9D 66 EE 2F 2F A7 E5 2D 16 19 EF 70 50
		3C D9 7C 4B 8B B2 79 8F 8C F7 91 F2 41 B9 7C 2C
		8E 44 82 96 5F 59 11 F1 9B 9B CC FD C5 F0 D5 D5
		74 36 2A DE A3 E2 7D 08 28 F8 88 A0 9B 0D 29 BF
		BE 2E E2 B7 B6 98 FB E5 CD 17 C7 AB B3 21 F8 41
		82 57 0F AC 80 78 8C 96 4F A5 C4 F2 2A 7E 67 87
		B9 DF B0 7C 6D AD AE BC B8 79 7D F9 A0 BA F9 98
		AE 7C 1E 7E 6F 8F B9 DF 62 41 65 65 51 7C 47 87
		82 EF 53 CA 07 10 0C 52 7C 22 81 99 19 1D 7E 63
		83 E2 0F 0E 98 FB 2B AD F8 AA 39 B0 85 78 72 55
		0E 6A CB 87 21 4C 52 FC DC 9C 31 7E 7F 1F 47 47
		CC FD 5F 1F 2C DF E3 41 1F 29 3F 2C E3 C3 E3 14
		AF 9D 4D 3A 6D 80 3F 3E 66 EE 77 FC 49 F1 F5 F5
		14 DF 45 0E 6C 1F BC 83 F9 B3 89 3E A1 FC E1 A1
		88 3F 3D 65 EE FF FE 9D E2 9D 4E 05 DF 45 CB 0F
		FB E0 97 F0 E1 30 26 F5 B3 59 5C 7C 04 7F 7E CE
		DC 5F 55 65 80 F7 68 F0 01 15 2F 20 16 C5 D4 14
		C5 2F 2F CB F8 4C 06 DB DB 06 F8 AB 2B E6 FE A2
		78 69 36 7E 3F 82 FF 3E 52 3E 93 31 28 7F 71 81
		CB 4B 5C 5F 33 F7 93 CD BB 5C 74 F3 04 AF 2B 5F
		80 D7 96 D7 E2 B3 59 B9 3C C1 DF DE 32 F7 3B 9D
		32 DE 2D 95 FF 8B E0 C9 6D A3 94 8F 46 F3 CB AF
		AD D1 D9 EC EE E6 CF 46 C5 DF DD 31 F7 BB 5C 68
		6F 97 0E 6C 0F C5 FB 0B F0 D3 D3 B4 BC 21 3E 9B
		C5 C9 49 3E FE FE 9E B9 5F BE 6D 7A C4 D9 0C 0C
		48 E5 FD 74 36 D1 98 78 60 9F 88 3F 3B 93 F1 37
		37 32 3E 97 C3 7F 01 00 00 FF FF FB 12 FC A0 0A
		endstream
		endobj
		6 0 obj <</Type/XObject/Subtype/Image
		/Width 64/Height 64/ColorSpace/DeviceRGB/BitsPerComponent 8
		/Filter/FlateDecode/Length 872>> stream
		0A 78 9C EC D4 CD 4B 5B 59 1C C6 F1 E7 AF 99 81
		E9 A2 8B 0E 54 48 40 C1 2E 5A 68 A1 85 24 54 99
		84 69 68 4B 5B 68 41 25 4A 7C 25 11 75 C6 97 C4
		49 24 11 93 C1 48 AE 24 A2 A2 01 15 23 2A 1A 50
		88 A0 A0 0B 05 85 08 0A 2A 28 B8 D0 85 0B 21 C3
		BD E7 DC 57 6F 74 5C 9C 81 81 13 EE FA E6 C3 97
		E7 77 8B 00 D3 87 ED DB 8B AC 5F CF FD DC CF FD
		DC CF FD DC CF FD DC CF FD DC CF FD DC CF FD DC
		CF FD DC CF FD DC CF FD DC CF FD DC FF BF F4 FF
		07 7F F0 A8 5F F1 91 3F EE E7 7E EE E7 7E EE E7
		7E EE 37 79 7E FE 09 BF 3C C1 D3 A7 78 F6 0C 65
		65 B0 58 51 51 81 17 2F F0 EA 15 DE BC C1 BB 77
		70 D8 51 5D 0D 97 0B 6E 37 3E 7F C6 B7 6F F8 F1
		03 75 75 68 68 40 53 13 DA DA D0 DE 8E AE 3F D0
		D3 8B FE 7E 0C 0C 60 70 10 7F C7 31 32 82 D1 51
		8C 8D 31 F7 3F 91 F1 CF 9F C3 2A E3 5F BE C4 EB
		D7 22 DE EE 40 75 15 5C CE 7B F1 5D E8 E9 D1 E1
		13 09 08 12 7E 72 92 B9 5F C4 FF 2A E2 2D 16 1D
		FE ED 5B 38 1C A8 AA 82 53 5F BE B6 56 C4 37 37
		8B 78 BF 06 1F 96 F0 71 52 5E A0 F8 4C 86 B9 9F
		CC C6 6A 45 B9 1E 6F B7 8B 78 97 13 1F 3E DC 5B
		BE 57 C5 8B E5 47 D4 F2 99 0C E6 E6 98 FB CB 48
		F9 72 13 BC 53 8F 27 E5 55 7C A7 54 3E 88 70 98
		96 4F 24 E8 E6 09 7E 76 16 0B 0B CC FD 16 AB 71
		36 14 7F 67 36 F5 F5 14 EF F7 D3 D9 04 83 18 08
		23 1A 95 F1 82 11 BF BC CC DC 5F 5E 81 CA 4A CD
		C1 DA 68 79 B7 1B 9F 3E 99 94 57 36 1F 94 CA 47
		E5 F2 82 BE 7C 36 2B E2 73 39 E6 FE 4A 4D 79 9B
		66 36 0A 9E 6C BE B1 51 9A 8D 1F 9D D2 6C 82 D2
		D7 26 1A 45 4C 53 7E 62 42 2D BF B4 24 E2 D7 D7
		99 FB 4B E1 BF 7E 55 67 43 F1 ED E8 D4 97 27 78
		41 C0 58 DA 58 7E 75 55 C4 6F 6C 30 F7 53 BC AD
		24 5E D9 3C 2D 4F 0E 56 DE BC 20 20 9D 56 CB 67
		B3 62 79 05 BF B5 C5 DC 6F 8A FF FE 5D 57 9E E2
		BB E5 F2 51 C4 63 2A 5E 5B DE 80 DF D9 61 EE B7
		D9 F1 FE 7D 49 7C 6B 2B C5 77 9B CD 86 94 9F 9E
		D6 E1 D7 D6 54 FC DE 1E 73 BF E1 60 EF E2 BB 0C
		E5 E3 18 1E 46 52 83 9F 99 31 C7 EF EE E2 E0 80
		B9 DF F9 DB 43 E5 25 7C 28 24 CF 46 83 D7 CE 26
		97 33 C1 1F 1E 32 F7 FF AE C1 D7 D4 68 F0 3E 8A
		0F 04 68 F9 58 CC 38 9B 7B CA EF EF 8B F8 E3 63
		E6 FE 8F 1F 55 BC C7 43 F1 3E 9F 5A 5E C1 8B B3
		49 EA F0 F3 F3 0F E0 4F 4F 99 FB BF 7C B9 83 57
		66 13 40 38 24 E3 13 10 92 48 A5 31 3E AE E2 17
		17 29 3E 9F C7 E6 A6 09 FE E2 82 B9 BF 54 79 32
		9B 48 14 43 43 26 07 AB 2D 9F CF 9B 94 3F 3B C3
		F9 39 2E 2F 99 FB C9 E6 BD 5E 79 F3 1D E8 FE 53
		C2 87 34 B3 11 90 4E E9 F0 DA F2 5A 7C A1 40 CB
		13 FC F5 35 73 BF C7 43 F1 3E 82 97 CA 93 AF 8D
		52 3E 95 32 96 5F 59 51 67 B3 BD 6D 9C 8D 82 BF
		B9 61 EE F7 7A D1 D2 22 E2 3B C8 6C FA 24 7C 44
		C6 27 29 7E 6A 4A 2D 6F 8A 2F 14 70 74 64 C4 DF
		DE 32 F7 93 F2 04 DF 17 40 E8 2F 44 22 88 C9 F8
		74 4A 3C D8 7F 89 3F 39 A1 F8 AB 2B 8A 2F 16 FF
		09 00 00 FF FF 2B A0 FC A9 0A
		endstream
		endobj
		xref
		0 7
		0000000000 65535 f
		0000000010 00000 n
		0000000056 00000 n
		0000000130 00000 n
		0000000244 00000 n
		0000000379 00000 n
		0000001399 00000 n
		trailer
		<</Size 7/Root 1 0 R>>
		startxref
		2429
		%%EOF
		`
		doc := NewPDF("20cm x 20cm")
		doc.SetCompression(true).
			SetUnits("cm").
			DrawImage(x, y, height, "./image/rgbt64.png", "Yellow").
			DrawImage(x, y+height+1, height, "./image/rgbt64.png", "Cyan")
		failIfHasErrors(t, doc.Errors)
		pdfCompare(t, doc.Bytes(), wantTransparent)
	}()
	// wrong argument in fileNameOrBytes
	func() {
		var (
			doc             = NewPDF("20cm x 20cm")
			fileNameOrBytes = []int{0xBAD, 0xBAD, 0xBAD}
		)
		const want = `
		%PDF-1.4
		1 0 obj <</Type/Catalog/Pages 2 0 R>>
		endobj
		2 0 obj <</Type/Pages/Count 1/MediaBox[0 0 566 566]/Kids[3 0 R]>>
		endobj
		3 0 obj <</Type/Page/Parent 2 0 R/Contents 4 0 R>>
		endobj
		4 0 obj <</Filter/FlateDecode/Length 11>> stream
		0A 78 9C 01 00 00 FF FF 00 00 00 01 0A
		endstream
		endobj
		xref
		0 5
		0000000000 65535 f
		0000000010 00000 n
		0000000056 00000 n
		0000000130 00000 n
		0000000189 00000 n
		trailer
		<</Size 5/Root 1 0 R>>
		startxref
		269
		%%EOF
		`
		{
			doc.SetCompression(true).
				SetUnits("cm").
				DrawImage(x, y, height, fileNameOrBytes)
		}
		{
			pdfCompare(t, doc.Bytes(), want)
		}
		tEqual(t, len(doc.Errors()), 1)
		if len(doc.Errors()) > 0 {
			tEqual(t, doc.Errors()[0], fmt.Errorf(
				`Invalid type in fileNameOrBytes`+
					` "[]int = [2989 2989 2989]" @DrawImage`))
		}
	}()
	// drawing an image on one page and another image on second page should work
	// (catch bug where image name in '/XObject' and '/IMG Do' do not match)
	func() {
		const want = `
		%PDF-1.4
		1 0 obj <</Type/Catalog/Pages 2 0 R>>
		endobj
		2 0 obj <</Type/Pages/Count 2/MediaBox[0 0 566 566]/Kids[3 0 R 5 0 R]>>
		endobj
		3 0 obj <</Type/Page/Parent 2 0 R/Contents 4 0 R
		/Resources <</XObject <</IMG0 7 0 R>> >> >>
		endobj
		4 0 obj <</Length 52>> stream
		q
		283.465 0 0 283.465 141.732 141.732 cm
		/IMG0 Do
		Q
		endstream
		endobj
		5 0 obj <</Type/Page/Parent 2 0 R/Contents 6 0 R
		/Resources <</XObject <</IMG1 8 0 R>> >> >>
		endobj
		6 0 obj <</Length 52>> stream
		q
		283.465 0 0 283.465 141.732 141.732 cm
		/IMG1 Do
		Q
		endstream
		endobj
		7 0 obj <</Type/XObject/Subtype/Image
		/Width 64/Height 64/ColorSpace/DeviceRGB/BitsPerComponent 8
		/Filter/FlateDecode/Length 89>> stream
		0A 78 9C EC CE B1 0D 00 30 08 C4 40 F6 5F 9A 0C
		90 DA 05 D2 BD A8 CD ED 4C 7A 6D 7D EB 3C 3F 3F
		3F 3F 3F 3F 3F 3F 3F 3F 3F 3F 3F 3F 3F 3F 3F 3F
		3F 3F 3F FF 49 7F FD A0 1E 3F 3F 3F 3F 3F 3F 3F
		3F 3F 3F 3F 3F 3F 3F 3F 3F 3F 3F 3F 3F 3F 3F FF
		BF 17 00 00 FF FF 93 C7 E9 5A 0A
		endstream
		endobj
		8 0 obj <</Type/XObject/Subtype/Image
		/Width 64/Height 64/ColorSpace/DeviceRGB/BitsPerComponent 8
		/Filter/FlateDecode/Length 89>> stream
		0A 78 9C EC CE B1 0D 00 30 08 C4 40 F6 5F 9A 0C
		90 DA 05 D2 BD A8 CD ED 4C 7A 6D 7D EB 3C 3F 3F
		3F 3F 3F 3F 3F 3F 3F 3F 3F 3F 3F 3F 3F 3F 3F 3F
		3F 3F 3F FF 49 7F FD A0 1E 3F 3F 3F 3F 3F 3F 3F
		3F 3F 3F 3F 3F 3F 3F 3F 3F 3F 3F 3F 3F 3F 3F FF
		BF 17 00 00 FF FF 93 C7 E9 5A 0A
		endstream
		endobj
		xref
		0 9
		0000000000 65535 f
		0000000010 00000 n
		0000000056 00000 n
		0000000136 00000 n
		0000000237 00000 n
		0000000339 00000 n
		0000000440 00000 n
		0000000542 00000 n
		0000000788 00000 n
		trailer
		<</Size 9/Root 1 0 R>>
		startxref
		1034
		%%EOF
		`
		// create two slightly-different images by appending a byte to each
		var (
			pngData1 = append(pngData, 1)
			pngData2 = append(pngData, 2)
			doc      = NewPDF("20cm x 20cm")
		)
		doc.SetCompression(false).
			SetUnits("cm").
			DrawImage(x, y, height, pngData1).
			AddPage().
			DrawImage(x, y, height, pngData2)
		failIfHasErrors(t, doc.Errors)
		pdfCompare(t, doc.Bytes(), want)
	}()
} //                                                         Test_PDF_DrawImage_

// Test_PDF_DrawText_ is the unit test for
// DrawText(s string) *PDF
func Test_PDF_DrawText_(t *testing.T) {
	func() {
		doc := NewPDF("A4")
		{
			doc.SetCompression(false).
				SetUnits("cm").
				SetColumnWidths(1, 4, 9).
				SetColor("#006B3C CadmiumGreen").
				SetFont("Helvetica-Bold", 10).
				SetX(5).
				SetY(5).
				DrawText("FIRST").
				DrawText("SECOND").
				DrawText("THIRD")
		}
		const want = `
		%PDF-1.4
		1 0 obj <</Type/Catalog/Pages 2 0 R>>
		endobj
		2 0 obj <</Type/Pages/Count 1/MediaBox[0 0 595 841]/Kids[3 0 R]>>
		endobj
		3 0 obj <</Type/Page/Parent 2 0 R/Contents 4 0 R
		/Resources <</Font <</FNT1 5 0 R>> >> >>
		endobj
		4 0 obj <</Length 143>> stream
		BT /FNT1 10 Tf ET
		0.000 0.420 0.235 rg
		0.000 0.420 0.235 RG
		BT 0 700 Td (FIRST) Tj ET
		BT 28 700 Td (SECOND) Tj ET
		BT 141 700 Td (THIRD) Tj ET
		endstream
		endobj
		5 0 obj <</Type/Font/Subtype/Type1/Name/FNT1
		/BaseFont/Helvetica-Bold
		/Encoding/StandardEncoding>>
		endobj
		xref
		0 6
		0000000000 65535 f
		0000000010 00000 n
		0000000056 00000 n
		0000000130 00000 n
		0000000228 00000 n
		0000000422 00000 n
		trailer
		<</Size 6/Root 1 0 R>>
		startxref
		528
		%%EOF
		`
		pdfCompare(t, doc.Bytes(), want)
	}()
	func() {
		doc := NewPDF("A4")
		{
			doc.SetCompression(false).
				SetUnits("cm").
				SetFont("Ye-Olde-Scriptte", 10).
				SetXY(5, 5).
				SetHorizontalScaling(150).
				DrawText("Ye-Olde-Scriptte")
		}
		const want = `
		%PDF-1.4
		1 0 obj <</Type/Catalog/Pages 2 0 R>>
		endobj
		2 0 obj <</Type/Pages/Count 1/MediaBox[0 0 595 841]/Kids[3 0 R]>>
		endobj
		3 0 obj <</Type/Page/Parent 2 0 R/Contents 4 0 R
		/Resources <</Font <</FNT1 5 0 R>> >> >>
		endobj
		4 0 obj <</Length 113>> stream
		BT /FNT1 10 Tf ET
		BT 150 Tz ET
		0.000 0.000 0.000 rg
		0.000 0.000 0.000 RG
		BT 141 700 Td (Ye-Olde-Scriptte) Tj ET
		endstream
		endobj
		5 0 obj <</Type/Font/Subtype/Type1/Name/FNT1
		/BaseFont/Helvetica
		/Encoding/StandardEncoding>>
		endobj
		xref
		0 6
		0000000000 65535 f
		0000000010 00000 n
		0000000056 00000 n
		0000000130 00000 n
		0000000228 00000 n
		0000000392 00000 n
		trailer
		<</Size 6/Root 1 0 R>>
		startxref
		493
		%%EOF
		`
		pdfCompare(t, doc.Bytes(), want)
		tEqual(t, len(doc.Errors()), 1)
		tEqual(t, doc.PullError(),
			fmt.Errorf(`Invalid font "Ye-Olde-Scriptte" @DrawText`))
	}()
} //                                                          Test_PDF_DrawText_

// Test_PDF_DrawTextAt_ is the unit test for
// DrawTextAt(x, y float64, text string) *PDF
func Test_PDF_DrawTextAt_(t *testing.T) {
	doc := NewPDF("A4")
	{
		doc.SetCompression(false).
			SetUnits("cm").
			SetColor("#36454F Charcoal").
			SetFont("Helvetica-Bold", 20).
			DrawTextAt(5, 5, ""). // no effect
			DrawTextAt(5, 5, "(5,5)").
			DrawTextAt(10, 10, ""). // no effect
			DrawTextAt(10, 10, "(10,10)").
			DrawTextAt(15, 15, ""). // no effect
			DrawTextAt(15, 15, "(15,15)").
			SetColor("#E03C31 CGRed").
			FillBox(5, 5, 0.1, 0.1).
			FillBox(10, 10, 0.1, 0.1).
			FillBox(15, 15, 0.1, 0.1)
	}
	const want = `
	%PDF-1.4
	1 0 obj <</Type/Catalog/Pages 2 0 R>>
	endobj
	2 0 obj <</Type/Pages/Count 1/MediaBox[0 0 595 841]/Kids[3 0 R]>>
	endobj
	3 0 obj <</Type/Page/Parent 2 0 R/Contents 4 0 R
	/Resources <</Font <</FNT1 5 0 R>> >> >>
	endobj
	4 0 obj <</Length 297>> stream
	BT /FNT1 20 Tf ET
	0.212 0.271 0.310 rg
	0.212 0.271 0.310 RG
	BT 141 700 Td (\(5,5\)) Tj ET
	BT 283 558 Td (\(10,10\)) Tj ET
	BT 425 416 Td (\(15,15\)) Tj ET
	0.878 0.235 0.192 rg
	0.878 0.235 0.192 RG
	141.732 697.323 2.835 2.835 re b
	283.465 555.591 2.835 2.835 re b
	425.197 413.858 2.835 2.835 re b
	endstream
	endobj
	5 0 obj <</Type/Font/Subtype/Type1/Name/FNT1
	/BaseFont/Helvetica-Bold
	/Encoding/StandardEncoding>>
	endobj
	xref
	0 6
	0000000000 65535 f
	0000000010 00000 n
	0000000056 00000 n
	0000000130 00000 n
	0000000228 00000 n
	0000000576 00000 n
	trailer
	<</Size 6/Root 1 0 R>>
	startxref
	682
	%%EOF
	`
	pdfCompare(t, doc.Bytes(), want)
} //                                                        Test_PDF_DrawTextAt_

// Test_PDF_DrawTextInBox_ is the unit test for
// DrawTextInBox(
//     x, y, width, height float64, align, text string) *PDF
func Test_PDF_DrawTextInBox_(t *testing.T) {
	const (
		lorem = "Lorem ipsum dolor sit amet," +
			" consectetur adipiscing elit," +
			" sed do eiusmod tempor incididunt ut" +
			" labore et dolore magna aliqua." +
			" Ut enim ad minim veniam," +
			" quis nostrud exercitation ullamco laboris" +
			" nisi ut aliquip ex ea commodo consequat." +
			" Duis aute irure dolor in reprehenderit in voluptate velit" +
			" esse cillum dolore eu fugiat nulla pariatur." +
			" Excepteur sint occaecat cupidatat non proident," +
			" sunt in culpa qui officia deserunt mollit anim id est laborum."
	)
	func() {
		doc := NewPDF("A4")
		{
			doc.SetCompression(false).
				SetUnits("cm").
				SetFont("Helvetica", 10).
				SetColor("Light Gray").
				FillBox(5, 5, 3, 15).
				SetColor("Black").
				DrawTextInBox(5, 5, 3, 15, "C", ""). // no effect
				DrawTextInBox(5, 5, 3, 15, "C", lorem)
		}
		const want = `
		%PDF-1.4
		1 0 obj <</Type/Catalog/Pages 2 0 R>>
		endobj
		2 0 obj <</Type/Pages/Count 1/MediaBox[0 0 595 841]/Kids[3 0 R]>>
		endobj
		3 0 obj <</Type/Page/Parent 2 0 R/Contents 4 0 R
		/Resources <</Font <</FNT1 5 0 R>> >> >>
		endobj
		4 0 obj <</Length 1252>> stream
		0.827 0.827 0.827 rg
		0.827 0.827 0.827 RG
		141.732 274.961 85.039 425.197 re b
		BT /FNT1 10 Tf ET
		0.000 0.000 0.000 rg
		0.000 0.000 0.000 RG
		BT 153 624 Td (Lorem ipsum ) Tj ET
		BT 151 614 Td (dolor sit amet, ) Tj ET
		BT 157 604 Td (consectetur ) Tj ET
		BT 142 594 Td (adipiscing elit, sed ) Tj ET
		BT 157 584 Td (do eiusmod ) Tj ET
		BT 144 574 Td (tempor incididunt ) Tj ET
		BT 142 564 Td (ut labore et dolore ) Tj ET
		BT 145 554 Td (magna aliqua. Ut ) Tj ET
		BT 150 544 Td (enim ad minim ) Tj ET
		BT 154 534 Td (veniam, quis ) Tj ET
		BT 166 524 Td (nostrud ) Tj ET
		BT 157 514 Td (exercitation ) Tj ET
		BT 149 504 Td (ullamco laboris ) Tj ET
		BT 147 494 Td (nisi ut aliquip ex ) Tj ET
		BT 153 484 Td (ea commodo ) Tj ET
		BT 147 474 Td (consequat. Duis ) Tj ET
		BT 143 464 Td (aute irure dolor in ) Tj ET
		BT 147 454 Td (reprehenderit in ) Tj ET
		BT 152 444 Td (voluptate velit ) Tj ET
		BT 142 434 Td (esse cillum dolore ) Tj ET
		BT 151 424 Td (eu fugiat nulla ) Tj ET
		BT 164 414 Td (pariatur. ) Tj ET
		BT 151 404 Td (Excepteur sint ) Tj ET
		BT 162 394 Td (occaecat ) Tj ET
		BT 152 384 Td (cupidatat non ) Tj ET
		BT 147 374 Td (proident, sunt in ) Tj ET
		BT 148 364 Td (culpa qui officia ) Tj ET
		BT 150 354 Td (deserunt mollit ) Tj ET
		BT 139 344 Td (anim id est laborum.) Tj ET
		endstream
		endobj
		5 0 obj <</Type/Font/Subtype/Type1/Name/FNT1
		/BaseFont/Helvetica
		/Encoding/StandardEncoding>>
		endobj
		xref
		0 6
		0000000000 65535 f
		0000000010 00000 n
		0000000056 00000 n
		0000000130 00000 n
		0000000228 00000 n
		0000001532 00000 n
		trailer
		<</Size 6/Root 1 0 R>>
		startxref
		1633
		%%EOF
		`
		pdfCompare(t, doc.Bytes(), want)
	}()
	func() {
		doc := NewPDF("A4")
		doc.SetCompression(false).
			SetUnits("cm").
			SetFont("Courier", 10).
			SetColor("Light Gray").
			FillBox(5, 5, 3, 15).
			SetColor("Black").
			DrawTextInBox(5, 5, 3, 15, "C", ""). // no effect
			DrawTextInBox(5, 5, 3, 15, "C", lorem)
		const want = `
		%PDF-1.4
		1 0 obj <</Type/Catalog/Pages 2 0 R>>
		endobj
		2 0 obj <</Type/Pages/Count 1/MediaBox[0 0 595 841]/Kids[3 0 R]>>
		endobj
		3 0 obj <</Type/Page/Parent 2 0 R/Contents 4 0 R
		/Resources <</Font <</FNT1 5 0 R>> >> >>
		endobj
		4 0 obj <</Length 1505>> stream
		0.827 0.827 0.827 rg
		0.827 0.827 0.827 RG
		141.732 274.961 85.039 425.197 re b
		BT /FNT1 10 Tf ET
		0.000 0.000 0.000 rg
		0.000 0.000 0.000 RG
		BT 148 679 Td (Lorem ipsum ) Tj ET
		BT 154 669 Td (dolor sit ) Tj ET
		BT 166 659 Td (amet, ) Tj ET
		BT 148 649 Td (consectetur ) Tj ET
		BT 151 639 Td (adipiscing ) Tj ET
		BT 145 629 Td (elit, sed do ) Tj ET
		BT 160 619 Td (eiusmod ) Tj ET
		BT 163 609 Td (tempor ) Tj ET
		BT 142 599 Td (incididunt ut ) Tj ET
		BT 154 589 Td (labore et ) Tj ET
		BT 145 579 Td (dolore magna ) Tj ET
		BT 151 569 Td (aliqua. Ut ) Tj ET
		BT 142 559 Td (enim ad minim ) Tj ET
		BT 145 549 Td (veniam, quis ) Tj ET
		BT 160 539 Td (nostrud ) Tj ET
		BT 145 529 Td (exercitation ) Tj ET
		BT 160 519 Td (ullamco ) Tj ET
		BT 145 509 Td (laboris nisi ) Tj ET
		BT 142 499 Td (ut aliquip ex ) Tj ET
		BT 151 489 Td (ea commodo ) Tj ET
		BT 151 479 Td (consequat. ) Tj ET
		BT 154 469 Td (Duis aute ) Tj ET
		BT 148 459 Td (irure dolor ) Tj ET
		BT 175 449 Td (in ) Tj ET
		BT 142 439 Td (reprehenderit ) Tj ET
		BT 145 429 Td (in voluptate ) Tj ET
		BT 151 419 Td (velit esse ) Tj ET
		BT 142 409 Td (cillum dolore ) Tj ET
		BT 154 399 Td (eu fugiat ) Tj ET
		BT 166 389 Td (nulla ) Tj ET
		BT 154 379 Td (pariatur. ) Tj ET
		BT 154 369 Td (Excepteur ) Tj ET
		BT 142 359 Td (sint occaecat ) Tj ET
		BT 142 349 Td (cupidatat non ) Tj ET
		BT 154 339 Td (proident, ) Tj ET
		BT 142 329 Td (sunt in culpa ) Tj ET
		BT 148 319 Td (qui officia ) Tj ET
		BT 157 309 Td (deserunt ) Tj ET
		BT 148 299 Td (mollit anim ) Tj ET
		BT 139 289 Td (id est laborum.) Tj ET
		endstream
		endobj
		5 0 obj <</Type/Font/Subtype/Type1/Name/FNT1
		/BaseFont/Courier
		/Encoding/StandardEncoding>>
		endobj
		xref
		0 6
		0000000000 65535 f
		0000000010 00000 n
		0000000056 00000 n
		0000000130 00000 n
		0000000228 00000 n
		0000001785 00000 n
		trailer
		<</Size 6/Root 1 0 R>>
		startxref
		1884
		%%EOF
		`
		pdfCompare(t, doc.Bytes(), want)
	}()
} //                                                     Test_PDF_DrawTextInBox_

// Test_PDF_DrawUnitGrid_ is the unit test for PDF.DrawUnitGrid()
func Test_PDF_DrawUnitGrid_(t *testing.T) {
	got := func() []byte {
		doc := NewPDF("A4")
		doc.SetCompression(false).SetUnits("cm").DrawUnitGrid()
		return doc.Bytes()
	}()
	const want = `
	%PDF-1.4
	1 0 obj <</Type/Catalog/Pages 2 0 R>>
	endobj
	2 0 obj <</Type/Pages/Count 1/MediaBox[0 0 595 841]/Kids[3 0 R]>>
	endobj
	3 0 obj <</Type/Page/Parent 2 0 R/Contents 4 0 R
	/Resources <</Font <</FNT1 5 0 R>> >> >>
	endobj
	4 0 obj <</Length 7417>> stream
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	0.100 w
	0.000 841.890 m 0.000 0.000 l S
	BT /FNT1 8 Tf ET
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 2 833 Td (0) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	28.346 841.890 m 28.346 0.000 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 31 833 Td (1) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	56.693 841.890 m 56.693 0.000 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 59 833 Td (2) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	85.039 841.890 m 85.039 0.000 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 87 833 Td (3) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	113.386 841.890 m 113.386 0.000 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 116 833 Td (4) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	141.732 841.890 m 141.732 0.000 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 144 833 Td (5) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	170.079 841.890 m 170.079 0.000 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 172 833 Td (6) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	198.425 841.890 m 198.425 0.000 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 201 833 Td (7) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	226.772 841.890 m 226.772 0.000 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 229 833 Td (8) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	255.118 841.890 m 255.118 0.000 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 257 833 Td (9) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	283.465 841.890 m 283.465 0.000 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 286 833 Td (10) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	311.811 841.890 m 311.811 0.000 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 314 833 Td (11) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	340.157 841.890 m 340.157 0.000 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 342 833 Td (12) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	368.504 841.890 m 368.504 0.000 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 371 833 Td (13) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	396.850 841.890 m 396.850 0.000 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 399 833 Td (14) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	425.197 841.890 m 425.197 0.000 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 428 833 Td (15) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	453.543 841.890 m 453.543 0.000 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 456 833 Td (16) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	481.890 841.890 m 481.890 0.000 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 484 833 Td (17) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	510.236 841.890 m 510.236 0.000 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 513 833 Td (18) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	538.583 841.890 m 538.583 0.000 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 541 833 Td (19) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	566.929 841.890 m 566.929 0.000 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 569 833 Td (20) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	0.000 841.890 m 595.276 841.890 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 2 833 Td (0) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	0.000 813.543 m 595.276 813.543 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 2 805 Td (1) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	0.000 785.197 m 595.276 785.197 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 2 776 Td (2) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	0.000 756.850 m 595.276 756.850 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 2 748 Td (3) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	0.000 728.504 m 595.276 728.504 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 2 720 Td (4) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	0.000 700.157 m 595.276 700.157 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 2 691 Td (5) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	0.000 671.811 m 595.276 671.811 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 2 663 Td (6) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	0.000 643.465 m 595.276 643.465 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 2 634 Td (7) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	0.000 615.118 m 595.276 615.118 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 2 606 Td (8) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	0.000 586.772 m 595.276 586.772 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 2 578 Td (9) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	0.000 558.425 m 595.276 558.425 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 2 549 Td (10) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	0.000 530.079 m 595.276 530.079 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 2 521 Td (11) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	0.000 501.732 m 595.276 501.732 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 2 493 Td (12) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	0.000 473.386 m 595.276 473.386 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 2 464 Td (13) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	0.000 445.039 m 595.276 445.039 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 2 436 Td (14) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	0.000 416.693 m 595.276 416.693 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 2 408 Td (15) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	0.000 388.346 m 595.276 388.346 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 2 379 Td (16) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	0.000 360.000 m 595.276 360.000 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 2 351 Td (17) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	0.000 331.654 m 595.276 331.654 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 2 323 Td (18) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	0.000 303.307 m 595.276 303.307 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 2 294 Td (19) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	0.000 274.961 m 595.276 274.961 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 2 266 Td (20) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	0.000 246.614 m 595.276 246.614 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 2 238 Td (21) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	0.000 218.268 m 595.276 218.268 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 2 209 Td (22) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	0.000 189.921 m 595.276 189.921 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 2 181 Td (23) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	0.000 161.575 m 595.276 161.575 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 2 153 Td (24) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	0.000 133.228 m 595.276 133.228 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 2 124 Td (25) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	0.000 104.882 m 595.276 104.882 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 2 96 Td (26) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	0.000 76.535 m 595.276 76.535 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 2 68 Td (27) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	0.000 48.189 m 595.276 48.189 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 2 39 Td (28) Tj ET
	0.784 0.784 0.784 rg
	0.784 0.784 0.784 RG
	0.000 19.843 m 595.276 19.843 l S
	0.294 0.000 0.510 rg
	0.294 0.000 0.510 RG
	BT 2 11 Td (29) Tj ET
	endstream
	endobj
	5 0 obj <</Type/Font/Subtype/Type1/Name/FNT1
	/BaseFont/Helvetica
	/Encoding/StandardEncoding>>
	endobj
	xref
	0 6
	0000000000 65535 f
	0000000010 00000 n
	0000000056 00000 n
	0000000130 00000 n
	0000000228 00000 n
	0000007697 00000 n
	trailer
	<</Size 6/Root 1 0 R>>
	startxref
	7798
	%%EOF
	`
	pdfCompare(t, got, want)
} //                                                      Test_PDF_DrawUnitGrid_

// Test_PDF_Errors_ tests PDF.Errors()
func Test_PDF_Errors_(t *testing.T) {
	// Errors() should be []error{} on a non-initialized PDF:
	func() {
		var doc PDF // uninitialized PDF
		//
		//        got                want
		tEqual(t, len(doc.Errors()), 0)
		tEqual(t, doc.Errors(), []error{})
	}()
	// same as above for a PDF properly initialized with NewPDF()
	// (also, call Errors() without chaining)
	func() {
		doc := NewPDF("A4")
		tEqual(t, len(doc.Errors()), 0)
		tEqual(t, doc.Errors(), []error{})
	}()
} //                                                            Test_PDF_Errors_

// Test_PDF_FillBox_ is the unit test for
// PDF.FillBox(x, y, width, height float64) *PDF
//
// Runs the test by filling the shape of a Monolith from 2001 Space Odyssey
func Test_PDF_FillBox_(t *testing.T) {
	var (
		doc    = NewPDF("A4")
		x      = 6.5
		y      = 6.0
		width  = 8.0
		height = 18.0
	)
	doc.SetCompression(false).
		SetUnits("cm").
		SetColor("#1B1B1B EerieBlack").
		FillBox(x, y, width, height)
	const want = `
	%PDF-1.4
	1 0 obj <</Type/Catalog/Pages 2 0 R>>
	endobj
	2 0 obj <</Type/Pages/Count 1/MediaBox[0 0 595 841]/Kids[3 0 R]>>
	endobj
	3 0 obj <</Type/Page/Parent 2 0 R/Contents 4 0 R>>
	endobj
	4 0 obj <</Length 80>> stream
	0.106 0.106 0.106 rg
	0.106 0.106 0.106 RG
	184.252 161.575 226.772 510.236 re b
	endstream
	endobj
	xref
	0 5
	0000000000 65535 f
	0000000010 00000 n
	0000000056 00000 n
	0000000130 00000 n
	0000000189 00000 n
	trailer
	<</Size 5/Root 1 0 R>>
	startxref
	319
	%%EOF
	`
	pdfCompare(t, doc.Bytes(), want)
} //                                                           Test_PDF_FillBox_

// Test_PDF_FillCircle_ is the unit test for
// PDF.FillCircle(x, y, radius float64, fill ...bool) *PDF
//
// Runs the test by drawing the flag of Japan using correct proportions
func Test_PDF_FillCircle_(t *testing.T) {
	var (
		doc    = NewPDF("30cm x 20cm")
		x, y   = 15.0, 10.0         // center of page
		radius = (20.0 * 3 / 5) / 2 // diameter = 3/5 of height
	)
	doc.SetCompression(false).
		SetUnits("cm").
		SetColor("#BC002D (close to #BE0032 CrimsonGlory)").
		FillCircle(x, y, radius)
	const want = `
	%PDF-1.4
	1 0 obj <</Type/Catalog/Pages 2 0 R>>
	endobj
	2 0 obj <</Type/Pages/Count 1/MediaBox[0 0 850 566]/Kids[3 0 R]>>
	endobj
	3 0 obj <</Type/Page/Parent 2 0 R/Contents 4 0 R>>
	endobj
	4 0 obj <</Length 267>> stream
	0.737 0.000 0.176 rg
	0.737 0.000 0.176 RG
	255.118 283.465 m
	255.118 377.396 331.265 453.543 425.197 453.543 c
	519.129 453.543 595.276 377.396 595.276 283.465 c
	595.276 189.533 519.129 113.386 425.197 113.386 c
	331.265 113.386 255.118 189.533 255.118 283.465 c
	b
	endstream
	endobj
	xref
	0 5
	0000000000 65535 f
	0000000010 00000 n
	0000000056 00000 n
	0000000130 00000 n
	0000000189 00000 n
	trailer
	<</Size 5/Root 1 0 R>>
	startxref
	507
	%%EOF
	`
	pdfCompare(t, doc.Bytes(), want)
} //                                                        Test_PDF_FillCircle_

// Test_PDF_FontName_ is the unit test for
func Test_PDF_FontName_(t *testing.T) {
	//
	// (ob *PDF) FontName() string
	//
	func() {
		var doc PDF // uninitialized PDF
		tEqual(t, doc.FontName(), "Helvetica")
	}()
	func() {
		doc := NewPDF("A4") // initialized PDF
		tEqual(t, doc.FontName(), "Helvetica")
	}()
	//
	// (ob *PDF) SetFontName(name string) *PDF
	//
	func() {
		var doc PDF // uninitialized PDF
		tEqual(t, doc.SetFontName("Courier").FontName(), "Courier")
	}()
	func() {
		doc := NewPDF("A4") // initialized PDF
		tEqual(t, doc.SetFontName("Courier").FontName(), "Courier")
	}()
	//
	// Test PDF generation
	//
	func() {
		doc := NewPDF("A4")
		doc.SetCompression(false).
			SetUnits("cm").
			SetXY(1, 1).
			SetFont("Helvetica", 10).
			SetFontName("TimesRoman").
			DrawText("Hello World!")
		const want = `
		%PDF-1.4
		1 0 obj <</Type/Catalog/Pages 2 0 R>>
		endobj
		2 0 obj <</Type/Pages/Count 1/MediaBox[0 0 595 841]/Kids[3 0 R]>>
		endobj
		3 0 obj <</Type/Page/Parent 2 0 R/Contents 4 0 R
		/Resources <</Font <</FNT1 5 0 R>> >> >>
		endobj
		4 0 obj <</Length 95>> stream
		BT /FNT1 10 Tf ET
		0.000 0.000 0.000 rg
		0.000 0.000 0.000 RG
		BT 28 813 Td (Hello World!) Tj ET
		endstream
		endobj
		5 0 obj <</Type/Font/Subtype/Type1/Name/FNT1
		/BaseFont/Times-Roman
		/Encoding/StandardEncoding>>
		endobj
		xref
		0 6
		0000000000 65535 f
		0000000010 00000 n
		0000000056 00000 n
		0000000130 00000 n
		0000000228 00000 n
		0000000373 00000 n
		trailer
		<</Size 6/Root 1 0 R>>
		startxref
		476
		%%EOF
		`
		pdfCompare(t, doc.Bytes(), want)
	}()
} //                                                          Test_PDF_FontName_

// Test_PDF_FontSize_ is the unit test for PDF.FontSize() and SetFontSize()
func Test_PDF_FontSize_(t *testing.T) {
	//
	// (ob *PDF) FontSize() string
	//
	func() {
		var doc PDF // uninitialized PDF
		tEqual(t, doc.FontSize(), 10)
	}()
	func() {
		doc := NewPDF("A4") // initialized PDF
		tEqual(t, doc.FontSize(), 10)
	}()
	//
	// (ob *PDF) SetFontSize(name string) *PDF
	//
	func() {
		var doc PDF // uninitialized PDF
		tEqual(t, doc.SetFontSize(15).FontSize(), 15)
	}()
	func() {
		doc := NewPDF("A4") // initialized PDF
		tEqual(t, doc.SetFontSize(20).FontSize(), 20)
	}()
	//
	// Test PDF generation
	//
	func() {
		doc := NewPDF("A4")
		doc.SetCompression(false).SetUnits("cm")
		pt10 := doc.ToUnits(10) // 10pt
		const w = 10
		doc.SetFont("Times-Bold", 20).SetXY(1, 1).DrawText("Font Sizes")
		for i, size := range []float64{5, 6, 7, 8, 9, 10, 15, 20, 25, 30} {
			y := 2 + float64(i)*3*pt10
			doc.SetLineWidth(0.1).
				SetXY(1, y+0.5).SetColor("Gray").
				DrawBox(1, y+0*pt10, w, pt10).
				DrawBox(1, y+1*pt10, w, pt10).
				DrawBox(1, y+2*pt10, w, pt10).
				SetXY(1, y+0.5).SetColor("Black").SetFont("Helvetica", size).
				DrawBox(1, y, w, 3*pt10).
				DrawTextInBox(1, y, w, 3*pt10, "TC",
					fmt.Sprintf("Helvetica %1.0f", size))
		}
		const want = `
		%PDF-1.4
		1 0 obj <</Type/Catalog/Pages 2 0 R>>
		endobj
		2 0 obj <</Type/Pages/Count 1/MediaBox[0 0 595 841]/Kids[3 0 R]>>
		endobj
		3 0 obj <</Type/Page/Parent 2 0 R/Contents 4 0 R
		/Resources <</Font <<
		/FNT1 5 0 R
		/FNT2 6 0 R>> >> >>
		endobj
		4 0 obj <</Length 2440>> stream
		BT /FNT1 20 Tf ET
		0.000 0.000 0.000 rg
		0.000 0.000 0.000 RG
		BT 28 813 Td (Font Sizes) Tj ET
		0.745 0.745 0.745 RG
		0.100 w
		28.346 775.197 283.465 10.000 re S
		28.346 765.197 283.465 10.000 re S
		28.346 755.197 283.465 10.000 re S
		0.000 0.000 0.000 RG
		28.346 755.197 283.465 30.000 re S
		BT /FNT2 5 Tf ET
		BT 157 780 Td (Helvetica 5) Tj ET
		0.745 0.745 0.745 RG
		28.346 745.197 283.465 10.000 re S
		28.346 735.197 283.465 10.000 re S
		28.346 725.197 283.465 10.000 re S
		0.000 0.000 0.000 RG
		28.346 725.197 283.465 30.000 re S
		BT /FNT2 6 Tf ET
		BT 155 749 Td (Helvetica 6) Tj ET
		0.745 0.745 0.745 RG
		28.346 715.197 283.465 10.000 re S
		28.346 705.197 283.465 10.000 re S
		28.346 695.197 283.465 10.000 re S
		0.000 0.000 0.000 RG
		28.346 695.197 283.465 30.000 re S
		BT /FNT2 7 Tf ET
		BT 152 718 Td (Helvetica 7) Tj ET
		0.745 0.745 0.745 RG
		28.346 685.197 283.465 10.000 re S
		28.346 675.197 283.465 10.000 re S
		28.346 665.197 283.465 10.000 re S
		0.000 0.000 0.000 RG
		28.346 665.197 283.465 30.000 re S
		BT /FNT2 8 Tf ET
		BT 150 687 Td (Helvetica 8) Tj ET
		0.745 0.745 0.745 RG
		28.346 655.197 283.465 10.000 re S
		28.346 645.197 283.465 10.000 re S
		28.346 635.197 283.465 10.000 re S
		0.000 0.000 0.000 RG
		28.346 635.197 283.465 30.000 re S
		BT /FNT2 9 Tf ET
		BT 147 656 Td (Helvetica 9) Tj ET
		0.745 0.745 0.745 RG
		28.346 625.197 283.465 10.000 re S
		28.346 615.197 283.465 10.000 re S
		28.346 605.197 283.465 10.000 re S
		0.000 0.000 0.000 RG
		28.346 605.197 283.465 30.000 re S
		BT /FNT2 10 Tf ET
		BT 142 625 Td (Helvetica 10) Tj ET
		0.745 0.745 0.745 RG
		28.346 595.197 283.465 10.000 re S
		28.346 585.197 283.465 10.000 re S
		28.346 575.197 283.465 10.000 re S
		0.000 0.000 0.000 RG
		28.346 575.197 283.465 30.000 re S
		BT /FNT2 15 Tf ET
		BT 128 590 Td (Helvetica 15) Tj ET
		0.745 0.745 0.745 RG
		28.346 565.197 283.465 10.000 re S
		28.346 555.197 283.465 10.000 re S
		28.346 545.197 283.465 10.000 re S
		0.000 0.000 0.000 RG
		28.346 545.197 283.465 30.000 re S
		BT /FNT2 20 Tf ET
		BT 115 555 Td (Helvetica 20) Tj ET
		0.745 0.745 0.745 RG
		28.346 535.197 283.465 10.000 re S
		28.346 525.197 283.465 10.000 re S
		28.346 515.197 283.465 10.000 re S
		0.000 0.000 0.000 RG
		28.346 515.197 283.465 30.000 re S
		BT /FNT2 25 Tf ET
		BT 101 520 Td (Helvetica 25) Tj ET
		0.745 0.745 0.745 RG
		28.346 505.197 283.465 10.000 re S
		28.346 495.197 283.465 10.000 re S
		28.346 485.197 283.465 10.000 re S
		0.000 0.000 0.000 RG
		28.346 485.197 283.465 30.000 re S
		BT /FNT2 30 Tf ET
		BT 87 485 Td (Helvetica 30) Tj ET
		endstream
		endobj
		5 0 obj <</Type/Font/Subtype/Type1/Name/FNT1
		/BaseFont/Times-Bold
		/Encoding/StandardEncoding>>
		endobj
		6 0 obj <</Type/Font/Subtype/Type1/Name/FNT2
		/BaseFont/Helvetica
		/Encoding/StandardEncoding>>
		endobj
		xref
		0 7
		0000000000 65535 f
		0000000010 00000 n
		0000000056 00000 n
		0000000130 00000 n
		0000000241 00000 n
		0000002733 00000 n
		0000002835 00000 n
		trailer
		<</Size 7/Root 1 0 R>>
		startxref
		2936
		%%EOF
		`
		pdfCompare(t, doc.Bytes(), want)
		// doc.SaveFile("~~Test_PDF_FontSize_.pdf")
	}()
} //                                                          Test_PDF_FontSize_

// Test_PDF_HorizontalScaling_ is the unit test for PDF.HorizontalScaling()
func Test_PDF_HorizontalScaling_(t *testing.T) {
	//
	// Horizontal Scaling of new PDF must be 100
	func() {
		var doc PDF
		tEqual(t, doc.HorizontalScaling(), 100)
	}()
	func() {
		doc := NewPDF("A4")
		tEqual(t, doc.HorizontalScaling(), 100)
	}()
	//
	// SetHorizontalScaling() has effect on the property?
	func() {
		var doc PDF
		doc.SetHorizontalScaling(149)
		tEqual(t, doc.HorizontalScaling(), 149)
	}()
	func() {
		doc := NewPDF("A4")
		doc.SetHorizontalScaling(149)
		tEqual(t, doc.HorizontalScaling(), 149)
	}()
	//
	// Test PDF generation
	//
	func() {
		doc := NewPDF("A4-L")
		{
			doc.SetCompression(false).
				SetUnits("cm").
				SetFont("Times-Bold", 20).
				SetXY(1, 1).
				DrawText("Horizontal Scaling Property")
		}
		for i, hscaling := range []int{50, 100, 150, 200, 250} {
			y := 2.5 + float64(i)*2.5
			doc.SetXY(1, y).
				SetFont("Helvetica", 10).
				SetHorizontalScaling(100).
				DrawText(fmt.Sprintf("Horizontal Scaling = %d", hscaling)).
				SetXY(1, y+0.7).
				SetHorizontalScaling(uint16(hscaling)).
				SetFontSize(20).
				DrawText("Five hexing wizard bots jump quickly")
		}
		const want = `
		%PDF-1.4
		1 0 obj <</Type/Catalog/Pages 2 0 R>>
		endobj
		2 0 obj <</Type/Pages/Count 1/MediaBox[0 0 841 595]/Kids[3 0 R]>>
		endobj
		3 0 obj <</Type/Page/Parent 2 0 R/Contents 4 0 R
		/Resources <</Font <<
		/FNT1 5 0 R
		/FNT2 6 0 R>> >> >>
		endobj
		4 0 obj <</Length 899>> stream
		BT /FNT1 20 Tf ET
		0.000 0.000 0.000 rg
		0.000 0.000 0.000 RG
		BT 28 566 Td (Horizontal Scaling Property) Tj ET
		BT /FNT2 10 Tf ET
		BT 28 524 Td (Horizontal Scaling = 50) Tj ET
		BT /FNT2 20 Tf ET
		BT 50 Tz ET
		BT 28 504 Td (Five hexing wizard bots jump quickly) Tj ET
		BT /FNT2 10 Tf ET
		BT 100 Tz ET
		BT 28 453 Td (Horizontal Scaling = 100) Tj ET
		BT /FNT2 20 Tf ET
		BT 28 433 Td (Five hexing wizard bots jump quickly) Tj ET
		BT /FNT2 10 Tf ET
		BT 28 382 Td (Horizontal Scaling = 150) Tj ET
		BT /FNT2 20 Tf ET
		BT 150 Tz ET
		BT 28 362 Td (Five hexing wizard bots jump quickly) Tj ET
		BT /FNT2 10 Tf ET
		BT 100 Tz ET
		BT 28 311 Td (Horizontal Scaling = 200) Tj ET
		BT /FNT2 20 Tf ET
		BT 200 Tz ET
		BT 28 291 Td (Five hexing wizard bots jump quickly) Tj ET
		BT /FNT2 10 Tf ET
		BT 100 Tz ET
		BT 28 240 Td (Horizontal Scaling = 250) Tj ET
		BT /FNT2 20 Tf ET
		BT 250 Tz ET
		BT 28 221 Td (Five hexing wizard bots jump quickly) Tj ET
		endstream
		endobj
		5 0 obj <</Type/Font/Subtype/Type1/Name/FNT1
		/BaseFont/Times-Bold
		/Encoding/StandardEncoding>>
		endobj
		6 0 obj <</Type/Font/Subtype/Type1/Name/FNT2
		/BaseFont/Helvetica
		/Encoding/StandardEncoding>>
		endobj
		xref
		0 7
		0000000000 65535 f
		0000000010 00000 n
		0000000056 00000 n
		0000000130 00000 n
		0000000241 00000 n
		0000001191 00000 n
		0000001293 00000 n
		trailer
		<</Size 7/Root 1 0 R>>
		startxref
		1394
		%%EOF
		`
		pdfCompare(t, doc.Bytes(), want)
		doc.SaveFile("~~test_horizontal_scaling.pdf")
	}()
} //                                                 Test_PDF_HorizontalScaling_

// Test_PDF_LineWidth_ is the unit test for PDF.LineWidth()
// go test --run Test_PDF_LineWidth_
func Test_PDF_LineWidth_(t *testing.T) {
	//
	// LineWidth of new PDF must be 1 point
	func() {
		var doc PDF
		tEqual(t, doc.LineWidth(), 1)
	}()
	func() {
		doc := NewPDF("A4")
		tEqual(t, doc.LineWidth(), 1)
	}()
	//
	// SetLineWidth() has effect on the property?
	func() {
		var doc PDF
		doc.SetLineWidth(42)
		tEqual(t, doc.LineWidth(), 42)
	}()
	func() {
		doc := NewPDF("A4")
		doc.SetLineWidth(7)
		tEqual(t, doc.LineWidth(), 7)
	}()
	//
	// Test PDF generation
	//
	func() {
		doc := NewPDF("A4")
		doc.SetCompression(false).
			SetUnits("cm").
			//DrawUnitGrid().
			SetXY(1, 1.5).SetColor("Indigo").
			SetFont("Helvetica", 16).
			DrawText("Test PDF.LineWidth()").
			SetFont("Helvetica", 9)
		y := 2.0
		for _, w := range []float64{0.1, 0.2, 0.5, 1, 5, 10, 15, 20, 25} {
			doc.SetColor("Dark Gray").
				SetXY(1.0, y+0.3).DrawText(" y = "+strconv.Itoa(int(y))).
				SetXY(2.3, y+0.3).DrawText(" w = "+fmt.Sprintf("%0.1f", w)).
				SetColor("Gray").SetLineWidth(0.1).DrawLine(1, y, 20, y).
				SetColor("Indigo").SetLineWidth(w).DrawLine(4, y, 15, y)
			y += 1
		}
		const want = `
		%PDF-1.4
		1 0 obj <</Type/Catalog/Pages 2 0 R>>
		endobj
		2 0 obj <</Type/Pages/Count 1/MediaBox[0 0 595 841]/Kids[3 0 R]>>
		endobj
		3 0 obj <</Type/Page/Parent 2 0 R/Contents 4 0 R
		/Resources <</Font <</FNT1 5 0 R>> >> >>
		endobj
		4 0 obj <</Length 2617>> stream
		BT /FNT1 16 Tf ET
		0.294 0.000 0.510 rg
		0.294 0.000 0.510 RG
		BT 28 799 Td (Test PDF.LineWidth\(\)) Tj ET
		BT /FNT1 9 Tf ET
		0.663 0.663 0.663 rg
		0.663 0.663 0.663 RG
		BT 28 776 Td ( y = 2) Tj ET
		BT 65 776 Td ( w = 0.1) Tj ET
		0.745 0.745 0.745 rg
		0.745 0.745 0.745 RG
		0.100 w
		28.346 785.197 m 566.929 785.197 l S
		0.294 0.000 0.510 rg
		0.294 0.000 0.510 RG
		113.386 785.197 m 425.197 785.197 l S
		0.663 0.663 0.663 rg
		0.663 0.663 0.663 RG
		BT 28 748 Td ( y = 3) Tj ET
		BT 65 748 Td ( w = 0.2) Tj ET
		0.745 0.745 0.745 rg
		0.745 0.745 0.745 RG
		28.346 756.850 m 566.929 756.850 l S
		0.294 0.000 0.510 rg
		0.294 0.000 0.510 RG
		0.200 w
		113.386 756.850 m 425.197 756.850 l S
		0.663 0.663 0.663 rg
		0.663 0.663 0.663 RG
		BT 28 720 Td ( y = 4) Tj ET
		BT 65 720 Td ( w = 0.5) Tj ET
		0.745 0.745 0.745 rg
		0.745 0.745 0.745 RG
		0.100 w
		28.346 728.504 m 566.929 728.504 l S
		0.294 0.000 0.510 rg
		0.294 0.000 0.510 RG
		0.500 w
		113.386 728.504 m 425.197 728.504 l S
		0.663 0.663 0.663 rg
		0.663 0.663 0.663 RG
		BT 28 691 Td ( y = 5) Tj ET
		BT 65 691 Td ( w = 1.0) Tj ET
		0.745 0.745 0.745 rg
		0.745 0.745 0.745 RG
		0.100 w
		28.346 700.157 m 566.929 700.157 l S
		0.294 0.000 0.510 rg
		0.294 0.000 0.510 RG
		1.000 w
		113.386 700.157 m 425.197 700.157 l S
		0.663 0.663 0.663 rg
		0.663 0.663 0.663 RG
		BT 28 663 Td ( y = 6) Tj ET
		BT 65 663 Td ( w = 5.0) Tj ET
		0.745 0.745 0.745 rg
		0.745 0.745 0.745 RG
		0.100 w
		28.346 671.811 m 566.929 671.811 l S
		0.294 0.000 0.510 rg
		0.294 0.000 0.510 RG
		5.000 w
		113.386 671.811 m 425.197 671.811 l S
		0.663 0.663 0.663 rg
		0.663 0.663 0.663 RG
		BT 28 634 Td ( y = 7) Tj ET
		BT 65 634 Td ( w = 10.0) Tj ET
		0.745 0.745 0.745 rg
		0.745 0.745 0.745 RG
		0.100 w
		28.346 643.465 m 566.929 643.465 l S
		0.294 0.000 0.510 rg
		0.294 0.000 0.510 RG
		10.000 w
		113.386 643.465 m 425.197 643.465 l S
		0.663 0.663 0.663 rg
		0.663 0.663 0.663 RG
		BT 28 606 Td ( y = 8) Tj ET
		BT 65 606 Td ( w = 15.0) Tj ET
		0.745 0.745 0.745 rg
		0.745 0.745 0.745 RG
		0.100 w
		28.346 615.118 m 566.929 615.118 l S
		0.294 0.000 0.510 rg
		0.294 0.000 0.510 RG
		15.000 w
		113.386 615.118 m 425.197 615.118 l S
		0.663 0.663 0.663 rg
		0.663 0.663 0.663 RG
		BT 28 578 Td ( y = 9) Tj ET
		BT 65 578 Td ( w = 20.0) Tj ET
		0.745 0.745 0.745 rg
		0.745 0.745 0.745 RG
		0.100 w
		28.346 586.772 m 566.929 586.772 l S
		0.294 0.000 0.510 rg
		0.294 0.000 0.510 RG
		20.000 w
		113.386 586.772 m 425.197 586.772 l S
		0.663 0.663 0.663 rg
		0.663 0.663 0.663 RG
		BT 28 549 Td ( y = 10) Tj ET
		BT 65 549 Td ( w = 25.0) Tj ET
		0.745 0.745 0.745 rg
		0.745 0.745 0.745 RG
		0.100 w
		28.346 558.425 m 566.929 558.425 l S
		0.294 0.000 0.510 rg
		0.294 0.000 0.510 RG
		25.000 w
		113.386 558.425 m 425.197 558.425 l S
		endstream
		endobj
		5 0 obj <</Type/Font/Subtype/Type1/Name/FNT1
		/BaseFont/Helvetica
		/Encoding/StandardEncoding>>
		endobj
		xref
		0 6
		0000000000 65535 f
		0000000010 00000 n
		0000000056 00000 n
		0000000130 00000 n
		0000000228 00000 n
		0000002897 00000 n
		trailer
		<</Size 6/Root 1 0 R>>
		startxref
		2998
		%%EOF
        `
		pdfCompare(t, doc.Bytes(), want)
		// doc.SaveFile("~~Test_PDF_LineWidth_.pdf")
	}()
} //                                                         Test_PDF_LineWidth_

// Test_NewPDF_ is the unit test for PDF.NewPDF
func Test_NewPDF_(t *testing.T) {
	const want = `
	%PDF-1.4
	1 0 obj <</Type/Catalog/Pages 2 0 R>>
	endobj
	2 0 obj <</Type/Pages/Count 1/MediaBox[0 0 595 841]/Kids[3 0 R]>>
	endobj
	3 0 obj <</Type/Page/Parent 2 0 R/Contents 4 0 R>>
	endobj
	4 0 obj <</Length 0>> stream
	endstream
	endobj
	xref
	0 5
	0000000000 65535 f
	0000000010 00000 n
	0000000056 00000 n
	0000000130 00000 n
	0000000189 00000 n
	trailer
	<</Size 5/Root 1 0 R>>
	startxref
	238
	%%EOF
	`
	// test NewPDF() and Bytes() while calling AddPage()
	func() {
		doc := NewPDF("A4")
		got := doc.SetCompression(false).AddPage().Bytes()
		pdfCompare(t, got, want)
	}()
	// test NewPDF() and Bytes() without calling AddPage()
	func() {
		doc := NewPDF("A4")
		got := doc.SetCompression(false).Bytes()
		pdfCompare(t, got, want)
	}()
} //                                                                Test_NewPDF_

// Test_PDF_PageCount_ tests PDF.PageCount()
func Test_PDF_PageCount_(t *testing.T) {
	//                                                 uninitialized PDF: 1 page
	func() {
		var doc PDF
		tEqual(t, doc.PageCount(), 1)
	}()
	//                  uninitialized PDF with initial call to AddPage(): 1 page
	func() {
		var doc PDF
		doc.AddPage()
		tEqual(t, doc.PageCount(), 1)
	}()
	//                                                   initialized PDF: 1 page
	func() {
		doc := NewPDF("A4")
		tEqual(t, doc.PageCount(), 1)
	}()
	//                    initialized PDF with initial call to AddPage(): 1 page
	func() {
		doc := NewPDF("A4")
		doc.AddPage()
		tEqual(t, doc.PageCount(), 1)
	}()
	//                              initialized PDF with a single method: 1 page
	func() {
		doc := NewPDF("LETTER")
		doc.SetXY(1, 1)
		tEqual(t, doc.PageCount(), 1)
	}()
	//                  calling AddPage() after any method, increases page count
	func() {
		var doc PDF //                                     uninitialized PDF
		doc.SetXY(1, 1)
		doc.AddPage()
		tEqual(t, doc.PageCount(), 2)
	}()
	func() {
		doc := NewPDF("LETTER") //                           initialized PDF
		doc.SetXY(1, 1)
		doc.AddPage()
		tEqual(t, doc.PageCount(), 2)
	}()
	//                  after calling AddPage() 10 times, PageCount() must be 10
	func() {
		var doc PDF //                                     uninitialized PDF
		for i := 0; i < 10; i++ {
			doc.AddPage()
		}
		tEqual(t, doc.PageCount(), 10)
	}()
	func() {
		doc := NewPDF("LETTER") //                           initialized PDF
		for i := 0; i < 10; i++ {
			doc.AddPage()
		}
		tEqual(t, doc.PageCount(), 10)
	}()
} //                                                         Test_PDF_PageCount_

// Test_PDF_PageHeight_ tests PDF.PageHeight()
func Test_PDF_PageHeight_(t *testing.T) {
	func() {
		var doc PDF // uninitialized PDF
		tEqual(t, doc.PageHeight(), 0.0)
	}()
	//
	// 2.83464566929134 points per mm
	// 1 inch / 25.4mm per inch * 72 points per inch
	//
	func() {
		doc := NewPDF("A4")                     // initialized PDF
		tEqual(t, doc.PageHeight(), 841.889764) // points
		//
		// A4 = 210mm width x 297mm height = 841.8897637795276 points
	}()
	func() {
		doc := NewPDF("LETTER")                 // initialized PDF
		tEqual(t, doc.PageHeight(), 790.866142) // points
		//
		// LETTER = 216mm width x 279mm height = 790.8661417322835 points
	}()
} //                                                        Test_PDF_PageHeight_

// Test_PDF_PageWidth_ tests PDF.PageWidth()
func Test_PDF_PageWidth_(t *testing.T) {
	func() {
		var doc PDF // uninitialized PDF
		tEqual(t, doc.PageWidth(), 0.0)
	}()
	//
	// 2.83464566929134 points per mm
	// 1 inch / 25.4mm per inch * 72 points per inch
	//
	func() {
		doc := NewPDF("A4")                    // initialized PDF
		tEqual(t, doc.PageWidth(), 595.275591) // points
		//
		// A4 = 210mm width x 297mm height = 595.2755905511811 points
	}()
	func() {
		doc := NewPDF("LETTER")                // initialized PDF
		tEqual(t, doc.PageWidth(), 612.283465) // points
		//
		// LETTER = 216mm width x 279mm height = 612.2834645669291 points
	}()
} //                                                         Test_PDF_PageWidth_

// Test_PDF_PullError_ is the unit test for PullError() error
func Test_PDF_PullError_(t *testing.T) {
	func() {
		doc := NewPDF("Papyrus")
		//
		// errors should have one error
		tEqual(t, len(doc.Errors()), 1)
		//
		// fetch and remove this error from Errors()
		err := doc.PullError()
		tEqual(t, err, fmt.Errorf(`Unknown paper size "Papyrus" @NewPDF`))
		//
		// Errors() should now be empty
		tEqual(t, len(doc.Errors()), 0)
		//
		// if we try to pull another error (there is none), 'err' will be nil
		err = doc.PullError()
		tEqual(t, len(doc.Errors()), 0)
		tEqual(t, err, nil)
	}()
} //                                                         Test_PDF_PullError_

// Test_PDF_Reset_ tests PDF.Reset()
func Test_PDF_Reset_(t *testing.T) {
	//
	// prepare a test PDF before calling Reset()
	doc := NewPDF("A4")
	{
		doc.SetCompression(false).
			SetUnits("cm").
			SetColumnWidths(1, 4, 9).
			SetColor("#006B3C CadmiumGreen").
			SetFont("Helvetica-Bold", 10).
			SetX(5).
			SetY(5).
			DrawText("FIRST").
			DrawText("SECOND").
			DrawText("THIRD")
		const want = `
		%PDF-1.4
		1 0 obj <</Type/Catalog/Pages 2 0 R>>
		endobj
		2 0 obj <</Type/Pages/Count 1/MediaBox[0 0 595 841]/Kids[3 0 R]>>
		endobj
		3 0 obj <</Type/Page/Parent 2 0 R/Contents 4 0 R
		/Resources <</Font <</FNT1 5 0 R>> >> >>
		endobj
		4 0 obj <</Length 143>> stream
		BT /FNT1 10 Tf ET
		0.000 0.420 0.235 rg
		0.000 0.420 0.235 RG
		BT 0 700 Td (FIRST) Tj ET
		BT 28 700 Td (SECOND) Tj ET
		BT 141 700 Td (THIRD) Tj ET
		endstream
		endobj
		5 0 obj <</Type/Font/Subtype/Type1/Name/FNT1
		/BaseFont/Helvetica-Bold
		/Encoding/StandardEncoding>>
		endobj
		xref
		0 6
		0000000000 65535 f
		0000000010 00000 n
		0000000056 00000 n
		0000000130 00000 n
		0000000228 00000 n
		0000000422 00000 n
		trailer
		<</Size 6/Root 1 0 R>>
		startxref
		528
		%%EOF
		`
		got := doc.Bytes()
		pdfCompare(t, got, want)
	}
	{
		doc.Reset()
		//
		// after calling Reset(), the PDF should just be a blank page:
		const want = `
		%PDF-1.4
		1 0 obj <</Type/Catalog/Pages 2 0 R>>
		endobj
		2 0 obj <</Type/Pages/Count 1/MediaBox[0 0 595 841]/Kids[3 0 R]>>
		endobj
		3 0 obj <</Type/Page/Parent 2 0 R/Contents 4 0 R>>
		endobj
		4 0 obj <</Length 0>> stream
		endstream
		endobj
		xref
		0 5
		0000000000 65535 f
		0000000010 00000 n
		0000000056 00000 n
		0000000130 00000 n
		0000000189 00000 n
		trailer
		<</Size 5/Root 1 0 R>>
		startxref
		238
		%%EOF
		`
		got := doc.SetCompression(false).Bytes()
		pdfCompare(t, got, want)
	}
	// TODO: add more test cases, test each property's state
} //                                                             Test_PDF_Reset_

// Test_PDF_SetFont_ is the unit test for PDF.SetFont()
func Test_PDF_SetFont_(t *testing.T) {
	//
	// setting font on uninitialized pdf
	for _, tc := range []struct {
		size    float64
		inName  string
		expName string
	}{
		{size: 10, inName: "Courier", expName: "Courier"},
		{size: 20, inName: "Zapf-Dingbats", expName: "Zapf-Dingbats"},
		{size: 30, inName: "ZapfDingbats", expName: "ZapfDingbats"},
		{size: 40, inName: "YeOldeScript", expName: "YeOldeScript"},
	} {
		var doc PDF // uninitialized PDF
		doc.SetFont(tc.inName, tc.size)
		tEqual(t, doc.FontName(), tc.expName)
		tEqual(t, doc.FontSize(), tc.size)
	}
	//
	// Test PDF generation
	//
	func() {
		doc := NewPDF("A4")
		doc.SetCompression(false).
			SetUnits("cm").
			SetFont("Times-Bold", 20).
			SetXY(1, 1).
			DrawText("Built-in PDF Fonts")
		for i, font := range []string{
			"Courier",
			"Courier-Bold",
			"Courier-BoldOblique",
			"Courier-Oblique",
			"Helvetica",
			"Helvetica-Bold",
			"Helvetica-BoldOblique",
			"Helvetica-Oblique",
			"Symbol",
			"Times-Bold",
			"Times-BoldItalic",
			"Times-Italic",
			"Times-Roman",
			"ZapfDingbats",
		} {
			y := 2.5 + float64(i)*1.8
			doc.SetXY(1, y).
				SetFont("Helvetica", 10).
				DrawText(font).
				SetXY(1, y+0.7).
				SetFont(font, 20).
				DrawText("Five hexing wizard bots jump quickly")

		}
		const want = `
		%PDF-1.4
		1 0 obj <</Type/Catalog/Pages 2 0 R>>
		endobj
		2 0 obj <</Type/Pages/Count 1/MediaBox[0 0 595 841]/Kids[3 0 R]>>
		endobj
		3 0 obj <</Type/Page/Parent 2 0 R/Contents 4 0 R
		/Resources <</Font <<
		/FNT1 5 0 R
		/FNT2 6 0 R
		/FNT3 7 0 R
		/FNT4 8 0 R
		/FNT5 9 0 R
		/FNT6 10 0 R
		/FNT7 11 0 R
		/FNT8 12 0 R
		/FNT9 13 0 R
		/FNT10 14 0 R
		/FNT11 15 0 R
		/FNT12 16 0 R
		/FNT13 17 0 R
		/FNT14 18 0 R>> >> >>
		endobj
		4 0 obj <</Length 1910>> stream
		BT /FNT1 20 Tf ET
		0.000 0.000 0.000 rg
		0.000 0.000 0.000 RG
		BT 28 813 Td (Built-in PDF Fonts) Tj ET
		BT /FNT2 10 Tf ET
		BT 28 771 Td (Courier) Tj ET
		BT /FNT3 20 Tf ET
		BT 28 751 Td (Five hexing wizard bots jump quickly) Tj ET
		BT /FNT2 10 Tf ET
		BT 28 720 Td (Courier-Bold) Tj ET
		BT /FNT4 20 Tf ET
		BT 28 700 Td (Five hexing wizard bots jump quickly) Tj ET
		BT /FNT2 10 Tf ET
		BT 28 668 Td (Courier-BoldOblique) Tj ET
		BT /FNT5 20 Tf ET
		BT 28 649 Td (Five hexing wizard bots jump quickly) Tj ET
		BT /FNT2 10 Tf ET
		BT 28 617 Td (Courier-Oblique) Tj ET
		BT /FNT6 20 Tf ET
		BT 28 598 Td (Five hexing wizard bots jump quickly) Tj ET
		BT /FNT2 10 Tf ET
		BT 28 566 Td (Helvetica) Tj ET
		BT /FNT2 20 Tf ET
		BT 28 547 Td (Five hexing wizard bots jump quickly) Tj ET
		BT /FNT2 10 Tf ET
		BT 28 515 Td (Helvetica-Bold) Tj ET
		BT /FNT7 20 Tf ET
		BT 28 496 Td (Five hexing wizard bots jump quickly) Tj ET
		BT /FNT2 10 Tf ET
		BT 28 464 Td (Helvetica-BoldOblique) Tj ET
		BT /FNT8 20 Tf ET
		BT 28 445 Td (Five hexing wizard bots jump quickly) Tj ET
		BT /FNT2 10 Tf ET
		BT 28 413 Td (Helvetica-Oblique) Tj ET
		BT /FNT9 20 Tf ET
		BT 28 394 Td (Five hexing wizard bots jump quickly) Tj ET
		BT /FNT2 10 Tf ET
		BT 28 362 Td (Symbol) Tj ET
		BT /FNT10 20 Tf ET
		BT 28 342 Td (Five hexing wizard bots jump quickly) Tj ET
		BT /FNT2 10 Tf ET
		BT 28 311 Td (Times-Bold) Tj ET
		BT /FNT1 20 Tf ET
		BT 28 291 Td (Five hexing wizard bots jump quickly) Tj ET
		BT /FNT2 10 Tf ET
		BT 28 260 Td (Times-BoldItalic) Tj ET
		BT /FNT11 20 Tf ET
		BT 28 240 Td (Five hexing wizard bots jump quickly) Tj ET
		BT /FNT2 10 Tf ET
		BT 28 209 Td (Times-Italic) Tj ET
		BT /FNT12 20 Tf ET
		BT 28 189 Td (Five hexing wizard bots jump quickly) Tj ET
		BT /FNT2 10 Tf ET
		BT 28 158 Td (Times-Roman) Tj ET
		BT /FNT13 20 Tf ET
		BT 28 138 Td (Five hexing wizard bots jump quickly) Tj ET
		BT /FNT2 10 Tf ET
		BT 28 107 Td (ZapfDingbats) Tj ET
		BT /FNT14 20 Tf ET
		BT 28 87 Td (Five hexing wizard bots jump quickly) Tj ET
		endstream
		endobj
		5 0 obj <</Type/Font/Subtype/Type1/Name/FNT1
		/BaseFont/Times-Bold
		/Encoding/StandardEncoding>>
		endobj
		6 0 obj <</Type/Font/Subtype/Type1/Name/FNT2
		/BaseFont/Helvetica
		/Encoding/StandardEncoding>>
		endobj
		7 0 obj <</Type/Font/Subtype/Type1/Name/FNT3
		/BaseFont/Courier
		/Encoding/StandardEncoding>>
		endobj
		8 0 obj <</Type/Font/Subtype/Type1/Name/FNT4
		/BaseFont/Courier-Bold
		/Encoding/StandardEncoding>>
		endobj
		9 0 obj <</Type/Font/Subtype/Type1/Name/FNT5
		/BaseFont/Courier-BoldOblique
		/Encoding/StandardEncoding>>
		endobj
		10 0 obj <</Type/Font/Subtype/Type1/Name/FNT6
		/BaseFont/Courier-Oblique
		/Encoding/StandardEncoding>>
		endobj
		11 0 obj <</Type/Font/Subtype/Type1/Name/FNT7
		/BaseFont/Helvetica-Bold
		/Encoding/StandardEncoding>>
		endobj
		12 0 obj <</Type/Font/Subtype/Type1/Name/FNT8
		/BaseFont/Helvetica-BoldOblique
		/Encoding/StandardEncoding>>
		endobj
		13 0 obj <</Type/Font/Subtype/Type1/Name/FNT9
		/BaseFont/Helvetica-Oblique
		/Encoding/StandardEncoding>>
		endobj
		14 0 obj <</Type/Font/Subtype/Type1/Name/FNT10
		/BaseFont/Symbol
		/Encoding/StandardEncoding>>
		endobj
		15 0 obj <</Type/Font/Subtype/Type1/Name/FNT11
		/BaseFont/Times-BoldItalic
		/Encoding/StandardEncoding>>
		endobj
		16 0 obj <</Type/Font/Subtype/Type1/Name/FNT12
		/BaseFont/Times-Italic
		/Encoding/StandardEncoding>>
		endobj
		17 0 obj <</Type/Font/Subtype/Type1/Name/FNT13
		/BaseFont/Times-Roman
		/Encoding/StandardEncoding>>
		endobj
		18 0 obj <</Type/Font/Subtype/Type1/Name/FNT14
		/BaseFont/ZapfDingbats
		/Encoding/StandardEncoding>>
		endobj
		xref
		0 19
		0000000000 65535 f
		0000000010 00000 n
		0000000056 00000 n
		0000000130 00000 n
		0000000399 00000 n
		0000002361 00000 n
		0000002463 00000 n
		0000002564 00000 n
		0000002663 00000 n
		0000002767 00000 n
		0000002878 00000 n
		0000002986 00000 n
		0000003093 00000 n
		0000003207 00000 n
		0000003317 00000 n
		0000003417 00000 n
		0000003527 00000 n
		0000003633 00000 n
		0000003738 00000 n
		trailer
		<</Size 19/Root 1 0 R>>
		startxref
		3844
		%%EOF
		`
		pdfCompare(t, doc.Bytes(), want)
		doc.SaveFile("~~font_sample.pdf")
	}()
} //                                                           Test_PDF_SetFont_

// Test_PDF_SetXY_ is the unit test for PDF.SetXY()
func Test_PDF_SetXY_(t *testing.T) {
	//
	// SetXY() sets X and Y properties properly?
	func() {
		var doc PDF
		doc.SetXY(123, 456)
		tEqual(t, doc.X(), 123)
		tEqual(t, doc.Y(), 456)
	}()
	func() {
		doc := NewPDF("A4")
		doc.SetXY(123, 456)
		tEqual(t, doc.X(), 123)
		tEqual(t, doc.Y(), 456)
	}()
	//
	// Test PDF generation
	//
	func() {
		doc := NewPDF("A4")
		doc.SetCompression(false).SetUnits("cm").SetFont("Helvetica", 10).
			SetXY(1, 3).DrawText("X=1cm Y=3cm").
			SetXY(3, 1).DrawText("X=3cm Y=1cm").
			SetXY(10, 5).DrawText("X=10cm Y=5cm").
			SetXY(5, 10).DrawText("X=5cm Y=10cm")
		const want = `
		%PDF-1.4
		1 0 obj <</Type/Catalog/Pages 2 0 R>>
		endobj
		2 0 obj <</Type/Pages/Count 1/MediaBox[0 0 595 841]/Kids[3 0 R]>>
		endobj
		3 0 obj <</Type/Page/Parent 2 0 R/Contents 4 0 R
		/Resources <</Font <</FNT1 5 0 R>> >> >>
		endobj
		4 0 obj <</Length 197>> stream
		BT /FNT1 10 Tf ET
		0.000 0.000 0.000 rg
		0.000 0.000 0.000 RG
		BT 28 756 Td (X=1cm Y=3cm) Tj ET
		BT 85 813 Td (X=3cm Y=1cm) Tj ET
		BT 283 700 Td (X=10cm Y=5cm) Tj ET
		BT 141 558 Td (X=5cm Y=10cm) Tj ET
		endstream
		endobj
		5 0 obj <</Type/Font/Subtype/Type1/Name/FNT1
		/BaseFont/Helvetica
		/Encoding/StandardEncoding>>
		endobj
		xref
		0 6
		0000000000 65535 f
		0000000010 00000 n
		0000000056 00000 n
		0000000130 00000 n
		0000000228 00000 n
		0000000476 00000 n
		trailer
		<</Size 6/Root 1 0 R>>
		startxref
		577
		%%EOF
		`
		pdfCompare(t, doc.Bytes(), want)
		doc.SaveFile("~~Test_PDF_SetXY_.pdf")
	}()
} //                                                             Test_PDF_SetXY_

// Test_PDF_ToColor_1_ is the unit test for
// (ob *PDF) ToColor(nameOrHTMLColor string) (color.RGBA, error)
func Test_PDF_ToColor_1_(t *testing.T) {
	func() {
		var doc PDF
		got, err := doc.ToColor("")
		tEqual(t, got, color.RGBA{A: 255}) // black
		// error is returned in `err`, but does not affect Errors()
		tEqual(t, len(doc.Errors()), 0)
		tEqual(t, err.Error(),
			fmt.Errorf(`Unknown color name "" @ToColor`).Error())
	}()
	func() {
		var doc PDF
		got, err := doc.ToColor("#uvwxyz")
		tEqual(t, got, color.RGBA{A: 255}) // black
		// error is returned in `err`, but does not affect Errors()
		tEqual(t, len(doc.Errors()), 0)
		tEqual(t, err.Error(),
			fmt.Errorf(`Bad color code "#uvwxyz" @ToColor`).Error())
	}()
	func() {
		// map copied from PDFColorNames, but color names in lower case
		m := map[string]color.RGBA{
			"aliceblue":            {R: 240, G: 248, B: 255}, // #F0F8FF
			"antiquewhite":         {R: 250, G: 235, B: 215}, // #FAEBD7
			"aqua":                 {R: 000, G: 255, B: 255}, // #00FFFF
			"aquamarine":           {R: 127, G: 255, B: 212}, // #7FFFD4
			"azure":                {R: 240, G: 255, B: 255}, // #F0FFFF
			"beige":                {R: 245, G: 245, B: 220}, // #F5F5DC
			"bisque":               {R: 255, G: 228, B: 196}, // #FFE4C4
			"black":                {R: 000, G: 000, B: 000}, // #000000
			"blanchedalmond":       {R: 255, G: 235, B: 205}, // #FFEBCD
			"blue":                 {R: 000, G: 000, B: 255}, // #0000FF
			"blueviolet":           {R: 138, G: 43, B: 226},  // #8A2BE2
			"brown":                {R: 165, G: 42, B: 42},   // #A52A2A
			"burlywood":            {R: 222, G: 184, B: 135}, // #DEB887
			"cadetblue":            {R: 95, G: 158, B: 160},  // #5F9EA0
			"chartreuse":           {R: 127, G: 255, B: 000}, // #7FFF00
			"chocolate":            {R: 210, G: 105, B: 30},  // #D2691E
			"coral":                {R: 255, G: 127, B: 80},  // #FF7F50
			"cornflowerblue":       {R: 100, G: 149, B: 237}, // #6495ED
			"cornsilk":             {R: 255, G: 248, B: 220}, // #FFF8DC
			"crimson":              {R: 220, G: 20, B: 60},   // #DC143C
			"cyan":                 {R: 000, G: 255, B: 255}, // #00FFFF
			"darkblue":             {R: 000, G: 000, B: 139}, // #00008B
			"darkcyan":             {R: 000, G: 139, B: 139}, // #008B8B
			"darkgoldenrod":        {R: 184, G: 134, B: 11},  // #B8860B
			"darkgray":             {R: 169, G: 169, B: 169}, // #A9A9A9
			"darkgreen":            {R: 000, G: 100, B: 000}, // #006400
			"darkkhaki":            {R: 189, G: 183, B: 107}, // #BDB76B
			"darkmagenta":          {R: 139, G: 000, B: 139}, // #8B008B
			"darkolivegreen":       {R: 85, G: 107, B: 47},   // #556B2F
			"darkorange":           {R: 255, G: 140, B: 000}, // #FF8C00
			"darkorchid":           {R: 153, G: 50, B: 204},  // #9932CC
			"darkred":              {R: 139, G: 000, B: 000}, // #8B0000
			"darksalmon":           {R: 233, G: 150, B: 122}, // #E9967A
			"darkseagreen":         {R: 143, G: 188, B: 143}, // #8FBC8F
			"darkslateblue":        {R: 72, G: 61, B: 139},   // #483D8B
			"darkslategray":        {R: 47, G: 79, B: 79},    // #2F4F4F
			"darkturquoise":        {R: 000, G: 206, B: 209}, // #00CED1
			"darkviolet":           {R: 148, G: 000, B: 211}, // #9400D3
			"deeppink":             {R: 255, G: 20, B: 147},  // #FF1493
			"deepskyblue":          {R: 000, G: 191, B: 255}, // #00BFFF
			"dimgray":              {R: 105, G: 105, B: 105}, // #696969
			"dodgerblue":           {R: 30, G: 144, B: 255},  // #1E90FF
			"firebrick":            {R: 178, G: 34, B: 34},   // #B22222
			"floralwhite":          {R: 255, G: 250, B: 240}, // #FFFAF0
			"forestgreen":          {R: 34, G: 139, B: 34},   // #228B22
			"fuchsia":              {R: 255, G: 000, B: 255}, // #FF00FF
			"gainsboro":            {R: 220, G: 220, B: 220}, // #DCDCDC
			"ghostwhite":           {R: 248, G: 248, B: 255}, // #F8F8FF
			"gold":                 {R: 255, G: 215, B: 000}, // #FFD700
			"goldenrod":            {R: 218, G: 165, B: 32},  // #DAA520
			"gray":                 {R: 190, G: 190, B: 190}, // #BEBEBE
			"green":                {R: 000, G: 255, B: 000}, // #00FF00
			"greenyellow":          {R: 173, G: 255, B: 47},  // #ADFF2F
			"honeydew":             {R: 240, G: 255, B: 240}, // #F0FFF0
			"hotpink":              {R: 255, G: 105, B: 180}, // #FF69B4
			"indianred":            {R: 205, G: 92, B: 92},   // #CD5C5C
			"indigo":               {R: 75, G: 000, B: 130},  // #4B0082
			"ivory":                {R: 255, G: 255, B: 240}, // #FFFFF0
			"khaki":                {R: 240, G: 230, B: 140}, // #F0E68C
			"lavender":             {R: 230, G: 230, B: 250}, // #E6E6FA
			"lavenderblush":        {R: 255, G: 240, B: 245}, // #FFF0F5
			"lawngreen":            {R: 124, G: 252, B: 000}, // #7CFC00
			"lemonchiffon":         {R: 255, G: 250, B: 205}, // #FFFACD
			"lightblue":            {R: 173, G: 216, B: 230}, // #ADD8E6
			"lightcoral":           {R: 240, G: 128, B: 128}, // #F08080
			"lightcyan":            {R: 224, G: 255, B: 255}, // #E0FFFF
			"lightgoldenrodyellow": {R: 250, G: 250, B: 210}, // #FAFAD2
			"lightgray":            {R: 211, G: 211, B: 211}, // #D3D3D3
			"lightgreen":           {R: 144, G: 238, B: 144}, // #90EE90
			"lightpink":            {R: 255, G: 182, B: 193}, // #FFB6C1
			"lightsalmon":          {R: 255, G: 160, B: 122}, // #FFA07A
			"lightseagreen":        {R: 32, G: 178, B: 170},  // #20B2AA
			"lightskyblue":         {R: 135, G: 206, B: 250}, // #87CEFA
			"lightslategray":       {R: 119, G: 136, B: 153}, // #778899
			"lightsteelblue":       {R: 176, G: 196, B: 222}, // #B0C4DE
			"lightyellow":          {R: 255, G: 255, B: 224}, // #FFFFE0
			"lime":                 {R: 000, G: 255, B: 000}, // #00FF00
			"limegreen":            {R: 50, G: 205, B: 50},   // #32CD32
			"linen":                {R: 250, G: 240, B: 230}, // #FAF0E6
			"magenta":              {R: 255, G: 000, B: 255}, // #FF00FF
			"maroon":               {R: 176, G: 48, B: 96},   // #B03060
			"mediumaquamarine":     {R: 102, G: 205, B: 170}, // #66CDAA
			"mediumblue":           {R: 000, G: 000, B: 205}, // #0000CD
			"mediumorchid":         {R: 186, G: 85, B: 211},  // #BA55D3
			"mediumpurple":         {R: 147, G: 112, B: 219}, // #9370DB
			"mediumseagreen":       {R: 60, G: 179, B: 113},  // #3CB371
			"mediumslateblue":      {R: 123, G: 104, B: 238}, // #7B68EE
			"mediumspringgreen":    {R: 000, G: 250, B: 154}, // #00FA9A
			"mediumturquoise":      {R: 72, G: 209, B: 204},  // #48D1CC
			"mediumvioletred":      {R: 199, G: 21, B: 133},  // #C71585
			"midnightblue":         {R: 25, G: 25, B: 112},   // #191970
			"mintcream":            {R: 245, G: 255, B: 250}, // #F5FFFA
			"mistyrose":            {R: 255, G: 228, B: 225}, // #FFE4E1
			"moccasin":             {R: 255, G: 228, B: 181}, // #FFE4B5
			"navajowhite":          {R: 255, G: 222, B: 173}, // #FFDEAD
			"navy":                 {R: 000, G: 000, B: 128}, // #000080
			"oldlace":              {R: 253, G: 245, B: 230}, // #FDF5E6
			"olive":                {R: 128, G: 128, B: 000}, // #808000
			"olivedrab":            {R: 107, G: 142, B: 35},  // #6B8E23
			"orange":               {R: 255, G: 165, B: 000}, // #FFA500
			"orangered":            {R: 255, G: 69, B: 000},  // #FF4500
			"orchid":               {R: 218, G: 112, B: 214}, // #DA70D6
			"palegoldenrod":        {R: 238, G: 232, B: 170}, // #EEE8AA
			"palegreen":            {R: 152, G: 251, B: 152}, // #98FB98
			"paleturquoise":        {R: 175, G: 238, B: 238}, // #AFEEEE
			"palevioletred":        {R: 219, G: 112, B: 147}, // #DB7093
			"papayawhip":           {R: 255, G: 239, B: 213}, // #FFEFD5
			"peachpuff":            {R: 255, G: 218, B: 185}, // #FFDAB9
			"peru":                 {R: 205, G: 133, B: 63},  // #CD853F
			"pink":                 {R: 255, G: 192, B: 203}, // #FFC0CB
			"plum":                 {R: 221, G: 160, B: 221}, // #DDA0DD
			"powderblue":           {R: 176, G: 224, B: 230}, // #B0E0E6
			"purple":               {R: 160, G: 32, B: 240},  // #A020F0
			"rebeccapurple":        {R: 102, G: 51, B: 153},  // #663399
			"red":                  {R: 255, G: 000, B: 000}, // #FF0000
			"rosybrown":            {R: 188, G: 143, B: 143}, // #BC8F8F
			"royalblue":            {R: 65, G: 105, B: 225},  // #4169E1
			"saddlebrown":          {R: 139, G: 69, B: 19},   // #8B4513
			"salmon":               {R: 250, G: 128, B: 114}, // #FA8072
			"sandybrown":           {R: 244, G: 164, B: 96},  // #F4A460
			"seagreen":             {R: 46, G: 139, B: 87},   // #2E8B57
			"seashell":             {R: 255, G: 245, B: 238}, // #FFF5EE
			"sienna":               {R: 160, G: 82, B: 45},   // #A0522D
			"silver":               {R: 192, G: 192, B: 192}, // #C0C0C0
			"skyblue":              {R: 135, G: 206, B: 235}, // #87CEEB
			"slateblue":            {R: 106, G: 90, B: 205},  // #6A5ACD
			"slategray":            {R: 112, G: 128, B: 144}, // #708090
			"snow":                 {R: 255, G: 250, B: 250}, // #FFFAFA
			"springgreen":          {R: 000, G: 255, B: 127}, // #00FF7F
			"steelblue":            {R: 70, G: 130, B: 180},  // #4682B4
			"tan":                  {R: 210, G: 180, B: 140}, // #D2B48C
			"teal":                 {R: 000, G: 128, B: 128}, // #008080
			"thistle":              {R: 216, G: 191, B: 216}, // #D8BFD8
			"tomato":               {R: 255, G: 99, B: 71},   // #FF6347
			"turquoise":            {R: 64, G: 224, B: 208},  // #40E0D0
			"violet":               {R: 238, G: 130, B: 238}, // #EE82EE
			"webgray":              {R: 128, G: 128, B: 128}, // #808080
			"webgreen":             {R: 000, G: 128, B: 000}, // #008000
			"webmaroon":            {R: 127, G: 000, B: 000}, // #7F0000
			"webpurple":            {R: 127, G: 000, B: 127}, // #7F007F
			"wheat":                {R: 245, G: 222, B: 179}, // #F5DEB3
			"white":                {R: 255, G: 255, B: 255}, // #FFFFFF
			"whitesmoke":           {R: 245, G: 245, B: 245}, // #F5F5F5
			"yellow":               {R: 255, G: 255, B: 000}, // #FFFF00
			"yellowgreen":          {R: 154, G: 205, B: 50},  // #9ACD32
		}
		var doc PDF
		for key, val := range m {
			val.A = 255 // make opaque, not transparent
			color, err := doc.ToColor(key)
			tEqual(t, color, val)
			tEqual(t, err, nil)
		}
	}()
} //                                                         Test_PDF_ToColor_1_

// Test_PDF_ToColor_2_ is the second unit test for PDF.ToColor()
func Test_PDF_ToColor_2_(t *testing.T) {
	testCases := []struct {
		description string
		input       string
		color       color.RGBA
		errMsg      string
		errVal      string
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
			errMsg:      "Bad color code",
			errVal:      "#wrongcolor",
		},
		// X is not a valid hex char. Only valid values are: 0-9 and A-F
		{
			description: "hex with an invalid character",
			input:       "#845X76",
			color:       color.RGBA{A: 255},
			errMsg:      "Bad color code",
			errVal:      "#845X76",
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
			errMsg:      "Unknown color name",
			errVal:      "picasso",
		},
	}
	for _, test := range testCases {
		var doc PDF
		t.Run(test.description, func(t *testing.T) {
			color, err := doc.ToColor(test.input)
			inf := doc.ErrorInfo(err)
			if inf.Msg != test.errMsg {
				t.Fatalf("expected error message %q got %q",
					test.errMsg, inf.Msg)
			}
			if inf.Val != test.errVal {
				t.Fatalf("expected error message %v got %v",
					test.errVal, inf.Val)
			}
			if test.color != color {
				t.Fatalf("expected color %v got %v", test.color, color)
			}
		})
	}
} //                                                         Test_PDF_ToColor_2_

// Test_PDF_ToPoints_ is the unit test for PDF.ToPoints()
func Test_PDF_ToPoints_(t *testing.T) {
	//
	test := func(wantVal float64, wantErr error, inputParts ...[]string) {
		for _, s := range permuteStrings(inputParts...) {
			var doc PDF // uninitialized PDF
			gotVal, gotErr := doc.ToPoints(s)
			tEqual(t,
				fmt.Sprintf("%0.03f", gotVal),
				fmt.Sprintf("%0.03f", wantVal),
			)
			tEqual(t, gotErr, wantErr)
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
	one := []string{"1"}
	test(72, nil, spc, one, spc, inches, spc)   // 1 inch = 72 points
	test(2.835, nil, spc, one, spc, mm, spc)    // 1 mm = 2.835 points
	test(28.346, nil, spc, one, spc, cm, spc)   // 1 cm = 28.346 points
	test(0.050, nil, spc, one, spc, twips, spc) // 1 twip = 0.05 points
	test(1, nil, spc, one, spc, points, spc)    // 1 point = 1 point :)
	//
	// test negative number with decimals
	negative := []string{"-12.345"}
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
	test(0, fmt.Errorf(`Unknown measurement units "km"`), []string{"1km"})
	test(0, fmt.Errorf(`Invalid number "1.0.1"`), []string{"1.0.1mm"})
} //                                                          Test_PDF_ToPoints_

// Test_PDF_ToUnits_ is the unit test for
// ToUnits(points float64) float64
func Test_PDF_ToUnits_(t *testing.T) {
	func() {
		var doc PDF
		tEqual(t, doc.ToUnits(1), 1)
		//
		doc.SetUnits("cm")
		tEqual(t, doc.ToUnits(1), 0.035278)         // 1 point = 0.035278 cm
		tEqual(t, doc.ToUnits(28.3464566929134), 1) // ~28.3 points = 1 cm
		//
		doc.SetUnits("in")
		tEqual(t, doc.ToUnits(1), 0.0138888888888889) // 1 point = ~0.0138 in.
		tEqual(t, doc.ToUnits(72), 1)                 // 72 points = 1 inch
		//
		doc.SetUnits("mm")
		tEqual(t, doc.ToUnits(1), 0.3527777777777776) // 1 point = ~0.3527 mm
		tEqual(t, doc.ToUnits(2.83464566929134), 1)   // ~2.8346 points = 1 mm
		//
		doc.SetUnits("point")
		tEqual(t, doc.ToUnits(1), 1) // 1 point = 1 point
		//
		doc.SetUnits("twip")
		tEqual(t, doc.ToUnits(1), 20)   // 1 point = 20 twips
		tEqual(t, doc.ToUnits(0.05), 1) // 0.05 points = 1 twip
	}()
} //                                                           Test_PDF_ToUnits_

// Test_PDF_Units_ tests PDF.Units() and SetUnits()
func Test_PDF_Units_(t *testing.T) {
	//
	// (ob *PDF) Units() string
	//
	func() {
		var doc PDF // uninitialized PDF
		tEqual(t, doc.Units(), "POINT")
	}()
	func() {
		doc := NewPDF("A4")
		tEqual(t, doc.Units(), "POINT")
	}()
	//
	// (ob *PDF) SetUnits(units string) *PDF
	//
	func() {
		doc := NewPDF("A4")
		tEqual(t, len(doc.Errors()), 0)
		doc.SetUnits("cm")
		tEqual(t, len(doc.Errors()), 0)
		tEqual(t, doc.Units(), "CM")
	}()
	func() {
		doc := NewPDF("A4")
		tEqual(t, len(doc.Errors()), 0)
		doc.SetUnits("fathoms")
		tEqual(t, len(doc.Errors()), 1)
		//
		if len(doc.Errors()) == 1 {
			tEqual(t,
				doc.Errors()[0],
				fmt.Errorf(`Unknown measurement units "fathoms" @SetUnits`))
		}
		tEqual(t, doc.Units(), "POINT")
	}()
} //                                                             Test_PDF_Units_

// Test_PDF_X_ is the unit test for PDF.X()
func Test_PDF_X_(t *testing.T) {
	//
	// X of new PDF must be -1
	func() {
		var doc PDF
		tEqual(t, doc.X(), -1)
	}()
	func() {
		doc := NewPDF("A4")
		tEqual(t, doc.X(), -1)
	}()
	// SetX() sets the property?
	func() {
		var doc PDF
		doc.SetX(123)
		tEqual(t, doc.X(), 123)
	}()
	func() {
		doc := NewPDF("A4")
		doc.SetX(456)
		tEqual(t, doc.X(), 456)
	}()
	//
	// Test PDF generation
	//
	func() {
		doc := NewPDF("A4")
		doc.SetCompression(false).
			SetUnits("cm").
			SetXY(10, 1).
			SetFont("Times-Bold", 20).
			DrawText("X=10 Y=1")
		const want = `
		%PDF-1.4
		1 0 obj <</Type/Catalog/Pages 2 0 R>>
		endobj
		2 0 obj <</Type/Pages/Count 1/MediaBox[0 0 595 841]/Kids[3 0 R]>>
		endobj
		3 0 obj <</Type/Page/Parent 2 0 R/Contents 4 0 R
		/Resources <</Font <</FNT1 5 0 R>> >> >>
		endobj
		4 0 obj <</Length 92>> stream
		BT /FNT1 20 Tf ET
		0.000 0.000 0.000 rg
		0.000 0.000 0.000 RG
		BT 283 813 Td (X=10 Y=1) Tj ET
		endstream
		endobj
		5 0 obj <</Type/Font/Subtype/Type1/Name/FNT1
		/BaseFont/Times-Bold
		/Encoding/StandardEncoding>>
		endobj
		xref
		0 6
		0000000000 65535 f
		0000000010 00000 n
		0000000056 00000 n
		0000000130 00000 n
		0000000228 00000 n
		0000000370 00000 n
		trailer
		<</Size 6/Root 1 0 R>>
		startxref
		472
		%%EOF
		`
		pdfCompare(t, doc.Bytes(), want)
		// doc.SaveFile("~~Test_PDF_X_.pdf")
	}()
} //                                                                 Test_PDF_X_

// Test_PDF_Y_ is the unit test for PDF.Y()
func Test_PDF_Y_(t *testing.T) {
	//
	// Y of new PDF must be -1
	func() {
		var doc PDF
		tEqual(t, doc.Y(), -1)
	}()
	func() {
		doc := NewPDF("A4")
		tEqual(t, doc.Y(), -1)
	}()
	// SetY() sets the property?
	func() {
		var doc PDF
		doc.SetY(321)
		tEqual(t, doc.Y(), 321)
	}()
	func() {
		doc := NewPDF("A4")
		doc.SetY(654)
		tEqual(t, doc.Y(), 654)
	}()
	//
	// Test PDF generation
	//
	func() {
		doc := NewPDF("A4")
		doc.SetCompression(false).
			SetUnits("cm").
			SetXY(1, 10).
			SetFont("Times-Bold", 20).
			DrawText("X=1 Y=10")
		const want = `
		%PDF-1.4
		1 0 obj <</Type/Catalog/Pages 2 0 R>>
		endobj
		2 0 obj <</Type/Pages/Count 1/MediaBox[0 0 595 841]/Kids[3 0 R]>>
		endobj
		3 0 obj <</Type/Page/Parent 2 0 R/Contents 4 0 R
		/Resources <</Font <</FNT1 5 0 R>> >> >>
		endobj
		4 0 obj <</Length 91>> stream
		BT /FNT1 20 Tf ET
		0.000 0.000 0.000 rg
		0.000 0.000 0.000 RG
		BT 28 558 Td (X=1 Y=10) Tj ET
		endstream
		endobj
		5 0 obj <</Type/Font/Subtype/Type1/Name/FNT1
		/BaseFont/Times-Bold
		/Encoding/StandardEncoding>>
		endobj
		xref
		0 6
		0000000000 65535 f
		0000000010 00000 n
		0000000056 00000 n
		0000000130 00000 n
		0000000228 00000 n
		0000000369 00000 n
		trailer
		<</Size 6/Root 1 0 R>>
		startxref
		471
		%%EOF
		`
		pdfCompare(t, doc.Bytes(), want)
		// doc.SaveFile("~~Test_PDF_Y_.pdf")
	}()
} //                                                                 Test_PDF_Y_

// -----------------------------------------------------------------------------
// # Internal Tests

// go test --run Test_getPapreSize_
func Test_getPapreSize_(t *testing.T) {
	//
	// subtest: tests a specific paper size against width and height in points
	subtest := func(paperSize, permuted string, w, h float64, err error) {
		var doc PDF
		doc.SetUnits("mm")
		got, gotErr := doc.getPaperSize(permuted)
		if gotErr != err {
			t.Errorf("'error' mismatch: expected: %v returned %v", err, gotErr)
			t.Fail()
		}
		if got.name != paperSize {
			mismatch(t, permuted+" 'name'", paperSize, got.name)
		}
		if floatStr(got.widthPt) != floatStr(w) {
			mismatch(t, permuted+" 'widthPt'", w, got.widthPt)
		}
		if floatStr(got.heightPt) != floatStr(h) {
			mismatch(t, permuted+" 'heightPt'", h, got.heightPt)
		}
	}
	// test: tests the given paper size in portrait and landscape orientations
	// w and h are the paper width and height in mm
	// - permutes the paper size by including spaces
	// - converts units from mm to points
	test := func(paperSize string, w, h float64, err error) {
		const PTperMM = 2.83464566929134
		spaces := []string{"", " ", "  ", "   ", "\r", "\n", "\t"}
		for _, orient := range []string{"", "-l", "-L"} {
			permuted := permuteStrings(
				spaces,
				[]string{
					strings.ToLower(paperSize),
					strings.ToUpper(paperSize),
				},
				spaces,
				[]string{orient},
				spaces,
			)
			for _, s := range permuted {
				size := paperSize + strings.ToUpper(orient)
				if orient == "-l" || orient == "-L" {
					subtest(size, s, h*PTperMM, w*PTperMM, err)
				} else {
					subtest(size, s, w*PTperMM, h*PTperMM, err)
				}
			}
		}
	}
	test("A4", 210, 297, nil)
	test("A0", 841, 1189, nil)
	test("A1", 594, 841, nil)
	test("A2", 420, 594, nil)
	test("A3", 297, 420, nil)
	test("A4", 210, 297, nil)
	test("A5", 148, 210, nil)
	test("A6", 105, 148, nil)
	test("A7", 74, 105, nil)
	test("A8", 52, 74, nil)
	test("A9", 37, 52, nil)
	test("A10", 26, 37, nil)
	test("B0", 1000, 1414, nil)
	test("B1", 707, 1000, nil)
	test("B2", 500, 707, nil)
	test("B3", 353, 500, nil)
	test("B4", 250, 353, nil)
	test("B5", 176, 250, nil)
	test("B6", 125, 176, nil)
	test("B7", 88, 125, nil)
	test("B8", 62, 88, nil)
	test("B9", 44, 62, nil)
	test("B10", 31, 44, nil)
	test("C0", 917, 1297, nil)
	test("C1", 648, 917, nil)
	test("C2", 458, 648, nil)
	test("C3", 324, 458, nil)
	test("C4", 229, 324, nil)
	test("C5", 162, 229, nil)
	test("C6", 114, 162, nil)
	test("C7", 81, 114, nil)
	test("C8", 57, 81, nil)
	test("C9", 40, 57, nil)
	test("C10", 28, 40, nil)
	test("LEDGER", 432, 279, nil)
	test("LEGAL", 216, 356, nil)
	test("LETTER", 216, 279, nil)
	test("TABLOID", 279, 432, nil)
} //                                                          Test_getPapreSize_

// -----------------------------------------------------------------------------
// # Helper Functions

// callerList returns a human-friendly list of strings showing the
// call stack with each calling method or function's name and line number.
//
// The most immediate callers are listed first, followed by their callers,
// and so on. For brevity, 'runtime.*' and 'syscall.*'
// and other top-level callers are not included.
func callerList() []string {
	var ret []string
	i := 0
mainLoop:
	for {
		i++
		programCounter, filename, lineNo, _ := runtime.Caller(i)
		funcName := runtime.FuncForPC(programCounter).Name()
		//
		// end loop on reaching a top-level runtime function
		for _, s := range []string{
			"", "runtime.goexit", "runtime.main", "testing.tRunner",
		} {
			if funcName == s {
				break mainLoop
			}
		}
		if strings.Contains(funcName, "HandlerFunc.ServeHTTP") {
			break
		}
		// skip runtime/syscall functions, but continue the loop
		for _, s := range []string{
			".Callers", ".callerList", ".Error", ".Log", ".logAsync",
			"mismatch", "runtime.", "syscall.",
		} {
			if strings.Contains(funcName, s) {
				continue mainLoop
			}
		}
		switch showFileNames {
		case 1:
			filename = filepath.Base(filename)
		case 2:
			// let the file name's path use the right kind of OS path separator
			// (by default, the file name contains '/' on all platforms)
			if string(os.PathSeparator) != "/" {
				filename = strings.Replace(filename,
					"/", string(os.PathSeparator), -1)
			}
		}
		// remove parent module/function names
		if index := strings.LastIndex(funcName, "/"); index != -1 {
			funcName = funcName[index+1:]
		}
		if strings.Count(funcName, ".") > 1 {
			funcName = funcName[strings.Index(funcName, ".")+1:]
		}
		// remove unneeded punctuation from function names
		for _, find := range []string{"(", ")", "*"} {
			if strings.Contains(funcName, find) {
				funcName = strings.Replace(funcName, find, "", -1)
			}
		}
		line := fmt.Sprintf(":%d %s()", lineNo, funcName)
		if showFileNames > 0 {
			line = filename + line
		}
		ret = append(ret, line)
	}
	return ret
} //                                                                  callerList

// failIfHasErrors raises a test failure if the supplied PDF has errors
func failIfHasErrors(t *testing.T, errors func() []error) {
	if len(errors()) == 0 {
		return
	}
	for i, err := range errors() {
		t.Errorf("ERROR %d: %s\n\n", i+1, err)
	}
	t.Fail()
} //                                                             failIfHasErrors

// floatStr returns a float64 as a string, with val rounded to 3 decimals
func floatStr(val float64) string {
	return fmt.Sprintf("%0.3f", val)
} //                                                                    floatStr

// formatLines accepts an uncompressed PDF document as a string,
// and returns an array of trimmed, non-empty lines
func formatLines(s string, formatStreams bool) []string {
	//
	// format streams
	if formatStreams {
		s = pdfFormatStreams(s)
	}
	//
	// change all newlines to "\n"
	s = strings.Replace(s, "\r\n", "\n", -1)
	s = strings.Replace(s, "\r", "\n", -1)
	//
	// change all other white-spaces to spaces
	for _, space := range "\a\b\f\t\v" {
		s = strings.Replace(s, string(space), " ", -1)
	}
	// remove all repeated spaces
	for strings.Contains(s, "  ") {
		s = strings.Replace(s, "  ", " ", -1)
	}
	// trim and copy non-blank lines to result
	// also, continue lines that end with '\'
	var (
		ar   = strings.Split(s, "\n")
		ret  = make([]string, 0, len(ar))
		prev = ""
	)
	for _, line := range ar {
		line = strings.Trim(line, " \a\b\f\n\r\t\v")
		if line == "" {
			continue
		}
		// join lines split with '\'
		if prev != "" {
			line = prev + line
		}
		if strings.HasSuffix(line, "\\") {
			prev = strings.TrimRight(line, "\\")
			continue
		}
		// append line to result
		prev = ""
		ret = append(ret, line)
	}
	return ret
} //                                                                 formatLines

// getStack returns a list of line numbers and function names on the call stack
func getStack() string {
	buf := make([]byte, 8192)
	runtime.Stack(buf, true)
	var ar []string
	for _, s := range strings.Split(string(buf), "\n") {
		if strings.Contains(s, "\t") && !strings.Contains(s, "/testing.go") {
			ar = append(ar, "<- "+filepath.Base(s))
		}
	}
	return strings.Join(ar, "\n")
} //                                                                    getStack

// mismatch formats and raises a test error
func mismatch(t *testing.T, tag string, want, got interface{}) {
	ws := fmt.Sprintf("%v", want)
	gs := fmt.Sprintf("%v", got)
	t.Errorf("%s mismatch: expected: %s got: %s\n%s",
		tag, ws, gs, getStack())
} //                                                                    mismatch

// pdfCompare compares generated result bytes to the expected PDF content:
// - convert result ('got') to a string
// - format both result and expected string using formatLines()
// - compare result and expected lines ('got' and 'want')
// - raise an error if there are diffs (report up to 5 differences)
func pdfCompare(t *testing.T, got []byte, want string) {
	//
	const formatStreams = true
	var (
		gotAr    = formatLines(string(got), formatStreams)
		wantAr   = formatLines(want, !formatStreams)
		errCount = 0
		mismatch = false
		max      = len(gotAr)
	)
	if max < len(wantAr) {
		max = len(wantAr)
	}
	for i := 0; i < max; i++ {
		//
		// get the expected and the result line at i
		// if the slice is too short, leave it blank
		var want, got string
		if i < len(wantAr) {
			want = wantAr[i]
		}
		if i < len(gotAr) {
			got = gotAr[i]
		}
		if want == got { // no problem, move along
			continue
		}
		// only report the first 5 mismatches
		mismatch = true
		errCount++
		if errCount > 5 {
			break
		}
		t.Errorf("%s",
			"\n"+
				"/*\n"+
				"LOCATION: "+tCaller()+":\n"+
				"MISMATCH: L"+strconv.Itoa(i+1)+":\n"+
				"EXPECTED: "+want+"\n"+
				"PRODUCED: "+got+"\n"+
				"*/\n")
	}
	if mismatch {
		t.Errorf("%s",
			"\n"+
				"// RETURNED-PDF:\n"+
				"// "+tCaller()+"\n"+
				"`\n"+
				strings.Join(gotAr, "\n")+"\n"+
				"`\n")
	}
} //                                                                  pdfCompare

// pdfFormatStreams formats content of all streams in s as hex strings
func pdfFormatStreams(s string) string {
	const (
		STREAM    = ">> stream"
		ENDSTREAM = "endstream"
		BPL       = 16 // bytes per line
	)
	buf := bytes.NewBuffer(make([]byte, 0, len(s)))
	for part, s := range strings.Split(s, " obj ") {
		if part > 0 {
			buf.WriteString(" obj ")
		}
		// write the stream as-is if not compressed/image
		i := strings.Index(s, STREAM)
		if i == -1 ||
			(!strings.Contains(s[:i], "/FlateDecode") &&
				!strings.Contains(s[:i], "/Image")) {
			buf.WriteString(s)
			continue
		}
		// write the part before stream's data without changing it
		i += len(STREAM)
		buf.WriteString(s[:i])
		s = s[i:]
		//
		// write the stream's data as hex numbers (each line with BPL columns)
		buf.WriteString("\n")
		n := strings.Index(s, ENDSTREAM)
		if n == -1 {
			n = len(s)
		}
		c := 0
		for _, b := range []byte(s[:n]) {
			buf.WriteString(fmt.Sprintf(" %02X", b))
			c++
			if c >= BPL {
				buf.WriteString("\n")
				c = 0
			}
		}
		buf.WriteString("\n")
		//
		// write the part after the stream's data
		s = s[n:]
		if len(s) > 0 {
			buf.WriteString(s)
		}
	}
	return buf.String()
} //                                                            pdfFormatStreams

// permuteStrings returns all combinations of strings in 'parts'
func permuteStrings(parts ...[]string) (ret []string) {
	{
		n := 1
		for _, ar := range parts {
			n *= len(ar)
		}
		ret = make([]string, 0, n)
	}
	at := make([]int, len(parts))
	var buf bytes.Buffer
loop:
	for {
		// increment position counters
		for i := len(parts) - 1; i >= 0; i-- {
			if at[i] > 0 && at[i] >= len(parts[i]) {
				if i == 0 || (i == 1 && at[i-1] == len(parts[0])-1) {
					break loop
				}
				at[i] = 0
				at[i-1]++
			}
		}
		// construct permutated string
		buf.Reset()
		for i, ar := range parts {
			j := at[i]
			if j >= 0 && j < len(ar) {
				buf.WriteString(ar[j])
			}
		}
		ret = append(ret, buf.String())
		at[len(parts)-1]++
	}
	return ret
} //                                                              permuteStrings

// tCaller returns the name of the unit test function.
func tCaller() string {
	for _, caller := range callerList() {
		if strings.Contains(caller, "util.tCaller") ||
			strings.Contains(caller, "util.tEqual") ||
			strings.Contains(caller, "util.pdfCompare") {
			continue
		}
		return caller
	}
	return "<no-caller>"
} //                                                                     tCaller

const (
	showFileNames = 1
	// 0 - Don't show file names
	// 1 - Show only file name
	// 2 - Show file name and path
)

// tEqual asserts that 'got' is equal to 'want'.
// Provides a slightly-altered tEqual() function (and functions it uses)
// from Zircon-Go lib: github.com/balacode/zr
func tEqual(t *testing.T, got interface{}, want interface{}) bool {
	makeStr := func(value interface{}) string {
		switch v := value.(type) {
		case nil:
			{
				return "nil"
			}
		case bool:
			{
				if v {
					return "true"
				}
				return "false"
			}
		case int, int8, int16, int32, int64,
			uint, uint8, uint16, uint32, uint64, uintptr:
			{
				return fmt.Sprintf("%d", v)
			}
		case float64, float32:
			{
				s := fmt.Sprintf("%.4f", v)
				if strings.Contains(s, ".") {
					for strings.HasSuffix(s, "0") {
						s = s[:len(s)-1]
					}
					for strings.HasSuffix(s, ".") {
						s = s[:len(s)-1]
					}
				}
				return s
			}
		case error:
			{
				return v.Error()
			}
		case string:
			{
				return v
			}
		case time.Time: // use date part without time and time zone
			{
				s := v.Format(time.RFC3339)[:19] // "2006-01-02T15:04:05Z07:00"
				if strings.HasSuffix(s, "T00:00:00") {
					s = s[:10]
				}
				return s
			}
		case fmt.Stringer:
			return v.String()
		}
		return fmt.Sprintf("(type: %v value: %v)", reflect.TypeOf(value), value)
	}
	if makeStr(got) != makeStr(want) {
		t.Logf("\n"+"LOCATION: %s\n"+"EXPECTED: %s\n"+"RETURNED: %s\n",
			tCaller(), makeStr(want), makeStr(got))
		t.Fail()
		return false
	}
	return true
} //                                                                      tEqual

//end
