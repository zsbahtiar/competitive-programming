package leetcode

// Given a string s consisting of words and spaces, return the length of the last word in the string.

// A word is a maximal
// substring
//  consisting of non-space characters only.

// Example 1:

// Input: s = "Hello World"
// Output: 5
// Explanation: The last word is "World" with length 5.
// Example 2:

// Input: s = "   fly me   to   the moon  "
// Output: 4
// Explanation: The last word is "moon" with length 4.
// Example 3:

// Input: s = "luffy is still joyboy"
// Output: 6
// Explanation: The last word is "joyboy" with length 6.

// Constraints:

// 1 <= s.length <= 104
// s consists of only English letters and spaces ' '.
// There will be at least one word in s.
func LengthOfLastWord(s string) int {
	counter := 0
	x := len(s) - 1 // because last word, we can count from n-1 char

	// we need last index char but not whitespace
	// image "moon "  last index is " " space then set the pointer to index "n"
	for x >= 0 && s[x] == ' ' {
		x--
	}

	// contra from loop ^^^
	// we need break counter when found space
	for x >= 0 && s[x] != ' ' {
		counter++
		x--
	}

	return counter
}
