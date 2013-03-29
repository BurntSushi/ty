package fun

import (
	"reflect"

	"github.com/BurntSushi/ty"
)

// Union has a parametric type:
//
//	func(a map[A]bool, b map[A]bool) map[A]bool
//
// Union returns the union of two sets, where a set is represented as a
// `map[A]bool`. The sets `a` and `b` are not modified.
func Union(a, b interface{}) interface{} {
	uni := ty.Unify(
		new(func(map[ty.A]bool, map[ty.A]bool) map[ty.A]bool),
		a, b)
	va, vb, tc := uni.Args[0], uni.Args[1], uni.Returns[0]

	vtrue := reflect.ValueOf(true)
	vc := reflect.MakeMap(tc)
	for _, vkey := range va.MapKeys() {
		vc.SetMapIndex(vkey, vtrue)
	}
	for _, vkey := range vb.MapKeys() {
		vc.SetMapIndex(vkey, vtrue)
	}
	return vc.Interface()
}

// Intersection has a parametric type:
//
//	func(a map[A]bool, b map[A]bool) map[A]bool
//
// Intersection returns the intersection of two sets, where a set is
// represented as a `map[A]bool`. The sets `a` and `b` are not modified.
func Intersection(a, b interface{}) interface{} {
	uni := ty.Unify(
		new(func(map[ty.A]bool, map[ty.A]bool) map[ty.A]bool),
		a, b)
	va, vb, tc := uni.Args[0], uni.Args[1], uni.Returns[0]

	vtrue := reflect.ValueOf(true)
	vc := reflect.MakeMap(tc)
	for _, vkey := range va.MapKeys() {
		if vb.MapIndex(vkey).IsValid() {
			vc.SetMapIndex(vkey, vtrue)
		}
	}
	for _, vkey := range vb.MapKeys() {
		if va.MapIndex(vkey).IsValid() {
			vc.SetMapIndex(vkey, vtrue)
		}
	}
	return vc.Interface()
}
