// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-04-03 00:06:12 66EB99                              [one_file_pdf.go]
// -----------------------------------------------------------------------------

package pdf

// # Main Structure and Constructor
//   PDF struct
//   NewPDF(paperSize string) PDF
//
// # Read-Only Properties (ob *PDF)
//   CurrentPage() int
//   PageHeight() float64
//   PageWidth() float64
//
// # Property Getters (ob *PDF)
//   Color() color.RGBA             FontName() string
//   Compression() bool             FontSize() float64
//   DocAuthor() string             HorizontalScaling() uint16
//   DocCreator() string            LineWidth() float64
//   DocKeywords() string           Units() string
//   DocSubject() string            X() float64
//   DocTitle() string              Y() float64
//
// # Property Setters (ob *PDF)
//   SetColor(nameOrHTMLColor string) *PDF
//   SetColorRGB(red, green, blue int) *PDF
//   SetCompression(val bool) *PDF
//   SetDocAuthor(s string) *PDF
//   SetDocCreator(s string) *PDF
//   SetDocKeywords(s string) *PDF
//   SetDocSubject(s string) *PDF
//   SetDocTitle(s string) *PDF
//   SetFont(name string, points float64) *PDF
//   SetFontName(name string) *PDF
//   SetFontSize(points float64) *PDF
//   SetHorizontalScaling(percent uint16) *PDF
//   SetLineWidth(points float64) *PDF
//   SetUnits(unitName string) *PDF
//   SetX(x float64) *PDF
//   SetXY(x, y float64) *PDF
//   SetY(y float64) *PDF
//
// # Methods (ob *PDF)
//   AddPage() *PDF
//   Bytes() []byte
//   DrawBox(x, y, width, height float64, fill ...bool) *PDF
//   DrawCircle(x, y, radius float64, fill ...bool) *PDF
//   DrawEllipse(x, y, xRadius, yRadius float64, fill ...bool) *PDF
//   DrawImage(x, y, height float64, fileNameOrBytes interface{},
//       backColor ...string) *PDF
//   DrawLine(x1, y1, x2, y2 float64) *PDF
//   DrawText(s string) *PDF
//   DrawTextAlignedToBox(
//       x, y, width, height float64, align, text string) *PDF
//   DrawTextAt(x, y float64, text string) *PDF
//   DrawTextInBox(
//       x, y, width, height float64, align, text string ) *PDF
//   DrawUnitGrid() *PDF
//   FillBox(x, y, width, height float64) *PDF
//   FillCircle(x, y, radius float64) *PDF
//   FillEllipse(x, y, xRadius, yRadius float64) *PDF
//   NextLine() *PDF
//   Reset() *PDF
//   SaveFile(filename string) error
//   SetColumnWidths(widths ...float64) *PDF
//
// # Metrics Methods (ob *PDF)
//   TextWidth(s string) float64
//   ToColor(nameOrHTMLColor string) (color.RGBA, error)
//   ToPoints(numberAndUnit string) (float64, error)
//   ToUnits(points float64) float64
//   WrapTextLines(width float64, text string) (ret []string)
//
// # Error Handling Methods (ob *PDF)
//   Clean() *PDF
//   Errors() []error
//   PullError() error
//   (*PDF) ErrorInfo(err error) (ret struct {
//       ID            int
//       Msg, Src, Val string
//   })
//
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// # Internal Structures
//   pdfError struct
//       (err pdfError) Error() string
//   pdfFont struct
//   pdfImage struct
//   pdfPage struct
//   pdfPaperSize struct
//
// # Internal Methods (ob *PDF)
//   applyFont() (err error)
//   drawTextLine(s string) *PDF
//   drawTextBox(x, y, width, height float64,
//       wrapText bool, align, text string) *PDF
//   init() *PDF
//   loadImage(fileNameOrBytes interface{}, back color.RGBA,
//       ) (img pdfImage, idx int, err error)
//   makeImage(source image.Image, back color.RGBA,
//       ) (widthPx, heightPx int, isGray bool, data []byte)
//   reservePage() *PDF
//   textWidthPt(s string) float64
//
// # Internal Generation Methods (ob *PDF)
//   nextObj() int
//   write(format string, args ...interface{}) *PDF
//   writeCurve(x1, y1, x2, y2, x3, y3 float64) *PDF
//   writeEndobj() *PDF
//   writeMode(fill ...bool) (mode string)
//   writeObj(objType string) *PDF
//   writePages(pagesIndex, fontsIndex, imagesIndex int) *PDF
//   writeStream(content []byte) *PDF
//   writeStreamData(content []byte) *PDF
//
// # Internal Functions (*PDF) - just attached to PDF, but not using its data
//   escape(s string) []byte
//   isWhiteSpace(s string) bool
//   splitLines(s string) []string
//   toUpperLettersDigits(s, extras string) string
//   (ob *PDF):
//   getPaperSize(name string) (pdfPaperSize, error)
//   getPointsPerUnit(unitName string) (ret float64, err error)
//   putError(id int, msg, val string) *PDF
//
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// # Constants
//   PDFColorNames = map[string]color.RGBA
//
// # Internal Constants
//   pdfFontNames = []string
//   pdfFontWidths = [][]int
//   pdfStandardPaperSizes = []pdfPaperSize

import (
	"bytes"
	"compress/zlib"
	"crypto/sha512"
	"fmt"
	"image"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png" // init image decoders
	"io/ioutil"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"unicode" // only uses IsDigit(), IsLetter(), IsSpace()
)

// -----------------------------------------------------------------------------
// # Main Structure and Constructor

// PDF is the main structure representing a PDF document.
type PDF struct {
	paperSize    pdfPaperSize  // paper size used in this PDF
	pageNo       int           // current page number
	ppage        *pdfPage      // pointer to the current page
	pages        []pdfPage     // all the pages added to this PDF
	fonts        []pdfFont     // all the fonts used in this PDF
	images       []pdfImage    // all the images used in this PDF
	columnWidths []float64     // user-set column widths (like tab stops)
	columnNo     int           // index of the current column
	unitName     string        // name of active measurement unit
	ptPerUnit    float64       // number of points per measurement unit
	color        color.RGBA    // current drawing color
	lineWidth    float64       // current line width (in points)
	fontName     string        // current font's name
	fontSizePt   float64       // current font's size (in points)
	horzScaling  uint16        // horizontal scaling factor (in %)
	compression  bool          // enable stream compression?
	content      bytes.Buffer  // content buffer where PDF is written
	pbuf         *bytes.Buffer // pointer to PDF/current page's buffer
	objOffsets   []int         // used by Bytes() and write..()
	objNo        int           // used by Bytes() and write..()
	errors       []error       // errors that occurred during method calls
	isInit       bool          // has the PDF been initialized?
	//
	// document metadata fields
	docAuthor, docCreator, docKeywords, docSubject, docTitle string
} //                                                                         PDF

// NewPDF creates and initializes a new PDF object. Specify paperSize as:
// A, B, C series (e.g. "A4") or "LEGAL", "TABLOID", "LETTER", or "LEDGER".
// To specify a landscape orientation, add "-L" suffix e.g. "A4-L".
// You can also specify custom paper sizes using "width unit x height unit",
// for example "20 cm x 20 cm" or even "15cm x 10inch", etc.
func NewPDF(paperSize string) PDF {
	var ob PDF
	var size, err = ob.init().getPaperSize(paperSize)
	if err, isT := err.(pdfError); isT {
		ob.putError(0xE52F92, err.msg, paperSize)
		ob.paperSize, _ = ob.getPaperSize("A4")
	}
	ob.paperSize = size
	return ob
} //                                                                      NewPDF

// -----------------------------------------------------------------------------
// # Read-Only Properties (ob *PDF)

// CurrentPage returns the current page's number, 1 being the first page.
func (ob *PDF) CurrentPage() int { return ob.pageNo + 1 }

// PageHeight returns the height of the current page in selected units.
func (ob *PDF) PageHeight() float64 { return ob.ToUnits(ob.paperSize.heightPt) }

// PageWidth returns the width of the current page in selected units.
func (ob *PDF) PageWidth() float64 { return ob.ToUnits(ob.paperSize.widthPt) }

// -----------------------------------------------------------------------------
// # Property Getters (ob *PDF)

// Color returns the current color, which is used for text, lines and fills.
func (ob *PDF) Color() color.RGBA { ob.init(); return ob.color }

// Compression returns the current compression mode. If it is true,
// all PDF content will be compressed when the PDF is generated. If
// false, most PDF content (excluding images) will be in plain text,
// which is useful for debugging or to study PDF commands.
func (ob *PDF) Compression() bool { ob.init(); return ob.compression }

// DocAuthor returns the optional 'document author' metadata entry.
func (ob *PDF) DocAuthor() string { ob.init(); return ob.docAuthor }

// DocCreator returns the optional 'document creator' metadata entry.
func (ob *PDF) DocCreator() string { ob.init(); return ob.docCreator }

// DocKeywords returns the optional 'document keywords' metadata entry.
func (ob *PDF) DocKeywords() string { ob.init(); return ob.docKeywords }

// DocSubject returns the optional 'document subject' metadata entry.
func (ob *PDF) DocSubject() string { ob.init(); return ob.docSubject }

// DocTitle returns the optional 'document subject' metadata entry.
func (ob *PDF) DocTitle() string { ob.init(); return ob.docTitle }

// FontName returns the name of the currently-active typeface.
func (ob *PDF) FontName() string { ob.init(); return ob.fontName }

// FontSize returns the current font size in points.
func (ob *PDF) FontSize() float64 { ob.init(); return ob.fontSizePt }

// HorizontalScaling returns the current horizontal scaling in percent.
func (ob *PDF) HorizontalScaling() uint16 { ob.init(); return ob.horzScaling }

// LineWidth returns the current line width in points.
func (ob *PDF) LineWidth() float64 { ob.init(); return ob.lineWidth }

// Units returns the currently selected measurement units.
// E.g.: mm cm " in inch inches tw twip twips pt point points
func (ob *PDF) Units() string { ob.init(); return ob.unitName }

// X returns the X-coordinate of the current drawing position.
func (ob *PDF) X() float64 { return ob.reservePage().ToUnits(ob.ppage.x) }

// Y returns the Y-coordinate of the current drawing position.
func (ob *PDF) Y() float64 {
	return ob.reservePage().ToUnits(ob.paperSize.heightPt - ob.ppage.y)
} //                                                                           Y

// -----------------------------------------------------------------------------
// # Property Setters (ob *PDF)

// SetColor sets the current color using a web/X11 color name
// (e.g. "HoneyDew") or HTML color value such as "#191970"
// for midnight blue (#RRGGBB). The current color is used
// for subsequent text and line drawing and fills.
// If the name is unknown or invalid, sets color to black.
func (ob *PDF) SetColor(nameOrHTMLColor string) *PDF {
	var color, err = ob.init().ToColor(nameOrHTMLColor)
	if err, isT := err.(pdfError); isT {
		ob.putError(0xE5B3A5, err.msg, nameOrHTMLColor)
	}
	ob.color = color
	return ob
} //                                                                    SetColor

// SetColorRGB sets the current color using red, green and blue values.
// The current color is used for subsequent text/line drawing and fills.
func (ob *PDF) SetColorRGB(r, g, b byte) *PDF {
	ob.init()
	ob.color = color.RGBA{r, g, b, 255}
	return ob
} //                                                                 SetColorRGB

// SetCompression sets the compression mode used to generate the PDF.
// If set to true, all PDF steams will be compressed when the PDF is
// generated. If false, most content (excluding images) will be in
// plain text, which is useful for debugging or to study PDF commands.
func (ob *PDF) SetCompression(val bool) *PDF {
	ob.init()
	ob.compression = val
	return ob
} //                                                              SetCompression

// SetDocAuthor sets the optional 'document author' metadata entry.
func (ob *PDF) SetDocAuthor(s string) *PDF { ob.docAuthor = s; return ob }

