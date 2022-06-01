package godownload

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/febriliankr/godownload/entities"
)

// Download To Path accepts "fullFileURL" as the download url and "downloadToPath" as where the file will be downloaded to. The "downloadToPath" parameter is in this format "tmp/downloads", without any '/' at the end. Use Download(fullFileURL, ".", ) to download to the root directory.
func Download(fileURL string, downloadToPath string, maxFileSizeInMB int64) (entities.DownloadToPathResponse, error) {
	var response entities.DownloadToPathResponse

	parsedURL, err := url.Parse(fileURL)
	if err != nil {
		return response, err
	}
	path := parsedURL.Path
	segments := strings.Split(path, "/")
	fileName := segments[len(segments)-1]

	destinationFilePath := fmt.Sprintf("%s/%s", downloadToPath, fileName)

	headSize, err := getContentLength(fileURL)

	if headSize > maxFileSizeInMB*1000*1000 {
		return response, errors.New("file size is bigger than maximum file size")
	}

	if err != nil {
		return response, err
	}

	log.Println(headSize)

	// Create a blank file
	file, err := os.Create(destinationFilePath)
	if err != nil {
		return response, err
	}

	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}

	// Put content on file
	resp, err := client.Get(fileURL)
	if err != nil {
		return response, err
	}

	defer resp.Body.Close()

	size, err := io.Copy(file, resp.Body)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	response.Size = size
	response.FilePath = destinationFilePath

	return response, nil
}

func getContentLength(url string) (int64, error) {

	var contentLength int64
	res, err := http.Head(url)
	if err != nil {
		return contentLength, err
	}

	contentLength = res.ContentLength
	return contentLength, nil
}
