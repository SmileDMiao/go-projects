package logic

import (
	"context"

	"bookstore/services/user/internal/svc"
	"bookstore/services/user/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type UpdateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserLogic) UpdateUser(in *user.UpdateUserReq) (*user.UserResp, error) {
	// todo: add your logic here and delete this line

	return &user.UserResp{}, nil
}
