package fun

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

var (
	pf  = fmt.Printf
	rng *rand.Rand
)

func init() {
	rng = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func assertDeep(t *testing.T, v1, v2 interface{}) {
	if !reflect.DeepEqual(v1, v2) {
		t.Fatalf("%v != %v", v1, v2)
	}
}

func randIntSlice(size int) []int {
	slice := make([]int, size)
	for i := 0; i < size; i++ {
		slice[i] = rng.Intn(1000000)
	}
	return slice
}
