package service

import (
	"errors"
	"gin_template/internal/app/domain/service/dto"
	"gin_template/internal/app/utils"
	"gorm.io/gorm"
)

func NewUser() UserService {
	return UserService{}
}

type UserService struct{}

/**
  注册
*/
func (s UserService) Register(params dto.RegisterParams) (success bool, err error) {
	user, err := userRepository.FindOne("username = ?", params.Username)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return false, errors.New("用户名已注册")
	}

	roles, err := roleRepository.Find("id in ?", params.RoleIds)
	if err != nil {
		return false, errors.New("不存在该角色")
	}

	user.Username = params.Username
	user.Mobile = params.Mobile
	user.Email = params.Email
	user.Roles = roles
	user.SetPssword(params.Password)
	success, err = userRepository.Create(user)

	return success, err
}

// 登录
func (s UserService) Login(params dto.LoginParams) (token string, err error) {
	user, err := userRepository.FindOne("username = ?", params.Username)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return "", errors.New("用户名或密码错误")
	}
	// 校验密码
	if isRight := user.CheckPassword(params.Password); !isRight {
		return "", errors.New("用户名或密码错误")
	}

	// 获取角色id
	roleIds := []string{}
	for _, v := range user.Roles {
		roleIds = append(roleIds, v.ID)
	}

	// 生成 token
	return utils.GenerateToken(user.ID, user.Username, roleIds, user.IsAdmin)
}
