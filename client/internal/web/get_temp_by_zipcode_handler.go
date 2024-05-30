package web

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/spf13/viper"
	"github.com/victorbrugnolo/golang-temp-zipcode-client/internal/entity"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

func GetTempByZipcodeHandler(w http.ResponseWriter, r *http.Request) {
	tracer := otel.Tracer("get-temp-by-zipcode-tracer")
	requestNameOTEL := viper.GetString("REQUEST_NAME_OTEL")

	carrier := propagation.HeaderCarrier(r.Header)
	ctx := r.Context()
	ctx = otel.GetTextMapPropagator().Extract(ctx, carrier)

	zipcode := &entity.GetTemperatureByZipcodeRequest{}
	var err error
	var req *http.Request

	err = json.NewDecoder(r.Body).Decode(zipcode)

	if err != nil || !validateZipcode(zipcode.Zipcode) {
		http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
		return
	}

	ctx, span := tracer.Start(ctx, "get_temp_on_server_"+requestNameOTEL)
	defer span.End()

	url := viper.GetString("SERVER_URL")

	if url == "" {
		urlWithPath := "http://localhost:8080/" + zipcode.Zipcode + "/temperature"
		req, err = http.NewRequestWithContext(ctx, "GET", urlWithPath, nil)
	} else {
		urlWithPath := url + zipcode.Zipcode + "/temperature"
		req, err = http.NewRequestWithContext(ctx, "GET", urlWithPath, nil)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(req.Header))
	resp, err := http.DefaultClient.Do(req)

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
