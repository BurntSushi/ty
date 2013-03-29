package fun

import (
	"reflect"
	"sort"

	"github.com/BurntSushi/ty"
)

// QuickSort has a parametric type:
//
//	func QuickSort(less func(x1 A, x2 A) bool, []A) []A
//
// QuickSort applies the "quicksort" algorithm to return a new sorted list
// of `xs`, where `xs` is not modified.
//
// `less` should be a function that returns true if and only if `x1` is less
// than `x2`.
func QuickSort(less, xs interface{}) interface{} {
	uni := ty.Unify(
		new(func(func(ty.A, ty.A) bool, []ty.A) []ty.A),
		less, xs)
	vless, vxs, tys := uni.Args[0], uni.Args[1], uni.Returns[0]

	vys := reflect.MakeSlice(tys, vxs.Len(), vxs.Len())
	reflect.Copy(vys, vxs)

	var qsort func(left, right int)
	var partition func(left, right, pivot int) int
	swapper := ty.SwapperOf(vys.Type().Elem())

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
		swapper.Swap(vys.Index(pivot), vys.Index(right))
		vpivot := vys.Index(right)

		ind := left
		for i := left; i < right; i++ {
			temp := vys.Index(i)
			if ty.Call1(vless, temp, vpivot).Bool() {
				swapper.Swap(temp, vys.Index(ind))
				ind++
			}
		}
		swapper.Swap(vys.Index(ind), vys.Index(right))
		return ind
	}
	qsort(0, vys.Len()-1)
	return vys.Interface()
}

// Sort has a parametric type:
//
//	func Sort(less func(x1 A, x2 A) bool, []A)
//
// Sort uses the standard library `sort` package to sort `xs` in place.
//
// `less` should be a function that returns true if and only if `x1` is less
// than `x2`.
func Sort(less, xs interface{}) {
	uni := ty.Unify(
		new(func(func(ty.A, ty.A) bool, []ty.A)),
		less, xs)

	vless, vxs := uni.Args[0], uni.Args[1]
	sort.Sort(&sortable{vless, vxs, ty.SwapperOf(vxs.Type().Elem())})
}

type sortable struct {
	less    reflect.Value
	xs      reflect.Value
	swapper ty.Swapper
}

func (s *sortable) Less(i, j int) bool {
	ith, jth := s.xs.Index(i), s.xs.Index(j)
	return ty.Call1(s.less, ith, jth).Bool()
}

func (s *sortable) Swap(i, j int) {
	s.swapper.Swap(s.xs.Index(i), s.xs.Index(j))
}

func (s *sortable) Len() int {
	return s.xs.Len()
}
