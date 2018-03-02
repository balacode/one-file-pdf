// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-03-02 16:44:08 4021D0                              [one_file_pdf.go]
// -----------------------------------------------------------------------------

package pdf

// # Structures
//   PDF struct
//   PDFColor struct
//   PDFColorName struct
//   PDFFont struct
//   PDFImage struct
//   PDFPage struct
//
//   PDFPageSize struct
//   PDFPageSizeOf(pageSize string) PDFPageSize
//
// # Constants
//   const PDFNoPage = -1
//   var PDFBuiltInFontNames = []string
//   var PDFColorNames = []PDFColorName
//   var PDFStandardPageSizes = []PDFPageSize
//   var pdfFontWidths = [][]int
//
// # Constructor
//   NewPDF(pageSize string) PDF
//
// # Destructor
//   (pdf *PDF) Reset() *PDF
//
// # Read-Only Properties
//   (pdf *PDF) CurrentPage() int
//   (pdf *PDF) PageHeight() float64
//   (pdf *PDF) PageWidth() float64
//
// # Property Getters
//   (pdf *PDF) Color() PDFColor
//   (pdf *PDF) Compression() bool
//   (pdf *PDF) DocAuthor() string
//   (pdf *PDF) DocCreator() string
//   (pdf *PDF) DocKeywords() string
//   (pdf *PDF) DocSubject() string
//   (pdf *PDF) DocTitle() string
//   (pdf *PDF) FontName() string
//   (pdf *PDF) FontSize() float64
//   (pdf *PDF) HorizontalScaling() uint16
//   (pdf *PDF) LineWidth() float64
//   (pdf *PDF) Units() string
//   (pdf *PDF) X() float64
//   (pdf *PDF) Y() float64
//
// # Property Setters
//   (pdf *PDF) SetColor(nameOrHTMLValue string) *PDF
//   (pdf *PDF) SetColorRGB(red, green, blue int) *PDF
//   (pdf *PDF) SetCompression(compress bool) *PDF
//   (pdf *PDF) SetDocAuthor(s string) *PDF
//   (pdf *PDF) SetDocCreator(s string) *PDF
//   (pdf *PDF) SetDocKeywords(s string) *PDF
//   (pdf *PDF) SetDocSubject(s string) *PDF
//   (pdf *PDF) SetDocTitle(s string) *PDF
//   (pdf *PDF) SetFont(name string, points float64) *PDF
//   (pdf *PDF) SetFontName(name string) *PDF
//   (pdf *PDF) SetFontSize(points float64) *PDF
//   (pdf *PDF) SetHorizontalScaling(percent uint16) *PDF
//   (pdf *PDF) SetLineWidth(points float64) *PDF
//   (pdf *PDF) SetUnits(unitName string) *PDF
//   (pdf *PDF) SetX(x float64) *PDF
//   (pdf *PDF) SetXY(x, y float64) *PDF
//   (pdf *PDF) SetY(y float64) *PDF
//
// # Methods
//   (pdf *PDF) AddPage() *PDF
//   (pdf *PDF) Bytes() []byte
//   (pdf *PDF) DrawBox(x, y, width, height float64) *PDF
//   (pdf *PDF) DrawImage(
//                 x, y, height float64, fileNameOrBytes interface{},
//              ) *PDF
//   (pdf *PDF) DrawLine(x1, y1, x2, y2 float64) *PDF
//   (pdf *PDF) DrawText(text string) *PDF
//   (pdf *PDF) DrawTextAlignedToBox(
//                 x, y, width, height float64, align, text string,
//              ) *PDF
//   (pdf *PDF) DrawTextAt(x, y float64, text string) *PDF
//   (pdf *PDF) DrawTextInBox(
//                 x, y, width, height float64, align, text string,
//              ) *PDF
//   (pdf *PDF) DrawUnitGrid() *PDF
//   (pdf *PDF) FillBox(x, y, width, height float64) *PDF
//   (pdf *PDF) NextLine() *PDF
//   (pdf *PDF) SaveFile(filename string) *PDF
//   (pdf *PDF) SetColumnWidths(widths ...float64) *PDF
//   (pdf *PDF) TextWidth(text string) float64
//   (pdf *PDF) WrapTextLines(width float64, text string) []string
//
// # Functions:
//   (pdf *PDF) ToPoints(numberAndUnit string) float64
//
// # Private Methods
//   (pdf *PDF) applyLineWidth() *PDF
//   (pdf *PDF) applyNonStrokeColor() *PDF
//   (pdf *PDF) applyStrokeColor() *PDF
//   (pdf *PDF) drawTextLine(text string) *PDF
//   (pdf *PDF) drawTextBox(
//                 x, y, width, height float64,
//                 wrapText bool, align, text string,
//              ) *PDF
//   (pdf *PDF) printf(format string, args ...interface{}) *PDF
//   (pdf *PDF) setCurrentPage(pageNo int) *PDF
//   (pdf *PDF) textWidthPt1000(text string) float64
//   (pdf *PDF) warnIfNoPage() bool
//
// # Private Functions
//   (*PDF) colorEqual(a, b PDFColor) bool
//   (*PDF) escape(s string) []byte
//   (*PDF) getPointsPerUnit(unitName string) float64

import "bytes"         // standard
import "compress/zlib" // standard
import "fmt"           // standard
import "image"         // standard
import "io/ioutil"     // standard
import "strconv"       // standard
import "strings"       // standard
import _ "image/png"   // standard

// -----------------------------------------------------------------------------
// # Structures

// PDF is the main structure representing a PDF document.
type PDF struct {
	docAuthor         string
	docCreator        string
	docKeywords       string
	docSubject        string
	docTitle          string
	pages             []PDFPage
	fonts             []PDFFont
	images            []PDFImage
	pageSize          PDFPageSize
	pageNo            int
	pagePtr           *PDFPage
	columnNo          int
	color             PDFColor
	fontName          string
	fontSizePt        float64
	lineWidth         float64
	horizontalScaling uint16
	compressStreams   bool // enable stream compression?
	content           bytes.Buffer
	contentPtr        *bytes.Buffer
	//
	// extra features
	unitName      string  // name of measurement unit
	pointsPerUnit float64 // number of points per measurement unit
	pageWidthPt   float64 // page width in points
	pageHeightPt  float64 // page height in points
	columnWidths  []float64
} //                                                                         PDF

// PDFColor represents a color value.
type PDFColor struct {
	Red   uint8
	Green uint8
	Blue  uint8
} //                                                                    PDFColor

// PDFColorName represents a color name and associated color value.
type PDFColorName struct {
	name string
	val  PDFColor
} //                                                                PDFColorName

// PDFFont represents a font name and its appearance.
type PDFFont struct {
	fontID    int
	fontName  string
	isBuiltIn bool
	isBold    bool
	isItalic  bool
} //                                                                     PDFFont

// PDFImage represents an image.
type PDFImage struct {
	name      string
	width     int
	height    int
	data      []byte
	grayscale bool
} //                                                                    PDFImage

// PDFPage is an internal structure that holds details for each page.
type PDFPage struct {
	pageSize          PDFPageSize
	pageContent       bytes.Buffer
	fontIDs           []int
	imageNos          []int
	x                 float64
	y                 float64
	lineWidth         float64
	fontSizePt        float64
	fontID            int
	strokeColor       PDFColor
	nonStrokeColor    PDFColor
	horizontalScaling uint16
} //                                                                     PDFPage

// PDFPageSize represents a page size name
// and its width and height in points.
type PDFPageSize struct {
	Name     string
	WidthPt  float64
	HeightPt float64
} //                                                                 PDFPageSize

// PDFPageSizeOf returns a PDFPageSize struct based on
// the specified page size string. If the page size is
// not found, returns a zero-initialized structure.
func PDFPageSizeOf(pageSize string) PDFPageSize {
	pageSize = strings.Trim(strings.ToUpper(pageSize), " \a\b\f\n\r\t\v")
	for _, size := range PDFStandardPageSizes {
		if pageSize == size.Name {
			return size
		}
	}
	return PDFPageSize{}
} //                                                               PDFPageSizeOf

// -----------------------------------------------------------------------------
// # Constants

// PDFNoPage specifies there is no current page.
const PDFNoPage = -1

// PDFBuiltInFontNames lists font names
// available on all PDF implementations.
var PDFBuiltInFontNames = []string{
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
} //                                                         PDFBuiltInFontNames

