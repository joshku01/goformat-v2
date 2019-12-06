package models

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"goformat-v2/app/global"
	"goformat-v2/app/global/errorcode"
	"log"
	"sync"
	"time"
)

// redisPool 存放redis連線池的全域變數
var redisPool *redis.Pool

type redisConnect struct {
	// another stuff , may be needed by middleware
}

// Redis 存取值
type Redis struct {
	redisConn redisConnect
}

var redisSingleton *Redis
var redisOnce sync.Once

// NewRedisPool redis 連接池初始化
func (*redisConnect) NewRedisPool() *redis.Pool {

	redisPool = &redis.Pool{
		MaxIdle:     global.Config.RedisPool.MaxIdle,                                  // int 最大可允許的閒置連線數
		MaxActive:   global.Config.RedisPool.MaxActive,                                // int 最大建立的連線數，默認為0不限制(reids 預設最大連線量)
		IdleTimeout: time.Duration(global.Config.RedisPool.IdleTimeout) * time.Second, // 連線過期時間，默認為0表示不做過期限制
		Wait:        true,                                                             // 當連線超出限制數量後，是否等待到空閒連線釋放
		Dial: func() (c redis.Conn, err error) {
			// 使用redis封裝的Dial進行tcp連接
			host := global.Config.Redis.RedisHost
			port := global.Config.Redis.RedisPort
			pwd := global.Config.Redis.RedisPwd

			// 組合連接資訊
			var connectionString = fmt.Sprintf("%s:%s", host, port)
			c, err = redis.Dial(
				"tcp",
				connectionString,
				redis.DialPassword(pwd),
				redis.DialConnectTimeout(5*time.Second), // 建立連線 time out 時間 5 秒
				redis.DialReadTimeout(5*time.Second),    // 讀取資料 time out 時間 5 秒
				redis.DialWriteTimeout(5*time.Second),   // 寫入資料 time out 時間 5 秒
			)

			if err != nil {
				errorcode.ErrorHandler("REDIS_CONNECT_ERROR", err)
				return
			}
			return
		}, // 連接redis的函数
		TestOnBorrow: func(redis redis.Conn, t time.Time) (err error) {
			// 每5秒ping一次redis
			if time.Since(t) < (5 * time.Second) {
				return
			}

			_, err = redis.Do("PING")
			if err != nil {
				errorcode.ErrorHandler("REDIS_PING_ERROR", err)
				return
			}

			return
		}, // 定期對 redis server 做 ping/pong 測試
	}
	return redisPool
}

// RedisPoolConnect 回傳連線池的 Redis 連線
func (r *redisConnect) RedisPoolConnect() *redis.Pool {
	return r.NewRedisPool()
}

// RedisIns 獲得單例對象
//// NewRedisConnect will create new an RedisConnection object representation of IRedis interface
func NewRedisConnect(r redisConnect) IRedis {
	redisOnce.Do(func() {
		redisSingleton = &Redis{
			redisConn: r,
		}
	})
	return redisSingleton
}

// IRedis interface
type IRedis interface {
	Exists(key string) (ok bool, apiErr errorcode.Error)
	Set(key string, value interface{}, expiretime int) (apiErr errorcode.Error)
	Get(key string) (value string, apiErr errorcode.Error)
	Delete(key string) (apiErr errorcode.Error)
	Append(key string, value interface{}) (n interface{}, apiErr errorcode.Error)
	HashSet(hkey string, key interface{}, value interface{}, time int) (apiErr errorcode.Error)
	HashGet(hkey string, field interface{}) (value string, apiErr errorcode.Error)
	RedisPing()
}

// RedisPing 檢查Redis是否啟動
func (r *Redis) RedisPing() {
	RedisPool := r.redisConn.RedisPoolConnect()
	conn := RedisPool.Get()
	defer conn.Close()

	_, err := conn.Do("PING")
	if err != nil {
		log.Fatalf("🔔🔔🔔 REDIS CONNECT ERROR: %v 🔔🔔🔔", err.Error())
	}
}

// Exists 檢查key是否存在
func (r *Redis) Exists(key string) (ok bool, apiErr errorcode.Error) {
	RedisPool := r.redisConn.RedisPoolConnect()
	conn := RedisPool.Get()
	defer conn.Close()

	chkExisits, _ := conn.Do("EXISTS", key)
	ok, err := redis.Bool(chkExisits, nil)
	if err != nil {
		apiErr = errorcode.ErrorHandler("REDIS_CHECK_EXIST_ERROR", err)
		return
	}

	return
}

// Set 存入redis值
func (r *Redis) Set(key string, value interface{}, expiretime int) (apiErr errorcode.Error) {
	RedisPool := r.redisConn.RedisPoolConnect()
	conn := RedisPool.Get()
	defer conn.Close()

	if _, err := conn.Do("SET", key, value, "EX", expiretime); err != nil {
		apiErr = errorcode.ErrorHandler("REDIS_INSERT_ERROR", err)
		return
	}
	return
}

// Get 取出redis值
func (r *Redis) Get(key string) (value string, apiErr errorcode.Error) {
	RedisPool := r.redisConn.RedisPoolConnect()
	conn := RedisPool.Get()
	defer conn.Close()

	value, err := redis.String(conn.Do("GET", key))
	if err != nil {
		errorcode.ErrorHandler("REDIS_GET_VALUE_ERROR", err)
	}

	return
}

// Delete 刪除redis值
func (r *Redis) Delete(key string) (apiErr errorcode.Error) {
	RedisPool := r.redisConn.RedisPoolConnect()
	conn := RedisPool.Get()
	defer conn.Close()

	if _, err := conn.Do("DEL", key); err != nil {
		apiErr = errorcode.ErrorHandler("REDIS_DELETE_ERROR", err)
		return
	}

	return
}

// Append 在相同key新增多個值
func (r *Redis) Append(key string, value interface{}) (n interface{}, apiErr errorcode.Error) {
	RedisPool := r.redisConn.RedisPoolConnect()
	conn := RedisPool.Get()
	defer conn.Close()

	n, err := conn.Do("APPEND", key, value)
	if err != nil {
		apiErr = errorcode.ErrorHandler("REDIS_APPEND_ERROR", err)
		return
	}

	return
}

// HashSet Hash方式存入redis值
func (r *Redis) HashSet(hkey string, key interface{}, value interface{}, time int) (apiErr errorcode.Error) {
	RedisPool := r.redisConn.RedisPoolConnect()
	conn := RedisPool.Get()
	defer conn.Close()

	// 存值
	if _, err := conn.Do("hset", hkey, key, value); err != nil {
		apiErr = errorcode.ErrorHandler("REDIS_INSERT_ERROR", err)
		return
	}

	// 設置過期時間
	if _, err := conn.Do("EXPIRE", hkey, time); err != nil {
		apiErr = errorcode.ErrorHandler("REDIS_SET_EXPIRE_ERROR", err)
		return
	}

	return
}

// HashGet Hash方式取出redis值
func (r *Redis) HashGet(hkey string, field interface{}) (value string, apiErr errorcode.Error) {
	RedisPool := r.redisConn.RedisPoolConnect()
	conn := RedisPool.Get()
	defer conn.Close()

	// 取值
	value, err := redis.String(conn.Do("HGET", hkey, field))
	if err != nil {
		errorcode.ErrorHandler("REDIS_GET_VALUE_ERROR", err)
	}

	return
}
