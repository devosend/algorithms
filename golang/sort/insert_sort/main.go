package insert_sort

func InsertSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	for i := 0; i < len(nums); i++ {
		val := nums[i]
		j := i - 1
		for ; j >= 0; j-- {
			if nums[j] > val {
				nums[j+1] = nums[j]
			} else {
				break
			}
		}
		nums[j+1] = val
	}
}
