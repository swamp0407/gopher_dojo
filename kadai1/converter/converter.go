package converter

import (
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func image2png(img image.Image, savefile io.Writer) error {
	png := png.Encode(savefile, img)
	return png
}

func Convert() error {
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) == ".jpg" {
			j, err := os.Open(path)
			if err != nil {
				return err
			}
			defer j.Close()
			img, err := jpeg.Decode(j)
			if err != nil {
				return err
			}
			outf := filepath.Join(filepath.Dir(path), strings.TrimRight(filepath.Base(path), ".jpg")+".png")
			savefile, err := os.Create(outf)
			if err != nil {
				return err
			}
			err = image2png(img, savefile)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	return nil
}
