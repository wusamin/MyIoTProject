package handler

import (
	"log"
	"maid/pkg/service"

	"github.com/gin-gonic/gin"
)

func SetRouting(r *gin.Engine) {
	scenario := r.Group("/maid/scenario")
	scenario.GET("/standbymyhome", standByMyHome)
	scenario.GET("/goodnight", goodNight)
	scenario.GET("/ceiling/:status", turnOnCeilingLight)
	scenario.GET("/turnontv", turnOnTVOnMorinig)
	scenario.GET("/bedlight/:status", toggleBedlight)
	scenario.GET("/adjust-brightness", adjustLightBrightness)
}

var standByMyHome = func(cxt *gin.Context) {
	code, body := service.TurnOnLightOnStandByMyHome()

	cxt.JSON(code, gin.H{"message": body["message"]})
}

var goodNight = func(cxt *gin.Context) {
	service.AdjustLightBeforeSleeping()
	cxt.JSON(204, gin.H{})
}

var turnOnCeilingLight = func(c *gin.Context) {
	switch c.Param("status") {
	case "on":
		service.TurnOnCeilingLight()
	case "off":
		service.TurnOffCeilingLight()
	}

	c.JSON(204, gin.H{})
}

var turnOnTVOnMorinig = func(cxt *gin.Context) {
	service.TurnOnTVonMorning()
	cxt.JSON(204, gin.H{})
}

var toggleBedlight = func(c *gin.Context) {
	service.ToggleBedlight()

	c.JSON(204, gin.H{})
}
var adjustLightBrightness = func(c *gin.Context) {
	log.Println("started2")
	service.AdjustLightBrightness()

	c.JSON(204, gin.H{})
}
