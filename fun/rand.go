package fun

import (
	"math/rand"
	"time"

	"github.com/BurntSushi/ty"
)

var randNumGen *rand.Rand

func init() {
	randNumGen = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// ShuffleGen has a parametric type:
//
//	func ShuffleGen(xs []A, rng *rand.Rand)
//
// ShuffleGen shuffles `xs` in place using the given random number
// generator `rng`.
func ShuffleGen(xs interface{}, rng *rand.Rand) {
	chk := ty.Check(
		new(func([]ty.A, *rand.Rand)),
		xs, rng)
	vxs := chk.Args[0]

	// Implements the Fisher-Yates shuffle: http://goo.gl/Hb9vg
	xsLen := vxs.Len()
	swapper := swapperOf(vxs.Type().Elem())
	for i := xsLen - 1; i >= 1; i-- {
		j := rng.Intn(i + 1)
		swapper.swap(vxs.Index(i), vxs.Index(j))
	}
}

// Shuffle has a parametric type:
//
//	func Shuffle(xs []A)
//
// Shuffle shuffles `xs` in place using a default random number
// generator seeded once at program initialization.
func Shuffle(xs interface{}) {
	ShuffleGen(xs, randNumGen)
}
