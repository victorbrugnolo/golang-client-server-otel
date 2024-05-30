package usecase

import (
	"github.com/victorbrugnolo/golang-temp-zipcode/internal/entity"
)

type GetTemperatureByZipcodeUseCase struct {
	zipCodeRepository    ZipCodeRepositoryInterface
	weatherApiRepository WeatherApiRepositoryInterface
}

func NewGetTemperatureByZipcodeUseCase(zipCodeRepository ZipCodeRepositoryInterface, weatherApiRepository WeatherApiRepositoryInterface) *GetTemperatureByZipcodeUseCase {
	return &GetTemperatureByZipcodeUseCase{
		zipCodeRepository:    zipCodeRepository,
		weatherApiRepository: weatherApiRepository,
	}
}

func (g *GetTemperatureByZipcodeUseCase) Execute(zipcode string) (*entity.GetTemperatureByZipcodeResponse, *entity.ErrorResponse) {
	zipcodeData, err := g.zipCodeRepository.GetZipcodeData(zipcode)

	if err != nil {
		return nil, err
	}

	weatherApiResponse, err := g.weatherApiRepository.GetWeatherData(zipcodeData.Localidade)

	if err != nil {
		return nil, err
	}

	celsiusTemp := weatherApiResponse.Current.TempC
	fahrenheitTemp := (celsiusTemp * 1.8) + 32
	kelvinTemp := celsiusTemp + 273

	getTemperatureByZipcodeResponse := entity.GetTemperatureByZipcodeResponse{
		TempC: celsiusTemp,
		TempF: fahrenheitTemp,
		TempK: kelvinTemp,
	}

	return &getTemperatureByZipcodeResponse, nil
}
