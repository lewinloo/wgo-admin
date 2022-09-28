package utils

import (
	"gin_template/internal/global"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/dgrijalva/jwt-go"
)

var JwtSecret = []byte(global.CONFIG.Jwt.Secret)

type Claims struct {
	Id       uint     `json:"id"`
	Username string   `json:"username"`
	RoleIds  []string `json:"roleIds"`
	jwt.StandardClaims
}

// 生成token
func GenerateToken(id uint, username string, roleIds []string) (string, error) {
	expireTime := time.Now().Add(time.Duration(global.CONFIG.Jwt.ExpireTime) * time.Hour)
	claims := Claims{
		Id:       id,
		Username: username,
		RoleIds:  roleIds,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    global.CONFIG.Jwt.Issuer,
			Subject:   global.CONFIG.Jwt.Subject,
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(JwtSecret)
	return token, err
}

// 解析token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return JwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

func GetUserRoleIds(c *gin.Context) []string {
	if roleIds, exists := c.Get("roleIds"); !exists {
		return []string{}
	} else {
		return roleIds.([]string)
	}
}
