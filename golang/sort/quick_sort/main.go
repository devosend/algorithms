package quick_sort

func QuickSort(nums []int) {
	quickSort(nums, 0, len(nums)-1)
}

func quickSort(nums []int, start int, end int) {
	if start >= end {
		return
	}
	mid := partition(nums, start, end)
	quickSort(nums, start, mid-1)
	quickSort(nums, mid+1, end)
}

func partition(nums []int, start int, end int) int {
	i := start
	point := nums[end]
	for j := i + 1; j < end; j++ {
		if nums[j] < point {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	nums[i], nums[end] = nums[end], nums[i]
	return i
}
