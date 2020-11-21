package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/gabriel-vasile/mimetype"
)

// WCount is the type definition for both sync and concurrent count methods
type WCount struct {
	files         []string
	ignoreFlag    bool
	nDisplay      int
	orderAsc      bool
	sync          bool
	directory     bool
	directoryPath string
	wcMap         map[string]int
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

func isTextFile(path string) bool {
	mimeType, _ := mimetype.DetectFile(path)
	if strings.HasPrefix(mimeType.String(), "text") {
		return true
	}
	return false
}

func (wc *WCount) registerTextFiles() {
	walkFunction := func(path string, info os.FileInfo, err error) error {
		if isTextFile(path) {
			wc.files = append(wc.files, path)
		}
		if err != nil {
			log.Fatal(err)
		}
		return err
	}
	filepath.Walk(wc.directoryPath, walkFunction)
	fmt.Printf("Parsing %d text files\n", len(wc.files))
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
