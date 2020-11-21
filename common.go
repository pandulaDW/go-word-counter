package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
)

// WCount is the type definition for both sync and concurrent count methods
type WCount struct {
	files      []string
	ignoreFlag bool
	nDisplay   int
	orderAsc   bool
	wcMap      map[string]int
}

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

func (wc *WCount) setFlags() {
	ignoreFlag := flag.Bool("i", false, "ignore case distinctions")
	nDisplay := flag.Int("n", 20, "No of words to display. By default top 20 words will be displayed")
	order := flag.Bool("asc", false, "display result ascending")

	// parse the flag
	flag.Parse()

	// parse the arguments
	files := flag.Args()

	wc.files = files
	wc.ignoreFlag = *ignoreFlag
	wc.nDisplay = *nDisplay
	wc.orderAsc = *order
}

func (wc *WCount) verifyFiles() {
	if len(wc.files) == 0 {
		log.Fatal("No files specified")
	}
	for _, file := range wc.files {
		info, err := os.Stat(file)
		if err != nil || info.IsDir() {
			log.Fatal(err)
		}
	}
}

func readFile(path string) *bufio.Reader {
	file, err := os.Open(path)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	rd := bufio.NewReader(file)
	return rd
}

func combineMaps(m1, m2 map[string]int) {
	for key, val := range m2 {
		m1[key] = m1[key] + val
	}
}