// PDFColorNames is an array of web color names,
// derived from X11 color names.
var PDFColorNames = []PDFColorName{
	{"ALICEBLUE", PDFColor{240, 248, 255}},
	{"ANTIQUEWHITE", PDFColor{250, 235, 215}},
	{"AQUA", PDFColor{0, 255, 255}},
	{"AQUAMARINE", PDFColor{127, 255, 212}},
	{"AZURE", PDFColor{240, 255, 255}},
	{"BEIGE", PDFColor{245, 245, 220}},
	{"BISQUE", PDFColor{255, 228, 196}},
	{"BLACK", PDFColor{0, 0, 0}},
	{"BLANCHEDALMOND", PDFColor{255, 235, 205}},
	{"BLANCHEDALMOND", PDFColor{255, 255, 205}},
	{"BLUE", PDFColor{0, 0, 255}},
	{"BLUEVIOLET", PDFColor{138, 43, 226}},
	{"BROWN", PDFColor{165, 42, 42}},
	{"BURLYWOOD", PDFColor{222, 184, 135}},
	{"CADETBLUE", PDFColor{95, 158, 160}},
	{"CHARTREUSE", PDFColor{127, 255, 0}},
	{"CHOCOLATE", PDFColor{210, 105, 30}},
	{"CORAL", PDFColor{255, 127, 80}},
	{"CORNFLOWERBLUE", PDFColor{100, 149, 237}},
	{"CORNSILK", PDFColor{255, 248, 220}},
	{"CRIMSON", PDFColor{220, 20, 60}},
	{"CYAN", PDFColor{0, 255, 255}},
	{"DARKBLUE", PDFColor{0, 0, 139}},
	{"DARKCYAN", PDFColor{0, 139, 139}},
	{"DARKGOLDENROD", PDFColor{184, 134, 11}},
	{"DARKGRAY", PDFColor{169, 169, 169}},
	{"DARKGREEN", PDFColor{0, 100, 0}},
	{"DARKKHAKI", PDFColor{189, 183, 107}},
	{"DARKMAGENTA", PDFColor{139, 0, 139}},
	{"DARKOLIVEGREEN", PDFColor{85, 107, 47}},
	{"DARKORANGE", PDFColor{255, 140, 0}},
	{"DARKORCHID", PDFColor{153, 50, 204}},
	{"DARKRED", PDFColor{139, 0, 0}},
	{"DARKSALMON", PDFColor{233, 150, 122}},
	{"DARKSEAGREEN", PDFColor{143, 188, 143}},
	{"DARKSLATEBLUE", PDFColor{72, 61, 139}},
	{"DARKSLATEGRAY", PDFColor{47, 79, 79}},
	{"DARKTURQUOISE", PDFColor{0, 206, 209}},
	{"DARKVIOLET", PDFColor{148, 0, 211}},
	{"DEEPPINK", PDFColor{255, 20, 147}},
	{"DEEPSKYBLUE", PDFColor{0, 191, 255}},
	{"DIMGRAY", PDFColor{105, 105, 105}},
	{"DODGERBLUE", PDFColor{30, 144, 255}},
	{"FIREBRICK", PDFColor{178, 34, 34}},
	{"FLORALWHITE", PDFColor{255, 250, 240}},
	{"FORESTGREEN", PDFColor{34, 139, 34}},
	{"FUCHSIA", PDFColor{255, 0, 255}},
	{"GAINSBORO", PDFColor{220, 220, 220}},
	{"GHOSTWHITE", PDFColor{248, 248, 255}},
	{"GOLD", PDFColor{255, 215, 0}},
	{"GOLDENROD", PDFColor{218, 165, 32}},
	{"GRAY", PDFColor{127, 127, 127}},
	{"GRAY", PDFColor{128, 128, 128}},
	{"GREEN", PDFColor{0, 128, 0}},
	{"GREENYELLOW", PDFColor{173, 255, 47}},
	{"HONEYDEW", PDFColor{240, 255, 240}},
	{"HOTPINK", PDFColor{255, 105, 180}},
	{"INDIANRED", PDFColor{205, 92, 92}},
	{"INDIGO", PDFColor{75, 0, 130}},
	{"IVORY", PDFColor{255, 255, 240}},
	{"KHAKI", PDFColor{195, 176, 145}},
	{"KHAKI", PDFColor{240, 230, 140}},
	{"LAVENDER", PDFColor{230, 230, 250}},
	{"LAVENDERBLUSH", PDFColor{255, 240, 245}},
	{"LAWNGREEN", PDFColor{124, 252, 0}},
	{"LEMONCHIFFON", PDFColor{255, 250, 205}},
	{"LIGHTBLUE", PDFColor{173, 216, 230}},
	{"LIGHTCORAL", PDFColor{240, 128, 128}},
	{"LIGHTCYAN", PDFColor{224, 255, 255}},
	{"LIGHTGOLDENRODYELLOW", PDFColor{250, 250, 210}},
	{"LIGHTGRAY", PDFColor{211, 211, 211}},
	{"LIGHTGREEN", PDFColor{144, 238, 144}},
	{"LIGHTPINK", PDFColor{255, 182, 193}},
	{"LIGHTSALMON", PDFColor{255, 160, 122}},
	{"LIGHTSEAGREEN", PDFColor{32, 178, 170}},
	{"LIGHTSKYBLUE", PDFColor{135, 206, 250}},
	{"LIGHTSLATEGRAY", PDFColor{119, 136, 153}},
	{"LIGHTSTEELBLUE", PDFColor{176, 196, 222}},
	{"LIGHTYELLOW", PDFColor{255, 255, 224}},
	{"LIME", PDFColor{0, 255, 0}},
	{"LIMEGREEN", PDFColor{50, 205, 50}},
	{"LINEN", PDFColor{250, 240, 230}},
	{"MAGENTA", PDFColor{255, 0, 255}},
	{"MAROON", PDFColor{128, 0, 0}},
	{"MEDIUMAQUAMARINE", PDFColor{102, 205, 170}},
	{"MEDIUMBLUE", PDFColor{0, 0, 205}},
	{"MEDIUMORCHID", PDFColor{186, 85, 211}},
	{"MEDIUMPURPLE", PDFColor{147, 112, 219}},
	{"MEDIUMSEAGREEN", PDFColor{60, 179, 113}},
	{"MEDIUMSLATEBLUE", PDFColor{123, 104, 238}},
	{"MEDIUMSPRINGGREEN", PDFColor{0, 250, 154}},
	{"MEDIUMTURQUOISE", PDFColor{72, 209, 204}},
	{"MEDIUMVIOLETRED", PDFColor{199, 21, 133}},
	{"MIDNIGHTBLUE", PDFColor{25, 25, 112}},
	{"MINTCREAM", PDFColor{245, 255, 250}},
	{"MISTYROSE", PDFColor{255, 228, 225}},
	{"MOCCASIN", PDFColor{255, 228, 181}},
	{"NAVAJOWHITE", PDFColor{255, 222, 173}},
	{"NAVY", PDFColor{0, 0, 128}},
	{"OLDLACE", PDFColor{253, 245, 230}},
	{"OLIVE", PDFColor{128, 128, 0}},
	{"OLIVEDRAB", PDFColor{107, 142, 35}},
	{"ORANGE", PDFColor{255, 165, 0}},
	{"ORANGERED", PDFColor{255, 69, 0}},
	{"ORCHID", PDFColor{218, 112, 214}},
	{"PALEGOLDENROD", PDFColor{238, 232, 170}},
	{"PALEGREEN", PDFColor{152, 251, 152}},
	{"PALETURQUOISE", PDFColor{175, 238, 238}},
	{"PALEVIOLETRED", PDFColor{219, 112, 147}},
	{"PAPAYAWHIP", PDFColor{255, 239, 213}},
	{"PEACHPUFF", PDFColor{255, 218, 185}},
	{"PERU", PDFColor{205, 133, 63}},
	{"PINK", PDFColor{255, 192, 203}},
	{"PLUM", PDFColor{221, 160, 221}},
	{"POWDERBLUE", PDFColor{176, 224, 230}},
	{"PURPLE", PDFColor{128, 0, 128}},
	{"RED", PDFColor{255, 0, 0}},
	{"ROSYBROWN", PDFColor{188, 143, 143}},
	{"ROYALBLUE", PDFColor{65, 105, 225}},
	{"SADDLEBROWN", PDFColor{139, 69, 19}},
	{"SALMON", PDFColor{250, 128, 114}},
	{"SANDYBROWN", PDFColor{244, 164, 96}},
	{"SEAGREEN", PDFColor{46, 139, 87}},
	{"SEASHELL", PDFColor{255, 245, 238}},
	{"SIENNA", PDFColor{160, 82, 45}},
	{"SILVER", PDFColor{192, 192, 192}},
	{"SKYBLUE", PDFColor{135, 206, 235}},
	{"SLATEBLUE", PDFColor{106, 90, 205}},
	{"SLATEGRAY", PDFColor{112, 128, 144}},
	{"SNOW", PDFColor{255, 250, 250}},
	{"SPRINGGREEN", PDFColor{0, 255, 127}},
	{"STEELBLUE", PDFColor{70, 130, 180}},
	{"TAN", PDFColor{210, 180, 140}},
	{"TEAL", PDFColor{0, 128, 128}},
	{"THISTLE", PDFColor{216, 191, 216}},
	{"TOMATO", PDFColor{255, 99, 71}},
	{"TURQUOISE", PDFColor{64, 224, 208}},
	{"VIOLET", PDFColor{238, 130, 238}},
	{"WHEAT", PDFColor{245, 222, 179}},
	{"WHITE", PDFColor{255, 255, 255}},
	{"WHITESMOKE", PDFColor{245, 245, 245}},
	{"YELLOW", PDFColor{255, 255, 0}},
	{"YELLOWGREEN", PDFColor{139, 205, 50}},
	{"YELLOWGREEN", PDFColor{154, 205, 50}},
} //                                                               PDFColorNames

