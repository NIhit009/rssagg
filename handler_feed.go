package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"github.com/NIhit009/rssagg/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User ){
	type feedParams struct{
		Name string `json:"name"`
		URL string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)
	params := feedParams{}

	err := decoder.Decode(&params)
	if err != nil {
		errResponse(w, 400, fmt.Sprintf("Something went while decoding: %v", err))
		return
	}
	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: params.Name,
		Url: params.URL,
		UserID: user.ID,
	})
	if err != nil{
		errResponse(w, 400, fmt.Sprintf("Could not create the feed: %v", err))
	}
	jsonResponse(w, 201, databaseFeedtoFeed(feed))
}

func (apiCfg *apiConfig) handlerGetAllFeeds(w http.ResponseWriter, r *http.Request){
	feeds, err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil{
		errResponse(w, 400, fmt.Sprintf("Error while retreving the feeds %v", err))
		return
	}
	jsonResponse(w, 200, databaseFeedstoFeeds(feeds))
}