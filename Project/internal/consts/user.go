package consts

import "time"

// HTTP Header 常量
const (
	AccessTokenHeader    = "access_token"
	RefreshTokenHeader   = "refresh_token"
	HeaderForwardedProto = "X-Forwarded-Proto"
)

// Cookie 设置常量
const (
	MaxAge = 3600
)

// Token 过期时间设置
const (
	AccessTokenExpireDuration  = 15 * time.Minute   // Access Token 过期时间
	RefreshTokenExpireDuration = 7 * 24 * time.Hour // Refresh Token 过期时间
)

// 错误码定义
const (
	SUCCESS                 = 200
	InvalidParams           = 400
	ErrorAuthCheckTokenFail = 401
)
