package app

import (
	"github.com/geiqin/supports/auth"
	"github.com/geiqin/supports/cache"
	"github.com/geiqin/supports/config"
	"github.com/geiqin/supports/database"
	"github.com/geiqin/supports/session"
	"log"
)

var appOption *Option

//var once sync.Once

type ConfigMode string

const (
	LocalMode ConfigMode = "local"
	CloudMode ConfigMode = "cloud"
)

type Option struct {
	Flag       string
	Private    bool
	ConfigMode ConfigMode
	ConfigPath string
}

func Run(flag string, private bool, option ...Option) {
	log.Println("app flag [" + flag + "] is running")

	opt := &Option{}
	if option != nil {
		opt = &option[0]
	}
	opt.Flag = flag
	opt.Private = private
	if opt.ConfigMode == "" || opt.ConfigPath == "" {
		opt.ConfigPath = "./configs"
	}
	appOption = opt
	config.Load(opt.ConfigPath)
	session.Load()
	cache.Load()
	database.Load(opt.Flag)
	auth.Load()
}

func Flag() string {
	return appOption.Flag
}

func Private() bool {
	return appOption.Private
}

func GetOption() *Option {
	return appOption
}
