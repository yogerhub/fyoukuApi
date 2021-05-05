package redisClient

import (
	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
	"time"
)

// Connect 直接连接
func Connect() redis.Conn {
	pool, _ := redis.Dial("tcp", beego.AppConfig.String("redisdb"))
	return pool
}

// PoolConnect 通过连接池连接
func PoolConnect() redis.Conn {
	pool := &redis.Pool{
		MaxIdle:     1,                 //最大的空闲连接数
		MaxActive:   10,                //最大连接数
		IdleTimeout: 180 * time.Second, //空闲连接超时时间
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", beego.AppConfig.String("redisdb"))
			if err != nil {
				return nil, err
			}
			return c, nil
		},
	}
	return pool.Get()
}
