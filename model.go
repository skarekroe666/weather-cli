package main

type WeatherInfo struct {
	Location struct {
		Name    string `json:"name"`
		Region  string `json:"region"`
		Country string `json:"country"`
	} `json:"location"`
	Current struct {
		TempC     float64 `json:"temp_c"`
		IsDay     int     `json:"is_day"`
		Condition struct {
			Text string `json:"text"`
		} `json:"condition"`
		WindKph  float64 `json:"wind_kph"`
		Humidity int     `json:"humidity"`
		Uv       float64 `json:"uv"`
	} `json:"current"`
}
