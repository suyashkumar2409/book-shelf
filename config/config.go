package config

import "sync"

var(
	configFile string
	c * Config
	configLock sync.RWMutex
)

type Config struct{
	DBName string `json:DBName`
}

func parseConfigFile(){

}

func GetDBName() string{
	configLock.RLock()
	defer configLock.RUnlock()

	return c.DBName
}

