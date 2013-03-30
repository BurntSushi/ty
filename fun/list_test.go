package fun

import (
	"testing"
)

func TestMap(t *testing.T) {
	square := func(x int) int { return x * x }
	squares := Map(square, []int{1, 2, 3, 4, 5}).([]int)

	assertDeep(t, squares, []int{1, 4, 9, 16, 25})
}

func TestPointers(t *testing.T) {
	type temp struct {
		val int
	}
	square := func(t *temp) *temp { return &temp{t.val * t.val} }
	squares := Map(square, []*temp{
		{1}, {2}, {3}, {4}, {5},
	})

	assertDeep(t, squares, []*temp{
		{1}, {4}, {9}, {16}, {25},
	})
}

func BenchmarkMapReflect(b *testing.B) {
	b.StopTimer()
	list := randIntSlice(1000)
	square := func(a int) int {
		return a * a
	}
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		_ = Map(square, list).([]int)
	}
}

func BenchmarkMapBuiltin(b *testing.B) {
	b.StopTimer()
	list := randIntSlice(1000)
	square := func(a int) int {
		return a * a
	}
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		ret := make([]int, len(list))
		for i := 0; i < len(list); i++ {
			ret[i] = square(list[i])
		}
	}
}
