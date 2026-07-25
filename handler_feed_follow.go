package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/NIhit009/rssagg/internal/database"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type feedFollow struct{
		FeedId uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)
	params := feedFollow{}
	err := decoder.Decode(params)
	if err != nil{
		errResponse(w, 400, fmt.Sprintf("Error while creating a feed follow: %v", err))
		return
	}
	feedFollows, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID: user.ID,
		FeedID: params.FeedId,
	})
	if err != nil{
		errResponse(w, 400, fmt.Sprintf("Error while creating feed follow: %v", err))
		return
	}
	jsonResponse(w, 201, databasefeedFollowtoFeedFollow(feedFollows))
}

func (apiCfg *apiConfig) handlerGetFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feed_follows, err := apiCfg.DB.GetFeedFollow(r.Context(), user.ID)
	if err != nil {
		errResponse(w, 400, fmt.Sprintf("Error while retreiving the feed follow: %v", err))
		return
	}
	jsonResponse(w, 200, databasefeedFollowstoFeedFollows(feed_follows))
}

func (apiCfg *apiConfig) handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type result struct{
		sucess bool
		message string
	}
	feedFollowId := chi.URLParam(r, "feedFollowId")
		if feedFollowId == ""{
		errResponse(w, 404, "Error could not get feed follow id")
		return
	}
	NewfeedFollowId, err := uuid.Parse(feedFollowId)
	if err != nil{
		errResponse(w, 400, fmt.Sprintf("Error while parsing the feed follow id: %v", NewfeedFollowId))
		return
	}

	erro := apiCfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID: NewfeedFollowId,
		UserID: user.ID,
	})
	if erro != nil{
		errResponse(w, 400, fmt.Sprintf("Error while deleting the feed follow: %v", erro))
		return
	}
	msg := result{
		sucess: true,
		message: "Deletion completed successfully",
	}
	jsonResponse(w, 201, msg)
}