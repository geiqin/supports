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

type Option struct {
	Flag string
	Name string
}

func Run(flag string,option ...Option)  {
	log.Println("app flag ["+flag+"] is running")
	if option !=nil{
		appOption  =&option[0]
		appOption.Flag =flag
	} else{
		appOption  =&Option{
			Flag:flag,
		}
	}
	config.Load()
	session.Load()
	cache.Load()
	database.Load()
	token.Load()
}

func Flag() string {
	return appOption.Flag
}

func GetOption() *Option {
	return appOption
}
