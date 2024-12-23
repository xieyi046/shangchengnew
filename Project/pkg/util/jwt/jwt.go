package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/shangcheng/Project/internal/consts"
)

var jwtSecret = []byte("FanOne") // 建议从环境变量或配置文件加载

type Claims struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenerateToken 签发用户Token
func GenerateToken(id int, username string) (accessToken, refreshToken string, err error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(consts.AccessTokenExpireDuration)
	rtExpireTime := nowTime.Add(consts.RefreshTokenExpireDuration)

	claims := Claims{
		ID:       id,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "mall",
		},
	}

	accessToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtSecret)
	if err != nil {
		return "", "", err
	}

	refreshClaims := jwt.StandardClaims{
		ExpiresAt: rtExpireTime.Unix(),
		Issuer:    "mall",
	}
	refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString(jwtSecret)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

// ParseToken 验证用户token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

// RefreshTokens 验证并刷新用户的 Token
func ParseRefreshToken(aToken, rToken string) (newAToken, newRToken string, err error) {
	accessClaim, err := ParseToken(aToken)
	if err != nil {
		return "", "", err
	}

	refreshClaim, err := ParseToken(rToken)
	if err != nil {
		return "", "", err
	}

	currentTime := time.Now().Unix()

	if accessClaim.ExpiresAt > currentTime || refreshClaim.ExpiresAt > currentTime {
		return GenerateToken(accessClaim.ID, accessClaim.Username)
	}

	return "", "", errors.New("身份过期，重新登录")
}
