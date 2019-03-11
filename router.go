package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	models "github.com/aosousa/ops-downloader/models"
)

/*NewRouter creates a new mux outer with the routes defined
 * in the method below
 */
func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/webhooks/twitter", handleWebhook).Methods("GET", "POST")
	router.HandleFunc("/test", test).Methods("GET", "POST")

	return router
}

func handleWebhook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CALLED WEBHOOK")
	crcToken := mux.Vars(r)["crc_token"]

	hash := generateHMACHash(config.APISecret, crcToken)
	setResponse(w, models.TwitterResponse{ResponseToken: hash})
	return
}

func test(w http.ResponseWriter, r *http.Request) {
	setResponse(w, models.HTTPResponse{Success: true, Data: "test"})
	return
}

/*SetResponse sets the ResponseWriter's headers and body to send in a JSON response. Receives:
 * w (http.ResponseWriter) - ResponseWriter struct
 * body (interface{}) - Body of the response (can be a HTTPResponse or TwitterResponse struct)
 */
func setResponse(w http.ResponseWriter, body interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(body)
}
