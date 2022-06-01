package godownload

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type DownloadToPathResponse struct {
	Size     int64
	FilePath string
}

// Download To Path accepts "fullFileURL" as the download url and "downloadToPath" as where the file will be downloaded to. The "downloadToPath" parameter is in this format "tmp/downloads", without any '/' at the end. Use DownloadToPath(fullFileURL, ".") to download to the root directory.
func DownloadToPath(fileURL string, downloadToPath string) (DownloadToPathResponse, error) {
	var response DownloadToPathResponse
	parsedURL, err := url.Parse(fileURL)
	if err != nil {
		return response, err
	}
	path := parsedURL.Path
	segments := strings.Split(path, "/")
	fileName := segments[len(segments)-1]

	destinationFilePath := fmt.Sprintf("%s/%s", downloadToPath, fileName)

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

	response = DownloadToPathResponse{
		Size:     size,
		FilePath: destinationFilePath,
	}
	return response, nil
}
