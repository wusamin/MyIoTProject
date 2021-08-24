package util

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

var LogBracket = "-----------------------------------------"

func LogSetting(dir, prefix string) {
	sysDate := time.Now()
	logFileSuffix := fmt.Sprintf("%d-%02d-%02d", sysDate.Year(), sysDate.Month(), sysDate.Day())
	logFileDirectory := dir
	logFileName := prefix

	logfile, err := os.OpenFile(logFileDirectory+logFileName+logFileSuffix+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		fmt.Println(err)
		return
	}

	log.SetOutput(io.MultiWriter(logfile, os.Stdout, os.Stderr))
	log.SetFlags(log.Ldate | log.Ltime)
}

func Println(v interface{}) {
	log.Println(v)
}
