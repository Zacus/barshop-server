package validator

import (
	"github.com/go-playground/validator/v10"
	"regexp"
	"strings"
	"unicode"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
	
	// 注册自定义验证器
	_ = validate.RegisterValidation("phone_cn", validateChinesePhone)
	_ = validate.RegisterValidation("username", validateUsername)
	_ = validate.RegisterValidation("password_strength", validatePasswordStrength)
}

// ValidateStruct 验证结构体
func ValidateStruct(s interface{}) error {
	return validate.Struct(s)
}

// CleanString 清理字符串（去除首尾空格，将多个空格替换为单个空格）
func CleanString(s string) string {
	return strings.TrimSpace(regexp.MustCompile(`\s+`).ReplaceAllString(s, " "))
}

// CleanPhone 清理电话号码（只保留数字）
func CleanPhone(phone string) string {
	return regexp.MustCompile(`[^\d]`).ReplaceAllString(phone, "")
}

// CleanEmail 清理邮箱地址（转为小写并去除空格）
func CleanEmail(email string) string {
	return strings.ToLower(strings.TrimSpace(email))
}

// validateChinesePhone 验证中国手机号
func validateChinesePhone(fl validator.FieldLevel) bool {
	phone := fl.Field().String()
	match, _ := regexp.MatchString(`^1[3-9]\d{9}$`, phone)
	return match
}

// validateUsername 验证用户名（字母开头，只能包含字母、数字和下划线，长度4-20）
func validateUsername(fl validator.FieldLevel) bool {
	username := fl.Field().String()
	if len(username) < 4 || len(username) > 20 {
		return false
	}
	match, _ := regexp.MatchString(`^[a-zA-Z][a-zA-Z0-9_]*$`, username)
	return match
}

// validatePasswordStrength 验证密码强度
func validatePasswordStrength(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	if len(password) < 8 {
		return false
	}
	
	var (
		hasUpper   bool
		hasLower   bool
		hasNumber  bool
		hasSpecial bool
	)
	
	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}
	
	return hasUpper && hasLower && hasNumber && hasSpecial
}