// SetDocCreator sets the optional 'document creator' metadata entry.
func (ob *PDF) SetDocCreator(s string) *PDF { ob.docCreator = s; return ob }

// SetDocKeywords sets the optional 'document keywords' metadata entry.
func (ob *PDF) SetDocKeywords(s string) *PDF { ob.docKeywords = s; return ob }

// SetDocSubject sets the optional 'document subject' metadata entry.
func (ob *PDF) SetDocSubject(s string) *PDF { ob.docSubject = s; return ob }

// SetDocTitle sets the optional 'document title' metadata entry.
func (ob *PDF) SetDocTitle(s string) *PDF { ob.docTitle = s; return ob }

// SetFont changes the current font name and size in points.
// For the font name, use one of the standard font names, e.g. 'Helvetica'.
// This font will be used for subsequent text drawing.
func (ob *PDF) SetFont(name string, points float64) *PDF {
	return ob.SetFontName(name).SetFontSize(points)
} //                                                                     SetFont

// SetFontName changes the current font, while using the
// same font size as the previous font. Use one of the
// standard font names, such as 'Helvetica'.
func (ob *PDF) SetFontName(name string) *PDF {
	ob.init()
	ob.fontName = name
	return ob
} //                                                                 SetFontName

// SetFontSize changes the current font size in points,
// without changing the currently-selected font typeface.
func (ob *PDF) SetFontSize(points float64) *PDF {
	ob.init()
	ob.fontSizePt = points
	return ob
} //                                                                 SetFontSize

// SetHorizontalScaling changes the horizontal scaling in percent.
// For example, 200 will stretch text to double its normal width.
func (ob *PDF) SetHorizontalScaling(percent uint16) *PDF {
	ob.init()
	ob.horzScaling = percent
	return ob
} //                                                        SetHorizontalScaling

// SetLineWidth changes the line width in points.
func (ob *PDF) SetLineWidth(points float64) *PDF {
	ob.init()
	ob.lineWidth = points
	return ob
} //                                                                SetLineWidth

// SetUnits changes the current measurement units:
// mm cm " in inch inches tw twip twips pt point points (can be in any case)
func (ob *PDF) SetUnits(unitName string) *PDF {
	var ppu, err = ob.init().getPointsPerUnit(unitName)
	if err, isT := err.(pdfError); isT {
		return ob.putError(0xEB4AAA, err.msg, unitName)
	}
	ob.ptPerUnit, ob.unitName = ppu, ob.toUpperLettersDigits(unitName, "")
	return ob
} //                                                                    SetUnits

// SetX changes the X-coordinate of the current drawing position.
func (ob *PDF) SetX(x float64) *PDF {
	ob.init().reservePage()
	ob.ppage.x = x * ob.ptPerUnit
	return ob
} //                                                                        SetX

// SetXY changes both X- and Y-coordinates of the current drawing position.
func (ob *PDF) SetXY(x, y float64) *PDF { return ob.SetX(x).SetY(y) }

// SetY changes the Y-coordinate of the current drawing position.
func (ob *PDF) SetY(y float64) *PDF {
	ob.init().reservePage()
	ob.ppage.y = ob.paperSize.heightPt - y*ob.ptPerUnit
	return ob
} //                                                                        SetY

// -----------------------------------------------------------------------------
// # Methods (ob *PDF)

// AddPage appends a new blank page to the PDF and makes it the current page.
func (ob *PDF) AddPage() *PDF {
	var COLOR = color.RGBA{1, 0, 1, 0x01} // unlikely default color
	ob.pages = append(ob.pages, pdfPage{
		x: -1, y: -1, lineWidth: 1, strokeColor: COLOR, nonStrokeColor: COLOR,
		fontSizePt: 10, horzScaling: 100,
	})
	ob.pageNo = len(ob.pages) - 1
	ob.ppage = &ob.pages[ob.pageNo]
	ob.pbuf = &ob.ppage.content
	return ob
} //                                                                     AddPage

// Bytes generates the PDF document from various page and
// auxiliary objects and returns it in an array of bytes,
// identical to the content of a PDF file. This method is where
// you'll find the core structure of a PDF document.
func (ob *PDF) Bytes() []byte {
	// free any existing generated content and write PDF header
	ob.reservePage()
	const pagesIndex = 3
	var fontsIndex = pagesIndex + len(ob.pages)*2
	var imagesIndex = fontsIndex + len(ob.fonts)
	var infoIndex int // set when metadata found
	var prevBuf = ob.pbuf
	ob.content.Reset()
	ob.pbuf = &ob.content
	ob.objOffsets = []int{}
	ob.objNo = 0
	ob.write("%%PDF-1.4\n").
		writeObj("/Catalog").write("/Pages 2 0 R").writeEndobj()
	//
	//  write /Pages object (2 0 obj), page count, page size and the pages
	ob.writePages(pagesIndex, fontsIndex, imagesIndex)
	//
	// write fonts
	for _, iter := range ob.fonts {
		ob.writeObj("/Font")
		if iter.builtInIndex >= 0 {
			ob.write("/Subtype/Type1/Name/F%d/BaseFont/%s"+
				"/Encoding/WinAnsiEncoding", iter.fontID, iter.fontName)
		}
		ob.writeEndobj()
	}
	// write images
	for _, iter := range ob.images {
		var colorSpace = "DeviceRGB"
		if iter.isGray {
			colorSpace = "DeviceGray"
		}
		ob.writeObj("/XObject").
			write("/Subtype/Image/Width %d/Height %d/ColorSpace/%s"+
				"/BitsPerComponent 8", iter.widthPx, iter.heightPx, colorSpace).
			writeStreamData(iter.data).write("\nendobj\n")
	}
	// write info object
	if ob.docTitle != "" || ob.docSubject != "" ||
		ob.docKeywords != "" || ob.docAuthor != "" || ob.docCreator != "" {
		//
		infoIndex = imagesIndex + len(ob.images)
		ob.writeObj("/Info")
		for _, iter := range [][]string{
			{"/Title ", ob.docTitle}, {"/Subject ", ob.docSubject},
			{"/Keywords ", ob.docKeywords}, {"/Author ", ob.docAuthor},
			{"/Creator ", ob.docCreator},
		} {
			if iter[1] != "" {
				ob.write(iter[0]).
					write("(").write(string(ob.escape(iter[1]))).write(")")
			}
		}
		ob.writeEndobj()
	}
	// write cross-reference table at end of document
	var startXref = ob.content.Len()
	ob.write("xref\n0 %d\n0000000000 65535 f \n", len(ob.objOffsets))
	for _, offset := range ob.objOffsets[1:] {
		ob.write("%010d 00000 n \n", offset)
	}
	// write the trailer
	ob.write("trailer\n<</Size %d/Root 1 0 R", len(ob.objOffsets))
	if infoIndex > 0 {
		ob.write("/Info %d 0 R", infoIndex) // optional reference to info
	}
	ob.write(">>\nstartxref\n%d\n", startXref).write("%%%%EOF\n")
	ob.pbuf = prevBuf
	return ob.content.Bytes()
} //                                                                       Bytes

// DrawBox draws a rectangle.
func (ob *PDF) DrawBox(x, y, width, height float64, fill ...bool) *PDF {
	width, height = width*ob.ptPerUnit, height*ob.ptPerUnit
	x, y = x*ob.ptPerUnit, ob.paperSize.heightPt-y*ob.ptPerUnit-height
	var mode = ob.writeMode(fill...)
	return ob.write("%.3f %.3f %.3f %.3f re %s\n", x, y, width, height, mode)
	// re: construct a rectangular path
} //                                                                     DrawBox

// DrawCircle draws a circle of radius r centered on (x, y),
// by drawing 4 Bézier curves (PDF has no circle primitive)
func (ob *PDF) DrawCircle(x, y, radius float64, fill ...bool) *PDF {
	return ob.DrawEllipse(x, y, radius, radius, fill...)
} //                                                                  DrawCircle

// DrawEllipse draws an ellipse centered on (x, y),
// with horizontal radius xRadius and vertical radius yRadius
// by drawing 4 Bézier curves (PDF has no ellipse primitive)
func (ob *PDF) DrawEllipse(x, y, xRadius, yRadius float64, fill ...bool) *PDF {
	x, y = x*ob.ptPerUnit, ob.paperSize.heightPt-y*ob.ptPerUnit
	const ratio = 0.552284749830794  // (4/3) * tan(PI/8)
	var r = xRadius * ob.ptPerUnit   // horizontal radius
	var v = yRadius * ob.ptPerUnit   // vertical radius
	var m, n = r * ratio, v * ratio  // ratios for control points
	var mode = ob.writeMode(fill...) // prepare colors/line width
	//
	return ob.write(" %.3f %.3f m", x-r, y). // x0 y0 m: move to point (x0, y0)
		//         control-1 control-2 endpoint
		writeCurve(x-r, y+n, x-m, y+v, x+0, y+v). // top left arc
		writeCurve(x+m, y+v, x+r, y+n, x+r, y+0). // top right
		writeCurve(x+r, y-n, x+m, y-v, x+0, y-v). // bottom right
		writeCurve(x-m, y-v, x-r, y-n, x-r, y+0). // bottom left
		write(" %s\n", mode)                      // b: fill or S: stroke
} //                                                                 DrawEllipse

// DrawImage draws a PNG image. x, y, height specify the position and height
// of the image. Width is scaled to match the image's aspect ratio.
// fileNameOrBytes is either a string specifying a file name,
// or a byte slice with PNG image data.
func (ob *PDF) DrawImage(x, y, height float64, fileNameOrBytes interface{},
	backColor ...string) *PDF {
	//
	var back = color.RGBA{R: 255, G: 255, B: 255, A: 255} // white by default
	if len(backColor) > 0 {
		back, _ = ob.ToColor(backColor[0])
	}
	// add the image to the current page, if not already referenced
	ob.reservePage()
	var img, idx, err = ob.loadImage(fileNameOrBytes, back)
	if err, isT := err.(pdfError); isT {
		return ob.putError(0xE8F375, err.msg, err.val)
	}
	var found bool
	for _, iter := range ob.ppage.imageNos {
		if iter == idx {
			found = true
			break
		}
	}
	if !found {
		ob.ppage.imageNos = append(ob.ppage.imageNos, idx)
	}
	// draw the image
	var h = height * ob.ptPerUnit
	var w = float64(img.widthPx) / float64(img.heightPx) * h
	x, y = x*ob.ptPerUnit, ob.paperSize.heightPt-y*ob.ptPerUnit-h
	return ob.write("q\n %f 0 0 %f %f %f cm\n/IMG%d Do\nQ\n", w, h, x, y, idx)
	//                   w      h  x  y
	//                q: save graphics state
	//               cm: concatenate matrix to current transform matrix
	//               Do: invoke named XObject (/IMGn)
	//                Q: restore graphics state
} //                                                                   DrawImage

// DrawLine draws a straight line from point (x1, y1) to point (x2, y2).
func (ob *PDF) DrawLine(x1, y1, x2, y2 float64) *PDF {
	x1, y1 = x1*ob.ptPerUnit, ob.paperSize.heightPt-y1*ob.ptPerUnit
	x2, y2 = x2*ob.ptPerUnit, ob.paperSize.heightPt-y2*ob.ptPerUnit
	ob.writeMode(true) // prepare color/line width
	return ob.write("%.3f %.3f m %.3f %.3f l S\n", x1, y1, x2, y2)
	// m: move  S: stroke path (for lines)
} //                                                                    DrawLine

// DrawText draws a text string at the current position (X, Y).
func (ob *PDF) DrawText(s string) *PDF {
	if len(ob.columnWidths) == 0 {
		return ob.drawTextLine(s)
	}
	var x = 0.0
	for i := 0; i < ob.columnNo; i, x = i+1, x+ob.columnWidths[i] {
	}
	ob.SetX(x).drawTextLine(s)
	if ob.columnNo < len(ob.columnWidths)-1 {
		ob.columnNo++
		return ob
	}
	return ob.NextLine()
} //                                                                    DrawText

