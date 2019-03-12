package merge_sort

func merge(nums []int, start int, mid int, end int) {
	arr := make([]int, end-start+1)
	i := start
	j := mid + 1
	k := 0
	for ; i <= mid && j <= end; k++ {
		if nums[i] <= nums[j] {
			arr = append(arr, nums[k])
			i++
		} else {
			arr = append(arr, nums[j+k])
			j++
		}
	}

	for ; i <= mid; i++ {
		arr = append(arr, nums[i])
	}

	for ; j <= end; j++ {
		arr = append(arr, nums[j])
	}

	copy(nums[start:end+1], arr)
}

func mergeSort(nums []int, start int, end int) {
	if start >= end {
		return
	}

	mid := (start + end) / 2
	mergeSort(nums, start, mid)
	mergeSort(nums, mid+1, end)
	merge(nums, start, mid, end)
}

func MergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	mergeSort(nums, 0, len(nums)-1)
	return nums
}
