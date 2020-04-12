package xconfig

import (
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/reader"
	grpcConfig "github.com/micro/go-plugins/config/source/grpc/v2"
	"log"
)

type ConfigManger struct {
	address string
	conf    *Configuration
	names   []string
}

func (b *ConfigManger) makeApp(read reader.Value) error {
	var info *AppInfo
	read.Scan(&info)
	b.conf.AppInfo = info
	return nil
}

func (b *ConfigManger) makeAppSession(read reader.Value) error {
	var info *SessionInfo
	read.Scan(&info)
	b.conf.SessionInfo = info
	return nil
}

func (b *ConfigManger) makeToken(read reader.Value) error {
	var list map[string]*TokenInfo
	read.Scan(&list)
	b.conf.TokenList = list
	return nil
}

func (b *ConfigManger) makeRedis(read reader.Value) error {
	var list map[string]*RedisInfo
	read.Scan(&list)
	b.conf.RedisList = list
	return nil
}

func (b *ConfigManger) makeWxPay(read reader.Value) error {
	var info *WxPayInfo
	read.Scan(&info)
	b.conf.WxPayInfo = info
	return nil
}

func (b *ConfigManger) makeAliPay(read reader.Value) error {
	var info *AliPayInfo
	read.Scan(&info)
	b.conf.AliPayInfo = info
	return nil
}

func (b *ConfigManger) makFilesystem(read reader.Value) error {
	var list map[string]*FileSystemInfo
	read.Scan(&list)
	b.conf.FileSystemList = list
	return nil
}

func (b *ConfigManger) makeDatabase(read reader.Value) error {
	var list map[string]*DatabaseInfo
	read.Scan(&list)
	b.conf.DatabaseList = list
	return nil
}

func (b *ConfigManger) Load() *Configuration {
	for _, app := range b.names {
		if err := config.Load(grpcConfig.NewSource(
			grpcConfig.WithAddress(b.address),
			grpcConfig.WithPath("/"+app),
		)); err != nil {
			log.Fatalf("[ConfigManger] load files error，%s", err)
			return nil
		}
		switch app {
		case "app":
			//b.makeDatabase(config.Get("app"))
			b.makeAppSession(config.Get("session"))
			break
		case "database":
			b.makeDatabase(config.Get("connections"))
			b.makeRedis(config.Get("redis"))
			break
		case "filesystem":
			b.makFilesystem(config.Get("clouds"))
			break
		case "token":
			b.makeToken(config.Get("tokens"))
			break
		case "payment":
			b.makeWxPay(config.Get("providers", "weixin"))
			b.makeAliPay(config.Get("providers", "alipay"))
			break
		}
	}
	return b.conf
}

func NewConfigManager(address string, names []string) *ConfigManger {
	obj := &ConfigManger{
		address: address,
		conf:    &Configuration{},
		names:   names,
	}
	return obj
}
