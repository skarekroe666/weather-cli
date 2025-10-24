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

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("could not get the url:", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		log.Fatalf("Weather API not available. Status: %d, Body: %s", resp.StatusCode, string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Empty body, check the link:", err)
	}

	var weather WeatherInfo
	err = json.Unmarshal(body, &weather)
	if err != nil {
		log.Fatal("failed to unmarshal data:", err)
	}
	// fmt.Println(weather)

	msg := fmt.Sprintf("%s, %s: %.0fC, %s, Wind: %.0fkmph, Humidity: %d\n",
		weather.Location.Name, weather.Location.Country,
		weather.Current.TempC, weather.Current.Condition.Text,
		weather.Current.WindKph, weather.Current.Humidity)

	if weather.Current.TempC <= 20 {
		fmt.Print(msg)
	} else {
		color.Red(msg)
	}
}
