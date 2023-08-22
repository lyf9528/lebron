package logic

import (
	"context"
	"lebron/apps/product/admin/internal/svc"
	"lebron/apps/product/admin/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadImageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

const (
	imageFileName = "image"
	bucketName    = "lebron-mall"
)

func NewUploadImageLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *UploadImageLogic {
	return &UploadImageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *UploadImageLogic) UploadImage() (resp *types.UploadImageResponse, err error) {
	// 获取上传的文件
	file, fileHeader, err := l.r.FormFile("file")
	if err != nil {
		return
	}
	if err != nil {
		return nil, err
	}
	defer file.Close()
	bucket, err := l.svcCtx.OssClient.Bucket(bucketName)
	if err = bucket.PutObject(fileHeader.Filename, file); err != nil {
		return nil, err
	}
	return &types.UploadImageResponse{Success: true}, nil
}
