package main

import (
	"fmt"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

// TODO: https://docs.google.com/presentation/d/1KiU14z2owLUoiTYz5pKo-gP8RnP2BINmucVYJ6DfxTs/edit#slide=id.g4cc8086b3f_0_541
func main() {
	root := "./img"
	err := filepath.Walk(root,
		func(path string, info os.FileInfo, err error) error {
			fmt.Println(path)
			fmt.Println(info)

			ext := filepath.Ext(path)
			if ext != ".png" {
				return nil
			}

			pngFile, err := os.Open(path)
			if err != nil {
				return err
			}
			defer pngFile.Close()

			img, err := png.Decode(pngFile)
			if err != nil {
				return err
			}

			jpgFile, err := os.Create(strings.TrimSuffix(path, ext) + ".jpg")
			if err != nil {
				return err
			}
			defer jpgFile.Close()

			err = jpeg.Encode(jpgFile, img, &jpeg.Options{Quality: 100})
			if err != nil {
				return err
			}

			return nil
		})
	if err != nil {
		panic(err)
	}
}
