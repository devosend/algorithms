package bucket_sort

func BucketSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	max := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > nums[max] {
			max = i
		}
	}

	arr := make([]int, max)
	for j := 0; j < len(nums); j++ {
		arr[nums[j]]++
	}

	ret := make([]int, len(nums))
	for k := 0; k < len(arr); k++ {
		count := arr[k]
		for ; count > 0; count-- {
			ret = append(ret, k)
		}
	}
	return ret
}
