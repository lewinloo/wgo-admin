package service

import (
	"errors"
	"gin_template/internal/app/domain/repository"
	"gin_template/internal/app/domain/service/dto"
	"gin_template/internal/app/domain/service/vo"
	"gin_template/internal/app/utils"
	"github.com/google/wire"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var UserSet = wire.NewSet(wire.Struct(new(UserService), "*"))

type UserService struct {
	UserRepo *repository.UserRepository
	RoleRepo *repository.RoleRepository
	Log      *zap.Logger
}

/**
  注册
*/
func (s UserService) Register(params dto.RegisterParams) (success bool, err error) {
	user, err := s.UserRepo.FindOne("username = ?", params.Username)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return false, errors.New("用户名已注册")
	}

	roles, err := s.RoleRepo.Find("id in ?", params.RoleIds)
	if err != nil {
		return false, errors.New("不存在该角色")
	}

	user.Username = params.Username
	user.Mobile = params.Mobile
	user.Email = params.Email
	user.Roles = roles
	_ = user.SetPssword(params.Password)
	success, err = s.UserRepo.Create(user)

	return success, err
}

// 登录
func (s UserService) Login(params dto.LoginParams) (token string, err error) {
	user, err := s.UserRepo.FindOne("username = ?", params.Username)
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

// 查询用户列表
func (s UserService) GetList(params dto.QueryUserListParams) (vo.ListResult, error) {
	return s.UserRepo.FindList(params)
}

// 删除用户
func (s UserService) DeleteUsers(params dto.Ids) (success bool) {
	success, err := s.UserRepo.DeleteByIds(params.Ids)
	if !success {
		s.Log.Error("删除用户失败：" + err.Error())
		return false
	}
	return true
}
