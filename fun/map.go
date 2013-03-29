package fun

import (
	"reflect"

	"github.com/BurntSushi/ty"
)

func Keys(m interface{}) interface{} {
	uni := ty.Unify(
		new(func(map[ty.A]ty.B) []ty.A),
		m)
	vm, tkeys := uni.Args[0], uni.Returns[0]

	vkeys := reflect.MakeSlice(tkeys, vm.Len(), vm.Len())
	for i, vkey := range vm.MapKeys() {
		vkeys.Index(i).Set(vkey)
	}
	return vkeys.Interface()
}

func Values(m interface{}) interface{} {
	uni := ty.Unify(
		new(func(map[ty.A]ty.B) []ty.B),
		m)
	vm, tvals := uni.Args[0], uni.Returns[0]

	vvals := reflect.MakeSlice(tvals, vm.Len(), vm.Len())
	for i, vkey := range vm.MapKeys() {
		vvals.Index(i).Set(vm.MapIndex(vkey))
	}
	return vvals.Interface()
}

// func MapMerge(m1, m2 interface{}) interface{} {
// }
