package handler

import (
	"log"
	"maid/pkg/service"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"

	conf "maid/pkg/config"
	repo "maid/pkg/repository"
	s "maid/pkg/structs"
)

// RestController is routing REST API.
func RestController(r *gin.Engine) {
	r.Use(cors.Default())

	r.GET("/", getTest)
	r.POST("/", postTest)
	r.GET("/timesignal", timeSignal)
	r.GET("/speech/timesignal/:hour", timeSignalHour)
	// r.GET("/weatherforecast", forecast)
	r.GET("/reloadconfig", reloadConfig)
	r.GET("/exit", exit)

	{
		maid := r.Group("/maid")
		maid.GET("/weather", weatherReport)
		maid.GET("/file/:voice/:fileName", returnFile)
		maid.GET("/device/:device/power/:status", toggleDevicePower)
		maid.GET("/device/:device/setting/:key/:value", updateDeviceSetting)
		maid.GET("/sensorvalue/:sensorType/:selectType", getSensorValueReported)
		maid.GET("/sensor/value/:deviceID/:sensorType", fetchSensorValue)
		maid.POST("/key/:action", operateKey)
		maid.POST("/key", doorGirl)
		maid.POST("/timer", registerTimer)
		// maid.GET("/nowweather", getNowWeather)
		maid.GET("/weatherforecast/:vendor", getWeatherforecast)

	}

	{
		dashboard := r.Group("/maid/dashboard")
		dashboard.GET("/screen/:power", dashboardScreen)
		dashboard.GET("/cast/voice", castDashboard)
	}

	{
		light := r.Group("/maid/light")
		light.GET("/on/:bulbName", changeLighbulb)
		light.GET("/off/:bulbName", turnOffLighbulb)
		light.GET("/ip", getTplinkIP)
		light.GET("/control/:bulbName", controlLightBulb)
	}
}

// reloadConfig is endpoint reloading config.toml.
var reloadConfig = func(c *gin.Context) {
	conf.Config = conf.LoadConfig()
	c.JSON(200, gin.H{
		"status": "Reloading config has successed.",
	})
}

// registerTimer registers record to timer table.
var registerTimer = func(c *gin.Context) {
	var t s.TimerRequest

	err := c.BindJSON(&t)

	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"status": "An error has occured on getting request body.",
		})
		return
	}

	service.RecordTimer(&t)
}

// dashboardScreen controls dashboard screen on/off.
var dashboardScreen = func(c *gin.Context) {
	switch c.Param("power") {
	case "on":
		service.ChangeScreenStatus(service.On)
	case "off":
		service.ChangeScreenStatus(service.Off)
	default:
		c.JSON(400, gin.H{
			"message": "The parameter 'power' permit only 'on' or 'off'",
		})
		return
	}

	c.JSON(204, gin.H{})
}

// castDashboard casts voice to voiceroid on desktop from dashboard.
var castDashboard = func(c *gin.Context) {
	manuscript := c.Query("manuscript")

	service.CastVoiceDashboard(manuscript)
	c.JSON(204, gin.H{})
}

// doorGirl is endpoint toggling Sesame status on/off.
var doorGirl = func(c *gin.Context) {
	if !service.CollateReuqstToken(c.Request.Header["Token"][0]) {
		c.JSON(400, gin.H{
			"message": "Calling API has been failed. Try again later.",
		})
		return
	}

	code, mes := service.Key()
	c.JSON(code, gin.H{"message": mes["message"]})
}

// doorGirl is endpoint controlling Sesame by path parameter.
var operateKey = func(c *gin.Context) {
	if !service.CollateReuqstToken(c.Request.Header["Token"][0]) {
		c.JSON(400, gin.H{
			"message": "Calling API has been failed. Try again later.",
		})
		return
	}

	act := c.Param("action")

	switch act {
	case "lock":
		fallthrough
	case "unlock":
		ret, err := service.SmartLock(act)

		if err != nil {
			log.Println(err)
			c.JSON(400, gin.H{"message": "An error has occured on action \"" + act + "\""})
			return
		}

		c.JSON(200, gin.H{
			"data": ret,
		})
		return

	case "status":
		ret, err := service.GetKeyStatus()

		if err != nil {
			log.Println(err)
			c.JSON(400, gin.H{"message": "An error has occured on action \"" + act + "\""})
			return
		}

		c.JSON(200, gin.H{
			"data": ret,
		})
		return

	default:
		c.JSON(400, gin.H{
			"message": "A path " + act + " is not suitable.",
		})
		return
	}
}

