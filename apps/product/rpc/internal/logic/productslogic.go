package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mr"
	"lebron/apps/product/model"
	"lebron/apps/product/rpc/internal/svc"
	"lebron/apps/product/rpc/product"
)

type ProductsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductsLogic {
	return &ProductsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductsLogic) Products(in *product.ProductsRequest) (*product.ProductsResponse, error) {
	products := make([]*product.ProductItem, 0)
	ps, err := mr.MapReduce[any, any, any](func(source chan<- any) {
		for _, pid := range in.ProductIds {
			source <- pid
		}
	}, func(item any, writer mr.Writer[any], cancel func(error)) {
		pid := item.(int64)
		p, err := l.svcCtx.ProductModel.FindOne(l.ctx, pid)
		if err != nil {
			cancel(err)
			return
		}
		writer.Write(p)
	}, func(pipe <-chan interface{}, writer mr.Writer[interface{}], cancel func(error)) {
		var r []*model.Product
		for p := range pipe {
			r = append(r, p.(*model.Product))
		}
		writer.Write(r)
	})
	if err != nil {
		return nil, err
	}
	for _, p := range ps.([]*model.Product) {
		products = append(products, &product.ProductItem{
			ProductId: p.Id,
			Name:      p.Name,
		})
	}
	return &product.ProductsResponse{Products: products}, nil
}
