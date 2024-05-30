package repository

import (
	"github.com/stretchr/testify/mock"
	"github.com/victorbrugnolo/golang-temp-zipcode/internal/entity"
)

type WeatherApiRepositoryMock struct {
	mock.Mock
}

func (m *WeatherApiRepositoryMock) GetWeatherData(city string) (*entity.WeatherApiResponse, *entity.ErrorResponse) {
	args := m.Called(city)

	if args.Get(1) != nil {
		return nil, args.Get(1).(*entity.ErrorResponse)
	}

	return args.Get(0).(*entity.WeatherApiResponse), nil
}