// DrawTextAlignedToBox draws 'text' within a rectangle specified
// by 'x', 'y', 'width' and 'height'. If 'align' is blank, the
// text is center-aligned both vertically and horizontally.
// Specify 'L' or 'R' to align the text left or right, and 'T' or
// 'B' to align the text to the top or bottom of the box.
func (ob *PDF) DrawTextAlignedToBox(
	x, y, width, height float64, align, text string) *PDF {
	return ob.drawTextBox(x, y, width, height, false, align, text)
} //                                                        DrawTextAlignedToBox

// DrawTextAt draws text at the specified point (x, y).
func (ob *PDF) DrawTextAt(x, y float64, text string) *PDF {
	return ob.SetXY(x, y).DrawText(text)
} //                                                                  DrawTextAt

// DrawTextInBox draws word-wrapped text within a rectangle
// specified by 'x', 'y', 'width' and 'height'. If 'align' is blank,
// the text is center-aligned both vertically and horizontally.
// Specify 'L' or 'R' to align the text left or right, and 'T' or
// 'B' to align the text to the top or bottom of the box.
func (ob *PDF) DrawTextInBox(
	x, y, width, height float64, align, text string) *PDF {
	return ob.drawTextBox(x, y, width, height, true, align, text)
} //                                                               DrawTextInBox

// DrawUnitGrid draws a light-gray grid demarcated in the
// current measurement unit. The grid fills the entire page.
// It helps with item positioning.
func (ob *PDF) DrawUnitGrid() *PDF {
	var pw, ph = ob.PageWidth(), ob.PageHeight()
	ob.SetLineWidth(0.1).SetFont("Helvetica", 8)
	for i, x := 0, 0.0; x < pw; i, x = i+1, x+1 { //            vertical lines |
		ob.SetColorRGB(200, 200, 200).DrawLine(x, 0, x, ph).
			SetColor("Indigo").SetXY(x+0.1, 0.3).DrawText(strconv.Itoa(i))
	}
	for i, y := 0, 0.0; y < ph; i, y = i+1, y+1 { //          horizontal lines -
		ob.SetColorRGB(200, 200, 200).DrawLine(0, y, pw, y).
			SetColor("Indigo").SetXY(0.1, y+0.3).DrawText(strconv.Itoa(i))
	}
	return ob
} //                                                                DrawUnitGrid

// FillBox fills a rectangle with the current color.
func (ob *PDF) FillBox(x, y, width, height float64) *PDF {
	return ob.DrawBox(x, y, width, height, true)
} //                                                                     FillBox

// FillCircle fills a circle of radius r centered on (x, y),
// by drawing 4 Bézier curves (PDF has no circle primitive)
func (ob *PDF) FillCircle(x, y, radius float64) *PDF {
	return ob.DrawEllipse(x, y, radius, radius, true)
} //                                                                  FillCircle

// FillEllipse fills a Ellipse of radius r centered on (x, y),
// by drawing 4 Bézier curves (PDF has no ellipse primitive)
func (ob *PDF) FillEllipse(x, y, xRadius, yRadius float64) *PDF {
	return ob.DrawEllipse(x, y, xRadius, yRadius, true)
} //                                                                 FillEllipse

// NextLine advances the text writing position to the next line.
// I.e. the Y increases by the height of the font and
// the X-coordinate is reset to zero.
func (ob *PDF) NextLine() *PDF {
	var x, y = 0.0, ob.Y() + ob.FontSize()*ob.ptPerUnit
	if len(ob.columnWidths) > 0 {
		x = ob.columnWidths[0]
	}
	if y > ob.paperSize.heightPt*ob.ptPerUnit {
		ob.AddPage()
		y = 0
	}
	ob.columnNo = 0
	return ob.SetXY(x, y)
} //                                                                    NextLine

// Reset releases all resources and resets all variables, except paper size.
func (ob *PDF) Reset() *PDF {
	ob.ppage, ob.pbuf = nil, nil
	*ob = NewPDF(ob.paperSize.name)
	return ob
} //                                                                       Reset

// SaveFile generates and saves the PDF document to a file.
func (ob *PDF) SaveFile(filename string) error {
	var err = ioutil.WriteFile(filename, ob.Bytes(), 0644)
	if err != nil {
		ob.putError(0xED3F6D, "Failed writing file", err.Error())
	}
	return err
} //                                                                    SaveFile

// SetColumnWidths creates column positions (tab stops) along the X-axis.
// To remove all column positions, call this method without any argument.
func (ob *PDF) SetColumnWidths(widths ...float64) *PDF {
	ob.init()
	ob.columnWidths = widths
	return ob
} //                                                             SetColumnWidths

// -----------------------------------------------------------------------------
// # Metrics Methods (ob *PDF)

// TextWidth returns the width of the text in current units.
func (ob *PDF) TextWidth(s string) float64 {
	return ob.ToUnits(ob.textWidthPt(s))
} //                                                                   TextWidth

// ToColor returns an RGBA color value from a web/X11 color name
// (e.g. "HoneyDew") or HTML color value such as "#191970"
// If the name or code is unknown or invalid, returns zero value (black).
func (ob *PDF) ToColor(nameOrHTMLColor string) (color.RGBA, error) {
	//
	// if name starts with '#' treat it as HTML color code (#RRGGBB)
	var s = ob.toUpperLettersDigits(nameOrHTMLColor, "#")
	if len(s) >= 7 && s[0] == '#' {
		var hex [6]byte
		for i, r := range s[1:7] {
			switch {
			case r >= '0' && r <= '9':
				hex[i] = byte(r - '0')
			case r >= 'A' && r <= 'F':
				hex[i] = byte(r - 'A' + 10)
			default:
				return pdfBlack, pdfError{id: 0xEED50B, src: "ToColor",
					msg: "Bad color code", val: nameOrHTMLColor}
			}
		}
		return color.RGBA{
			hex[0]*16 + hex[1],
			hex[2]*16 + hex[3],
			hex[4]*16 + hex[5], 255}, nil
	}
	// otherwise search for color name
	var c, found = PDFColorNames[s]
	if found {
		return color.RGBA{c.R, c.G, c.B, 255}, nil
	}
	return pdfBlack, pdfError{id: 0xE00982, src: "ToColor",
		msg: "Unknown color name", val: nameOrHTMLColor}
} //                                                                     ToColor

// ToPoints converts a string composed of a number and unit to points.
// For example '1 cm' or '1cm' becomes 28.346 points.
// Recognised units: mm cm " in inch inches tw twip twips pt point points
func (ob *PDF) ToPoints(numberAndUnit string) (float64, error) {
	var num, unit string //                              extract number and unit
	for _, r := range numberAndUnit {
		switch {
		case r == '-', r == '.', unicode.IsDigit(r):
			num += string(r)
		case r == '"', unicode.IsLetter(r):
			unit += string(r)
		}
	}
	var ppu = 1.0
	if unit != "" {
		var err error
		ppu, err = ob.getPointsPerUnit(unit)
		if err, isT := err.(pdfError); isT {
			return 0, err
		}
	}
	var n, err = strconv.ParseFloat(num, 64)
	if err != nil {
		return 0, pdfError{id: 0xE154AC, msg: "Invalid number", val: num}
	}
	return n * ppu, nil
} //                                                                    ToPoints

// ToUnits converts points to the currently selected unit of measurement.
func (ob *PDF) ToUnits(points float64) float64 {
	if int(ob.ptPerUnit*100) == 0 {
		return points
	}
	return points / ob.ptPerUnit
} //                                                                     ToUnits

// WrapTextLines splits a string into multiple lines so that the text
// fits in the specified width. The text is wrapped on word boundaries.
// Newline characters (CR and "\n") also cause text to be split.
// You can find out the number of lines needed to wrap some
// text by checking the length of the returned array.
func (ob *PDF) WrapTextLines(width float64, text string) (ret []string) {
	var fit = func(s string, step, n int, width float64) int {
		for max := len(s); n > 0 && n <= max; {
			var w = ob.TextWidth(s[:n])
			switch step {
			case 1, 3: //       keep halving (or - 1) until n chars fit in width
				if w <= width {
					return n
				}
				n--
				if step == 1 {
					n /= 2
				}
			case 2: //               increase n until n chars won't fit in width
				if w > width {
					return n
				}
				n = int((float64(n) * 1.2)) //                 increase n by 20%
			}
		}
		return 0
	}
	// split text into lines. then break lines based on text width
	for _, iter := range ob.splitLines(text) {
		for ob.TextWidth(iter) > width {
			var n = len(iter) // reduce, increase, then reduce n to get best fit
			for i := 1; i <= 3; i++ {
				n = fit(iter, i, n, width)
			}
			// move to the last word (if white-space is found)
			var found, max = false, n
			for n > 0 {
				if ob.isWhiteSpace(iter[n-1 : n]) {
					found = true
					break
				}
				n--
			}
			if !found {
				n = max
			}
			if n <= 0 {
				break
			}
			ret = append(ret, iter[:n])
			iter = iter[n:]
		}
		ret = append(ret, iter)
	}
	return ret
} //                                                               WrapTextLines

// -----------------------------------------------------------------------------
// # Error Handling Methods (ob *PDF)

// Clean clears all accumulated errors.
func (ob *PDF) Clean() *PDF { ob.errors = nil; return ob }

// ErrorInfo extracts and returns additional error details from PDF errors
func (*PDF) ErrorInfo(err error) (ret struct {
	ID            int
	Msg, Src, Val string
}) {
	if err, isT := err.(pdfError); isT {
		ret.ID, ret.Msg, ret.Src, ret.Val = err.id, err.msg, err.src, err.val
	}
	return ret
} //                                                                   ErrorInfo

// Errors returns a slice of all accumulated errors.
func (ob *PDF) Errors() []error { return ob.errors }

// PullError removes and returns the first error from the errors collection.
func (ob *PDF) PullError() error {
	if len(ob.errors) == 0 {
		return nil
	}
	var ret = ob.errors[0]
	ob.errors = ob.errors[1:]
	return ret
} //                                                                   PullError

// -----------------------------------------------------------------------------
// # Internal Structures

// pdfError stores extended error details for errors in this package.
type pdfError struct {
	id            int    // unique ID of the error (only within package)
	msg, src, val string // the error message, source method and invalid value
} //                                                                    pdfError

// Error creates and returns an error message from pdfError details
func (err pdfError) Error() string {
	var ret = fmt.Sprintf("%s %q", err.msg, err.val)
	if err.src != "" {
		ret += " @" + err.src
	}
	return ret
} //                                                                       Error

// pdfFont represents a font name and its appearance
type pdfFont struct {
	fontID           int
	fontName         string
	builtInIndex     int
	isBold, isItalic bool
} //                                                                     pdfFont

// pdfImage represents an image
type pdfImage struct {
	filename          string     // name of file from which image was read
	widthPx, heightPx int        // width and height in pixels
	data              []byte     // image data
	hash              [64]byte   // hash of data (used to compare images)
	backColor         color.RGBA // background color (used to compare images)
	isGray            bool       // image is grayscale? (if false, color image)
} //                                                                    pdfImage

// pdfPage holds references, state and the stream buffer for each page
type pdfPage struct {
	fontIDs, imageNos           []int        // references to fonts and images
	x, y, lineWidth, fontSizePt float64      // current drawing state
	strokeColor, nonStrokeColor color.RGBA   // "
	fontID                      int          // "
	horzScaling                 uint16       // "
	content                     bytes.Buffer // write..() calls send output here
} //                                                                     pdfPage

// pdfPaperSize represents a page size name and its dimensions in points
type pdfPaperSize struct {
	name              string  // paper size: e.g. 'Letter', 'A4', etc.
	widthPt, heightPt float64 // width and height in points
} //                                                                pdfPaperSize

// -----------------------------------------------------------------------------
// # Internal Methods (ob *PDF)

