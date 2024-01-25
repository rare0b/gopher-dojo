package main

import (
	"flag"
	"fmt"
	"os"
)

func init() {
	flag.String("n", "", "files")
}

func main() {
	flag.Parse()

	flagArgs := flag.Args()
	osArgs := os.Args
	fmt.Println(flagArgs)
	fmt.Println(osArgs)
	fmt.Println("Hello")
}
