/*
 * @Author: zs
 * @Date: 2025-06-04 20:35:23
 * @LastEditors: zs
 * @LastEditTime: 2025-06-04 20:35:23
 * @FilePath: /barshop-server/internal/cache/redis.go
 * @Description: Redis缓存管理器
 * 
 * Copyright (c) 2025 by zs, All Rights Reserved. 
 */
package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/zacus/barshop-server/internal/config"
)

var (
	Client *redis.Client
	ctx    = context.Background()
)

// InitRedis 初始化Redis连接
func InitRedis(cfg *config.Config) error {
	Client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Cache.Host, cfg.Cache.Port),
		Password: cfg.Cache.Password,
		DB:       cfg.Cache.DB,
		PoolSize: cfg.Cache.PoolSize,
	})

	// 测试连接
	if err := Client.Ping(ctx).Err(); err != nil {
		return fmt.Errorf("failed to connect to redis: %v", err)
	}

	return nil
}

// Set 设置缓存
func Set(key string, value interface{}, expiration time.Duration) error {
	// 将值转换为JSON
	jsonValue, err := json.Marshal(value)
	if err != nil {
		return err
	}

	// 设置缓存
	return Client.Set(ctx, key, jsonValue, expiration).Err()
}

// Get 获取缓存
func Get(key string, value interface{}) error {
	// 获取缓存
	jsonValue, err := Client.Get(ctx, key).Result()
	if err != nil {
		return err
	}

	// 解析JSON到目标结构
	return json.Unmarshal([]byte(jsonValue), value)
}

// Delete 删除缓存
func Delete(key string) error {
	return Client.Del(ctx, key).Err()
}

// Exists 检查键是否存在
func Exists(key string) bool {
	result, _ := Client.Exists(ctx, key).Result()
	return result > 0
}

// SetNX 仅当键不存在时设置缓存（可用于实现分布式锁）
func SetNX(key string, value interface{}, expiration time.Duration) (bool, error) {
	jsonValue, err := json.Marshal(value)
	if err != nil {
		return false, err
	}

	return Client.SetNX(ctx, key, jsonValue, expiration).Result()
}

// GetTTL 获取键的剩余生存时间
func GetTTL(key string) (time.Duration, error) {
	return Client.TTL(ctx, key).Result()
}

// Incr 将键的整数值加1
func Incr(key string) (int64, error) {
	return Client.Incr(ctx, key).Result()
}

// IncrBy 将键的整数值加上指定的增量
func IncrBy(key string, increment int64) (int64, error) {
	return Client.IncrBy(ctx, key, increment).Result()
}

// Decr 将键的整数值减1
func Decr(key string) (int64, error) {
	return Client.Decr(ctx, key).Result()
}

// DecrBy 将键的整数值减去指定的减量
func DecrBy(key string, decrement int64) (int64, error) {
	return Client.DecrBy(ctx, key, decrement).Result()
}

// HSet 设置哈希表字段的值
func HSet(key, field string, value interface{}) error {
	jsonValue, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return Client.HSet(ctx, key, field, jsonValue).Err()
}

// HGet 获取哈希表字段的值
func HGet(key, field string, value interface{}) error {
	jsonValue, err := Client.HGet(ctx, key, field).Result()
	if err != nil {
		return err
	}

	return json.Unmarshal([]byte(jsonValue), value)
}

// HDelete 删除哈希表字段
func HDelete(key string, fields ...string) error {
	return Client.HDel(ctx, key, fields...).Err()
}

// Close 关闭Redis连接
func Close() error {
	if Client != nil {
		return Client.Close()
	}
	return nil
} 