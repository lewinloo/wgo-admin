package dto

type RegisterParams struct {
	Username string   `json:"username" binding:"required,min=4,max=16"`
	Mobile   string   `json:"mobile" binding:"required"`
	Email    string   `json:"email" binding:"required"`
	Password string   `json:"password" binding:"required,min=6,max=32"`
	RoleIds  []string `json:"roleIds" binding:"required"`
}

type LoginParams struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type QueryUserListParams struct {
	PageParams
	Username string `json:"username"` // 用户名
	Email    string `json:"email"`    // 邮箱
	Status   uint   `json:"status"`   // 1：正常 2: 禁用
}
