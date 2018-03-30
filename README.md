## one-file-pdf - A minimalist PDF generator in &lt;2K lines and 1 file
[![Go Report Card](https://goreportcard.com/badge/github.com/balacode/one-file-pdf)](https://goreportcard.com/report/github.com/balacode/one-file-pdf)
[![Build Status](https://travis-ci.org/balacode/one-file-pdf.svg?branch=master)](https://travis-ci.org/balacode/one-file-pdf)
[![Test Coverage](https://coveralls.io/repos/github/balacode/one-file-pdf/badge.svg?branch=master&service=github)](https://coveralls.io/github/balacode/one-file-pdf?branch=master)
[![Gitter chat](https://badges.gitter.im/balacode/one-file-pdf.png)](https://gitter.im/one-file-pdf/Lobby)
[![godoc](https://godoc.org/github.com/balacode/one-file-pdf?status.svg)](https://godoc.org/github.com/balacode/one-file-pdf)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)    

The main idea behind this project was:  
*"How small can I make a PDF generator for it to still be useful for 80% of common PDF generation needs?"*

The result is a single .go file with less than 1999 lines of code, about 400 of which are color and glyph-size constants, and ~350 are comments.

- All the basics for generating PDF documents, enough for generating mundane business reports.
- It's easier to learn about the internals of the PDF format with a small, concise library.
- The current version of the file is indicated in the header (the timestamp).

### To Install:  

```bash
    go get github.com/balacode/one-file-pdf
```

### Naming Convention:  
All types in are prefixed with PDF for public, and 'pdf' for private types.
The only type you need to use is PDF, while PDFColorNames are left public for reference.

### Features:  
- Use all built-in PDF fonts: Courier, Helvetica, Symbol, Times, ZapfDingbats, and their variants
- Recognises 144 web colo(u)r names, or any RGB value
- Stream compression can be turned on or off (PDF files normally compress streams to reduce file size, but turning it off helps in debugging or learning about PDF commands)
- Metadata properties: author, creator, keywords, subject and title
- Set the measurement units you want: mm, cm, inches, twips or points
- Draw lines with different thickness
- Filled or outline rectangles, circles and ellipses
- JPEG, GIF and transparent PNG images (filled with specified background color)
- Set columns for text (like tab stops on the page)
- Built-in grid option to help measurement and positioning

### Not Supported:  
- Unicode (requires font embedding)
- Font embedding
- PDF encryption
- Paths, curves and complex graphics

### Hello World:  

```go
package main 

import "fmt"
import "github.com/balacode/one-file-pdf"

func main() {
	fmt.Println(`Generating a "Hello World" PDF...`)

	// create a new PDF using 'A4' page size
	var pdf = pdf.NewPDF("A4")

	// set the measurement units to centimeters
	pdf.SetUnits("cm")

	// draw a grid to help us align stuff (just a guide, not necessary)
	pdf.DrawUnitGrid()

	// draw the word 'HELLO' in orange, using 100pt bold Helvetica font
	// - text is placed on top of, not below the Y-coordinate
	// - you can use method chaining
	pdf.SetFont("Helvetica-Bold", 100).
		SetXY(5, 5).
		SetColor("Orange").
		DrawText("HELLO")

	// draw the word 'WORLD' in blue-violet, using 100pt Helvetica font
	// note that here we use the colo(u)r hex code instead
	// of its name, using the CSS/HTML format: #RRGGBB
	pdf.SetXY(5, 9).
		SetColor("#8A2BE2").
		SetFont("Helvetica", 100).
		DrawText("WORLD!")

	// draw a flower icon using 300pt Zapf-Dingbats font
	pdf.SetX(7).SetY(17).
		SetColorRGB(255, 0, 0).
		SetFont("ZapfDingbats", 300).
		DrawText("a")

	// save the file:
	// if the file exists, it will be overwritten
	// if the file is in use, prints an error message
	pdf.SaveFile("hello.pdf")
} //                                                                        main
```

### Samples:
Click on a sample to see the PDF in more detail.

[!["Hello World!" sample image](demo/samples/hello.png)](demo/samples/hello.pdf)  

[!["Synergy Ipsum" sample image](demo/samples/corporate.png)](demo/samples/corporate.pdf)  

### Changelog (excluding internal changes):

**ALTERED API:**
- **Removed SetErrorLogger() method**
- **ToPoints(): added error return value**  

- Initialize PDF automatically, even when NewPDF() wasn't called. The paper size
  is A4, and the units CM by default. You have to use NewPDF() to specify a paper size.
- No need to add the first page with AddPage(). It is inserted automatically.
- New error handling methods Clean(), Errors(), ErrorInfo() and PullError().
- SetColumnWidths(): can be called without arguments when you need to reset all columns.
- Added various unit tests, for 95% code coverage.
- Fixed text wrapping bug that could cause PDF to freeze.

See [changelog.md](./doc/changelog.md) for changes made earlier.
