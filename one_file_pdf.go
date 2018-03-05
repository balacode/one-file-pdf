// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-03-06 00:28:18 6FA61A                              [one_file_pdf.go]
// -----------------------------------------------------------------------------

package pdf

// # Structures
//   PDF struct
//   PDFColor struct
//   PDFPageSize struct
//   PDFPageSizeOf(pageSize string) PDFPageSize
//
// # Constants
//   PDFColorNames = map[string]PDFColor
//   PDFNoPage = -1
//   PDFStandardPageSizes = []PDFPageSize
//
// # Internal Structures
//   pdfFont struct
//   pdfImage struct
//   pdfPage struct
//
// # Internal Constants
//   pdfFontNames = []string
//   pdfFontWidths = [][]int
//   pdfPagesIndex = 3
//
// # Constructor
//   NewPDF(pageSize string) PDF
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
//   (pdf *PDF) SetColor(nameOrHTMLColor string) *PDF
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
//   (pdf *PDF) DrawImage(x, y, height float64, fileNameOrBytes interface{},
//              ) *PDF
//   (pdf *PDF) DrawLine(x1, y1, x2, y2 float64) *PDF
//   (pdf *PDF) DrawText(text string) *PDF
//   (pdf *PDF) DrawTextAlignedToBox(
//                 x, y, width, height float64, align, text string,
//              ) *PDF
//   (pdf *PDF) DrawTextAt(x, y float64, text string) *PDF
//   (pdf *PDF) DrawTextInBox(x, y, width, height float64, align, text string,
//              ) *PDF
//   (pdf *PDF) DrawUnitGrid() *PDF
//   (pdf *PDF) FillBox(x, y, width, height float64) *PDF
//   (pdf *PDF) NextLine() *PDF
//   (pdf *PDF) Reset() *PDF
//   (pdf *PDF) SaveFile(filename string) *PDF
//   (pdf *PDF) SetColumnWidths(widths ...float64) *PDF
//   (pdf *PDF) TextWidth(text string) float64
//   (pdf *PDF) WrapTextLines(width float64, text string) []string
//
// # Functions:
//   (pdf *PDF) ToPoints(numberAndUnit string) float64
//
// # Private Methods
//   (pdf *PDF) applyFont()
//   (pdf *PDF) applyLineWidth() *PDF
//   (pdf *PDF) applyNonStrokeColor() *PDF
//   (pdf *PDF) applyStrokeColor() *PDF
//   (pdf *PDF) drawTextLine(text string) *PDF
//   (pdf *PDF) drawTextBox(
//                 x, y, width, height float64,
//                 wrapText bool, align, text string,
//              ) *PDF
//   (pdf *PDF) setCurrentPage(pageNo int) *PDF
//   (pdf *PDF) textWidthPt1000(text string) float64
//   (pdf *PDF) warnIfNoPage() bool
//
// # Private Generation Methods
//   (pdf *PDF) nextObj() int
//   (pdf *PDF) write(format string, args ...interface{}) *PDF
//   (pdf *PDF) writeEndobj() *PDF
//   (pdf *PDF) writeObj(objectType string) *PDF
//   (pdf *PDF) writePages(fontsIndex, imagesIndex int) *PDF
//   (pdf *PDF) writeStream(content []byte) *PDF
//   (pdf *PDF) writeStreamData(content []byte) *PDF
//
// # Private Functions
//   (*PDF) escape(s string) []byte
//   (*PDF) getPointsPerUnit(unitName string) float64
//   (*PDF) isWhiteSpace(s string) bool
//   (*PDF) splitLines(s string) []string
//   (PDF) logError(a ...interface{})

import "bytes"         // standard
import "compress/zlib" // standard
import "fmt"           // standard
import "image"         // standard
import "io/ioutil"     // standard
import "strconv"       // standard
import "strings"       // standard
import _ "image/png"   // standard

// PDFErrorHandler is the function that handles errors.
// You can redefine it as needed or set to nil to mute error messages.
var PDFErrorHandler = fmt.Println

// -----------------------------------------------------------------------------
// # Structures

// PDF is the main structure representing a PDF document.
type PDF struct {
	docAuthor         string        // 'author' metadata entry
	docCreator        string        // 'creator' metadata entry
	docKeywords       string        // 'keywords' metadata entry
	docSubject        string        // 'subject' metadata entry
	docTitle          string        // 'title' metadata entry
	pageSize          PDFPageSize   // page size used in this PDF
	pageNo            int           // current page number
	pagePtr           *pdfPage      // pointer to the current page
	pages             []pdfPage     // all the pages added to this PDF
	fonts             []pdfFont     // all the fonts used in this PDF
	images            []pdfImage    // all the images used in this PDF
	columnWidths      []float64     // user-set column widths (like tab stops)
	columnNo          int           // index of the current column
	unitName          string        // name of active measurement unit
	pointsPerUnit     float64       // number of points per measurement unit
	color             PDFColor      // current drawing color
	lineWidth         float64       // current line width (in points)
	fontName          string        // current font's name
	fontSizePt        float64       // current font's size (in points)
	horizontalScaling uint16        // horizontal scaling factor (in %)
	compressStreams   bool          // enable stream compression?
	content           bytes.Buffer  // content buffer where PDF is written
	contentPtr        *bytes.Buffer // pointer to PDF/current page's buffer
	objOffsets        []int         // used by Bytes() and write..()
	objNo             int           // used by Bytes() and write..()
} //                                                                         PDF

// PDFColor represents a color value.
type PDFColor struct {
	Red   uint8
	Green uint8
	Blue  uint8
} //                                                                    PDFColor

