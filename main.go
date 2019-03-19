package main

import (
	"fmt"
	"net/url"
	"path/filepath"
	"strconv"

	"github.com/ChimeraCoder/anaconda"

	utils "github.com/aosousa/ops-downloader/utils"
)

func init() {
	// set up Config struct before continuing
	fmt.Println("Configuration file: Loading")
	initConfig()
	fmt.Println("Configuration file: OK")
	fmt.Println()
}

func main() {
	var fileURL, fileName, filePath string

	api := anaconda.NewTwitterApiWithCredentials(config.AccessToken, config.AccessTokenSecret, config.APIKey, config.APISecret)

	// set URL params
	urlParams := url.Values{}
	urlParams.Set("screen_name", "OnePerfectShot")
	urlParams.Set("since_id", config.SinceID)

	// get tweets
	tweets, err := api.GetUserTimeline(urlParams)
	if err != nil {
		utils.HandleError(err)
	}

	if len(tweets) > 0 {
		for _, tweet := range tweets {
			if len(tweet.Entities.Media) > 0 {
				fileURL = tweet.Entities.Media[0].Media_url
				fileName = filepath.Base(fileURL)
				filePath = config.OutputFolder + fileName

				err = downloadImage(filePath, fileURL)
				if err == nil {
					fmt.Printf("Downloaded file: %s\n", fileName)
				} else {
					fmt.Printf("An error occurred while trying to download file %s: %s\n", fileName, err)
				}
			}
		}

		// save last tweet ID
		config.SinceID = strconv.Itoa(int(tweets[0].Id))
		config.Save()
	} else {
		fmt.Println("No new tweets.")
	}
}
