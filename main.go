package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type CityWeather struct {
	Location struct {
		City    string `json:"city"`
		Country string `json:"country"`
	} `json:"location"`
	Current struct {
		Temperature float64 `json:"temp_c"`
		Condition   struct {
			Description string `json:"description"`
		} `json:"condition"`
	} `json:"current"`
	Forecast struct {
		Daily []struct {
			Hours []struct {
				TimeEpoch    int64   `json:"time_epoch"`
				TemperatureC float64 `json:"temp_c"`
				Condition    struct {
					Description string `json:"description"`
				} `json:"condition"`
				ChanceOfRain float64 `json:"chance_of_rain"`
			} `json:"hours"`
		} `json:"daily"`
	} `json:"forecast"`
}

func main() {
	query := "Kolkata"

	if len(os.Args) >= 2 {
		query = os.Args[1]
	}
	response, err := http.Get("http://api.weatherapi.com/v1/forecast.json?key=36ab413e8fa94626894182233241005&q=" + query + "&days=1&aqi=no&alerts=no")
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		panic("weather data not available")
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	var cityWeather CityWeather
	err = json.Unmarshal(body, &cityWeather)
	if err != nil {
		panic(err)
	}

	location, current, dailyForecast := cityWeather.Location, cityWeather.Current, cityWeather.Forecast.Daily[0].Hours

	fmt.Printf("%s, %s: %.0f°C. %s\n",
		location.City,
		location.Country,
		current.Temperature,
		current.Condition.Description)

	for _, hour := range dailyForecast {
		date := time.Unix(hour.TimeEpoch, 0)

		if date.Before(time.Now()) {
			continue
		}

		if hour.ChanceOfRain < 40 {
			fmt.Printf("%s - %.0f°C, %.0f%%, %s\n",
				date.Format("15:04"),
				hour.TemperatureC,
				hour.ChanceOfRain,
				hour.Condition.Description)
		} else {
			fmt.Printf("\033[31m%s - %.0f°C, %.0f%%, %s\033[0m\n",
				date.Format("15:04"),
				hour.TemperatureC,
				hour.ChanceOfRain,
				hour.Condition.Description)
		}
	}
}
