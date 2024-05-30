package usecase

import "github.com/victorbrugnolo/golang-temp-zipcode/internal/entity"

type ZipCodeRepositoryInterface interface {
	GetZipcodeData(zipcode string) (*entity.ZipcodeDataResponse, *entity.ErrorResponse)
}

type WeatherApiRepositoryInterface interface {
	GetWeatherData(city string) (*entity.WeatherApiResponse, *entity.ErrorResponse)
}
