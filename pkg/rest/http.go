package rest

/*
	This is a helper to parse response to content-type json
*/
import (
	"encoding/json"
	"net/http"
)

func RespondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
