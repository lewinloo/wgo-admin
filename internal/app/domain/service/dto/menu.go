package dto

type CreateMenuParams struct {
	Name              string `json:"name" binding:"required"`
	Title             string `json:"title" binding:"required"`
	ComponentFilePath string `json:"componentFilePath" binding:"required"`
	Path              string `json:"path" binding:"required"`
	Icon              string `json:"icon"`
	Redirect          string `json:"redirect"`
	Sort              uint   `json:"sort"`
	Enable            bool   `json:"enable"`
	Hidden            bool   `json:"hidden"`
	KeepAlive         bool   `json:"keepAlive"`
	Breadcrumb        bool   `json:"breadcrumb"`
	ParentId          uint   `json:"parentId"`
}