// PDFPageSize represents a page size name and its dimensions in points.
type PDFPageSize struct {
	Name     string
	WidthPt  float64 // width in points
	HeightPt float64 // height in points
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

// PDFColorNames maps web (X11) color names to values.
// (from https://en.wikipedia.org/wiki/X11_color_names)
var PDFColorNames = map[string]PDFColor{
	"ALICEBLUE":            {240, 248, 255}, // #F0F8FF
	"ANTIQUEWHITE":         {250, 235, 215}, // #FAEBD7
	"AQUA":                 {0, 255, 255},   // #00FFFF
	"AQUAMARINE":           {127, 255, 212}, // #7FFFD4
	"AZURE":                {240, 255, 255}, // #F0FFFF
	"BEIGE":                {245, 245, 220}, // #F5F5DC
	"BISQUE":               {255, 228, 196}, // #FFE4C4
	"BLACK":                {0, 0, 0},       // #000000
	"BLANCHEDALMOND":       {255, 235, 205}, // #FFEBCD
	"BLUE":                 {0, 0, 255},     // #0000FF
	"BLUEVIOLET":           {138, 43, 226},  // #8A2BE2
	"BROWN":                {165, 42, 42},   // #A52A2A
	"BURLYWOOD":            {222, 184, 135}, // #DEB887
	"CADETBLUE":            {95, 158, 160},  // #5F9EA0
	"CHARTREUSE":           {127, 255, 0},   // #7FFF00
	"CHOCOLATE":            {210, 105, 30},  // #D2691E
	"CORAL":                {255, 127, 80},  // #FF7F50
	"CORNFLOWERBLUE":       {100, 149, 237}, // #6495ED
	"CORNSILK":             {255, 248, 220}, // #FFF8DC
	"CRIMSON":              {220, 20, 60},   // #DC143C
	"CYAN":                 {0, 255, 255},   // #00FFFF
	"DARKBLUE":             {0, 0, 139},     // #00008B
	"DARKCYAN":             {0, 139, 139},   // #008B8B
	"DARKGOLDENROD":        {184, 134, 11},  // #B8860B
	"DARKGRAY":             {169, 169, 169}, // #A9A9A9
	"DARKGREEN":            {0, 100, 0},     // #006400
	"DARKKHAKI":            {189, 183, 107}, // #BDB76B
	"DARKMAGENTA":          {139, 0, 139},   // #8B008B
	"DARKOLIVEGREEN":       {85, 107, 47},   // #556B2F
	"DARKORANGE":           {255, 140, 0},   // #FF8C00
	"DARKORCHID":           {153, 50, 204},  // #9932CC
	"DARKRED":              {139, 0, 0},     // #8B0000
	"DARKSALMON":           {233, 150, 122}, // #E9967A
	"DARKSEAGREEN":         {143, 188, 143}, // #8FBC8F
	"DARKSLATEBLUE":        {72, 61, 139},   // #483D8B
	"DARKSLATEGRAY":        {47, 79, 79},    // #2F4F4F
	"DARKTURQUOISE":        {0, 206, 209},   // #00CED1
	"DARKVIOLET":           {148, 0, 211},   // #9400D3
	"DEEPPINK":             {255, 20, 147},  // #FF1493
	"DEEPSKYBLUE":          {0, 191, 255},   // #00BFFF
	"DIMGRAY":              {105, 105, 105}, // #696969
	"DODGERBLUE":           {30, 144, 255},  // #1E90FF
	"FIREBRICK":            {178, 34, 34},   // #B22222
	"FLORALWHITE":          {255, 250, 240}, // #FFFAF0
	"FORESTGREEN":          {34, 139, 34},   // #228B22
	"FUCHSIA":              {255, 0, 255},   // #FF00FF
	"GAINSBORO":            {220, 220, 220}, // #DCDCDC
	"GHOSTWHITE":           {248, 248, 255}, // #F8F8FF
	"GOLD":                 {255, 215, 0},   // #FFD700
	"GOLDENROD":            {218, 165, 32},  // #DAA520
	"GRAY":                 {190, 190, 190}, // #BEBEBE X11 Version
	"GREEN":                {0, 255, 0},     // #00FF00 X11 Version
	"GREENYELLOW":          {173, 255, 47},  // #ADFF2F
	"HONEYDEW":             {240, 255, 240}, // #F0FFF0
	"HOTPINK":              {255, 105, 180}, // #FF69B4
	"INDIANRED":            {205, 92, 92},   // #CD5C5C
	"INDIGO":               {75, 0, 130},    // #4B0082
	"IVORY":                {255, 255, 240}, // #FFFFF0
	"KHAKI":                {240, 230, 140}, // #F0E68C
	"LAVENDER":             {230, 230, 250}, // #E6E6FA
	"LAVENDERBLUSH":        {255, 240, 245}, // #FFF0F5
	"LAWNGREEN":            {124, 252, 0},   // #7CFC00
	"LEMONCHIFFON":         {255, 250, 205}, // #FFFACD
	"LIGHTBLUE":            {173, 216, 230}, // #ADD8E6
	"LIGHTCORAL":           {240, 128, 128}, // #F08080
	"LIGHTCYAN":            {224, 255, 255}, // #E0FFFF
	"LIGHTGOLDENRODYELLOW": {250, 250, 210}, // #FAFAD2
	"LIGHTGRAY":            {211, 211, 211}, // #D3D3D3
	"LIGHTGREEN":           {144, 238, 144}, // #90EE90
	"LIGHTPINK":            {255, 182, 193}, // #FFB6C1
	"LIGHTSALMON":          {255, 160, 122}, // #FFA07A
	"LIGHTSEAGREEN":        {32, 178, 170},  // #20B2AA
	"LIGHTSKYBLUE":         {135, 206, 250}, // #87CEFA
	"LIGHTSLATEGRAY":       {119, 136, 153}, // #778899
	"LIGHTSTEELBLUE":       {176, 196, 222}, // #B0C4DE
	"LIGHTYELLOW":          {255, 255, 224}, // #FFFFE0
	"LIME":                 {0, 255, 0},     // #00FF00
	"LIMEGREEN":            {50, 205, 50},   // #32CD32
	"LINEN":                {250, 240, 230}, // #FAF0E6
	"MAGENTA":              {255, 0, 255},   // #FF00FF
	"MAROON":               {176, 48, 96},   // #B03060 X11 Version
	"MEDIUMAQUAMARINE":     {102, 205, 170}, // #66CDAA
	"MEDIUMBLUE":           {0, 0, 205},     // #0000CD
	"MEDIUMORCHID":         {186, 85, 211},  // #BA55D3
	"MEDIUMPURPLE":         {147, 112, 219}, // #9370DB
	"MEDIUMSEAGREEN":       {60, 179, 113},  // #3CB371
	"MEDIUMSLATEBLUE":      {123, 104, 238}, // #7B68EE
	"MEDIUMSPRINGGREEN":    {0, 250, 154},   // #00FA9A
	"MEDIUMTURQUOISE":      {72, 209, 204},  // #48D1CC
	"MEDIUMVIOLETRED":      {199, 21, 133},  // #C71585
	"MIDNIGHTBLUE":         {25, 25, 112},   // #191970
	"MINTCREAM":            {245, 255, 250}, // #F5FFFA
	"MISTYROSE":            {255, 228, 225}, // #FFE4E1
	"MOCCASIN":             {255, 228, 181}, // #FFE4B5
	"NAVAJOWHITE":          {255, 222, 173}, // #FFDEAD
	"NAVY":                 {0, 0, 128},     // #000080
	"OLDLACE":              {253, 245, 230}, // #FDF5E6
	"OLIVE":                {128, 128, 0},   // #808000
	"OLIVEDRAB":            {107, 142, 35},  // #6B8E23
	"ORANGE":               {255, 165, 0},   // #FFA500
	"ORANGERED":            {255, 69, 0},    // #FF4500
	"ORCHID":               {218, 112, 214}, // #DA70D6
	"PALEGOLDENROD":        {238, 232, 170}, // #EEE8AA
	"PALEGREEN":            {152, 251, 152}, // #98FB98
	"PALETURQUOISE":        {175, 238, 238}, // #AFEEEE
	"PALEVIOLETRED":        {219, 112, 147}, // #DB7093
	"PAPAYAWHIP":           {255, 239, 213}, // #FFEFD5
	"PEACHPUFF":            {255, 218, 185}, // #FFDAB9
	"PERU":                 {205, 133, 63},  // #CD853F
	"PINK":                 {255, 192, 203}, // #FFC0CB
	"PLUM":                 {221, 160, 221}, // #DDA0DD
	"POWDERBLUE":           {176, 224, 230}, // #B0E0E6
	"PURPLE":               {160, 32, 240},  // #A020F0 X11 Version
	"REBECCAPURPLE":        {102, 51, 153},  // #663399
	"RED":                  {255, 0, 0},     // #FF0000
	"ROSYBROWN":            {188, 143, 143}, // #BC8F8F
	"ROYALBLUE":            {65, 105, 225},  // #4169E1
	"SADDLEBROWN":          {139, 69, 19},   // #8B4513
	"SALMON":               {250, 128, 114}, // #FA8072
	"SANDYBROWN":           {244, 164, 96},  // #F4A460
	"SEAGREEN":             {46, 139, 87},   // #2E8B57
	"SEASHELL":             {255, 245, 238}, // #FFF5EE
	"SIENNA":               {160, 82, 45},   // #A0522D
	"SILVER":               {192, 192, 192}, // #C0C0C0
	"SKYBLUE":              {135, 206, 235}, // #87CEEB
	"SLATEBLUE":            {106, 90, 205},  // #6A5ACD
	"SLATEGRAY":            {112, 128, 144}, // #708090
	"SNOW":                 {255, 250, 250}, // #FFFAFA
	"SPRINGGREEN":          {0, 255, 127},   // #00FF7F
	"STEELBLUE":            {70, 130, 180},  // #4682B4
	"TAN":                  {210, 180, 140}, // #D2B48C
	"TEAL":                 {0, 128, 128},   // #008080
	"THISTLE":              {216, 191, 216}, // #D8BFD8
	"TOMATO":               {255, 99, 71},   // #FF6347
	"TURQUOISE":            {64, 224, 208},  // #40E0D0
	"VIOLET":               {238, 130, 238}, // #EE82EE
	"WEBGRAY":              {128, 128, 128}, // #808080 Web Version
	"WEBGREEN":             {0, 128, 0},     // #008000 Web Version
	"WEBMAROON":            {127, 0, 0},     // #7F0000 Web Version
	"WEBPURPLE":            {127, 0, 127},   // #7F007F Web Version
	"WHEAT":                {245, 222, 179}, // #F5DEB3
	"WHITE":                {255, 255, 255}, // #FFFFFF
	"WHITESMOKE":           {245, 245, 245}, // #F5F5F5
	"YELLOW":               {255, 255, 0},   // #FFFF00
	"YELLOWGREEN":          {154, 205, 50},  // #9ACD32
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

// -----------------------------------------------------------------------------
// # Internal Structures

// pdfFont represents a font name and its appearance.
type pdfFont struct {
	fontID    int
	fontName  string
	isBuiltIn bool
	isBold    bool
	isItalic  bool
} //                                                                     pdfFont

// pdfImage represents an image.
type pdfImage struct {
	name      string
	width     int
	height    int
	data      []byte
	grayscale bool
} //                                                                    pdfImage

// pdfPage holds details for each page.
type pdfPage struct {
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
} //                                                                     pdfPage

// -----------------------------------------------------------------------------
// # Internal Constants

// pdfFontNames lists the names of built-in font names
// these fonts must be available on all PDF implementations
var pdfFontNames = []string{
	"Courier",               // fixed-width
	"Courier-Bold",          // fixed-width
	"Courier-BoldOblique",   // fixed-width
	"Courier-Oblique",       // fixed-width
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
} //                                                                pdfFontNames

// pdfFontWidths specifies the widths of build-in fonts
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

// pdfPagesIndex is the starting page object index (after Catalog and Pages)
const pdfPagesIndex = 3

// -----------------------------------------------------------------------------
// # Constructor

// NewPDF creates and initializes a new PDF object.
func NewPDF(pageSize string) PDF {
	//
	// get and store dimensions of specified page size (in points)
	var size = PDFPageSizeOf(pageSize)
	if size.Name == "" {
		PDF{}.logError("Unknown page size ", pageSize, ". Setting to 'A4'.\n")
		size = PDFPageSizeOf("A4")
	}
	// create a new PDF object
	var pdf = PDF{
		pageNo:            PDFNoPage,
		pageSize:          size,
		horizontalScaling: 100,
		compressStreams:   true,
	}
	// set default units, otherwise pointsPerUnit, x and y will be 0
	pdf.SetUnits("point")
	return pdf
} //                                                                      NewPDF

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
// E.g.: MM CM " IN INCH INCHES TW TWIP TWIPS PT POINT POINTS
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
	return (pdf.pageSize.HeightPt - pdf.pagePtr.y) / pdf.pointsPerUnit
} //                                                                           Y

// -----------------------------------------------------------------------------
// # Property Setters

// SetColor sets the current color using a web/X11 color name
// (e.g. "HoneyDew") or HTML color value such as "#191970"
// for midnight blue (#RRGGBB). The current color is used
// for subsequent text and line drawing and fills.
// If the name is unknown or valid, sets the current color to black.
func (pdf *PDF) SetColor(nameOrHTMLColor string) *PDF {
	//
	// if name starts with '#' treat it as HTML color (#RRGGBB)
	nameOrHTMLColor = strings.ToUpper(nameOrHTMLColor)
	if nameOrHTMLColor != "" && nameOrHTMLColor[0] == '#' {
		var hex [6]uint8
		for i, ch := range nameOrHTMLColor[1:] {
			if i > 6 {
				break
			}
			if ch >= '0' && ch <= '9' {
				hex[i] = uint8(ch - '0')
				continue
			}
			if ch >= 'A' && ch <= 'F' {
				hex[i] = uint8(ch - 'A' + 10)
				continue
			}
			pdf.logError("Invalid color code '" + nameOrHTMLColor + "'." +
				"Setting to black.")
			pdf.SetColorRGB(0, 0, 0)
			return pdf
		}
		pdf.SetColorRGB(
			int(hex[0]*16+hex[1]),
			int(hex[2]*16+hex[3]),
			int(hex[4]*16+hex[5]),
		)
		return pdf
	}
	// otherwise search for color name
	var color, exists = PDFColorNames[nameOrHTMLColor]
	if exists {
		pdf.SetColorRGB(int(color.Red), int(color.Green), int(color.Blue))
		return pdf
	}
	pdf.logError("Color name '" + nameOrHTMLColor + "' not known." +
		"Setting to black.")
	pdf.SetColorRGB(0, 0, 0)
	return pdf
} //                                                                    SetColor

// SetColorRGB sets the current color using separate
// red, green and blue values. Each component value
// can range from 0 to 255. The current color is used
// for subsequent text and line drawing and fills.
func (pdf *PDF) SetColorRGB(red, green, blue int) *PDF {
	pdf.color = PDFColor{uint8(red), uint8(green), uint8(blue)}
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
	return pdf.SetFontName(name).SetFontSize(points)
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
// Can be upper/lowercase:
// MM CM " IN INCH INCHES TW TWIP TWIPS PT POINT POINTS.
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
	x, y = x*pdf.pointsPerUnit, pdf.pageSize.HeightPt-y*pdf.pointsPerUnit
	pdf.pagePtr.x, pdf.pagePtr.y = x, y
	return pdf
} //                                                                       SetXY

// SetY changes the Y-coordinate of the current drawing position.
func (pdf *PDF) SetY(y float64) *PDF {
	if pdf.warnIfNoPage() {
		pdf.pagePtr.y = (pdf.pageSize.HeightPt * pdf.pointsPerUnit) -
			(y * pdf.pointsPerUnit)
	}
	return pdf
} //                                                                        SetY

// -----------------------------------------------------------------------------
// # Methods

// AddPage appends a new blank page to the document and makes it selected.
func (pdf *PDF) AddPage() *PDF {
	var i = len(pdf.pages)
	pdf.pages = append(pdf.pages, pdfPage{})
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
// Called by: SaveFile()
func (pdf *PDF) Bytes() []byte {
	pdf.objOffsets = []int{}
	pdf.objNo = 0
	//
	// free any existing generated content and write beginning of document
	var fontsIndex = pdfPagesIndex + len(pdf.pages)*2
	var imagesIndex = fontsIndex + len(pdf.fonts)
	var infoIndex int // set when metadata found
	pdf.content.Reset()
	pdf.setCurrentPage(PDFNoPage).write("%%PDF-1.4\n").
		writeObj("/Catalog").write("/Pages 2 0 R").writeEndobj()
	//
	//  write /Pages object (2 0 obj), page count, page size and pages
	pdf.writeObj("/Pages").
		write("/Count %d/MediaBox[0 0 %d %d]",
							len(pdf.pages),
							int(pdf.pageSize.WidthPt),
							int(pdf.pageSize.HeightPt)).
		writePages(fontsIndex, imagesIndex) // page numbers and pages
	//
	// write fonts
	for _, iter := range pdf.fonts {
		pdf.writeObj("/Font")
		if iter.isBuiltIn {
			pdf.write("/Subtype/Type1/Name/F%d"+
				"/BaseFont/%s/Encoding/WinAnsiEncoding",
				iter.fontID, iter.fontName)
		}
		pdf.writeEndobj()
	}
	// write images
	for _, iter := range pdf.images {
		pdf.writeObj("/XObject").
			write("/Subtype/Image/Width %d/Height %d"+
				"/ColorSpace/DeviceGray/BitsPerComponent 8",
				iter.width, iter.height).
			writeStreamData(iter.data).
			write("\nendobj\n")
	}
	// write info object
	if pdf.docTitle != "" || pdf.docSubject != "" ||
		pdf.docKeywords != "" || pdf.docAuthor != "" || pdf.docCreator != "" {
		//
		infoIndex = imagesIndex + len(pdf.images)
		pdf.writeObj("/Info")
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
				pdf.write(iter.label).
					write("(").write(string(pdf.escape(iter.field))).write(")")
			}
		}
		pdf.writeEndobj() // info
	}
	// write cross-reference table at end of document
	var startXref = pdf.content.Len()
	pdf.setCurrentPage(PDFNoPage).
		write("xref\n0 %d\n0000000000 65535 f \n", len(pdf.objOffsets))
	for _, offset := range pdf.objOffsets[1:] {
		pdf.write("%010d 00000 n \n", offset)
	}
	// write the trailer
	pdf.write("trailer\n<</Size %d/Root 1 0 R", len(pdf.objOffsets))
	if infoIndex > 0 {
		pdf.write("/Info %d 0 R", infoIndex) // optional reference to info
	}
	pdf.write(">>\nstartxref\n%d\n", startXref).write("%%%%EOF\n")
	return pdf.content.Bytes()
} //                                                                       Bytes

