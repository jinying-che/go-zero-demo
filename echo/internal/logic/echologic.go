package logic

import (
	"context"

	"echo/internal/svc"
	"echo/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EchoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEchoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EchoLogic {
	return &EchoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EchoLogic) Echo(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	l.Logger.Info(req)
	ret := &types.Response{
		Message: req.Name,
	}
	return ret, nil
}
