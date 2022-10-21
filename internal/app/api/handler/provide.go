package handler

import (
	"github.com/google/wire"
)

var ProvideSet = wire.NewSet(UserSet, HelloSet, MenuSet, RoleSet)
