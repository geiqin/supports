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

func (b *ConfigManger) makeAuth(read reader.Value) error {
	var list map[string]*AuthInfo
	read.Scan(&list)
	b.conf.AuthList = list
	return nil
}

func (b *ConfigManger) makeRedis(read reader.Value) error {
	var list map[string]*RedisInfo
	read.Scan(&list)
	b.conf.RedisList = list
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
		case "database":
			b.makeDatabase(config.Get("connections"))
			b.makeRedis(config.Get("redis"))
			break
		case "filesystem":
			b.makFilesystem(config.Get("clouds"))
			break
		case "auth":
			b.makFilesystem(config.Get("providers"))
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
