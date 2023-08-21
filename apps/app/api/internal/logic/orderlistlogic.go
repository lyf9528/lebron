package logic

import (
	"context"
	"lebron/apps/order/rpc/order"
	"lebron/apps/product/rpc/product"
	"strconv"

	"lebron/apps/app/api/internal/svc"
	"lebron/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderListLogic {
	return &OrderListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderListLogic) OrderList(req *types.OrderListRequest) (resp *types.OrderListResponse, err error) {
	ordersRes, err := l.svcCtx.OrderRPC.Orders(l.ctx, &order.OrdersRequest{
		UserId: req.UID,
		Status: req.Status,
		Cursor: req.Cursor,
		Ps:     int32(req.Ps),
	})
	if err != nil {
		return nil, err
	}
	var pids []string
	for _, o := range ordersRes.Orders {
		pids = append(pids, strconv.FormatInt(o.ProductId, 10))
	}
	productsRes, err := l.svcCtx.ProductRPC.Products(l.ctx, &product.ProductsRequest{ProductIds: pids})
	if err != nil {
		return nil, err
	}
	products := make(map[int64]*product.ProductItem, 0)
	for _, p := range productsRes.Products {
		products[p.ProductId] = p
	}
	orders := make([]*types.Order, 0)
	for _, o := range ordersRes.Orders {
		orders = append(orders, &types.Order{
			OrderID:            o.OrderId,
			Status:             o.Status,
			Quantity:           o.Quantity,
			Payment:            float64(o.Payment),
			TotalPrice:         0,
			CreateTime:         o.CreateTime,
			ProductID:          products[o.ProductId].ProductId,
			ProductName:        products[o.ProductId].Name,
			ProductImage:       products[o.ProductId].ImageUrl,
			ProductDescription: products[o.ProductId].Description,
		})
	}
	return &types.OrderListResponse{
		Orders:    orders,
		IsEnd:     false,
		OrderTime: 0,
	}, nil
}
