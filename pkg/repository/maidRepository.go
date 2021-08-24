package repository

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/go-gorp/gorp"
	"github.com/jinzhu/gorm"

	c "maid/pkg/config"
	"maid/pkg/structs"
)

func dbConnect() (*gorp.DbMap, error) {
	if conn, err := sqlOpen(); err != nil {
		return nil, err
	} else {
		return &gorp.DbMap{Db: conn, Dialect: gorp.MySQLDialect{Engine: "InnoDB", Encoding: "UTF8"}}, nil
	}

}

func sqlOpen() (*sql.DB, error) {
	DBMS := c.Config.DataBase.DBMS
	USER := c.Config.DataBase.User
	PASS := c.Config.DataBase.Password
	PROTPCOL := c.Config.DataBase.Protocol
	DBNAME := c.Config.DataBase.Dbname
	OPTION := c.Config.DataBase.Option
	CONNECT := USER + ":" + PASS + "@" + PROTPCOL + "/" + DBNAME + "?" + OPTION

	db, err := sql.Open(DBMS, CONNECT)

	if err != nil {
		return nil, err
	}
	return db, nil
}

func connectGorm() (*gorm.DB, error) {
	DBMS := c.Config.DataBase.DBMS
	USER := c.Config.DataBase.User
	PASS := c.Config.DataBase.Password
	PROTPCOL := c.Config.DataBase.Protocol
	DBNAME := c.Config.DataBase.Dbname
	OPTION := c.Config.DataBase.Option
	CONNECT := USER + ":" + PASS + "@" + PROTPCOL + "/" + DBNAME + "?" + OPTION

	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		return nil, err
	}
	return db, nil
}

// SelectTimeSignal returns for manuscript of time signal.
func SelectTimeSignal(hour int) string {
	dbmap, err := dbConnect()
	if err != nil {
		log.Println(err)
		return ""
	}
	defer dbmap.Db.Close()

	var dto structs.MTimeSignal

	// 本当はエラーハンドリングが必要
	dbmap.SelectOne(&dto, `
		select
			m.*
		from
			M_TIME_SIGNAL m
				cross join (
					select
						floor(RAND() * count(*)) as rand
					from
						M_TIME_SIGNAL
					where
						THE_HOUR = ?
				) as rand
		where
			THE_HOUR = ?
			and HOUR_ID = rand.rand
	`, hour, hour)

	return dto.Text
}

// FindSensorVal gets sensor val specified sensorType.
func FindSensorVal(sensorType string, limit int) []*structs.TNatureRemoSensor {
	dbmap, err := dbConnect()
	if err != nil {
		log.Println(err)
		return []*structs.TNatureRemoSensor{}
	}

	defer dbmap.Db.Close()

	var dto structs.TNatureRemoSensor

	selected, err := dbmap.Select(&dto,
		`select
			recorded_at
			, device_id
			, sensor_type
			, val
			, created_at
			, ins_date
			, upd_date 
		from
			t_nature_remo_sensor 
		where
			sensor_type = ? 
		order by
			recorded_at desc 
		limit
			?
	`, sensorType, limit)

	if err != nil {
		fmt.Println(err)
	}

	var ret []*structs.TNatureRemoSensor

	for _, p := range selected {
		if q, ok := p.(*structs.TNatureRemoSensor); ok {
			ret = append(ret, q)
		}
	}

	return ret
}

func FindSensorVals(deviceID string, sensorType string, limit int) []*structs.TNatureRemoSensor {
	db, err := connectGorm()
	if err != nil {
		log.Println(err)
		return []*structs.TNatureRemoSensor{}
	}

	defer db.Close()

	var r []*structs.TNatureRemoSensor
	db.Table("t_nature_remo_sensor").
		Where("device_id = ? and sensor_type = ?", deviceID, sensorType).
		Order("recorded_at desc").
		Limit(limit).
		Find(&r)

	return r
}

// SelectSensorValForStandBy is for "StandBy" method searching for structs.TNatureRemoSensor .
func SelectSensorValForStandBy() []interface{} {
	dbmap, err := dbConnect()
	if err != nil {
		log.Println(err)
		return []interface{}{}
	}

	defer dbmap.Db.Close()

	var dto structs.TNatureRemoSensor

	selected, err := dbmap.Select(&dto,
		`select
			recorded_at
			, device_id
			, sensor_type
			, val
			, created_at
			, ins_date
			, upd_date 
		from
			t_nature_remo_sensor 
		where
			sensor_type = 'mo' 
			and device_id = 'natureremo'
		order by
			recorded_at desc 
		limit
			3
	`)

	if err != nil {
		fmt.Println(err)
	}

	return selected
}

