package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/gookit/color"
	"github.com/joho/godotenv"

	models "go_outside/models"
)

func getData(url string) string {
	resp, err := http.Get(url)
	printErr(err)

	body, err := ioutil.ReadAll(resp.Body)
	printErr(err)


	sb := string(body)

	return sb
}

func loadConfig(key string) string {
	err := godotenv.Load(".env")

	printErr(err)

	return os.Getenv(key)
}

func main() {
	location := getLocation("http://ipinfo.io")
	coords := strings.Split(location.Loc, ",")

	key := loadConfig("API_KEY")
	lat := coords[0]
	lng := coords[1]

	var url string = "https://api.openweathermap.org/data/2.5/weather?units=imperial&lat=" + lat + "&lon=" + lng + "&appid=" + key

	var body = getData(url)

	weatherResponse := models.BuildWeatherResponse(body)

	printWeather(weatherResponse)
}

func printWeather(weatherData models.WeatherResponse) {
	processCondition(weatherData.Weather[0].Id)
	fmt.Println()
	fmt.Printf("\nCurrent weather in %s: \n", weatherData.Name)
	fmt.Println()

	color.Cyan.Printf("%.1f degrees F,", weatherData.Main.Temp)
	color.FgLightRed.Printf(" feels like %.1f, ", weatherData.Main.Feels_like)
	color.Magenta.Printf("%s: %s \n", weatherData.Weather[0].Main, weatherData.Weather[0].Description)
	fmt.Println()
}

func getLocation(url string) models.Location {
	var location models.Location
	err := json.Unmarshal([]byte(getData(url)), &location)
	printErr(err)

	return location
}

// this is so gross
func processCondition(id int) {
	if id > 800 {
		models.GetClouds()
	} else if id == 800 {
		models.GetClear()
	} else if id >= 700 && id <= 741 {
		models.GetFog()
	} else if id > 741 && id < 781 {
		models.GetAsh()
	} else if id == 781 {
		models.GetTornado()
	} else if id >= 600 && id <= 700 {
		models.GetSnow()
	} else if id >= 500 && id < 600 {
		models.GetRain()
	} else if id >= 300 && id < 400 {
		models.GetDrizzle()
	} else {
		models.GetThunderstorm()
	}
}

func printErr(err error) {
	if err != nil {
		fmt.Errorf(err.Error())
	}
}