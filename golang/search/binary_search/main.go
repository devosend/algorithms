package binary_search

func search(nums []int, start int, end int, n int) int {
	if start > end {
		return -1
	}
	mid := (start + end) / 2
	if n == nums[mid] {
		return mid
	} else if n < nums[mid] {
		return search(nums, start, mid-1, n)
	} else {
		return search(nums, mid+1, end, n)
	}

}
func BinarySearch(nums []int, n int) int {
	start := 0
	end := len(nums) - 1
	return search(nums, start, end, n)
}
