package handler

import (
	"log"
	"maid/pkg/repository"
	"maid/pkg/service"
	"time"

	"github.com/carlescere/scheduler"
	"github.com/go-resty/resty/v2"
)

// Schedule is scheduling methods.
func Schedule() {
	log.Println("scheduling...")
	scheduler.Every(2).Minutes().Run(reportingSensorValue)
	scheduler.Every(4).Minutes().Run(controlAircon)
	scheduler.Every(1).Minutes().Run(timer)
	scheduler.Every(5).Seconds().Run(standBy)
	scheduler.Every().Day().At("05:00").Run(purge)
	scheduler.Every(30).Minutes().Run(lighting)
	scheduler.Every().Day().At("07:00").Run(sleepingVoice)
}

var sleepingVoice = func() {
	service.RegisterSleepVoiceSchedule()
}

var reportingSensorValue = func() {
	service.SuggestTurnOnLight()
	service.ControlFan()
}

var controlAircon = func() {
	service.Airconditionning()
}

var standBy = func() {
	service.StandByMe()
}

var timer = func() {
	runTimer()
}

var purge = func() {
	service.PurgeSensorVal()
	service.RegisterTimeSignal()
}

var lighting = func() {
	service.AdjustLightBrightness()
}

// runTimer executes GET with URL scheduled.
func runTimer() {
	now := time.Now()
	r := resty.New().SetTimeout(5 * time.Second).R()

	// タイマー用のテーブルを検索する処理
	for _, schedule := range repository.FindScheduledURL() {
		if schedule.StartUpTime.After(now) {
			continue
		}

		_, err := r.Get(schedule.TargetURL)

		log.Println("-----------------------------------------")
		log.Println("Scheduled URL was executed.")
		log.Printf("%-9s : %s", "URL", schedule.TargetURL)
		log.Printf("%-9s : %s", "Start up", schedule.StartUpTime)

		if err != nil {
			log.Println(err)
		}

		log.Println("-----------------------------------------")

		repository.DeleteScheduledURL(schedule)
	}
}
