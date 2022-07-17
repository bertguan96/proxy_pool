package test

import (
	"project/proxy_pool/db"
	"testing"
)

// TestHSet 测试redis set能力
func TestHSet(t *testing.T) {
	db.HSet("test", "11", "12")
	str := db.HGet("test", "11")
	if str == "12" {
		t.Failed()
	}
}

// TestHExits 测试Key是否存在
func TestHExits(t *testing.T) {
	res := db.HExists("test", "11")
	if res == false {
		t.Failed()
	}
}
