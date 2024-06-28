package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

var (
	MysqlDBConnection string
)

func Initialize() {
	filePath := "files/config/config.json"

	jsonFile, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Failed to get config file ", err)
	}
	config := apiAmarthaDB{}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal([]byte(byteValue), &config)

	MysqlDBConnection = config.ConfigDB.Data.MysqlDBConnection
}
