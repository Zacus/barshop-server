/*
 * @Author: zs
 * @Date: 2025-06-03 18:01:24
 * @LastEditors: zs
 * @LastEditTime: 2025-06-04 17:55:36
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

type Logger struct {
	zapLogger *zap.Logger
}

var defaultLogger *Logger

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
		// 开发模式：输出到控制台
		consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)
		core = zapcore.NewCore(
			consoleEncoder,
			zapcore.AddSync(os.Stdout),
			logLevel,
		)
	} else {
		// 生产模式：同时输出到文件和控制台
		now := time.Now()
		logfile := filepath.Join(logDir, fmt.Sprintf("%s.log", now.Format("2006-01-02")))
		writer, err := os.OpenFile(logfile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			return fmt.Errorf("failed to create log file: %v", err)
		}

		jsonEncoder := zapcore.NewJSONEncoder(encoderConfig)
		core = zapcore.NewTee(
			zapcore.NewCore(jsonEncoder, zapcore.AddSync(writer), logLevel),
			zapcore.NewCore(jsonEncoder, zapcore.AddSync(os.Stdout), logLevel),
		)
	}

	// 创建日志记录器
	zapLogger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	defaultLogger = &Logger{zapLogger: zapLogger}
	return nil
}

// Debug 输出Debug级别日志
func Debug(msg string, args ...interface{}) {
	fields := makeFields(args...)
	defaultLogger.zapLogger.Debug(msg, fields...)
}

// Info 输出Info级别日志
func Info(msg string, args ...interface{}) {
	fields := makeFields(args...)
	defaultLogger.zapLogger.Info(msg, fields...)
}

// Warn 输出Warn级别日志
func Warn(msg string, args ...interface{}) {
	fields := makeFields(args...)
	defaultLogger.zapLogger.Warn(msg, fields...)
}

// Error 输出Error级别日志
func Error(msg string, args ...interface{}) {
	fields := makeFields(args...)
	defaultLogger.zapLogger.Error(msg, fields...)
}

// Fatal 输出Fatal级别日志
func Fatal(msg string, args ...interface{}) {
	fields := makeFields(args...)
	defaultLogger.zapLogger.Fatal(msg, fields...)
}

// Sync 同步日志缓冲
func Sync() error {
	return defaultLogger.zapLogger.Sync()
}

// makeFields 将参数转换为zap.Field
func makeFields(args ...interface{}) []zap.Field {
	fields := make([]zap.Field, 0, len(args))
	for i := 0; i < len(args); i += 2 {
		if i+1 < len(args) {
			fields = append(fields, zap.Any(fmt.Sprint(args[i]), args[i+1]))
		}
	}
	return fields
} 