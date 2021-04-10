package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gookit/color"
	"github.com/joho/godotenv"

	models "go_outside/models"
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

	var url string = "https://api.openweathermap.org/data/2.5/weather?units=imperial&lat=" + lat + "&lon=" + lng + "&appid=" + key

	var body = getData(url)

	var weatherResponse models.WeatherResponse
	weatherResponse = models.BuildWeatherResponse(body)

	printWeather(weatherResponse)
}

func printWeather(weatherData models.WeatherResponse) {
	models.GetClouds()
	fmt.Println()
	fmt.Printf("\nCurrent weather in %s: \n", weatherData.Name)
	fmt.Println()

	color.Cyan.Printf("%.1f degrees F,", weatherData.Main.Temp)
	color.FgLightRed.Printf(" feels like %.1f, ", weatherData.Main.Feels_like)
	color.Magenta.Printf("%s \n", weatherData.Weather[0].Main)
	fmt.Println()

}
