package main

import (
	"fmt"
	"net/url"
	"strconv"

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
	urlParams.Set("since_id", config.SinceID)

	// get tweets
	result, err := api.GetUserTimeline(urlParams)
	if err != nil {
		utils.HandleError(err)
	}

	for _, tweet := range result {
		fmt.Println(tweet.Entities.Media)
	}

	if len(result) > 0 {
		config.SinceID = strconv.Itoa(int(result[0].Id))
		config.Save()
	}
}
