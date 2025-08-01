package auth

import (
	"Accounting/internal/database"
	"Accounting/internal/entities"
	"encoding/json"
	"net/http"
)

type AuthRequest struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	var req AuthRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	var terminal entities.Terminal
	if err := database.DB.Where("client_id = ? AND client_secret = ?", req.ClientID, req.ClientSecret).First(&terminal).Error; err != nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	token, err := GenerateJWT(terminal.ID, terminal.ClientID)
	if err != nil {
		http.Error(w, "could not generate token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}
