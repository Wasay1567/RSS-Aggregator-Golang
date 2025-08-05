package main

import (
	"net/http"

	"github.com/wasay1567/rssagg/internal/auth"
	"github.com/wasay1567/rssagg/internal/database"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *ApiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

		handler(w, r, user)
	}
}
