package main

import (
	"encoding/json"
	"net/http"
)

type stripePayload struct {
	Currency string `json:"currency"`
	Amount   string `json:"amount"`
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
	Content string `json:"content"`
	ID      int    `json:"id"`
}

func (app *application) GetPaymentIntent(rw http.ResponseWriter, r *http.Request) {
	j := jsonResponse{
		OK: true,
	}

	out, err := json.MarshalIndent(j, "", "   ")
	if err != nil {
		app.errorLog.Println(err)
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.Write(out)
}