// applyFont writes a font change command, provided the font has
// been changed since the last operation that uses fonts.
//
// This should be called just before a font needs to be used.
// This way, if a font is picked with SetFontName() property, but
// never used to draw text, no font selection command is output.
//
// Before calling this method, the font name must be already
// set by SetFontName(), which is stored in ob.font.fontName
//
// What this method does:
// - Validates the current font name and determines if it is a
//   standard (built-in) font like Helvetica or a TrueType font.
// - Fills the document-wide list of fonts (ob.fonts).
// - Adds items to the list of font ID's used on the current page.
func (ob *PDF) applyFont() (err error) {
	var font pdfFont
	var name = ob.toUpperLettersDigits(ob.fontName, "")
	var valid = name != ""
	if valid {
		valid = false
		for i, iter := range pdfFontNames {
			iter = ob.toUpperLettersDigits(iter, "")
			if iter != name {
				continue
			}
			var has = strings.Contains
			font = pdfFont{
				fontName:     pdfFontNames[i],
				builtInIndex: i,
				isBold:       has(iter, "BOLD"),
				isItalic:     has(iter, "OBLIQUE") || has(iter, "ITALIC"),
			}
			valid = true
			break
		}
	}
	// if there is no selected font or it's invalid, use Helvetica
	if !valid {
		err = pdfError{id: 0xE86819, msg: "Invalid font", val: ob.fontName}
		ob.fontName = "Helvetica"
		ob.applyFont()
		return err
	}
	// has the font been added to the global list? if not, add it:
	for _, iter := range ob.fonts {
		if font.fontName == iter.fontName {
			font.fontID = iter.fontID
			break
		}
	}
	if font.fontID == 0 {
		font.fontID = 1 + len(ob.fonts)
		ob.fonts = append(ob.fonts, font)
	}
	if ob.ppage.fontID == font.fontID &&
		int(ob.ppage.fontSizePt*100) == int(ob.fontSizePt)*100 {
		return err
	}
	// add the font ID to the current page, if not already referenced
	var alreadyUsedOnPage bool
	for _, id := range ob.ppage.fontIDs {
		if id == font.fontID {
			alreadyUsedOnPage = true
			break
		}
	}
	if !alreadyUsedOnPage {
		ob.ppage.fontIDs = append(ob.ppage.fontIDs, 0)
		ob.ppage.fontIDs[len(ob.ppage.fontIDs)-1] = font.fontID
	}
	ob.ppage.fontID = font.fontID
	ob.ppage.fontSizePt = ob.fontSizePt
	ob.write("BT /FNT%d %d Tf ET\n", ob.ppage.fontID, int(ob.ppage.fontSizePt))
	// BT: begin text   /FNT0 i0 Tf: set font to FNT0 index i0   ET: end text
	return err
} //                                                                   applyFont

// drawTextLine writes a line of text at the current coordinates to the
// current page's content stream, using a sequence of raw PDF commands
func (ob *PDF) drawTextLine(s string) *PDF {
	if s == "" {
		return ob
	}
	// draw the text
	if err, isT := ob.applyFont().(pdfError); isT {
		ob.putError(0xEAEAC4, err.msg, err.val)
	}
	if ob.ppage.horzScaling != ob.horzScaling {
		ob.ppage.horzScaling = ob.horzScaling
		ob.write("BT %d Tz ET\n", ob.ppage.horzScaling)
		// BT: begin text   n0 Tz: set horiz. text scaling to n0%   ET: end text
	}
	ob.writeMode(true) // fill/nonStroke
	ob.write("BT %d %d Td (%s) Tj ET\n",
		int(ob.ppage.x), int(ob.ppage.y), ob.escape(s))
	// BT: begin text   Td: move text position   Tj: show text   ET: end text
	ob.ppage.x += ob.textWidthPt(s)
	return ob
} //                                                                drawTextLine

// drawTextBox draws a line of text, or a word-wrapped block of text.
// align: specify up to 2 flags: L R T B to align left, right, top or bottom
// the default (blank) is C center, both vertically and horizontally
func (ob *PDF) drawTextBox(x, y, width, height float64,
	wrapText bool, align, text string,
) *PDF {
	if text == "" {
		return ob
	}
	ob.reservePage()
	if err, isT := ob.applyFont().(pdfError); isT {
		ob.putError(0xE0737C, err.msg, err.val)
	}
	var lines []string
	if wrapText {
		lines = ob.WrapTextLines(width, text)
	} else {
		lines = []string{text}
	}
	align = strings.ToUpper(align)
	var lineHeight = ob.FontSize()
	var allLinesHeight = lineHeight * float64(len(lines))
	//
	// calculate aligned y-axis position of text (top, bottom, center)
	y, height = y*ob.ptPerUnit+ob.fontSizePt, height*ob.ptPerUnit
	if strings.Contains(align, "B") { // bottom
		y += height - allLinesHeight - 4 //                           4pt margin
	} else if !strings.Contains(align, "T") {
		y += height/2 - allLinesHeight/2 - ob.fontSizePt*0.15 //         center
	}
	y = ob.paperSize.heightPt - y
	//
	// calculate x-axis position of text (left, right, center)
	x, width = x*ob.ptPerUnit, width*ob.ptPerUnit
	for _, line := range lines {
		var off = 0.0 //                                x-offset to align in box
		if strings.Contains(align, "L") {
			off = ob.fontSizePt / 6 //                              left margin
		} else if strings.Contains(align, "R") {
			off = width - ob.textWidthPt(line) - ob.fontSizePt/6
		} else {
			off = width/2 - ob.textWidthPt(line)/2 //                     center
		}
		ob.ppage.x, ob.ppage.y = x+off, y
		ob.drawTextLine(line)
		y -= lineHeight
	}
	return ob
} //                                                                 drawTextBox

// init initializes the PDF object, if not initialized already
func (ob *PDF) init() *PDF {
	if ob.isInit {
		return ob
	}
	ob.unitName = "POINT"
	ob.paperSize, _ = ob.getPaperSize("A4")
	ob.ptPerUnit, _ = ob.getPointsPerUnit(ob.unitName)
	ob.color, ob.lineWidth = pdfBlack, 1 // point
	ob.fontName, ob.fontSizePt = "Helvetica", 10
	ob.horzScaling, ob.compression = 100, true
	ob.isInit = true
	return ob
} //                                                                        init

// loadImage reads an image from a file or byte array, stores its data in
// the PDF's images array, and returns a pdfImage and its reference index
func (ob *PDF) loadImage(fileNameOrBytes interface{}, back color.RGBA,
) (img pdfImage, idx int, err error) {
	var buf *bytes.Buffer
	switch val := fileNameOrBytes.(type) {
	case string:
		for i, iter := range ob.images {
			if iter.filename == val && iter.backColor == back {
				return iter, i, nil
			}
		}
		img.filename = val
		var data, err = ioutil.ReadFile(val)
		if err != nil {
			return pdfImage{}, -1, pdfError{id: 0xE9F387,
				msg: "Failed reading file", val: err.Error()}
		}
		buf = bytes.NewBuffer(data)
		img.hash = sha512.Sum512(data)
	case []byte:
		buf = bytes.NewBuffer(val)
		img.hash = sha512.Sum512(val)
	default:
		return pdfImage{}, -1,
			pdfError{id: 0xEE3E42, msg: "Invalid type in fileNameOrBytes",
				val: fmt.Sprintf("%s = %v",
					reflect.TypeOf(fileNameOrBytes), fileNameOrBytes)}
	}
	for i, iter := range ob.images {
		if bytes.Equal(iter.hash[:], img.hash[:]) && iter.backColor == back {
			return iter, i, nil
		}
	}
	var decoded, _, err2 = image.Decode(buf)
	if err2 != nil {
		return pdfImage{}, -1,
			pdfError{id: 0xE64335, msg: "Image not decoded", val: err2.Error()}
	}
	img.backColor = back
	img.widthPx, img.heightPx, img.isGray, img.data = makeImage(decoded, back)
	ob.images = append(ob.images, img)
	return img, len(ob.images) - 1, nil
} //                                                                   loadImage

// makeImage encodes the source image in a PDF image data stream
func makeImage(source image.Image, back color.RGBA,
) (widthPx, heightPx int, isGray bool, ar []byte) {
	//
	// blends color into the background 'back', using opacity (alpha) value
	var blend = func(color, alpha uint32, back byte) byte {
		var c, a = float64(color), 65535 - float64(alpha) // range 0-65535
		return byte((c + (float64(back)*255-c)/65536*a) / 65536 * 255)
	}
	widthPx, heightPx = source.Bounds().Max.X, source.Bounds().Max.Y
	var model = source.ColorModel()
	isGray = model == color.GrayModel || model == color.Gray16Model
	for y := 0; y < heightPx; y++ {
		for x := 0; x < widthPx; x++ {
			var r, g, b, a = source.At(x, y).RGBA() //      value range: 0-65535
			switch {
			case isGray:
				ar = append(ar, byte(float64(r)))
			case a == 65535: //                                 if fully opaque:
				ar = append(ar, byte(r), byte(g), byte(b)) //    use pixel color
			case a == 0: //                                      if transparent:
				ar = append(ar, back.R, back.G, back.B) //  use background color
			default: //                                     if semi-transparent:
				ar = append(ar,
					blend(r, a, back.R), //          blend pixel and back colors
					blend(g, a, back.G), //             based on the alpha value
					blend(b, a, back.B))
			}
		}
	}
	return widthPx, heightPx, isGray, ar
} //                                                                   makeImage

// reservePage ensures there is at least one page in the PDF
func (ob *PDF) reservePage() *PDF {
	if len(ob.pages) == 0 {
		ob.AddPage()
	}
	return ob
} //                                                                  reservePage

// textWidthPt returns the width of text in points
func (ob *PDF) textWidthPt(s string) float64 {
	if s == "" {
		return 0
	}
	var w = 0.0
	for i, r := range s {
		if r < 0 || r > 255 {
			ob.putError(0xE31046, "Rune out of range",
				fmt.Sprintf("at %d = '%s'", i, string(r)))
			break
		}
		var id = ob.fonts[ob.ppage.fontID-1].builtInIndex
		if id >= 0 && id <= 9 {
			w += float64(pdfFontWidths[r][id])
		} else {
			w += 600 // for Courier font
		}
	}
	return w * ob.fontSizePt / 1000.0 * float64(ob.horzScaling) / 100.0
} //                                                                 textWidthPt

// -----------------------------------------------------------------------------
// # Internal Generation Methods (ob *PDF)

// nextObj increases the object serial no. and stores its offset in array
func (ob *PDF) nextObj() int {
	ob.objNo++
	for len(ob.objOffsets) <= ob.objNo {
		ob.objOffsets = append(ob.objOffsets, ob.content.Len())
	}
	return ob.objNo
} //                                                                     nextObj

// write writes formatted strings (like fmt.Sprintf) to the current page's
// content stream or to the final generated PDF, if there is no active page
func (ob *PDF) write(format string, args ...interface{}) *PDF {
	ob.reservePage()
	ob.pbuf.Write([]byte(fmt.Sprintf(format, args...)))
	return ob
} //                                                                       write

// writeCurve writes a Bézier curve using the 'c' PDF primitive.
// The starting point is the current (x, y) position.
// (x1, y1) and (x2, y2) are the two control points, (x3, y3) the end point.
func (ob *PDF) writeCurve(x1, y1, x2, y2, x3, y3 float64) *PDF {
	return ob.write(" %.3f %.3f %.3f %.3f %.3f %.3f c", x1, y1, x2, y2, x3, y3)
	// x1 y1 x2 y2 x3 y3 c: append cubic Bézier curve to the current path
} //                                                                  writeCurve

// writeEndobj writes 'endobj' (PDF object end marker)
func (ob *PDF) writeEndobj() *PDF {
	return ob.write(">>\nendobj\n")
} //                                                                 writeEndobj

