package main

import (
	"flag"
	"log"
	"os"
)

func (wc *WCount) setFlags() {
	ignoreFlag := flag.Bool("i", true, "ignore case distinctions")
	nDisplay := flag.Int("n", 20, "no of words to display")
	order := flag.Bool("asc", false, "display the result in an ascending order")
	sync := flag.Bool("sync", false, "whether to run the program sequentially")
	directory := flag.Bool("d", false, "whether to run in the directory mode")

	// parse the flag
	flag.Parse()

	// parse the arguments
	files := flag.Args()

	// validate if directory is set at least one argument is given
	if *directory == true && len(files) == 0 {
		log.Fatal("At least one directory should be specified")
	}

	// validate if directory is set, only one argument is given
	if *directory == true && len(files) != 1 {
		log.Fatal("Only one directory should be given")
	}

	wc.files = files
	wc.ignoreFlag = *ignoreFlag
	wc.nDisplay = *nDisplay
	wc.orderAsc = *order
	wc.sync = *sync
	wc.directory = *directory
	wc.directoryPath = files[0]

	// if the argument provided is not a directory, exit with an error
	info, err := os.Stat(wc.directoryPath)
	if wc.directory && (err != nil || !info.IsDir()) {
		log.Fatal("Provided path is not a directory")
	}
}
