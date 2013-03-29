package fun

import (
	"reflect"

	"github.com/BurntSushi/ty"
)

// Map has a parametric type:
//
//	func Map(f func(A) B, xs []A) []B
//
// Map returns the list corresponding to the return value of applying
// `f` to each element in `xs`.
func Map(f, xs interface{}) interface{} {
	uni := ty.Unify(
		new(func(func(ty.A) ty.B, []ty.A) []ty.B),
		f, xs)
	vf, vxs, tys := uni.Args[0], uni.Args[1], uni.Returns[0]

	xsLen := vxs.Len()
	vys := reflect.MakeSlice(tys, xsLen, xsLen)
	for i := 0; i < xsLen; i++ {
		vy := ty.Call1(vf, vxs.Index(i))
		vys.Index(i).Set(vy)
	}
	return vys.Interface()
}
