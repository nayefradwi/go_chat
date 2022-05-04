package goChatUtil

import (
	"encoding/json"
	"gochat/errorHandling"
	"net/http"
)

func WriteErrorResponse(w http.ResponseWriter, err *errorHandling.BaseError) {
	response := err.GenerateResponse()
	w.WriteHeader(err.Status)
	w.Write(response)
}

func WriteEmptyCreatedResponse(w http.ResponseWriter) {
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(make(map[string]string))
}
