package service

import (
	"fmt"
	"log"
	c "maid/pkg/config"
	repo "maid/pkg/repository"
	"maid/pkg/structs"
	"maid/pkg/util"
	u "maid/pkg/util"
	"strconv"
	"time"
)

// TurnOnCeilingLight turns on lightbulbs on ceil.
func TurnOnCeilingLight() {
	m, e := u.ReadJSONFile(c.Config.WebSetting.JSONFilePath + "/tplink.json")

	if e != nil {
		log.Println(e)
		return
	}

	go ChangeLightbulbIlumination(m["bulb-ceiling-1"], 100, 2000)
	go ChangeLightbulbIlumination(m["bulb-ceiling-2"], 100, 2000)
	go ChangeLightbulbIlumination(m["bulb-ceiling-3"], 100, 2000)
	ChangeLightbulbIlumination(m["bulb-ceiling-4"], 100, 2000)
}

// TurnOffCeilingLight turns off lightbulbs on ceil.
func TurnOffCeilingLight() {
	m, e := u.ReadJSONFile(c.Config.WebSetting.JSONFilePath + "/tplink.json")

	if e != nil {
		log.Println(e)
		return
	}

	go TurnOffLightBulb(m["bulb-ceiling-1"], 2000)
	go TurnOffLightBulb(m["bulb-ceiling-2"], 2000)
	go TurnOffLightBulb(m["bulb-ceiling-3"], 2000)
	TurnOffLightBulb(m["bulb-ceiling-4"], 2000)
}

// TurnOnLightOnStandByMyHome turns on ceilnglight returning home.
func TurnOnLightOnStandByMyHome() (int, map[string]string) {
	// 電気を点けるシナリオ
	now := time.Now()

	// 日中帯に電気の制御を行わない
	if hour := now.Hour(); 8 <= hour && hour <= 15 {
		return 200, map[string]string{"message": "Now hour is within range 8 to 15."}
	}

	// 深夜帯は別の処理で制御したい
	if hour := now.Hour(); 0 <= hour && hour <= 7 {
		return 200, map[string]string{"message": "Now hour is within range 0 to 7."}
	}

	ilSensorVal := repo.FindSensorVal("il", 1)

	sensorVal, err := strconv.ParseFloat(ilSensorVal[0].Val, 64)

	borderLightOn, err := strconv.ParseFloat(c.Config.MaidSetting.BorderSensorValTurnOnLight, 64)
	if err != nil {
		log.Println(err)
		return 400, map[string]string{"message": "An error has occured."}
	}

	// ボーダー以上の明るさなら終了
	if borderLightOn < sensorVal {
		return 200, map[string]string{"message": "Now ilmination is over 'borderLightOn'."}
	}

	moVals := repo.FindSensorVals("natureremo", "mo", 2)

	dur := moVals[0].CreatedAt.Sub(moVals[1].CreatedAt)

	lastDetected := moVals[0].CreatedAt
	lastDetectedDur := now.Sub(lastDetected)

	t := 10 * time.Minute

	// 最後に検知した時刻
	if lastDetectedDur.Nanoseconds() < t.Nanoseconds() {
		// 最後に検知した時刻が現在時間よりも
		// 最新の人感センサの値で、記録した時間と最後に検知した時間の差が10分以内なら電気を点けない
		if dur.Hours() < 2.0 && dur.Minutes() < 10 {
			return 200, map[string]string{"message": "This request is within 10 minute of last detection."}
		}
	}

	m, e := u.ReadJSONFile(c.Config.WebSetting.JSONFilePath + "/tplink.json")

	if e != nil {
		log.Println(e)
		return 400, map[string]string{"message": "An error has occured on reading json file."}
	}

	go ChangeScreenStatus(On)
	go ChangeLightbulbIlumination(m["bulb-ceiling-1"], 100, 2000)
	go ChangeLightbulbIlumination(m["bulb-ceiling-2"], 100, 2000)
	go ChangeLightbulbIlumination(m["bulb-ceiling-3"], 100, 2000)
	go ChangeLightbulbIlumination(m["bulb-ceiling-4"], 100, 2000)
	time.Sleep(2)
	repo.InsertLog("TurnOnLightOnStandByMyHome", "OK")

	return 204, map[string]string{}
}

