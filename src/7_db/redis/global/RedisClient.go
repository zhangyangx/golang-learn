package global

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

//var RedisClient redis.Conn

var RedisClient *redis.Pool

func InitRedisClient() {

	RedisClient = &redis.Pool{
		// 连接方法
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "127.0.0.1:6379")
			if err != nil {
				return nil, err
			}
			_, err = c.Do("SELECT", 0)
			if err != nil {
				return nil, err
			}
			return c, nil
		},
		//DialContext:     nil,
		//TestOnBorrow:    nil,
		// 最大的空闲连接数，表示即使没有redis连接时依然可以保持N个空闲的连接，而不被清除，随时处于待命状态。
		MaxIdle: 1,
		// 最大的激活连接数，表示同时最多有N个连接
		MaxActive: 10,
		// 最大的空闲连接等待时间，超过此时间后，空闲连接将被关闭
		IdleTimeout: 180 * time.Second,
		//Wait:            false,
		//MaxConnLifetime: 0,
	}

	//第一种连接方法
	//c1, err := redis.Dial("tcp", "127.0.0.1:6379")
	//if err != nil {
	//	panic(err)
	//}
	//RedisClient = c1
	//defer func(c1 redis.Conn) {
	//	err := c1.Close()
	//	if err != nil {
	//		panic(err)
	//	}
	//}(c1)
	//第二种连接方法 其实这个最后还是会调用Dial
	//c2,err := redis.DialURL("redis://127.0.0.1:6379")
	//if err != nil {
	//	panic(err)
	//}
}
