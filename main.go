package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func loadEnv() (string, error) {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	key := os.Getenv("URL")

	return key, nil
}

func main() {
	fmt.Println("lets start the cli")

	key, err := loadEnv()
	if err != nil {
		panic(err)
	}

	city := "london"
	url := fmt.Sprintf("https://api.weatherapi.com/v1/current.json?key=%s&q=%s&day=1&aqi=no&alerts=no", key, city)

	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	res.Body.Close()

	if res.StatusCode != 200 {
		panic("Weathe api not avalaible")
	}

}
