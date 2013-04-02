package exp

import (
	"fmt"
	"reflect"
	"testing"
)

var pf = fmt.Printf

func TestOrdMap(t *testing.T) {
	omap := OrderedMap(new(string), new(int))

	omap.Put("kaitlyn", 24)
	omap.Put("andrew", 25)
	omap.Put("lauren", 20)
	omap.Put("jen", 24)
	omap.Put("brennan", 25)

	omap.Delete("kaitlyn")

	assertDeep(t, omap.Keys(), []string{"andrew", "lauren", "jen", "brennan"})
	assertDeep(t, omap.Values(), []int{25, 20, 24, 25})
}

func assertDeep(t *testing.T, v1, v2 interface{}) {
	if !reflect.DeepEqual(v1, v2) {
		t.Fatalf("%v != %v", v1, v2)
	}
}
