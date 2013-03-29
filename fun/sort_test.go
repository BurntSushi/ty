package fun

import (
	"testing"
)

func TestSort(t *testing.T) {
	tosort := []int{10, 3, 5, 1, 15, 6}
	Sort(func(a, b int) bool { return b < a }, tosort)

	assertDeep(t, tosort, []int{15, 10, 6, 5, 3, 1})
}

func TestQuickSort(t *testing.T) {
	tosort := []int{10, 3, 5, 1, 15, 6}
	sorted := QuickSort(
		func(a, b int) bool {
			return b < a
		}, tosort).([]int)

	assertDeep(t, sorted, []int{15, 10, 6, 5, 3, 1})
}
