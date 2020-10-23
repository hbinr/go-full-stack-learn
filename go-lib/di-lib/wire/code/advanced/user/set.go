package user

import (
	"github.com/google/wire"
	"hb.study/go-lib/di-lib/wire/code/advanced/user/dao"
	"hb.study/go-lib/di-lib/wire/code/advanced/user/service"
)

var Set = wire.NewSet(dao.NewUserDaoSet, service.UserServiceSet)
