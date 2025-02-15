package leetcode

// Write a function to find the longest common prefix string amongst an array of strings.
// If there is no common prefix, return an empty string "".

// Example 1:
// Input: strs = ["flower","flow","flight"]
// Output: "fl"
// Example 2:
// Input: strs = ["dog","racecar","car"]
// Output: ""
// Explanation: There is no common prefix among the input strings.

// Constraints:
// 1 <= strs.length <= 200
// 0 <= strs[i].length <= 200
// strs[i] consists of only lowercase English letters if it is non-empty.

func longestCommonPrefix(strs []string) string {
	if len(strs) < 1 {
		return ""
	}

	// "common prefix" is a characters same at first index of all strings
	// so we can iterate over the first string and compare it with the rest
	prefix := strs[0]
	for _, str := range strs[1:] {
		for x := 0; x < len(prefix); x++ {
			if x == len(str) || str[x] != prefix[x] {
				prefix = prefix[:x]
				break
			}
		}
	}

	return prefix
}
