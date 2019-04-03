// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2019-04-03 10:52:12 05EDE7                    one-file-pdf/demo/[demo.go]
// -----------------------------------------------------------------------------

package main

// This demo generates the sample PDF files
// and demonstrates the API of PDF

import (
	"fmt"
	str "strings"

	"github.com/balacode/one-file-pdf"
)

func main() {
	/*1*/ helloWorld()
	/*2*/ corporateIpsum()
	/*3*/ pngImages()
	/*4*/ dingbats()
} //                                                                        main

func helloWorld() {
	fmt.Println(`Generating a "Hello World" PDF...`)
	//
	// create a new PDF using 'A4' page size
	doc := pdf.NewPDF("A4")
	//
	// you can call AddPage() to add the first page, but it is
	// not required: the first page is added automatically
	// doc.AddPage()
	//
	// set the measurement units to centimeters
	doc.SetUnits("cm")
	//
	// draw a grid to help us align stuff (just a guide, not necessary)
	doc.DrawUnitGrid()
	//
	// draw the word 'HELLO' in orange, using 100pt bold Helvetica font
	// - text is placed on top of, not below the Y-coordinate
	// - you can use method chaining
	doc.SetFont("Helvetica-Bold", 100).
		SetXY(5, 5).
		SetColor("Orange").
		DrawText("HELLO")
	//
	// draw the word 'WORLD' in blue-violet, using 100pt Helvetica font
	// note that here we use the colo(u)r hex code instead
	// of its name, using the CSS/HTML format: #RRGGBB
	doc.SetXY(5, 9).
		SetColor("#8A2BE2").
		SetFont("Helvetica", 100).
		DrawText("WORLD!")
	//
	// draw a flower icon using 300pt Zapf-Dingbats font
	doc.SetXY(7, 17).
		SetColor("Red").
		SetFont("ZapfDingbats", 300).
		DrawText("a")
	//
	// save the file:
	// if the file exists, it will be overwritten
	// if the file is in use, prints an error message
	doc.SaveFile("hello.pdf")
} //                                                                  helloWorld

func corporateIpsum() {
	const FILENAME = "corporate.pdf"
	fmt.Println("Generating sample PDF:", FILENAME, "...")
	doc := pdf.NewPDF("A4") // create a new PDF using 'A4' page size
	doc.SetUnits("cm")
	//
	// draw the heading
	doc.SetColor("#002FA7 InternationalKleinBlue").
		FillBox(0, 1.5, 21, 1.5).
		SetFont("Helvetica-Bold", 50).
		SetColor("White").
		SetXY(3.5, 2.7).DrawText("Synergy Ipsum")
	//
	// draw the green circle
	doc.SetColor("#74C365 Mantis").FillCircle(21, 21, 10) // x, y, radius
	//
	// draw the left column of text (in a box)
	col1 := str.Replace(CorporateFiller1, "\n", " ", -1)
	doc.SetColor("#73C2FB MayaBlue").
		FillBox(0, 4, 10, 15). // xywh
		SetColor("black").
		SetFont("times-roman", 11).
		DrawTextInBox(0.5, 4.5, 9, 15, "LT", col1)
	//
	// draw the right column of text
	col2 := str.Replace(CorporateFiller2, "\n", " ", -1)
	doc.SetColor("black").
		SetFont("Times-Italic", 11).
		DrawTextInBox(10.5, 4, 9, 28, "LT", col2)
	//
	// draw the bottom-left box with a checkmark
	doc.SetColor("#EAA221 Marigold").
		FillBox(0, 25, 5, 5). // xywh
		SetFont("zapfdingbats", 50).
		SetColor("white").
		DrawTextInBox(0, 25, 5, 5, "C", string(rune(063)))
	//
	// save the file
	doc.SaveFile(FILENAME)
} //                                                              corporateIpsum

func pngImages() {
	const FILENAME = "png_images.pdf"
	fmt.Println("Generating sample PDF:", FILENAME, "...")
	doc := pdf.NewPDF("A4")
	doc.SetUnits("cm")
	//
	// draw background pattern
	for x := 0.0; x < doc.PageWidth(); x += 6 {
		for y := 0.0; y < doc.PageHeight(); y += 5 {
			doc.DrawImage(x, y, 5, "../image/gophers.png", "cyan")
		}
	}
	// draw dice
	doc.SetColor("WHITE").FillBox(3.5, 4.5, 14.7, 17).
		//
		DrawImage(4, 5, 5, "../image/dice.png", "WHITE").
		DrawImage(11, 5, 5, "../image/dice.png", "RED").
		//
		DrawImage(4, 10.5, 5, "../image/dice.png", "GREEN").
		DrawImage(11, 10.5, 5, "../image/dice.png", "BLUE").
		//
		DrawImage(4, 16, 5, "../image/dice.png", "BLACK").
		SetFont("Helvetica-Bold", 50).
		SetXY(3, 3).SetColor("#009150").
		DrawText("PNG Image Demo")
	//
	doc.SaveFile(FILENAME)
} //                                                                   pngImages

// dingbats generates a useful table of icon codes for Zapf-Dingbats
func dingbats() {
	const filename = "zapf_dingbats_table.pdf"
	fmt.Println("Generating", filename, "...")
	//
	// create a new PDF using 'A4' page size
	doc := pdf.NewPDF("A4")
	doc.SetUnits("cm")
	//
	const boxSize = 1.2 // cm
	x, y := 1.0, 1.0    // cm
	//
	doc.SetFont("Helvetica-Bold", 100)
	doc.SetLineWidth(0.02)
	//
	for row := 0; row < 16; row++ {
		x = 1.0
		for col := 0; col < 16; col++ {
			//
			// draw border around each icon
			doc.SetColor("gray")
			doc.DrawBox(x, y, boxSize, boxSize)
			//
			// draw hex code
			doc.SetColor("dark green")
			doc.SetFont("Helvetica", 7)
			doc.DrawTextInBox(x+0.1, y, boxSize, boxSize, "TL",
				fmt.Sprintf("%02Xh", row*16+col))
			//
			// draw decimmal code
			doc.SetColor("dark violet")
			doc.SetFont("Helvetica", 7)
			doc.DrawTextInBox(x, y, boxSize, boxSize, "TR",
				fmt.Sprintf("%d", row*16+col))
			//
			// this is the right way to use a dingbat code (0-255):
			// (casting int to rune to string won't work as expected)
			code := row*16 + col
			s := string([]byte{byte(code)})
			//
			// draw the dingbat icon
			doc.SetColor("black")
			doc.SetFont("ZapfDingbats", 20)
			doc.DrawTextInBox(x, y+0.1, boxSize, boxSize, "C", s)
			//
			x += boxSize
		}
		y += boxSize
	}
	doc.SaveFile(filename)
} //                                                                    dingbats

//end
