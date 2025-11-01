package api

import (
	"encoding/json"
	"game-assistant-server/internal/db"
	"game-assistant-server/internal/model"
	"net/http"
	"strconv"

	"database/sql"

	"github.com/gorilla/mux"
)

type CardHandler struct {
	DB *sql.DB
}

// GET /api/cards
func (h *CardHandler) GetCards(w http.ResponseWriter, r *http.Request) {
	card, err := db.GetAllCards(h.DB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(card)
}

// GET /api/cards/{id}
func (h *CardHandler) GetCard(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	card, err := db.GetCard(h.DB, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(card)
}

// POST /api/cards
func (h *CardHandler) InsertCard(w http.ResponseWriter, r *http.Request) {
	var card model.Card
	if err := json.NewDecoder(r.Body).Decode(&card); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id, err := db.InsertCard(h.DB, &card)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	card.ID = id
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(card)
}
