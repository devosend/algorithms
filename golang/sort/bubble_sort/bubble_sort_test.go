package bubble_sort

import (
	"testing"
)

func TestBubbleSort(t *testing.T) {
	nums := []int{3, 4, 1, 8, 6, 9}
	BubbleSort(nums)
	t.Log(nums)
}
