package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"lebron/apps/app/api/internal/config"
	"lebron/apps/order/rpc/orderclient"
	"lebron/apps/product/rpc/productclient"
)

type ServiceContext struct {
	Config     config.Config
	OrderRPC   orderclient.Order
	ProductRPC productclient.Product
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		OrderRPC:   orderclient.NewOrder(zrpc.MustNewClient(c.OrderRPC)),
		ProductRPC: productclient.NewProduct(zrpc.MustNewClient(c.ProductRPC)),
	}
}
