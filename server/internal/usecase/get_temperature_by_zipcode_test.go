package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/victorbrugnolo/golang-temp-zipcode/internal/entity"
	"github.com/victorbrugnolo/golang-temp-zipcode/internal/repository"
)

func TestExecuteAndGetZipcodeDataError(t *testing.T) {
	zipCodeRepositoryMock := &repository.ZipCodeRepositoryMock{}
	weatherApiRepositoryMock := &repository.WeatherApiRepositoryMock{}

	zipCodeRepositoryMock.On("GetZipcodeData", "12345678").Return(&entity.ZipcodeDataResponse{}, &entity.ErrorResponse{Message: "Zipcode not found"})

	getTemperatureByZipcodeUseCase := NewGetTemperatureByZipcodeUseCase(zipCodeRepositoryMock, weatherApiRepositoryMock)

	_, err := getTemperatureByZipcodeUseCase.Execute("12345678")

	assert.NotNil(t, err)
	assert.Equal(t, "Zipcode not found", err.Message)

	zipCodeRepositoryMock.AssertExpectations(t)
}

func TestExecuteAndGetWeatherDataError(t *testing.T) {
	zipCodeRepositoryMock := &repository.ZipCodeRepositoryMock{}
	weatherApiRepositoryMock := &repository.WeatherApiRepositoryMock{}

	zipCodeRepositoryMock.On("GetZipcodeData", "12345678").Return(&entity.ZipcodeDataResponse{Localidade: "São Paulo"}, nil)
	weatherApiRepositoryMock.On("GetWeatherData", "São Paulo").Return(&entity.WeatherApiResponse{}, &entity.ErrorResponse{Message: "City not found"})

	getTemperatureByZipcodeUseCase := NewGetTemperatureByZipcodeUseCase(zipCodeRepositoryMock, weatherApiRepositoryMock)

	_, err := getTemperatureByZipcodeUseCase.Execute("12345678")

	assert.NotNil(t, err)
	assert.Equal(t, "City not found", err.Message)

	zipCodeRepositoryMock.AssertExpectations(t)
	weatherApiRepositoryMock.AssertExpectations(t)
}

func TestExecuteSuccess(t *testing.T) {
	zipCodeRepositoryMock := &repository.ZipCodeRepositoryMock{}
	weatherApiRepositoryMock := &repository.WeatherApiRepositoryMock{}

	zipCodeRepositoryMock.On("GetZipcodeData", "12345678").Return(&entity.ZipcodeDataResponse{Localidade: "São Paulo"}, nil)
	weatherApiRepositoryMock.On("GetWeatherData", "São Paulo").Return(&entity.WeatherApiResponse{Current: entity.Current{TempC: 10}}, nil)

	getTemperatureByZipcodeUseCase := NewGetTemperatureByZipcodeUseCase(zipCodeRepositoryMock, weatherApiRepositoryMock)

	response, err := getTemperatureByZipcodeUseCase.Execute("12345678")

	assert.Nil(t, err)
	assert.Equal(t, 10.0, response.TempC)
	assert.Equal(t, 50.0, response.TempF)
	assert.Equal(t, 283.0, response.TempK)

	zipCodeRepositoryMock.AssertExpectations(t)
	weatherApiRepositoryMock.AssertExpectations(t)
}
