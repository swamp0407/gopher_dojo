package converter

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"

	"io"
	"os"
	"path/filepath"
)

func image2png(img image.Image, savefile io.Writer) error {
	png := png.Encode(savefile, img)
	return png
}

func image2jpg(img image.Image, savefile io.Writer) error {
	jpg := jpeg.Encode(savefile, img, &jpeg.Options{Quality: 100})
	return jpg
}

func replaceExt(filePath, from, to string) string {
	ext := filepath.Ext(filePath)
	if len(from) > 0 && ext != "."+from {
		return filePath
	}
	return filePath[:len(filePath)-len(ext)] + "." + to
}

func Convert(dir, ext_from, ext_to string) error {
	fmt.Println("Converting...")
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) == "."+ext_from {
			fmt.Printf("%s\n", path)

			f, err := os.Open(path)
			if err != nil {
				return err
			}
			defer f.Close()
			img, _, err := image.Decode(f)
			if err != nil {
				return err
			}
			dstfileName := replaceExt(filepath.Base(path), ext_from, ext_to)
			dstfile, err := os.Create(filepath.Join("output", dstfileName))

			if err != nil {
				return err
			}
			defer dstfile.Close()

			if ext_to == "png" {
				if err := image2png(img, dstfile); err != nil {
					return err
				}
			}
			if ext_to == "jpg" || ext_to == "jpeg" {
				if err := image2jpg(img, dstfile); err != nil {
					return err
				}
			}
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	return nil
}