func FindSensorValForStandBy() []*structs.TNatureRemoSensor {
	db, err := connectGorm()
	if err != nil {
		log.Println(err)
		return []*structs.TNatureRemoSensor{}
	}

	defer db.Close()

	var r []*structs.TNatureRemoSensor
	db.Table("t_motion_sensor").
		Where("device_id = ? and val = ?", "iws600cm", "2").
		Order("recorded_at desc").
		Limit(2).
		Find(&r)

	return r
}

// FindSensorValAverage gets sensor value average.
func FindSensorValAverage(deviceID string, sensorType string) structs.TNatureRemoSensor {
	dbmap, e := dbConnect()
	if e != nil {
		log.Println(e)
		return structs.TNatureRemoSensor{}
	}

	defer dbmap.Db.Close()

	var dto structs.TNatureRemoSensor

	// 過去6レコードの平均値を取得する
	err := dbmap.SelectOne(&dto,
		`
		select
			avg(t.val) as val
		from ( 
			select
				val
			from
				t_nature_remo_sensor 
			where
				sensor_type = ? 
				and device_id = ?
			order by
				recorded_at desc 
			limit 6
		) t
		`, sensorType, deviceID)

	if err != nil {
		fmt.Println(err)
	}

	return dto
}

// FindDeviceStatus searches m_home_appliances for records specified deviceName.
func FindDeviceStatus(deviceName string) structs.MHomeAppliances {
	dbmap, e := dbConnect()
	if e != nil {
		log.Println(e)
		return structs.MHomeAppliances{}
	}

	dbmap.AddTableWithName(structs.MHomeAppliances{}, "m_home_appliances")
	defer dbmap.Db.Close()

	var dto structs.MHomeAppliances

	err := dbmap.SelectOne(&dto,
		`
		select
			device_name,
			device_type,
			status,
			update_date,
			INS_DATE,
			UPD_DATE  
		from
			m_home_appliances 
		where
			device_type = ?
		`, deviceName)

	if err != nil {
		println(err.Error())
	}

	return dto
}

// FindSensorValRanged searches t_natureremo_sensor for records specified by sensorType within range of "from" and "to".
func FindSensorValRanged(sensorType string, deviceID string, from time.Time, to time.Time) []*structs.TNatureRemoSensor {
	dbmap, err := dbConnect()
	if err != nil {
		log.Println(err)
		return []*structs.TNatureRemoSensor{}
	}

	defer dbmap.Db.Close()

	var dto structs.TNatureRemoSensor
	// 条件で指定された範囲のレコードを2レコード毎に纏める
	selected, err := dbmap.Select(&dto,
		`
				select
					recorded_at
					, sensor_type
					, val
					, created_at
				from
					t_nature_remo_sensor, (select @rownum:=0) as dummy
				where
					sensor_type = ? 
					and device_id = ?
					and recorded_at between ? and ?
				order by
					recorded_at
	`, sensorType, deviceID, from, to)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	var retList []*structs.TNatureRemoSensor

	for _, tNatureRemoSensor := range selected {
		if v, ok := tNatureRemoSensor.(*structs.TNatureRemoSensor); ok {
			retList = append(retList, v)
		}
	}

	return retList
}

func FindSensorValBetween(sensorType string, from time.Time, to time.Time) []*structs.TNatureRemoSensor {
	gorm, err := connectGorm()
	if err != nil {
		log.Println(err)
		return []*structs.TNatureRemoSensor{}
	}

	defer gorm.Close()

	var ret []*structs.TNatureRemoSensor

	gorm.Table("t_nature_remo_sensor").
		Where("sensor_type = ? and recorded_at between ? and ?", sensorType, from, to).
		Find(&ret)

	return ret
}

// FindScheduledURL select timer table.
func FindScheduledURL() []*structs.TMaidTimerScheduled {
	gorm, err := connectGorm()
	if err != nil {
		log.Println(err)
		return []*structs.TMaidTimerScheduled{}
	}

	defer gorm.Close()

	var ret []*structs.TMaidTimerScheduled

	gorm.Table("t_maid_timer_scheduled").
		Where("deleted = ?", "0").
		Find(&ret)

	return ret
}

func UpdateScheduledURL(pk string) error {
	gorm, err := connectGorm()
	if err != nil {
		log.Println(err)
		return err
	}

	defer gorm.Close()

	gorm.Table("t_maid_timer_scheduled").
		Where("pk = ?", pk)
	// Update()

	return nil
}

func FindWeatherForecast() *structs.OneCall {
	bytes, err := ioutil.ReadFile(c.Config.WebSetting.JSONFilePath + "/open-weather-map-onecall.json")
	if err != nil {
		// return map[string]string{}, err
	}
	var r structs.OneCall

	if err := json.Unmarshal(bytes, &r); err != nil {
		log.Println("An error has occuerd.")
		log.Println(err)
		// return map[string]string{}, err
	}

	return &r
}
