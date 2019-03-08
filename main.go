package main

import (
	"fmt"
	"net/http"
)

func init() {
	// set up Config struct before continuing
	fmt.Println("Configuration file: Loading")
	initConfig()
	fmt.Println("Configuration file: OK")
}

func main() {
	/*fileURL := "https://golangcode.com/images/avatar.jpg"

	if err := DownloadImage("test.jpg", fileURL); err != nil {
		panic(err)
	}*/

	router := NewRouter()
	fmt.Println("Serving on port 8081")
	http.ListenAndServe(":8081", router)
}
