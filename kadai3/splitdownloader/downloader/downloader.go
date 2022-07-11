package downloader

import (
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
	"strconv"
	"sync"

	"golang.org/x/sync/errgroup"
)

type Downloader struct {
	url string
}

func New(url string) *Downloader {
	return &Downloader{url: url}
}

var wg sync.WaitGroup
var proc = 4

var TempDir = "./tmp"

func (d *Downloader) Run() int {
	if err := os.MkdirAll(TempDir, os.ModePerm); err != nil {
		panic(err)
	}
	len, err := getContentLength(d.url)
	if err != nil {
		fmt.Println(err)
		return 1
	}
	eg, ctx := errgroup.WithContext(context.Background())
	for i := 0; i < proc; i++ {
		i := i
		from := i * len / proc
		to := (i + 1) * len / proc
		if i == proc-1 {
			to = len
		}
		eg.Go(func() error {
			return d.rangeRequest(ctx, from, to, i)
		})
	}
	if err := eg.Wait(); err != nil {
		fmt.Println(err)
		return 1
	}
	if err := d.merge(); err != nil {
		fmt.Println(err)
		return 1
	}

	return 0
}

func (d *Downloader) rangeRequest(ctx context.Context, from, to, chunk int) error {
	defer wg.Done()
	req, err := http.NewRequest("GET", d.url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", from, to))
	req = req.WithContext(ctx)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusPartialContent {
		return errors.New("status code is not 206")
	}
	file, err := os.Create(fmt.Sprintf("%s/%d.tmp", TempDir, chunk))
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}
	return nil
}

func (d *Downloader) merge() error {
	files, err := ioutil.ReadDir(TempDir)
	if err != nil {
		return err
	}
	sort.Slice(files, func(i, j int) bool {
		a, _ := strconv.Atoi(files[i].Name())
		b, _ := strconv.Atoi(files[j].Name())
		return a < b
	})
	mergedFile, err := os.Create(fmt.Sprintf("%s/%s", TempDir, "merged.tmp"))
	if err != nil {
		return err
	}
	for _, file := range files {
		f, err := os.Open(file.Name())
		if err != nil {
			return err
		}
		io.Copy(mergedFile, f)
		f.Close()
	}
	mergedFile.Close()

	return nil
}

func getContentLength(url string) (int, error) {
	resp, err := http.Head(url)
	if err != nil {
		return 0, errors.New("failed to access url: " + url + err.Error())
		// return 0, errors.Wrap(err, "failed to access url: "+url)
	}

	if resp.Header.Get("Accept-Ranges") != "bytes" {
		return 0, errors.New("this site does not support a range request")
	}
	len, err := strconv.Atoi(resp.Header.Get("Content-Length"))
	if err != nil {
		return 0, errors.New("getContentLength: " + err.Error())
	}

	return len, nil
}
