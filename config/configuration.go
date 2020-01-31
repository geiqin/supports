package config

import (
	"github.com/micro/go-micro/config"
	"github.com/mitchellh/mapstructure"
	"log"
)

var conf map[string]interface{}

func init()  {
	err := config.LoadFile("./config/config.json")
	if err != nil {
		log.Fatal("could not load config file: %s", err.Error())
	}
	conf = config.Map()
}

func GetConfig(path string) map[string]interface{} {
	if path !="" {
		return conf[path].(map[string]interface{})
	}
	return conf
}


func ConvertStruct(key string,structModel interface{}) interface{} {
	cfg :=GetConfig(key)
	if cfg ==nil {
		log.Fatal("convert structModel is failed key:",key)
	}
	mapstructure.Decode(cfg,structModel)
	return structModel
}
