package conf

import (
	"encoding/json"
	"log"
	"os"
)

const (
	confFilePath = "./config/config.json"
)

// Conf holds configuration of the server
type config struct {
	LogPath string `json:"logpath"`
}

// Conf variable contains the config info
var Conf config

func init() {
	loadConfig()
	validateConfig()
}

func loadConfig() {
	file, err := os.Open(confFilePath)
	if err != nil {
		log.Fatalln("unable to open conf file, error:", err)
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&Conf)
	if err != nil {
		log.Fatalln("unable to decode conf file, error:", err)
	}
}

func validateConfig() {
	_, err := os.Stat(Conf.LogPath)

	if os.IsNotExist(err) {
		err = os.MkdirAll(Conf.LogPath, 0755)
		if err != nil {
			log.Fatalln("Failed to create Log Directory, err:", err)
		}
	}
}