// AdjustLightBeforeSleeping adjust lightbulbs for sleeping before sleeping.
func AdjustLightBeforeSleeping() {
	m, e := u.ReadJSONFile(c.Config.WebSetting.JSONFilePath + "/tplink.json")

	if e != nil {
		log.Println(e)
		return
	}

	// 寝る前に電気を色々するシナリオ
	go ChangeScreenStatus(Off)
	go TurnOffLightBulb(m["bulb-ceiling-1"], 2000)
	go TurnOffLightBulb(m["bulb-ceiling-2"], 2000)
	go TurnOffLightBulb(m["bulb-ceiling-3"], 2000)
	TurnOffLightBulb(m["bulb-ceiling-4"], 2000)
	ChangeLightbulbIlumination(m["bulb-living-1"], 4, 2000)
	repo.InsertLog("AdjustLightBeforeSleeping", "OK")
}

func TurnOnTVonMorning() {
	// 一定範囲内の時間にセンサに反応があったらテレビを点ける
	// 22 - 5時？
	// amah := CallAmah()
	now := time.Now()

	now.AddDate(0, 0, -1)

	// applieances, err := amah.NatureremoClient.ApplianceService.GetAll(amah.Context)

	// if err != nil {
	// 	return
	// }

}

// AdjustLightBrightness changes lightbulbs status by hour.
func AdjustLightBrightness() {
	repo.InsertLog("AdjustLightBrightness", "START")
	log.Println(u.LogBracket)
	log.Println("AdjustLightBrightness started")
	defer log.Println(u.LogBracket)

	execLightbulbCmd := func(ips []map[string]string, c *structs.TplinkLightbulbCmd) {
		for cnt, v := range ips {
			c.IP = v["ip"]
			if cnt-1 != len(ips) {

				if v["colorTemp"] == fmt.Sprint(c.Kelvin) {
					continue
				}

				log.Printf("%-8s : %s %s", "bulb ip", v, strconv.Itoa(c.Kelvin))
				util.TplinkLightbulbExecute(c)
			}
		}
	}
	// 時刻に合わせて電気の明るさを変える
	// 23-00なら少し暗くする
	// 00-01ならPCと反対側の電球を更に暗くする
	t := time.Now()

	c := structs.TplinkLightbulbCmd{}
	c.Command = structs.Temp
	c.Transition = 3000

	bulbState := GetActiveLightbulbIP()

	// Finish when Any lightbulb has not been off.
	if len(bulbState) == 0 {
		log.Println("No light bulbs were available.")
		return
	}

	log.Println("Lightbulbs have been changed.")

	switch {
	case 18 <= t.Hour() && t.Hour() <= 22:
		c.Kelvin = 3500
		log.Printf("Kelvin : %d", c.Kelvin)
		execLightbulbCmd(bulbState, &c)
	case 23 <= t.Hour():
		c.Kelvin = 3000
		log.Printf("Kelvin : %d", c.Kelvin)
		execLightbulbCmd(bulbState, &c)
	case 0 <= t.Hour() && t.Hour() <= 2:
		c.Kelvin = 2500
		log.Printf("Kelvin : %d", c.Kelvin)
		execLightbulbCmd(bulbState, &c)
	case t.Hour() < 18:
		c.Kelvin = 4000
		log.Printf("Kelvin : %d", c.Kelvin)
		execLightbulbCmd(bulbState, &c)
	default:
		log.Println("Now is out of target time.")
	}
}

// ToggleBedlight is turning bedlight off.
func ToggleBedlight() {
	m, e := u.ReadJSONFile(c.Config.WebSetting.JSONFilePath + "/tplink.json")

	if e != nil {
		log.Println(e)
		return
	}

	TurnOffLightBulb(m["bulb-living-1"], 2000)

	repo.InsertLog("Sleeping", "OK")
}
