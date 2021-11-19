package main

import (
	"fmt"
	"net/http"
)

// the application is structure + interfaces for the space
func (app *application) indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to QUOTOS")
}
