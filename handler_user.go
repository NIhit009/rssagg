package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"github.com/NIhit009/rssagg/internal/auth"
	"github.com/NIhit009/rssagg/internal/database"
	"github.com/google/uuid"
)

func (apiCfg apiConfig)handlerUser(w http.ResponseWriter, r *http.Request){
	type parameters struct{
		Name string `name`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	err := decoder.Decode(&params)
	if err != nil{
		errResponse(w, 400, fmt.Sprintf("Could not decode the body into the correct parameters: %v", err))
		return
	}	
	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: params.Name,
	})
	if err != nil{
		errResponse(w, 400, fmt.Sprintf("Error while creating a user: %v", err))
		return 
	}
	jsonResponse(w, 201, databaseUserToUser(user))
}

func (apiCfg apiConfig) HandlerGetUserByApiKey(w http.ResponseWriter, r  *http.Request){
	apiKey, err := auth.GetApiKey(r.Header)
	if err != nil{
		errResponse(w, 403, fmt.Sprintf("some error was found while logging: %v", err))
		return
	}
	user, err := apiCfg.DB.GetUserByAPIKEY(r.Context(), apiKey)
	if err != nil{
		errResponse(w, 400, fmt.Sprintf("Error while retrieving the user: %v",err))
		return
	}
	jsonResponse(w, 200, databaseUserToUser(user))
}