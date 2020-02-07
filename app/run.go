package app

import (
	"github.com/geiqin/supports/cache"
	"github.com/geiqin/supports/config"
	"github.com/geiqin/supports/database"
	"github.com/geiqin/supports/session"
	"github.com/geiqin/supports/token"
	"log"
)

var appOption *Option
//var once sync.Once

type ConfigMode string
const (
	LocalMode ConfigMode ="local"
	CloudMode ConfigMode ="cloud"
)

type Option struct {
	Flag string
	//Name string
	ConfigMode ConfigMode
	ConfigPath string
}

func Run(flag string,option ...Option)  {
	log.Println("app flag ["+flag+"] is running")

	opt:=&Option{}
	if option!=nil{
		opt =&option[0]
	}
	opt.Flag=flag
	if  opt.ConfigMode=="" || opt.ConfigPath==""{
		opt.ConfigPath="./configs"
	}

	config.Load(opt.ConfigPath)
	session.Load()
	cache.Load()
	database.Load(opt.Flag)
	token.Load()
}

func Flag() string {
	return appOption.Flag
}

func GetOption() *Option {
	return appOption
}
