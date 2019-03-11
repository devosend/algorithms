package insert_sort

func InsertSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	for i := 0; i < len(nums); i++ {
		j := i + 1
		for ; j > 0; j-- {
			if nums[j] < nums[j-1] {
				nums[j] = nums[j-1]
			} else {
				break
			}
		}
		nums[j] = nums[i+1]
	}
}
