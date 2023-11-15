package main

import (
	"fmt"
	"io"
	"net/http"
)

type Weather struct {
	Location struct {
		Name    string `json:"name"`
		Country string `json:"country"`
		Region  string `json:"region"`
	} `json:"location"`

	Current struct {
		Name    string `json:"name"`
		Country string `json:"country"`
		Region  string `json:"region"`
	} `json:"current"`
}

func main() {
	res, err := http.Get("http://api.weatherapi.com/v1/forecast.json?key=9fffb20495d14590a7761327231511&q=bandung")

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic("Weather API not Available !")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
