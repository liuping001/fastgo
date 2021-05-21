// Author: coolliu
// Date: 2021/5/21

package myredis

import (
	"flag"
	"github.com/go-redis/redis/v7"
	"strings"
)

var (
	redisAddr = flag.String("redisAddr", "127.0.0.1:6379", "redis 地址")
	redisPass = flag.String("redisPass", "", "redis 密码")
)

func NewRedisClient(addr string, pass string) redis.Cmdable {
	addrList := strings.Split(addr, ",")
	if len(addrList) > 1 {
		return redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:    addrList,
			Password: pass,
		})
	} else {
		return redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: pass,
		})
	}
}

var Client redis.Cmdable

func Init() {
	Client = NewRedisClient(*redisAddr, *redisPass)
}
