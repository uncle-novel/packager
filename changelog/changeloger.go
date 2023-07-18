package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/parkr/changelog"
)

func main() {
	// Read options
	var filename string
	flag.StringVar(&filename, "file", "", "The path to your changelog")
	var output string
	flag.StringVar(&output, "out", "", "Where to write the changelog")
	var verbose bool
	flag.BoolVar(&verbose, "v", false, "Whether to print verbose output")
	flag.Parse()

	changelog.SetVerbose(verbose)

	// Find History.markdown
	if filename == "" {
		filename = changelog.HistoryFilename()
	}

	changes, err := changelog.NewChangelogFromFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println(changes.Versions[0].String())
	// 写出
	err = os.WriteFile(output, []byte(changes.Versions[0].String()), 0644)
}
