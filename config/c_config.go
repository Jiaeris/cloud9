package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

var (
	JwtTokenKey = "cloudKeyOkOkNoNoNo"
)

type uusConfig struct {
	Db DbConfig
}

type DbConfig struct {
	User     string
	Password string
	Address  string
	Port     int
	DbName   string
}

var UusConfig *uusConfig = nil

//time layout 2006-01-02 15:04:05
func init() {
	if UusConfig != nil {
		return
	}
	file, err := os.Open("config.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	configBuf, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	config := &uusConfig{}
	err = json.Unmarshal(configBuf, config)
	if err != nil {
		panic(err)
	}
	UusConfig = config
}
