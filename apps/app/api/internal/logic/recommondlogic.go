package logic

import (
	"context"

	"lebron/apps/app/api/internal/svc"
	"lebron/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RecommondLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRecommondLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RecommondLogic {
	return &RecommondLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RecommondLogic) Recommond(req *types.RecommondRequest) (resp *types.RecommondResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
