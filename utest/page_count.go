// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2019-04-03 10:03:38 C9F56A             one-file-pdf/utest/[page_count.go]
// -----------------------------------------------------------------------------

package utest

import (
	"fmt"
	"testing"

	"github.com/balacode/one-file-pdf"
)

// Test_PDF_PageCount_ tests PDF.PageCount()
func Test_PDF_PageCount_(t *testing.T) {
	fmt.Println("Test PDF.PageCount()")
	//                                                 uninitialized PDF: 1 page
	func() {
		var doc pdf.PDF
		TEqual(t, doc.PageCount(), 1)
	}()
	//                  uninitialized PDF with initial call to AddPage(): 1 page
	func() {
		var doc pdf.PDF
		doc.AddPage()
		TEqual(t, doc.PageCount(), 1)
	}()
	//                                                   initialized PDF: 1 page
	func() {
		doc := pdf.NewPDF("A4")
		TEqual(t, doc.PageCount(), 1)
	}()
	//                    initialized PDF with initial call to AddPage(): 1 page
	func() {
		doc := pdf.NewPDF("A4")
		doc.AddPage()
		TEqual(t, doc.PageCount(), 1)
	}()
	//                              initialized PDF with a single method: 1 page
	func() {
		doc := pdf.NewPDF("LETTER")
		doc.SetXY(1, 1)
		TEqual(t, doc.PageCount(), 1)
	}()
	//                  calling AddPage() after any method, increases page count
	func() {
		var doc pdf.PDF //                                     uninitialized PDF
		doc.SetXY(1, 1)
		doc.AddPage()
		TEqual(t, doc.PageCount(), 2)
	}()
	func() {
		doc := pdf.NewPDF("LETTER") //                           initialized PDF
		doc.SetXY(1, 1)
		doc.AddPage()
		TEqual(t, doc.PageCount(), 2)
	}()
	//                  after calling AddPage() 10 times, PageCount() must be 10
	func() {
		var doc pdf.PDF //                                     uninitialized PDF
		for i := 0; i < 10; i++ {
			doc.AddPage()
		}
		TEqual(t, doc.PageCount(), 10)
	}()
	func() {
		doc := pdf.NewPDF("LETTER") //                           initialized PDF
		for i := 0; i < 10; i++ {
			doc.AddPage()
		}
		TEqual(t, doc.PageCount(), 10)
	}()
} //                                                         Test_PDF_PageCount_

//end
