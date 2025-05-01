package logic

import (
	"context"

	"template/gozero/src/internal/svc"
	"template/gozero/src/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SrcLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSrcLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SrcLogic {
	return &SrcLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SrcLogic) Src(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
