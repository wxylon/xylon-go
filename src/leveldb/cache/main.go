package main

import (
	"fmt"
)

func main() {
	cases := []struct {
		key   uint64
		value string
	}{
		{1, "a"},
		{2, "b"},
		{3, "c"},
		{4, "d"},
	}

	setfin := 0
	c := NewLRUCache(100)
	ns := c.GetNamespace(0)

	for i, x := range cases {
		set(ns, x.key, x.value, len(x.value), func() {
			setfin++
		}).Release()

		for j, y := range cases {
			r, ok := ns.Get(y.key, nil)
			if j <= i {
				// should hit
				if !ok {
					fmt.Errorf("case '%d' iteration '%d' is miss", i, j)
				} else if r.Value().(string) != y.value {
					fmt.Errorf("case '%d' iteration '%d' has invalid value got '%s', want '%s'", i, j, r.Value().(string), y.value)
				}
			} else {
				// should miss
				if ok {
					fmt.Errorf("case '%d' iteration '%d' is hit , value '%s'", i, j, r.Value().(string))
				}
			}
			if ok {
				r.Release()
			}
		}
	}

	//fmt.Println(cases, c, ns)
}

func set(ns Namespace, key uint64, value interface{}, charge int, fin func()) Object {
	obj, _ := ns.Get(key, func() (bool, interface{}, int, SetFin) {
		return true, value, charge, fin
	})
	return obj
}
