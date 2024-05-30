package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/victorbrugnolo/golang-temp-zipcode-client/internal/web"
)

func main() {
	r := chi.NewRouter()
	r.Post("/temperature", web.GetTempByZipcodeHandler)

	http.ListenAndServe(":8081", r)
}
