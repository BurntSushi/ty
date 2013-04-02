/*
Package ty provides utilities for writing type parametric functions with run
time safety.

Go tip/1.1 is required

The very foundation of this package only recently became possible with the
addition of 3 new functions in the standard library `reflect` package:
SliceOf, MapOf and ChanOf. In particular, it provides the ability to
dynamically construct types at run time from component types.

Further extensions to this package can be made if similar functions are added
for structs and functions(?).
*/
package ty
