package logic

import (
	"context"

	"lebron/apps/app/api/internal/svc"
	"lebron/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductCommontLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductCommontLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductCommontLogic {
	return &ProductCommontLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductCommontLogic) ProductCommont(req *types.ProductCommentRequest) (resp *types.ProductCommentResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
