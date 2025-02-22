package general

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
