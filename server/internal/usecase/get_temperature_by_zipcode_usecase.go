package usecase

import (
	"context"

	"github.com/spf13/viper"
	"github.com/victorbrugnolo/golang-temp-zipcode/internal/entity"
	"go.opentelemetry.io/otel/trace"
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

func (g *GetTemperatureByZipcodeUseCase) Execute(ctx context.Context, zipcode string, tracer trace.Tracer) (*entity.GetTemperatureByZipcodeResponse, *entity.ErrorResponse) {
	requestNameOTEL := viper.GetString("REQUEST_NAME_OTEL")

	_, span := tracer.Start(ctx, "get_zipcode_data_on_external_api_"+requestNameOTEL)
	zipcodeData, err := g.zipCodeRepository.GetZipcodeData(zipcode)
	span.End()

	if err != nil {
		return nil, err
	}

	_, span = tracer.Start(ctx, "get_weather_data_on_external_api_"+requestNameOTEL)
	weatherApiResponse, err := g.weatherApiRepository.GetWeatherData(zipcodeData.Localidade)
	span.End()

	if err != nil {
		return nil, err
	}

	celsiusTemp := weatherApiResponse.Current.TempC
	fahrenheitTemp := (celsiusTemp * 1.8) + 32
	kelvinTemp := celsiusTemp + 273

	getTemperatureByZipcodeResponse := entity.GetTemperatureByZipcodeResponse{
		City:  zipcodeData.Localidade,
		TempC: celsiusTemp,
		TempF: fahrenheitTemp,
		TempK: kelvinTemp,
	}

	return &getTemperatureByZipcodeResponse, nil
}
