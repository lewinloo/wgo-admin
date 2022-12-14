package utils

import (
	"gin_template/internal/app/config"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/dgrijalva/jwt-go"
)

var JwtSecret = []byte(config.C.Jwt.Secret)

type Claims struct {
	Id       uint     `json:"id"`
	Username string   `json:"username"`
	RoleIds  []string `json:"roleIds"`
	IsAdmin  bool     `json:"isAdmin"`
	jwt.StandardClaims
}

// 生成token
func GenerateToken(id uint, username string, roleIds []string, isAdmin bool) (string, error) {
	expireTime := time.Now().Add(time.Duration(config.C.Jwt.ExpireTime) * time.Hour)
	claims := Claims{
		Id:       id,
		Username: username,
		RoleIds:  roleIds,
		IsAdmin:  isAdmin,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    config.C.Jwt.Issuer,
			Subject:   config.C.Jwt.Subject,
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

func GetUserIsAdmin(c *gin.Context) bool {
	if isAdmin, exists := c.Get("isAdmin"); !exists {
		return false
	} else {
		return isAdmin.(bool)
	}
}
