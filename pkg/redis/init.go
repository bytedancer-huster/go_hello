package redis

import "github.com/go-redis/redis"

var Client *redis.Client

func init()  {
	Client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // use default Addr
		Password: "",               // no password set
	})
}