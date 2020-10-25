package user

import (
	"github.com/google/wire"
	"hb.study/go-lib/di-lib/wire/code/advanced/internal/app/user/dao"
	"hb.study/go-lib/di-lib/wire/code/advanced/internal/app/user/service"
)

var Set = wire.NewSet(dao.UserDaoSet, service.UserServiceSet)
