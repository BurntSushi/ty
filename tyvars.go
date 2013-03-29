package ty

import (
	"reflect"
)

// TypeVariable is the underlying type of every type variable used in
// parametric types. It should not be used directly. Instead, use
//
//	type myOwnTypeVariable TypeVariable
//
// to create your own type variable. For your convenience, this package
// defines type variables for each letter in the range A-Z.
type TypeVariable struct {
	noImitation struct{}
}

// tyvarUnderlyingType is used to discover types that are type variables.
// Namely, any type variable must be convertible to `TypeVariable`.
var tyvarUnderlyingType = reflect.TypeOf(TypeVariable{})

type A TypeVariable
type B TypeVariable
type C TypeVariable
type D TypeVariable
type E TypeVariable
type F TypeVariable
type G TypeVariable
type H TypeVariable
type I TypeVariable
type J TypeVariable
type K TypeVariable
type L TypeVariable
type M TypeVariable
type N TypeVariable
type O TypeVariable
type P TypeVariable
type Q TypeVariable
type R TypeVariable
type S TypeVariable
type T TypeVariable
type U TypeVariable
type V TypeVariable
type W TypeVariable
type X TypeVariable
type Y TypeVariable
type Z TypeVariable
