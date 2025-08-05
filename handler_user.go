package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/wasay1567/rssagg/internal/auth"
	"github.com/wasay1567/rssagg/internal/database"
)

func (apiCfg *ApiConfig) handlerUserCreate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		responseWithJson(w, 400, map[string]string{"error": "Invalid request body"})
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		Name:      params.Name,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})

	if err != nil {
		log.Printf("Error creating user: %v", err)
		responseWithJson(w, 400, map[string]string{"error": "couldn't create user"})
		return
	}

	responseWithJson(w, 201, databaseUserToUser(user))
}

func (apiCfg *ApiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request) {
	apikey, err := auth.GetAPIKey(r.Header)
	if err != nil {
		responseWithJson(w, 403, map[string]string{"error": "auth error: " + err.Error()})
		return
	}
	user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apikey)
	if err != nil {
		responseWithJson(w, 404, map[string]string{"error": "user not found"})
		return
	}

	responseWithJson(w, 200, databaseUserToUser(user))
}
