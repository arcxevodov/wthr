package wthr

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Weather struct {
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
	} `json:"main"`
	Name string `json:"name"`
}

func GetWeather(city, token, units string) (Weather, error) {
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&units=%s&appid=%s", city, units, token)
	resp, err := http.Get(url)
	if err != nil {
		return Weather{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Weather{}, err
	}
	var weather Weather
	err = json.Unmarshal(body[:], &weather)
	if err != nil {
		return Weather{}, err
	}
	return weather, nil
}
