## Detailed Changelog  

These are the changes in the functionality of the package, 
not including internal changes. Internal changes are are 
best seen in the commits history.  

**2018-MAR-30**  
- **ALTERED API: Removed SetErrorLogger() method**
- **ALTERED API: ToPoints(): added error return value**  
- Initialize PDF automatically, even when NewPDF() wasn't called. The paper size
  is A4, and the units CM by default. To specify a different paper size, use NewPDF().
- No need to add the first page with AddPage(). It is inserted automatically.
- New error handling methods Clean(), Errors(), ErrorInfo() and PullError().
- SetColumnWidths(): can be called without arguments when you need to reset all columns.
- Added various unit tests, for 95% code coverage.
- Fixed text wrapping bug that could cause PDF to freeze.

**2018-MAR-14**
- Added support for color JPEG, GIF and PNG images with transparency blending
- Added all standard A, B, and C paper sizes, and US Tabloid and Ledger
- DrawImage(): added backColor optional parameter (using ...) so you can specify the background color for transparent PNGs
- DrawImage(): changed to draw images down from the Y-coordinate position (below, not above Y)
- Created "PNG Images Demo", which outputs to png_images.pdf
- Created ToColor() function to convert named colors and HTML color codes to RGBA color values
- Created ToUnits() method to convert points to the currently-active units of measurement
- Removed PDFNoPage constant
- Created basic unit tests and test helper functions
- Various internal changes, reducing file length by about 60 lines

**2018-MAR-08**
- New methods DrawCircle(), DrawEllipse(), FillCircle(), FillEllipse()
- New demo demonstrating circles and text wrapping: corporate.pdf ("Synergy Ipsum")
- SetColor(): now allows HTML color values like `"#4C9141 MayGreen"`, ignores the extra chars.
- Log an error when the name of a selected font is unknown.
- Log the specified measurement unit's name when its name is not valid.

**2018-MAR-09**
- Replaced PDFColor type with standard lib's color.RGBA (import "image/color")
- SetColorRGB(): changed parameters from int to uint8
- Changed PDFPageSize and PDFStandardPageSizes to private structures
- SaveFile() now returns `error` instead of `*PDF`, to allow caller to check for IO errors
- SetColumnWidths() is no longer limited to 100 columns
- Font names and color names can be specified with spaces, underscores or '-' delimiting words
- Removed module-global PDFErrorHandler, created SetErrorLogger() to set the handler for each PDF instance
