/*
 * @Author: zs
 * @Date: 2025-05-30 11:58:06
 * @LastEditors: zs
 * @LastEditTime: 2025-06-09 17:17:56
 * @FilePath: /barshop-server/internal/config/config.go
 * @Description: 
 * 
 * Copyright (c) 2025 by zs, All Rights Reserved. 
 */
package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
	Redis    RedisConfig    `mapstructure:"redis"`
	JWT      JWTConfig     `mapstructure:"jwt"`
	Log      LogConfig     `mapstructure:"log"`
}

type ServerConfig struct {
	Port string `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
}

type DatabaseConfig struct {
	Driver       string        `mapstructure:"driver"`
	Host         string        `mapstructure:"host"`
	Port         int          `mapstructure:"port"`
	User         string        `mapstructure:"user"`
	Password     string        `mapstructure:"password"`
	DBName       string        `mapstructure:"dbname"`
	Charset      string        `mapstructure:"charset"`
	SSLMode      string        `mapstructure:"sslmode"`
	AutoMigrate  bool          `mapstructure:"auto_migrate"`
	Options      DatabaseOptions `mapstructure:"options"`
}

type DatabaseOptions struct {
	MaxIdleConns     int   `mapstructure:"max_idle_conns"`
	MaxOpenConns     int   `mapstructure:"max_open_conns"`
	ConnMaxLifetime  int   `mapstructure:"conn_max_lifetime"`
	ConnMaxIdleTime  int   `mapstructure:"conn_max_idle_time"`
	Debug           bool  `mapstructure:"debug"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`     // Redis主机地址
	Port     int    `mapstructure:"port"`     // Redis端口
	Password string `mapstructure:"password"` // Redis密码
	DB       int    `mapstructure:"db"`       // Redis数据库编号
	PoolSize int    `mapstructure:"poolsize"` // 连接池大小
}

type JWTConfig struct {
	Secret string `mapstructure:"secret"`
	Expire string `mapstructure:"expire"`
}

type LogConfig struct {
	Level    string `mapstructure:"level"`
	IsDev    bool   `mapstructure:"is_dev"`
}

func LoadConfig(path string) (*Config, error) {
	viper.SetConfigFile(path)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	config := &Config{}
	if err := viper.Unmarshal(config); err != nil {
		return nil, err
	}

	return config, nil
} 