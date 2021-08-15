package redisutils

import (
	"context"
	"encoding/json"

	"github.com/go-redis/redis/v8"
)

func RedisNewClient(host, pass string, db int, ctx context.Context) (client *redis.Client, err error) {

	client = redis.NewClient(&redis.Options{
		Addr:     host,
		Password: pass,
		DB:       db,
	})

	_, err = client.Ping(ctx).Result()

	return
}

func SaveToRedis(client *redis.Client, key string, value interface{}, ctx context.Context) error {
	status := client.Set(ctx, key, value, 0)
	return status.Err()
}

func GetRedisValue(client *redis.Client, key string, ctx context.Context) (result string, err error) {
	result, err = client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil
	}
	return
}

func GetRedisValueWithUnmarshal(client *redis.Client, key string, resultStruct interface{}, ctx context.Context) (result string, err error) {
	result, err = GetRedisValue(client, key, ctx)

	if err != nil {
		return
	}

	if resultStruct != nil && len(result) > 0 {
		err = json.Unmarshal([]byte(result), resultStruct)
	}
	return
}
