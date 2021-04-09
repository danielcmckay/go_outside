package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	weather "go_outside/models"
)

func getData(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	sb := string(body)

	return sb
}

func loadConfig(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func main() {
	key := loadConfig("API_KEY")
	lat := "48.209785"
	lng := "-114.308106"

	fmt.Println(key)

	var url string = "https://api.openweathermap.org/data/2.5/weather?lat=" + lat + "&lon=" + lng + "&appid=" + key
	// var url string = "https://api.tomorrow.io/v4/timelines?location=" + lat + "," + lng + "&fields=temperature&timesteps=1h&units=imperial&apikey=" + key;

	var body = getData(url)

	// fmt.Println(body)

	var weatherResponse weather.WeatherResponse
	weatherResponse = weather.BuildWeatherResponse(body)

	fmt.Println(weatherResponse)

}
