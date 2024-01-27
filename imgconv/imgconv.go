package main

import (
	"os"
	"path/filepath"
)

func main() {
	root := "."
	err := filepath.Walk(root,
		func(path string, info os.FileInfo, err error) error {
			ext := filepath.Ext(path)
			if ext == ".png" {
				//TODO
			}
			return nil
		})
	if err != nil {
		panic(err)
	}
}
