package main

import "net/http"

func (app *application) VirtualTerminal(rw http.ResponseWriter, r *http.Request) {
	if err := app.renderTemplate(rw, r, "terminal", nil); err != nil {
		app.errorLog.Println(err)
	}
}
