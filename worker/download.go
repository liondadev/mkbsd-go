package worker

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const DownloadFolder = "./download"

// Download can be called to download the file to the ./download folder
func Download(url string, name string) error {
	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to make http get request: %v", err)
	}
	defer res.Body.Close()

	f, err := os.OpenFile(DownloadFolder+"/"+name, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		return fmt.Errorf("failed to open file %s: %v", DownloadFolder+"/"+name, err)
	}

	n, err := io.Copy(f, res.Body)
	if err != nil {
		return fmt.Errorf("failed to write to download folder: %v", err)
	}

	log.Printf("Downloaded wallpaper %s (%d bytes)\n", name, n)

	return nil
}
