package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/go-redis/redis/v8"
)

const (
	// UserKeyPrefix 表示用户键的前缀
	UserKeyPrefix = "user:"

	// SessionKeyPrefix 表示会话键的前缀
	SessionKeyPrefix = "session:"

	// CounterKey 表示计数器键
	CounterKey = "counter"
)

var (
	once       sync.Once
	redisCli   *redis.Client
	ctx        = context.Background()
	redisAddr  string
	redisPwd   string
	redisDBNum int
)

// GetRedisClient 获取 Redis 客户端连接实例
func GetRedisClient() (*redis.Client, error) {

	var err error
	once.Do(func() {
		redisAddr = fmt.Sprintf("%s:%d", config.Redis.Host, config.Redis.Port)
		redisPwd = config.Redis.Password
		redisDBNum = config.Redis.DB

		// 创建 Redis 客户端连接
		redisCli = redis.NewClient(&redis.Options{
			Addr:     redisAddr,
			Password: redisPwd,
			DB:       redisDBNum,
		})

		// 检查是否能够连接到 Redis
		_, err = redisCli.Ping(ctx).Result()
		if err != nil {
			fmt.Printf("failed to connect to Redis: %v\n", err)
			return
		}
	})

	return redisCli, err
}

// Set 设置键值对
func SetRedisKey(key, value string) error {
	// 获取 Redis 客户端连接实例
	redisCli, err := GetRedisClient()
	if err != nil {
		return err
	}

	// 设置键值对
	return redisCli.Set(ctx, key, value, 0).Err()
}

// Get 获取键值对
func GetRedisByKey(key string) (string, error) {
	// 获取 Redis 客户端连接实例
	redisCli, err := GetRedisClient()
	if err != nil {
		return "", err
	}

	// 获取键值对
	return redisCli.Get(ctx, key).Result()
}

// Del 删除键
func DelRedisByKey(key string) error {
	// 获取 Redis 客户端连接实例
	redisCli, err := GetRedisClient()
	if err != nil {
		return err
	}

	// 删除键
	return redisCli.Del(ctx, key).Err()
}

// Close 关闭 Redis 客户端连接
func Close() error {
	// 获取 Redis 客户端连接实例
	redisCli, err := GetRedisClient()
	if err != nil {
		return err
	}

	// 关闭 Redis 客户端连接
	return redisCli.Close()
}

func SetRedisKeyByList(key string, values ...interface{}) error {
	redisCli, err := GetRedisClient()
	if err != nil {
		return err
	}

	// 设置键值对
	return redisCli.RPush(ctx, key, values).Err()
}

func GetRedisListByKey(key string) error {
	redisCli, err := GetRedisClient()
	if err != nil {
		return err
	}

	return redisCli.LRange(ctx, key, 0, -1).Err()
}

func PopByKey(key string) (string, error) {
	redisCli, err := GetRedisClient()
	if err != nil {
		return "", err
	}

	val, err := redisCli.RPop(ctx, key).Result()

	if err != nil {
		return "", err
	}

	return val, nil
}
