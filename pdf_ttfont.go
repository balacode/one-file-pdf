// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-05-05 01:22:38 BD60FC                   one-file-pdf/[pdf_ttfont.go]
// -----------------------------------------------------------------------------

// THIS FILE IS A WORK IN PROGRESS

// This file contains a TTF font parser and PDF font-related functionality.
// It augments PDF in pdf_core.go to support Unicode and font embedding,
// but is not required for basic PDF functionality.

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
	Data []byte
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

// -----------------------------------------------------------------------------
// # TTF Parsing Methods (ob *pdfTTFont)

// readTTF __
func (ob *pdfTTFont) readTTF(reader io.Reader) {
	if ob.Err != nil {
		return
	}
	var err error
	ob.Data, err = ioutil.ReadAll(reader)
	if err != nil {
		ob.Err = err
		return
	}
	var rd = bytes.NewReader(ob.Data)
	var ver = ob.read(rd, 4)
	if !bytes.Equal(ver, []byte{0, 1, 0, 0}) {
		//TODO: 0xE0E9AE: error
		return
	}
	for _, fn := range []func(*bytes.Reader){
		ob.readHEAD, ob.readHHEA, ob.readMAXP, ob.readHMTX, ob.readCMAP,
		ob.readNAME, ob.readOS2, ob.readPOST, ob.readLOCA} {
		if ob.Err != nil {
			/// 0xE2257E: log error
			break
		}
		fn(rd)
	}
} //                                                                     readTTF

// readHEAD __
func (ob *pdfTTFont) readHEAD(rd *bytes.Reader) {
	if ob.Err != nil {
		return
	}
	//TODO: implement
} //                                                                    readHEAD

// readHHEA __
func (ob *pdfTTFont) readHHEA(rd *bytes.Reader) {
	if ob.Err != nil {
		return
	}
	//TODO: implement
} //                                                                    readHHEA

// readMAXP __
func (ob *pdfTTFont) readMAXP(rd *bytes.Reader) {
	if ob.Err != nil {
		return
	}
	//TODO: implement
} //                                                                    readMAXP

// readHMTX __
func (ob *pdfTTFont) readHMTX(rd *bytes.Reader) {
	if ob.Err != nil {
		return
	}
	//TODO: implement
} //                                                                    readHMTX

// readCMAP __
func (ob *pdfTTFont) readCMAP(rd *bytes.Reader) {
	if ob.Err != nil {
		return
	}
	//TODO: implement
} //                                                                    readCMAP

// readNAME __
func (ob *pdfTTFont) readNAME(rd *bytes.Reader) {
	if ob.Err != nil {
		return
	}
	//TODO: implement
} //                                                                    readNAME

// readOS2 __
func (ob *pdfTTFont) readOS2(rd *bytes.Reader) {
	if ob.Err != nil {
		return
	}
	//TODO: implement
} //                                                                     readOS2

// readPOST __
func (ob *pdfTTFont) readPOST(rd *bytes.Reader) {
	if ob.Err != nil {
		return
	}
	//TODO: implement
} //                                                                    readPOST

// readLOCA __
func (ob *pdfTTFont) readLOCA(rd *bytes.Reader) {
	if ob.Err != nil {
		return
	}
	//TODO: implement
} //                                                                    readLOCA

// read __
func (ob *pdfTTFont) read(rd *bytes.Reader, size int, useData ...bool) []byte {
	if ob.Err != nil {
		return nil
	}
	if len(useData) > 0 && useData[0] == false {
		_, err := rd.Seek(int64(size), 1)
		if err != nil {
			ob.Err = err
		}
		return nil
	}
	var ret = make([]byte, size)
	var n, err = rd.Read(ret)
	if err != nil {
		ob.Err = err
		return nil
	}
	if n != size {
		/// 0xE9B50D: error "END OF FILE DURING READING"
		return nil
	}
	return ret
} //                                                                        read

// readI16 __
func (ob *pdfTTFont) readI16(rd *bytes.Reader) int16 {
	var ar = ob.read(rd, 2)
	if ob.Err != nil {
		return 0
	}
	var ret = int(uint16(ar[0])<<8 | uint16(ar[1]))
	if ret >= 32768 {
		ret -= 65536
	}
	return int16(ret)
} //                                                                     readI16

// readUI16 __
func (ob *pdfTTFont) readUI16(rd *bytes.Reader) uint16 {
	var ar = ob.read(rd, 2)
	if ob.Err != nil {
		return 0
	}
	return uint16(ar[0])<<8 | uint16(ar[1])
} //                                                                    readUI16

// readUI32 __
func (ob *pdfTTFont) readUI32(rd *bytes.Reader) uint32 {
	var ar = ob.read(rd, 4)
	if ob.Err != nil {
		return 0
	}
	return uint32(ar[0])<<24 | uint32(ar[1])<<16 |
		uint32(ar[2])<<8 | uint32(ar[3])
} //                                                                    readUI32

//end
