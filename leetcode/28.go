package leetcode

// Given two strings needle and haystack, return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.

// Example 1:

// Input: haystack = "sadbutsad", needle = "sad"
// Output: 0
// Explanation: "sad" occurs at index 0 and 6.
// The first occurrence is at index 0, so we return 0.
// Example 2:

// Input: haystack = "leetcode", needle = "leeto"
// Output: -1
// Explanation: "leeto" did not occur in "leetcode", so we return -1.

// Constraints:

// 1 <= haystack.length, needle.length <= 104
// haystack and needle consist of only lowercase English characters.
// option : Knuth-Morris-Pratt(KMP), Naive String Matching, Boyer-Moore,Rabin-Karp
// try with KMP
// detail: https://cp-algorithms.com/string/prefix-function.html
func FindTheIndexOfTheFirstOccurrenceInAString(haystack string, needle string) int {
	if len(needle) < 1 {
		return 0
	}

	nHaystack, nNeedle := len(haystack), len(needle)
	if nNeedle > nHaystack {
		return -1
	}

	// lps
	// step 1: create lps table (the cheatsheet)
	// example case: needle = "sad"
	longestPrefixSuffix := make([]int, nNeedle) // creates [0,0,0]

	// why x = 1:
	// 1. lps[0] always 0
	// 2. need min 2-char to compare prefix and suffix
	x := 1
	lengthOfLps := 0 // tracks the length of prev lps

	// example case: "sad"
	// for x = 1: 'a' compared with 's' (index 0)
	// not same: set lps[1] = 0, move to x = 2
	// for x = 2: 'd' compared with 's' (index 0)
	// not same: set lps[2] = 0, move to x = 3
	// final lps: [0,0,0]
	for x < nNeedle {
		if needle[x] == needle[lengthOfLps] {
			lengthOfLps++
			longestPrefixSuffix[x] = lengthOfLps
			x++
		} else {
			if lengthOfLps > 0 {
				lengthOfLps = longestPrefixSuffix[lengthOfLps-1]
			} else {
				longestPrefixSuffix[x] = 0
				x++
			}
		}
	}

	// step 2: use lps to find pattern
	// for haystack = "sadbutsad" and needle = "sad"
	x = 0  // index for haystack
	y := 0 // index for needle/pattern

	for x < nHaystack {
		// if chars match, move both pointers
		if haystack[x] == needle[y] {
			x++
			y++
		}

		// if we reached end of pattern, we found it!
		// example: x = 3 (after matching 'sad')
		//          y = 3 (length of 'sad')
		//          so x-y = 0 (start of pattern)
		if y == nNeedle {
			return x - y
		} else if x < nHaystack && haystack[x] != needle[y] {
			// if mismatch occurs
			if y > 0 {
				// use lps table to skip comparisons
				y = longestPrefixSuffix[y-1]
			} else {
				// if at start of pattern, just move x
				x++
			}
		}
	}

	return -1
}
