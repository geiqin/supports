package cache

import (
	"github.com/geiqin/supports/xconfig"
	"log"
	"sync"
	"time"
)

var cache Cache
var once sync.Once

type Cache interface {
	HashSet(key string, value interface{}, duration time.Duration) error
	HashGet(key string) interface{}
	Set(key string, value string, duration time.Duration) error
	Get(key string) string
	Has(key string) bool
	Increment(key string, step ...int64) int64
	Decrement(key string, step ...int64) int64
	Remember(key string, duration time.Duration, fn func(args ...interface{}) string)
	Delete(key string)
}

func Load() {
	cfg := xconfig.GetCacheCfg()
	if cfg == nil {
		log.Println("load cache config failed")
		return
	}
	log.Println("load cache config succeed")
	LoadRedis(cfg)
}

func GetCache() Cache {
	return redisStore
}

func Register(provider Cache) {
	if provider == nil {
		panic("session: Register provide is nil")
	}
	cache = provider
}

/*
//注册 由实现Provider接口的结构体调用
func Register(name string, provider Cache) {
	if provider == nil {
		panic("session: Register provide is nil")
	}
	cache = provider
}

*/
