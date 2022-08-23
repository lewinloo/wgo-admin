package utils

import (
	"gin_template/internal/global"
	"time"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"

	"github.com/dgrijalva/jwt-go"
)

var JwtSecret = []byte(global.CONFIG.Jwt.Secret)

// TODO 根据需要改
type Claims struct {
	Id       uint      `json:"id"`
	UUID     uuid.UUID `json:"uuid"`
	UserName string    `json:"username"`
	jwt.StandardClaims
}

// 生成token
func GenerateToken(id uint, uuid uuid.UUID, username string) (string, error) {
	expireTime := time.Now().Add(time.Duration(global.CONFIG.Jwt.ExpireTime) * time.Hour)
	claims := Claims{
		Id:       id,
		UserName: username,
		UUID:     uuid,
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

func GetUserRoleId(c *gin.Context) string {
	if roleId, exists := c.Get("roleId"); !exists {
		return ""
	} else {
		return roleId.(string)
	}
}

func GetUserUUID(c *gin.Context) uuid.UUID {
	if id, exists := c.Get("uuid"); !exists {
		return id.(uuid.UUID)
	} else {
		return id.(uuid.UUID)
	}
}