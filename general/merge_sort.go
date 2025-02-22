package general

// divide and conqueror
func MergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	mid := len(nums) / 2
	left := nums[:mid]
	right := nums[mid:]

	left = MergeSort(left)
	right = MergeSort(right)

	var (
		result []int
		x, y   int
	)

	for x < len(left) && y < len(right) {
		if left[x] <= right[y] {
			result = append(result, left[x])
			x++
		} else {
			result = append(result, right[y])
			y++
		}
	}

	for x < len(left) {
		result = append(result, left[x])
		x++
	}

	for y < len(right) {
		result = append(result, right[y])
		y++
	}

	return result
}
