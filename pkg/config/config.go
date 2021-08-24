package config

import (
	"log"

	"github.com/BurntSushi/toml"
	"gopkg.in/olahol/melody.v1"

	"maid/pkg/structs"
)

var WebSocket = melody.New()

var WebSocketVoice = melody.New()

// Config is mapping config.toml.
var Config = LoadConfig()

// LoadConfig loads config.toml.
func LoadConfig() structs.Config {
	var c structs.Config
	_, err := toml.DecodeFile("./files/conf/config.toml", &c)

	if err != nil {
		log.Println(err)
		return c
	}

	return c
}
