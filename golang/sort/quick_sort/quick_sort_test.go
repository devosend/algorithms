package quick_sort

import (
	"testing"
)

func TestQuickSort(t *testing.T) {
	nums := []int{3, 4, 1, 8, 6, 9}
	QuickSort(nums)
	t.Log(nums)
}
