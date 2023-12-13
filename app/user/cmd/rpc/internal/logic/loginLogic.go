package logic

import (
	"context"

	"go-zero/app/user/cmd/rpc/internal/svc"
	"go-zero/app/user/cmd/rpc/pd"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *pd.LoginReq) (*pd.LoginResp, error) {
	// todo: add your logic here and delete this line

	return &pd.LoginResp{}, nil
}
