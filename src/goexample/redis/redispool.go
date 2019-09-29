package main

import (
	"github.com/garyburd/redigo/redis"
	"fmt"
	"time"
)
 
var (
	// 定义常量
	RedisClient     *redis.Pool
	REDIS_HOST string
	REDIS_DB   int
	REDIS_PASS string
)
 
func initPool() {
	// 从配置文件获取redis的ip以及db
	REDIS_HOST = "103.242.202.197:6379"
	REDIS_DB = 1
	REDIS_PASS = "o1Akq6ShccAHcwSbYFjq"
	options := redis.DialPassword("o1Akq6ShccAHcwSbYFjq")
	// 建立连接池
	RedisClient = &redis.Pool{
		// 从配置文件获取maxidle以及maxactive，取不到则用后面的默认值
		MaxIdle:     1,
		MaxActive:   10,
		IdleTimeout: 180 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", REDIS_HOST, options)
			if err != nil {
				return nil, err
			}
			// 选择db
			c.Do("SELECT", REDIS_DB)
			return c, nil
		},
	}
}

func main()  {
	initPool()
	
	rc := RedisClient.Get()
	defer rc.Close()
	rc.Do("LPUSH", "netgame:info", "yanetao")

	

	netgame, err = rc.Do("GET", "netgame:info")
	
	if err != nil {
        fmt.Println("redis get failed:", err)
    } else {
        fmt.Printf("Get netgame: %v \n", netgame)
    }

}
