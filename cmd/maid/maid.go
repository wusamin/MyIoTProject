package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"maid/pkg/config"
	"maid/pkg/handler"
	"maid/pkg/service"
)

func main() {
	// set log
	logSetting()
	defer recovery()
	log.Println("maid has been started...")
	config.Config = config.LoadConfig()
	handler.Schedule()

	router := gin.Default()
	handler.RestController(router)
	handler.SetRouting(router)
	handler.ViewController(router)

	service.CastBootVoice()

	router.Run(":" + config.Config.WebSetting.Port)
}

func logSetting() {
	sysDate := time.Now()
	logFileSuffix := fmt.Sprintf("%02d-%02d-%02d", sysDate.Year(), sysDate.Month(), sysDate.Day())
	logFileDirectory := "./logs/"
	logFileName := "maid-"

	logfile, err := os.OpenFile(logFileDirectory+logFileName+logFileSuffix+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		fmt.Println(err)
		return
	}

	log.SetOutput(io.MultiWriter(logfile, os.Stdout, os.Stderr))
	log.SetFlags(log.Ldate | log.Ltime)

	exec.Command("ln", "-nfs", "/opt/MAID/logs/"+logFileName+logFileSuffix+".log", logFileDirectory+"maid.log").Run()
}

func recovery() {
	if err := recover(); err != nil {
		log.Printf("Panic has happened: %v", err)
		for depth := 0; ; depth++ {
			_, file, line, ok := runtime.Caller(depth)
			if !ok {
				break
			}
			log.Printf("======> %d: %v:%d", depth, file, line)
		}
	}

	os.Exit(100)
}
