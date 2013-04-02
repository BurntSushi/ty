/*
Package fun provides type parametric utility functions for lists, sets,
channels and maps.

The major contribution of this package is a set of functions that operate
on values without depending on their types while maintaining type safety at
run time using the `reflect` package.

The use cases of this package still aren't clear, but there are two primary
concerns: the loss of compile time type safety and performance. In particular,
with regard to performance, most functions here are much slower than their
built-in counter parts. However, there are a couple where the overhead of
reflection is relatively insignificant: AsyncChan and ParMap.

In terms of code structure and organization, the price is mostly paid inside
of the package due to the annoyances of operating with `reflect`. The caller
usually only has one obligation other than to provide values consistent with
the type of the function: type assert the result to the desired type.

When the caller provides values that are inconsistent with the parametric type
of the function, the function will panic with a `TypeError`. (Either because
the types cannot be unified or because they cannot be constructed due to
limitations of the `reflect` package.)

Examples

Squaring each integer in a slice:

	square := func(x int) int { return x * x }
	nums := []int{1, 2, 3, 4, 5}
	squares := Map(square, nums).([]int)

Reversing any slice:

	slice := []string{"a", "b", "c"}
	reversed := Reverse(slice).([]string)

Sorting any slice:

	// Sort a slice of structs with first class functions.
	type Album struct {
		Title string
		Year int
	}
	albums := []Album{
		{"Born to Run", 1975},
		{"WIESS",       1973},
		{"Darkness",    1978},
		{"Greetings",   1973},
	}
	sorted := QuickSort(
		func(a, b Album) bool { return a.Year < b.Year },
		albums).([]Album)

Parallel map:

	// Compute the prime factorization in parallel
	// for every integer in [1000, 10000].
	primeFactors := func(n int) []int { // compute prime factors }
	factors := ParMap(primeFactors, Range(1000, 10001)).([]int)

Asynchronous channel without a fixed size buffer:

	s, r := AsyncChan(new(chan int))
	send, recv := s.(chan<- int), r.(<-chan int)

	// Send as much as you want.
	for i := 0; i < 100; i++ {
		s <- i
	}
	close(s)
	for i := range recv {
		// do something with `i`
	}

*/
package fun
