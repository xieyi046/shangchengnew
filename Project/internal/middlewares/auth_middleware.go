package middlewares

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/shangcheng/Project/Project/internal/consts"
	ctl "github.com/shangcheng/Project/Project/pkg/ctl"
	"github.com/shangcheng/Project/Project/pkg/util/jwt"
)

// AuthMiddleware token验证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := c.GetHeader(consts.AccessTokenHeader)
		refreshToken := c.GetHeader(consts.RefreshTokenHeader)

		if accessToken == "" {
			resp := ctl.RespError(c, errors.New("token不能为空"), "Token 为空")
			c.JSON(200, resp)
			c.Abort()
			return
		}

		newAccessToken, newRefreshToken, err := jwt.ParseRefreshToken(accessToken, refreshToken)
		if err != nil {
			resp := ctl.RespError(c, err, "鉴权失败")
			c.JSON(200, resp)
			c.Abort()
			return
		}

		claims, err := jwt.ParseToken(newAccessToken)
		if err != nil {
			resp := ctl.RespError(c, err, "解析 access token 失败")
			c.JSON(200, resp)
			c.Abort()
			return
		}

		SetToken(c, newAccessToken, newRefreshToken)
		userInfo := &ctl.UserInfo{ID: claims.ID}
		c.Request = c.Request.WithContext(ctl.NewContext(c.Request.Context(), userInfo))
		ctl.InitUserInfo(c.Request.Context())
		c.Next()
	}
}

func SetToken(c *gin.Context, accessToken, refreshToken string) {
	secure := IsHttps(c)
	c.Header(consts.AccessTokenHeader, accessToken)
	c.Header(consts.RefreshTokenHeader, refreshToken)
	c.SetCookie(consts.AccessTokenHeader, accessToken, consts.MaxAge, "/", "", secure, true)
	c.SetCookie(consts.RefreshTokenHeader, refreshToken, consts.MaxAge, "/", "", secure, true)
}

func IsHttps(c *gin.Context) bool {
	return c.GetHeader(consts.HeaderForwardedProto) == "https" || c.Request.TLS != nil
}
