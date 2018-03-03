## one-file-pdf - A minimalist PDF generator in &lt;2K lines and 1 file
[![Go Report Card](https://goreportcard.com/badge/github.com/balacode/one-file-pdf)](https://goreportcard.com/report/github.com/balacode/one-file-pdf)  

The main idea behind this project was:  
"How small can I make a PDF generator for it to still be useful for 80% of common PDF generation needs?"

The result is a single .go file with less than 1999 lines of code.

- All the basics for generating PDF documents, enough for generating mundane business reports. 
- It's easier to learn about the internals of the PDF format with a smaller library.
- You can just drop the one_file_pdf.go file in your Go project. No need to manage dependencies. (But remember to change the package to 'main' or whatever package you are adding it to.)
- The current version of the file is indicated in the header (the timestamp).

**Supported Features (the fundamentals):**
- You can use all built-in PDF fonts: Courier, Helvetica, Symbol, Times, ZapfDingbats, and their variants
- Draw lines with different thickness
- Draw filled or outline rectangles
- Draw grayscale PNG images

**Not Supported (everything else):**
- Unicode (requires font embedding)
- Font embedding
- PDF encryption
- Paths, Circles and complex graphics

*... adding usage documentation / examples ...*
