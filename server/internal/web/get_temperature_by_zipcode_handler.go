package web

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/spf13/viper"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

type GetTemperatureByZipCodeHandler struct {
	getTemperatureByZipCodeUseCase GetTemperatureByZipCodeUseCaseInterface
}

func NewGetTemperatureByZipCodeHandler(getTemperatureByZipCodeUseCase GetTemperatureByZipCodeUseCaseInterface) *GetTemperatureByZipCodeHandler {
	return &GetTemperatureByZipCodeHandler{
		getTemperatureByZipCodeUseCase: getTemperatureByZipCodeUseCase,
	}
}

func (h *GetTemperatureByZipCodeHandler) GetTemperatureByZipcodeHandler(w http.ResponseWriter, r *http.Request) {
	tracer := otel.Tracer("get-temp-by-zipcode-tracer")
	requestNameOTEL := viper.GetString("REQUEST_NAME_OTEL")

	carrier := propagation.HeaderCarrier(r.Header)
	ctx := r.Context()
	ctx = otel.GetTextMapPropagator().Extract(ctx, carrier)

	cep := chi.URLParam(r, "zipcode")

	if cep == "" || len(cep) != 8 {
		http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
		return
	}

	_, span := tracer.Start(ctx, "get_temp_on_server_"+requestNameOTEL)
	defer span.End()

	resp, err := h.getTemperatureByZipCodeUseCase.Execute(cep)

	if err != nil {
		http.Error(w, err.Message, err.StatusCode)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
