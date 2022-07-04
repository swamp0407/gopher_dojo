package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/swamp0407/gopher_dojo/kadai1/converter"
)

func main() {
	flag.Usage = usage

	flag.Parse()

	if len(os.Args) != 4 {
		flag.Usage()
		os.Exit(1)
	}
	if err := os.MkdirAll("output", 0755); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	ext_from := flag.Arg(0)
	ext_to := flag.Arg(1)
	dir := flag.Arg(2)
	if err := converter.Convert(dir, ext_from, ext_to); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println("Convertion completed.")
	os.Exit(0)

}

func usage() {
	fmt.Println("Usage: main <extension(from)> <extension(to)> <directory>")
}
