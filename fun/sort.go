package fun

import (
	"reflect"
	"sort"

	"github.com/BurntSushi/ty"
)

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
