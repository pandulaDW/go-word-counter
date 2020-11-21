package main

import (
	"fmt"
	"math"
	"sort"
)

func (wc *WCount) display() {
	type kv struct {
		Key   string
		Value int
	}

	var ss []kv
	for k, v := range wc.wcMap {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		if wc.orderAsc {
			return ss[i].Value < ss[j].Value
		}
		return ss[i].Value > ss[j].Value
	})

	nDisplay := math.Min(float64(len(wc.wcMap)), float64(wc.nDisplay))
	for i := 0; i < int(nDisplay); i++ {
		fmt.Printf("%s -> %d\n", ss[i].Key, ss[i].Value)
	}
}
