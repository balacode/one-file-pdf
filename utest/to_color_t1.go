// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-04-29 23:42:24 E4B30C            one-file-pdf/utest/[to_color_t1.go]
// -----------------------------------------------------------------------------

package utest

import (
	"fmt"
	"image/color"
	"testing"

	"github.com/balacode/one-file-pdf"
)

// Test_PDF_ToColor_1_ is the unit test for
// (ob *PDF) ToColor(nameOrHTMLColor string) (color.RGBA, error)
func Test_PDF_ToColor_1_(t *testing.T) {
	fmt.Println("Test PDF.ToColor() [1]")

	func() {
		var doc pdf.PDF
		var got, err = doc.ToColor("")
		TEqual(t, got, color.RGBA{A: 255}) // black
		// error is returned in `err`, but does not affect Errors()
		TEqual(t, len(doc.Errors()), 0)
		TEqual(t, err.Error(),
			fmt.Errorf(`Unknown color name "" @ToColor`).Error())
	}()

	func() {
		var doc pdf.PDF
		var got, err = doc.ToColor("#uvwxyz")
		TEqual(t, got, color.RGBA{A: 255}) // black
		// error is returned in `err`, but does not affect Errors()
		TEqual(t, len(doc.Errors()), 0)
		TEqual(t, err.Error(),
			fmt.Errorf(`Bad color code "#uvwxyz" @ToColor`).Error())
	}()

	func() {
		// map copied from PDFColorNames, but color names in lower case
		var m = map[string]color.RGBA{
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
		var doc pdf.PDF
		for key, val := range m {
			val.A = 255 // make opaque, not transparent
			var color, err = doc.ToColor(key)
			TEqual(t, color, val)
			TEqual(t, err, nil)
		}
	}()

} //                                                         Test_PDF_ToColor_1_

//end