// PDFStandardPageSizes is an array of standard page sizes,
// specifying the size name, width and height in points.
var PDFStandardPageSizes = []PDFPageSize{
	{"A3", 841.89, 1190.55},
	{"A3-L", 1190.55, 841.89}, // A3-Landscape, etc.
	{"A4", 595.28, 841.89},
	{"A4-L", 841.89, 595.28},
	{"A5", 420.94, 595.28},
	{"A5-L", 595.28, 420.94},
	{"LEGAL", 612, 1008},
	{"LEGAL-L", 1008, 612},
	{"LETTER", 612, 792},
	{"LETTER-L", 792, 612},
} //                                                        PDFStandardPageSizes

// Built-in Font Widths:
// 0 Helvetica
// 1 HelveticaBold
// 2 HelveticaBoldOblique
// 3 HelveticaOblique
// 4 Symbol
// 5 TimesBold
// 6 TimesBoldItalic
// 7 TimesItalic
// 8 TimesRoman
// 9 ZapfDingbats
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
	{722, 722, 722, 722, 612, 722, 722, 722, 722, 788},         // 068 d
	{667, 667, 667, 667, 611, 667, 667, 611, 611, 790},         // 069 E
	{611, 611, 611, 611, 763, 611, 667, 611, 556, 793},         // 070 F
	{778, 778, 778, 778, 603, 778, 722, 722, 722, 794},         // 071 G
	{722, 722, 722, 722, 722, 778, 778, 722, 722, 816},         // 072 h
	{278, 278, 278, 278, 333, 389, 389, 333, 333, 823},         // 073 i
	{500, 556, 556, 500, 631, 500, 500, 444, 389, 789},         // 074 J
	{667, 722, 722, 667, 722, 778, 667, 667, 722, 841},         // 075 K
	{556, 611, 611, 556, 686, 667, 611, 556, 611, 823},         // 076 L
	{833, 833, 833, 833, 889, 944, 889, 833, 889, 833},         // 077 M
	{722, 722, 722, 722, 722, 722, 722, 667, 722, 816},         // 078 N
	{778, 778, 778, 778, 722, 778, 722, 722, 722, 831},         // 079 O
	{667, 667, 667, 667, 768, 611, 611, 611, 556, 923},         // 080 p
	{778, 778, 778, 778, 741, 778, 722, 722, 722, 744},         // 081 Q
	{722, 722, 722, 722, 556, 722, 667, 611, 667, 723},         // 082 R
	{667, 667, 667, 667, 592, 556, 556, 500, 556, 749},         // 083 S
	{611, 611, 611, 611, 611, 667, 611, 556, 611, 790},         // 084 T
	{722, 722, 722, 722, 690, 722, 722, 722, 722, 792},         // 085 u
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

// -----------------------------------------------------------------------------
// # Constructor

// NewPDF creates and initializes a new PDF object.
func NewPDF(pageSize string) PDF {
	var pdf PDF //                                        create a new PDF object
	pdf.compressStreams = true
	pdf.horizontalScaling = 100
	pdf.pageNo = PDFNoPage
	//
	// get and store dimensions of specified page size (in points)
	var size = PDFPageSizeOf(pageSize)
	if size.Name == "" {
		pdf.err("Unknown page_size ^", pageSize, ". Defaulting to 'A4'.")
		size = PDFPageSizeOf("A4")
	}
	pdf.pageSize = size
	//
	// store page dimensions
	pdf.pageWidthPt = pdf.pageSize.WidthPt
	pdf.pageHeightPt = pdf.pageSize.HeightPt
	//
	// set default units, otherwise pointsPerUnit and all x and y will be 0
	pdf.SetUnits("point")
	return pdf
} //                                                                      NewPDF

// -----------------------------------------------------------------------------
// # Destructor

// Reset releases all resources and resets all variables.
func (pdf *PDF) Reset() *PDF {
	for _, page := range pdf.pages {
		if len(page.fontIDs) > 0 {
			page.fontIDs = []int{}
		}
		if len(page.imageNos) > 0 {
			page.imageNos = []int{}
		}
		if page.pageContent.Len() > 0 {
			page.pageContent.Reset()
		}
	}
	if len(pdf.pages) > 0 {
		pdf.pages = []PDFPage{}
	}
	if len(pdf.fonts) > 0 {
		pdf.fonts = []PDFFont{}
	}
	if len(pdf.images) > 0 {
		pdf.images = []PDFImage{}
	}
	pdf.content.Reset()
	pdf.docAuthor, pdf.docCreator, pdf.docKeywords, pdf.docSubject,
		pdf.docTitle, pdf.unitName = "", "", "", "", "", ""
	if len(pdf.columnWidths) > 0 {
		pdf.columnWidths = []float64{}
	}
	return pdf
} //                                                                       Reset

// -----------------------------------------------------------------------------
// # Read-Only Properties

// CurrentPage returns the number of the current page,
// 1 being the first page.
func (pdf *PDF) CurrentPage() int {
	return pdf.pageNo + 1
} //                                                                 CurrentPage

// PageHeight returns the height of the current page in selected units.
func (pdf *PDF) PageHeight() float64 {
	if pdf.pageNo == PDFNoPage || pdf.pageNo > (len(pdf.pages)-1) ||
		pdf.pagePtr == nil {
		return pdf.pageSize.HeightPt / pdf.pointsPerUnit
	}
	return pdf.pagePtr.pageSize.HeightPt / pdf.pointsPerUnit
} //                                                                  PageHeight

// PageWidth returns the width of the current page in selected units.
func (pdf *PDF) PageWidth() float64 {
	if pdf.pageNo == PDFNoPage || pdf.pageNo > (len(pdf.pages)-1) ||
		pdf.pagePtr == nil {
		return pdf.pageSize.WidthPt / pdf.pointsPerUnit
	}
	return pdf.pagePtr.pageSize.WidthPt / pdf.pointsPerUnit
} //                                                                   PageWidth

// -----------------------------------------------------------------------------
// # Property Getters

// Color returns the current color, which is used for text, lines and fills.
func (pdf *PDF) Color() PDFColor {
	return pdf.color
} //                                                                       Color

// Compression returns the current compression mode.
// If it is true, all PDF content will be compressed when
// the PDF is generated. If false, most PDF content
// (excluding images) will be in plain text, which is
// useful for debugging or to study PDF commands.
func (pdf *PDF) Compression() bool {
	return pdf.compressStreams
} //                                                                 Compression

// DocAuthor returns the optional 'document author' metadata entry.
func (pdf *PDF) DocAuthor() string {
	return pdf.docAuthor
} //                                                                   DocAuthor

