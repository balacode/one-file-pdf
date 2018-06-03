// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-04-29 23:42:24 7C1E53           one-file-pdf/utest/[current_page.go]
// -----------------------------------------------------------------------------

package utest

import (
	"fmt"
	"testing"

	"github.com/balacode/one-file-pdf"
)

// Test_PDF_CurrentPage_ tests PDF.CurrentPage()
func Test_PDF_CurrentPage_(t *testing.T) {
	fmt.Println("Test PDF.CurrentPage()")

	func() {
		var doc pdf.PDF // uninitialized PDF
		//
		// before calling AddPage(), returns 1
		TEqual(t, doc.CurrentPage(), 1)
		//
		// since AddPage() is called without any drawing method,
		// the page is added implicitly: therefore still on page 1
		doc.AddPage()
		TEqual(t, doc.CurrentPage(), 1)
		//
		// the next call to AddPage(), returns 2, and so on
		doc.AddPage()
		TEqual(t, doc.CurrentPage(), 2)
		//
		doc.AddPage()
		TEqual(t, doc.CurrentPage(), 3)
		//
		doc.AddPage()
		TEqual(t, doc.CurrentPage(), 4)
	}()

	func() {
		var doc = pdf.NewPDF("LETTER")
		//
		// before calling AddPage(), returns 1
		TEqual(t, doc.CurrentPage(), 1)
		//
		// since AddPage() is called without any drawing method,
		// the page is added implicitly: therefore still on page 1
		doc.AddPage()
		TEqual(t, doc.CurrentPage(), 1)
		//
		// the next call to AddPage(), returns 2, and so on
		doc.AddPage()
		TEqual(t, doc.CurrentPage(), 2)
		//
		doc.AddPage()
		TEqual(t, doc.CurrentPage(), 3)
		//
		doc.AddPage()
		TEqual(t, doc.CurrentPage(), 4)
	}()

} //                                                       Test_PDF_CurrentPage_

//end
