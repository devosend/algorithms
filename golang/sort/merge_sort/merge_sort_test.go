package merge_sort

import (
	"testing"
)

func TestMergeSort(t *testing.T) {
	nums := []int{3, 4, 1, 8, 6, 9}
	MergeSort(nums)
	t.Log(nums)
}
