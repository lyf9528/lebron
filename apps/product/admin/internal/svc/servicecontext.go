package svc

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"lebron/apps/product/admin/internal/config"
)

type ServiceContext struct {
	Config    config.Config
	OssClient *oss.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	oc, err := oss.New(c.OssConfig.OSSEndpoint, c.OssConfig.AccessKeyID, c.OssConfig.AccessKeySecret)
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config:    c,
		OssClient: oc,
	}
}
