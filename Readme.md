# Godownload

Simple library to download file from url and write them to your destined local directory.

## Installation

Install as a library

`go get github.com/febriliankr/godownload`

## Usage

```
var url = "https://www.google.com/images/branding/googlelogo/2x/googlelogo_color_272x92dp.png"
var folderPath = "tmp"
res, err := godownload.Download(url, folderPath)

log.Println("File downloaded! Size:", res.Size, "Path:", res.FilePath)
```
