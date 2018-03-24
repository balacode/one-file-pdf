// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-03-24 18:56:20 90D1FF                        [utest/current_page.go]
// -----------------------------------------------------------------------------

package utest

import "fmt"     // standard
import "testing" // standard

import "github.com/balacode/one-file-pdf"

func CurrentPage(t *testing.T) {
	fmt.Println("utest.CurrentPage")

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

} //                                                                 CurrentPage

//end
