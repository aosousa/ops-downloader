package models

/*HTTPResponse is the struct used to send back JSON responses in the API. Contains:
 * Success (bool) - Whether the request was successful or not
 * Data (interface{}) - Interface that has the relevant information to return. Can be a slice, struct, string, etc.
 * Error (string) - Error message in case an error occurred
 */
type HTTPResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Error   string      `json:"error"`
}

/*TwitterResponse is the struct used to send back JSON responses for Twitter's
 * CRC requests every hour for webhook validation. Contains:
 * ResponseToken (string) - The base64 encoded HMAC SHA-256 hash required for the webhook validation
 */
type TwitterResponse struct {
	ResponseToken string `json:"response_token"`
}
