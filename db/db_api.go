package db

import (
	"math/rand"
	"project/proxy_pool/config"
	"time"
)

// GetIp 随机获取一个IP，返回的事一个IP JSON需要反序列化
func GetIp() string {
	rand.Seed(time.Now().Unix())
	keys := HKeys(config.DBKey) // 这个Key是随机生成的ID
	if len(keys) == 0 {
		return ""
	}
	val := rand.Intn(len(keys))
	return HGet(config.DBKey, keys[val])
}

// AddIp 添加一个IP
func AddIp(key string, value string) {
	HSet(config.DBKey, key, value)
}

// DelIp 删除IP
func DelIp(key string) bool {
	return HDel(config.DBKey, key)
}

// ClearAll 清除所有
func ClearAll() bool {
	return HClear(config.DBKey)
}

// GetAll 获取所有IP内容
func GetAll() map[string]string {
	return HGetAll(config.DBKey)
}

// SetExpire 设置超时时间
func SetExpire(time time.Duration) bool {
	return Expire(config.DBKey, time)
}

// Exists 判断key是否存在
func Exists(key string) bool {
	return HExists(config.DBKey, key)
}
