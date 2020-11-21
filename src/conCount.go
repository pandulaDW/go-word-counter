package main

import (
	"sync"
)

// forking a goroutine for each file and synchronizing them using mutexes
func (wc *WCount) runConCount() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	wcMaps := make([]map[string]int, 0, len(wc.files))

	for _, file := range wc.files {
		wg.Add(1)
		go func(file string) {
			wcMap := wc.syncCount(file)
			mu.Lock()
			wcMaps = append(wcMaps, wcMap)
			defer mu.Unlock()
			defer wg.Done()
		}(file)
	}

	wg.Wait()

	cumMap := map[string]int{}
	for _, wcMap := range wcMaps {
		combineMaps(cumMap, wcMap)
	}

	wc.wcMap = cumMap
}
