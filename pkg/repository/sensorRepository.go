package repository

import (
	"log"
	"time"
)

// MoveSensaorVal is moving sensor val records work table to purge table.
func MoveSensorVal(d time.Time) error {
	db, err := dbConnect()
	if err != nil {
		return err
	}

	defer db.Db.Close()

	if _, err := db.Query(`
		insert into purge_sensor_val
		select *
		from t_nature_remo_sensor
		where recorded_at < ?`, d); err != nil {
		log.Println(err)
		log.Println("An error has occured on executing MoveSensaorVal select-insert.")
		return err
	}
	//where recorded_at < str_to_date('2019-8-17 00:30:00', '%Y-%m-%d %T')`)

	if _, err := db.Query(`
		delete from t_nature_remo_sensor
		where recorded_at < ?
		`, d); err != nil {
		log.Println(err)
		log.Println("An error has occured on executing MoveSensaorVal delete.")
		return err
	}

	return nil
}

func RemoveSystemTempRecords() error {
	db, err := dbConnect()
	if err != nil {
		return err
	}

	defer db.Db.Close()

	if _, err := db.Query(`
	delete 
	from
		purge_sensor_val
	where device_id = 'system_temp'
		and(
			DATE_FORMAT(recorded_at, '%i') in ('00', '02', '04', '06', '08') 
			or DATE_FORMAT(recorded_at, '%i') in ('20', '22', '24', '26', '28')
			or DATE_FORMAT(recorded_at, '%i') in ('30', '32', '34', '36', '38')
			or DATE_FORMAT(recorded_at, '%i') in ('40', '42', '44', '46', '48')
			or DATE_FORMAT(recorded_at, '%i') in ('50', '52', '54', '56', '58')
			or DATE_FORMAT(recorded_at, '%i') in ('10', '12', '14', '16', '18')
		)
	;`); err != nil {
		log.Println(err)
		log.Println("An error has occured on executing RemoveSystemTempRecords.")
		return err
	}

	return nil
}

func RemoveCO2Records() error {
	db, err := dbConnect()
	if err != nil {
		return err
	}

	defer db.Db.Close()

	if _, err := db.Query(`
	delete
	from purge_sensor_val
	where device_id = 'co2mini'
	and(
		DATE_FORMAT(recorded_at, '%i') in ('00', '02', '04', '06', '08') 
		or DATE_FORMAT(recorded_at, '%i') in ('10', '12', '14', '16', '18')
		or DATE_FORMAT(recorded_at, '%i') in ('20', '22', '24', '26', '28')
		or DATE_FORMAT(recorded_at, '%i') in ('30', '32', '34', '36', '38')
		or DATE_FORMAT(recorded_at, '%i') in ('40', '42', '44', '46', '48')
		or DATE_FORMAT(recorded_at, '%i') in ('50', '52', '54', '56', '58')
    )
	;`); err != nil {
		log.Println(err)
		log.Println("An error has occured on executing RemoveSystemTempRecords.")
		return err
	}

	return nil
}

func RemoveMotionSensorVal(d time.Time) error {
	db, err := dbConnect()
	if err != nil {
		return err
	}

	defer db.Db.Close()

	if _, err := db.Query(`
		delete from t_motion_sensor
		where recorded_at < ?
		`, d); err != nil {
		log.Println(err)
		log.Println("An error has occured on executing RemoveMotionSensorVal delete.")
		return err
	}

	return nil
}
