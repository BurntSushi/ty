package fun

import (
	"sort"
	"testing"
)

func BenchmarkSortReflect(b *testing.B) {
	less := func(a, b int) bool { return a < b }

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		list := randIntSlice(1000)
		b.StartTimer()

		Sort(less, list)
	}
}

func BenchmarkSortBuiltin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		list := randIntSlice(1000)
		b.StartTimer()

		sort.Sort(sort.IntSlice(list))
	}
}

func BenchmarkQuickSortReflect(b *testing.B) {
	less := func(a, b int) bool { return a < b }

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		list := randIntSlice(1000)
		b.StartTimer()

		_ = QuickSort(less, list)
	}
}

func BenchmarkQuickSortBuiltin(b *testing.B) {
	less := func(a, b int) bool { return a < b }

	quicksort := func(xs []int) []int {
		ys := make([]int, len(xs))
		copy(ys, xs)

		var qsort func(left, right int)
		var partition func(left, right, pivot int) int

		qsort = func(left, right int) {
			if left >= right {
				return
			}
			pivot := (left + right) / 2
			pivot = partition(left, right, pivot)

			qsort(left, pivot-1)
			qsort(pivot+1, right)
		}
		partition = func(left, right, pivot int) int {
			vpivot := ys[pivot]
			ys[pivot], ys[right] = ys[right], ys[pivot]

			ind := left
			for i := left; i < right; i++ {
				if less(ys[i], vpivot) {
					ys[i], ys[ind] = ys[ind], ys[i]
					ind++
				}
			}
			ys[ind], ys[right] = ys[right], ys[ind]
			return ind
		}

		return ys
	}

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		list := randIntSlice(1000)
		b.StartTimer()

		_ = quicksort(list)
	}
}
