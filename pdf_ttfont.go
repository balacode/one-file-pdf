// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-04-24 00:53:39 9EAB5D                                [pdf_ttfont.go]
// -----------------------------------------------------------------------------

// THIS FILE IS A WORK IN PROGRESS

// This file contains a TTF font parser and PDF font-related functionality.

// # Module Initialization
//   init()
//
// # pdfFontHandler Interface (ob *pdfTTFont)
//   readFont(owner *PDF, font interface{}) bool
//   textWidthPt(s string) float64
//   writeText(s string)
//   writeFontObjects(font *pdfFont)

package pdf

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"reflect"
)

// pdfTTFont __
type pdfTTFont struct {
	Name string
	Err  error
	HEAD struct { //              font header table: general info about the font
		UnitsPerEm uint16
		XMin       int16
		YMin       int16
		XMax       int16
		YMax       int16
	}
	HHEA struct { //     horizontal header: layout of horizontally-written fonts
		HMetricCount uint16
		Ascent       int16
		Descent      int16
	}
	MAXP struct { //        maximum profile table: specifies memory requirements
		NumGlyphs uint16
	}
	HMTX struct { //               horizontal metrics table: width of each glyph
		Widths []uint16
	}
	CMAP struct { //                       maps character codes to glyph indices
		Chars map[int]uint16
	}
	NAME struct { //                                            font names table
		PostScriptName string
	}
	OS2 struct { //                                     OS/2 compatibility table
		Version        uint16
		STypoAscender  int16
		STypoDescender int16
		STypoLineGap   int16
	}
	POST struct { //                        glyph name and PostScript font table
		ItalicAngle int16
	}
	LOCA []uint32 //                                   glyph data location table
	pdf  *PDF
} //                                                                   pdfTTFont

// -----------------------------------------------------------------------------
// # Module Initialization

// init __
func init() {
	//TODO: uncomment this line when font handler is implemented
	// pdfNewFontHandler = func() pdfFontHandler { return &pdfTTFont{}	}
} //                                                                        init

// -----------------------------------------------------------------------------
// # pdfFontHandler Interface (ob *pdfTTFont)

// readFont loads a font from a file name, slice of bytes, or io.Reader
func (ob *pdfTTFont) readFont(owner *PDF, font interface{}) bool {
	ob.Err = nil
	ob.pdf = owner
	var src string
	var rd io.Reader
	switch arg := font.(type) {
	case string:
		src = arg
		var data, err = ioutil.ReadFile(arg)
		if err != nil {
			ob.pdf.putError(0xE5445B, "Failed reading font file", src)
			return false
		}
		rd = bytes.NewReader(data)
	case []byte:
		src = fmt.Sprintf("[]byte len(%d)", len(arg))
		rd = bytes.NewReader(arg)
	case io.Reader:
		src = "io.Reader"
		rd = arg
	default:
		ob.pdf.putError(0xECEB7B, "Invalid type in arg",
			reflect.TypeOf(font).String())
		return false
	}
	_ = rd //TODO: remove when reader is used
	return ob.Err == nil
} //                                                                    readFont

// textWidthPt returns the width of text 's' in points
func (ob *pdfTTFont) textWidthPt(s string) float64 {
	ob.Err = nil
	var ret float64
	for _, r := range s {
		_ = r
		var glyph, found = 0, false //TODO: find out glyph entry
		if !found {
			//TODO: 0xE0074A: error
			continue
		}
		var w = float64(ob.HMTX.Widths[glyph])
		if ob.HEAD.UnitsPerEm != 1000 {
			w = w * 1000.0 / float64(ob.HEAD.UnitsPerEm)
		}
		ret += w / 1000.0 * ob.pdf.fontSizePt
	}
	return ret
} //                                                                 textWidthPt

// writeText encodes text in the string 's'
func (ob *pdfTTFont) writeText(s string) {
	ob.Err = nil
	ob.pdf.write("BT ", ob.pdf.page.x, " ", ob.pdf.page.y, " Td ")
	//
	//TODO: add each rune of s, to determine glyphs to embed
	//
	// write hex encoded text to PDF
	ob.pdf.write("[<")
	for _, r := range s {
		_ = r
		var glyph, found = 0, false //TODO: find out glyph entry
		if !found {
			//TODO: 0xE1DC96: error
			return
		}
		ob.pdf.write(fmt.Sprintf("%04X", glyph))
	}
	ob.pdf.write(">] TJ ET\n")
} //                                                                   writeText

// writeFontObjects writes the PDF objects that define the embedded font
func (ob *pdfTTFont) writeFontObjects(font *pdfFont) {
	ob.Err = nil
	//TODO: write font-related objects here (call writer methods)
	if ob.Err != nil {
		ob.pdf.putError(0xED2CDF, ob.Err.Error(), "")
	}
} //                                                            writeFontObjects

//end
