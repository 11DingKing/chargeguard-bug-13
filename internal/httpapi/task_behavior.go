package httpapi

import (
	"chargeguard/internal/charging"
	"encoding/json"
	"net/http"
)

func TaskHTTPHandler(w http.ResponseWriter, r *http.Request) {
	history := charging.InspectionHistory()
	if current := r.URL.Query().Get("current"); current != "" {
		history[len(history)-1] = current
		history = append(history, "current")
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(history)
}
