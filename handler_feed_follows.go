package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/wasay1567/rssagg/internal/database"
)

func (apiCfg *ApiConfig) handlerFeedFollowsCreate(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		responseWithJson(w, 400, map[string]string{"error": "Invalid request body"})
		return
	}

	feed_follow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})

	if err != nil {
		log.Printf("Error creating feed follow: %v", err)
		responseWithJson(w, 400, map[string]string{"error": "couldn't create feed follow"})
		return
	}

	responseWithJson(w, 201, databaseFeedFollowToFeedFollow(feed_follow))
}

func (apiCfg *ApiConfig) handlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feeds_follows, err := apiCfg.DB.GetFeedFollows(r.Context(), user.ID)
	if err != nil {
		responseWithJson(w, 400, "couldn't fetch the feeds follows")
		return
	}

	result := []FeedsFollow{}
	for _, feeds_follow := range feeds_follows {
		result = append(result, databaseFeedFollowToFeedFollow(feeds_follow))
	}

	responseWithJson(w, 200, result)
}

func (apiCfg *ApiConfig) handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIDStr := chi.URLParam(r, "feedFollowId")
	feedFollowID, err := uuid.Parse(feedFollowIDStr)
	if err != nil {
		responseWithJson(w, 400, map[string]string{"error": "can't parse feed follow id"})
		return
	}
	err = apiCfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID:     feedFollowID,
		UserID: user.ID,
	})
	if err != nil {
		responseWithJson(w, 401, map[string]string{"error": "unable to delete feed follow"})
		return
	}

	responseWithJson(w, 200, struct{}{})
}
