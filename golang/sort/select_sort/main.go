package select_sort

func SelectSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	for i := 0; i < len(nums); i++ {
		min := 0
		for j := i; j < len(nums); j++ {
			if nums[j] < nums[min] {
				min = j
			}
		}
		tmp := nums[i]
		nums[i] = nums[min]
		nums[min] = tmp
	}
}
