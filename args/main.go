package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	commandLineArgs := os.Args[1:]
	abcd, err := os.Open(commandLineArgs[0])
	if err != nil {
		fmt.Printf("Error is : %v", err)
	}
	io.Copy(os.Stdout, abcd)
}
