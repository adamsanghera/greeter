package greeter

import (
	"encoding/json"
	"net/http"
	"os/exec"
	"strconv"
)

const (
	apikey = "5375b5c08950c1bf50072eebd7fcd54c"
)

// ComplexGreeter looks up the hostname, and the weather at lat and lon from openweathermap.org
type ComplexGreeter struct {
	lat float64
	lon float64
}

type weatherDesc struct {
	ID          int    `json:"ID"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type ret struct {
	Weather []weatherDesc `json:"weather"`
}

// NewComplexGreeter returns a ComplexGreeter with a given lat and lon.
func NewComplexGreeter(lat float64, lon float64) ComplexGreeter {
	return ComplexGreeter{
		lat: lat,
		lon: lon,
	}
}

// SayHello returns a look-up of the hostname
func (cg ComplexGreeter) SayHello() string {
	c := exec.Command("hostname")
	s, _ := (c.Output())
	return "Hi " + string(s)[:len(s)-1] // Sliced off the last character, because it was a newline.
}

// TellWeather hits the openweathermap api
func (cg ComplexGreeter) TellWeather() string {
	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?lat=" + strconv.FormatFloat(cg.lat, 'f', 6, 32) + "&lon=" + strconv.FormatFloat(cg.lon, 'f', 6, 32) + "&APPID=" + apikey)
	if err != nil {
		panic(err)
	}

	r := ret{}
	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		panic(err)
	}
	return r.Weather[0].Description
}
