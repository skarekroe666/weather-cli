package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/fatih/color"
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

func getResp() string {
	key, err := loadEnv()
	if err != nil {
		// panic(err)
		log.Fatal("could not load .env file:", err)
	}

	city := "London"
	if len(os.Args) >= 2 {
		city = os.Args[1]
	}
	return fmt.Sprintf("https://api.weatherapi.com/v1/current.json?key=%s&q=%s&days=1&aqi=yes&alerts=no", key, city)
}

func main() {
	fmt.Println("lets start the cli")

	url := getResp()

	res, err := http.Get(url)
	if err != nil {
		log.Fatal("could not get the url:", err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		body, _ := io.ReadAll(res.Body)
		log.Fatalf("Weather API not available. Status: %d, Body: %s", res.StatusCode, string(body))
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Empty body, check the link:", err)
	}

	var weather WeatherInfo
	err = json.Unmarshal(body, &weather)
	if err != nil {
		log.Fatal("failed to unmarshal data:", err)
	}
	// fmt.Println(weather)

	msg := fmt.Sprintf("%s, %s: %.0fC, %s, Wind: %.0fkmph, Humidity: %d",
		weather.Location.Name, weather.Location.Country,
		weather.Current.TempC, weather.Current.Condition.Text,
		weather.Current.WindKph, weather.Current.Humidity)

	if weather.Current.TempC <= 20 {
		fmt.Print(msg)
	} else {
		color.Red(msg)
	}
}