// writeMode sets the stroking or non-stroking color and line width.
// 'fill' arg specifies non-stroking (true) or stroking mode (none/false)
func (ob *PDF) writeMode(fill ...bool) (mode string) {
	ob.reservePage()
	mode = "S" // S: stroke path (for lines)
	if len(fill) > 0 && fill[0] {
		mode = "b" // b: fill / text
		if pv := &ob.ppage.nonStrokeColor; *pv != ob.color {
			*pv = ob.color
			ob.write(" %.3f %.3f %.3f rg\n", // rg: set non-stroking/text color
				float64(pv.R)/255, float64(pv.G)/255, float64(pv.B)/255)
		}
	}
	if pv := &ob.ppage.strokeColor; *pv != ob.color {
		*pv = ob.color
		ob.write("%.3f %.3f %.3f RG\n", // RG: set stroke (line) color
			float64(pv.R)/255, float64(pv.G)/255, float64(pv.B)/255)
	}
	if pv := &ob.ppage.lineWidth; int(*pv*100) != int(ob.lineWidth*100) {
		*pv = ob.lineWidth
		ob.write("%.3f w\n", float64(*pv)) // n0 w: set line width to n0
	}
	return mode
} //                                                                   writeMode

// writeObj writes an object header. objType must start with '/', e.g. /Catalog
func (ob *PDF) writeObj(objType string) *PDF {
	return ob.write("%d 0 obj<</Type%s", ob.nextObj(), objType)
} //                                                                    writeObj

// writePages writes all PDF pages
func (ob *PDF) writePages(pagesIndex, fontsIndex, imagesIndex int) *PDF {
	ob.writeObj("/Pages").write("/Count %d/MediaBox[0 0 %d %d]",
		len(ob.pages), int(ob.paperSize.widthPt), int(ob.paperSize.heightPt))
	//                                                        write page numbers
	if len(ob.pages) > 0 {
		var pageObjNo = pagesIndex
		ob.write("/Kids[")
		for i := range ob.pages {
			if i > 0 {
				ob.write(" ")
			}
			ob.write("%d 0 R", pageObjNo)
			pageObjNo += 2 //                           1 for page, 1 for stream
		}
		ob.write("]")
	}
	ob.writeEndobj()
	for _, pg := range ob.pages { //                             write each page
		ob.writeObj("/Page").write("/Parent 2 0 R/Contents %d 0 R", ob.objNo+1)
		if len(pg.fontIDs) > 0 || len(pg.imageNos) > 0 {
			ob.write("/Resources<<")
		}
		if len(pg.fontIDs) > 0 {
			ob.write("/Font <<")
			for fontNo := range ob.fonts {
				ob.write("/FNT%d %d 0 R", fontNo+1, fontsIndex+fontNo)
			}
			ob.write(">>")
		}
		if len(pg.imageNos) > 0 {
			ob.write("/XObject<<")
			for imageNo := range pg.imageNos {
				ob.write("/IMG%d %d 0 R", imageNo, imagesIndex+imageNo)
			}
			ob.write(">>")
		}
		if len(pg.fontIDs) > 0 || len(pg.imageNos) > 0 {
			ob.write(">>")
		}
		ob.writeEndobj().writeStream([]byte(pg.content.String()))
	}
	return ob
} //                                                                  writePages

// writeStream outputs a stream object to the document's main buffer
func (ob *PDF) writeStream(content []byte) *PDF {
	return ob.write("%d 0 obj <<", ob.nextObj()).writeStreamData(content)
} //                                                                 writeStream

// writeStreamData writes a stream or image stream
func (ob *PDF) writeStreamData(ar []byte) *PDF {
	var s string // filter
	if ob.compression {
		var buf bytes.Buffer
		var wr = zlib.NewWriter(&buf)
		var _, err = wr.Write([]byte(ar))
		if err != nil {
			return ob.putError(0xE782A2, "Failed compressing", err.Error())
		}
		wr.Close() // don't defer, close before reading Bytes()
		ar = buf.Bytes()
		s = "/Filter/FlateDecode"
	}
	return ob.write("%s/Length %d>>stream\n%s\nendstream\n", s, len(ar), ar)
} //                                                             writeStreamData

// -----------------------------------------------------------------------------
// # Internal Functions (just attached to PDF, but not using it)

// escape escapes special characters '(', '(' and '\' in strings
// in order to avoid them interfering with PDF commands
func (*PDF) escape(s string) []byte {
	var has = strings.Contains
	if !has(s, "(") && !has(s, ")") && !has(s, "\\") {
		return []byte(s)
	}
	var buf = bytes.NewBuffer(make([]byte, 0, len(s)))
	for _, r := range s {
		if r == '(' || r == ')' || r == '\\' {
			buf.WriteRune('\\')
		}
		buf.WriteRune(r)
	}
	return buf.Bytes()
} //                                                                      escape

// isWhiteSpace returns true if all the chars. in 's' are white-spaces
func (*PDF) isWhiteSpace(s string) bool {
	for _, r := range s {
		if !unicode.IsSpace(r) {
			return false
		}
	}
	return len(s) > 0
} //                                                                isWhiteSpace

// splitLines splits 's' into several lines using line breaks in 's'
func (*PDF) splitLines(s string) []string {
	var split = func(ar []string, sep string) (ret []string) {
		for _, iter := range ar {
			if strings.Contains(iter, sep) {
				ret = append(ret, strings.Split(iter, sep)...)
				continue
			}
			ret = append(ret, iter)
		}
		return ret
	}
	return split(split(split([]string{s}, "\r\n"), "\r"), "\n")
} //                                                                  splitLines

// toUpperLettersDigits returns letters and digits from s, in upper case
func (*PDF) toUpperLettersDigits(s, extras string) string {
	var buf = bytes.NewBuffer(make([]byte, 0, len(s)))
	for _, r := range strings.ToUpper(s) {
		if unicode.IsLetter(r) || unicode.IsDigit(r) ||
			strings.ContainsRune(extras, r) {
			buf.WriteRune(r)
		}
	}
	return buf.String()
} //                                                        toUpperLettersDigits

// getPaperSize returns a pdfPaperSize based on the specified paper name.
// Specify custom paper sizes using "width x height", e.g. "9cm x 20cm"
// If the paper size is not found, returns a zero-initialized structure
func (ob *PDF) getPaperSize(name string) (pdfPaperSize, error) {
	var s = strings.ToUpper(name)
	if strings.Contains(s, " X ") {
		var wh = strings.Split(s, " X ")
		var w, err = ob.ToPoints(wh[0])
		if err, isT := err.(pdfError); isT {
			return pdfPaperSize{}, err
		}
		var h float64
		h, err = ob.ToPoints(wh[1])
		if err, isT := err.(pdfError); isT {
			return pdfPaperSize{}, err
		}
		return pdfPaperSize{s, w, h}, nil
	}
	s = ob.toUpperLettersDigits(s, "-")
	var landscape = strings.HasSuffix(s, "-L")
	s = ob.toUpperLettersDigits(s, "")
	if landscape {
		s = s[:len(s)-1] // "-" is already removed above. now remove the "L"
	}
	var wh, found = pdfStandardPaperSizes[s]
	if !found {
		return pdfPaperSize{},
			pdfError{id: 0xEE42FB, msg: "Unknown paper size", val: name}
	}
	// convert mm to points: div by 25.4mm/inch; mul by 72 points/inch
	var w, h = float64(wh[0]) / 25.4 * 72, float64(wh[1]) / 25.4 * 72
	if landscape {
		return pdfPaperSize{s + "-L", h, w}, nil
	}
	return pdfPaperSize{s, w, h}, nil
} //                                                                getPaperSize

// getPointsPerUnit returns number of points per named measurement unit
func (ob *PDF) getPointsPerUnit(unitName string) (ret float64, err error) {
	switch ob.toUpperLettersDigits(unitName, `"`) {
	case "CM":
		ret = 28.3464566929134 //          " / 2.54cm per " * 72 points per inch
	case "IN", "INCH", "INCHES", `"`:
		ret = 72.0 //                                         72 points per inch
	case "MM":
		ret = 2.83464566929134 //     1 inch / 25.4mm per " * 72 points per inch
	case "PT", "POINT", "POINTS":
		ret = 1.0 //                                                     1 point
	case "TW", "TWIP", "TWIPS":
		ret = 0.05 //                               1 point / 20 twips per point
	default:
		err = pdfError{id: 0xEE34DA, msg: "Unknown unit name", val: unitName}
	}
	return ret, err
} //                                                            getPointsPerUnit

// putError appends an error to the errors collection
func (ob *PDF) putError(id int, msg, val string) *PDF {
	var fn string //                                  get the public method name
	for i := 0; i < 10; i++ {
		var programCounter, _, _, _ = runtime.Caller(i)
		fn = runtime.FuncForPC(programCounter).Name()
		fn = fn[strings.LastIndex(fn, ".")+1:]
		if unicode.IsLower(rune(fn[0])) {
			continue
		}
		break
	}
	ob.errors = append(ob.errors,
		pdfError{id: id, src: fn, msg: msg, val: val})
	return ob
} //                                                                    putError

// -----------------------------------------------------------------------------
// # Constants

