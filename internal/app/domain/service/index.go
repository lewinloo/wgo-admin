package service

import "gin_template/internal/app/domain/repository"

var (
	userRepository = repository.NewUser()
	roleRepository = repository.NewRole()
)
