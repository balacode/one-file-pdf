// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-05-04 23:51:59 B20A7E                        [utest/util/t_equal.go]
// -----------------------------------------------------------------------------

package util

// Provides a slightly-altered TEqual() function (and functions it uses)
// from Zircon-Go lib: github.com/balacode/zr

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"testing"
	"time"
)

const showFileNames = 1

// 0 - Don't show file names
// 1 - Show only file name
// 2 - Show file name and path

// PL is fmt.Println() but is used only for debugging.
var PL = fmt.Println

// TEqual asserts that result is equal to expect.
func TEqual(t *testing.T, result interface{}, expect interface{}) bool {
	var makeStr = func(val interface{}) string {
		switch val := val.(type) {
		case nil:
			return "nil"
		case bool:
			if val {
				return "true"
			}
			return "false"
		case int, int8, int16, int32, int64,
			uint, uint8, uint16, uint32, uint64, uintptr:
			return fmt.Sprintf("%d", val)
		case float32, float64:
			var s = fmt.Sprintf("%.4f", val)
			if strings.Contains(s, ".") {
				for strings.HasSuffix(s, "0") {
					s = s[:len(s)-1]
				}
				for strings.HasSuffix(s, ".") {
					s = s[:len(s)-1]
				}
			}
			return s
		case error:
			return val.Error()
		case string:
			return val
		case time.Time: // use date part without time and time zone
			var s = val.Format(time.RFC3339)[:19] // "2006-01-02T15:04:05Z07:00"
			if strings.HasSuffix(s, "T00:00:00") {
				s = s[:10]
			}
			return s
		case fmt.Stringer:
			return val.String()
		}
		return fmt.Sprintf("(type: %v value: %v)", reflect.TypeOf(val), val)
	}
	if makeStr(result) != makeStr(expect) {
		t.Logf("\n"+"LOCATION: %s\n"+"EXPECTED: %s\n"+"RETURNED: %s\n",
			TCaller(), makeStr(expect), makeStr(result))
		t.Fail()
		return false
	}
	return true
} //                                                                      TEqual

// -----------------------------------------------------------------------------

// CallerList returns a human-friendly list of strings showing the
// call stack with each calling method or function's name and line number.
//
// The most immediate callers are listed first, followed by their callers,
// and so on. For brevity, 'runtime.*' and 'syscall.*'
// and other top-level callers are not included.
func CallerList() []string {
	var ret []string
	var i = 0
mainLoop:
	for {
		i++
		var programCounter, filename, lineNo, _ = runtime.Caller(i)
		var funcName = runtime.FuncForPC(programCounter).Name()
		//
		// end loop on reaching a top-level runtime function
		for _, s := range []string{
			"", "runtime.goexit", "runtime.main", "testing.tRunner",
		} {
			if funcName == s {
				break mainLoop
			}
		}
		if strings.Contains(funcName, "HandlerFunc.ServeHTTP") {
			break
		}
		// skip runtime/syscall functions, but continue the loop
		for _, s := range []string{
			".Callers", ".CallerList", ".Error", ".Log", ".logAsync",
			"mismatch", "runtime.", "syscall.",
		} {
			if strings.Contains(funcName, s) {
				continue mainLoop
			}
		}
		switch showFileNames {
		case 1:
			filename = filepath.Base(filename)
		case 2:
			// let the file name's path use the right kind of OS path separator
			// (by default, the file name contains '/' on all platforms)
			if string(os.PathSeparator) != "/" {
				filename = strings.Replace(filename,
					"/", string(os.PathSeparator), -1)
			}
		}
		// remove parent module/function names
		if index := strings.LastIndex(funcName, "/"); index != -1 {
			funcName = funcName[index+1:]
		}
		if strings.Count(funcName, ".") > 1 {
			funcName = funcName[strings.Index(funcName, ".")+1:]
		}
		// remove unneeded punctuation from function names
		for _, find := range []string{"(", ")", "*"} {
			if strings.Contains(funcName, find) {
				funcName = strings.Replace(funcName, find, "", -1)
			}
		}
		var line = fmt.Sprintf(":%d %s()", lineNo, funcName)
		if showFileNames > 0 {
			line = filename + line
		}
		ret = append(ret, line)
	}
	return ret
} //                                                                  CallerList

// TCaller returns the name of the unit test function.
func TCaller() string {
	for _, iter := range CallerList() {
		if strings.Contains(iter, "util.TCaller") ||
			strings.Contains(iter, "util.TEqual") ||
			strings.Contains(iter, "util.ComparePDF") {
			continue
		}
		return iter
	}
	return "<no-caller>"
} //                                                                     TCaller

//end