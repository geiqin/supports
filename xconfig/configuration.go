package xconfig

import "os"

var conf *Configuration

//注册配置
func Register(names ...string) {
	defNames := []string{"database"}
	if names != nil {
		defNames = append(defNames, names...)
	}

	address := os.Getenv("CFG_SERVER_ADDRESS")
	if address == "" {
		address = "127.0.0.1:8600"
	}

	mgr := NewConfigManager(address, defNames)
	conf = mgr.Load()
}

//获取全部配置
func GetConfig() *Configuration {
	return conf
}

//获取数据库配置
func GetDatabaseCfg(name string) *DatabaseInfo {
	cfg := conf.DatabaseList[name]
	return cfg
}

//获取店铺数据库配置
func GetStoreDatabaseCfg(name string, storeFlag string) *DatabaseInfo {
	cfg := *conf.DatabaseList[name]
	if &cfg != nil {
		cfg.Database = storeFlag
	}
	return &cfg
}

//获取缓存配置
func GetCacheCfg() *RedisInfo {
	cfg := conf.RedisList["cache"]
	return cfg
}

//获取会话配置
func GetSessionCfg() *RedisInfo {
	cfg := conf.RedisList["session"]
	return cfg
}

//获取授权配置
func GetAuthCfg(name string) *AuthInfo {
	cfg := conf.AuthList[name]
	return cfg
}
