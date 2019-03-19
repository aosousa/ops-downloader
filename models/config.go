package models

import (
	"encoding/json"
	"io/ioutil"
	"os"

	utils "github.com/aosousa/ops-downloader/utils"
)

/*Config struct contains all the necessary configurations for the application to run. Contains:
 * APIKey (string) - Twitter application API key
 * APISecret (string) - Twitter application API secret
 * AccessToken (string) - Twitter application access token
 * AccessTokenSecret (string) - Twitter application secret token secret
 * SinceID (string) - Last Tweet ID returned in the search
 * OutputFolder (string) - Path where the images should be stored
 */
type Config struct {
	APIKey            string `json:"apiKey"`
	APISecret         string `json:"apiSecret"`
	AccessToken       string `json:"accessToken"`
	AccessTokenSecret string `json:"accessTokenSecret"`
	SinceID           string `json:"sinceId"`
	OutputFolder      string `json:"outputFolder"`
}

/*CreateConfig adds information from a configuration file to a Config struct. Returns:
 * Config - Configuration struct
 */
func CreateConfig() Config {
	var config Config
	jsonFile, err := ioutil.ReadFile("./config.json")
	if err != nil {
		utils.HandleError(err)
		os.Exit(1)
	}

	err = json.Unmarshal(jsonFile, &config)
	if err != nil {
		utils.HandleError(err)
		os.Exit(1)
	}

	return config
}

// Save saves the Config struct into config.json file
func (c Config) Save() {
	file, _ := json.MarshalIndent(c, "", " ")
	_ = ioutil.WriteFile("config.json", file, 0644)
}
