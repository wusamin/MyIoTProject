package service

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"strconv"
	"time"

	"maid/pkg/structs"

	"github.com/go-resty/resty/v2"
	"github.com/ikasamah/homecast"
	"github.com/wusamin/dsnip"
	nobyapi "github.com/wusamin/go-noby-api"

	repo "maid/pkg/repository"

	c "maid/pkg/config"
)

// TimeSignal returns manuscript for timesignal.
func TimeSignal() string {
	time := time.Now()

	if 59 <= time.Minute() {
		return repo.SelectTimeSignal(time.Hour() + 1)
	}

	return repo.SelectTimeSignal(time.Hour())
}

// TimeSignalTest returns manuscript for timesignal.
func TimeSignalTest(hour int) string { return repo.SelectTimeSignal(hour) }

// TimeSignalWeather get text of weatherReport
func TimeSignalWeather() string {

	result := WeatherReport()
	// forecasts := result["forecasts"].([]interface{})[0].(map[string]interface{})
	forecasts := result["daily"].([]interface{})[0].(map[string]interface{})

	text := forecasts["weather"].([]interface{})[0].(map[string]interface{})["description"]
	temp := forecasts["temp"].(map[string]interface{})
	max := temp["max"]
	min := temp["min"]

	return "ゼロゼロゼロゼロです。明日の天気は" + text.(string) + "です。" + "最高気温は" + strconv.FormatFloat(max.(float64), 'f', 1, 64) + "度、最低気温は" + strconv.FormatFloat(min.(float64), 'f', 1, 64) + "度です。"
}

// StandByMe make Google Home to say "Welcome back!" by iws600cm value.
func StandByMe() {
	now := time.Now()
	// センサーが検知した、検知していないときのそれぞれの最新のレコードを取得する
	selected := repo.FindSensorValForStandBy()

	if len(selected) < 2 {
		return
	}

	ret0 := selected[0]
	ret1 := selected[1]
	s := ret0.RecordedAt.Sub(ret1.RecordedAt)

	border := 1 * time.Hour

	// 前回検知した時間から3時間以上経過しているか
	if s < border {
		return
	}

	if 5*time.Second < now.Sub(ret0.RecordedAt) {
		return
	}

	log.Println("-----------------------------------------")
	log.Println("A voice 'StandByMe' casted ")
	CastVoice2Dashboard(c.Config.WebSetting.StandByVoice)
	log.Println("-----------------------------------------")
}

// CastBootVoice cast boot voice to google home.
func CastBootVoice() {
	sendVoiceManuscript(c.Config.Manuscript.Boot)
}

// CastVoiceDashboard casts voice of manuscript to dashboard.
func CastVoiceDashboard(manuscript string) {
	sendVoiceManuscript(manuscript)
}

// BroadcastMessageToDashboard broadcast strings to dashboard.
func BroadcastMessageToDashboard(text string) {
	resty.New().R().SetQueryParam(
		"message", text,
	).Get("")
}

func CastVoice2Dashboard(rawurl string) {
	resty.New().R().SetQueryParam(
		"message", rawurl,
	).Get("")
}

func castVoiceForGoogleHome(rawurl string) {
	ctx := context.Background()
	devices := homecast.LookupAndConnect(ctx)

	u, err := url.Parse(rawurl)

	if err != nil {
		return
	}

	var isFound bool = false

	for _, device := range devices {
		log.Printf("%-8s : %s", "Address", device.AddrV4)
		log.Printf("%-8s : %d", "Port", device.Port)
		log.Printf("%-8s : %s", "Name", device.Name)
		fmt.Printf("Device: [%s:%d]%s", device.AddrV4, device.Port, device.Name)

		if err := device.Play(ctx, u); err != nil {
			log.Printf("Failed to speak: %v", err)
		}

		isFound = true
	}

	if !isFound {
		log.Println("Google Home is not found.")
	}
}

func TestVoiceroid() {
	castVoiceForGoogleHome(c.Config.WebSetting.StandByVoice)
}

