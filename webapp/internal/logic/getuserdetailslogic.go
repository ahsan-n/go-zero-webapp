package logic

import (
	"context"

	"bazaartechnologies/rating-service/go-zero-webapp/webapp/internal/svc"
	"bazaartechnologies/rating-service/go-zero-webapp/webapp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserDetailsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserDetailsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserDetailsLogic {
	return &GetUserDetailsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserDetailsLogic) GetUserDetails() (resp *types.UserDetails, err error) {
	// todo: add your logic here and delete this line

	return
}
