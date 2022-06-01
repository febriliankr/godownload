package godownload

import "testing"

func TestDownload(t *testing.T) {
	_, err := Download("https://www.google.com/images/branding/googlelogo/2x/googlelogo_color_272x92dp.png", "tmp/downloads", 100)

	if err != nil {
		t.Error(err)
	}
}
func TestDownloadLarge(t *testing.T) {
	_, err := Download("https://www.google.com/images/branding/googlelogo/2x/googlelogo_color_272x92dp.png", "tmp/downloads", 100000000)

	if err != nil {
		t.Error(err)
	}
}
