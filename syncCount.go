package main

import (
	"strings"
)

func (wc *WCount) syncCount(path string) map[string]int {
	rd := readFile(path)
	wcMap := map[string]int{}

	for {
		line, err := rd.ReadString('\n')
		if err != nil {
			break
		}
		words := strings.Split(line, " ")
		for _, val := range words {
			if wc.ignoreFlag {
				val = strings.TrimSpace(val)
			} else {
				val = strings.ToLower(strings.TrimSpace(val))
			}
			val = strings.ToLower(strings.TrimSpace(val))
			wcMap[val]++
		}
	}

	return wcMap
}

func (wc *WCount) runSyncCount() {
	cumMap := map[string]int{}

	for _, file := range wc.files {
		wcMap := wc.syncCount(file)
		combineMaps(cumMap, wcMap)
	}

	wc.wcMap = cumMap
}