// DocCreator returns the optional 'document creator' metadata entry.
func (pdf *PDF) DocCreator() string {
	return pdf.docCreator
} //                                                                  DocCreator

// DocKeywords returns the optional 'document keywords' metadata entry.
func (pdf *PDF) DocKeywords() string {
	return pdf.docKeywords
} //                                                                 DocKeywords

// DocSubject returns the optional 'document subject' metadata entry.
func (pdf *PDF) DocSubject() string {
	return pdf.docSubject
} //                                                                  DocSubject

// DocTitle returns the optional 'document subject' metadata entry.
func (pdf *PDF) DocTitle() string {
	return pdf.docTitle
} //                                                                    DocTitle

// FontName returns the name of the currently-active typeface.
func (pdf *PDF) FontName() string {
	return pdf.fontName
} //                                                                    FontName

// FontSize returns the current font size in points.
func (pdf *PDF) FontSize() float64 {
	return pdf.fontSizePt
} //                                                                    FontSize

// HorizontalScaling returns the current horizontal scaling in percent.
func (pdf *PDF) HorizontalScaling() uint16 {
	return pdf.horizontalScaling
} //                                                           HorizontalScaling

// LineWidth returns the current line width in points.
func (pdf *PDF) LineWidth() float64 {
	return pdf.lineWidth
} //                                                                   LineWidth

// Units returns the currently selected measurement units.
// Can be upper/lowercase:
// MM CM " in inch inches tw twip twips pt point points.
func (pdf *PDF) Units() string {
	return pdf.unitName
} //                                                                       Units

// X returns the X-coordinate of the current drawing position.
func (pdf *PDF) X() float64 {
	if pdf.warnIfNoPage() {
		return 0
	}
	return pdf.pagePtr.x / pdf.pointsPerUnit
} //                                                                           X

// Y returns the Y-coordinate of the current drawing position.
func (pdf *PDF) Y() float64 {
	if pdf.warnIfNoPage() {
		return 0
	}
	return (pdf.pageHeightPt - pdf.pagePtr.y) / pdf.pointsPerUnit
} //                                                                           Y

// -----------------------------------------------------------------------------
// # Property Setters

// SetColor sets the current color using a web/X11 color name
// (e.g. "HoneyDew") or HTML color value such as "#191970"
// for midnight blue (#RRGGBB). The current color is used
// for subsequent text and line drawing and fills.
// If the name is unknown or valid, sets the current color to black.
func (pdf *PDF) SetColor(nameOrHTMLValue string) *PDF {
	nameOrHTMLValue = strings.ToUpper(nameOrHTMLValue)
	var color = (func() PDFColor {
		//
		// if name starts with '#' treat it as HTML color (#RRGGBB)
		if nameOrHTMLValue[0] == '#' {
			var hex [6]uint8
			for i, r := range nameOrHTMLValue[1:] {
				if i > 6 {
					break
				}
				if (r >= '0' && r <= '9') || (r >= 'a' && r <= 'f') {
					if r >= '0' && r <= '9' {
						hex[i] = uint8(r - '0')
					} else if r >= 'A' && r <= 'F' {
						hex[i] = uint8(r - 'A' + 10)
					} else {
						hex[i] = 255
					}
				} else {
					return PDFColor{} // black - conversion failed
				}
			}
			return PDFColor{
				Red:   hex[0]*16 + hex[1],
				Green: hex[2]*16 + hex[3],
				Blue:  hex[4]*16 + hex[5],
			}
		}
		for _, color := range PDFColorNames { // otherwise search for color name
			if nameOrHTMLValue == color.name {
				return color.val
			}
		}
		return PDFColor{} // black - color name not found
	})() //                                                                 IIFE
	pdf.SetColorRGB(int(color.Red), int(color.Green), int(color.Blue))
	return pdf
} //                                                                    SetColor

// SetColorRGB sets the current color using separate
// red, green and blue values. Each component value
// can range from 0 to 255. The current color is used
// for subsequent text and line drawing and fills.
func (pdf *PDF) SetColorRGB(red, green, blue int) *PDF {
	pdf.color = PDFColor{Red: uint8(red), Green: uint8(green), Blue: uint8(blue)}
	return pdf
} //                                                                 SetColorRGB

// SetCompression sets the compression mode used to generate
// the PDF. If set to true, all PDF content will be
// compressed when the PDF is generated. If false, most PDF
// content (excluding images) will be in plain text, which
// is useful for debugging or to study PDF commands.
func (pdf *PDF) SetCompression(compress bool) *PDF {
	pdf.compressStreams = compress
	return pdf
} //                                                              SetCompression

// SetDocAuthor sets the optional 'document author' metadata entry.
func (pdf *PDF) SetDocAuthor(s string) *PDF {
	pdf.docAuthor = s
	return pdf
} //                                                                SetDocAuthor

// SetDocCreator sets the optional 'document creator' metadata entry.
func (pdf *PDF) SetDocCreator(s string) *PDF {
	pdf.docCreator = s
	return pdf
} //                                                               SetDocCreator

// SetDocKeywords sets the optional 'document keywords' metadata entry.
func (pdf *PDF) SetDocKeywords(s string) *PDF {
	pdf.docKeywords = s
	return pdf
} //                                                              SetDocKeywords

// SetDocSubject sets the optional 'document subject' metadata entry.
func (pdf *PDF) SetDocSubject(s string) *PDF {
	pdf.docSubject = s
	return pdf
} //                                                               SetDocSubject

// SetDocTitle sets the optional 'document title' metadata entry.
func (pdf *PDF) SetDocTitle(s string) *PDF {
	pdf.docTitle = s
	return pdf
} //                                                                 SetDocTitle

// SetFont changes the current font name and size
// in points. For the font name, use one of the
// standard font names, such as 'Helvetica'. This
// font will be used for subsequent text drawing.
func (pdf *PDF) SetFont(name string, points float64) *PDF {
	pdf.SetFontName(name).SetFontSize(points)
	return pdf
} //                                                                     SetFont

// SetFontName changes the current font, while using the
// same font size as the previous font. Use one of the
// standard font names, such as 'Helvetica'.
func (pdf *PDF) SetFontName(name string) *PDF {
	pdf.fontName = name
	return pdf
} //                                                                 SetFontName

// SetFontSize changes the current font size in points,
// without changing the currently-selected font typeface.
func (pdf *PDF) SetFontSize(points float64) *PDF {
	pdf.fontSizePt = points
	return pdf
} //                                                                 SetFontSize

// SetHorizontalScaling changes the horizontal scaling in percent.
// For example, 200 will stretch text to double its normal width.
func (pdf *PDF) SetHorizontalScaling(percent uint16) *PDF {
	pdf.horizontalScaling = percent
	return pdf
} //                                                        SetHorizontalScaling

// SetLineWidth changes the line width in points.
func (pdf *PDF) SetLineWidth(points float64) *PDF {
	pdf.lineWidth = points
	return pdf
} //                                                                SetLineWidth

// SetUnits changes the current measurement units.
// Can be upper/lowercase: MM CM " IN INCH INCHES TW TWIP TWIPS PT POINT POINTS.
func (pdf *PDF) SetUnits(unitName string) *PDF {
	pdf.unitName = strings.ToUpper(strings.Trim(unitName, " \a\b\f\n\r\t\v"))
	pdf.pointsPerUnit = pdf.getPointsPerUnit(pdf.unitName)
	return pdf
} //                                                                    SetUnits

// SetX changes the X-coordinate of the current drawing position.
func (pdf *PDF) SetX(x float64) *PDF {
	if !pdf.warnIfNoPage() {
		pdf.pagePtr.x = x * pdf.pointsPerUnit
	}
	return pdf
} //                                                                        SetX

// SetXY changes both X- and Y-coordinates of the current drawing position.
func (pdf *PDF) SetXY(x, y float64) *PDF {
	if pdf.warnIfNoPage() {
		return pdf
	}
	x, y = x*pdf.pointsPerUnit, pdf.pageHeightPt-y*pdf.pointsPerUnit
	pdf.pagePtr.x, pdf.pagePtr.y = x, y
	return pdf
} //                                                                       SetXY

// SetY changes the Y-coordinate of the current drawing position.
func (pdf *PDF) SetY(y float64) *PDF {
	if pdf.warnIfNoPage() {
		pdf.pagePtr.y = (pdf.pageHeightPt * pdf.pointsPerUnit) -
			(y * pdf.pointsPerUnit)
	}
	return pdf
} //                                                                        SetY