// PDFColorNames maps web (X11) color names to values.
// From https://en.wikipedia.org/wiki/X11_color_names
var PDFColorNames = map[string]color.RGBA{
	"ALICEBLUE":            {R: 240, G: 248, B: 255}, // #F0F8FF
	"ANTIQUEWHITE":         {R: 250, G: 235, B: 215}, // #FAEBD7
	"AQUA":                 {R: 000, G: 255, B: 255}, // #00FFFF
	"AQUAMARINE":           {R: 127, G: 255, B: 212}, // #7FFFD4
	"AZURE":                {R: 240, G: 255, B: 255}, // #F0FFFF
	"BEIGE":                {R: 245, G: 245, B: 220}, // #F5F5DC
	"BISQUE":               {R: 255, G: 228, B: 196}, // #FFE4C4
	"BLACK":                {R: 000, G: 000, B: 000}, // #000000
	"BLANCHEDALMOND":       {R: 255, G: 235, B: 205}, // #FFEBCD
	"BLUE":                 {R: 000, G: 000, B: 255}, // #0000FF
	"BLUEVIOLET":           {R: 138, G: 43, B: 226},  // #8A2BE2
	"BROWN":                {R: 165, G: 42, B: 42},   // #A52A2A
	"BURLYWOOD":            {R: 222, G: 184, B: 135}, // #DEB887
	"CADETBLUE":            {R: 95, G: 158, B: 160},  // #5F9EA0
	"CHARTREUSE":           {R: 127, G: 255, B: 000}, // #7FFF00
	"CHOCOLATE":            {R: 210, G: 105, B: 30},  // #D2691E
	"CORAL":                {R: 255, G: 127, B: 80},  // #FF7F50
	"CORNFLOWERBLUE":       {R: 100, G: 149, B: 237}, // #6495ED
	"CORNSILK":             {R: 255, G: 248, B: 220}, // #FFF8DC
	"CRIMSON":              {R: 220, G: 20, B: 60},   // #DC143C
	"CYAN":                 {R: 000, G: 255, B: 255}, // #00FFFF
	"DARKBLUE":             {R: 000, G: 000, B: 139}, // #00008B
	"DARKCYAN":             {R: 000, G: 139, B: 139}, // #008B8B
	"DARKGOLDENROD":        {R: 184, G: 134, B: 11},  // #B8860B
	"DARKGRAY":             {R: 169, G: 169, B: 169}, // #A9A9A9
	"DARKGREEN":            {R: 000, G: 100, B: 000}, // #006400
	"DARKKHAKI":            {R: 189, G: 183, B: 107}, // #BDB76B
	"DARKMAGENTA":          {R: 139, G: 000, B: 139}, // #8B008B
	"DARKOLIVEGREEN":       {R: 85, G: 107, B: 47},   // #556B2F
	"DARKORANGE":           {R: 255, G: 140, B: 000}, // #FF8C00
	"DARKORCHID":           {R: 153, G: 50, B: 204},  // #9932CC
	"DARKRED":              {R: 139, G: 000, B: 000}, // #8B0000
	"DARKSALMON":           {R: 233, G: 150, B: 122}, // #E9967A
	"DARKSEAGREEN":         {R: 143, G: 188, B: 143}, // #8FBC8F
	"DARKSLATEBLUE":        {R: 72, G: 61, B: 139},   // #483D8B
	"DARKSLATEGRAY":        {R: 47, G: 79, B: 79},    // #2F4F4F
	"DARKTURQUOISE":        {R: 000, G: 206, B: 209}, // #00CED1
	"DARKVIOLET":           {R: 148, G: 000, B: 211}, // #9400D3
	"DEEPPINK":             {R: 255, G: 20, B: 147},  // #FF1493
	"DEEPSKYBLUE":          {R: 000, G: 191, B: 255}, // #00BFFF
	"DIMGRAY":              {R: 105, G: 105, B: 105}, // #696969
	"DODGERBLUE":           {R: 30, G: 144, B: 255},  // #1E90FF
	"FIREBRICK":            {R: 178, G: 34, B: 34},   // #B22222
	"FLORALWHITE":          {R: 255, G: 250, B: 240}, // #FFFAF0
	"FORESTGREEN":          {R: 34, G: 139, B: 34},   // #228B22
	"FUCHSIA":              {R: 255, G: 000, B: 255}, // #FF00FF
	"GAINSBORO":            {R: 220, G: 220, B: 220}, // #DCDCDC
	"GHOSTWHITE":           {R: 248, G: 248, B: 255}, // #F8F8FF
	"GOLD":                 {R: 255, G: 215, B: 000}, // #FFD700
	"GOLDENROD":            {R: 218, G: 165, B: 32},  // #DAA520
	"GRAY":                 {R: 190, G: 190, B: 190}, // #BEBEBE X11 Version
	"GREEN":                {R: 000, G: 255, B: 000}, // #00FF00 X11 Version
	"GREENYELLOW":          {R: 173, G: 255, B: 47},  // #ADFF2F
	"HONEYDEW":             {R: 240, G: 255, B: 240}, // #F0FFF0
	"HOTPINK":              {R: 255, G: 105, B: 180}, // #FF69B4
	"INDIANRED":            {R: 205, G: 92, B: 92},   // #CD5C5C
	"INDIGO":               {R: 75, G: 000, B: 130},  // #4B0082
	"IVORY":                {R: 255, G: 255, B: 240}, // #FFFFF0
	"KHAKI":                {R: 240, G: 230, B: 140}, // #F0E68C
	"LAVENDER":             {R: 230, G: 230, B: 250}, // #E6E6FA
	"LAVENDERBLUSH":        {R: 255, G: 240, B: 245}, // #FFF0F5
	"LAWNGREEN":            {R: 124, G: 252, B: 000}, // #7CFC00
	"LEMONCHIFFON":         {R: 255, G: 250, B: 205}, // #FFFACD
	"LIGHTBLUE":            {R: 173, G: 216, B: 230}, // #ADD8E6
	"LIGHTCORAL":           {R: 240, G: 128, B: 128}, // #F08080
	"LIGHTCYAN":            {R: 224, G: 255, B: 255}, // #E0FFFF
	"LIGHTGOLDENRODYELLOW": {R: 250, G: 250, B: 210}, // #FAFAD2
	"LIGHTGRAY":            {R: 211, G: 211, B: 211}, // #D3D3D3
	"LIGHTGREEN":           {R: 144, G: 238, B: 144}, // #90EE90
	"LIGHTPINK":            {R: 255, G: 182, B: 193}, // #FFB6C1
	"LIGHTSALMON":          {R: 255, G: 160, B: 122}, // #FFA07A
	"LIGHTSEAGREEN":        {R: 32, G: 178, B: 170},  // #20B2AA
	"LIGHTSKYBLUE":         {R: 135, G: 206, B: 250}, // #87CEFA
	"LIGHTSLATEGRAY":       {R: 119, G: 136, B: 153}, // #778899
	"LIGHTSTEELBLUE":       {R: 176, G: 196, B: 222}, // #B0C4DE
	"LIGHTYELLOW":          {R: 255, G: 255, B: 224}, // #FFFFE0
	"LIME":                 {R: 000, G: 255, B: 000}, // #00FF00
	"LIMEGREEN":            {R: 50, G: 205, B: 50},   // #32CD32
	"LINEN":                {R: 250, G: 240, B: 230}, // #FAF0E6
	"MAGENTA":              {R: 255, G: 000, B: 255}, // #FF00FF
	"MAROON":               {R: 176, G: 48, B: 96},   // #B03060 X11 Version
	"MEDIUMAQUAMARINE":     {R: 102, G: 205, B: 170}, // #66CDAA
	"MEDIUMBLUE":           {R: 000, G: 000, B: 205}, // #0000CD
	"MEDIUMORCHID":         {R: 186, G: 85, B: 211},  // #BA55D3
	"MEDIUMPURPLE":         {R: 147, G: 112, B: 219}, // #9370DB
	"MEDIUMSEAGREEN":       {R: 60, G: 179, B: 113},  // #3CB371
	"MEDIUMSLATEBLUE":      {R: 123, G: 104, B: 238}, // #7B68EE
	"MEDIUMSPRINGGREEN":    {R: 000, G: 250, B: 154}, // #00FA9A
	"MEDIUMTURQUOISE":      {R: 72, G: 209, B: 204},  // #48D1CC
	"MEDIUMVIOLETRED":      {R: 199, G: 21, B: 133},  // #C71585
	"MIDNIGHTBLUE":         {R: 25, G: 25, B: 112},   // #191970
	"MINTCREAM":            {R: 245, G: 255, B: 250}, // #F5FFFA
	"MISTYROSE":            {R: 255, G: 228, B: 225}, // #FFE4E1
	"MOCCASIN":             {R: 255, G: 228, B: 181}, // #FFE4B5
	"NAVAJOWHITE":          {R: 255, G: 222, B: 173}, // #FFDEAD
	"NAVY":                 {R: 000, G: 000, B: 128}, // #000080
	"OLDLACE":              {R: 253, G: 245, B: 230}, // #FDF5E6
	"OLIVE":                {R: 128, G: 128, B: 000}, // #808000
	"OLIVEDRAB":            {R: 107, G: 142, B: 35},  // #6B8E23
	"ORANGE":               {R: 255, G: 165, B: 000}, // #FFA500
	"ORANGERED":            {R: 255, G: 69, B: 000},  // #FF4500
	"ORCHID":               {R: 218, G: 112, B: 214}, // #DA70D6
	"PALEGOLDENROD":        {R: 238, G: 232, B: 170}, // #EEE8AA
	"PALEGREEN":            {R: 152, G: 251, B: 152}, // #98FB98
	"PALETURQUOISE":        {R: 175, G: 238, B: 238}, // #AFEEEE
	"PALEVIOLETRED":        {R: 219, G: 112, B: 147}, // #DB7093
	"PAPAYAWHIP":           {R: 255, G: 239, B: 213}, // #FFEFD5
	"PEACHPUFF":            {R: 255, G: 218, B: 185}, // #FFDAB9
	"PERU":                 {R: 205, G: 133, B: 63},  // #CD853F
	"PINK":                 {R: 255, G: 192, B: 203}, // #FFC0CB
	"PLUM":                 {R: 221, G: 160, B: 221}, // #DDA0DD
	"POWDERBLUE":           {R: 176, G: 224, B: 230}, // #B0E0E6
	"PURPLE":               {R: 160, G: 32, B: 240},  // #A020F0 X11 Version
	"REBECCAPURPLE":        {R: 102, G: 51, B: 153},  // #663399
	"RED":                  {R: 255, G: 000, B: 000}, // #FF0000
	"ROSYBROWN":            {R: 188, G: 143, B: 143}, // #BC8F8F
	"ROYALBLUE":            {R: 65, G: 105, B: 225},  // #4169E1
	"SADDLEBROWN":          {R: 139, G: 69, B: 19},   // #8B4513
	"SALMON":               {R: 250, G: 128, B: 114}, // #FA8072
	"SANDYBROWN":           {R: 244, G: 164, B: 96},  // #F4A460
	"SEAGREEN":             {R: 46, G: 139, B: 87},   // #2E8B57
	"SEASHELL":             {R: 255, G: 245, B: 238}, // #FFF5EE
	"SIENNA":               {R: 160, G: 82, B: 45},   // #A0522D
	"SILVER":               {R: 192, G: 192, B: 192}, // #C0C0C0
	"SKYBLUE":              {R: 135, G: 206, B: 235}, // #87CEEB
	"SLATEBLUE":            {R: 106, G: 90, B: 205},  // #6A5ACD
	"SLATEGRAY":            {R: 112, G: 128, B: 144}, // #708090
	"SNOW":                 {R: 255, G: 250, B: 250}, // #FFFAFA
	"SPRINGGREEN":          {R: 000, G: 255, B: 127}, // #00FF7F
	"STEELBLUE":            {R: 70, G: 130, B: 180},  // #4682B4
	"TAN":                  {R: 210, G: 180, B: 140}, // #D2B48C
	"TEAL":                 {R: 000, G: 128, B: 128}, // #008080
	"THISTLE":              {R: 216, G: 191, B: 216}, // #D8BFD8
	"TOMATO":               {R: 255, G: 99, B: 71},   // #FF6347
	"TURQUOISE":            {R: 64, G: 224, B: 208},  // #40E0D0
	"VIOLET":               {R: 238, G: 130, B: 238}, // #EE82EE
	"WEBGRAY":              {R: 128, G: 128, B: 128}, // #808080 Web Version
	"WEBGREEN":             {R: 000, G: 128, B: 000}, // #008000 Web Version
	"WEBMAROON":            {R: 127, G: 000, B: 000}, // #7F0000 Web Version
	"WEBPURPLE":            {R: 127, G: 000, B: 127}, // #7F007F Web Version
	"WHEAT":                {R: 245, G: 222, B: 179}, // #F5DEB3
	"WHITE":                {R: 255, G: 255, B: 255}, // #FFFFFF
	"WHITESMOKE":           {R: 245, G: 245, B: 245}, // #F5F5F5
	"YELLOW":               {R: 255, G: 255, B: 000}, // #FFFF00
	"YELLOWGREEN":          {R: 154, G: 205, B: 50},  // #9ACD32
} //                                                               PDFColorNames

// -----------------------------------------------------------------------------
// # Internal Constants

var pdfBlack = color.RGBA{A: 255}

// pdfFontNames contains font names available on all PDF implementations
var pdfFontNames = []string{
	"Helvetica",             // 0
	"Helvetica-Bold",        // 1
	"Helvetica-BoldOblique", // 2
	"Helvetica-Oblique",     // 3
	"Symbol",                // 4
	"Times-Bold",            // 5
	"Times-BoldItalic",      // 6
	"Times-Italic",          // 7
	"Times-Roman",           // 8
	"ZapfDingbats",          // 9
	"Courier",               // <- keep fixed-width Courier
	"Courier-Bold",          // font at the end of the list
	"Courier-BoldOblique",
	"Courier-Oblique",
} //                                                                pdfFontNames

