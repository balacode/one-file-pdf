// -----------------------------------------------------------------------------
// github.com/balacode/one-file-pdf                 one-file-pdf/[pdf_ttfont.go]
// (c) balarabe@protonmail.com                                      License: MIT
// -----------------------------------------------------------------------------

// THIS FILE IS A WORK IN PROGRESS

// This file contains a TTF font parser and PDF font-related functionality.
// It augments PDF in pdf_core.go to support Unicode and font embedding,
// but is not required for basic PDF functionality.

// # Module Initialization
//   init()
//
// # pdfFontHandler Interface (f *pdfTTFont)
//   readFont(owner *PDF, font interface{}) bool
//   textWidthPt(s string) float64
//   writeText(s string)
//   writeFontObjects(font *pdfFont)

package pdf

import (
	"bytes"
	"fmt"
	"io"
	"os"
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
	// TODO: uncomment this line when font handler is implemented
	// pdfNewFontHandler = func() pdfFontHandler { return &pdfTTFont{} }
} //                                                                        init

// -----------------------------------------------------------------------------
// # pdfFontHandler Interface (f *pdfTTFont)

// readFont loads a font from a file name, slice of bytes, or io.Reader
func (f *pdfTTFont) readFont(owner *PDF, font interface{}) bool {
	f.Err = nil
	f.pdf = owner
	var (
		src string
		rd  io.Reader
	)
	switch arg := font.(type) {
	case string:
		{
			src = arg
			data, err := os.ReadFile(arg)
			if err != nil {
				f.pdf.putError(0xE5445B, "Failed reading font file", src)
				return false
			}
			rd = bytes.NewReader(data)
		}
	case []byte:
		{
			src = fmt.Sprintf("[]byte len(%d)", len(arg))
			rd = bytes.NewReader(arg)
		}
	case io.Reader:
		{
			src = "io.Reader"
			rd = arg
		}
	default:
		f.pdf.putError(0xECEB7B, "Invalid type in arg",
			reflect.TypeOf(font).String())
		return false
	}
	_ = rd // TODO: remove when reader is used
	return f.Err == nil
} //                                                                    readFont

// textWidthPt returns the width of text 's' in points
func (f *pdfTTFont) textWidthPt(s string) float64 {
	f.Err = nil
	var ret float64
	for _, r := range s {
		_ = r
		glyph, found := 0, false // TODO: find out glyph entry
		if !found {
			// TODO: 0xE0074A: error
			continue
		}
		w := float64(f.HMTX.Widths[glyph])
		if f.HEAD.UnitsPerEm != 1000 {
			w = w * 1000.0 / float64(f.HEAD.UnitsPerEm)
		}
		ret += w / 1000.0 * f.pdf.fontSizePt
	}
	return ret
} //                                                                 textWidthPt

// writeText encodes text in the string 's'
func (f *pdfTTFont) writeText(s string) {
	f.Err = nil
	f.pdf.write("BT ", f.pdf.page.x, " ", f.pdf.page.y, " Td ")
	//
	// TODO: add each rune of s, to determine glyphs to embed
	//
	// write hex encoded text to PDF
	f.pdf.write("[<")
	for _, r := range s {
		_ = r
		glyph, found := 0, false // TODO: find out glyph entry
		if !found {
			// TODO: 0xE1DC96: error
			return
		}
		f.pdf.write(fmt.Sprintf("%04X", glyph))
	}
	f.pdf.write(">] TJ ET\n")
} //                                                                   writeText

// writeFontObjects writes the PDF objects that define the embedded font
func (f *pdfTTFont) writeFontObjects(font *pdfFont) {
	f.Err = nil
	// TODO: write font-related objects here (call writer methods)
	if f.Err != nil {
		f.pdf.putError(0xED2CDF, f.Err.Error(), "")
	}
} //                                                            writeFontObjects

// -----------------------------------------------------------------------------
// # TTF Parsing Methods (f *pdfTTFont)

