package main

import (
	"io"
	"net/http"
	"os"

	models "github.com/aosousa/ops-downloader/models"
)

var config models.Config

func initConfig() {
	config = models.CreateConfig()
}

// DownloadImage will download a url to a local file.
// https://golangcode.com/download-a-file-from-a-url/
func downloadImage(filepath string, url string) error {
	// get the image
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// check if directory exists, create it if it doesn't
	if _, err := os.Stat(config.OutputFolder); os.IsNotExist(err) {
		os.Mkdir(config.OutputFolder, os.ModeDir)
	}

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
