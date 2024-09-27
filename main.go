package main

import (
	"encoding/json"
	"fmt"
	"github.com/liondadev/mkbsd-go/worker"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
	"sync"
)

const Url = "https://storage.googleapis.com/panels-api/data/20240916/media-1a-i-p~s"

type resp struct {
	Version int                       `json:"version"`
	Data    map[int]map[string]string `json:"data"`
}

func main() {
	// Make sure the downloads folder exists
	_, err := os.Stat(worker.DownloadFolder)
	if err != nil {
		if os.IsNotExist(err) {
			log.Println("Downloads folder doesn't exist, attempting creation...")
			err := os.Mkdir(worker.DownloadFolder, 0777)
			if err != nil {
				panic(err)
			}
		} else {
			panic(err)
		}
	}

	res, err := http.Get(Url)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var pBody resp
	if err := json.Unmarshal(body, &pBody); err != nil {
		panic(err)
	}

	log.Printf("Got API Response. (Version: %d)\n", pBody.Version)

	var wg sync.WaitGroup
	wg.Add(len(pBody.Data))
	for id, d := range pBody.Data {
		go func() {
			defer wg.Done()

			hd, ok := d["dhd"]
			if !ok {
				return
			}

			url, err := url.Parse(hd)
			if err != nil {
				fmt.Printf("failed: %v\n", err)
				return
			}

			ext := path.Ext(url.Path)
			err = worker.Download(hd, strconv.Itoa(id)+ext)
			if err != nil {
				fmt.Printf("Failed to download wallpaper %d: %v\n", id, err)
			}
		}()
	}
	wg.Wait()
	log.Println("Finished downloading!")
}
