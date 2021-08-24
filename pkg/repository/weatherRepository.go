package repository

import (
	"encoding/json"
	"log"

	"github.com/go-resty/resty/v2"
)

func GetWeatherForecast() map[string]interface{} {
	url := "https://api.openweathermap.org/data/2.5/onecall"
	latitude := "35.728471"
	longitude := "139.840651"
	exclude := "minutely"
	apiKey := ""
	lang := "ja"
	units := "metric"

	queryParam := map[string]string{}

	queryParam["lat"] = latitude
	queryParam["lon"] = longitude
	queryParam["exclude"] = exclude
	queryParam["appid"] = apiKey
	queryParam["lang"] = lang
	queryParam["units"] = units

	client := resty.New()

	res, err := client.R().
		SetQueryParams(queryParam).
		SetHeader("Accept", "application/json").
		Get(url)

	if err != nil {
		log.Println(err)
		return map[string]interface{}{}
	}

	var ret map[string]interface{}
	json.Unmarshal(res.Body(), &ret)

	return ret
}
