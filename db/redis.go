package db

import (
	"context"
	"github.com/go-redis/redis/v8"
	"project/proxy_pool/config"
	"time"
)

var ctx = context.Background()

var rdb *redis.Client

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     config.DBHost,
		Password: config.DBPassword, // 设置密码
		DB:       config.DB,         // 选择存储的DB
	})
}

// HSet 新增数据
func HSet(key string, field string, value string) {
	//redis的键的set方法
	err := rdb.HSet(ctx, key, field, value).Err()
	if err != nil {
		panic(err)
	}
}

// HExists 判断Key是否存在
func HExists(key string, value string) bool {
	val, err := rdb.HExists(ctx, key, value).Result()
	if err != nil {
		panic(err)
	}
	return val
}

// HDel 删除key
func HDel(key string, field string) bool {
	err := rdb.HDel(ctx, key, field).Err()
	if err != nil {
		return false
	}
	return true
}

// HGetAll 获取所有
func HGetAll(key string) map[string]string {
	result, err := rdb.HGetAll(ctx, key).Result()
	if err != nil {
		return map[string]string{}
	}
	return result
}

// HClear 一键删除所有的Key
func HClear(key string) bool {
	err := rdb.Del(ctx, key).Err()
	if err != nil {
		return false
	}
	return true
}

// HGet 获取对应的Key
func HGet(key string, field string) string {
	result, err := rdb.HGet(ctx, key, field).Result()
	if err != nil {
		return ""
	}
	return result
}

// Expire 设置超时时间
func Expire(key string, expiration time.Duration) bool {
	result, err := rdb.Expire(ctx, key, expiration).Result()
	if err != nil {
		return false
	}
	return result
}

// HKeys 获取所有的key
func HKeys(key string) []string {
	result, err := rdb.HKeys(ctx, key).Result()
	if err != nil {
		return []string{}
	}
	return result
}