// DrawBox draws a rectangle.
func (pdf *PDF) DrawBox(x, y, width, height float64) *PDF {
	if pdf.warnIfNoPage() {
		return pdf
	}
	width, height = width*pdf.pointsPerUnit, height*pdf.pointsPerUnit
	x *= pdf.pointsPerUnit
	y = pdf.pageSize.HeightPt - y*pdf.pointsPerUnit - height
	pdf.applyLineWidth().applyStrokeColor()
	//
	// 're' = construct a rectangular path, 'S' = stroke path
	return pdf.write("%.3f %.3f %.3f %.3f re S\n", x, y, width, height)
} //                                                                     DrawBox

// DrawImage draws an image.
// For now, only grayscale PNG images are supported.
// 'x' and 'y' specify the position of the image.
// 'height' specifies its height.
// (width is scaled to match the image's aspect ratio).
// fileNameOrBytes is either a string specifying a
// file name, or a byte slice with PNG image data.
func (pdf *PDF) DrawImage(x, y, height float64, fileNameOrBytes interface{},
) *PDF {
	if pdf.warnIfNoPage() {
		return pdf
	}
	var imgName string
	var imgBuf *bytes.Buffer
	switch val := fileNameOrBytes.(type) {
	case string:
		var data, err = ioutil.ReadFile(val)
		if err != nil {
			pdf.logError("File", val, ":", err)
			return pdf
		}
		imgName = val
		imgBuf = bytes.NewBuffer(data)
	case []byte:
		var n = len(val)
		if n > 32 {
			n = 32
		}
		imgName = fmt.Sprintf("%x", val[:n])
		imgBuf = bytes.NewBuffer(val)
	}
	var pg = pdf.pagePtr
	var img pdfImage
	var imgNo = -1
	for i, iter := range pdf.images {
		if iter.name == imgName {
			imgNo, img = i, iter
		}
	}
	if imgNo == -1 {
		var decoded, _, err = image.Decode(imgBuf)
		if err != nil {
			fmt.Printf("image not decoded: %v\n", err)
			return pdf
		}
		var bounds = decoded.Bounds()
		var w, h = bounds.Max.X, bounds.Max.Y
		var data []byte
		for y := 0; y < h; y++ {
			for x := 0; x < w; x++ {
				var rd, gr, bl, _ = decoded.At(x, y).RGBA()
				data = append(data, byte(
					float64(rd)*0.2126+float64(gr)*0.7152+float64(bl)*0.0722,
				))
			}
		}
		img = pdfImage{
			name:      imgName,
			width:     w,
			height:    h,
			data:      data,
			grayscale: true, //TODO: determine actual grayscale mode
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
	x *= pdf.pointsPerUnit
	y = pdf.pageSize.HeightPt - y*pdf.pointsPerUnit - height
	height *= pdf.pointsPerUnit
	var width = float64(img.width) / float64(img.height) * height
	//
	// write command to draw the image
	return pdf.write("q\n"+ // save graphics state
		// w      h x  y
		" %f 0 0 %f %f %f cm\n"+
		"/Im%d Do\n"+
		"Q\n", // restore graphics state
		width, height, x, y, imgNo)
} //                                                                   DrawImage

// DrawLine draws a straight line from point (x1, y1) to point (x2, y2).
func (pdf *PDF) DrawLine(x1, y1, x2, y2 float64) *PDF {
	if pdf.warnIfNoPage() {
		return pdf
	}
	x1, y1 = x1*pdf.pointsPerUnit, pdf.pageSize.HeightPt-y1*pdf.pointsPerUnit
	x2, y2 = x2*pdf.pointsPerUnit, pdf.pageSize.HeightPt-y2*pdf.pointsPerUnit
	pdf.applyLineWidth().applyStrokeColor()
	return pdf.write("%.3f %.3f m %.3f %.3f l S\n", x1, y1, x2, y2)
	// 'm' = move, 'S' = stroke path
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
		pdf.logError("No current page.")
		return pdf
	}
	return pdf.drawTextBox(x, y, width, height, false, align, text)
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
func (pdf *PDF) DrawTextInBox(x, y, width, height float64, align, text string,
) *PDF {
	if pdf.pageNo < 0 {
		pdf.logError("No current page.")
		return pdf
	}
	return pdf.drawTextBox(x, y, width, height, true, align, text)
} //                                                               DrawTextInBox

// DrawUnitGrid draws a light-gray grid demarcated in the current
// measurement unit. The grid fills the entire page.
// It helps with item positioning.
func (pdf *PDF) DrawUnitGrid() *PDF {
	var x, y, pgWidth, pgHeight = 0.0, 0.0, pdf.PageWidth(), pdf.PageHeight()
	if pdf.pageNo < 0 { // ensure there is a current page
		pdf.logError("No current page.")
		return pdf
	}
	pdf.SetLineWidth(0.1).SetFont("Helvetica", 8)
	var xh = 0.1
	var yh = 0.5
	var xv = 0.3
	var yv = 0.3
	var i = 0
	//
	// draw vertical lines
	for x = 0; x < pgWidth; x++ {
		pdf.SetColorRGB(200, 200, 200).DrawLine(x, 0, x, pgHeight)
		if i > 0 {
			pdf.SetColor("Indigo").SetXY(x+xh, y+yh).DrawText(strconv.Itoa(i))
		}
		i++
	}
	// draw horizontal lines
	i = 0
	for y = 0; y < pgHeight; y++ {
		pdf.SetColorRGB(200, 200, 200).DrawLine(0, y, pgWidth, y)
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
	width, height = width*pdf.pointsPerUnit, height*pdf.pointsPerUnit
	x *= pdf.pointsPerUnit
	y = pdf.pageSize.HeightPt - y*pdf.pointsPerUnit - height
	return pdf.applyLineWidth().applyNonStrokeColor().
		write("%.3f %.3f %.3f %.3f re f\n", x, y, width, height)
	// 're' = construct a rectangular path, 'f' = fill
} //                                                                     FillBox

// NextLine advances the text writing position to the next line.
// I.e. the Y increases by the height of the font and
// the X-coordinate is reset to zero.
func (pdf *PDF) NextLine() *PDF {
	var y = pdf.Y() + pdf.FontSize()*pdf.pointsPerUnit
	if y > pdf.pageSize.HeightPt*pdf.pointsPerUnit {
		pdf.AddPage()
		y = 0
	}
	pdf.columnNo = 0
	var x = 0.0
	if len(pdf.columnWidths) > 0 {
		x = pdf.columnWidths[0]
	}
	return pdf.SetXY(x, y)
} //                                                                    NextLine

// Reset releases all resources and resets all variables.
func (pdf *PDF) Reset() *PDF {
	for _, page := range pdf.pages {
		page.fontIDs = []int{}
		page.imageNos = []int{}
		page.pageContent.Reset()
	}
	pdf.docAuthor, pdf.docCreator, pdf.docKeywords = "", "", ""
	pdf.docSubject, pdf.docTitle, pdf.unitName = "", "", ""
	pdf.pages = []pdfPage{}
	pdf.fonts = []pdfFont{}
	pdf.images = []pdfImage{}
	pdf.columnWidths = []float64{}
	pdf.content.Reset()
	return pdf
} //                                                                       Reset

// SaveFile generates and saves the PDF document to a file.
func (pdf *PDF) SaveFile(filename string) *PDF {
	var err = ioutil.WriteFile(filename, pdf.Bytes(), 0644)
	if err != nil {
		pdf.logError("Failed writing to file", filename, ":", err)
	}
	return pdf
} //                                                                    SaveFile

// SetColumnWidths creates columns along the X-axis.
func (pdf *PDF) SetColumnWidths(widths ...float64) *PDF {
	if len(widths) < 1 {
		pdf.logError("len(widths) < 1")
		return pdf
	}
	if len(widths) > 100 {
		pdf.logError("len(widths) > 100")
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
	//
	// first handle the newlines...
	var ar = pdf.splitLines(text)
	//
	// ...then break lines based on text width
	var ret []string
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
			for ln > 0 && ln < max && pdf.TextWidth(iter[:ln]) <= width {
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
			pdf.logError("Unknown unit name.")
		}
		ret *= ppu
	}
	return ret
} //                                                                    ToPoints

// -----------------------------------------------------------------------------
// # Private Methods

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
// What this method does:
// - Validates the current font name and determines if it is a
//   standard (built-in) font like Helvetica or a TrueType font.
// - Fills the document-wide list of fonts (pdf.fonts).
// - Adds items to the list of font ID's used on the current page.
//
// called by drawTextLine()
func (pdf *PDF) applyFont() {
	var isValid = pdf.fontName != ""
	var font pdfFont
	if isValid {
		isValid = false
		for i, name := range pdfFontNames {
			name = strings.ToUpper(name)
			if strings.ToUpper(pdf.fontName) != name {
				continue
			}
			font.fontName = pdfFontNames[i]
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
	pdf.write("BT /F%d %d Tf ET\n", pg.fontID, int(pg.fontSizePt))
	// 'BT' = __, F = __, 'Tf' = __, 'ET' = __
} //                                                                   applyFont

// applyLineWidth writes a line-width PDF command ('w') to the current
// page's content stream, if the value changed since last call.
// called by: DrawBox(), DrawLine(), FillBox()
func (pdf *PDF) applyLineWidth() *PDF {
	var val = &pdf.pagePtr.lineWidth
	if int(*val*10000) != int(pdf.lineWidth*10000) {
		*val = pdf.lineWidth
		pdf.write("%.3f w\n", float64(*val))
		// 'w' = __
	}
	return pdf
} //                                                              applyLineWidth

// applyNonStrokeColor writes a text color selection PDF command ('rg') to
// the current page's content stream, if the value changed since last call.
// called by: drawTextLine(), FillBox()
func (pdf *PDF) applyNonStrokeColor() *PDF {
	var val = &pdf.pagePtr.nonStrokeColor
	if *val == pdf.color {
		return pdf
	}
	*val = pdf.color
	pdf.write("%.3f %.3f %.3f rg\n", // 'rg' = set non-stroking (text) color
		float64(val.Red)/255,
		float64(val.Green)/255,
		float64(val.Blue)/255)
	return pdf
} //                                                         applyNonStrokeColor

// applyStrokeColor writes a line color selection PDF
// command ('RG') to the current page's content
// stream, if the value changed since last call.
// called by: DrawBox(), DrawLine()
func (pdf *PDF) applyStrokeColor() *PDF {
	var val = &pdf.pagePtr.strokeColor
	if *val == pdf.color {
		return pdf
	}
	*val = pdf.color
	pdf.write("%.3f %.3f %.3f RG\n", // 'RG' = stroke (line) color
		float64(val.Red)/255,
		float64(val.Green)/255,
		float64(val.Blue)/255)
	return pdf
} //                                                            applyStrokeColor

// drawTextLine writes a line of text at the current coordinates to the
// current page's content stream, using a sequence of raw PDF commands.
// called by: DrawText(), drawTextBox()
func (pdf *PDF) drawTextLine(text string) *PDF {
	if text == "" {
		return pdf
	}
	// draw the text:
	pdf.applyFont()
	if pdf.pagePtr.horizontalScaling != pdf.horizontalScaling {
		pdf.pagePtr.horizontalScaling = pdf.horizontalScaling
		pdf.write("BT %d Tz ET\n", pdf.pagePtr.horizontalScaling)
		// 'BT' = __ 'Tz' = __ 'ET' = __
	}
	pdf.applyNonStrokeColor()
	var pg = pdf.pagePtr
	if pg.x < 0 || pg.y < 0 {
		pdf.SetXY(0, 0)
	}
	// 'BT' = __, 'Td' = __, 'Tj' = __, 'ET' = __
	pdf.write("BT %d %d Td (%s) Tj ET\n",
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
		pdf.logError("No current page.")
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
		for _, ch := range align {
			if ch == 'l' || ch == 'L' {
				return pdf.fontSizePt / 6 // add margin
			} else if ch == 'r' || ch == 'R' {
				return width - pdf.textWidthPt1000(text) - pdf.fontSizePt/6
			}
		}
		return width/2 - pdf.textWidthPt1000(text)/2 // center
	}
	// calculate aligned y-axis position of text (top, bottom, center)
	y, height = y*pdf.pointsPerUnit, height*pdf.pointsPerUnit
	y += pdf.fontSizePt + (func() float64 {
		for _, ch := range align {
			if ch == 't' || ch == 'T' {
				return 0 // top
			} else if ch == 'b' || ch == 'B' {
				return height - allLinesHeight - 4 // bottom
			}
		}
		return height/2 - allLinesHeight/2 - pdf.fontSizePt*0.15 // center
	})() // IIFE
	y = pdf.pageSize.HeightPt - y
	for _, line := range lines {
		pdf.pagePtr.x, pdf.pagePtr.y = x+alignX(line), y
		pdf.drawTextLine(line)
		y -= lineHeight
	}
	return pdf
} //                                                                 drawTextBox

// setCurrentPage selects the currently-active page.
// called by: AddPage(), Bytes()
func (pdf *PDF) setCurrentPage(pageNo int) *PDF {
	if pageNo != PDFNoPage && pageNo > (len(pdf.pages)-1) {
		pdf.logError("Page number out of range.")
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
		pdf.logError("No current page.")
		return 0
	}
	if text == "" {
		return 0
	}
	var w = 0.0
	for i, ch := range text {
		if ch < 0 || ch > 255 {
			pdf.logError("char out of range at %d: %d", i, ch)
			break
		}
		w += float64(pdfFontWidths[ch][0])
		// TODO: [0] is not considering the current font!
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
		pdf.logError("No current page.")
		return true
	}
	return false
} //                                                                warnIfNoPage

// -----------------------------------------------------------------------------
// # Private Generation Methods

// nextObj increases the object serial no. and stores its offset in array
func (pdf *PDF) nextObj() int {
	pdf.objNo++
	for len(pdf.objOffsets) <= pdf.objNo {
		pdf.objOffsets = append(pdf.objOffsets, pdf.content.Len())
	}
	return pdf.objNo
} //                                                                     nextObj

// write writes formatted strings (like fmt.Sprintf) to the current page's
// content stream or to the final generated PDF, if there is no active page.
// called by: Bytes(), DrawBox(), DrawLine(), drawTextLine(), FillBox()
//            applyLineWidth(), applyNonStrokeColor(), applyStrokeColor()
func (pdf *PDF) write(format string, args ...interface{}) *PDF {
	var buf *bytes.Buffer
	if pdf.pageNo == PDFNoPage {
		buf = pdf.contentPtr
	} else if pdf.pageNo > (len(pdf.pages) - 1) {
		pdf.logError("Invalid page number.")
		return pdf
	} else {
		buf = &pdf.pagePtr.pageContent
	}
	buf.Write([]byte(fmt.Sprintf(format, args...)))
	return pdf
} //                                                                       write

// writeEndobj writes 'endobj' (PDF object end marker).
func (pdf *PDF) writeEndobj() *PDF {
	return pdf.write(">>\nendobj\n")
} //                                                                 writeEndobj

// writeObj outputs an object header
func (pdf *PDF) writeObj(objectType string) *PDF {
	pdf.setCurrentPage(PDFNoPage)
	var objNo = pdf.nextObj()
	if objectType == "" {
		pdf.write("%d 0 obj<<", objNo)
	} else if objectType[0] == '/' {
		pdf.write("%d 0 obj<</Type%s", objNo, objectType)
	} else {
		pdf.logError("objectType should begin with '/' or be a blank string.")
	}
	return pdf
} //                                                                    writeObj

// writePages writes all PDF pages.
func (pdf *PDF) writePages(fontsIndex, imagesIndex int) *PDF {
	if len(pdf.pages) > 0 { //                                write page numbers
		var pageObjectNo = pdfPagesIndex
		pdf.write("/Kids[")
		for i := range pdf.pages {
			if i > 0 {
				pdf.write(" ")
			}
			pdf.write("%d 0 R", pageObjectNo)
			pageObjectNo += 2 // 1 for page, 1 for stream
		}
		pdf.write("]")
	}
	pdf.writeEndobj()
	for _, pg := range pdf.pages { //                            write each page
		if pg.pageContent.Len() == 0 {
			fmt.Printf("WARNING: Empty Page\n")
			continue
		}
		pdf.writeObj("/Page").
			write("/Parent 2 0 R/Contents %d 0 R", pdf.objNo+1)
		if len(pg.fontIDs) > 0 || len(pg.imageNos) > 0 {
			pdf.write("/Resources<<")
		}
		if len(pg.fontIDs) > 0 {
			pdf.write("/Font <<")
			for fontNo := range pdf.fonts {
				pdf.write("/F%d %d 0 R", fontNo+1, fontsIndex+fontNo)
			}
			pdf.write(">>")
		}
		if len(pg.imageNos) > 0 {
			pdf.write("/XObject<<")
			for imageNo := range pg.imageNos {
				pdf.write("/Im%d %d 0 R", imageNo, imagesIndex+imageNo)
			}
			pdf.write(">>")
		}
		if len(pg.fontIDs) > 0 || len(pg.imageNos) > 0 {
			pdf.write(">>")
		}
		pdf.writeEndobj().writeStream([]byte(pg.pageContent.String()))
	}
	return pdf
} //                                                                  writePages

// writeStream outputs a stream object to the document's main buffer.
func (pdf *PDF) writeStream(content []byte) *PDF {
	return pdf.setCurrentPage(PDFNoPage).
		write("%d 0 obj <<", pdf.nextObj()).writeStreamData(content)
} //                                                                 writeStream

// writeStreamData writes a stream or image stream.
func (pdf *PDF) writeStreamData(content []byte) *PDF {
	pdf.setCurrentPage(PDFNoPage)
	var filter string
	if pdf.compressStreams {
		filter = "/Filter/FlateDecode"
		var buf bytes.Buffer
		var writer = zlib.NewWriter(&buf)
		var _, err = writer.Write([]byte(content))
		if err != nil {
			pdf.logError("Failed compressing:", err)
			return pdf
		}
		writer.Close() // don't use defer, close immediately
		content = buf.Bytes()
	}
	return pdf.write("%s/Length %d>>stream\n%s\nendstream\n",
		filter, len(content), content)
} //                                                             writeStreamData

// -----------------------------------------------------------------------------
// # Private Functions

// escape escapes special characters '(', '(' and '\' in strings
// in order to avoid them interfering with PDF commands.
// called by: Bytes(), drawTextLine()
func (*PDF) escape(s string) []byte {
	if strings.Contains(s, "(") || strings.Contains(s, ")") ||
		strings.Contains(s, "\\") {
		//
		var writer = bytes.NewBuffer(make([]byte, 0, len(s)))
		for _, ch := range s {
			if ch == '(' || ch == ')' || ch == '\\' {
				writer.WriteRune('\\')
			}
			writer.WriteRune(ch)
		}
		return writer.Bytes()
	}
	return []byte(s)
} //                                                                      escape

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

// isWhiteSpace returns true if all the chars. in 's' are white-spaces.
func (*PDF) isWhiteSpace(s string) bool {
	if s == "" {
		return false
	}
	for _, ch := range s {
		if ch != ' ' && ch != '\a' && ch != '\b' && ch != '\f' &&
			ch != '\n' && ch != '\r' && ch != '\t' && ch != '\v' {
			return false
		}
	}
	return true
} //                                                                isWhiteSpace

// splitLines splits 's' into several lines using line breaks in 's'.
func (*PDF) splitLines(s string) []string {
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
	return split(split(split([]string{s}, "\r\n"), "\r"), "\n")
} //                                                                  splitLines

// logError reports an error by calling PDFErrorHandler,
// which is set to fmt.Println by default.
func (PDF) logError(a ...interface{}) {
	if PDFErrorHandler != nil {
		PDFErrorHandler(a...)
	}
} //                                                                    logError

//end
