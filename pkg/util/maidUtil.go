package util

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/tenntenn/natureremo"

	c "maid/pkg/config"
	"maid/pkg/structs"
)

// TransFloat convert value of float64(Exponential notation) to string(xx.xx).
func TransFloat(value float64) string {
	return strconv.FormatFloat(value, 'f', 2, 64)
}

// TplinkLightbulb send command for Tplight or Tplight.exe.
func TplinkLightbulb(cmds ...string) ([]byte, error) {
	return exec.Command(c.Config.WebSetting.TplightPath, cmds...).Output()
}

type LightState struct {
	OnOff      int    `json:"on_off"`
	Mode       string `json:"mode"`
	Hue        int    `json:"hue"`
	Saturation int    `json:"saturation"`
	ColorTemp  int    `json:"color_temp"`
	Brightness int    `json:"brightness"`
}

func TplinkLightbulbExecute(t *structs.TplinkLightbulbCmd) map[string]interface{} {
	go CmdKill("tplight", 5)
	switch t.Command {
	case structs.Details:
		cmds := []string{}
		cmds = append(cmds, string(t.Command))
		cmds = append(cmds, t.IP)

		if b, err := exec.Command(c.Config.WebSetting.TplightPath, cmds...).Output(); err != nil {
			return nil
		} else {
			// fmt.Println(string(b))
			var r map[string]interface{}

			if err := json.Unmarshal(b, &r); err != nil {
				log.Println(err)
				return r
			}
			return r["light_state"].(map[string]interface{})
		}
	case structs.Scan:
		cmds := []string{}
		cmds = append(cmds, string(t.Command))
		if t.Timeout != 0 {
			cmds = append(cmds, "-t", strconv.Itoa(t.Transition))
		}

		if b, err := exec.Command(c.Config.WebSetting.TplightPath, cmds...).Output(); err != nil {
			return nil
		} else {
			fmt.Println(string(b))
		}
	case structs.Temp:
		cmds := []string{}
		cmds = append(cmds, string(t.Command))
		cmds = append(cmds, t.IP)
		cmds = append(cmds, strconv.Itoa(t.Kelvin))
		if t.Transition != 0 {
			cmds = append(cmds, "-t", strconv.Itoa(t.Transition))
		}

		exec.Command(c.Config.WebSetting.TplightPath, cmds...).Output()
	}

	return nil
}

// CallAmah gets struct that has context, natureremo client.
func CallAmah() *structs.Amah {
	return &structs.Amah{Context: context.Background(), NatureremoClient: natureremo.NewClient(c.Config.WebSetting.NatureRemoToken)}
}

// ExportJSON2File export interface to JSON file specified by fileName.
func ExportJSON2File(fileName string, v interface{}) {
	ebytes, err := json.MarshalIndent(v, "", "    ")

	if err != nil {
		log.Println(err)
		return
	}

	if err := ioutil.WriteFile(fileName, ebytes, os.ModePerm); err != nil {
		log.Println(err)
	}
}

// ReadJSONFile returns map json unmarshalled.
func ReadJSONFile(fileName string) (map[string]string, error) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return map[string]string{}, err
	}
	var r map[string]string

	if err := json.Unmarshal(bytes, &r); err != nil {
		log.Println(err)
		return map[string]string{}, err
	}

	return r, nil
}

func CmdKill(processName string, timeout int64) {
	if b, err := exec.Command(c.Config.WebSetting.ShPath+"/ProcessKill.sh", processName, strconv.FormatInt(timeout, 10)).Output(); err == nil {
		log.Println(strings.ReplaceAll(string(b), "\n", ""))
	} else {
		log.Println(err)
	}
}

func CmdExecWithRetry(cmd *exec.Cmd, tryTime int64) {
	if err := cmd.Start(); err != nil {

	}

	cmd.Process.Kill()
}
