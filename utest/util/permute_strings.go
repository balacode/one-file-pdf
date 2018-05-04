// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-05-04 23:53:15 26AD8A                [utest/util/permute_strings.go]
// -----------------------------------------------------------------------------

package util

import (
	"bytes"
)

// PermuteStrings returns all combinations of strings in 'parts'
func PermuteStrings(parts ...[]string) (ret []string) {
	{
		var n = 1
		for _, ar := range parts {
			n *= len(ar)
		}
		ret = make([]string, 0, n)
	}
	var at = make([]int, len(parts))
	var buf bytes.Buffer
loop:
	for {
		// increment position counters
		for i := len(parts) - 1; i >= 0; i-- {
			if at[i] > 0 && at[i] >= len(parts[i]) {
				if i == 0 || (i == 1 && at[i-1] == len(parts[0])-1) {
					break loop
				}
				at[i] = 0
				at[i-1]++
			}
		}
		// construct permutated string
		buf.Reset()
		for i, ar := range parts {
			var j = at[i]
			if j >= 0 && j < len(ar) {
				buf.WriteString(ar[j])
			}
		}
		ret = append(ret, buf.String())
		at[len(parts)-1]++
	}
	return ret
} //                                                              PermuteStrings

//end
