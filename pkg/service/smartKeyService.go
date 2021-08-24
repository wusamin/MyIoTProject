package service

import (
	"encoding/json"
	"log"
	c "maid/pkg/config"

	"github.com/go-resty/resty/v2"
)

// GetKeyStatus gets status of Sesame.
func GetKeyStatus() (map[string]interface{}, error) {
	res, err := resty.New().R().
		SetHeader("Accept", "application/json").
		SetHeader("Authorization", c.Config.WebSetting.SesameToken).
		Get("https://api.candyhouse.co/public/sesame/" + c.Config.WebSetting.SesameID)

	if err != nil {
		return map[string]interface{}{}, err
	}

	var dat map[string]interface{}
	json.Unmarshal(res.Body(), &dat)

	return dat, nil
}

// Key controls Sesame.
func Key() (int, map[string]string) {
	res, err := resty.New().R().
		SetHeader("Accept", "application/json").
		SetHeader("Authorization", c.Config.WebSetting.SesameToken).
		Get("https://api.candyhouse.co/public/sesame/" + c.Config.WebSetting.SesameID)

	if err != nil {
		return 400, map[string]string{"message": "An error has occured on getting Sesame status."}
	}

	var dat map[string]interface{}
	json.Unmarshal(res.Body(), &dat)

	if locked, ok := dat["locked"].(bool); ok {
		if locked {
			if _, err := SmartLock("unlock"); err != nil {
				log.Println(err)
				return 400, map[string]string{"message": "An error has occured on unlocking Sesame."}
			}
		} else {
			if _, err := SmartLock("lock"); err != nil {
				log.Println(err)
				return 400, map[string]string{"message": "An error has occured on locking Sesame."}
			}
		}
	}

	return 204, map[string]string{}
}

// SmartLock locks or unlock SmartLock.
func SmartLock(status string) (map[string]interface{}, error) {
	res, err := resty.New().R().SetBody(map[string]interface{}{"command": status}).
		SetHeader("Accept", "application/json").
		SetHeader("Authorization", c.Config.WebSetting.SesameToken).
		Post("https://api.candyhouse.co/public/sesame/" + c.Config.WebSetting.SesameID)

	if err != nil {
		log.Println(err)
		return map[string]interface{}{}, err
	}

	var dat map[string]interface{}
	if err := json.Unmarshal(res.Body(), &dat); err != nil {
		log.Println(err)
		return map[string]interface{}{}, err
	}

	log.Println("-----------------------------------------")
	log.Println("Key status was changed.")
	log.Printf("%-8s : %s", "command", status)
	log.Printf("%-8s : %s", "Task-ID", dat["task_id"])
	log.Println("-----------------------------------------")

	return dat, nil
}
