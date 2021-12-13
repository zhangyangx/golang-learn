package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"golang-learn/src/7_db/redis/global"
)

func main() {
	// 初始化redis客户端
	global.InitRedisClient()
	client := global.RedisClient.Get()

	// 创建永不过期key
	res, err := client.Do("Set", "myKey", "hello golang")
	if err != nil {
		panic(err)
	}
	// res:OK
	fmt.Println(res)

	// 创建key，设置5秒过期
	//res, err := client.Do("Set", "myKey", "hello golang", "EX", 5)
	//if err != nil {
	//	panic(err)
	//}
	//// res:OK
	//fmt.Println(res)

	// 根据key获取缓存
	//res, err := client.Do("Get", "myKey")
	//if err != nil {
	//	panic(err)
	//}
	//if res != nil {
	//	// res:hello golang
	//	fmt.Println(string(res.([]byte)))
	//}

	// 删除key
	//res, err := client.Do("DEL", "myKey")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(res)

	// 为key设置过期时间
	//res, err = client.Do("EXPIRE", "myKey", 7200)
	//if err != nil {
	//	panic(err)
	//}
	//// res:1
	//fmt.Println(res)

	// list操作
	_, _ = client.Do("lpush", "mylist", "redis")
	_, _ = client.Do("lpush", "mylist", "mongodb")
	_, _ = client.Do("lpush", "mylist", "mysql")
	values, _ := redis.Values(client.Do("lrange", "mylist", "0", "-1"))
	for _, v := range values {
		fmt.Println(string(v.([]byte)))
	}

}
