package ty

import (
	"reflect"
)

func ZeroValue(typ reflect.Type) reflect.Value {
	return reflect.New(typ).Elem()
}

type Swapper reflect.Value

func SwapperOf(typ reflect.Type) Swapper {
	return Swapper(ZeroValue(typ))
}

func (s Swapper) Swap(a, b reflect.Value) {
	vs := reflect.Value(s)
	vs.Set(a)
	a.Set(b)
	b.Set(vs)
}

func Call(f reflect.Value, args ...reflect.Value) {
	f.Call(args)
}

func Call1(f reflect.Value, args ...reflect.Value) reflect.Value {
	return f.Call(args)[0]
}

func Call2(f reflect.Value, args ...reflect.Value) (
	reflect.Value, reflect.Value) {

	ret := f.Call(args)
	return ret[0], ret[1]
}
