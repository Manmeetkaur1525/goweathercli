package main

import (
	// "fmt"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/fatih/color"
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
		Forecastday []struct //array of object , slice of struct
		{
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

func main() {
	 q := "India"
	 if len(os.Args) >= 2{
		q = os.Args[1]
	 }
	res, err := http.Get("http://api.weatherapi.com/v1/forecast.json?key="+os.Getenv("API_KEY")+"&q="+q+"&days=1")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		panic("Weather Api not available")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	// fmt.Println(string(body))

	var weather Weather
	err = json.Unmarshal(body, &weather) //this will convert it into the body to whatever we pass into the unmarchal function
	if err != nil {
		panic(err)

	}
	// fmt.Println(weather)
	location, current, hours := weather.Location, weather.Current, weather.Forecast.Forecastday[0].Hour

	fmt.Printf("%s, %s : %.0fC , %s\n", location.Country, location.Name, current.TempC, current.Condition.Text)

	loc, _ := time.LoadLocation("Asia/Kolkata")
	today := time.Now().In(loc)
	fmt.Println("Today is " , today.Format("02 Jan 2006"))
	for _, hour := range hours {
		date := time.Unix(hour.TimeEpoch, 0).In(loc)
		// fmt.Println(hour.TimeEpoch)
		if date.Before(today) {
			continue
		}

		
//printf allows only 5 format specifiers 
		msg := fmt.Sprintf( "time %s - %.0fC, %.0f%% , %s \n",
			date.Format("15:04"),
			hour.TempC,
			hour.ChanceOfRain,
			hour.Condition.Text,
		)
		if hour.ChanceOfRain < 40 {
			fmt.Println(msg)
		}else{
					color.Red(msg)
		}

	}
}

//variables that are easier to work with
