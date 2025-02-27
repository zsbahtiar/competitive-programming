package leetcode

// Given two binary strings a and b, return their sum as a binary string.

// Example 1:

// Input: a = "11", b = "1"
// Output: "100"
// Example 2:

// Input: a = "1010", b = "1011"
// Output: "10101"

// Constraints:

// 1 <= a.length, b.length <= 104
// a and b consist only of '0' or '1' characters.
// Each string does not contain leading zeros except for the zero itself.
func AddBinary(a, b string) string {
	var result string

	x := len(a) - 1
	y := len(b) - 1
	carry := 0
	for x >= 0 || y >= 0 || carry > 0 {
		if x >= 0 {
			carry += int(a[x] - '0')
			x--
		}
		if y >= 0 {
			carry += int(b[y] - '0')
			y--
		}
		result = string('0'+uint8(carry%2)) + result
		carry /= 2

	}

	return result
}