// sendVoiceManuscript delivers manuscript for voiceroid.
func sendVoiceManuscript(manuscript string) {
	client := resty.New()

	_, err := client.SetTimeout(2*time.Second).
		R().
		SetQueryParams(map[string]string{"manuscript": manuscript}).
		SetHeader("Accept", "application/json").
		Get(c.Config.WebSetting.SpeechVoiceURL)

	if err != nil {
		log.Println(err)
	}
}

// SuggestTurnOnLight send voice "Do you want to turn on room lights ?" to PC when room becomes be getting dark at evening.
func SuggestTurnOnLight() {
	// 時刻が適正な範囲か
	nowHour := time.Now().Hour()

	// 時刻が指定の範囲の場合のみ実行する
	if !(15 <= nowHour && nowHour <= 18) {
		return
	}

	// natureremoの人感センサで5分以上いない場合は実行しない
	if v := repo.FindSensorVals("natureremo", "mo", 1)[0]; 5*time.Minute < v.RecordedAt.Sub(v.CreatedAt) {
		return
	}

	borderLightOn, err := strconv.ParseFloat(c.Config.MaidSetting.BorderSensorValTurnOnLight, 64)
	if err != nil {
		log.Println(err)
		return
	}

	ilValNow := repo.FindSensorVal("il", 1)

	sensorValNow, err := strconv.ParseFloat(ilValNow[0].Val, 64)

	// 最新の照度センサの値が閾値より上か。上の場合は終了する
	if borderLightOn <= sensorValNow {
		return
	}

	// 照度センサの値は閾値以下か
	ilVal := repo.FindSensorValAverage("natureremo", "il")
	sensorVal, err := strconv.ParseFloat(ilVal.Val, 64)

	if err != nil {
		log.Println("Parsing ilmination sensor value to float64.")
		log.Println(err)
		return
	}

	if !(sensorVal < borderLightOn) {
		return
	}

	// テキストを送る
	sendVoiceManuscript("少し暗くなってきましたね。電気を点けてはいかがですか？")
}

// RegisterSleepVoiceSchedule register schedules casting sleeping voice.
func RegisterSleepVoiceSchedule() {
	// 最新のおやすみAPIの時間を取得する
	latest := dsnip.AddDay(repo.FindLatestLog("AdjustLightBeforeSleeping").InsDate, 1)

	// 昨日寝た時間＋1日
	firstTime := dsnip.AddMinute(latest, -15)
	log.Printf("first:%v", firstTime)

	second := dsnip.AddMinute(latest, -5)
	log.Printf("second:%v", second)

	first := structs.TimerRequest{}

	first.URL = c.Config.WebSetting.SpeechVoiceURL + "?manuscript=" + url.QueryEscape("そろそろ寝る時間ですよ。寝ましょうね。")
	log.Printf("%04v%02v%02v%02v%02v%02v", firstTime.Year(), int(firstTime.Month()), firstTime.Day(), firstTime.Hour(), firstTime.Minute(), firstTime.Second())
	time, _ := strconv.ParseInt(fmt.Sprintf("%04v%02v%02v%02v%02v%02v", firstTime.Year(), int(firstTime.Month()), firstTime.Day(), firstTime.Hour(), firstTime.Minute(), firstTime.Second()), 10, 64)
	first.Time = time
	RegisterSchedule(&first)

	first.URL = c.Config.WebSetting.SpeechVoiceURL + "?manuscript=" + url.QueryEscape("昨日はこのくらいの時間に寝ていましたよ。さあ、寝ましょう。")
	log.Printf("%04v%02v%02v%02v%02v%02v", second.Year(), int(second.Month()), second.Day(), second.Hour(), second.Minute(), second.Second())
	time, _ = strconv.ParseInt(fmt.Sprintf("%04v%02v%02v%02v%02v%02v", second.Year(), int(second.Month()), second.Day(), second.Hour(), second.Minute(), second.Second()), 10, 64)
	first.Time = time
	RegisterSchedule(&first)
}

func BroadcastConversationText() {
	req := nobyapi.NobyRequest{}
	req.Appkey = ""
	req.Text = "夜だね"
	req.Persona = nobyapi.Normal

	if res, err := nobyapi.Call(&req); err != nil {
		log.Println(err)
		return
	} else {
		log.Println(res)
		BroadcastMessageToDashboard(res.Text)
	}
}
