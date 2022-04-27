package goChatUtil

import (
	"gochat/errorHandling"
	"net/http"
)

func WriteErrorResponse(w http.ResponseWriter, err *errorHandling.BaseError) {
	response := err.GenerateResponse()
	w.WriteHeader(err.Status)
	w.Write(response)
}
