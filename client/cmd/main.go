package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/spf13/viper"
	"github.com/victorbrugnolo/golang-temp-zipcode-client/internal/web"
)

func init() {
	viper.AutomaticEnv()
}

func main() {
	r := chi.NewRouter()
	r.Post("/temperature", web.GetTempByZipcodeHandler)

	http.ListenAndServe(":8081", r)
}
