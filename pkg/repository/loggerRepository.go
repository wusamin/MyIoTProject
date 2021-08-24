package repository

import (
	"fmt"
	"log"
	"maid/pkg/structs"
	"time"
)

// InsertLog is method to record for log API has been called.
func InsertLog(apiName, res string) {
	dbmap, e := dbConnect()
	if e != nil {
		log.Println(e)
		return
	}

	dbmap.AddTableWithName(structs.TApiCalledLog{}, "T_API_CALLED_LOG")
	defer dbmap.Db.Close()

	dto := &structs.TApiCalledLog{APIName: apiName, OperationResult: res, InsDate: time.Now()}

	err := dbmap.Insert(dto)

	if err != nil {
		println(err.Error())
	}
}

// FindLatestLog gets a latest record specified apiName.
func FindLatestLog(apiName string) *structs.TApiCalledLog {
	dbmap, e := dbConnect()
	if e != nil {
		log.Println(e)
		return &structs.TApiCalledLog{}
	}

	defer dbmap.Db.Close()

	var dto structs.TApiCalledLog

	// 過去6レコードの平均値を取得する
	err := dbmap.SelectOne(&dto,
		`
		select
			t1.API_NAME
			, t1.operation_result
			, t1.INS_DATE 
		from
			T_API_CALLED_LOG t1 
			inner join ( 
				select
					max(api.ID) id 
				from
					T_API_CALLED_LOG api 
				where
					api.API_NAME = ?
				) t2 
			on t1.ID = t2.id
		`, apiName)

	if err != nil {
		fmt.Println(err)
	}

	return &dto
}
