package models

import (
	"encoding/json"
	"fmt"
)

type WeatherResponse struct {
	Weather []struct {
		Id          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`

	Main struct {
		Temp       float64 `json:"temp"`
		Feels_like float64 `json:"feels_like"`
		Temp_min   float64 `json:"temp_min"`
		Temp_max   float64 `json:"temp_max"`
		Pressure   int     `json:"pressure"`
		Humidity   int     `json:"humidity"`
	}

	Sys struct {
		Id      int `json:"id"`
		Sunrise int `json:"sunrise"`
		Sunset  int `json:"sunset"`
	} `json:"sys"`

	Name string `json:"name"`
}

func BuildWeatherResponse(body string) WeatherResponse {
	var weatherResponse WeatherResponse
	err := json.Unmarshal([]byte(body), &weatherResponse)

	if err != nil {
		fmt.Println(err)
	}

	return weatherResponse
}
