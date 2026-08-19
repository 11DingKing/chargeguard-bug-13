package httpapi

import (
	"chargeguard/internal/charging"
	"encoding/json"
	"net/http"
)

func TaskHTTPHandler(w http.ResponseWriter, r *http.Request) {
	history := charging.InspectionHistory()
	if current := r.URL.Query().Get("current"); current != "" {
		view := make([]string, 0, len(history)+1)
		view = append(view, history...)
		view = append(view, current)
		history = view
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(history)
}
