package fun

import (
	"reflect"
	"runtime"
	"sync"

	"github.com/BurntSushi/ty"
)

// Map has a parametric type:
//
//	func Map(f func(A) B, xs []A) []B
//
// Map returns the list corresponding to the return value of applying
// `f` to each element in `xs`.
func Map(f, xs interface{}) interface{} {
	uni := ty.Check(
		new(func(func(ty.A) ty.B, []ty.A) []ty.B),
		f, xs)
	vf, vxs, tys := uni.Args[0], uni.Args[1], uni.Returns[0]

	xsLen := vxs.Len()
	vys := reflect.MakeSlice(tys, xsLen, xsLen)
	for i := 0; i < xsLen; i++ {
		vy := call1(vf, vxs.Index(i))
		vys.Index(i).Set(vy)
	}
	return vys.Interface()
}

// Filter has a parametric type:
//
//	func Filter(p func(A) bool, xs []A) []A
//
// Filter returns a new list only containing the elements of `xs` that satisfy
// the predicate `p`.
func Filter(p, xs interface{}) interface{} {
	uni := ty.Check(
		new(func(func(ty.A) bool, []ty.A) []ty.A),
		p, xs)
	vp, vxs, tys := uni.Args[0], uni.Args[1], uni.Returns[0]

	xsLen := vxs.Len()
	vys := reflect.MakeSlice(tys, 0, xsLen)
	for i := 0; i < xsLen; i++ {
		vx := vxs.Index(i)
		if call1(vp, vx).Bool() {
			vys = reflect.Append(vys, vx)
		}
	}
	return vys.Interface()
}

// Foldl has a parametric type:
//
//	func Foldl(f func(A, B) B, init B, xs []A) B
//
// Foldl reduces a list of A to a single element B using a left fold with
// an initial value `init`.
func Foldl(f, init, xs interface{}) interface{} {
	uni := ty.Check(
		new(func(func(ty.A, ty.B) ty.B, ty.B, []ty.A) ty.B),
		f, init, xs)
	vf, vinit, vxs, tb := uni.Args[0], uni.Args[1], uni.Args[2], uni.Returns[0]

	xsLen := vxs.Len()
	vb := zeroValue(tb)
	vb.Set(vinit)
	if xsLen == 0 {
		return vb.Interface()
	}

	vb.Set(call1(vf, vxs.Index(0), vb))
	for i := 1; i < xsLen; i++ {
		vb.Set(call1(vf, vxs.Index(i), vb))
	}
	return vb.Interface()
}

// Foldr has a parametric type:
//
//	func Foldr(f func(A, B) B, init B, xs []A) B
//
// Foldr reduces a list of A to a single element B using a right fold with
// an initial value `init`.
func Foldr(f, init, xs interface{}) interface{} {
	uni := ty.Check(
		new(func(func(ty.A, ty.B) ty.B, ty.B, []ty.A) ty.B),
		f, init, xs)
	vf, vinit, vxs, tb := uni.Args[0], uni.Args[1], uni.Args[2], uni.Returns[0]

	xsLen := vxs.Len()
	vb := zeroValue(tb)
	vb.Set(vinit)
	if xsLen == 0 {
		return vb.Interface()
	}

	vb.Set(call1(vf, vxs.Index(xsLen-1), vb))
	for i := xsLen - 2; i >= 0; i-- {
		vb.Set(call1(vf, vxs.Index(i), vb))
	}
	return vb.Interface()
}

// Concat has a parametric type:
//
//	func Concat(xs [][]A) []A
//
// Concat returns a new flattened list by appending all elements of `xs`.
func Concat(xs interface{}) interface{} {
	uni := ty.Check(
		new(func([][]ty.A) []ty.A),
		xs)
	vxs, tflat := uni.Args[0], uni.Returns[0]

	xsLen := vxs.Len()
	vflat := reflect.MakeSlice(tflat, 0, xsLen*3)
	for i := 0; i < xsLen; i++ {
		vflat = reflect.AppendSlice(vflat, vxs.Index(i))
	}
	return vflat.Interface()
}

// ParMap has a parametric type:
//
//	func Map(f func(A) B, xs []A) []B
//
// ParMap is just like Map, except it applies `f` to each element in `xs`
// concurrently using N worker goroutines (where N is the number of CPUs
// available reported by the Go runtime).
//
// It is important that `f` not be a trivial operation, otherwise the overhead
// of executing it concurrently will result in worse performance than using
// a `Map`.
func ParMap(f, xs interface{}) interface{} {
	uni := ty.Check(
		new(func(func(ty.A) ty.B, []ty.A) []ty.B),
		f, xs)
	vf, vxs, tys := uni.Args[0], uni.Args[1], uni.Returns[0]

	xsLen := vxs.Len()
	ys := reflect.MakeSlice(tys, xsLen, xsLen)

	N := runtime.NumCPU()
	if N < 1 {
		N = 1
	}
	work := make(chan int, N)
	wg := new(sync.WaitGroup)
	for i := 0; i < N; i++ {
		wg.Add(1)
		go func() {
			for j := range work {
				ys.Index(j).Set(call1(vf, vxs.Index(j)))
			}
			wg.Done()
		}()
	}
	for i := 0; i < xsLen; i++ {
		work <- i
	}
	close(work)
	wg.Wait()
	return ys.Interface()
}

// Range generates a list of integers corresponding to every integer in
// the half-open interval [x, y).
//
// Range will panic if `end < start`.
func Range(start, end int) []int {
	if end < start {
		panic("range must have end greater than or equal to start")
	}
	r := make([]int, end-start)
	for i := start; i < end; i++ {
		r[i-start] = i
	}
	return r
}
