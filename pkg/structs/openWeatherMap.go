package structs

type OneCall struct {
	Lat            float64  `json:"lat"`
	Lon            float64  `json:"lon"`
	Timezone       string   `json:"timezone"`
	TimezoneOffset int64    `json:"timezone_offset"`
	Current        Current  `json:"current"`
	Hourly         []Hourly `json:"hourly"`
	Daily          []Daily  `json:"daily"`
}

type Current struct {
	Dt         int64     `json:"dt"`
	Sunrise    int64     `json:"sunrise"`
	Sunset     int64     `json:"sunset"`
	Temp       float64   `json:"temp"`
	FeelsLike  float64   `json:"feels_like"`
	Pressure   int64     `json:"pressure"`
	Humidity   float64   `json:"humidity"`
	DewPoint   float64   `json:"dew_point"`
	Uvi        float64   `json:"uvi"`
	Clouds     float64   `json:"clouds"`
	Visibility int64     `json:"visibility"`
	WindSpeed  float64   `json:"wind_speed"`
	WindDeg    int64     `json:"wind_deg"`
	WindGust   float64   `json:"wind_gust"`
	Weather    []Weather `json:"weather"`
}

type Weather struct {
	Id          int64  `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Hourly struct {
	Dt         int64     `json:"dt"`
	Temp       float64   `json:"temp"`
	FeelsLike  float64   `json:"feelsLike"`
	Pressure   int64     `json:"pressure"`
	Humidity   float64   `json:"humidity"`
	DewPoint   float64   `json:"dewPoint"`
	Uvi        float64   `json:"uvi"`
	Clouds     float64   `json:"clouds"`
	Visibility int64     `json:"visibility"`
	WindSpeed  float64   `json:"windSpeed"`
	WindDeg    int64     `json:"windDeg"`
	WindGust   float64   `json:"windGust"`
	Weather    []Weather `json:"weather"`
	Pop        float64   `json:"pop"`
}

type Daily struct {
	Dt        int64          `json:"dt"`
	Sunrise   int64          `json:"sunrise"`
	Sunset    int64          `json:"sunset"`
	Temp      DailyTemp      `json:"temp"`
	FeelsLike DailyFeelsLike `json:"feelsLike"`
	Pressure  int64          `json:"pressure"`
	Humidity  float64        `json:"humidity"`
	DewPoint  float64        `json:"dewPoint"`
	WindSpeed float64        `json:"windSpeed"`
	WindDeg   int64          `json:"windDeg"`
	Weather   []Weather      `json:"weather"`
	Clouds    float64        `json:"clouds"`
	Pop       float64        `json:"pop"`
	Rain      float64        `json:"rain"`
	Uvi       float64        `json:"uvi"`
}

type DailyTemp struct {
	Day   float64 `json:"day"`
	Min   float64 `json:"min"`
	Max   float64 `json:"max"`
	Night float64 `json:"night"`
	Eve   float64 `json:"eve"`
	Morn  float64 `json:"morn"`
}

type DailyFeelsLike struct {
	Day   float64 `json:"day"`
	Night float64 `json:"night"`
	Eve   float64 `json:"eve"`
	Morn  float64 `json:"morn"`
}
