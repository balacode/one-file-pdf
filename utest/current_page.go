// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-03-17 23:51:20 3CA33A                        [utest/current_page.go]
// -----------------------------------------------------------------------------

package utest

import "fmt"     // standard
import "testing" // standard

import "github.com/balacode/one-file-pdf"

func CurrentPage(t *testing.T) {
	fmt.Println("utest.CurrentPage")
	//
	func() {
		var ob pdf.PDF
		//
		// before calling AddPage(), returns 1
		TEqual(t, ob.CurrentPage(), 1)
		//
		// since AddPage() is called without any drawing method,
		// the page is added implicitly: therefore still on page 1
		ob.AddPage()
		TEqual(t, ob.CurrentPage(), 1)
		//
		// the next call to AddPage(), returns 2, and so on
		ob.AddPage()
		TEqual(t, ob.CurrentPage(), 2)
		//
		ob.AddPage()
		TEqual(t, ob.CurrentPage(), 3)
		//
		ob.AddPage()
		TEqual(t, ob.CurrentPage(), 4)
	}()
	//
	func() {
		var ob = pdf.NewPDF("LETTER")
		//
		// before calling AddPage(), returns 1
		TEqual(t, ob.CurrentPage(), 1)
		//
		// since AddPage() is called without any drawing method,
		// the page is added implicitly: therefore still on page 1
		ob.AddPage()
		TEqual(t, ob.CurrentPage(), 1)
		//
		// the next call to AddPage(), returns 2, and so on
		ob.AddPage()
		TEqual(t, ob.CurrentPage(), 2)
		//
		ob.AddPage()
		TEqual(t, ob.CurrentPage(), 3)
		//
		ob.AddPage()
		TEqual(t, ob.CurrentPage(), 4)
	}()
} //                                                                 CurrentPage

//end
