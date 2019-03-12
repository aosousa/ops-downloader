package main

import (
	"fmt"
	"net/url"

	"github.com/ChimeraCoder/anaconda"

	utils "github.com/aosousa/ops-downloader/utils"
)

func init() {
	// set up Config struct before continuing
	fmt.Println("Configuration file: Loading")
	initConfig()
	fmt.Println("Configuration file: OK")
}

func main() {
	api := anaconda.NewTwitterApiWithCredentials(config.AccessToken, config.AccessTokenSecret, config.APIKey, config.APISecret)

	// set URL params
	urlParams := url.Values{}
	urlParams.Set("screen_name", "OnePerfectShot")
	urlParams.Set("count", "30")

	// get tweets
	result, err := api.GetUserTimeline(urlParams)
	if err != nil {
		utils.HandleError(err)
	}

	for _, tweet := range result {
		fmt.Println(tweet.Entities.Media)
	}
}
