package main

import (
	"encoding/json"
	"net/http"
)

func history(w http.ResponseWriter, req *http.Request) {

	var history History
	history.Data = translations
	js, err := json.Marshal(history)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}