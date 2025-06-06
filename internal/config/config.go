/*
 * @Author: zs
 * @Date: 2025-05-30 11:58:06
 * @LastEditors: zs
 * @LastEditTime: 2025-06-04 20:29:17
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
	Cache    CacheConfig   `mapstructure:"cache"`
}

type ServerConfig struct {
	Port string `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
}

type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
	SSLMode  string `mapstructure:"sslmode"`
}

type RedisConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
	DB   int    `mapstructure:"db"`
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