package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func getData(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	sb := string(body)
	log.Printf(sb)
}

func loadConfig(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
    log.Fatalf("Error loading .env file")
  }


  return os.Getenv(key)
}

func main() {
	url := loadConfig("URL")

	fmt.Println(url)

	getData(url)
}
