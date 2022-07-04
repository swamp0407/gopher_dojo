package converter

import (
	"fmt"
	"image"
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	log.Println("Do stuff BEFORE the tests!")
	if err := os.MkdirAll("output", 0777); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	exitVal := m.Run()
	log.Println("Do stuff AFTER the tests!")
	if err := os.RemoveAll("output"); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	os.Exit(exitVal)
}

func Test_image2png(t *testing.T) {
	type args struct {
		img image.Image
	}
	tests := []struct {
		name         string
		args         args
		wantSavefile string
		wantErr      bool
	}{
		// TODO: Add test cases.
		{"generate test.png", args{image.NewRGBA(image.Rect(0, 0, 100, 100))}, "./output/test.png", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			savefile, err := os.Create(tt.wantSavefile)
			if err != nil {
				t.Errorf("failed to open a test.png")
				return
			}
			if err := image2png(tt.args.img, savefile); (err != nil) != tt.wantErr {
				t.Errorf("image2png() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_image2jpg(t *testing.T) {
	type args struct {
		img image.Image
	}
	tests := []struct {
		name         string
		args         args
		wantSavefile string
		wantErr      bool
	}{
		// TODO: Add test cases.
		{"generate test.jpg", args{image.NewRGBA(image.Rect(0, 0, 100, 100))}, "./output/test.jpg", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			savefile, err := os.Create(tt.wantSavefile)
			if err != nil {
				t.Errorf("failed to open a test.jpg,%s", err)
				return
			}
			if err := image2jpg(tt.args.img, savefile); (err != nil) != tt.wantErr {
				t.Errorf("image2jpg() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_replaceExt(t *testing.T) {
	type args struct {
		filePath string
		from     string
		to       string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := replaceExt(tt.args.filePath, tt.args.from, tt.args.to); got != tt.want {
				t.Errorf("replaceExt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvert(t *testing.T) {
	type args struct {
		dir      string
		ext_from string
		ext_to   string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Convert(tt.args.dir, tt.args.ext_from, tt.args.ext_to); (err != nil) != tt.wantErr {
				t.Errorf("Convert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
