package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/wasay1567/rssagg/internal/database"
)

func (apiCfg *ApiConfig) handlerFeedCreate(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		responseWithJson(w, 400, map[string]string{"error": "Invalid request body"})
		return
	}

	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		Name:      params.Name,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Url:       params.Url,
		UserID:    user.ID,
	})

	if err != nil {
		log.Printf("Error creating feed: %v", err)
		responseWithJson(w, 400, map[string]string{"error": "couldn't create feed"})
		return
	}

	responseWithJson(w, 201, databaseFeedToFeed(feed))
}

func (apiCfg *ApiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := apiCfg.DB.GetFeed(r.Context())
	if err != nil {
		responseWithJson(w, 400, "couldn't fetch the feeds")
		return
	}

	result := []Feed{}
	for _, feed := range feeds {
		result = append(result, databaseFeedToFeed(feed))
	}

	responseWithJson(w, 200, result)
}
