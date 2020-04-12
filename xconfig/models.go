package xconfig

type Configuration struct {
	AppInfo        AppInfo
	DatabaseList   map[string]*DatabaseInfo
	RedisList      map[string]*RedisInfo
	FileSystemList map[string]*FileSystemInfo
	TokenList      map[string]*TokenInfo
	ProviderList   map[string]*ProviderInfo
	SmsList        map[string]*SmsInfo
	MailList       map[string]*MailInfo
}

type DatabaseInfo struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
	Prefix   string `json:"prefix"`
}

type RedisInfo struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database int    `json:"database"`
}

type ProviderInfo struct {
	AppId     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
	MchId     string `json:"mch_id"`
	Md5key    string `json:"md_5_key"`
}

type FileSystemInfo struct {
	Driver    string `json:"driver"`
	Domain    string `json:"domain"`
	AccessKey string `json:"access_key"`
	SecretKey string `json:"secret_key"`
	Bucket    string `json:"bucket"`
	Transport string `json:"transport"`
}

type TokenInfo struct {
	Issuer     string `json:"issuer"`
	Audience   string `json:"audience"`
	PrivateKey string `json:"private_key"`
	ExpireTime int    `json:"expire_time"`
}

type SmsInfo struct {
	Address string `json:"address"`
	Port    int    `json:"port"`
}

type MailInfo struct {
}

type AppInfo struct {
}
