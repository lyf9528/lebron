package logic

import (
	"context"

	"lebron/apps/product/rpc/internal/svc"
	"lebron/apps/product/rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductLogic {
	return &ProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductLogic) Product(in *product.ProductItemRequest) (*product.ProductItem, error) {
	one, err := l.svcCtx.ProductModel.FindOne(l.ctx, in.GetProductId())
	if err != nil {
		return nil, err
	}
	return &product.ProductItem{
		ProductId:   one.Id,
		Name:        one.Name,
		Description: one.Detail.String,
		ImageUrl:    one.Images.String,
	}, nil
}
