package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero/app/user/cmd/rpc/internal/config"
	"go-zero/app/user/model"
)

type ServiceContext struct {
	Config      config.Config
	RedisClient *redis.Redis

	UserModel     model.UserModel
	UserAuthModel model.UserAuthModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config:      c,
		RedisClient: redis.MustNewRedis(c.Redis.RedisConf),

		UserAuthModel: model.NewUserAuthModel(sqlConn),
		UserModel:     model.NewUserModel(sqlConn),
	}
}
