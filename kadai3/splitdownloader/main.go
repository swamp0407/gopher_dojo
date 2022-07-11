package main

import (
	"fmt"
	"os"

	"github.com/swamp0407/gopher_dojo/kadai3/splitdownloader/downloader"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
	}()
	if len(os.Args) < 2 {
		fmt.Println("Usage: splitdownloader <url>")
		os.Exit(1)
	}

	url := os.Args[1]

	cli := downloader.New(url)
	os.Exit(cli.Run())
}
