package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/victorbrugnolo/golang-temp-zipcode/internal/entity"
)

type WeatherApiRepository struct{}

func NewWeatherApiRepository() *WeatherApiRepository {
	return &WeatherApiRepository{}
}

func (r *WeatherApiRepository) GetWeatherData(city string) (*entity.WeatherApiResponse, *entity.ErrorResponse) {
	key := os.Getenv("WEATHER_API_KEY")
	url := fmt.Sprintf("https://api.weatherapi.com/v1/current.json?key=%s&q=%s&aqi=no", url.QueryEscape(key), url.QueryEscape(city))
	resp, err := http.Get(url)

	if err != nil {
		return nil, entity.BuildErrorResponse(http.StatusInternalServerError, err)
	}

	if resp.StatusCode == 400 {
		errorMessage := fmt.Sprintf("can not find weather for city %s", city)
		return nil, entity.BuildErrorResponse(http.StatusNotFound, errors.New(errorMessage))
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, entity.BuildErrorResponse(http.StatusInternalServerError, err)
	}

	var weatherApiResponse entity.WeatherApiResponse
	err = json.Unmarshal(body, &weatherApiResponse)

	if err != nil {
		return nil, entity.BuildErrorResponse(http.StatusInternalServerError, err)
	}

	return &weatherApiResponse, nil
}
