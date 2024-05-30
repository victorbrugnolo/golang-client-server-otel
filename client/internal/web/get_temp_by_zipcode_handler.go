package web

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/spf13/viper"
	"github.com/victorbrugnolo/golang-temp-zipcode-client/internal/entity"
)

func GetTempByZipcodeHandler(w http.ResponseWriter, r *http.Request) {
	zipcode := &entity.GetTemperatureByZipcodeRequest{}
	var err error
	var resp *http.Response

	err = json.NewDecoder(r.Body).Decode(zipcode)

	if err != nil || !validateZipcode(zipcode.Zipcode) {
		http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
		return
	}

	url := viper.GetString("SERVER_URL")

	if url == "" {
		resp, err = http.Get("http://localhost:8080/" + zipcode.Zipcode + "/temperature")

	} else {
		resp, err = http.Get(url + zipcode.Zipcode + "/temperature")

	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if resp.StatusCode != 200 {
		http.Error(w, string(body), resp.StatusCode)
		return
	}

	var getTemperatureByZipcodeResponse entity.GetTemperatureByZipcodeResponse
	err = json.Unmarshal(body, &getTemperatureByZipcodeResponse)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(getTemperatureByZipcodeResponse)

}

func validateZipcode(zipcode string) bool {
	return len(zipcode) == 8
}
