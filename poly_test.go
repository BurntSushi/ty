package poly

import (
	"fmt"
	"testing"
)

func square(x int) int {
	return x * x
}

func TestSort(t *testing.T) {
	tosort := []int{10, 3, 5, 1, 15, 6}
	Sort(func(a, b int) bool {
		return b < a
	}, tosort)
	for _, n := range tosort {
		fmt.Printf("%d ", n)
	}
	fmt.Println()
}

func TestQuickSort(t *testing.T) {
	tosort := []int{10, 3, 5, 1, 15, 6}
	sorted := QuickSort(func(a, b int) bool {
		return b < a
	}, tosort).([]int)
	for _, n := range sorted {
		fmt.Printf("%d ", n)
	}
	fmt.Println()
}

func TestMap(t *testing.T) {
	squares := Map(square, []int{1, 2, 3, 4, 5}).([]int)
	for _, sq := range squares {
		fmt.Printf("%d ", sq)
	}
	fmt.Println()
	// fmt.Printf("%v\n", Map(square, []int{1, 2, 3, 4, 5}))
}
