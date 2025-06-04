/*
 * @Author: zs
 * @Date: 2025-06-03 18:01:24
 * @LastEditors: zs
 * @LastEditTime: 2025-06-03 18:09:27
 * @FilePath: /barshop-server/internal/logger/logger.go
 * @Description: 
 * 
 * Copyright (c) 2025 by zs, All Rights Reserved. 
 */
package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
	"time"
)

var Log *zap.Logger

// Error 将error转换为zap.Field
func Error(err error) zap.Field {
	return zap.Error(err)
}

// InitLogger 初始化日志系统
func InitLogger(level string, isDevelopment bool) error {
	// 解析日志级别
	logLevel, err := zapcore.ParseLevel(level)
	if err != nil {
		return fmt.Errorf("invalid log level: %v", err)
	}

	// 创建日志目录
	logDir := "logs"
	if err := os.MkdirAll(logDir, 0755); err != nil {
		return fmt.Errorf("failed to create log directory: %v", err)
	}

	// 配置编码器
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// 创建核心
	var core zapcore.Core
	if isDevelopment {
		// 开发模式：同时写入文件和控制台
		consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)
		consoleCore := zapcore.NewCore(
			consoleEncoder,
			zapcore.AddSync(os.Stdout),
			logLevel,
		)

		// 文件输出
		now := time.Now()
		logFile := filepath.Join(logDir, fmt.Sprintf("%s.log", now.Format("2006-01-02")))
		writer, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return fmt.Errorf("failed to open log file: %v", err)
		}

		fileEncoder := zapcore.NewJSONEncoder(encoderConfig)
		fileCore := zapcore.NewCore(
			fileEncoder,
			zapcore.AddSync(writer),
			logLevel,
		)

		core = zapcore.NewTee(consoleCore, fileCore)
	} else {
		// 生产模式：只写入文件
		now := time.Now()
		logFile := filepath.Join(logDir, fmt.Sprintf("%s.log", now.Format("2006-01-02")))
		writer, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return fmt.Errorf("failed to open log file: %v", err)
		}

		fileEncoder := zapcore.NewJSONEncoder(encoderConfig)
		core = zapcore.NewCore(
			fileEncoder,
			zapcore.AddSync(writer),
			logLevel,
		)
	}

	// 创建logger
	Log = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	return nil
} 