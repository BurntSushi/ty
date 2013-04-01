package fun

import (
	"testing"
)

func TestShuffle(t *testing.T) {
	nums := Range(0, 100)
	Shuffle(nums)

	assertDeep(t, Set(nums), Set(Range(0, 100)))
}
