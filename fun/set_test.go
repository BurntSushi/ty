package fun

import (
	"testing"
)

func TestUnion(t *testing.T) {
	a := map[string]bool{
		"springsteen": true,
		"jgeils":      true,
		"seger":       true,
		"metallica":   true,
	}
	b := map[string]bool{
		"metallica": true,
		"chesney":   true,
		"mcgraw":    true,
		"cash":      true,
	}
	c := Union(a, b).(map[string]bool)

	assertDeep(t, c, map[string]bool{
		"springsteen": true,
		"jgeils":      true,
		"seger":       true,
		"metallica":   true,
		"chesney":     true,
		"mcgraw":      true,
		"cash":        true,
	})
}

func TestIntersection(t *testing.T) {
	a := map[string]bool{
		"springsteen": true,
		"jgeils":      true,
		"seger":       true,
		"metallica":   true,
	}
	b := map[string]bool{
		"metallica": true,
		"chesney":   true,
		"mcgraw":    true,
		"cash":      true,
	}
	c := Intersection(a, b).(map[string]bool)

	assertDeep(t, c, map[string]bool{
		"metallica": true,
	})
}