// -----------------------------------------------------------------------------
// # Methods

// AddPage appends a new blank page to the document and makes it selected.
func (pdf *PDF) AddPage() *PDF {
	var i = len(pdf.pages)
	pdf.pages = append(pdf.pages, PDFPage{})
	var pg = &pdf.pages[i]
	pdf.setCurrentPage(i)
	pdf.pageNo = i
	pg.pageSize = pdf.pageSize
	pg.horizontalScaling = 100
	pg.pageContent.Reset()
	pg.x, pg.y = -1, -1 // these variables must default to -1
	pg.strokeColor, pg.nonStrokeColor = PDFColor{1, 1, 1}, PDFColor{1, 1, 1}
	pdf.SetXY(0, 0)
	return pdf
} //                                                                     AddPage

// Bytes generates the PDF document from various page and
// auxiliary objects and returns it in an array of bytes,
// identical to the content of a PDF file. This method is where
// you'll find the core structure of a PDF document.
func (pdf *PDF) Bytes() []byte {
	// called by: SaveFile()
	var objectOffsets []int
	var objectNo int
	//
	// increases the object serial number and stores its offset in the array
	var nextObject = func() int {
		objectNo++
		for len(objectOffsets) <= objectNo {
			objectOffsets = append(objectOffsets, pdf.content.Len())
		}
		return objectNo
	}
	// outputs an object header
	var obj = func(objectType string) {
		pdf.setCurrentPage(PDFNoPage)
		var objectNo = nextObject()
		if objectType == "" {
			pdf.printf("%d 0 obj<<", objectNo)
		} else if objectType[0] == '/' {
			pdf.printf("%d 0 obj<</Type%s", objectNo, objectType)
		} else {
			pdf.err("objectType should begin with '/' or be a blank string.")
		}
	}
	// (outputs a stream object to the document's main buffer)
	var streamObject = func(content []byte) {
		pdf.setCurrentPage(PDFNoPage).printf("%d 0 obj <<", nextObject())
		pdf.streamData(content)
	}
	var endobj = func() {
		pdf.printf(">>\nendobj\n")
	}
	// free any existing generated content and write beginning of document
	var pageObjectsIndex = 3
	var fontObjectsIndex = pageObjectsIndex + len(pdf.pages)*2
	var imageObjectsIndex = fontObjectsIndex + len(pdf.fonts)
	var infoObjectIndex = imageObjectsIndex + len(pdf.images)
	pdf.setCurrentPage(PDFNoPage)
	pdf.content.Reset()
	pdf.printf("%%PDF-1.4\n")
	obj("/Catalog")
	pdf.printf("/Pages 2 0 R")
	endobj()
	//
	//  write /Pages object, page count and page size
	obj("/Pages") // 2 0 obj
	pdf.printf("/Count %d/MediaBox[0 0 %d %d]", len(pdf.pages),
		int(pdf.pageSize.WidthPt), int(pdf.pageSize.HeightPt))
	//
	// page numbers
	if len(pdf.pages) > 0 {
		var pageObjectNo = pageObjectsIndex
		pdf.printf("/Kids[")
		for i := range pdf.pages {
			if i > 0 {
				pdf.printf(" ")
			}
			pdf.printf("%d 0 R", pageObjectNo)
			pageObjectNo += 2 // 1 for page, 1 for stream
		}
		pdf.printf("]")
	}
	endobj()
	//
	// write each page
	for _, pg := range pdf.pages {
		if pg.pageContent.Len() == 0 {
			fmt.Printf("WARNING: Empty Page\n")
			continue
		}
		obj("/Page")
		pdf.printf("/Parent 2 0 R/Contents %d 0 R", objectNo+1)
		if len(pg.fontIDs) > 0 || len(pg.imageNos) > 0 {
			pdf.printf("/Resources<<")
		}
		if len(pg.fontIDs) > 0 {
			pdf.printf("/Font <<")
			for fontNo := range pdf.fonts {
				pdf.printf("/F%d %d 0 R", fontNo+1, fontObjectsIndex+fontNo)
			}
			pdf.printf(">>")
		}
		if len(pg.imageNos) > 0 {
			pdf.printf("/XObject<<")
			for imageNo := range pg.imageNos {
				pdf.printf("/Im%d %d 0 R", imageNo, imageObjectsIndex+imageNo)
			}
			pdf.printf(">>")
		}
		if len(pg.fontIDs) > 0 || len(pg.imageNos) > 0 {
			pdf.printf(">>")
		}
		endobj()
		streamObject([]byte(pg.pageContent.String()))
	}
	// write fonts
	for _, iter := range pdf.fonts {
		obj("/Font")
		if iter.isBuiltIn {
			pdf.printf("/Subtype/Type1/Name/F%d"+
				"/BaseFont/%s/Encoding/WinAnsiEncoding",
				iter.fontID, iter.fontName)
		}
		endobj()
	}
	// write images
	for _, iter := range pdf.images {
		obj("/XObject")
		pdf.printf(
			"/Subtype/Image/Width %d/Height %d"+
				"/ColorSpace/DeviceGray/BitsPerComponent 8",
			iter.width, iter.height,
		)
		pdf.streamData(iter.data)
		pdf.printf("\nendobj\n")
	}
	// write info object
	if pdf.docTitle != "" || pdf.docSubject != "" ||
		pdf.docKeywords != "" || pdf.docAuthor != "" || pdf.docCreator != "" {
		for _, iter := range []struct {
			label string
			field string
		}{
			{"/Title ", pdf.docTitle},
			{"/Subject ", pdf.docSubject},
			{"/Keywords ", pdf.docKeywords},
			{"/Author ", pdf.docAuthor},
			{"/Creator ", pdf.docCreator},
		} {
			if iter.field != "" {
				pdf.printf(iter.label).printf("(").
					printf(string(pdf.escape(iter.field))).printf(")")
			}
		}
		endobj() // info
	}
	// write cross-reference table at end of document
	var startXref = pdf.content.Len()
	pdf.setCurrentPage(PDFNoPage).
		printf("xref\n0 %d\n0000000000 65535 f \n",
			len(objectOffsets))
	for _, offset := range objectOffsets[1:] {
		pdf.printf("%010d 00000 n \n", offset)
	}
	// write the trailer
	pdf.printf("trailer\n<</Size %d/Root 1 0 R", len(objectOffsets))
	if infoObjectIndex > 0 {
		pdf.printf("/Info %d 0 R", infoObjectIndex)
	}
	pdf.printf(">>\nstartxref\n%d\n", startXref).printf("%%%%EOF\n")
	return pdf.content.Bytes()
} //                                                                       Bytes

// DrawBox draws a rectangle.
func (pdf *PDF) DrawBox(x, y, width, height float64) *PDF {
	if pdf.warnIfNoPage() {
		return pdf
	}
	x, y, width, height = x*pdf.pointsPerUnit, y*pdf.pointsPerUnit,
		width*pdf.pointsPerUnit, height*pdf.pointsPerUnit
	y = pdf.pageHeightPt - y - height
	pdf.applyLineWidth().applyStrokeColor()
	// re = construct a rectangular path, S = stroke path
	pdf.printf("%.3f %.3f %.3f %.3f re S\n", x, y, width, height)
	return pdf
} //                                                                     DrawBox

