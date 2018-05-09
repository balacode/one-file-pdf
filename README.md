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

- It's easier to learn about the internals of the PDF format with a small, concise library.
- The current version of the file is indicated in the header (the timestamp).

## Features:  
- The essentials for generating PDF documents, sufficient for common business reports.
- Use all built-in PDF fonts: Courier, Helvetica, Symbol, Times, ZapfDingbats, and their variants
- Specify colo(u)rs by name (144 web colors), HTML codes (#RRGGBB) or RGB value
- Set columns for text (like tab stops on the page)
- Built-in grid option to help measurement and positioning
- Metadata properties: author, creator, keywords, subject and title
- Set the measurement units you want: mm, cm, inches, twips or points
- Draw lines with different thickness
- Filled or outline rectangles, circles and ellipses
- JPEG, GIF and transparent PNG images (filled with specified background color)
- Stream compression can be turned on or off (PDF files normally compress streams to reduce file size, but turning it off helps in debugging or learning about PDF commands)

## Not Yet Supported:  
- Unicode (requires font embedding)
- Font embedding
- PDF encryption
- Paths, curves and complex graphics

## Installation:  

```bash
    go get github.com/balacode/one-file-pdf
```

## Naming Convention:  
All types in are prefixed with PDF for public, and 'pdf' for private types.
The only type you need to use is PDF, while PDFColorNames are left public for reference.

## Hello World:  

```go
package main 

import (
	"fmt"
	"github.com/balacode/one-file-pdf"
)

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

## Samples:
Click on a sample to see the PDF in more detail.

[!["Hello World!" sample image](demo/samples/hello.png)](demo/samples/hello.pdf)  

[!["Synergy Ipsum" sample image](demo/samples/corporate.png)](demo/samples/corporate.pdf)  

## Changelog:  

These are the most recent changes in the functionality of the package,
not including internal changes which are best seen in the commits history.

**2018-04-14**
- Changed CurrentPage from read-only to read/write property: added SetCurrentPage()
- Created PageCount() read-only property
- Created dingbats() demo to generate `zapf_dingbats_table.pdf`.
  You can use this table to look up the hex code for each icon.
- Changed text encoding from /WinAnsiEncoding to /StandardEncoding

See [changelog.md](./doc/changelog.md) for changes made earlier.

## Roadmap:  

- Achieve 100% test coverage
- Create a unit test for every method
- Unicode support
- Partial font embedding
