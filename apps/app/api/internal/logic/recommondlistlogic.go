package logic

import (
	"context"

	"lebron/apps/app/api/internal/svc"
	"lebron/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RecommondListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRecommondListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RecommondListLogic {
	return &RecommondListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RecommondListLogic) RecommondList(req *types.RecommondRequest) (resp *types.RecommondResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
