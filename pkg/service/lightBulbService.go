package service

import (
	"fmt"
	"log"
	c "maid/pkg/config"
	"maid/pkg/structs"
	"maid/pkg/util"
	u "maid/pkg/util"
	"strconv"
	"strings"
)

// TurnOffLightBulb turns lightbulb off.
func TurnOffLightBulb(ip string, transition int) error {
	ret, err := u.TplinkLightbulb("off", ip, "-t", strconv.Itoa(transition))

	if err != nil {
		log.Println(err)
		return err
	}

	log.Println(string(ret))
	return nil
}

// ChangeLightBulbStatus sets status to lightbulb.
func ChangeLightBulbStatus(bulbName string, status map[string]string) (int, map[string]string) {
	bulbMap, e := u.ReadJSONFile(c.Config.WebSetting.JSONFilePath + "/tplink.json")

	if e != nil {
		log.Println(e)
		return 400, map[string]string{"message": "An error has occured on getting 'ChangeLightBulbStatus' gets bulbMap"}
	}

	for k, v := range status {
		switch k {
		case "il":
			if _, err := u.TplinkLightbulb("on", bulbMap[bulbName], "-b", v, "-t", status["transition"]); err != nil {
				log.Println(err)
				mes := "Get error on executing command 'il'"
				log.Println(mes)
				return 400, map[string]string{"message": mes}
			}
		case "temp":
			if _, err := u.TplinkLightbulb("temp", bulbMap[bulbName], v, "-t", status["transition"]); err != nil {
				log.Println(err)
				mes := "Get error on executing command 'temp'"
				log.Println(mes)
				return 400, map[string]string{"message": mes}
			}
		case "transition":
			break
		}

	}

	return 200, map[string]string{"message": "ok"}
}

func ChangeLightbulbIlumination(ip string, illuminance int, transition int) error {

	_, err := u.TplinkLightbulb("on", ip, "-b", strconv.Itoa(illuminance), "-t", strconv.Itoa(transition))

	if err != nil {
		log.Println(err)
		return err
	}

	log.Println("-----------------------------------------")
	log.Println("LightBulb status was changed.")
	log.Printf("%-11s : %s", "command", "on")
	log.Printf("%-11s : %s", "IP", ip)
	log.Printf("%-11s : %s", "Brightness", strconv.Itoa(illuminance))
	log.Printf("%-11s : %s", "Transition", strconv.Itoa(transition))
	// log.Println(string(ret))
	log.Println("-----------------------------------------")

	return nil
}

// RecordIPTPLink exports json file that lightbulb's IP address has been recoreded.
func RecordIPTPLink() error {
	b, err := util.TplinkLightbulb("scan", "-t", "2")

	if err != nil {
		return err
	}

	m := map[string]string{}

	for i, v := range strings.Split(string(b), "\n") {
		splitted := strings.Split(v, " - ")
		if 1 < len(splitted) {
			log.Printf("%d name:%s, addr:%s", i, splitted[1], splitted[0])
			m[splitted[1]] = splitted[0]
		}
	}

	util.ExportJSON2File(c.Config.WebSetting.JSONFilePath+"/tplink.json", m)

	return nil
}

func GetActiveLightbulbIP() []map[string]string {
	m, e := u.ReadJSONFile(c.Config.WebSetting.JSONFilePath + "/tplink.json")

	if e != nil {
		log.Println(e)

	}

	bulbs := []string{m["bulb-ceiling-1"], m["bulb-ceiling-2"], m["bulb-ceiling-3"], m["bulb-ceiling-4"]}

	c := structs.TplinkLightbulbCmd{}
	c.Command = structs.Details

	ret := []map[string]string{}

	for _, IP := range bulbs {
		c.IP = IP
		m := u.TplinkLightbulbExecute(&c)

		if m == nil {
			continue
		}

		record := map[string]string{}

		if m["on_off"].(float64) == 1 {
			record["colorTemp"] = fmt.Sprint(m["color_temp"])
			record["ip"] = IP
			ret = append(ret, record)
		}
	}

	return ret
}
