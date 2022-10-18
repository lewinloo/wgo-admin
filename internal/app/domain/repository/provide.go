package repository

import "github.com/google/wire"

var ProvideSet = wire.NewSet(UserSet, RoleSet)
