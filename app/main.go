package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/word", translateWord).Methods("POST")
	r.HandleFunc("/sentence", translateSentence).Methods("POST")
	r.HandleFunc("/history", history).Methods("GET")
	http.Handle("/", r)
	var port string = os.Args[1]

	fmt.Printf("Starting server on PORT: " + port + "\n")
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
