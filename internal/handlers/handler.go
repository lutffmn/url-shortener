package handlers

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type URLMapping struct {
	URL                 string    `json:"url"`
	Alias               string    `json:"alias"`
	ExpiringTime        time.Time `json:"expiring_time"`
	ExpiringTimeDisplay string    `json:"expiring_time_display"`
}

func shortenUrl(url string) string {
	hash := sha1.Sum([]byte(url))

	return base64.RawURLEncoding.EncodeToString(hash[:])[:8]
}

func ShortenHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var url URLMapping
	if err := json.NewDecoder(r.Body).Decode(&url); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
	}

	defer r.Body.Close()

	alias := shortenUrl(url.URL)

	url.Alias = fmt.Sprintf("%s/%s", r.Host, alias)
	url.ExpiringTime = time.Now().Add(24 * time.Hour)
	url.ExpiringTimeDisplay = url.ExpiringTime.Format("2006-01-02")

	fmt.Fprintf(w, "Result: %v", url)
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(http.StatusOK)
}
