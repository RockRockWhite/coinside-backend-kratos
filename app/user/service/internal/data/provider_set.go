package data

import (
	"github.com/google/wire"
	"github.com/ljxsteam/coinside-backend-kratos/app/user/service/config"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(config.NewConfig, NewDB, NewUserRepoNoCache)
