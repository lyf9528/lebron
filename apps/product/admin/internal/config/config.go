package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	OssConfig OssConfig
}

type OssConfig struct {
	OSSEndpoint     string
	AccessKeyID     string
	AccessKeySecret string
}
