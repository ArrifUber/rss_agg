package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/arrifuber/rss_agregator/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handleCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
	deCoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := deCoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("error parsing json: %v", err))
		return
	}

	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.URL,
		UserID:    user.ID,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error Cannot create feed: %v", err))
		return
	}

	respondWithJSON(w, 201, databaseFeedToFedd(feed))
}

func (apiCfg *apiConfig) handleGetFeed(w http.ResponseWriter, r *http.Request) {

	feeds, err := apiCfg.DB.GetFeed(r.Context())

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error get feeds: %v", err))
		return
	}

	respondWithJSON(w, 201, databaseFeedsToFedds(feeds))
}
