package poly

import (
	"fmt"
	"testing"
)

func square(x int) int {
	return x * x
}

func TestMap(t *testing.T) {
	squares := Map(square, []int{1, 2, 3, 4, 5}).([]int)
	for _, sq := range squares {
		fmt.Printf("%d ", sq)
	}
	fmt.Println()
	// fmt.Printf("%v\n", Map(square, []int{1, 2, 3, 4, 5}))
}
