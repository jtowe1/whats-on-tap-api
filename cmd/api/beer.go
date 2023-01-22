package main

import (
	"encoding/json"
	"errors"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"strconv"
)

func (app *application) showBeerHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)

	beer, err := app.BeerModel.Get(id)
	if err != nil {
		log.Fatal(err)
	}

	body, err := json.Marshal(beer)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (app *application) readIDParam(r *http.Request) (int64, error) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter")
	}

	return id, nil
}
