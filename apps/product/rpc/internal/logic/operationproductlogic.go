package logic

import (
	"context"

	"lebron/apps/product/rpc/internal/svc"
	"lebron/apps/product/rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
)

const (
	validStatus          = 1
	operationProductsKey = "operation#products"
)

type OperationProductLogic struct {
	ctx              context.Context
	svcCtx           *svc.ServiceContext
	productListLogic *ProductListLogic
	logx.Logger
}

func NewOperationProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OperationProductLogic {
	return &OperationProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OperationProductLogic) OperationProduct(in *product.OperationProductRequest) (*product.OperationProductResponse, error) {
	opProducts, ok := l.svcCtx.LocalCache.Get(operationProductsKey)
	if ok {
		return &product.OperationProductResponse{Products: opProducts.([]*product.ProductItem)}, nil
	}
	//没有命中本地缓存
	operationProducts, err := l.svcCtx.OperationModel.OperationProducts(l.ctx, validStatus)
	if err != nil {
		return nil, err
	}
	pIds := make([]int64, 0)
	for _, item := range operationProducts {
		pIds = append(pIds, item.ProductId)
	}
	products, err := l.productListLogic.productsByIds(l.ctx, pIds)
	if err != nil {
		return nil, err
	}
	pItems := make([]*product.ProductItem, 0)
	for _, item := range products {
		pItems = append(pItems, &product.ProductItem{
			ProductId: item.Id,
			Name:      item.Name,
		})
	}
	return &product.OperationProductResponse{Products: pItems}, nil
}
