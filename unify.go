package poly

import (
	"fmt"
	"reflect"
	"sort"
)

type PolyError string

func (pe PolyError) Error() string {
	return string(pe)
}

func pe(format string, v ...interface{}) PolyError {
	return PolyError(fmt.Sprintf(format, v...))
}

func ppe(format string, v ...interface{}) {
	panic(pe(format, v...))
}

type sortable struct {
	less reflect.Value
	xs   reflect.Value
}

func (s *sortable) Less(i, j int) bool {
	ith, jth := s.xs.Index(i), s.xs.Index(j)
	ret := s.less.Call([]reflect.Value{ith, jth})
	return ret[0].Bool()
}

func (s *sortable) Swap(i, j int) {
	ith := reflect.ValueOf(s.xs.Index(i).Interface())
	s.xs.Index(i).Set(s.xs.Index(j))
	s.xs.Index(j).Set(ith)
}

func (s *sortable) Len() int {
	return s.xs.Len()
}

func Sort(less, xs interface{}) {
	arg, _ := Unify(
		new(func(func(A, A) bool, []A)),
		less, xs)

	rless, rxs := arg[0], arg[1]
	sort.Sort(&sortable{rless, rxs})
}

func QuickSort(less, xs interface{}) interface{} {
	arg, ret := Unify(
		new(func(func(A, A) bool, []A) []A),
		less, xs)
	rless, rxs, tys := arg[0], arg[1], ret[0]

	var qsort func(reflect.Value) reflect.Value
	qsort = func(list reflect.Value) reflect.Value {
		listLen := list.Len()
		if listLen <= 1 {
			return list
		}

		pivot := list.Index(0)
		left := reflect.MakeSlice(tys, 0, listLen/2)
		right := reflect.MakeSlice(tys, 0, listLen/2)

		for i := 1; i < listLen; i++ {
			el := list.Index(i)
			if rless.Call([]reflect.Value{el, pivot})[0].Bool() {
				left = reflect.Append(left, el)
			} else {
				right = reflect.Append(right, el)
			}
		}
		return reflect.AppendSlice(reflect.Append(qsort(left), pivot),
			qsort(right))
	}
	return qsort(rxs).Interface()
}

func Map(f, xs interface{}) interface{} {
	var typ func(func(A) B, []A) []B
	arg, ret := Unify(typ, f, xs)
	rf, rxs, tys := arg[0], arg[1], ret[0]

	xsLen := rxs.Len()
	rys := reflect.MakeSlice(tys, xsLen, xsLen)
	for i := 0; i < xsLen; i++ {
		res := rf.Call([]reflect.Value{rxs.Index(i)})
		rys.Index(i).Set(res[0])
	}
	return rys.Interface()
}

func Unify(f interface{}, as ...interface{}) ([]reflect.Value, []reflect.Type) {
	rf := reflect.ValueOf(f)
	tf := rf.Type()

	if tf.Kind() == reflect.Ptr {
		rf = reflect.Indirect(rf)
		tf = rf.Type()
	}
	if tf.Kind() != reflect.Func {
		ppe("The type of `f` must be a function, but it is a '%s'.", tf.Kind())
	}
	if tf.NumIn() != len(as) {
		ppe("`f` expects %d arguments, but only %d were given.",
			tf.NumIn(), len(as))
	}

	// Populate the argument value list.
	args := make([]reflect.Value, len(as))
	for i := 0; i < len(as); i++ {
		args[i] = reflect.ValueOf(as[i])
	}

	// Populate our type variable environment through unification.
	tyenv := make(tyenv)
	for i := 0; i < len(args); i++ {
		tp := typePair{tyenv, tf.In(i), args[i].Type()}

		fmt.Printf("Unifying %s <-> %s\n", tp.param, tp.input)

		// Mutates the type variable environment.
		tp.unify(tp.param, tp.input)
	}

	// Now substitute those types into the return types of `f`.
	retTypes := make([]reflect.Type, tf.NumOut())
	for i := 0; i < tf.NumOut(); i++ {
		retTypes[i] = (&returnType{tyenv, tf.Out(i)}).tysubst(tf.Out(i))
	}
	return args, retTypes
}

type tyenv map[string]reflect.Type

type typePair struct {
	tyenv tyenv
	param reflect.Type
	input reflect.Type
}

func (tp typePair) panic(format string, v ...interface{}) {
	ppe("Error unifying type '%s' and '%s': %s",
		tp.param, tp.input, fmt.Sprintf(format, v...))
}

func (tp typePair) unify(param, input reflect.Type) {
	if tyname := tyvarName(input); len(tyname) > 0 {
		tp.panic("Type variables are not (yet) allowed in the types of " +
			"input arguments.")
	}
	if tyname := tyvarName(param); len(tyname) > 0 {
		if cur, ok := tp.tyenv[tyname]; ok && cur != input {
			tp.panic("Type variable %s cannot be bound to %s, as it is "+
				"already bound to %s.", tyname, input, cur)
		} else if !ok {
			fmt.Printf("Setting %s to %s\n", tyname, input)
			tp.tyenv[tyname] = input
		}
		return
	}
	if param.Kind() != input.Kind() {
		tp.panic("Cannot unify '%s' with '%s'.", param, input)
	}

	switch param.Kind() {
	case reflect.Func:
		if param.NumIn() != input.NumIn() || param.NumOut() != input.NumOut() {
			tp.panic("Cannot unify '%s' with '%s'.", param, input)
		}
		for i := 0; i < param.NumIn(); i++ {
			tp.unify(param.In(i), input.In(i))
		}
		for i := 0; i < param.NumOut(); i++ {
			tp.unify(param.Out(i), input.Out(i))
		}
	case reflect.Slice:
		tp.unify(param.Elem(), input.Elem())
	}
}

type returnType struct {
	tyenv tyenv
	typ   reflect.Type
}

func (rt returnType) panic(format string, v ...interface{}) {
	ppe("Error substituting in return type '%s': %s",
		rt.typ, fmt.Sprintf(format, v...))
}

func (rt returnType) tysubst(typ reflect.Type) reflect.Type {
	if tyname := tyvarName(typ); len(tyname) > 0 {
		if thetype, ok := rt.tyenv[tyname]; !ok {
			rt.panic("Unbound type variable %s.", tyname)
		} else {
			return thetype
		}
	}

	switch typ.Kind() {
	case reflect.Array:
		rt.panic("Cannot dynamically create Array types.")
	case reflect.Chan:
		return reflect.ChanOf(typ.ChanDir(), typ.Elem())
	case reflect.Func:
		rt.panic("Cannot dynamically create Function types.")
	case reflect.Interface:
		rt.panic("TODO")
	case reflect.Map:
		return reflect.MapOf(typ.Key(), typ.Elem())
	case reflect.Ptr:
		return reflect.PtrTo(rt.tysubst(typ.Elem()))
	case reflect.Slice:
		return reflect.SliceOf(rt.tysubst(typ.Elem()))
	case reflect.Struct:
		rt.panic("Cannot dynamically create Struct types.")
	case reflect.UnsafePointer:
		rt.panic("Cannot dynamically create unsafe.Pointer types.")
	}

	// We've covered all the composite types, so we're only left with
	// base types.
	return typ
}

func tyvarName(t reflect.Type) string {
	if !t.ConvertibleTo(tyvarUnderlyingType) {
		return ""
	}
	return t.Name()
}
