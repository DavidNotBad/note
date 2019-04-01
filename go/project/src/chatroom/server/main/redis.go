package main

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

var pool *redis.Pool

func initPool(address string, maxIdle int, maxActive int, idleTimeout time.Duration) {
	pool = &redis.Pool{
		MaxIdle: maxIdle, //最大空闲连接数
		MaxActive: maxActive, //表示和数据库的最大连接数， 0表示没有限制
		IdleTimeout: idleTimeout, //最大空闲时间
		Dial: func() (conn redis.Conn, e error) {
			return redis.Dial("tcp", address)
		},
	}
}



