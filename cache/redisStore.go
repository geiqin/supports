package cache

import (
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"time"
)
var redisCli *redis.Client
var redisStore *RedisStore

type RedisStore struct {}

func RedisInit(cfg *CacheConfig)   {
	serverAddr :=fmt.Sprintf("%s:%d",cfg.Host,cfg.Port)

	opts :=&redis.Options{
		Addr: serverAddr, // Redis地址
		Password: cfg.Password,  // Redis账号
		DB:    cfg.Database,  // Redis库
		PoolSize: 10,  // Redis连接池大小
		MaxRetries: 3, // 最大重试次数
		IdleTimeout: 10*time.Second, // 空闲链接超时时间
	}

	redisCli = redis.NewClient(opts)
	pong, err := redisCli.Ping().Result()

	if err == redis.Nil {
		log.Fatal("Redis异常")
	} else if err != nil {
		log.Fatal("失败:", err)
	} else {
		log.Println("redis init ",pong)
	}
}

func (h *RedisStore) Set(key string, value string,duration time.Duration) error  {
	if duration > 0{
		s :=redisCli.Set(key,value,duration)
		if s.Err() !=nil{
			fmt.Println("cache redis set failed :",s.Err())
			return s.Err()
		}
	} else{
		g :=redisCli.GetSet(key, value)
		if g.Err() != nil {
			fmt.Println("cache redis getset failed :",g.Err())
			return g.Err()
		}
	}

	return nil
}

func (h *RedisStore) Get(key string) string {
	c :=  redisCli.Get(key)
	if c.Err() != nil {
		fmt.Println("redis get failed:", c.Err())
		return ""
	}
	return c.Val()
}

func (h *RedisStore) Has(key string) bool  {
	rs :=redisCli.Exists(key)
	if rs.Err() !=nil{
		return false
	}
	if rs.Val() == 1{
		return true
	}
	return false
}


func (h *RedisStore) Increment(key string, step ...int64) int64{
	var rs *redis.IntCmd
	if len(step)>0 {
		rs = redisCli.IncrBy(key,step[0])
	} else{
		rs = redisCli.Incr(key)
	}
	return rs.Val()
}

func (h *RedisStore) Decrement(key string, step ...int64) int64{
	var rs *redis.IntCmd
	if len(step)>0 {
		rs = redisCli.DecrBy(key,step[0])
	}else{
		rs = redisCli.Decr(key)
	}
	return rs.Val()
}

func (h *RedisStore) Remember(key string,age time.Duration,fn func(args ...interface{}) string) {
	//return nil
}

func (h *RedisStore) Delete(key string)   {
	redisCli.Del(key)
}

func (h *RedisStore) HashSet(key string, value interface{}, duration time.Duration) error{
	if duration > 0{
		s :=redisCli.Set(key,value,duration)
		if s.Err() !=nil{
			fmt.Println("cache redis set failed :",s.Err())
			return s.Err()
		}
	} else{
		g :=redisCli.GetSet(key, value)
		if g.Err() != nil {
			fmt.Println("cache redis getset failed :",g.Err())
			return g.Err()
		}
	}

	return nil
}

func (h *RedisStore) HashGet(key string) interface{}{
	c :=  redisCli.Get(key)
	if c.Err() != nil {
		fmt.Println("redis get failed:", c.Err())
		return ""
	}
	return c.Val()
}
