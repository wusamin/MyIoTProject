package service

import (
	"log"
	c "maid/pkg/config"
	repo "maid/pkg/repository"
	"maid/pkg/structs"
	"time"
)

// GetSensorValueRecorded finds value sensor recorded between "from" and "to".
func GetSensorValueRecorded(sensorType string, selectType string, deviceID string, from time.Time, to time.Time) []*structs.SensorAPIReturn {
	sensorValues := repo.FindSensorValRanged(sensorType, deviceID, from, to)
	switch sensorType {
	case "mo":
		for _, sensorVal := range sensorValues {
			dur := sensorVal.RecordedAt.Sub(sensorVal.CreatedAt)
			if 4*time.Minute < dur {
				sensorVal.Val = "0"
			}
		}
	default:
	}

	var ret []*structs.SensorAPIReturn

	for _, v := range sensorValues {
		s := structs.SensorAPIReturn{}
		s.CreatedAt = v.CreatedAt
		s.DeviceID = v.DeviceID
		s.SensorType = v.SensorType
		s.Val = v.Val
		s.RecordedAt = v.RecordedAt

		ret = append(ret, &s)
	}

	return ret
}

func GetSensorValue(sensorType string, selectType string, deviceID string, from time.Time, to time.Time, recordNum int) []*structs.TNatureRemoSensor {
	sensorValues := repo.FindSensorValRanged(sensorType, deviceID, from, to)

	log.Println(len(sensorValues))

	return sensorValues
}

// PurgeSensorVal is purging sensor val records.
func PurgeSensorVal() {
	log.Println("-----------------------------------------")
	log.Println("Purge start.")

	if err := repo.MoveSensorVal(time.Now().AddDate(0, 0, -c.Config.MaidSetting.PurgeDate)); err != nil {
		log.Println(err)
	}
	log.Println("SensorVal was purged.")

	if err := repo.RemoveSystemTempRecords(); err != nil {
		log.Println(err)
	}
	log.Println("SystemTempRecords was purged.")

	if err := repo.RemoveCO2Records(); err != nil {
		log.Println(err)
	}
	log.Println("SystemTempRecords was purged.")

	if err := repo.RemoveMotionSensorVal(time.Now().AddDate(0, 0, -c.Config.MaidSetting.PurgeDate)); err != nil {
		log.Println(err)
	}
	log.Println("MotionSensorVal was purged.")

	log.Println("-----------------------------------------")
}
