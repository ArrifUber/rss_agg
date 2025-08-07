package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/arrifuber/rss_agregator/internal/auth"
	"github.com/arrifuber/rss_agregator/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handler_create_user(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	deCoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := deCoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("error parsing json: %v", err))
		return
	}

	user, err := apiCfg.DB.CreateUsesr(r.Context(), database.CreateUsesrParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error Cannot create user: %v", err))
		return
	}

	respondWithJSON(w, 201, databaseUserToUser(user))
}
func (apiCfg *apiConfig) handler_get_user_by_apikey(w http.ResponseWriter, r *http.Request) {
	apiKey, err := auth.GetApiKey(r.Header)
	if err != nil {
		respondWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
		return
	}

	user, err := apiCfg.DB.GetUserByApiKey(r.Context(), apiKey)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("cannot get user: %v", err))
		return
	}

	respondWithJSON(w, 200, databaseUserToUser(user))
}
