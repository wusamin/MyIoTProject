package structs

import (
	"context"
	"time"

	"github.com/tenntenn/natureremo"
)

// Config is root object of setting.
type Config struct {
	WebSetting  WebSettings
	DataBase    DatabaseSetting
	MaidSetting MaidSettings
	LightBulb   LightBulbs
	Manuscript  Manuscripts
}

// DatabaseSetting is values of database setting.
type DatabaseSetting struct {
	DBMS     string `toml:"dbms"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	Protocol string `toml:"protocol"`
	Dbname   string `toml:"dbname"`
	Option   string `toml:"option"`
}

// WebSettings is value of setting for web system.
type WebSettings struct {
	IP                    string `toml:"ip"`
	Port                  string `toml:"port"`
	ContainerPort         string `toml:"containerPort"`
	NatureRemoToken       string `toml:"natureRemoToken"`
	NatureRemoDeviceID    string `toml:"natureRemoDeviceID"`
	WelcomebackHour       int    `toml:"welcomebackHour"`
	StandByVoice          string `toml:"standByVoice"`
	SpeakingGoogleHomeURL string `toml:"speakingGoogleHomeUrl"`
	StaticFilePath        string `toml:"staticFilePath"`
	TplinkAddress         string `toml:"tplinkAddress"`
	TplinkPassword        string `toml:"tplinkPassword"`
	SpeechVoiceURL        string `toml:"speechVoiceURL"`
	TplightPath           string `toml:"tplightPath"`
	SesameToken           string `toml:"sesameToken"`
	SesameID              string `toml:"sesameID"`
	SesamePrivateToken    string `toml:"sesamePrivateToken"`
	TabletPushURL         string `toml:"tabletPushURL"`
	Iws600cmPath          string `toml:"iws600cmPath"`
	JSONFilePath          string `toml:"jsonFilePath"`
	DashboardPushURL      string `toml:"dashboardPushURL"`
	DashboardPushKey      string `toml:"dashboardPushKey"`
	ShPath                string `toml:"shPath"`
}

// MaidSettings is value of setting for maid.
type MaidSettings struct {
	OffTimeCiculator           int     `toml:"offTimeCiculator"`
	OnTimeCiculator            int     `toml:"onTimeCiculator"`
	BorderSensorValTurnOnLight string  `toml:"borderSensorValTurnOnLight"`
	PurgeDate                  int     `toml:"purgeDate"`
	TurningFanOn               float64 `toml:"turningFanOn"`
	TurningFanOff              float64 `toml:"turningFanOff"`
}

// LightBulbs represents IP address of lightbulb.
type LightBulbs struct {
	BulbLiving1  string `toml:"bulb-living-1"`
	BulbCeiling1 string `toml:"bulb-ceiling-1"`
	BulbCeiling2 string `toml:"bulb-ceiling-2"`
	BulbCeiling3 string `toml:"bulb-ceiling-3"`
	BulbCeiling4 string `toml:"bulb-ceiling-4"`
}

// Manuscripts represents manuscripts voicroid talks.
type Manuscripts struct {
	Boot string `toml:"boot"`
}

// MTimeSignal is binded to M_TIME_SIGNAL
type MTimeSignal struct {
	TheHour int
	HourID  int
	Text    string
	InsDate time.Time
	UpdDate time.Time
}

// TApiCalledLog is binded to T_API_CALLED_LOG
type TApiCalledLog struct {
	APIName         string    `db:"API_NAME"`
	OperationResult string    `db:"operation_result"`
	InsDate         time.Time `db:"INS_DATE"`
}

// TNatureRemoSensor is binded to t_nature_remo_sensor
type TNatureRemoSensor struct {
	RecordedAt time.Time `db:"recorded_at"`
	DeviceID   string    `db:"device_id"`
	SensorType string    `db:"sensor_type"`
	Val        string    `db:"val"`
	CreatedAt  time.Time `db:"created_at"`
	InsDate    time.Time `db:"ins_date"`
	UpdDate    time.Time `db:"upd_date"`
}

// MHomeAppliances is binded to m_home_appliances.
type MHomeAppliances struct {
	DeviceName string    `db:"device_name"`
	DeviceType string    `db:"device_type"`
	Status     string    `db:"status"`
	UpdateDate time.Time `db:"update_date"`
	InsDate    time.Time `db:"ins_date"`
	UpdDate    time.Time `db:"upd_date"`
}

// TNobyAPI is binded to T_NOBY_API
type TNobyAPI struct {
	CallText     string    `db:"call_text"`
	CommandID    string    `db:"command_id"`
	CommandName  string    `db:"command_name"`
	Text         string    `db:"text"`
	ReturnType   string    `db:"return_type"`
	Mood         float32   `db:"mood"`
	Negaposi     float32   `db:"negaposi"`
	NegaposiList string    `db:"negaposi_list"`
	Emotion      string    `db:"emotion"`
	EmotionList  string    `db:"emotion_list"`
	WordList     string    `db:"word_list"`
	Art          string    `db:"art"`
	Org          string    `db:"org"`
	Psn          string    `db:"psn"`
	Loc          string    `db:"loc"`
	Dat          string    `db:"dat"`
	Tim          string    `db:"tim"`
	InsDate      time.Time `db:"ins_date"`
	UpdDate      time.Time `db:"upd_date"`
}

// Amah has objects to control devices.
type Amah struct {
	Context          context.Context
	NatureremoClient *natureremo.Client
}

// TMaidTimerScheduled is binded t_maid_timer_scheduled.
type TMaidTimerScheduled struct {
	ID          int64     `db:"id"`
	TargetURL   string    `db:"target_url"`
	TimerVal    int64     `db:"timer_val"`
	StartUpTime time.Time `db:"start_up_time"`
	Result      string    `db:"result"`
	Deleted     string    `db:"deleted"`
	InsDate     time.Time `db:"ins_date" gorm:"primary_key"`
	UpdDate     time.Time `db:"upd_date"`
}

type DarkSkyOutput struct {
	Summary        string `json:"summary"`
	NowSummary     string `json:"nowSummary"`
	NowTemperature string `json:"nowTemperature"`
	Icon           string `json:"icon"`
	NowHumidity    string `json:"nowHumidity"`
}

// TimerRequest represents request parameter for timer api.
type TimerRequest struct {
	URL       string            `json:"url"`
	Time      int64             `json:"time"`
	Parameter map[string]string `json:"parameter"`
	Query     map[string]string `json:"query"`
}

type SensorAPIReturn struct {
	RecordedAt time.Time
	DeviceID   string
	SensorType string
	Val        string
	CreatedAt  time.Time
}
