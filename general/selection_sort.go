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
