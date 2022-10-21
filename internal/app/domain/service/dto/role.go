package dto

type CreateRoleParams struct {
  ID     string `json:"id" binding:"required"`
  Name   string `json:"name" binding:"required"`
  Desc   string `json:"desc" binding:"required"`
  Status bool   `json:"status" binding:"required"`
}

type QueryRoleListParams struct {
  PageParams
}
