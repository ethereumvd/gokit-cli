package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	loc := "Bhubaneshwar"

	if len(os.Args) == 2 {
		loc = os.Args[1]
	}

	api_key := os.Getenv("API_KEY")

	BaseURL := "https://api.weatherapi.com/v1/forecast.json?key=" + api_key + "&q=" + loc + "&days=1&aqi=no&alerts=no"

	resp, err := http.Get(BaseURL)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		panic("WeatherAPI Down or invalid city name!")
	}

	respbody, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// fmt.Println(string(respbody))

	weather := Weather{}

	err = json.Unmarshal(respbody, &weather)
	if err != nil {
		panic(err)
	}

	location, current, hrs := weather.Location, weather.Current, weather.ForeCast.ForecastDay[0].Hour

	fmt.Printf(
		"%s, %s: %.0fC, %s\n",
		location.Name,
		location.Country,
		current.TempC,
		current.Condition.Text,
	)

	for _, hr := range hrs {
		date := time.Unix(hr.TimeEpoch, 0)

		if date.Before(time.Now()) { continue }

		mess := fmt.Sprintf(
			"%s - %-.0fC, %.0f, %s\n",
			date.Format("15:04"),
			hr.TempC,
			hr.ChanceofRain,
			hr.Condition.Text,
		)

		if hr.ChanceofRain < 50 {
			fmt.Print(mess)
		} else {
			color.Blue(mess)
		}

	}
}
