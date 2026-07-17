package main

import (
	"encoding/json"
	"net/http"
	"log"
)

func errResponse(w http.ResponseWriter, code int, msg string){
	if code > 499{
		type errStruct struct{
			Error string `json:"error"`
		}
		jsonResponse(w, code, errStruct{
			Error: msg,
		})
	}
}

func jsonResponse(w http.ResponseWriter, code int, payload interface{}){
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error occurred while marshaling the json response %v", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)
}