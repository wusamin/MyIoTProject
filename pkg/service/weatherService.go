package service

import (
	"encoding/json"
	"io/ioutil"
	"log"
	c "maid/pkg/config"
	"maid/pkg/repository"
)

// WeatherReport get weatherreport from livedoor API.
func WeatherReport() map[string]interface{} {

	return repository.GetWeatherForecast()
}

// GetWeatherforecast gets Weatherforecast and return it.
func GetWeatherforecast() map[string]string {
	bytes, err := ioutil.ReadFile(c.Config.WebSetting.JSONFilePath + "/darksky.json")
	if err != nil {
		log.Println(err)
		return map[string]string{}
	}
	var r map[string]string

	if err := json.Unmarshal(bytes, &r); err != nil {
		log.Println(err)
		return map[string]string{}
	}

	return r
}

func FindWeatherForecast(vendor string) interface{} {
	switch vendor {
	case "owm":
		return repository.FindWeatherForecast()

	case "darksky":
		bytes, err := ioutil.ReadFile(c.Config.WebSetting.JSONFilePath + "/darksky.json")
		if err != nil {
			log.Println(err)
			return map[string]string{}
		}
		var r map[string]string

		if err := json.Unmarshal(bytes, &r); err != nil {
			log.Println(err)
			return map[string]string{}
		} else {
			return r
		}

	default:
		return map[string]string{}
	}
}
