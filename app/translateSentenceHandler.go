package main

import (
	"log"
	"net/http"
	"encoding/json"
	"strings"
)

func translateSentence(w http.ResponseWriter, req *http.Request) {
	var js []byte
	var err error
	var data InputSentence
	var translatedSentence string
	var customErr CustomError
	if err_ := json.NewDecoder(req.Body).Decode(&data); err_ != nil {
		log.Println(err_)
	}
	if data.EnglishSentence == "" {
		customErr.Msg = "Input Data Error"

		js, err = json.Marshal(customErr)

	} else {
		words := strings.Split(data.EnglishSentence, " ")

		for _, word := range words {
			translatedSentence += translate(word) + " "
		}

		if translations == nil {
			translations = make(map[string]string)
		}

		translations[data.EnglishSentence] = strings.TrimSpace(translatedSentence)

		var result OutputSentence
		result.GopherSentence = strings.TrimSpace(translatedSentence)

		js, err = json.Marshal(result)

	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}