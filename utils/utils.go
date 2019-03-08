package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// HandleError prints out an error that occurred
func HandleError(err error) {
	fmt.Printf("%s", err)
}

// DownloadImage will download a url to a local file.
// https://golangcode.com/download-a-file-from-a-url/
func DownloadImage(filepath string, url string) error {
	// get the image
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// create the file
	output, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer output.Close()

	// write the body to file
	_, err = io.Copy(output, res.Body)
	return err
}
