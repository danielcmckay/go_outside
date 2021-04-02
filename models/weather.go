package models

import (
	"encoding/json"
	"fmt"
)

type WeatherResponse struct {
	weather []string `json:"weather"`
	Main struct {
		temp       float64 `json:"temp"`
		feels_like float64 `json:"feels_like"`
		pressure   int     `json:"pressure"`
		humidity   int     `json:"humidity"`
	} `json:"main"`
	Sys struct {
		id      int `json:"id"`
		sunrise int `json:"sunrise"`
		sunset  int `json:"sunset"`
	} `json:"sys"`
}

func BuildWeatherResponse(body string) {
	var weatherResponse WeatherResponse
	err := json.Unmarshal([]byte(body), &weatherResponse)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(weatherResponse)
}
