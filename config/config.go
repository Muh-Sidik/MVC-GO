package config

import (
	"github.com/tkanos/gonfig"
)

type Configurations struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_PORT     string
	DB_HOST     string
	DB_NAME     string
}

func GetConfig() Configurations {
	config := Configurations{}
	err := gonfig.GetConf("config/config.json", &config)

	if err != nil {
		panic(err)
	}

	return config
}
