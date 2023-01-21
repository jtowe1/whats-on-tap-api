package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func (app *application) showBeerHandler(w http.ResponseWriter, r *http.Request) {
	body, err := json.Marshal(struct {
		Name string
	}{
		Name: "Test name",
	})

	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}
