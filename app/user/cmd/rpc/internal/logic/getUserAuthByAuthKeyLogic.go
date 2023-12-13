package logic

import (
	"context"

	"go-zero/app/user/cmd/rpc/internal/svc"
	"go-zero/app/user/cmd/rpc/pd"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserAuthByAuthKeyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserAuthByAuthKeyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserAuthByAuthKeyLogic {
	return &GetUserAuthByAuthKeyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserAuthByAuthKeyLogic) GetUserAuthByAuthKey(in *pd.GetUserAuthByAuthKeyReq) (*pd.GetUserAuthByAuthKeyResp, error) {
	// todo: add your logic here and delete this line

	return &pd.GetUserAuthByAuthKeyResp{}, nil
}