// DrawImage draws an image
func (pdf *PDF) DrawImage(
	x, y, height float64,
	fileNameOrBytes interface{},
) *PDF {
	//
	if pdf.warnIfNoPage() {
		return pdf
	}
	var imageName string
	var imageBuf *bytes.Buffer
	switch val := fileNameOrBytes.(type) {
	case string:
		var data, err = ioutil.ReadFile(val)
		if err != nil {
			pdf.err("File", val, ":", err)
			return pdf
		}
		imageName = val
		imageBuf = bytes.NewBuffer(data)
	case []byte:
		var l = len(val)
		if l > 32 {
			l = 32
		}
		imageName = fmt.Sprintf("%x", val[:l])
		imageBuf = bytes.NewBuffer(val)
	}
	var pg = pdf.pagePtr
	var imgNo = -1
	var img PDFImage
	for i, iter := range pdf.images {
		if iter.name == imageName {
			imgNo, img = i, iter
		}
	}
	if imgNo == -1 {
		var decoded, _, err = image.Decode(imageBuf)
		if err != nil {
			fmt.Printf("image not decoded: %v\n", err)
			return pdf
		}
		var bounds = decoded.Bounds()
		var w, h = bounds.Max.X, bounds.Max.Y
		var data []byte
		for y := 0; y < h; y++ {
			for x := 0; x < w; x++ {
				var r, g, b, _ = decoded.At(x, y).RGBA()
				data = append(data, byte(
					float64(r)*0.2126+float64(g)*0.7152+float64(b)*0.0722,
				))
			}
		}
		img = PDFImage{
			name:      imageName,
			width:     w,
			height:    h,
			data:      data,
			grayscale: true, //` determine grayscale
		}
		imgNo = len(pdf.images)
		pdf.images = append(pdf.images, img)
	}
	// add the image number to the current page, if not already referenced
	var found bool
	for _, iter := range pg.imageNos {
		if iter == imgNo {
			found = true
			break
		}
	}
	if !found {
		pg.imageNos = append(pg.imageNos, imgNo)
	}
	x, y, height = x*pdf.pointsPerUnit, y*pdf.pointsPerUnit,
		height*pdf.pointsPerUnit
	y = pdf.pageHeightPt - y - height
	var width = float64(img.width) / float64(img.height) * height
	//
	// write command to draw the image
	pdf.printf("q\n"+ // save graphics state
		// w      h  x  y
		" %f 0 0 %f %f %f cm\n"+
		"/Im%d Do\n"+
		"Q\n", // restore graphics state
		width, height, x, y, imgNo)
	return pdf
} //                                                                   DrawImage

// DrawLine draws a straight line from point (x1, y1) to point (x2, y2).
func (pdf *PDF) DrawLine(x1, y1, x2, y2 float64) *PDF {
	if pdf.warnIfNoPage() {
		return pdf
	}
	x1, y1, x2, y2 = x1*pdf.pointsPerUnit, y1*pdf.pointsPerUnit,
		x2*pdf.pointsPerUnit, y2*pdf.pointsPerUnit
	y1, y2 = pdf.pageHeightPt-y1, pdf.pageHeightPt-y2
	pdf.applyLineWidth().applyStrokeColor()
	//
	// send command to draw the line
	pdf.printf("%.3f %.3f m %.3f %.3f l S\n", x1, y1, x2, y2)
	return pdf
} //                                                                    DrawLine

// DrawText draws a text string at the current position (X, Y).
func (pdf *PDF) DrawText(text string) *PDF {
	if pdf.warnIfNoPage() {
		return pdf
	}
	if len(pdf.columnWidths) == 0 {
		pdf.drawTextLine(text)
		return pdf
	}
	var x = 0.0
	for i := 0; i < pdf.columnNo; i++ {
		x += pdf.columnWidths[i]
	}
	pdf.SetX(x)
	pdf.drawTextLine(text)
	if pdf.columnNo == (len(pdf.columnWidths) - 1) {
		pdf.NextLine()
	} else {
		pdf.columnNo++
	}
	return pdf
} //                                                                    DrawText

// DrawTextAlignedToBox draws 'text' within a rectangle specified
// by 'x', 'y', 'width' and 'height'. If 'align' is blank, the
// text is center-aligned both vertically and horizontally.
// Specify 'L' or 'R' to align the text left or right, and 'T' or
// 'B' to align the text to the top or bottom of the box.
func (pdf *PDF) DrawTextAlignedToBox(
	x, y, width, height float64, align, text string,
) *PDF {
	if pdf.pageNo < 0 {
		pdf.err("No current page.")
		return pdf
	}
	pdf.drawTextBox(x, y, width, height, false, align, text)
	return pdf
} //                                                        DrawTextAlignedToBox

// DrawTextAt draws text at the specified point (x, y).
func (pdf *PDF) DrawTextAt(x, y float64, text string) *PDF {
	return pdf.SetXY(x, y).DrawText(text)
} //                                                                  DrawTextAt

// DrawTextInBox draws word-wrapped text within a rectangle
// specified by 'x', 'y', 'width' and 'height'. If 'align' is blank,
// the text is center-aligned both vertically and horizontally.
// Specify 'L' or 'R' to align the text left or right, and 'T' or
// 'B' to align the text to the top or bottom of the box.
func (pdf *PDF) DrawTextInBox(
	x, y, width, height float64, align, text string,
) *PDF {
	if pdf.pageNo < 0 {
		pdf.err("No current page.")
		return pdf
	}
	pdf.drawTextBox(x, y, width, height, true, align, text)
	return pdf
} //                                                               DrawTextInBox

// DrawUnitGrid draws a light-gray grid demarcated in the current
// measurement unit. The grid fills the entire page.
// It helps with item positioning.
func (pdf *PDF) DrawUnitGrid() *PDF {
	var x, y, pageWidth, pageHeight = 0.0, 0.0, pdf.PageWidth(), pdf.PageHeight()
	if pdf.pageNo < 0 { // ensure there is a current page
		pdf.err("No current page.")
		return pdf
	}
	pdf.SetLineWidth(0.1).SetFont("Helvetica", 8)
	var xh = 0.1
	var yh = 0.5
	var xv = 0.3
	var yv = 0.3
	var i = 0
	// draw vertical lines
	for x = 0; x < pageWidth; x++ {
		pdf.SetColorRGB(200, 200, 200).DrawLine(x, 0, x, pageHeight)
		if i > 0 {
			pdf.SetColor("Indigo").SetXY(x+xh, y+yh).DrawText(strconv.Itoa(i))
		}
		i++
	}
	// draw horizontal lines
	i = 0
	for y = 0; y < pageHeight; y++ {
		pdf.SetColorRGB(200, 200, 200).DrawLine(0, y, pageWidth, y)
		if i > 0 {
			pdf.SetColor("Indigo").SetXY(xv, y+yv).DrawText(strconv.Itoa(i))
		}
		i++
	}
	return pdf
} //                                                                DrawUnitGrid

// FillBox fills a rectangle with the current color.
func (pdf *PDF) FillBox(x, y, width, height float64) *PDF {
	if pdf.warnIfNoPage() {
		return pdf
	}
	x, y, width, height = x*pdf.pointsPerUnit, y*pdf.pointsPerUnit,
		width*pdf.pointsPerUnit, height*pdf.pointsPerUnit
	y = pdf.pageHeightPt - y - height
	pdf.applyLineWidth().applyNonStrokeColor().
		printf("%.3f %.3f %.3f %.3f re f\n", x, y, width, height)
	// 're' = construct a rectangular path, f = fill
	return pdf
} //                                                                     FillBox

// NextLine advances the text writing position to the next line.
// I.e. the Y increases by the height of the font and
// the X-coordinate is reset to zero.
func (pdf *PDF) NextLine() *PDF {
	var x = pdf.X()
	var y = pdf.Y()
	var lineHeight = pdf.FontSize() * pdf.pointsPerUnit
	var pageHeight = pdf.pageHeightPt * pdf.pointsPerUnit
	y += lineHeight
	if y > pageHeight {
		pdf.AddPage()
		y = 0
	}
	pdf.columnNo = 0
	if len(pdf.columnWidths) == 0 {
		x = 0
	} else {
		x = pdf.columnWidths[0]
	}
	pdf.SetXY(x, y)
	return pdf
} //                                                                    NextLine

// SaveFile generates and saves the PDF document to a file.
func (pdf *PDF) SaveFile(filename string) *PDF {
	var err = ioutil.WriteFile(filename, pdf.Bytes(), 0)
	if err != nil {
		pdf.err("Failed writing to file", filename, ":", err)
	}
	return pdf
} //                                                                    SaveFile

// SetColumnWidths creates columns along the X-axis.
func (pdf *PDF) SetColumnWidths(widths ...float64) *PDF {
	if len(widths) < 1 {
		pdf.err("len(widths) < 1")
		return pdf
	}
	if len(widths) > 100 {
		pdf.err("len(widths) > 100")
		return pdf
	}
	pdf.columnWidths = widths
	return pdf
} //                                                             SetColumnWidths

// TextWidth returns the width of the text in current units.
func (pdf *PDF) TextWidth(text string) float64 {
	if pdf.warnIfNoPage() {
		return 0
	}
	return pdf.textWidthPt1000(text) / pdf.pointsPerUnit
} //                                                                   TextWidth

