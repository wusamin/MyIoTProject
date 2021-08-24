package service

import (
	"fmt"
	"log"
	"maid/pkg/repository"
	"maid/pkg/structs"
	"time"
)

// RecordTimer inserts a record timer request to table for timer.
func RecordTimer(req *structs.TimerRequest) {
	dto := structs.TMaidTimerScheduled{}
	dto.TargetURL = req.URL
	dto.TimerVal = req.Time
	dto.Deleted = "0"

	now := time.Now()
	dto.StartUpTime = now.Add(time.Duration(req.Time) * time.Millisecond)
	dto.InsDate = now
	dto.UpdDate = now

	repository.InsertScheduledURL(&dto)

	log.Println("-----------------------------------------")
	log.Println("Scheduled URL was reporded.")
	log.Printf("%-9s : %s", "URL", dto.TargetURL)
	log.Printf("%-9s : %d", "Time", dto.TimerVal)
	log.Printf("%-9s : %s", "Start up", dto.StartUpTime)
	log.Printf("%-9s : %s", "deleted", dto.Deleted)
	log.Println("-----------------------------------------")
}

func RegisterSchedule(req *structs.TimerRequest) {
	dto := structs.TMaidTimerScheduled{}
	dto.TargetURL = req.URL
	dto.TimerVal = req.Time
	dto.Deleted = "0"

	now := time.Now()
	start, err := time.Parse(YYYYMMDDHHmmSS+" MST", fmt.Sprint(req.Time)+" JST")
	if err != nil {
		log.Println(err)
		return
	}
	dto.StartUpTime = start
	dto.InsDate = now
	dto.UpdDate = now

	repository.InsertScheduledURL(&dto)

	log.Println("-----------------------------------------")
	log.Println("Scheduled URL was reporded.")
	log.Printf("%-9s : %s", "URL", dto.TargetURL)
	log.Printf("%-9s : %d", "Datetime", dto.TimerVal)
	log.Printf("%-9s : %s", "Start up", dto.StartUpTime)
	log.Printf("%-9s : %s", "deleted", dto.Deleted)
	log.Println("-----------------------------------------")
}

const (
	YYYYMMDD             = "20060102"
	YYYYMMDDHHmmSS       = "20060102150405"
	YYYYMMDDHHmmSS_Slash = "2006/01/02 15:04:05"
)
