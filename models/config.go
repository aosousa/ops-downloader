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
 */
type Config struct {
	APIKey            string `json:"apiKey"`
	APISecret         string `json:"apiSecret"`
	AccessToken       string `json:"accessToken"`
	AccessTokenSecret string `json:"accessTokenSecret"`
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
