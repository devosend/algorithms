package select_sort

func SelectSort(nums []int) {
	if len(nums) <= 1 {
		return
	}

	for i := 0; i < len(nums); i++ {
		min := i
		for j := i; j < len(nums); j++ {
			if nums[j] < nums[min] {
				min = j
			}
		}
		if i != nums[min] {
			nums[i], nums[min] = nums[min], nums[i]
		}
	}
}
