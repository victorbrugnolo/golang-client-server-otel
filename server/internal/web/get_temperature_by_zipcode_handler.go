package web

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
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
	cep := mux.Vars(r)["zipcode"]

	if cep == "" || len(cep) != 8 {
		http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
		return
	}

	resp, err := h.getTemperatureByZipCodeUseCase.Execute(cep)

	if err != nil {
		http.Error(w, err.Message, err.StatusCode)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
