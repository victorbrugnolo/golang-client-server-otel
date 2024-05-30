package entity

type Current struct {
	TempC float64 `json:"temp_c"`
}

type WeatherApiResponse struct {
	Current `json:"current"`
}