// readTTF __
func (f *pdfTTFont) readTTF(reader io.Reader) {
	if f.Err != nil {
		return
	}
	var err error
	f.Data, err = io.ReadAll(reader)
	if err != nil {
		f.Err = err
		return
	}
	rd := bytes.NewReader(f.Data)
	ver := f.read(rd, 4)
	if !bytes.Equal(ver, []byte{0, 1, 0, 0}) {
		// TODO: 0xE0E9AE: error
		return
	}
	for _, fn := range []func(*bytes.Reader){
		f.readHEAD, f.readHHEA, f.readMAXP, f.readHMTX, f.readCMAP,
		f.readNAME, f.readOS2, f.readPOST, f.readLOCA} {
		if f.Err != nil {
			/// 0xE2257E: log error
			break
		}
		fn(rd)
	}
} //                                                                     readTTF

// readHEAD __
func (f *pdfTTFont) readHEAD(rd *bytes.Reader) {
	if f.Err != nil {
		return
	}
	// TODO: implement
} //                                                                    readHEAD

// readHHEA __
func (f *pdfTTFont) readHHEA(rd *bytes.Reader) {
	if f.Err != nil {
		return
	}
	// TODO: implement
} //                                                                    readHHEA

// readMAXP __
func (f *pdfTTFont) readMAXP(rd *bytes.Reader) {
	if f.Err != nil {
		return
	}
	// TODO: implement
} //                                                                    readMAXP

// readHMTX __
func (f *pdfTTFont) readHMTX(rd *bytes.Reader) {
	if f.Err != nil {
		return
	}
	// TODO: implement
} //                                                                    readHMTX

// readCMAP __
func (f *pdfTTFont) readCMAP(rd *bytes.Reader) {
	if f.Err != nil {
		return
	}
	// TODO: implement
} //                                                                    readCMAP

// readNAME __
func (f *pdfTTFont) readNAME(rd *bytes.Reader) {
	if f.Err != nil {
		return
	}
	// TODO: implement
} //                                                                    readNAME

// readOS2 __
func (f *pdfTTFont) readOS2(rd *bytes.Reader) {
	if f.Err != nil {
		return
	}
	// TODO: implement
} //                                                                     readOS2

// readPOST __
func (f *pdfTTFont) readPOST(rd *bytes.Reader) {
	if f.Err != nil {
		return
	}
	// TODO: implement
} //                                                                    readPOST

// readLOCA __
func (f *pdfTTFont) readLOCA(rd *bytes.Reader) {
	if f.Err != nil {
		return
	}
	// TODO: implement
} //                                                                    readLOCA

// read __
func (f *pdfTTFont) read(rd *bytes.Reader, size int, useData ...bool) []byte {
	if f.Err != nil {
		return nil
	}
	if len(useData) > 0 && useData[0] == false {
		_, err := rd.Seek(int64(size), 1)
		if err != nil {
			f.Err = err
		}
		return nil
	}
	ret := make([]byte, size)
	n, err := rd.Read(ret)
	if err != nil {
		f.Err = err
		return nil
	}
	if n != size {
		/// 0xE9B50D: error "END OF FILE DURING READING"
		return nil
	}
	return ret
} //                                                                        read

// readI16 __
func (f *pdfTTFont) readI16(rd *bytes.Reader) int16 {
	ar := f.read(rd, 2)
	if f.Err != nil {
		return 0
	}
	ret := int(uint16(ar[0])<<8 | uint16(ar[1]))
	if ret >= 32768 {
		ret -= 65536
	}
	return int16(ret)
} //                                                                     readI16

// readUI16 __
func (f *pdfTTFont) readUI16(rd *bytes.Reader) uint16 {
	ar := f.read(rd, 2)
	if f.Err != nil {
		return 0
	}
	return uint16(ar[0])<<8 | uint16(ar[1])
} //                                                                    readUI16

// readUI32 __
func (f *pdfTTFont) readUI32(rd *bytes.Reader) uint32 {
	ar := f.read(rd, 4)
	if f.Err != nil {
		return 0
	}
	return uint32(ar[0])<<24 | uint32(ar[1])<<16 |
		uint32(ar[2])<<8 | uint32(ar[3])
} //                                                                    readUI32

// end
