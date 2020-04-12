package session

import (
	"encoding/json"
	"fmt"
	"github.com/geiqin/supports/helper"
	"github.com/geiqin/supports/xconfig"
	"github.com/go-redis/redis"
	"log"
	"sync"
	"time"
)

//session来自内存 实现
type FromRedis struct {
	Driver *redis.Client
	TTL    int64
	lock   sync.Mutex //用来锁
	//sessions map[string]*list.Element //用来存储在内存
	//list     *list.List               //用来做 gc
}

func LoadRedis(cfg *xconfig.SessionInfo) {
	if cfg == nil || cfg.Provider == nil {
		log.Println("load redis of session config failed")
		return
	}
	log.Println("load redis of session config succeed")

	server := fmt.Sprintf("%s:%d", cfg.Provider.Host, cfg.Provider.Port)

	var client = redis.NewClient(&redis.Options{
		Addr:     server,
		Password: "",                    // no password set
		DB:       cfg.Provider.Database, // use default DB
	})
	pder = &FromRedis{
		Driver: client,
		TTL:    cfg.MaxLifeTime,
		//lock:   sync.Mutex,
	}

	Register("redis", pder)
}

func (this *FromRedis) SessionInit(sid string) (Session, error) {
	this.lock.Lock()
	defer this.lock.Unlock()
	v := make(map[string]interface{}, 0)
	newSess := &SessionStore{sid: sid, LastAccessedTime: time.Now(), value: v}

	var h = this.Driver
	var bytes []byte
	//var err error
	var content string

	if h.Exists(sid).Val() == 1 {
		content = h.Get(sid).Val()
	} else {
		content = "{}"
	}
	json.Unmarshal([]byte(content), &v)

	bytes, _ = json.Marshal(v)
	h.Set(sid, string(bytes), time.Duration(this.TTL)*time.Second)
	return newSess, nil
}

func (this *FromRedis) SessionRead(sid string) (Session, error) {
	var h = this.Driver

	if h.Exists(sid).Val() == 1 {
		var content string
		v := make(map[string]interface{}, 0)
		content = h.Get(sid).Val()
		err := json.Unmarshal([]byte(content), &v)
		sess := &SessionStore{sid: sid, LastAccessedTime: time.Now(), value: v}
		return sess, err

	} else {
		sess, err := this.SessionInit(sid)
		return sess, err
	}

	return nil, nil
}

func (this *FromRedis) SessionDestroy(sid string) error {
	var h = this.Driver
	h.Del(sid).Val()
	return nil

}

func (this *FromRedis) SessionGC(maxLifeTime int64) {

}

func (this *FromRedis) SessionUpdate(sid string) error {
	var h = this.Driver
	h.Expire(sid, time.Duration(this.TTL)*time.Second)
	return nil
}

func (this *FromRedis) SessionSave(sid string, value map[string]interface{}) error {
	var h = this.Driver
	val := helper.JsonEncode(value)
	h.Set(sid, val, time.Duration(this.TTL)*time.Second)
	return nil
}
