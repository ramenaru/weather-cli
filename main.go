package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
)

type Weather struct {
	Location struct {
		Name    string `json:"name"`
		Country string `json:"country"`
	} `json:"location"`

	Current struct {
		TempC     float64 `json:"temp_c"`
		Condition struct {
			Text string `json:"text"`
		} `json:"condition"`
	} `json:"current"`

	Forecast struct {
		Forecastday []struct {
			Hour []struct {
				TimeEpoch int64   `json:"time_epoch"`
				TempC     float64 `json:"temp_c"`
				Condition struct {
					Text string `json:"text"`
				} `json:"condition"`
				ChanceOfRain float64 `json:"chance_of_rain"`
			} `json:"hour"`
		} `json:"forecastday"`
	} `json:"forecast"`
}

type WeatherLocation struct {
	Name    string `json:"name"`
	Country string `json:"country"`
}

type WeatherCurrent struct {
	TempC     float64 `json:"temp_c"`
	Condition struct {
		Text string `json:"text"`
	} `json:"condition"`
}

type WeatherHour struct {
	TimeEpoch int64   `json:"time_epoch"`
	TempC     float64 `json:"temp_c"`
	Condition struct {
		Text string `json:"text"`
	} `json:"condition"`
	ChanceOfRain float64 `json:"chance_of_rain"`
}

func printHeader(location WeatherLocation, current WeatherCurrent) {
	color.Cyan("\n🌦️ Weather Forecast for %s, %s\n", location.Name, location.Country)
	color.Yellow("🌡️ Current Temperature: %s, Condition: %s\n\n", colorizeTemperature(current.TempC), current.Condition.Text)
}

func printForecast(hours []WeatherHour) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Time", "Temperature (°C)", "Chance of Rain (%)", "Condition"})
	for _, hour := range hours {
		date := time.Unix(hour.TimeEpoch, 0)
		if date.Before(time.Now()) {
			continue
		}
		table.Append([]string{
			date.Format("15:04"),
			colorizeTemperature(hour.TempC),
			fmt.Sprintf("%.0f", hour.ChanceOfRain),
			hour.Condition.Text,
		})
	}
	table.Render()
}

func colorizeTemperature(tempC float64) string {
	if tempC < 10 {
		return color.BlueString("%.0f°C", tempC)
	} else if tempC > 30 {
		return color.RedString("%.0f°C", tempC)
	}
	return color.YellowString("%.0f°C", tempC)
}

func main() {
	q := ""

	locationPrompt := &survey.Input{
		Message: "Enter the location (e.g., city):",
	}
	_ = survey.AskOne(locationPrompt, &q)

	res, err := http.Get("http://api.weatherapi.com/v1/forecast.json?key=9fffb20495d14590a7761327231511&q=" + q + "&days=current&aqi=no&alerts=no")

	if err != nil {
		fmt.Println("Error: Failed to fetch weather data. Please check your internet connection and try again.")
		os.Exit(1)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		fmt.Println("Error: Weather API request failed. Please check your internet connection and try again.")
		os.Exit(1)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error: Failed to read weather data. Please try again.")
		os.Exit(1)
	}

	var weather Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		fmt.Println("Error: Failed to parse weather data. Please try again.")
		os.Exit(1)
	}

	location, current, hours := weather.Location, weather.Current, weather.Forecast.Forecastday[0].Hour

	var weatherHours []WeatherHour
	for _, hour := range hours {
		weatherHours = append(weatherHours, WeatherHour(hour))
	}

	printHeader(location, current)
	printForecast(weatherHours)
}
