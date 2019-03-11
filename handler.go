package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
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

/* Generates a base64 encoded HMAC SHA-256 hash created from the crcToken sent by Twitter
 * and the app's API secret key in order to validate the webhook. Receives:
 * secret (string) - App's API secret key
 * crcToken (string) - CRC token sent by Twitter
 *
 * Returns: base64 encoded HMAC SHA-256 hash (string)
 */
func generateHMACHash(secret, crcToken string) string {
	// create a new HMAC by defining the hash type and the key (as byte array)
	h := hmac.New(sha256.New, []byte(secret))

	// write data to it
	h.Write([]byte(crcToken))

	// get result and base64 encode it
	hash := base64.StdEncoding.EncodeToString([]byte(h.Sum(nil)))

	return "sha256=" + hash
}
