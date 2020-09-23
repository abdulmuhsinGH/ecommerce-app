package format

import (
	"encoding/json"
	"net/http"
	"os"
)

/*
Resp interface for response structure
*/
type Resp map[string]interface{}

/*
Send sends response of a request
*/
func Send(response http.ResponseWriter, status int, data Resp) {
	response.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	response.Header().Set("Access-Control-Allow-Origin", os.Getenv("RESOURCE_ALLOWED_ORIGIN"))
	response.WriteHeader(status)
	json.NewEncoder(response).Encode(data)
}

/*
Message format for response
*/
func Message(status bool, message string, data interface{}) Resp {
	return Resp{"status": status, "message": message, "data": data}
}
