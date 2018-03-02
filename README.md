## one-file-pdf - A minimalist PDF generator in &lt;1999 lines and 1 file
[![Go Report Card](https://goreportcard.com/badge/github.com/balacode/one-file-pdf)](https://goreportcard.com/report/github.com/balacode/one-file-pdf)  

The main idea behind this project was: "How small can I make a PDF generator for it to still be useful for 80% of common PDF generation needs?"

- This file has all the basics for generating PDF documents, enough for generating mundane business reports. 
- It's easier to learn about the internals of the PDF format with a smaller libary.
- You can just drop the one_file_pdf.go file in your Go project. No need to manage dependencies. (But remember to change the package to 'main' or whatever package you are adding it to.)
- The current version of the file is already visible in the header.

**Supported Features (i.e. the fundamentals):**
- You can use all built-in PDF fonts: Courier, Helvetica, Symbol, Times, ZapfDingbats
- Draw Lines with different thickness
- Draw Filled / outline rectangles
- Draw Grayscale PNG images

**Not Supported (i.e. everything else):**
- Unicode characters
- Font embedding
- PDF encryption
- Paths, Circles and complex graphics

*... adding usage documentation / examples ...*
