package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	wc := WCount{}
	wc.setFlags()
	wc.verifyFiles()

	if wc.sync {
		fmt.Println("running synchronized version")
		wc.runSyncCount()
	} else {
		fmt.Println("running concurrent version")
		wc.runConCount()
	}

	fmt.Printf("Process took %s\n", time.Since(start))
	wc.display()
}
