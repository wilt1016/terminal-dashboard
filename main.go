package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type WeatherResponse struct {
	Name string `json:"name"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
	} `json:"main"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
}

func clearScreen() {
	fmt.Println("\033[H\033[2J")
}

func getWeather(city string, apiKey string) WeatherResponse {

	//city := "Red+Deer"
	//apiKey := "9b107786106090c329d26b6f1fb4acb0"

	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", city, apiKey)

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("error fetching weather:", err)
		return WeatherResponse{}
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("error reading response", err)
		return WeatherResponse{}
	}

	var weatherData WeatherResponse
	err = json.Unmarshal(body, &weatherData)
	if err != nil {
		fmt.Println("error parsing weather:", err)
		return WeatherResponse{}
	}
	return weatherData
}

func main() {

	weatherData := getWeather("Red+Deer", "9b107786106090c329d26b6f1fb4acb0")
	lastFetch := time.Now()

	for {
		clearScreen()

		time.Since(lastFetch)
		currentTime := time.Now()
		//formats time so it isnt super odd. 15:04 ect. is actually the date and time needed foor it to mean format
		formattedTime := currentTime.Format("15:04 Jan 2")

		fmt.Println(formattedTime)

		if time.Since(lastFetch) > 5*time.Minute {
			weatherData = getWeather("Red+Deer", "9b107786106090c329d26b6f1fb4acb0")
			lastFetch = time.Now()
		}
		fmt.Printf("Temp: %.1f°C (feels like %.1f°C)\n", weatherData.Main.Temp, weatherData.Main.FeelsLike)
		fmt.Printf("Conditions: %s\n", weatherData.Weather[0].Description)

		time.Sleep(time.Second)
	}
}
