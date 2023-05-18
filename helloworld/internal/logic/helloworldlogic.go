package logic

import (
	"context"
	"time"

	"go-zero-demo-standalone/helloworld/internal/svc"
	"go-zero-demo-standalone/helloworld/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HelloworldLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHelloworldLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HelloworldLogic {
	return &HelloworldLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HelloworldLogic) Helloworld(req *types.Request) (resp *types.Response, err error) {
	return &types.Response{
		Message: "hello! You input: " + req.Name + ", current time: " + time.DateTime,
	}, nil
}
