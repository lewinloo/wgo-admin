package handler

import "gin_template/internal/app/domain/service"

// 初始化 service
var (
	userService = service.NewUser()
)
