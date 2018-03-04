// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-03-04 02:49:22 AC75C8                                      [demo.go]
// -----------------------------------------------------------------------------

package main

// This demo displays

import "fmt" // standard

import "github.com/balacode/one-file-pdf"

func main() {
	fmt.Println(`Generating a "Hello World" PDF...`)
	//
	// create a new PDF using 'A4' page size
	var pdf = pdf.NewPDF("A4")
	//
	// add a page: this must be done before drawing
	pdf.AddPage()
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
} //                                                                        main

//end
