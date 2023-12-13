package logic

import (
	"context"
	"github.com/pkg/errors"
	"go-zero/app/user/model"
	"go-zero/common/xerr"

	"go-zero/app/user/cmd/rpc/internal/svc"
	"go-zero/app/user/cmd/rpc/pd"

	"github.com/zeromicro/go-zero/core/logx"
)

var ErrUserAlreadyRegisterError = xerr.NewErrMsg("user has been registered")

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *pd.RegisterReq) (*pd.RegisterResp, error) {
	// todo: add your logic here and delete this line

	user, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, in.Mobile)
	if err != nil && !errors.Is(model.ErrNotFound, err) {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "mobile:%s,err:%v", in.Mobile, err)
	}
	if user != nil {
		return nil, errors.Wrapf(ErrUserAlreadyRegisterError, "Register user exists mobile:%s,err:%v", in.Mobile, err)
	}

	// todo: add user logic ...

	return &pd.RegisterResp{}, nil
}
