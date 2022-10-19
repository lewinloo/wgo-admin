package middleware

import (
	"gin_template/internal/app/common"
	"gin_template/internal/app/config"
	"gin_template/internal/app/utils"
	"github.com/casbin/casbin/v2"
	"net/http"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
)

var checkLock sync.Mutex

func CheckPermission(casbin *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		subs := utils.GetUserRoleIds(c)
		isAdmin := utils.GetUserIsAdmin(c)

		if isAdmin {
			c.Next()
			return
		}

		// 获得请求路径URL
		obj := strings.TrimPrefix(c.FullPath(), config.C.System.GlobalPrefix)
		// 获取请求方式
		act := c.Request.Method

		isPass := check(casbin, subs, obj, act)
		if !isPass {
			c.JSON(http.StatusForbidden, common.Result{
				Success: false,
				Message: "没有权限",
				Data:    nil,
				Code:    403,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

func check(casbin *casbin.Enforcer, subs []string, obj string, act string) bool {
	// 同一时间只允许一个请求执行校验, 否则可能会校验失败
	checkLock.Lock()
	defer checkLock.Unlock()
	isPass := false
	for _, sub := range subs {
		pass, _ := casbin.Enforce(sub, obj, act)
		if pass {
			isPass = true
			break
		}
	}
	return isPass
}
