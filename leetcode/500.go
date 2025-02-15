package leetcode

import "strings"

// Given an array of strings words, return the words that can be typed using letters of the alphabet on only one row of American keyboard like the image below.
// Note that the strings are case-insensitive, both lowercased and uppercased of the same letter are treated as if they are at the same row.
// In the American keyboard:
// the first row consists of the characters "qwertyuiop",
// the second row consists of the characters "asdfghjkl", and
// the third row consists of the characters "zxcvbnm".
func findWords(words []string) []string {
	var result []string
	for _, word := range words {
		if isOneRow(strings.ToLower(word)) {
			result = append(result, word)
		}
	}
	return result
}

func isOneRow(word string) bool {
	rows := []string{"qwertyuiop", "asdfghjkl", "zxcvbnm"}
	for _, row := range rows {
		if containsAll(word, row) {
			return true
		}
	}
	return false
}

func containsAll(word, row string) bool {
	for _, c := range word {
		if !contains(row, c) {
			return false
		}
	}
	return true
}

func contains(s string, r rune) bool {
	for _, c := range s {
		if c == r {
			return true
		}
	}
	return false
}
