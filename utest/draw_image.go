// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-05-13 01:54:23 9D8232             one-file-pdf/utest/[draw_image.go]
// -----------------------------------------------------------------------------

package utest

import (
	"fmt"
	"testing"

	"github.com/balacode/one-file-pdf"
)

// Test_PDF_DrawImage_ is the unit test for
// PDF.DrawImage(x, y, height float64, fileNameOrBytes interface{},
//     backColor ...string) *PDF
//
// Runs the test by drawing rgbw64.png:
// a small 64 x 64 PNG split into pure red, green, blue
// and transparent gradient squares
func Test_PDF_DrawImage_(t *testing.T) {
	fmt.Println("Test PDF.DrawImage()")
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

	const expectOpaque = `
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
		var doc = pdf.NewPDF("20cm x 20cm")
		doc.SetCompression(true).
			SetUnits("cm").
			DrawImage(x, y, height, pngData)
		FailIfHasErrors(t, doc.Errors)
		ComparePDF(t, doc.Bytes(), expectOpaque)
	}()

	// the same test, but reading direcly from PNG file
	func() {
		var doc = pdf.NewPDF("20cm x 20cm")
		doc.SetCompression(true).
			SetUnits("cm").
			DrawImage(x, y, height, "./image/rgbw64.png")
		FailIfHasErrors(t, doc.Errors)
		ComparePDF(t, doc.Bytes(), expectOpaque)
	}()

	// PNG transparency test
	func() {
		var x = 5.0
		var y = 5.0
		var height = 5.0

		const expectTransparent = `
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

		var doc = pdf.NewPDF("20cm x 20cm")
		doc.SetCompression(true).
			SetUnits("cm").
			DrawImage(x, y, height, "./image/rgbt64.png", "Yellow").
			DrawImage(x, y+height+1, height, "./image/rgbt64.png", "Cyan")
		FailIfHasErrors(t, doc.Errors)
		ComparePDF(t, doc.Bytes(), expectTransparent)
	}()

	// wrong argument in fileNameOrBytes
	func() {
		var doc = pdf.NewPDF("20cm x 20cm")
		var fileNameOrBytes = []int{0xBAD, 0xBAD, 0xBAD}

		const expect = `
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

		doc.SetCompression(true).
			SetUnits("cm").
			DrawImage(x, y, height, fileNameOrBytes)
		ComparePDF(t, doc.Bytes(), expect)
		//
		TEqual(t, len(doc.Errors()), 1)
		if len(doc.Errors()) > 0 {
			TEqual(t, doc.Errors()[0], fmt.Errorf(
				`Invalid type in fileNameOrBytes`+
					` "[]int = [2989 2989 2989]" @DrawImage`))
		}
	}()

	// drawing an image on one page and another image on second page should work
	// (catch bug where image name in '/XObject' and '/IMG Do' do not match)
	func() {

		const expect = `
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
		var pngData1 = append(pngData, 1)
		var pngData2 = append(pngData, 2)
		//
		var doc = pdf.NewPDF("20cm x 20cm")
		doc.SetCompression(false).
			SetUnits("cm").
			DrawImage(x, y, height, pngData1).
			AddPage().
			DrawImage(x, y, height, pngData2)
		FailIfHasErrors(t, doc.Errors)
		ComparePDF(t, doc.Bytes(), expect)
	}()

} //                                                         Test_PDF_DrawImage_

//end
