package cache

import (
	"context"
	"log"

	"github.com/go-redis/redis/v9"
)

var rdb *redis.Client

func NewRedisCache() *redis.Client {
	initRedisCache()
	return rdb
}

func initRedisCache() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	ctx := context.Background()
	err = rdb.Ping(ctx).Err()
	return
}

// Ping runs a health check...
func Ping(ctx context.Context) (err error) {
	if err = rdb.Ping(ctx).Err(); err != nil {
		log.Printf("ERROR : %+v", err)
		return err
	}
	return nil
}
