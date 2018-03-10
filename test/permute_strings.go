// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-03-11 00:26:09 5E7042                      [test/permute_strings.go]
// -----------------------------------------------------------------------------

package main

import "bytes" // standard

// permuteStrings returns all combinations of strings in 'parts'
func permuteStrings(parts ...[]string) (ret []string) {
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
			var p = at[i]
			if p >= 0 && p < len(ar) {
				buf.WriteString(ar[p])
			}
		}
		ret = append(ret, buf.String())
		at[len(parts)-1]++
	}
	return ret
} //                                                              permuteStrings

//end
