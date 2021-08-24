package structs

type TplinkLightbulbCmd struct {
	Command    Cmd
	IP         string
	Transition int
	Kelvin     int
	Brightness int
	Timeout    int
}

type Cmd string

const (
	Scan    Cmd = "cmd"
	On      Cmd = "on"
	Off     Cmd = "off"
	Temp    Cmd = "temp"
	Hex     Cmd = "hex"
	Hsb     Cmd = "hsb"
	Cloud   Cmd = "cloud"
	Raw     Cmd = "raw"
	Details Cmd = "details"
)
