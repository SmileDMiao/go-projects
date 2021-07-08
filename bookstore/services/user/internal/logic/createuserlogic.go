package logic

import (
	"context"

	"bookstore/services/user/internal/svc"
	"bookstore/services/user/model"
	"bookstore/services/user/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type CreateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUserLogic) CreateUser(in *user.CreateUserReq) (*user.UserResp, error) {
	u := model.User{
		Username: in.Name,
		Password: in.Password,
	}
	result := l.svcCtx.DbEngin.Create(&u)

	if result.Error != nil {
		return &user.UserResp{Ok: false}, nil
	}

	return &user.UserResp{Ok: true}, nil
}
