package main

import "net/http"

func handlerError(w http.ResponseWriter, r *http.Request){
	errResponse(w, 400, "Something went wrong on the server side..")
}