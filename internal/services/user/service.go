package userService

import (
	"errors"
	"gin_template/internal/dto"
	"gin_template/internal/repository"

	"gorm.io/gorm"
)

var (
	userRepository = repository.NewUser()
	roleRepository = repository.NewRole()
)

func Register(params dto.RegisterParams) (success bool, err error) {
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
