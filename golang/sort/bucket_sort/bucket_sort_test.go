package bucket_sort

import "testing"

func TestInsertSort(t *testing.T) {
	nums := []int{3, 4, 1, 8, 6, 9}
	nums = BucketSort(nums)
	t.Log(nums)
}
