package main

import (
	"errors"
	"github.com/jtowe1/whats-on-tap-api/internal/data"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func (app *application) showBeerHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)

	beer, err := app.BeerModel.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, beer, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) readIDParam(r *http.Request) (int64, error) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter")
	}

	return id, nil
}
