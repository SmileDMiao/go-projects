package logic

import (
	"context"

	"bookstore/api/internal/svc"
	"bookstore/api/internal/types"
	"bookstore/services/user/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type CreateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) CreateUserLogic {
	return CreateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateUserLogic) CreateUser(req types.CreateUserReq) (*types.UserResp, error) {
	resp, err := l.svcCtx.User.CreateUser(l.ctx, &user.CreateUserReq{
		Name:     req.Name,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	return &types.UserResp{
		Ok: resp.Ok,
	}, nil

}
