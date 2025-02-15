package leetcode

// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
// An input string is valid if:
// Open brackets must be closed by the same type of brackets.
// Open brackets must be closed in the correct order.
// Every close bracket has a corresponding open bracket of the same type.

// Example 1:
// Input: s = "()"
// Output: true
// Example 2:
// Input: s = "()[]{}"
// Output: true
// Example 3:
// Input: s = "(]"
// Output: false
// Example 4:
// Input: s = "([])"
// Output: true

func ValidParentheses(s string) bool {
	if len(s) == 0 {
		return true
	}

	type ptr struct {
		status  string
		partner string
	}

	var bracket map[string]ptr = map[string]ptr{
		"(": {
			status:  "start",
			partner: ")",
		},
		")": {
			status:  "end",
			partner: "(",
		},
		"{": {
			status:  "start",
			partner: "}",
		},
		"}": {
			status:  "end",
			partner: "{",
		},
		"[": {
			status:  "start",
			partner: "]",
		},
		"]": {
			status:  "end",
			partner: "[",
		},
	}

	var temp []string

	for _, char := range s {
		pt, exists := bracket[string(char)]
		if exists && pt.status == "start" {
			temp = append(temp, string(char))
		} else if exists && pt.status == "end" {
			if len(temp) > 0 && temp[len(temp)-1] == pt.partner {
				temp = temp[:len(temp)-1]
			} else {
				return false
			}
		}
	}

	// key of parentheses, using the LIFO
	return len(temp) == 0
}
