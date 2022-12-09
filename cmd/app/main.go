package main

import (
	"fetcher/cmd/app/fetch"
	"flag"
	"fmt"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("No urls provided")
		os.Exit(1)
	}

	showMetadata := flag.Bool("metadata", false, "Load pages metadata")
	loadAssets := flag.Bool("assets", false, "Load assets files")
	flag.Parse()

	fetch.ProcessUrls(flag.Args(), *showMetadata, *loadAssets)

}
