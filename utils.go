package main

import (
	"fmt"
	"flag"
	"os"
)

type Params struct {
	URL string
	wl 	string
	nThreads int
}

func ParseOptions(args [] string) (PARAMS Params){
	if( len(args) < 3){
		fmt.Printf("Usage %s -u <url> -w <wordlist> [-t threads]\n", args[0])
		os.Exit(1)
	}
	flag.StringVar(&PARAMS.URL, "u", "", "URL to FUZZ")
	flag.StringVar(&PARAMS.wl, "w", "", "wordlist")
	flag.IntVar(&PARAMS.nThreads, "t", 20, "Threads for concurrency")
	flag.Parse()
	return

}