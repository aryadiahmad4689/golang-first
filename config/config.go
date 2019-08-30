package config

import (
	"github.com/tkanos/gonfig"
)
// GetConfig func
func GetConfig() Configuration {
	configuration := Configuration{}
	gonfig.GetConf("config/env.json", &configuration)
	return configuration
}
