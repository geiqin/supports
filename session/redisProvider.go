package session

import (
	"encoding/json"
	"fmt"
	"github.com/geiqin/supports/config"
	"github.com/geiqin/supports/helper"
	"github.com/go-redis/redis"
	"log"
	"sync"
	"time"
)

//session来自内存 实现
type FromRedis struct {
	Driver *redis.Client
	TTL int64
	lock     sync.Mutex               //用来锁
	//sessions map[string]*list.Element //用来存储在内存
	//list     *list.List               //用来做 gc
}

type RedisType struct {
	Host string
	Port int
	Database int
}

func LoadRedis(cnf *SessionConfig) {

	redisCfg :=config.GetConfig("database","redis","session")

	if redisCfg ==nil{
		log.Println("load redis of session config failed")
		return
	}
	rd :=&RedisType{}
	helper.MapToStruct(redisCfg,rd)
	server :=fmt.Sprintf("%s:%d",rd.Host,rd.Port)

	var client = redis.NewClient(&redis.Options{
		Addr: server,
		Password: "", // no password set
		DB:       rd.Database,                              // use default DB
	})
	pder =&FromRedis{
		Driver: client,
		TTL:    cnf.MaxLifeTime,
		//lock:   sync.Mutex,
	}

	Register("redis", pder)
}



func (this *FromRedis) SessionInit(sid string) (Session, error) {
	this.lock.Lock()
	defer this.lock.Unlock()
	v := make(map[string]interface{}, 0)
	newsess := &SessionStore{sid: sid, LastAccessedTime: time.Now(), value: v}
	//element := this.list.PushBack(newsess)
	//this.sessions[sid] = element


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
	return newsess, nil
}

func (this *FromRedis) SessionRead(sid string) (Session, error) {
	var h = this.Driver

	if h.Exists(sid).Val() == 1 {
		var content string
		v := make(map[string]interface{}, 0)
		content = h.Get(sid).Val()
		err :=json.Unmarshal([]byte(content), &v)
		sess := &SessionStore{sid: sid, LastAccessedTime: time.Now(), value: v}
		return sess,err

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
	//var h = this.Driver
	//h.Expire(sid,time.Duration(this.TTL)*time.Second)
	return nil
}

func (this *FromRedis) SessionSave(sid string, value map[string]interface{}) error {
	var h = this.Driver
	val :=helper.JsonEncode(value)
	h.Set(sid, val, time.Duration(this.TTL)*time.Second)
	return nil
}