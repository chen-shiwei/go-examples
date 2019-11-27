package redis

import (
	"fmt"
	"github.com/go-redis/redis"
)

var (
	cli *redis.Client
)

func init() {
	_ = MustInit()
}

func MustInit() *redis.Client {
	//addr: 182.92.69.60:63778
	//password: redis2016*()$
	//db: 4
	//pool: 5
	redisOpt := redis.Options{
		Network:  "tcp",
		Addr:     "182.92.69.60:63778",
		Password: "redis2016*()$",
		DB:       14,
		//MaxRetries:         0,
		//MinRetryBackoff:    0,
		//MaxRetryBackoff:    0,
		//DialTimeout:        0,
		//ReadTimeout:        0,
		//WriteTimeout:       0,
		PoolSize: 2,
		//MinIdleConns:       0,
		//MaxConnAge:         0,
		//PoolTimeout:        0,
		//IdleTimeout:        0,
		//IdleCheckFrequency: 0,
		//TLSConfig:          nil,
	}

	cli = redis.NewClient(&redisOpt)

	pong, err := cli.Ping().Result()
	if err != nil {
		panic(fmt.Errorf("redis ping err: %s", err))
		return nil
	}
	fmt.Printf("redis： ping:%s \n", pong)
	return cli
}
