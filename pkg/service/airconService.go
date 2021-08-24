package service

import (
	"context"
	"log"
	c "maid/pkg/config"
	"maid/pkg/repository"
	"maid/pkg/structs"
	"maid/pkg/util"
	"os"
	"strconv"

	"github.com/tenntenn/natureremo"
)

var airconSetting string = c.Config.WebSetting.JSONFilePath + "/airconSetting.json"

func UpdateAirconSetting(key, value string) {
	log.Println("処理開始")
	_, err := os.Stat(airconSetting)

	// if there is not json file, create json file with blank map.
	if os.IsNotExist(err) {
		log.Println("ファイルなし")
		util.ExportJSON2File(airconSetting, map[string]string{})
	}

	setting, err := util.ReadJSONFile(airconSetting)

	if err != nil {
		log.Println(err)
		return
	}

	setting[key] = value

	util.ExportJSON2File(airconSetting, setting)
}

// Airconditionning controls aircon appliance.
func Airconditionning() {
	// log.Println("-----------------------------------------")
	// log.Println("Airconditionning start...")
	// defer log.Println("Airconditionning has finished...")
	// defer log.Println("-----------------------------------------")

	cli := natureremo.NewClient(c.Config.WebSetting.NatureRemoToken)
	ctx := context.Background()

	appliances, err := cli.ApplianceService.GetAll(ctx)

	// 機器情報取得でエラーがあった場合は終了
	if err != nil {
		log.Println(err)
		return
	}

	for _, appliance := range appliances {
		// 機器毎に制御処理を実行する
		switch appliance.Type {
		case natureremo.ApplianceTypeAirCon:
			aircon(ctx, appliance, cli)
		}
	}
}

func aircon(ctx context.Context, airconAppliance *natureremo.Appliance, cli *natureremo.Client) {
	// log.Println("controling aircon start...")
	// defer log.Println("controling aircon has finished...")

	if airconAppliance == nil {
		log.Println("Type aircon is not registered.")
		return
	}

	// 電源が点いているときのみ制御を行う
	if airconAppliance.AirConSettings.Button == "power-off" {
		return
	}

	var airconSetting *natureremo.AirConSettings

	// 室温の情報を取得する
	temprature := repository.FindSensorValAverage("natureremo", "te")

	settingedTemp := airconAppliance.AirConSettings.Temperature

	// エアコンの実行状態毎に行う処理を振り分ける
	switch airconAppliance.AirConSettings.OperationMode {
	case natureremo.OperationModeCool:
		airconSetting = controlCooler(airconAppliance, &temprature)

	case natureremo.OperationModeWarm:
	case natureremo.OperationModeDry:
	}

	// log.Print("aircon setting : ")
	// fmt.Println(airconSetting)
	// エアコン設定がnullの場合は処理を終了する
	if airconSetting == nil {
		return
	}

	log.Println("-----------------------------------------")
	log.Println("Aircon setting is changed.")
	log.Printf("%-11s : %s", "Mode", airconSetting.OperationMode)
	log.Printf("%-11s : %s", "Now temp", temprature.Val)
	log.Printf("%-11s : %s -> %s", "Temparature", settingedTemp, airconSetting.Temperature)
	log.Printf("%-11s : %s", "Direction", airconSetting.AirDirection)
	log.Printf("%-11s : %s", "Air volume", airconSetting.AirVolume)
	log.Printf("%-11s : %s", "Button", airconSetting.Button)
	log.Println("-----------------------------------------")

	if err := cli.ApplianceService.UpdateAirConSettings(ctx, airconAppliance, airconSetting); err != nil {
		log.Println(err)
		return
	}
}

func controlCooler(appliance *natureremo.Appliance, temparature *structs.TNatureRemoSensor) *natureremo.AirConSettings {

	setting, err := util.ReadJSONFile(airconSetting)

	if err != nil {
		log.Println(err)
		return nil
	}

	comfortTemparature, _ := strconv.ParseFloat(setting["comfortTemparature"], 64)
	nowTemp, _ := strconv.ParseFloat(temparature.Val, 64)
	tempSettinged, _ := strconv.ParseFloat(appliance.AirConSettings.Temperature, 64)

	// 快適な温度の範囲ならば、温度の制御を行わない
	if comfortTemparature-0.5 <= nowTemp && nowTemp <= comfortTemparature+0.3 {
		return nil
	}

	// 室温が快適な温度よりも高い場合
	if comfortTemparature < nowTemp {
		switch {
		// エアコンの設定温度が快適な温度になっていた場合、温度を1度下げる
		case comfortTemparature == tempSettinged:
			appliance.AirConSettings.Temperature = strconv.Itoa(int(tempSettinged - 1))
			appliance.AirConSettings.Button = ""
			// log.Println("set temparature :" + appliance.AirConSettings.Temperature)
			return appliance.AirConSettings

		// 室温が快適な温度よりも高く、エアコンの設定温度が快適な温度よりも高い場合、エアコンの設定温度を快適な温度に設定する
		case comfortTemparature < tempSettinged:
			appliance.AirConSettings.Temperature = strconv.Itoa(int(comfortTemparature))
			appliance.AirConSettings.Button = ""
			// log.Println("set temparature :" + appliance.AirConSettings.Temperature)
			return appliance.AirConSettings

		// 室温が快適な温度よりも高く、エアコンの設定温度が快適な温度よりも低い場合、エアコンの設定温度を操作しない
		case tempSettinged < comfortTemparature:
			return nil
		}
	}

	// 快適な温度よりも室温が低い場合
	if nowTemp < comfortTemparature {
		switch {

		// 室温が快適な温度よりも低く、エアコンの設定温度が快適な温度の場合、調整を行わない
		case tempSettinged == comfortTemparature:
			return nil

		// クーラーを点けてる状態で、快適な温度よりも室温が低い場合、エアコンの温度を快適な温度に設定する
		case tempSettinged < comfortTemparature:
			appliance.AirConSettings.Temperature = strconv.Itoa(int(comfortTemparature))
			appliance.AirConSettings.Button = ""
			return appliance.AirConSettings
		}
	}

	return nil
}