var getTest = func(c *gin.Context) {
}

var postTest = func(c *gin.Context) {
	service.RegisterTimeSignal()
}

var changeLighbulb = func(cxt *gin.Context) {
	brightness := cxt.Query("b")
	if brightness == "" {
		cxt.JSON(400, gin.H{
			"status": "brightnress is required parameter.",
		})
		return
	}

	b, err := strconv.Atoi(brightness)

	if err != nil {
		cxt.JSON(400, gin.H{
			"status": "b is permitted only numeric.",
		})
		return
	}

	transition := cxt.Query("t")

	t := 3000
	if transition != "" {
		_t, err := strconv.Atoi(transition)

		if err != nil {
			cxt.JSON(400, gin.H{
				"status": "transition is permitted only numeric.",
			})
			return
		}

		t = _t
	}

	bulbName := cxt.Param("bulbName")
	log.Println(bulbName)

	if err := service.ChangeLightbulbIlumination(bulbName, b, t); err != nil {
		cxt.JSON(400, gin.H{
			"status": "a error has occurred on changing bulb status.",
		})
		return
	}

	cxt.JSON(200, gin.H{
		"status": "changing lightbulb status has been suucessed.",
	})
	return
}

var turnOffLighbulb = func(cxt *gin.Context) {
	transition := cxt.Query("t")

	t := 3000
	if transition != "" {
		_t, err := strconv.Atoi(transition)

		if err != nil {
			cxt.JSON(400, gin.H{
				"status": "transition is permitted only numeric.",
			})
			return
		}

		t = _t
	}

	bulbName := cxt.Param("bulbName")

	if err := service.TurnOffLightBulb(bulbName, t); err != nil {
		cxt.JSON(400, gin.H{
			"status": "a error has occurred on changing bulb status.",
		})
		return
	}

	cxt.JSON(200, gin.H{
		"status": "changing lightbulb status has been suucessed.",
	})
	return

}

var controlLightBulb = func(cxt *gin.Context) {
	m := map[string]string{}

	if il := cxt.Query("illuminance"); il != "" {
		m["il"] = il
	}

	if temp := cxt.Query("temperature"); temp != "" {
		m["temp"] = temp
	}

	if t := cxt.Query("transition"); t != "" {
		m["transition"] = t
	} else {
		m["transition"] = "2000"
	}

	r, json := service.ChangeLightBulbStatus(cxt.Param("bulbName"), m)

	cxt.JSON(r, json)
}

var fetchSensorValue = func(cxt *gin.Context) {
	from := cxt.Query("from")
	to := cxt.Query("to")

	loc, err := time.LoadLocation("Asia/Tokyo")

	if err != nil {
		log.Println(err)
		cxt.JSON(400, gin.H{
			"status": "An error has occured.",
		})
		return
	}

	fromDate, err := time.ParseInLocation("20060102150405", from, loc)

	if err != nil {
		log.Println(err)
		cxt.JSON(400, gin.H{
			"status": "could not parse to day: " + from,
		})
		return
	}

	toDate, err := time.ParseInLocation("20060102150405", to, loc)

	if err != nil {
		log.Println(err)
		cxt.JSON(400, gin.H{
			"status": "could not parse to day: " + to,
		})
		return
	}

	sensorType := cxt.Param("sensorType")
	deviceID := cxt.Param("deviceID")

	retData := service.GetSensorValueRecorded(sensorType, "", deviceID, fromDate, toDate)

	cxt.JSON(200, gin.H{
		"data": retData,
	})
}

