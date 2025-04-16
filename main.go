package main

import (
	"flag"
	"fmt"
)

func main() {
	modePtr := flag.String("mode", "", "Mode to run in")
	persistencePtr := flag.Bool("persist", false, "Whether to run tool with persistence mode")
	flag.Parse()
	fmt.Println("Mode:", *modePtr)
	fmt.Println("Persistence:", *persistencePtr)
}
