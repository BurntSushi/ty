package fun

import (
	"testing"
)

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
