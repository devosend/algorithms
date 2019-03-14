package binary_search

import "testing"

func TestBinarySearch(t *testing.T) {
	nums := []int{1, 3, 4, 5, 33, 54, 67, 85}
	count := BinarySearch(nums, 33)
	if count == 4 {
		t.Log("success")
	} else {
		t.Fatal("search failed")
	}

	count = BinarySearch(nums, 34)
	if count == -1 {
		t.Log("success")
	} else {
		t.Fatal("search failed")
	}
}
