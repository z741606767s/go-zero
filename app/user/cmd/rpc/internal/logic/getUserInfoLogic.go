package logic

import (
	"context"

	"go-zero/app/user/cmd/rpc/internal/svc"
	"go-zero/app/user/cmd/rpc/pd"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *pd.GetUserInfoReq) (*pd.GetUserInfoResp, error) {
	// todo: add your logic here and delete this line

	return &pd.GetUserInfoResp{}, nil
}
