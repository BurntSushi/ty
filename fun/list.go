package fun

import (
	"reflect"

	"github.com/BurntSushi/ty"
)

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
