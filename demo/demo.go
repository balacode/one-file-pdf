// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-03-16 16:21:46 44A300                                      [demo.go]
// -----------------------------------------------------------------------------

package main

// This demo generates the sample PDF files
// and demonstrates the API of One-File-PDF

import "fmt"     // standard
import "strings" // standard

import "github.com/balacode/one-file-pdf"

func main() {
	helloWorld()
	corporateIpsum()
	pngImages()
} //                                                                        main

func helloWorld() {
	fmt.Println(`Generating a "Hello World" PDF...`)
	//
	// create a new PDF using 'A4' page size
	var pdf = pdf.NewPDF("A4")
	//
	// set the measurement units to centimeters
	pdf.SetUnits("cm")
	//
	// draw a grid to help us align stuff (just a guide, not necessary)
	pdf.DrawUnitGrid()
	//
	// draw the word 'HELLO' in orange, using 100pt bold Helvetica font
	// - text is placed on top of, not below the Y-coordinate
	// - you can use method chaining
	pdf.SetFont("Helvetica-Bold", 100).
		SetXY(5, 5).
		SetColor("Orange").
		DrawText("HELLO")
	//
	// draw the word 'WORLD' in blue-violet, using 100pt Helvetica font
	// note that here we use the colo(u)r hex code instead
	// of its name, using the CSS/HTML format: #RRGGBB
	pdf.SetXY(5, 9).
		SetColor("#8A2BE2").
		SetFont("Helvetica", 100).
		DrawText("WORLD!")
	//
	// draw a flower icon using 300pt Zapf-Dingbats font
	pdf.SetXY(7, 17).
		SetColor("Red").
		SetFont("ZapfDingbats", 300).
		DrawText("a")
	//
	// save the file:
	// if the file exists, it will be overwritten
	// if the file is in use, prints an error message
	pdf.SaveFile("hello.pdf")
} //                                                                  helloWorld

func corporateIpsum() {
	const FILENAME = "corporate.pdf"
	fmt.Println("Generating sample PDF:", FILENAME, "...")
	var pdf = pdf.NewPDF("A4") // create a new PDF using 'A4' page size
	pdf.SetUnits("cm")
	//
	// draw the heading
	pdf.SetColor("#002FA7 InternationalKleinBlue").
		FillBox(0, 1.5, 21, 1.5).
		SetFont("Helvetica-Bold", 50).
		SetColor("White").
		SetXY(3.5, 2.7).DrawText("Synergy Ipsum")
	//
	// draw the green circle
	pdf.SetColor("#74C365 Mantis").FillCircle(21, 21, 10) // x, y, radius
	//
	// draw the left column of text (in a box)
	var col1 = strings.Replace(CorporateFiller1, "\n", " ", -1)
	pdf.SetColor("#73C2FB MayaBlue").
		FillBox(0, 4, 10, 15). // xywh
		SetColor("black").
		SetFont("times-roman", 11).
		DrawTextInBox(0.5, 4.5, 9, 15, "LT", col1)
	//
	// draw the right column of text
	var col2 = strings.Replace(CorporateFiller2, "\n", " ", -1)
	pdf.SetColor("black").
		SetFont("Times-Italic", 11).
		DrawTextInBox(10.5, 4, 9, 28, "LT", col2)
	//
	// draw the bottom-left box with a checkmark
	pdf.SetColor("#EAA221 Marigold").
		FillBox(0, 25, 5, 5). // xywh
		SetFont("zapfdingbats", 50).
		SetColor("white").
		DrawTextInBox(0, 25, 5, 5, "C", string(rune(063)))
	//
	// save the file
	pdf.SaveFile(FILENAME)
} //                                                              corporateIpsum

func pngImages() {
	const FILENAME = "png_images.pdf"
	fmt.Println("Generating sample PDF:", FILENAME, "...")
	var pdf = pdf.NewPDF("A4").SetUnits("cm")
	//
	// draw background pattern
	for x := 0.0; x < pdf.PageWidth(); x += 6 {
		for y := 0.0; y < pdf.PageHeight(); y += 5 {
			pdf.DrawImage(x, y, 5, "../image/gophers.png", "cyan")
		}
	}
	// draw dice
	pdf.SetColor("WHITE")
	pdf.FillBox(3.5, 4.5, 14.7, 17)
	//
	pdf.DrawImage(4, 5, 5, "../image/dice.png", "WHITE")
	pdf.DrawImage(11, 5, 5, "../image/dice.png", "RED")
	//
	pdf.DrawImage(4, 10.5, 5, "../image/dice.png", "GREEN")
	pdf.DrawImage(11, 10.5, 5, "../image/dice.png", "BLUE")
	//
	pdf.DrawImage(4, 16, 5, "../image/dice.png", "BLACK")

	pdf.SetFont("Helvetica-Bold", 50)
	pdf.SetXY(3, 3).SetColor("#009150").
		DrawText("PNG Image Demo")

	//pdf.DrawUnitGrid()
	pdf.SaveFile(FILENAME)
} //                                                                   pngImages

//end
