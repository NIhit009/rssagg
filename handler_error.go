package main

import "net/http"

func handlerError(w http.ResponseWriter, r *http.Request){
	errResponse(w, 500, "Something went wrong on the server side..")
}