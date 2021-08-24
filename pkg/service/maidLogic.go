package service

import (
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/tenntenn/natureremo"
	"github.com/wusamin/dsnip"
	tplink "github.com/wusamin/tplink-api"

	c "maid/pkg/config"
	"maid/pkg/repository"
	"maid/pkg/structs"
	"maid/pkg/util"
	u "maid/pkg/util"
)

// Test is testing method.
func Test() {

}

// ControlDevicePower makes status specified by deviceType to be "power".
func ControlDevicePower(deviceType string, power string) {
	log.Printf("%-11s:%-8s", "device type", deviceType)
	log.Printf("%-11s:%-8s", "status", power)

	switch deviceType {
	case "tv":
		turnOnTV()
		return
	}

	amah := u.CallAmah()
	appliances, err := amah.NatureremoClient.ApplianceService.GetAll(amah.Context)

	// 機器情報取得でエラーがあった場合は終了
	if err != nil {
		log.Println(err)
		return
	}

	switch deviceType {
	case "aircon":
		for _, appliance := range appliances {
			if appliance.Type == natureremo.ApplianceTypeAirCon {
				if power == "off" {
					appliance.AirConSettings.Button = natureremo.ButtonPowerOff
				} else if power == "on" {
					appliance.AirConSettings.Button = natureremo.ButtonPowerOn
				} else {
					return
				}

				amah.NatureremoClient.ApplianceService.UpdateAirConSettings(amah.Context, appliance, appliance.AirConSettings)
			}
		}
	}
}

// DevicePower represents power status for devices.
type DevicePower string

// this represents power status for device.
const (
	PowerOn  DevicePower = "on"
	PowerOff DevicePower = "off"
)

// CollateReuqstToken collate token with token this app has.
func CollateReuqstToken(token string) bool {
	hashed := sha512.Sum512([]byte(token))
	return hex.EncodeToString(hashed[:]) == c.Config.WebSetting.SesamePrivateToken
}

// ScreenStatus represents dashboard screen status.
type ScreenStatus string

const (
	// On represents dashboard screen is on.
	On ScreenStatus = "s1"

	// Off represents dashboard screen is off.
	Off ScreenStatus = "so"
)

// ChangeScreenStatus change tablet screen.
func ChangeScreenStatus(destStatus ScreenStatus) {
	client := resty.New()

	_, err := client.R().
		SetHeader("Accept", "application/json").
		SetQueryParams(map[string]string{"key": c.Config.WebSetting.DashboardPushKey}).
		SetQueryParams(map[string]string{"message": string(destStatus)}).
		Get(c.Config.WebSetting.DashboardPushURL)

	if err != nil {
		log.Println(err)
	}
}

// TurnOnTV is pushing power button of TV remocon.
func turnOnTV() {

	turnOn := func() {
		client := resty.New()

		_, err := client.R().
			SetHeader("Accept", "application/json").
			SetHeader("Authorization", "Bearer "+c.Config.WebSetting.NatureRemoToken).
			SetFormData(map[string]string{"button": "power"}).
			Post("")

		if err != nil {
			log.Println(err)
		}
	}

	latest := repository.FindLatestLog("AdjustLightBeforeSleeping")

	log.Println(latest.InsDate)

	diff := time.Now().Sub(latest.InsDate)

	log.Println(diff)

	if diff <= 12*time.Hour {
		turnOn()
	} else {
		return
	}
}

// ControlFan controls cooling fan for raspberry pi.
func ControlFan() {
	// 部屋に人がおらず、ファンが止まっていたらファンを起動して終了する
	// 部屋に人がおらず、ファンが動いていたら終了する
	// 寝入るまでは止める

	// 3分間のCPU温度の平均が閾値未満であればファンを止める

	vals := repository.FindSensorVals("system_temp", "cpu", 10)

	// センサ検索結果の値の平均を取得する
	getAverage := func(a []*structs.TNatureRemoSensor) float64 {
		var tempSum float64

		for _, v := range a {
			f, e := strconv.ParseFloat(v.Val, 64)
			if e != nil {
				log.Println(e)
				continue
			}
			tempSum = tempSum + f
		}
		return tempSum / float64(len(a))
	}

	api, _ := tplink.Connect(c.Config.WebSetting.TplinkAddress, c.Config.WebSetting.TplinkPassword)
	hs105, err := api.GetHS105("plug-living-1")

	if err != nil {
		log.Println(err)
		return
	}

	info, _ := hs105.GetInfo()

	// 規定期間のCPU温度の平均が閾値以上であればファンを動かす
	// 温度が高い状態はなるべく短くしたいため、直近３分間の温度の平均を取る
	over := vals[0:3]
	tempAve := getAverage(over)

	if c.Config.MaidSetting.TurningFanOn <= tempAve && info.OnTime == 0 {

		// repository.InsertLog("system_fan_control", "Turning fan on")
		toggleSmartplug("", PowerOn)
		log.Println("-----------------------------------------")
		log.Println("Turning fan on.")
		log.Printf("%-8s : %s", "status", string(PowerOn))
		log.Printf("%-8s : %s", "temp", util.TransFloat(tempAve)+"°C")
		log.Println("-----------------------------------------")
		return
	}

	// 閾値以下の場合はファンを消す
	// ファンはなるべく長く動かしたいため、平均時間を長めに取る
	tempAve2 := getAverage(vals)

	if tempAve2 <= c.Config.MaidSetting.TurningFanOff && 0 < info.OnTime {
		// repository.InsertLog("system_fan_control", "Turning fan off")
		toggleSmartplug("", PowerOff)
		log.Println("-----------------------------------------")
		log.Println("Turning fan off.")
		log.Printf("%-8s : %s", "status", string(PowerOff))
		log.Printf("%-8s : %s", "temp", util.TransFloat(tempAve)+"°C")
		log.Println("-----------------------------------------")
		return
	}
}

func CallNobyAPI() {
	client := resty.New()

	r, err := client.R().
		SetHeader("Accept", "application/json").
		SetHeader("Authorization", "Bearer "+c.Config.WebSetting.NatureRemoToken).
		SetQueryParams(map[string]string{
			"appkey": "",
			"text":   "なにか話して",
		}).
		Get("https://app.cotogoto.ai/webapi/noby.json")

	if err != nil {
		fmt.Println(err)
		return
	}

	var ret structs.NobyResponse

	json.Unmarshal(r.Body(), &ret)

	fmt.Println(ret.Text)
}

// BroadcastTimesignal broadcasts timesignal voice and text.
func BroadcastTimesignal(hour int) {
	var text string

	if hour == 0 {
		text = TimeSignalWeather()
	} else {
		text = TimeSignalTest(hour)
	}
	resty.New().R().SetQueryParam(
		"message", strings.ReplaceAll(text, "ゼロゼロゼロゼロです。", ""),
	).Get("")

	sendVoiceManuscript(text)
}

// RegisterTimeSignal creates timer record of timesignal per an hour of tomorrow.
func RegisterTimeSignal() {
	targetDay := dsnip.AddDay(time.Now(), 1)

	first := structs.TimerRequest{}

	// log.Printf("%04v%02v%02v%02v%02v%02v", targetDay.Year(), int(targetDay.Month()), targetDay.Day(), 9, 59, 59)

	for i := 0; i < 24; i++ {
		first.URL = fmt.Sprint("", i)
		time, _ := strconv.ParseInt(fmt.Sprintf("%04v%02v%02v%02v%02v%02v", targetDay.Year(), int(targetDay.Month()), targetDay.Day(), i, 0, 0), 10, 64)
		first.Time = time
		RegisterSchedule(&first)
	}
}
