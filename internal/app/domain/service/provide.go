package service

import "github.com/google/wire"

var ProvideSet = wire.NewSet(UserSet, CasbinSet)
