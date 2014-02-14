package main

import (
	"fmt"
	//this import actually has a flags package but the go Idea plugin parser didn't like it because the end of the import
	//is usually the name of the package to be used. To make the red highlights go away I just aliased the import.
	flags "github.com/jessevdk/go-flags"
)

func main() {

	fmt.Println("GMF example program.\n")

	var opts struct {
	}

	parser := flags.NewParser(&opts, flags.Default)
	parser.Usage = "[OPTIONS]\n\n  examples:\n  \n"
	parser.Parse()
}
