package dto

type RegisterParams struct {
	Username string   `json:"username" binding:"required,min=4,max=16"`
	Mobile   string   `json:"mobile" binding:"required"`
	Email    string   `json:"email" binding:"required"`
	Password string   `json:"password" binding:"required,min=6,max=32"`
	RoleIds  []string `json:"roleIds" binding:"required"`
}
