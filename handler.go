package main

import (
	models "github.com/aosousa/ops-downloader/models"
)

var config models.Config

func initConfig() {
	config = models.CreateConfig()
}