// WrapTextLines splits a string into multiple lines so that the text
// fits in the specified width. The text is wrapped on word boundaries.
// Newline characters (CR and "\n") also cause text to be split.
// You can find out the number of lines needed to wrap some
// text by checking the length of the returned array.
func (pdf *PDF) WrapTextLines(width float64, text string) []string {
	var ar = []string{text} //                      first handle the newlines...
	{
		var split = func(ar []string, sep string) []string {
			var ret []string
			for _, iter := range ar {
				if strings.Contains(iter, sep) {
					ret = append(ret, strings.Split(iter, sep)...)
				} else {
					ret = append(ret, iter)
				}
			}
			return ret
		}
		ar = split(split(split(ar, "\r\n"), "\r"), "\n")
	}
	var ret []string //                  ...then break lines based on text width
	for _, iter := range ar {
		for pdf.TextWidth(iter) > width {
			var max = len(iter)
			var ln = max
			for ln > 0 { //                    take half less text until it fits
				ln /= 2
				if pdf.TextWidth(iter[:ln]) <= width {
					break
				}
			}
			var inc = ln / 2 //         take half more text until it doesn't fit
			ln += inc
			for ln > 0 && ln < max &&
				pdf.TextWidth(iter[:ln]) <= width {
				if inc > 1 {
					inc /= 2
				}
				ln += inc
			}
			ln-- //                     take less by 1 character until text fits
			for ln > 0 && pdf.TextWidth(iter[:ln]) > width {
				ln--
			}
			var found bool //    move to the last word (if white-space is found)
			max = ln
			for ln > 0 {
				if pdf.isWhiteSpace(iter[ln-1 : ln]) {
					found = true
					break
				}
				ln--
			}
			if !found {
				ln = max
			}
			ret = append(ret, iter[:ln])
			iter = iter[ln-1:]
		}
		ret = append(ret, iter)
	}
	return ret
} //                                                               WrapTextLines

// -----------------------------------------------------------------------------
// # Functions:

// ToPoints converts a string composed of a number and unit
// to points. For example '1 cm' or '1cm' becomes 28.346 points.
//
// Recognised units are:
// mm cm " in inch inches tw twip twips pt point points
func (pdf *PDF) ToPoints(numberAndUnit string) float64 {
	numberAndUnit = strings.ToUpper(
		strings.Trim(numberAndUnit, " \a\b\f\n\r\t\v"),
	)
	if numberAndUnit == "" {
		return 0
	}
	var num, unit string //            read value and unit into separate strings
	for _, ch := range numberAndUnit {
		if (ch >= '0' && ch <= '9') || ch == '.' || ch == '-' {
			num += string(ch)
			continue
		}
		if (ch >= 'A' && ch <= 'Z') ||
			(ch >= 'a' && ch <= 'z') ||
			ch == '"' {
			unit += strings.ToUpper(string(ch))
		}
	}
	var ret, _ = strconv.ParseFloat(num, 64)
	if unit != "" { //                       determine number of points per unit
		var ppu = pdf.getPointsPerUnit(unit)
		if int(ppu*1000000) == 0 {
			pdf.err("Unknown unit name.")
		}
		ret *= ppu
	}
	return ret
} //                                                                    ToPoints

// -----------------------------------------------------------------------------
// # Private Methods

// applyLineWidth writes a line-width PDF command ('w') to the current
// page's content stream, if the value changed since last call.
// called by: DrawBox(), DrawLine(), FillBox()
func (pdf *PDF) applyLineWidth() *PDF {
	var val = &pdf.pagePtr.lineWidth
	if int(*val*10000) != int(pdf.lineWidth*10000) {
		*val = pdf.lineWidth
		pdf.printf("%.3f w\n", float64(*val))
	}
	return pdf
} //                                                              applyLineWidth

// applyNonStrokeColor writes a text color selection PDF command ('rg') to
// the current page's content stream, if the value changed since last call.
// called by: drawTextLine(), FillBox()
func (pdf *PDF) applyNonStrokeColor() *PDF {
	var val = &pdf.pagePtr.nonStrokeColor
	if !pdf.colorEqual(*val, pdf.color) {
		*val = pdf.color
		pdf.printf(
			"%.3f %.3f %.3f rg\n", // non-stroking (text) color
			float64((*val).Red)/255,
			float64((*val).Green)/255,
			float64((*val).Blue)/255,
		)
	}
	return pdf
} //                                                         applyNonStrokeColor

// applyStrokeColor writes a line color selection PDF
// command ('RG') to the current page's content
// stream, if the value changed since last call.
// called by: DrawBox(), DrawLine()
func (pdf *PDF) applyStrokeColor() *PDF {
	var val = &pdf.pagePtr.strokeColor
	if !pdf.colorEqual(*val, pdf.color) {
		*val = pdf.color
		pdf.printf(
			"%.3f %.3f %.3f RG\n", // RG - stroke (line) color
			float64((*val).Red)/255,
			float64((*val).Green)/255,
			float64((*val).Blue)/255,
		)
	}
	return pdf
} //                                                            applyStrokeColor

// drawTextLine writes a line of text at the current coordinates to the
// current page's content stream, using a sequence of raw PDF commands.
// called by: DrawText(), drawTextBox()
func (pdf *PDF) drawTextLine(text string) *PDF {
	if text == "" {
		return pdf
	}
	// applyFont writes a font change command, provided the font has
	// been changed since the last operation that uses fonts.
	//
	// This should be called just before a font needs to be used.
	// This way, if a font is picked with SetFontName() property, but
	// never used to draw text, no font selection command is output.
	//
	// Before calling this method, the font name must be already
	// set by SetFontName(), which is stored in pdf.font.fontName
	//
	// What this function does:
	// - Validates the current font name and determines if it is a
	//   standard (built-in) font like Helvetica or a TrueType font.
	// - Fills the document-wide list of fonts (pdf.fonts).
	// - Adds items to the list of font ID's used on the current page.
	var applyFont = func() {
		var isValid = pdf.fontName != ""
		var font PDFFont
		if isValid {
			isValid = false
			for i, name := range PDFBuiltInFontNames {
				name = strings.ToUpper(name)
				if strings.ToUpper(pdf.fontName) != name {
					continue
				}
				font.fontName = PDFBuiltInFontNames[i]
				font.isBuiltIn = true
				font.isBold = strings.Contains(name, "BOLD")
				font.isItalic = strings.Contains(name, "OBLIQUE") ||
					strings.Contains(name, "ITALIC")
				isValid = true
				break
			}
		}
		// if there is no selected font or it's invalid, use Helvetica
		if !isValid {
			font.fontName = "Helvetica"
			font.isBuiltIn = true
			font.isBold = false
			font.isItalic = false
		}
		// has the font been added to the global list? If not, add it:
		for _, iter := range pdf.fonts {
			if font.fontName == iter.fontName {
				font.fontID = iter.fontID
				break
			}
		}
		if font.fontID == 0 {
			font.fontID = 1 + len(pdf.fonts)
			pdf.fonts = append(pdf.fonts, font)
		}
		var pg = pdf.pagePtr
		if pg.fontID == font.fontID &&
			(int(pg.fontSizePt*1000) == int(pdf.fontSizePt)*1000) {
			return
		}
		// add the font ID to the current page, if not already referenced
		var alreadyUsedOnPage bool
		for _, id := range pg.fontIDs {
			if id == font.fontID {
				alreadyUsedOnPage = true
				break
			}
		}
		if !alreadyUsedOnPage {
			pg.fontIDs = append(pg.fontIDs, 0)
			pg.fontIDs[len(pg.fontIDs)-1] = font.fontID
		}
		pg.fontID = font.fontID
		pg.fontSizePt = pdf.fontSizePt
		pdf.printf("BT /F%d %d Tf ET\n", pg.fontID, int(pg.fontSizePt))
	}
	// applyHorizontalScaling sets the horizontal text scaling
	var applyHorizontalScaling = func() {
		if pdf.pagePtr.horizontalScaling != pdf.horizontalScaling {
			pdf.pagePtr.horizontalScaling = pdf.horizontalScaling
			pdf.printf("BT %d Tz ET\n", pdf.pagePtr.horizontalScaling)
		}
	}
	// draw the text:
	applyFont()
	applyHorizontalScaling()
	pdf.applyNonStrokeColor()
	var pg = pdf.pagePtr
	if pg.x < 0 || pg.y < 0 {
		pdf.SetXY(0, 0)
	}
	pdf.printf("BT %d %d Td (%s) Tj ET\n",
		int(pg.x), int(pg.y), pdf.escape(text))
	pg.x += pdf.textWidthPt1000(text)
	return pdf
} //                                                                drawTextLine

