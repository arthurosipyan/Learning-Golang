package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// Open file from argument
	openedFile, err := os.Open(os.Args[1])

	// Read file
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, openedFile)
}
