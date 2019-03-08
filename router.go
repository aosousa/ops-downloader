package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	models "github.com/aosousa/ops-downloader/models"
)

/*NewRouter creates a new mux outer with the routes defined
 * in the method below
 */
func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/webhooks/twitter", handleWebhook).Methods("GET")
	router.HandleFunc("/test", test).Methods("GET", "POST")

	return router
}

func handleWebhook(w http.ResponseWriter, r *http.Request) {

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
