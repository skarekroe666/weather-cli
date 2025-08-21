# Weather CLI

A simple command-line tool written in Go to fetch and display the current weather for a specified city using the [WeatherAPI](https://www.weatherapi.com/).

## Features

- Fetches current weather data for any city.
- Displays temperature, weather condition, wind speed, and humidity.
- Colors the output red if the temperature is above 20°C.
- Uses a `.env` file to securely store your WeatherAPI key.

## Prerequisites

- [Go](https://golang.org/dl/) 1.18 or newer installed.
- A free API key from [WeatherAPI](https://www.weatherapi.com/).

## Setup

1. **Clone the repository:**
    ```sh
    git clone https://github.com/skarekroe666/weather-cli.git
    cd weather-cli
    ```

2. **Install dependencies:**
    ```sh
    go mod tidy
    ```

3. **Create a `.env` file in the project root:**
    ```
    URL=your_weatherapi_key_here
    ```
    Replace `your_weatherapi_key_here` with your actual WeatherAPI key.

## Usage

### 1. Run to check if it's working

```sh
go run main.go [city]
```

- Replace `[city]` with the name of the city you want the weather for (e.g., `London`, `Miami`).
- If no city is provided, it defaults to `London`.

**Example:**

```sh
go run main.go Miami
```

### 2. Build and run using the build

```sh
go build -o weather
./weather [city]
```

**Example:**

```sh
./weather Miami
```

## Example Output

```
Miami, United States: 32C, Partly cloudy, Wind: 15kmph, Humidity: 70
```

If the temperature is above 20°C, the output will be colored red.

## Troubleshooting

- If you see an error like `Weather API not available. Status: 400, Body: {"error":{"code":1006,"message":"No matching location found."}}`, check your API key and city name.
- Make sure your `.env` file is present and contains the correct key.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
