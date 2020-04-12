package xconfig

type Configuration struct {
	AppInfo        *AppInfo                   `json:"app_info"`
	SessionInfo    *SessionInfo               `json:"session_info"`
	DatabaseList   map[string]*DatabaseInfo   `json:"database_list"`
	RedisList      map[string]*RedisInfo      `json:"redis_list"`
	FileSystemList map[string]*FileSystemInfo `json:"file_system_list"`
	TokenList      map[string]*TokenInfo      `json:"token_list"`
	SmsList        map[string]*SmsInfo        `json:"sms_list"`
	MailList       map[string]*MailInfo       `json:"mail_list"`
	WxPayInfo      *WxPayInfo                 `json:"wx_pay_info"`
	AliPayInfo     *AliPayInfo                `json:"ali_pay_info"`
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
	PrivateKey []byte `json:"private_key"`
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

type CacheInfo struct {
}

type SessionInfo struct {
	Driver      string     `json:"driver"`
	CookieName  string     `json:"cookie_name"`
	MaxLifeTime int64      `json:"max_life_time"`
	Provider    *RedisInfo `json:"provider"`
}

type WxPayInfo struct {
	NotifyUrl    string `json:"notify_url"`
	QinAppId     string `json:"qin_app_id"`
	QinAppSecret string `json:"qin_app_secret"`
	QinMchId     string `json:"qin_mch_id"`
	QinMd5Key    string `json:"qin_md5_key"`
}

type AliPayInfo struct {
	AppId     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
	MchId     string `json:"mch_id"`
	Md5key    string `json:"md_5_key"`
}
