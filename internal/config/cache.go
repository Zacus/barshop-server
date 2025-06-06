/*
 * @Author: zs
 * @Date: 2025-06-04 20:30:23
 * @LastEditors: zs
 * @LastEditTime: 2025-06-04 20:30:23
 * @FilePath: /barshop-server/internal/config/cache.go
 * @Description: 缓存配置
 * 
 * Copyright (c) 2025 by zs, All Rights Reserved. 
 */
package config

// CacheConfig Redis缓存配置
type CacheConfig struct {
	Host     string `mapstructure:"host"`     // Redis主机地址
	Port     int    `mapstructure:"port"`     // Redis端口
	Password string `mapstructure:"password"` // Redis密码
	DB       int    `mapstructure:"db"`       // Redis数据库编号
	PoolSize int    `mapstructure:"poolsize"` // 连接池大小
} 