/*
 * @Author: zs
 * @Date: 2025-06-04 16:53:38
 * @LastEditors: zs
 * @LastEditTime: 2025-06-08 14:44:03
 * @FilePath: /barshop-server/internal/models/response.go
 * @Description: 
 * 
 * Copyright (c) 2025 by zs, All Rights Reserved. 
 */
package models

// Response 通用响应结构
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// NewResponse 创建新的响应
func NewResponse(code int, message string, data interface{}) *Response {
	return &Response{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

// SuccessResponse 成功响应
func SuccessResponse(data interface{}) *Response {
	return NewResponse(200, "success", data)
}

// ErrorResponse 错误响应
func ErrorResponse(code int, message string) *Response {
	return NewResponse(code, message, nil)
} 