package main

import (
	"log"
	"net/http"
	"encoding/json"
)

func translateWord(w http.ResponseWriter, req *http.Request) {

	var js []byte
	var err error
	var data InputWord
	var result OutputWord
	var customErr CustomError

	if err_ := json.NewDecoder(req.Body).Decode(&data); err_ != nil {
		log.Println(err_)
	}

	if data.EnglishWord == "" {
		customErr.Msg = "Input Data Error"
		js, err = json.Marshal(customErr)
	} else {
		translatedWord := translate(data.EnglishWord)

		if translations == nil {
			translations = make(map[string]string)
		}
		translations[data.EnglishWord] = translatedWord
		
		result.GopherWord = translatedWord

		js, err = json.Marshal(result)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}