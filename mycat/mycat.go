package main

import (
	"flag"
	"fmt"
)

func main() {
	args := flag.Args()
	fmt.Println(args)
	fmt.Println("Hello")
}
