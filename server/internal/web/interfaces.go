package web

import "github.com/victorbrugnolo/golang-temp-zipcode/internal/entity"

type GetTemperatureByZipCodeUseCaseInterface interface {
	Execute(zipcode string) (*entity.GetTemperatureByZipcodeResponse, *entity.ErrorResponse)
}
