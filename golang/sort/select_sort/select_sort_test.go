package select_sort

import (
	"testing"
)

func TestSelectSort(t *testing.T) {
	nums := []int{3, 4, 1, 8, 6, 9}
	SelectSort(nums)
	t.Log(nums)
}
