package middleware

import (
  "gin_template/internal/app/common"
  "github.com/gin-gonic/gin"
)

func Recovery() gin.HandlerFunc {
  return func(c *gin.Context) {
    defer func() {
      if err := recover(); err != any(nil) { //用于捕获panic
        common.ResponseFail(c, 500, "系统内部异常")
        c.Abort()
      }
    }()
    c.Next() // 调用下一个处理
  }
}
