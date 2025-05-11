package main

import (
	"encoding/json"
	"net/http"
)

func enableCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

type Response struct {
	Message string `json:"message"`
}

func newPageDataHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)
	w.Header().Set("Content-Type", "application/json")
	response := Response{Message: "Hello New Page!"}
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/newPageDataHandler", newPageDataHandler)

	http.ListenAndServe(":8080", nil)
}
