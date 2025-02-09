package handlers

import (
	"database/sql"
	"net/http"
	"github.com/go-chi/chi/v5"
)

type YapHandler struct {
	db *sql.DB
}

func NewYapHandler(db *sql.DB) *YapHandler {
	return &YapHandler{db: db}
}

func (h *YapHandler) IncreaseYapCounter(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "Missing trade ID", http.StatusBadRequest)
		return
	}

	query := `UPDATE trades SET yap_counter = yap_counter + 1 WHERE id = $1`
	result, err := h.db.Exec(query, id)
	if err != nil {
		http.Error(w, "Failed to update yap counter", http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, "Failed to get rows affected", http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, "Trade not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}