// pdfFontWidths specifies widths of built-in fonts,
// in thousandths of a point per point of height
var pdfFontWidths = [][]int{
	{278, 278, 278, 278, 250, 250, 250, 250, 250, 000},         // 000
	{278, 278, 278, 278, 250, 250, 250, 250, 250, 000},         // 001
	{278, 278, 278, 278, 250, 250, 250, 250, 250, 000},         // 002
	{278, 278, 278, 278, 250, 250, 250, 250, 250, 000},         // 003
	{278, 278, 278, 278, 250, 250, 250, 250, 250, 000},         // 004
	{278, 278, 278, 278, 250, 250, 250, 250, 250, 000},         // 005
	{278, 278, 278, 278, 250, 250, 250, 250, 250, 000},         // 006
	{278, 278, 278, 278, 250, 250, 250, 250, 250, 000},         // 007
	{278, 278, 278, 278, 250, 250, 250, 250, 250, 000},         // 008
	{278, 278, 278, 278, 250, 250, 250, 250, 250, 000},         // 009
	{278, 278, 278, 278, 250, 250, 250, 250, 250, 000},         // 010
	{278, 278, 278, 278, 250, 250, 250, 250, 250, 000},         // 011
	{278, 278, 278, 278, 250, 250, 250, 250, 250, 000},         // 012
	{278, 278, 278, 278, 250, 250, 250, 250, 250, 000},         // 013
	{278, 278, 278, 278, 250, 250, 250, 250, 250, 000},         // 014
	{278, 278, 278, 278, 250, 250, 250, 250, 250, 000},         // 015
	{278, 278, 278, 278, 250, 250, 250, 250, 250, 000},         // 016
	{278, 278, 278, 278, 250, 250, 250, 250, 250, 000},         // 017
	{278, 278, 278, 278, 250, 250, 250, 250, 250, 000},         // 018
	{278, 278, 278, 278, 250, 250, 250, 250, 250, 000},         // 019
	{278, 278, 278, 278, 250, 250, 250, 250, 250, 000},         // 020
	{278, 278, 278, 278, 250, 250, 250, 250, 250, 000},         // 021
	{278, 278, 278, 278, 250, 250, 250, 250, 250, 000},         // 022
	{278, 278, 278, 278, 250, 250, 250, 250, 250, 000},         // 023
	{278, 278, 278, 278, 250, 250, 250, 250, 250, 000},         // 024
	{278, 278, 278, 278, 250, 250, 250, 250, 250, 000},         // 025
	{278, 278, 278, 278, 250, 250, 250, 250, 250, 000},         // 026
	{278, 278, 278, 278, 250, 250, 250, 250, 250, 000},         // 027
	{278, 278, 278, 278, 250, 250, 250, 250, 250, 000},         // 028
	{278, 278, 278, 278, 250, 250, 250, 250, 250, 000},         // 029
	{278, 278, 278, 278, 250, 250, 250, 250, 250, 000},         // 030
	{278, 278, 278, 278, 250, 250, 250, 250, 250, 000},         // 031
	{278, 278, 278, 278, 250, 250, 250, 250, 250, 278},         // 032
	{278, 333, 333, 278, 333, 333, 389, 333, 333, 974},         // 033 !
	{355, 474, 474, 355, 713, 555, 555, 420, 408, 961},         // 034 "
	{556, 556, 556, 556, 500, 500, 500, 500, 500, 974},         // 035 #
	{556, 556, 556, 556, 549, 500, 500, 500, 500, 980},         // 036 $
	{889, 889, 889, 889, 833, 000, 833, 833, 833, 719},         // 037 %
	{667, 722, 722, 667, 778, 833, 778, 778, 778, 789},         // 038 &
	{191, 238, 238, 191, 439, 278, 278, 214, 180, 790},         // 039 '
	{333, 333, 333, 333, 333, 333, 333, 333, 333, 791},         // 040 (
	{333, 333, 333, 333, 333, 333, 333, 333, 333, 690},         // 041 )
	{389, 389, 389, 389, 500, 500, 500, 500, 500, 960},         // 042 *
	{584, 584, 584, 584, 549, 570, 570, 675, 564, 939},         // 043 +
	{278, 278, 278, 278, 250, 250, 250, 250, 250, 549},         // 044 ,
	{333, 333, 333, 333, 549, 333, 333, 333, 333, 855},         // 045 -
	{278, 278, 278, 278, 250, 250, 250, 250, 250, 911},         // 046 .
	{278, 278, 278, 278, 278, 278, 278, 278, 278, 933},         // 047 /
	{556, 556, 556, 556, 500, 500, 500, 500, 500, 911},         // 048 000
	{556, 556, 556, 556, 500, 500, 500, 500, 500, 945},         // 049 1
	{556, 556, 556, 556, 500, 500, 500, 500, 500, 974},         // 050 2
	{556, 556, 556, 556, 500, 500, 500, 500, 500, 755},         // 051 3
	{556, 556, 556, 556, 500, 500, 500, 500, 500, 846},         // 052 4
	{556, 556, 556, 556, 500, 500, 500, 500, 500, 762},         // 053 5
	{556, 556, 556, 556, 500, 500, 500, 500, 500, 761},         // 054 6
	{556, 556, 556, 556, 500, 500, 500, 500, 500, 571},         // 055 7
	{556, 556, 556, 556, 500, 500, 500, 500, 500, 677},         // 056 8
	{556, 556, 556, 556, 500, 500, 500, 500, 500, 763},         // 057 9
	{278, 333, 333, 278, 278, 333, 333, 333, 278, 760},         // 058 :
	{278, 333, 333, 278, 278, 333, 333, 333, 278, 759},         // 059 ;
	{584, 584, 584, 584, 549, 570, 570, 675, 564, 754},         // 060 <
	{584, 584, 584, 584, 549, 570, 570, 675, 564, 494},         // 061 =
	{584, 584, 584, 584, 549, 570, 570, 675, 564, 552},         // 062 >
	{556, 611, 611, 556, 444, 500, 500, 500, 444, 537},         // 063 ?
	{1015, 975, 975, 1015, 549, 930, 832, 920, 921, 577},       // 064 @
	{667, 722, 722, 667, 722, 722, 667, 611, 722, 692},         // 065 A
	{667, 722, 722, 667, 667, 667, 667, 611, 667, 786},         // 066 B
	{722, 722, 722, 722, 722, 722, 667, 667, 667, 788},         // 067 C
	{722, 722, 722, 722, 612, 722, 722, 722, 722, 788},         // 068 D
	{667, 667, 667, 667, 611, 667, 667, 611, 611, 790},         // 069 E
	{611, 611, 611, 611, 763, 611, 667, 611, 556, 793},         // 070 F
	{778, 778, 778, 778, 603, 778, 722, 722, 722, 794},         // 071 G
	{722, 722, 722, 722, 722, 778, 778, 722, 722, 816},         // 072 H
	{278, 278, 278, 278, 333, 389, 389, 333, 333, 823},         // 073 I
	{500, 556, 556, 500, 631, 500, 500, 444, 389, 789},         // 074 J
	{667, 722, 722, 667, 722, 778, 667, 667, 722, 841},         // 075 K
	{556, 611, 611, 556, 686, 667, 611, 556, 611, 823},         // 076 L
	{833, 833, 833, 833, 889, 944, 889, 833, 889, 833},         // 077 M
	{722, 722, 722, 722, 722, 722, 722, 667, 722, 816},         // 078 N
	{778, 778, 778, 778, 722, 778, 722, 722, 722, 831},         // 079 O
	{667, 667, 667, 667, 768, 611, 611, 611, 556, 923},         // 080 P
	{778, 778, 778, 778, 741, 778, 722, 722, 722, 744},         // 081 Q
	{722, 722, 722, 722, 556, 722, 667, 611, 667, 723},         // 082 R
	{667, 667, 667, 667, 592, 556, 556, 500, 556, 749},         // 083 S
	{611, 611, 611, 611, 611, 667, 611, 556, 611, 790},         // 084 T
	{722, 722, 722, 722, 690, 722, 722, 722, 722, 792},         // 085 U
	{667, 667, 667, 667, 439, 722, 667, 611, 722, 695},         // 086 V
	{944, 944, 944, 944, 768, 1000, 889, 833, 944, 776},        // 087 W
	{667, 667, 667, 667, 645, 722, 667, 611, 722, 768},         // 088 X
	{667, 667, 667, 667, 795, 722, 611, 556, 722, 792},         // 089 Y
	{611, 611, 611, 611, 611, 667, 611, 556, 611, 759},         // 090 Z
	{278, 333, 333, 278, 333, 333, 333, 389, 333, 707},         // 091 [
	{278, 278, 278, 278, 863, 278, 278, 278, 278, 708},         // 092 \
	{278, 333, 333, 278, 333, 333, 333, 389, 333, 682},         // 093 ]
	{469, 584, 584, 469, 658, 581, 570, 422, 469, 701},         // 094 ^
	{556, 556, 556, 556, 500, 500, 500, 500, 500, 826},         // 095 _
	{333, 333, 333, 333, 500, 333, 333, 333, 333, 815},         // 096 \x60
	{556, 556, 556, 556, 631, 500, 500, 500, 444, 789},         // 097 a
	{556, 611, 611, 556, 549, 556, 500, 500, 500, 789},         // 098 b
	{500, 556, 556, 500, 549, 444, 444, 444, 444, 707},         // 099 c
	{556, 611, 611, 556, 494, 556, 500, 500, 500, 687},         // 100 d
	{556, 556, 556, 556, 439, 444, 444, 444, 444, 696},         // 101 e
	{278, 333, 333, 278, 521, 333, 333, 278, 333, 689},         // 102 f
	{556, 611, 611, 556, 411, 500, 500, 500, 500, 786},         // 103 g
	{556, 611, 611, 556, 603, 556, 556, 500, 500, 787},         // 104 h
	{222, 278, 278, 222, 329, 278, 278, 278, 278, 713},         // 105 i
	{222, 278, 278, 222, 603, 333, 278, 278, 278, 791},         // 106 j
	{500, 556, 556, 500, 549, 556, 500, 444, 500, 785},         // 107 k
	{222, 278, 278, 222, 549, 278, 278, 278, 278, 791},         // 108 l
	{833, 889, 889, 833, 576, 833, 778, 722, 778, 873},         // 109 m
	{556, 611, 611, 556, 521, 556, 556, 500, 500, 761},         // 110 n
	{556, 611, 611, 556, 549, 500, 500, 500, 500, 762},         // 111 o
	{556, 611, 611, 556, 549, 556, 500, 500, 500, 762},         // 112 p
	{556, 611, 611, 556, 521, 556, 500, 500, 500, 759},         // 113 q
	{333, 389, 389, 333, 549, 444, 389, 389, 333, 759},         // 114 r
	{500, 556, 556, 500, 603, 389, 389, 389, 389, 892},         // 115 s
	{278, 333, 333, 278, 439, 333, 278, 278, 278, 892},         // 116 t
	{556, 611, 611, 556, 576, 556, 556, 500, 500, 788},         // 117 u
	{500, 556, 556, 500, 713, 500, 444, 444, 500, 784},         // 118 v
	{722, 778, 778, 722, 686, 722, 667, 667, 722, 438},         // 119 w
	{500, 556, 556, 500, 493, 500, 500, 444, 500, 138},         // 120 x
	{500, 556, 556, 500, 686, 500, 444, 444, 500, 277},         // 121 y
	{500, 500, 500, 500, 494, 444, 389, 389, 444, 415},         // 122 z
	{334, 389, 389, 334, 480, 394, 348, 400, 480, 392},         // 123 {
	{260, 280, 280, 260, 200, 220, 220, 275, 200, 392},         // 124 |
	{334, 389, 389, 334, 480, 394, 348, 400, 480, 668},         // 125 }
	{584, 584, 584, 584, 549, 520, 570, 541, 541, 668},         // 126 ~
	{350, 350, 350, 350, 000, 350, 350, 350, 350, 000},         // 127
	{556, 556, 556, 556, 000, 500, 500, 500, 500, 390},         // 128
	{350, 350, 350, 350, 000, 350, 350, 350, 350, 390},         // 129
	{222, 278, 278, 222, 000, 333, 333, 333, 333, 317},         // 130
	{556, 556, 556, 556, 000, 500, 500, 500, 500, 317},         // 131
	{333, 500, 500, 333, 000, 500, 500, 556, 444, 276},         // 132
	{1000, 1000, 1000, 1000, 000, 1000, 1000, 889, 1000, 276},  // 133
	{556, 556, 556, 556, 000, 500, 500, 500, 500, 509},         // 134
	{556, 556, 556, 556, 000, 500, 500, 500, 500, 509},         // 135
	{333, 333, 333, 333, 000, 333, 333, 333, 333, 410},         // 136
	{1000, 1000, 1000, 1000, 000, 1000, 1000, 1000, 1000, 410}, // 137
	{667, 667, 667, 667, 000, 556, 556, 500, 556, 234},         // 138
	{333, 333, 333, 333, 000, 333, 333, 333, 333, 234},         // 139
	{1000, 1000, 1000, 1000, 000, 1000, 944, 944, 889, 334},    // 140
	{350, 350, 350, 350, 000, 350, 350, 350, 350, 334},         // 141
	{611, 611, 611, 611, 000, 667, 611, 556, 611, 000},         // 142
	{350, 350, 350, 350, 000, 350, 350, 350, 350, 000},         // 143
	{350, 350, 350, 350, 000, 350, 350, 350, 350, 000},         // 144
	{222, 278, 278, 222, 000, 333, 333, 333, 333, 000},         // 145
	{222, 278, 278, 222, 000, 333, 333, 333, 333, 000},         // 146
	{333, 500, 500, 333, 000, 500, 500, 556, 444, 000},         // 147
	{333, 500, 500, 333, 000, 500, 500, 556, 444, 000},         // 148
	{350, 350, 350, 350, 000, 350, 350, 350, 350, 000},         // 149
	{556, 556, 556, 556, 000, 500, 500, 500, 500, 000},         // 150
	{1000, 1000, 1000, 1000, 000, 1000, 1000, 889, 1000, 000},  // 151
	{333, 333, 333, 333, 000, 333, 333, 333, 333, 000},         // 152
	{1000, 1000, 1000, 1000, 000, 1000, 1000, 980, 980, 000},   // 153
	{500, 556, 556, 500, 000, 389, 389, 389, 389, 000},         // 154
	{333, 333, 333, 333, 000, 333, 333, 333, 333, 000},         // 155
	{944, 944, 944, 944, 000, 722, 722, 667, 722, 000},         // 156
	{350, 350, 350, 350, 000, 350, 350, 350, 350, 000},         // 157
	{500, 500, 500, 500, 000, 444, 389, 389, 444, 000},         // 158
	{667, 667, 667, 667, 000, 722, 611, 556, 722, 000},         // 159
	{278, 278, 278, 278, 750, 250, 250, 250, 250, 000},         // 160
	{333, 333, 333, 333, 620, 333, 389, 389, 333, 732},         // 161
	{556, 556, 556, 556, 247, 500, 500, 500, 500, 544},         // 162
	{556, 556, 556, 556, 549, 500, 500, 500, 500, 544},         // 163
	{556, 556, 556, 556, 167, 500, 500, 500, 500, 910},         // 164
	{556, 556, 556, 556, 713, 500, 500, 500, 500, 667},         // 165
	{260, 280, 280, 260, 500, 220, 220, 275, 200, 760},         // 166
	{556, 556, 556, 556, 753, 500, 500, 500, 500, 760},         // 167
	{333, 333, 333, 333, 753, 333, 333, 333, 333, 776},         // 168
	{737, 737, 737, 737, 753, 747, 747, 760, 760, 595},         // 169
	{370, 370, 370, 370, 753, 300, 266, 276, 276, 694},         // 170
	{556, 556, 556, 556, 1042, 500, 500, 500, 500, 626},        // 171
	{584, 584, 584, 584, 987, 570, 606, 675, 564, 788},         // 172
	{333, 333, 333, 333, 603, 333, 333, 333, 333, 788},         // 173
	{737, 737, 737, 737, 987, 747, 747, 760, 760, 788},         // 174
	{333, 333, 333, 333, 603, 333, 333, 333, 333, 788},         // 175
	{400, 400, 400, 400, 400, 400, 400, 400, 400, 788},         // 176
	{584, 584, 584, 584, 549, 570, 570, 675, 564, 788},         // 177
	{333, 333, 333, 333, 411, 300, 300, 300, 300, 788},         // 178
	{333, 333, 333, 333, 549, 300, 300, 300, 300, 788},         // 179
	{333, 333, 333, 333, 549, 333, 333, 333, 333, 788},         // 180
	{556, 611, 611, 556, 713, 556, 576, 500, 500, 788},         // 181
	{537, 556, 556, 537, 494, 540, 500, 523, 453, 788},         // 182
	{278, 278, 278, 278, 460, 250, 250, 250, 250, 788},         // 183
	{333, 333, 333, 333, 549, 333, 333, 333, 333, 788},         // 184
	{333, 333, 333, 333, 549, 300, 300, 300, 300, 788},         // 185
	{365, 365, 365, 365, 549, 330, 300, 310, 310, 788},         // 186
	{556, 556, 556, 556, 549, 500, 500, 500, 500, 788},         // 187
	{834, 834, 834, 834, 1000, 750, 750, 750, 750, 788},        // 188
	{834, 834, 834, 834, 603, 750, 750, 750, 750, 788},         // 189
	{834, 834, 834, 834, 1000, 750, 750, 750, 750, 788},        // 190
	{611, 611, 611, 611, 658, 500, 500, 500, 444, 788},         // 191
	{667, 722, 722, 667, 823, 722, 667, 611, 722, 788},         // 192
	{667, 722, 722, 667, 686, 722, 667, 611, 722, 788},         // 193
	{667, 722, 722, 667, 795, 722, 667, 611, 722, 788},         // 194
	{667, 722, 722, 667, 987, 722, 667, 611, 722, 788},         // 195
	{667, 722, 722, 667, 768, 722, 667, 611, 722, 788},         // 196
	{667, 722, 722, 667, 768, 722, 667, 611, 722, 788},         // 197
	{1000, 1000, 1000, 1000, 823, 1000, 944, 889, 889, 788},    // 198
	{722, 722, 722, 722, 768, 722, 667, 667, 667, 788},         // 199
	{667, 667, 667, 667, 768, 667, 667, 611, 611, 788},         // 200
	{667, 667, 667, 667, 713, 667, 667, 611, 611, 788},         // 201
	{667, 667, 667, 667, 713, 667, 667, 611, 611, 788},         // 202
	{667, 667, 667, 667, 713, 667, 667, 611, 611, 788},         // 203
	{278, 278, 278, 278, 713, 389, 389, 333, 333, 788},         // 204
	{278, 278, 278, 278, 713, 389, 389, 333, 333, 788},         // 205
	{278, 278, 278, 278, 713, 389, 389, 333, 333, 788},         // 206
	{278, 278, 278, 278, 713, 389, 389, 333, 333, 788},         // 207
	{722, 722, 722, 722, 768, 722, 722, 722, 722, 788},         // 208
	{722, 722, 722, 722, 713, 722, 722, 667, 722, 788},         // 209
	{778, 778, 778, 778, 790, 778, 722, 722, 722, 788},         // 210
	{778, 778, 778, 778, 790, 778, 722, 722, 722, 788},         // 211
	{778, 778, 778, 778, 890, 778, 722, 722, 722, 894},         // 212
	{778, 778, 778, 778, 823, 778, 722, 722, 722, 838},         // 213
	{778, 778, 778, 778, 549, 778, 722, 722, 722, 1016},        // 214
	{584, 584, 584, 584, 250, 570, 570, 675, 564, 458},         // 215
	{778, 778, 778, 778, 713, 778, 722, 722, 722, 748},         // 216
	{722, 722, 722, 722, 603, 722, 722, 722, 722, 924},         // 217
	{722, 722, 722, 722, 603, 722, 722, 722, 722, 748},         // 218
	{722, 722, 722, 722, 1042, 722, 722, 722, 722, 918},        // 219
	{722, 722, 722, 722, 987, 722, 722, 722, 722, 927},         // 220
	{667, 667, 667, 667, 603, 722, 611, 556, 722, 928},         // 221
	{667, 667, 667, 667, 987, 611, 611, 611, 556, 928},         // 222
	{611, 611, 611, 611, 603, 556, 500, 500, 500, 834},         // 223
	{556, 556, 556, 556, 494, 500, 500, 500, 444, 873},         // 224
	{556, 556, 556, 556, 329, 500, 500, 500, 444, 828},         // 225
	{556, 556, 556, 556, 790, 500, 500, 500, 444, 924},         // 226
	{556, 556, 556, 556, 790, 500, 500, 500, 444, 924},         // 227
	{556, 556, 556, 556, 786, 500, 500, 500, 444, 917},         // 228
	{556, 556, 556, 556, 713, 500, 500, 500, 444, 930},         // 229
	{889, 889, 889, 889, 384, 722, 722, 667, 667, 931},         // 230
	{500, 556, 556, 500, 384, 444, 444, 444, 444, 463},         // 231
	{556, 556, 556, 556, 384, 444, 444, 444, 444, 883},         // 232
	{556, 556, 556, 556, 384, 444, 444, 444, 444, 836},         // 233
	{556, 556, 556, 556, 384, 444, 444, 444, 444, 836},         // 234
	{556, 556, 556, 556, 384, 444, 444, 444, 444, 867},         // 235
	{278, 278, 278, 278, 494, 278, 278, 278, 278, 867},         // 236
	{278, 278, 278, 278, 494, 278, 278, 278, 278, 696},         // 237
	{278, 278, 278, 278, 494, 278, 278, 278, 278, 696},         // 238
	{278, 278, 278, 278, 494, 278, 278, 278, 278, 874},         // 239
	{556, 611, 611, 556, 000, 500, 500, 500, 500, 000},         // 240
	{556, 611, 611, 556, 329, 556, 556, 500, 500, 874},         // 241
	{556, 611, 611, 556, 274, 500, 500, 500, 500, 760},         // 242
	{556, 611, 611, 556, 686, 500, 500, 500, 500, 946},         // 243
	{556, 611, 611, 556, 686, 500, 500, 500, 500, 771},         // 244
	{556, 611, 611, 556, 686, 500, 500, 500, 500, 865},         // 245
	{556, 611, 611, 556, 384, 500, 500, 500, 500, 771},         // 246
	{584, 584, 584, 584, 384, 570, 570, 675, 564, 888},         // 247
	{611, 611, 611, 611, 384, 500, 500, 500, 500, 967},         // 248
	{556, 611, 611, 556, 384, 556, 556, 500, 500, 888},         // 249
	{556, 611, 611, 556, 384, 556, 556, 500, 500, 831},         // 250
	{556, 611, 611, 556, 384, 556, 556, 500, 500, 873},         // 251
	{556, 611, 611, 556, 494, 556, 556, 500, 500, 927},         // 252
	{500, 556, 556, 500, 494, 500, 444, 444, 500, 970},         // 253
	{556, 611, 611, 556, 494, 556, 500, 500, 500, 918},         // 254
	{500, 556, 556, 500, 000, 500, 444, 444, 500, 000},         // 255
} //                                                               pdfFontWidths

