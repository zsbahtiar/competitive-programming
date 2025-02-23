package general

func SequentialSearch(nums []int, search int) bool {
	found := false
	for _, num := range nums {
		if num == search {
			found = true
			break
		}
	}
	return found
}

func BinarySearch(nums []int, search int) bool {
	nums = MergeSort(nums)
	left := 0
	right := len(nums) - 1
	found := false

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == search {
			found = true
			break
		} else if nums[mid] < search {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return found
}
