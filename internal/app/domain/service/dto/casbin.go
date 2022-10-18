package dto

type PermissionInfo struct {
	Path   string `json:"path"`   // 路径
	Method string `json:"method"` // 方法
}

type PermissionInReceive struct {
	RoleId          string           `json:"roleId"` // 权限id
	PermissionInfos []PermissionInfo `json:"permissionInfos"`
}

func DefaultCasbin() []PermissionInfo {
	return []PermissionInfo{
		{Path: "/user/login", Method: "POST"},
		{Path: "/user/register", Method: "POST"},
		// {Path: "/menu/getMenu", Method: "POST"},
		// {Path: "/user/changePassword", Method: "POST"},
		// {Path: "/user/setUserInfo", Method: "PUT"},
		// {Path: "/user/getUserInfo", Method: "GET"},
	}
}