var getSensorValueReported = func(cxt *gin.Context) {

	baseDate := cxt.Query("day")
	from := cxt.Query("from")
	to := cxt.Query("to")

	sensorType := cxt.Param("sensorType")

	switch sensorType {
	case "te":
	case "hu":
	case "il":
	case "mo":
	case "co2":
	case "temperature":
		break
	default:
		cxt.JSON(400, gin.H{
			"status": "sensorType is unsuitable:" + sensorType,
		})
		return
	}

	selectType := cxt.Param("selectType")

	switch selectType {
	case "daily":
	case "weekly":
	case "monthly":
		break
	default:
		cxt.JSON(400, gin.H{
			"status": "selectType is unsuitable:" + selectType,
		})
		return
	}

	deviceID := cxt.Query("deviceId")

	if deviceID == "" {
		deviceID = "natureremo"
	}

	if baseDate != "" {

		_, err := time.Parse("20060102", baseDate)

		if err != nil {
			cxt.JSON(400, gin.H{
				"status": "day couldn't parse to day: " + baseDate,
			})
			return
		}
	}

	loc, err := time.LoadLocation("Asia/Tokyo")

	if err != nil {
		return
	}

	fromDate, err := time.ParseInLocation("20060102150405", from, loc)

	if err != nil {
		log.Println(err)
		cxt.JSON(400, gin.H{
			"status": "day couldn't parse to day: " + from,
		})
		return
	}

	toDate, err := time.ParseInLocation("20060102150405", to, loc)

	if err != nil {
		log.Println(err)
		cxt.JSON(400, gin.H{
			"status": "day couldn't parse to day: " + to,
		})
		return
	}

	retData := service.GetSensorValueRecorded(sensorType, selectType, deviceID, fromDate, toDate)

	cxt.JSON(200, gin.H{
		"data": retData,
	})

}

// control power switch device has.
var toggleDevicePower = func(cxt *gin.Context) {
	service.ControlDevicePower(cxt.Param("device"), cxt.Param("status"))
	cxt.JSON(204, gin.H{})
}

// Get server's file specified by url path.
var returnFile = func(cxt *gin.Context) {
	filePath := "files/" + cxt.Param("voice") + "/" + cxt.Param("fileName")
	cxt.Header("Content-Type", "application/octet-stream")
	cxt.File(filePath)
}

var timeSignal = func(cxt *gin.Context) {
	var text string

	if time.Now().Hour() == 0 {
		text = service.TimeSignalWeather()
	} else {
		text = service.TimeSignal()
	}

	service.BroadcastMessageToDashboard(strings.ReplaceAll(text, "ゼロゼロゼロゼロです。", ""))

	cxt.JSON(200, gin.H{
		"text": text,
	})
}

// Get weatherreport.
var weatherReport = func(cxt *gin.Context) {
	apiResult := service.WeatherReport()
	println(service.TimeSignalWeather())
	cxt.JSON(http.StatusOK, gin.H{
		"status":      "posted",
		"title":       apiResult["title"],
		"description": apiResult["description"].(map[string]interface{})["text"].(string),
		"pinpoint":    apiResult["forecasts"].([]interface{})[0].(map[string]interface{})["telop"], //["telop"].(string) + "です。",
		"text":        service.TimeSignalWeather(),
	})
}

var forecast = func(cxt *gin.Context) {
	ret := service.GetWeatherforecast()
	a := map[string]interface{}{}
	for k, v := range ret {
		a[k] = v
	}
	cxt.JSON(200, a)
}

// Getting timesignal text for speach.
var timeSignalHour = func(c *gin.Context) {
	hour, _ := strconv.Atoi((c.Param("hour")))

	service.BroadcastTimesignal(hour)

	repo.InsertLog("TimeSignal", "OK")
	c.JSON(200, gin.H{})
}

// CORSMiddleware set gin settings for Cross-Origin Resource Sharing.
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

var getTplinkIP = func(c *gin.Context) {
	if e := service.RecordIPTPLink(); e != nil {
		log.Println(e)
		c.JSON(400, gin.H{"message": "An error has occured."})
	}

	c.JSON(204, gin.H{})
}

var exit = func(c *gin.Context) {
	log.Println("has received exit signal.")
	conf.WebSocket.Close()

	os.Exit(0)
}

var getWeatherforecast = func(c *gin.Context) {
	vendor := c.Param("vendor")
	ret := service.FindWeatherForecast(vendor)

	c.JSON(200, ret)
}

var updateDeviceSetting = func(c *gin.Context) {
	key := c.Param("key")
	value := c.Param("value")
	service.UpdateAirconSetting(key, value)

	c.JSON(204, gin.H{})
}
