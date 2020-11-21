package main

func main() {
	wc := WCount{}
	wc.setFlags()
	wc.verifyFiles()
	wc.runSyncCount()
	wc.display()
}
