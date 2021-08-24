package service

import (
	c "maid/pkg/config"

	tplink "github.com/wusamin/tplink-api"
)

func toggleSmartplug(plugName string, status DevicePower) error {
	api, err := tplink.Connect(c.Config.WebSetting.TplinkAddress, c.Config.WebSetting.TplinkPassword)

	if err != nil {
		return err
	}

	hs105, err := api.GetHS105("plug-living-1")
	if err != nil {
		return err
	}

	switch status {
	case PowerOn:
		if err := hs105.TurnOn(); err != nil {
			return err
		}

	case PowerOff:
		if err := hs105.TurnOff(); err != nil {
			return err
		}
	}

	return nil
}