// drawTextBox draws a line of text, or a word-wrapped block of text.
// called by: DrawTextAlignedToBox(), DrawTextInBox()
func (pdf *PDF) drawTextBox(
	x, y, width, height float64, wrapText bool, align, text string,
) *PDF {
	if text == "" {
		return pdf
	}
	if pdf.pageNo < 0 {
		pdf.err("No current page.")
		return pdf
	}
	var lines = (func() []string {
		if wrapText {
			return pdf.WrapTextLines(width, text)
		}
		return []string{text}
	})() // IIFE
	var lineHeight = pdf.FontSize()
	var allLinesHeight = lineHeight * float64(len(lines))
	//
	// calculate x-axis position of text (left, right, center)
	x, width = x*pdf.pointsPerUnit, width*pdf.pointsPerUnit
	var alignX = func(text string) float64 {
		for _, r := range align {
			if r == 'l' || r == 'L' {
				return pdf.fontSizePt / 6 // add margin
			} else if r == 'r' || r == 'R' {
				return width - pdf.textWidthPt1000(text) - pdf.fontSizePt/6
			}
		}
		return width/2 - pdf.textWidthPt1000(text)/2 // center
	}
	// calculate aligned y-axis position of text (top, bottom, center)
	y, height = y*pdf.pointsPerUnit, height*pdf.pointsPerUnit
	y += pdf.fontSizePt + (func() float64 {
		for _, r := range align {
			if r == 't' || r == 'T' {
				return 0 // top
			} else if r == 'b' || r == 'B' {
				return height - allLinesHeight - 4 // bottom
			}
		}
		return height/2 - allLinesHeight/2 - pdf.fontSizePt*0.15 // center
	})() // IIFE
	y = pdf.pageHeightPt - y
	for _, line := range lines {
		pdf.pagePtr.x, pdf.pagePtr.y = x+alignX(line), y
		pdf.drawTextLine(line)
		y -= lineHeight
	}
	return pdf
} //                                                                 drawTextBox

// printf writes formatted strings (like fmt.Sprintf) to the current page's
// content stream or to the final generated PDF, if there is no active page.
// called by: Bytes(), DrawBox(), DrawLine(), drawTextLine(), FillBox()
//            applyLineWidth(), applyNonStrokeColor(), applyStrokeColor()
func (pdf *PDF) printf(format string, args ...interface{}) *PDF {
	var buf *bytes.Buffer
	if pdf.pageNo == PDFNoPage {
		buf = pdf.contentPtr
	} else if pdf.pageNo > (len(pdf.pages) - 1) {
		pdf.err("Invalid page number.")
		return pdf
	} else {
		buf = &pdf.pagePtr.pageContent
	}
	buf.Write([]byte(fmt.Sprintf(format, args...)))
	return pdf
} //                                                                      printf

// setCurrentPage selects the currently-active page.
// called by: AddPage(), Bytes()
func (pdf *PDF) setCurrentPage(pageNo int) *PDF {
	if pageNo != PDFNoPage && pageNo > (len(pdf.pages)-1) {
		pdf.err("Page number out of range.")
	} else if pageNo == PDFNoPage {
		pdf.pagePtr = nil
		pdf.contentPtr = &pdf.content
	} else {
		pdf.pagePtr = &pdf.pages[pageNo]
		pdf.contentPtr = &pdf.pagePtr.pageContent
	}
	pdf.pageNo = pageNo
	return pdf
} //                                                              setCurrentPage

// textWidthPt1000 returns the width of text in thousands of a point.
// called by: drawTextLine(), drawTextBox(), TextWidth()
func (pdf *PDF) textWidthPt1000(text string) float64 {
	//
	// warn and return if there is no current page
	if pdf.pageNo == PDFNoPage || pdf.pageNo > (len(pdf.pages)-1) ||
		pdf.pagePtr == nil {
		pdf.err("No current page.")
		return 0
	}
	if text == "" {
		return 0
	}
	var w = 0.0
	var charWidths = pdfFontWidths[0]
	// TODO:                       ^ this is not considering the font name!
	//
	for i, r := range text {
		if r < 0 || r > 255 {
			pdf.err("char out of range at %d: %d", i, r)
			break
		}
		w += float64(charWidths[r])
	}
	return w * pdf.fontSizePt / 1000 * float64(pdf.horizontalScaling) / 100
} //                                                             textWidthPt1000

// warnIfNoPage outputs a warning and returns true if there is no
// active page. This can only happen when the user didn't call AddPage().
// called by: DrawBox(), DrawLine(), DrawText(), FillBox(),
//            SetX(), SetXY(), SetY(), TextWidth(), X(), Y()
func (pdf *PDF) warnIfNoPage() bool {
	if len(pdf.pages) == 0 || pdf.pageNo > (len(pdf.pages)-1) ||
		pdf.pagePtr == nil {
		pdf.err("No current page.")
		return true
	}
	return false
} //                                                                warnIfNoPage

// -----------------------------------------------------------------------------
// # Private Generation Methods

// streamData __
func (pdf *PDF) streamData(content []byte) {
	pdf.setCurrentPage(PDFNoPage)
	var filter string
	if pdf.compressStreams {
		filter = "/Filter/FlateDecode"
		var buf bytes.Buffer
		var writer = zlib.NewWriter(&buf)
		var _, err = writer.Write([]byte(content))
		if err != nil {
			pdf.err("Failed compressing:", err)
			return
		}
		writer.Close()
		content = buf.Bytes()
	}
	pdf.printf("%s/Length %d>>stream\n%s\nendstream\n",
		filter, len(content), content)
} //                                                                  streamData

// -----------------------------------------------------------------------------
// # Private Functions

// colorEqual compares two PDFColor values
// and returns true if they are equal.
//
// called by: applyNonStrokeColor(), applyStrokeColor()
func (*PDF) colorEqual(a, b PDFColor) bool {
	return a.Red == b.Red && a.Green == b.Green && a.Blue == b.Blue
} //                                                                  colorEqual

// err reports an error
func (*PDF) err(a ...interface{}) {
	fmt.Println(a)
} //                                                                         err

// escape escapes special characters '(', '(' and '\' in strings
// in order to avoid them interfering with PDF commands.
// called by: Bytes(), drawTextLine()
func (*PDF) escape(s string) []byte {
	if strings.Contains(s, "(") || strings.Contains(s, ")") ||
		strings.Contains(s, "\\") {
		//
		var writer = bytes.NewBuffer(make([]byte, 0, len(s)))
		for _, r := range s {
			if r == '(' || r == ')' || r == '\\' {
				writer.WriteRune('\\')
			}
			writer.WriteRune(r)
		}
		return writer.Bytes()
	}
	return []byte(s)
} //                                                                      escape

// isWhiteSpace returns true if all the
// characters in a string are white-spaces.
func (*PDF) isWhiteSpace(s string) bool {
	if s == "" {
		return false
	}
	for _, ch := range s {
		if ch != ' ' && ch != '\a' &&
			ch != '\b' && ch != '\f' &&
			ch != '\n' && ch != '\r' &&
			ch != '\t' && ch != '\v' {
			return false
		}
	}
	return true
} //                                                                isWhiteSpace

// getPointsPerUnit returns the number of points
// per a named unit of measurement.
// called by: SetUnits(), ToPoints()
func (*PDF) getPointsPerUnit(unitName string) float64 {
	switch strings.Trim(strings.ToUpper(unitName), " \a\b\f\n\r\t\v") {
	case "MM":
		return 2.83464566929134 // 1 inch / 25.4mm per inch * 72 points per in.
	case "CM":
		return 28.3464566929134 // 1 inch / 2.54cm per inch * 72 points per in.
	case "IN", "INCH", "INCHES", `"`:
		return 72.0 // points per inch
	case "TW", "TWIP", "TWIPS":
		return 0.05 // 1 point / 20 twips per point
	case "PT", "POINT", "POINTS":
		return 1.0 // point
	}
	return 0
} //                                                            getPointsPerUnit

//end
