package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var n bool

func init() {
	flag.BoolVar(&n, "n", false, "-nが指定されたか")
}

func readFile(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return []string{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return []string{}, err
	}

	return lines, nil
}

func main() {
	flag.Parse()
	args := flag.Args()

	var allLines []string
	for _, file := range args {
		lines, err := readFile(file)
		if err != nil {
			panic(err)
		}

		allLines = append(allLines, lines...)
	}

	if n {
		for i, line := range allLines {
			fmt.Printf("%d: %s\n", i+1, line)
		}
	} else {
		for _, line := range allLines {
			fmt.Printf("%s\n", line)
		}
	}
}
