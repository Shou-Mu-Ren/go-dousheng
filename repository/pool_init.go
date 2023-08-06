package repository

import (
	"github.com/gomodule/redigo/redis"
)

/*
	Redis 连接池
 */

var pool *redis.Pool // 创建redis连接池

func Pool_Init() {
	// 进行实例化连接池
	pool = &redis.Pool{
		MaxIdle: 16,	// 最初的连接数量
		MaxActive: 10,//最大连接数量
		// MaxActive: 0,	//连接池最大连接数量，不确定可以用0 （0表示自动定义）按需分配
		IdleTimeout: 24 * 3600,//连接关闭时间300秒 (300秒不使用的话就自动关闭)

		// 连接数据库
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp","localhost:6379")
		},
	}
}