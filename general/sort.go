package general

func SelectionSort(nums []int) []int {
	for x := 0; x < len(nums); x++ {
		for y := x + 1; y < len(nums); y++ {
			if nums[y] < nums[x] {
				nums[y], nums[x] = nums[x], nums[y]
			}
		}
	}
	return nums
}

func BubbleSort(nums []int) []int {
	for x := 0; x < len(nums)-1; x++ {
		for y := 0; y < len(nums)-x-1; y++ {
			if nums[y] > nums[y+1] {
				nums[y], nums[y+1] = nums[y+1], nums[y]
			}
		}
	}
	return nums
}

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
