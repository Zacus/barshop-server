/*
 * @Author: zs
 * @Date: 2025-06-04 20:40:23
 * @LastEditors: zs
 * @LastEditTime: 2025-06-04 20:40:23
 * @FilePath: /barshop-server/internal/cache/keys.go
 * @Description: 缓存键管理器
 * 
 * Copyright (c) 2025 by zs, All Rights Reserved. 
 */
package cache

import "fmt"

const (
	// 键前缀
	UserPrefix        = "user:"       // 用户相关的缓存键前缀
	ServicePrefix     = "service:"    // 服务相关的缓存键前缀
	AppointmentPrefix = "appointment:" // 预约相关的缓存键前缀
	StatisticsPrefix  = "stats:"      // 统计相关的缓存键前缀
	LockPrefix        = "lock:"       // 分布式锁相关的缓存键前缀

	// 过期时间（秒）
	UserExpire        = 3600      // 用户缓存过期时间
	ServiceExpire     = 7200      // 服务缓存过期时间
	AppointmentExpire = 1800      // 预约缓存过期时间
	StatisticsExpire  = 86400     // 统计缓存过期时间
	LockExpire        = 10        // 分布式锁过期时间
	TokenExpire       = 604800    // Token缓存过期时间（7天）
)

// 用户相关的键
func UserKey(userID uint) string {
	return fmt.Sprintf("%s%d", UserPrefix, userID)
}

func UserTokenKey(userID uint) string {
	return fmt.Sprintf("%stoken:%d", UserPrefix, userID)
}

func UserStatsKey(userID uint) string {
	return fmt.Sprintf("%sstats:%d", UserPrefix, userID)
}

// 服务相关的键
func ServiceKey(serviceID uint) string {
	return fmt.Sprintf("%s%d", ServicePrefix, serviceID)
}

func ServiceListKey() string {
	return fmt.Sprintf("%slist", ServicePrefix)
}

func ServiceStatsKey(serviceID uint) string {
	return fmt.Sprintf("%sstats:%d", ServicePrefix, serviceID)
}

// 预约相关的键
func AppointmentKey(appointmentID uint) string {
	return fmt.Sprintf("%s%d", AppointmentPrefix, appointmentID)
}

func AppointmentListKey(userID uint) string {
	return fmt.Sprintf("%slist:%d", AppointmentPrefix, userID)
}

func AppointmentDateKey(date string) string {
	return fmt.Sprintf("%sdate:%s", AppointmentPrefix, date)
}

// 统计相关的键
func DailyStatsKey(date string) string {
	return fmt.Sprintf("%sdaily:%s", StatisticsPrefix, date)
}

func MonthlyStatsKey(yearMonth string) string {
	return fmt.Sprintf("%smonthly:%s", StatisticsPrefix, yearMonth)
}

func YearlyStatsKey(year string) string {
	return fmt.Sprintf("%syearly:%s", StatisticsPrefix, year)
}

// 分布式锁相关的键
func LockKey(resource string) string {
	return fmt.Sprintf("%s%s", LockPrefix, resource)
}

func AppointmentLockKey(barberID uint, startTime string) string {
	return fmt.Sprintf("%sappointment:%d:%s", LockPrefix, barberID, startTime)
} 