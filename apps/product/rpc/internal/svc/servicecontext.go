package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"lebron/apps/product/model"
	"lebron/apps/product/rpc/internal/config"
)

type ServiceContext struct {
	Config       config.Config
	ProductModel model.ProductModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DataSource)
	return &ServiceContext{
		Config:       c,
		ProductModel: model.NewProductModel(conn),
	}
}
