package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func getFileSize(url string) (int, error) {
	resp, err := http.Head(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	contentLength := resp.Header.Get("Content-Length")
	if contentLength == "" {
		return 0, fmt.Errorf("Content-Length header is missing")
	}

	size, err := strconv.Atoi(contentLength)
	if err != nil {
		return 0, err
	}

	return size, nil
}

func main() {
	url := "https://raw.githubusercontent.com/golang/go/master/README.md"

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	out, err := os.Create("a.out")
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		panic(err)
	}
}
