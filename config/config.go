package config

import (
	"github.com/tkanos/gonfig"
)

// func GetConfig
func GetConfig() Configuration {
	configuration := Configuration{}
	gonfig.GetConf("config/env.json", &configuration)
	return configuration
}
