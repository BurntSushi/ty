package fun

import (
	"fmt"
)

func ExamplePrimeNumberSieve() {
	bound := 100

	primes := Foldr(
		func(n int, sieve []int) []int {
			return Filter(func(n2 int) bool {
				return n2 == n || n2%n != 0
			}, sieve).([]int)
		},
		Range(2, bound), Range(2, bound))

	fmt.Printf("%v\n", primes)
	// Output:
	// [2 3 5 7 11 13 17 19 23 29 31 37 41 43 47 53 59 61 67 71 73 79 83 89 97]
}
