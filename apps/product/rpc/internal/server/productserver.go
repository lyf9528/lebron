// Code generated by goctl. DO NOT EDIT.
// Source: product.proto

package server

import (
	"context"

	"lebron/apps/product/rpc/internal/logic"
	"lebron/apps/product/rpc/internal/svc"
	"lebron/apps/product/rpc/product"
)

type ProductServer struct {
	svcCtx *svc.ServiceContext
	product.UnimplementedProductServer
}

func NewProductServer(svcCtx *svc.ServiceContext) *ProductServer {
	return &ProductServer{
		svcCtx: svcCtx,
	}
}

func (s *ProductServer) Products(ctx context.Context, in *product.ProductsRequest) (*product.ProductsResponse, error) {
	l := logic.NewProductsLogic(ctx, s.svcCtx)
	return l.Products(in)
}

func (s *ProductServer) Product(ctx context.Context, in *product.ProductItemRequest) (*product.ProductItem, error) {
	l := logic.NewProductLogic(ctx, s.svcCtx)
	return l.Product(in)
}

func (s *ProductServer) ProductList(ctx context.Context, in *product.ProductListRequest) (*product.ProductListResponse, error) {
	l := logic.NewProductListLogic(ctx, s.svcCtx)
	return l.ProductList(in)
}

func (s *ProductServer) OperationProduct(ctx context.Context, in *product.OperationProductRequest) (*product.OperationProductResponse, error) {
	l := logic.NewOperationProductLogic(ctx, s.svcCtx)
	return l.OperationProduct(in)
}
