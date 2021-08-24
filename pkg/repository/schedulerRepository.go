package repository

import "maid/pkg/structs"

// InsertScheduledURL insert dto to TMaidTimerScheduled.
func InsertScheduledURL(dto *structs.TMaidTimerScheduled) error {
	gorm, err := connectGorm()
	if err != nil {
		return err
	}

	defer gorm.Close()

	gorm.Table("t_maid_timer_scheduled").Create(dto)

	return nil
}

// DeleteScheduledURL delete a record specified.
func DeleteScheduledURL(pk *structs.TMaidTimerScheduled) error {
	dbmap, e := dbConnect()
	if e != nil {
		return e
	}

	// dbmap.AddTableWithName(structs.MHomeAppliances{}, "m_home_appliances")
	defer dbmap.Db.Close()

	// var dto structs.MHomeAppliances

	_, err := dbmap.Exec("delete from t_maid_timer_scheduled where id = ?", pk.ID)

	if err != nil {
		println(err.Error())
	}

	return nil
}