// pdfStandardPaperSizes contains standard paper sizes in mm (width x height)
var pdfStandardPaperSizes = map[string][2]int{
	"A0": {841, 1189}, "B0": {1000, 1414}, "C0": {917, 1297}, // ISO-216
	"A1": {594, 841}, "B1": {707, 1000}, "C1": {648, 917},
	"A2": {420, 594}, "B2": {500, 707}, "C2": {458, 648},
	"A3": {297, 420}, "B3": {353, 500}, "C3": {324, 458},
	"A4": {210, 297}, "B4": {250, 353}, "C4": {229, 324},
	"A5": {148, 210}, "B5": {176, 250}, "C5": {162, 229},
	"A6": {105, 148}, "B6": {125, 176}, "C6": {114, 162},
	"A7": {74, 105}, "B7": {88, 125}, "C7": {81, 114},
	"A8": {52, 74}, "B8": {62, 88}, "C8": {57, 81},
	"A9": {37, 52}, "B9": {44, 62}, "C9": {40, 57},
	"A10": {26, 37}, "B10": {31, 44}, "C10": {28, 40},
	"LEGAL": {216, 356}, "TABLOID": {279, 432}, // US paper sizes
	"LETTER": {216, 279}, "LEDGER": {432, 279},
} //                                                       pdfStandardPaperSizes

//end
